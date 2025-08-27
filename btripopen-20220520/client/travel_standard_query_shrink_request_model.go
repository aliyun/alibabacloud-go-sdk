// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardQueryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromGroup(v bool) *TravelStandardQueryShrinkRequest
	GetFromGroup() *bool
	SetRuleCode(v int64) *TravelStandardQueryShrinkRequest
	GetRuleCode() *int64
	SetServiceTypeListShrink(v string) *TravelStandardQueryShrinkRequest
	GetServiceTypeListShrink() *string
}

type TravelStandardQueryShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	FromGroup *bool `json:"from_group,omitempty" xml:"from_group,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2006523763
	RuleCode *int64 `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
	// This parameter is required.
	ServiceTypeListShrink *string `json:"service_type_list,omitempty" xml:"service_type_list,omitempty"`
}

func (s TravelStandardQueryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryShrinkRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardQueryShrinkRequest) GetRuleCode() *int64 {
	return s.RuleCode
}

func (s *TravelStandardQueryShrinkRequest) GetServiceTypeListShrink() *string {
	return s.ServiceTypeListShrink
}

func (s *TravelStandardQueryShrinkRequest) SetFromGroup(v bool) *TravelStandardQueryShrinkRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardQueryShrinkRequest) SetRuleCode(v int64) *TravelStandardQueryShrinkRequest {
	s.RuleCode = &v
	return s
}

func (s *TravelStandardQueryShrinkRequest) SetServiceTypeListShrink(v string) *TravelStandardQueryShrinkRequest {
	s.ServiceTypeListShrink = &v
	return s
}

func (s *TravelStandardQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
