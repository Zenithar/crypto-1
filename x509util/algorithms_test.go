package x509util

import (
	"crypto/x509"
	"reflect"
	"testing"
)

func TestSignatureAlgorithm_Set(t *testing.T) {
	type args struct {
		c *x509.Certificate
	}
	tests := []struct {
		name string
		s    SignatureAlgorithm
		args args
		want *x509.Certificate
	}{
		{"ok", SignatureAlgorithm(x509.ECDSAWithSHA256), args{&x509.Certificate{}}, &x509.Certificate{SignatureAlgorithm: x509.ECDSAWithSHA256}},
		{"ok", SignatureAlgorithm(x509.PureEd25519), args{&x509.Certificate{}}, &x509.Certificate{SignatureAlgorithm: x509.PureEd25519}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Set(tt.args.c)
			if !reflect.DeepEqual(tt.args.c, tt.want) {
				t.Errorf("SignatureAlgorithm.Set() = %v, want %v", tt.args.c, tt.want)
			}
		})
	}
}

func TestSignatureAlgorithm_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		s       SignatureAlgorithm
		want    []byte
		wantErr bool
	}{
		{"UnknownSignatureAlgorithm", SignatureAlgorithm(x509.UnknownSignatureAlgorithm), []byte(`""`), false},
		{"MD2WithRSA", SignatureAlgorithm(x509.MD2WithRSA), []byte(`"` + MD2WithRSA + `"`), false},
		{"MD5WithRSA", SignatureAlgorithm(x509.MD5WithRSA), []byte(`"` + MD5WithRSA + `"`), false},
		{"SHA1WithRSA", SignatureAlgorithm(x509.SHA1WithRSA), []byte(`"` + SHA1WithRSA + `"`), false},
		{"SHA256WithRSA", SignatureAlgorithm(x509.SHA256WithRSA), []byte(`"` + SHA256WithRSA + `"`), false},
		{"SHA384WithRSA", SignatureAlgorithm(x509.SHA384WithRSA), []byte(`"` + SHA384WithRSA + `"`), false},
		{"SHA512WithRSA", SignatureAlgorithm(x509.SHA512WithRSA), []byte(`"` + SHA512WithRSA + `"`), false},
		{"DSAWithSHA1", SignatureAlgorithm(x509.DSAWithSHA1), []byte(`"` + DSAWithSHA1 + `"`), false},
		{"DSAWithSHA256", SignatureAlgorithm(x509.DSAWithSHA256), []byte(`"` + DSAWithSHA256 + `"`), false},
		{"ECDSAWithSHA1", SignatureAlgorithm(x509.ECDSAWithSHA1), []byte(`"` + ECDSAWithSHA1 + `"`), false},
		{"ECDSAWithSHA256", SignatureAlgorithm(x509.ECDSAWithSHA256), []byte(`"` + ECDSAWithSHA256 + `"`), false},
		{"ECDSAWithSHA384", SignatureAlgorithm(x509.ECDSAWithSHA384), []byte(`"` + ECDSAWithSHA384 + `"`), false},
		{"ECDSAWithSHA512", SignatureAlgorithm(x509.ECDSAWithSHA512), []byte(`"` + ECDSAWithSHA512 + `"`), false},
		{"SHA256WithRSAPSS", SignatureAlgorithm(x509.SHA256WithRSAPSS), []byte(`"` + SHA256WithRSAPSS + `"`), false},
		{"SHA384WithRSAPSS", SignatureAlgorithm(x509.SHA384WithRSAPSS), []byte(`"` + SHA384WithRSAPSS + `"`), false},
		{"SHA512WithRSAPSS", SignatureAlgorithm(x509.SHA512WithRSAPSS), []byte(`"` + SHA512WithRSAPSS + `"`), false},
		{"PureEd25519", SignatureAlgorithm(x509.PureEd25519), []byte(`"` + PureEd25519 + `"`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SignatureAlgorithm.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignatureAlgorithm.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestSignatureAlgorithm_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    SignatureAlgorithm
		wantErr bool
	}{
		{"MD2-RSA", args{[]byte(`"MD2-RSA"`)}, SignatureAlgorithm(x509.MD2WithRSA), false},
		{"MD5-RSA", args{[]byte(`"MD5-RSA"`)}, SignatureAlgorithm(x509.MD5WithRSA), false},
		{"SHA1-RSA", args{[]byte(`"SHA1-RSA"`)}, SignatureAlgorithm(x509.SHA1WithRSA), false},
		{"SHA256-RSA", args{[]byte(`"SHA256-RSA"`)}, SignatureAlgorithm(x509.SHA256WithRSA), false},
		{"SHA384-RSA", args{[]byte(`"SHA384-RSA"`)}, SignatureAlgorithm(x509.SHA384WithRSA), false},
		{"SHA512-RSA", args{[]byte(`"SHA512-RSA"`)}, SignatureAlgorithm(x509.SHA512WithRSA), false},
		{"SHA256-RSAPSS", args{[]byte(`"SHA256-RSAPSS"`)}, SignatureAlgorithm(x509.SHA256WithRSAPSS), false},
		{"SHA384-RSAPSS", args{[]byte(`"SHA384-RSAPSS"`)}, SignatureAlgorithm(x509.SHA384WithRSAPSS), false},
		{"SHA512-RSAPSS", args{[]byte(`"SHA512-RSAPSS"`)}, SignatureAlgorithm(x509.SHA512WithRSAPSS), false},
		{"DSA-SHA1", args{[]byte(`"DSA-SHA1"`)}, SignatureAlgorithm(x509.DSAWithSHA1), false},
		{"DSA-SHA256", args{[]byte(`"DSA-SHA256"`)}, SignatureAlgorithm(x509.DSAWithSHA256), false},
		{"ECDSA-SHA1", args{[]byte(`"ECDSA-SHA1"`)}, SignatureAlgorithm(x509.ECDSAWithSHA1), false},
		{"ECDSA-SHA256", args{[]byte(`"ECDSA-SHA256"`)}, SignatureAlgorithm(x509.ECDSAWithSHA256), false},
		{"ECDSA-SHA384", args{[]byte(`"ECDSA-SHA384"`)}, SignatureAlgorithm(x509.ECDSAWithSHA384), false},
		{"ECDSA-SHA512", args{[]byte(`"ECDSA-SHA512"`)}, SignatureAlgorithm(x509.ECDSAWithSHA512), false},
		{"Ed25519", args{[]byte(`"Ed25519"`)}, SignatureAlgorithm(x509.PureEd25519), false},
		{"lowercase", args{[]byte(`"ecdsa-sha256"`)}, SignatureAlgorithm(x509.ECDSAWithSHA256), false},
		{"empty", args{[]byte(`""`)}, SignatureAlgorithm(0), false},
		{"null", args{[]byte(`null`)}, SignatureAlgorithm(0), false},
		{"unknown", args{[]byte(`"unknown"`)}, SignatureAlgorithm(0), true},
		{"number", args{[]byte(`0`)}, SignatureAlgorithm(0), true},
		{"numberString", args{[]byte(`"0"`)}, SignatureAlgorithm(0), true},
		{"object", args{[]byte(`{}`)}, SignatureAlgorithm(0), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got SignatureAlgorithm
			if err := got.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SignatureAlgorithm.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignatureAlgorithm.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
