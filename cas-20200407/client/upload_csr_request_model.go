// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsr(v string) *UploadCsrRequest
	GetCsr() *string
	SetKey(v string) *UploadCsrRequest
	GetKey() *string
	SetName(v string) *UploadCsrRequest
	GetName() *string
}

type UploadCsrRequest struct {
	// The content of the CSR.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The private key content of the certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MII.... -----END RSA PRIVATE KEY-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the CSR.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UploadCsrRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadCsrRequest) GoString() string {
	return s.String()
}

func (s *UploadCsrRequest) GetCsr() *string {
	return s.Csr
}

func (s *UploadCsrRequest) GetKey() *string {
	return s.Key
}

func (s *UploadCsrRequest) GetName() *string {
	return s.Name
}

func (s *UploadCsrRequest) SetCsr(v string) *UploadCsrRequest {
	s.Csr = &v
	return s
}

func (s *UploadCsrRequest) SetKey(v string) *UploadCsrRequest {
	s.Key = &v
	return s
}

func (s *UploadCsrRequest) SetName(v string) *UploadCsrRequest {
	s.Name = &v
	return s
}

func (s *UploadCsrRequest) Validate() error {
	return dara.Validate(s)
}
