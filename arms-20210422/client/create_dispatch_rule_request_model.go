// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDispatchRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchRule(v string) *CreateDispatchRuleRequest
	GetDispatchRule() *string
	SetRegionId(v string) *CreateDispatchRuleRequest
	GetRegionId() *string
}

type CreateDispatchRuleRequest struct {
	// This parameter is required.
	DispatchRule *string `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDispatchRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleRequest) GetDispatchRule() *string {
	return s.DispatchRule
}

func (s *CreateDispatchRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDispatchRuleRequest) SetDispatchRule(v string) *CreateDispatchRuleRequest {
	s.DispatchRule = &v
	return s
}

func (s *CreateDispatchRuleRequest) SetRegionId(v string) *CreateDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDispatchRuleRequest) Validate() error {
	return dara.Validate(s)
}
