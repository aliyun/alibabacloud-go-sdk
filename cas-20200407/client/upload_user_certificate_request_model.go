// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadUserCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *UploadUserCertificateRequest
	GetCert() *string
	SetEncryptCert(v string) *UploadUserCertificateRequest
	GetEncryptCert() *string
	SetEncryptPrivateKey(v string) *UploadUserCertificateRequest
	GetEncryptPrivateKey() *string
	SetKey(v string) *UploadUserCertificateRequest
	GetKey() *string
	SetName(v string) *UploadUserCertificateRequest
	GetName() *string
	SetResourceGroupId(v string) *UploadUserCertificateRequest
	GetResourceGroupId() *string
	SetSignCert(v string) *UploadUserCertificateRequest
	GetSignCert() *string
	SetSignPrivateKey(v string) *UploadUserCertificateRequest
	GetSignPrivateKey() *string
	SetTags(v []*UploadUserCertificateRequestTags) *UploadUserCertificateRequest
	GetTags() []*UploadUserCertificateRequestTags
}

type UploadUserCertificateRequest struct {
	// The content of the certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIF...... -----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The content of the encryption certificate in PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCA***
	//
	// -----END CERTIFICATE-----
	EncryptCert *string `json:"EncryptCert,omitempty" xml:"EncryptCert,omitempty"`
	// The private key of the encryption certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN EC PRIVATE KEY-----
	//
	// MHcCAQEEI****
	//
	// -----END EC PRIVATE KEY-----
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	// The private key of the certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCAbagAw****
	//
	// -----END CERTIFICATE-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the certificate. The name can be up to 64 characters in length, and can contain all types of characters, such as letters, digits, and underscores (_).
	//
	// >  The name must be unique within an Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// cert-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// the resource group id.
	//
	// example:
	//
	// rg-ae****vty
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The content of the signing certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCAbagAw****
	//
	// -----END CERTIFICATE-----
	SignCert *string `json:"SignCert,omitempty" xml:"SignCert,omitempty"`
	// The private key of the signing certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN EC PRIVATE KEY-----
	//
	// MHcCAQEEILR****
	//
	// -----END EC PRIVATE KEY-----
	SignPrivateKey *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
	// The tags.
	Tags []*UploadUserCertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UploadUserCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadUserCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadUserCertificateRequest) GetCert() *string {
	return s.Cert
}

func (s *UploadUserCertificateRequest) GetEncryptCert() *string {
	return s.EncryptCert
}

func (s *UploadUserCertificateRequest) GetEncryptPrivateKey() *string {
	return s.EncryptPrivateKey
}

func (s *UploadUserCertificateRequest) GetKey() *string {
	return s.Key
}

func (s *UploadUserCertificateRequest) GetName() *string {
	return s.Name
}

func (s *UploadUserCertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UploadUserCertificateRequest) GetSignCert() *string {
	return s.SignCert
}

func (s *UploadUserCertificateRequest) GetSignPrivateKey() *string {
	return s.SignPrivateKey
}

func (s *UploadUserCertificateRequest) GetTags() []*UploadUserCertificateRequestTags {
	return s.Tags
}

func (s *UploadUserCertificateRequest) SetCert(v string) *UploadUserCertificateRequest {
	s.Cert = &v
	return s
}

func (s *UploadUserCertificateRequest) SetEncryptCert(v string) *UploadUserCertificateRequest {
	s.EncryptCert = &v
	return s
}

func (s *UploadUserCertificateRequest) SetEncryptPrivateKey(v string) *UploadUserCertificateRequest {
	s.EncryptPrivateKey = &v
	return s
}

func (s *UploadUserCertificateRequest) SetKey(v string) *UploadUserCertificateRequest {
	s.Key = &v
	return s
}

func (s *UploadUserCertificateRequest) SetName(v string) *UploadUserCertificateRequest {
	s.Name = &v
	return s
}

func (s *UploadUserCertificateRequest) SetResourceGroupId(v string) *UploadUserCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadUserCertificateRequest) SetSignCert(v string) *UploadUserCertificateRequest {
	s.SignCert = &v
	return s
}

func (s *UploadUserCertificateRequest) SetSignPrivateKey(v string) *UploadUserCertificateRequest {
	s.SignPrivateKey = &v
	return s
}

func (s *UploadUserCertificateRequest) SetTags(v []*UploadUserCertificateRequestTags) *UploadUserCertificateRequest {
	s.Tags = v
	return s
}

func (s *UploadUserCertificateRequest) Validate() error {
	return dara.Validate(s)
}

type UploadUserCertificateRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UploadUserCertificateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UploadUserCertificateRequestTags) GoString() string {
	return s.String()
}

func (s *UploadUserCertificateRequestTags) GetKey() *string {
	return s.Key
}

func (s *UploadUserCertificateRequestTags) GetValue() *string {
	return s.Value
}

func (s *UploadUserCertificateRequestTags) SetKey(v string) *UploadUserCertificateRequestTags {
	s.Key = &v
	return s
}

func (s *UploadUserCertificateRequestTags) SetValue(v string) *UploadUserCertificateRequestTags {
	s.Value = &v
	return s
}

func (s *UploadUserCertificateRequestTags) Validate() error {
	return dara.Validate(s)
}
