// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPCACertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *UploadPCACertRequest
	GetCert() *string
	SetName(v string) *UploadPCACertRequest
	GetName() *string
	SetPrivateKey(v string) *UploadPCACertRequest
	GetPrivateKey() *string
	SetWarehouseId(v int64) *UploadPCACertRequest
	GetWarehouseId() *int64
}

type UploadPCACertRequest struct {
	// The content of the certificate and its chain, in PEM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIEJDCCAwygAwIBAgIQITRHItTLTQizTyd3K7AMRTANBgkqhkiG9w0BAQsFADBe ***************	- 5/akmr2GK/Y= -----END CERTIFICATE----- -----BEGIN CERTIFICATE----- MIIDuzCCAqOgAwIBAgIQSEIWDPfWTDKZcWNyL2O+fjANBgkqhkiG9w0BAQsFADBf ***************	- URUHyMW5Qd5Q9g6Y4sDOIm6It9TF7EjpwMs42R30agcRYzuUsN72ZFBYFJwnBX8= -----END CERTIFICATE----- -----BEGIN CERTIFICATE----- MIIDizCCAnOgAwIBAgIRAMfjPkDKfELTo07l3A3cUSYwDQYJKoZIhvcNAQELBQAw ********	- CjWTnYPhCcO2uIcnqMt7zCVs5LXBK/XSwlAXKMvKT0uuzw9VxeMfEabflKu0By8= -----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// A custom name for the certificate.
	//
	// example:
	//
	// cert_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private key of the certificate, in PEM format.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MIIEowIBAAKCAQEA5SIfpNCBoiDrZhX1H39CHwQMVD0kBNeBTWfP9xkeesvfzbOz ******	- POVNFfDf9h7pJtQ5fRZNTYTDs/d+cH62Z3+nS74mNnEfff0nkvne -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The ID of the certificate warehouse.
	//
	// > Call [ListCertWarehouse](https://help.aliyun.com/document_detail/455805.html) to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WarehouseId *int64 `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
}

func (s UploadPCACertRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadPCACertRequest) GoString() string {
	return s.String()
}

func (s *UploadPCACertRequest) GetCert() *string {
	return s.Cert
}

func (s *UploadPCACertRequest) GetName() *string {
	return s.Name
}

func (s *UploadPCACertRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *UploadPCACertRequest) GetWarehouseId() *int64 {
	return s.WarehouseId
}

func (s *UploadPCACertRequest) SetCert(v string) *UploadPCACertRequest {
	s.Cert = &v
	return s
}

func (s *UploadPCACertRequest) SetName(v string) *UploadPCACertRequest {
	s.Name = &v
	return s
}

func (s *UploadPCACertRequest) SetPrivateKey(v string) *UploadPCACertRequest {
	s.PrivateKey = &v
	return s
}

func (s *UploadPCACertRequest) SetWarehouseId(v int64) *UploadPCACertRequest {
	s.WarehouseId = &v
	return s
}

func (s *UploadPCACertRequest) Validate() error {
	return dara.Validate(s)
}
