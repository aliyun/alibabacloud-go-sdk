// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachGatewayDomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDomainShrink(v string) *AttachGatewayDomainShrinkRequest
	GetCustomDomainShrink() *string
}

type AttachGatewayDomainShrinkRequest struct {
	// The custom domain name information.
	//
	// This parameter is required.
	CustomDomainShrink *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
}

func (s AttachGatewayDomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachGatewayDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachGatewayDomainShrinkRequest) GetCustomDomainShrink() *string {
	return s.CustomDomainShrink
}

func (s *AttachGatewayDomainShrinkRequest) SetCustomDomainShrink(v string) *AttachGatewayDomainShrinkRequest {
	s.CustomDomainShrink = &v
	return s
}

func (s *AttachGatewayDomainShrinkRequest) Validate() error {
	return dara.Validate(s)
}
