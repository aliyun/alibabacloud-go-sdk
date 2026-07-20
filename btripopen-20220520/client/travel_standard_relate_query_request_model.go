// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromGroup(v bool) *TravelStandardRelateQueryRequest
	GetFromGroup() *bool
	SetRuleId(v int64) *TravelStandardRelateQueryRequest
	GetRuleId() *int64
}

type TravelStandardRelateQueryRequest struct {
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

func (s TravelStandardRelateQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateQueryRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateQueryRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardRelateQueryRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *TravelStandardRelateQueryRequest) SetFromGroup(v bool) *TravelStandardRelateQueryRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardRelateQueryRequest) SetRuleId(v int64) *TravelStandardRelateQueryRequest {
	s.RuleId = &v
	return s
}

func (s *TravelStandardRelateQueryRequest) Validate() error {
	return dara.Validate(s)
}
