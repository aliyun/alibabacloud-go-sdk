// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDispatchRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchRule(v string) *UpdateDispatchRuleRequest
	GetDispatchRule() *string
	SetRegionId(v string) *UpdateDispatchRuleRequest
	GetRegionId() *string
}

type UpdateDispatchRuleRequest struct {
	// This parameter is required.
	DispatchRule *string `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateDispatchRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleRequest) GetDispatchRule() *string {
	return s.DispatchRule
}

func (s *UpdateDispatchRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDispatchRuleRequest) SetDispatchRule(v string) *UpdateDispatchRuleRequest {
	s.DispatchRule = &v
	return s
}

func (s *UpdateDispatchRuleRequest) SetRegionId(v string) *UpdateDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDispatchRuleRequest) Validate() error {
	return dara.Validate(s)
}
