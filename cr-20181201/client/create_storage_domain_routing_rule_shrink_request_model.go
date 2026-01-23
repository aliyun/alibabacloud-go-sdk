// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageDomainRoutingRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateStorageDomainRoutingRuleShrinkRequest
	GetInstanceId() *string
	SetRoutesShrink(v string) *CreateStorageDomainRoutingRuleShrinkRequest
	GetRoutesShrink() *string
}

type CreateStorageDomainRoutingRuleShrinkRequest struct {
	// The instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The route list
	//
	// This parameter is required.
	RoutesShrink *string `json:"Routes,omitempty" xml:"Routes,omitempty"`
}

func (s CreateStorageDomainRoutingRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageDomainRoutingRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStorageDomainRoutingRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateStorageDomainRoutingRuleShrinkRequest) GetRoutesShrink() *string {
	return s.RoutesShrink
}

func (s *CreateStorageDomainRoutingRuleShrinkRequest) SetInstanceId(v string) *CreateStorageDomainRoutingRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateStorageDomainRoutingRuleShrinkRequest) SetRoutesShrink(v string) *CreateStorageDomainRoutingRuleShrinkRequest {
	s.RoutesShrink = &v
	return s
}

func (s *CreateStorageDomainRoutingRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
