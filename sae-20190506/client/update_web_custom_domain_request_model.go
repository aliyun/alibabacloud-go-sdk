// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebCustomDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *UpdateWebCustomDomainRequest
	GetNamespaceId() *string
	SetBody(v *UpdateWebCustomDomainInput) *UpdateWebCustomDomainRequest
	GetBody() *UpdateWebCustomDomainInput
}

type UpdateWebCustomDomainRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The information about the custom domain name.
	//
	// This parameter is required.
	Body *UpdateWebCustomDomainInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebCustomDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebCustomDomainRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateWebCustomDomainRequest) GetBody() *UpdateWebCustomDomainInput {
	return s.Body
}

func (s *UpdateWebCustomDomainRequest) SetNamespaceId(v string) *UpdateWebCustomDomainRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateWebCustomDomainRequest) SetBody(v *UpdateWebCustomDomainInput) *UpdateWebCustomDomainRequest {
	s.Body = v
	return s
}

func (s *UpdateWebCustomDomainRequest) Validate() error {
	return dara.Validate(s)
}
