// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateAddShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddListShrink(v string) *TravelStandardRelateAddShrinkRequest
	GetAddListShrink() *string
	SetFromGroup(v bool) *TravelStandardRelateAddShrinkRequest
	GetFromGroup() *bool
	SetRuleId(v int64) *TravelStandardRelateAddShrinkRequest
	GetRuleId() *int64
}

type TravelStandardRelateAddShrinkRequest struct {
	AddListShrink *string `json:"add_list,omitempty" xml:"add_list,omitempty"`
	// example:
	//
	// false
	FromGroup *bool `json:"from_group,omitempty" xml:"from_group,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6516571
	RuleId *int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
}

func (s TravelStandardRelateAddShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateAddShrinkRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateAddShrinkRequest) GetAddListShrink() *string {
	return s.AddListShrink
}

func (s *TravelStandardRelateAddShrinkRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardRelateAddShrinkRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *TravelStandardRelateAddShrinkRequest) SetAddListShrink(v string) *TravelStandardRelateAddShrinkRequest {
	s.AddListShrink = &v
	return s
}

func (s *TravelStandardRelateAddShrinkRequest) SetFromGroup(v bool) *TravelStandardRelateAddShrinkRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardRelateAddShrinkRequest) SetRuleId(v int64) *TravelStandardRelateAddShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *TravelStandardRelateAddShrinkRequest) Validate() error {
	return dara.Validate(s)
}
