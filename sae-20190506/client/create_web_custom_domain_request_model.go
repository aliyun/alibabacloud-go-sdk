// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebCustomDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *CreateWebCustomDomainRequest
	GetNamespaceId() *string
	SetBody(v *CreateWebCustomDomainInput) *CreateWebCustomDomainRequest
	GetBody() *CreateWebCustomDomainInput
}

type CreateWebCustomDomainRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The information about custom domain name.
	//
	// This parameter is required.
	Body *CreateWebCustomDomainInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebCustomDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateWebCustomDomainRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateWebCustomDomainRequest) GetBody() *CreateWebCustomDomainInput {
	return s.Body
}

func (s *CreateWebCustomDomainRequest) SetNamespaceId(v string) *CreateWebCustomDomainRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateWebCustomDomainRequest) SetBody(v *CreateWebCustomDomainInput) *CreateWebCustomDomainRequest {
	s.Body = v
	return s
}

func (s *CreateWebCustomDomainRequest) Validate() error {
	return dara.Validate(s)
}
