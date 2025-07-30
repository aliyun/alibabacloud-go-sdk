// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCACertificateStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCACertificateStatusRequest
	GetClientToken() *string
	SetIdentifier(v string) *UpdateCACertificateStatusRequest
	GetIdentifier() *string
	SetStatus(v string) *UpdateCACertificateStatusRequest
	GetStatus() *string
}

type UpdateCACertificateStatusRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique identifier of the CA certificate whose status you want to change.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifiers of all CA certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The state to which you want to change the CA certificate. Set to the value to **REVOKE**. After this operation is called, the status of the CA certificate is changed to **REVOKE**.
	//
	// >  You can call this operation only if the status of a CA certificate is **ISSUE**. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the status of a CA certificate.
	//
	// example:
	//
	// REVOKE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateCACertificateStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCACertificateStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCACertificateStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCACertificateStatusRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateCACertificateStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateCACertificateStatusRequest) SetClientToken(v string) *UpdateCACertificateStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCACertificateStatusRequest) SetIdentifier(v string) *UpdateCACertificateStatusRequest {
	s.Identifier = &v
	return s
}

func (s *UpdateCACertificateStatusRequest) SetStatus(v string) *UpdateCACertificateStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateCACertificateStatusRequest) Validate() error {
	return dara.Validate(s)
}
