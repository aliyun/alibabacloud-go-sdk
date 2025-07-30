// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *DeleteClientCertificateRequest
	GetIdentifier() *string
}

type DeleteClientCertificateRequest struct {
	// The unique identifier of the client certificate or server certificate that you want to delete. The status of the certificate must be **REVOKE**.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers and status of all client certificates and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DeleteClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientCertificateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DeleteClientCertificateRequest) SetIdentifier(v string) *DeleteClientCertificateRequest {
	s.Identifier = &v
	return s
}

func (s *DeleteClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
