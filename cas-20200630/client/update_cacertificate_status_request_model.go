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
	// A client token used to ensure the idempotence of the request.
	//
	// Generate a unique parameter value from your client for each request. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 3838B684-3075-582B-9A45-8C99104029DF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique identifier of the CA certificate.
	//
	// > Call [DescribeCACertificateList](https://help.aliyun.com/document_detail/465957.html) to query the unique identifiers of all CA certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The operation to perform on the CA certificate. Set the value to **REVOKE**. This revokes the CA certificate and changes its status to **REVOKE**.
	//
	// > This operation is supported only when the CA certificate is in the **ISSUE*	- state. Call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to query the current status of the CA certificate.
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
