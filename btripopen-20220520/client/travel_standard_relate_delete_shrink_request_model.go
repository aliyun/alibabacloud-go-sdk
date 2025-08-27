// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateDeleteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromGroup(v bool) *TravelStandardRelateDeleteShrinkRequest
	GetFromGroup() *bool
	SetRemoveListShrink(v string) *TravelStandardRelateDeleteShrinkRequest
	GetRemoveListShrink() *string
	SetRuleId(v int64) *TravelStandardRelateDeleteShrinkRequest
	GetRuleId() *int64
}

type TravelStandardRelateDeleteShrinkRequest struct {
	// example:
	//
	// false
	FromGroup        *bool   `json:"from_group,omitempty" xml:"from_group,omitempty"`
	RemoveListShrink *string `json:"remove_list,omitempty" xml:"remove_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6523763
	RuleId *int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
}

func (s TravelStandardRelateDeleteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateDeleteShrinkRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateDeleteShrinkRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardRelateDeleteShrinkRequest) GetRemoveListShrink() *string {
	return s.RemoveListShrink
}

func (s *TravelStandardRelateDeleteShrinkRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *TravelStandardRelateDeleteShrinkRequest) SetFromGroup(v bool) *TravelStandardRelateDeleteShrinkRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardRelateDeleteShrinkRequest) SetRemoveListShrink(v string) *TravelStandardRelateDeleteShrinkRequest {
	s.RemoveListShrink = &v
	return s
}

func (s *TravelStandardRelateDeleteShrinkRequest) SetRuleId(v int64) *TravelStandardRelateDeleteShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *TravelStandardRelateDeleteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
