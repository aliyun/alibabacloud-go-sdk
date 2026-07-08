// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeWHClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *RevokeWHClientCertificateRequest
	GetIdentifier() *string
}

type RevokeWHClientCertificateRequest struct {
	// The unique identifier of the client or server certificate to revoke.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s RevokeWHClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeWHClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *RevokeWHClientCertificateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *RevokeWHClientCertificateRequest) SetIdentifier(v string) *RevokeWHClientCertificateRequest {
	s.Identifier = &v
	return s
}

func (s *RevokeWHClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
