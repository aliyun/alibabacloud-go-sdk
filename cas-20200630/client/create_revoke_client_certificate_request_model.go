// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRevokeClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *CreateRevokeClientCertificateRequest
	GetIdentifier() *string
}

type CreateRevokeClientCertificateRequest struct {
	// The unique identifier of the client or server certificate to revoke.
	//
	// > Call [ListClientCertificate](https://help.aliyun.com/document_detail/465990.html) to query the unique identifiers of all client and server certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s CreateRevokeClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRevokeClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateRevokeClientCertificateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateRevokeClientCertificateRequest) SetIdentifier(v string) *CreateRevokeClientCertificateRequest {
	s.Identifier = &v
	return s
}

func (s *CreateRevokeClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
