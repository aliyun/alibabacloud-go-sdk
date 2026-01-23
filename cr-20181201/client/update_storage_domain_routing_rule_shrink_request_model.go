// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStorageDomainRoutingRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateStorageDomainRoutingRuleShrinkRequest
	GetInstanceId() *string
	SetRoutesShrink(v string) *UpdateStorageDomainRoutingRuleShrinkRequest
	GetRoutesShrink() *string
	SetRuleId(v string) *UpdateStorageDomainRoutingRuleShrinkRequest
	GetRuleId() *string
}

type UpdateStorageDomainRoutingRuleShrinkRequest struct {
	// The instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The route list
	//
	// This parameter is required.
	RoutesShrink *string `json:"Routes,omitempty" xml:"Routes,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crsdr-b6thg027zmk1****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateStorageDomainRoutingRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStorageDomainRoutingRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStorageDomainRoutingRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateStorageDomainRoutingRuleShrinkRequest) GetRoutesShrink() *string {
	return s.RoutesShrink
}

func (s *UpdateStorageDomainRoutingRuleShrinkRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateStorageDomainRoutingRuleShrinkRequest) SetInstanceId(v string) *UpdateStorageDomainRoutingRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleShrinkRequest) SetRoutesShrink(v string) *UpdateStorageDomainRoutingRuleShrinkRequest {
	s.RoutesShrink = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleShrinkRequest) SetRuleId(v string) *UpdateStorageDomainRoutingRuleShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
