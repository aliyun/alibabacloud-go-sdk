// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGatewayDomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDomainShrink(v string) *DetachGatewayDomainShrinkRequest
	GetCustomDomainShrink() *string
}

type DetachGatewayDomainShrinkRequest struct {
	// The custom domain name information.
	//
	// This parameter is required.
	CustomDomainShrink *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
}

func (s DetachGatewayDomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachGatewayDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachGatewayDomainShrinkRequest) GetCustomDomainShrink() *string {
	return s.CustomDomainShrink
}

func (s *DetachGatewayDomainShrinkRequest) SetCustomDomainShrink(v string) *DetachGatewayDomainShrinkRequest {
	s.CustomDomainShrink = &v
	return s
}

func (s *DetachGatewayDomainShrinkRequest) Validate() error {
	return dara.Validate(s)
}
