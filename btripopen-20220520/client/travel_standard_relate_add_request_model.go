// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddList(v []*TravelStandardRelateAddRequestAddList) *TravelStandardRelateAddRequest
	GetAddList() []*TravelStandardRelateAddRequestAddList
	SetFromGroup(v bool) *TravelStandardRelateAddRequest
	GetFromGroup() *bool
	SetRuleId(v int64) *TravelStandardRelateAddRequest
	GetRuleId() *int64
}

type TravelStandardRelateAddRequest struct {
	AddList []*TravelStandardRelateAddRequestAddList `json:"add_list,omitempty" xml:"add_list,omitempty" type:"Repeated"`
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

func (s TravelStandardRelateAddRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateAddRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateAddRequest) GetAddList() []*TravelStandardRelateAddRequestAddList {
	return s.AddList
}

func (s *TravelStandardRelateAddRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardRelateAddRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *TravelStandardRelateAddRequest) SetAddList(v []*TravelStandardRelateAddRequestAddList) *TravelStandardRelateAddRequest {
	s.AddList = v
	return s
}

func (s *TravelStandardRelateAddRequest) SetFromGroup(v bool) *TravelStandardRelateAddRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardRelateAddRequest) SetRuleId(v int64) *TravelStandardRelateAddRequest {
	s.RuleId = &v
	return s
}

func (s *TravelStandardRelateAddRequest) Validate() error {
	return dara.Validate(s)
}

type TravelStandardRelateAddRequestAddList struct {
	// This parameter is required.
	//
	// example:
	//
	// 667104628
	EntityId *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s TravelStandardRelateAddRequestAddList) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateAddRequestAddList) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateAddRequestAddList) GetEntityId() *string {
	return s.EntityId
}

func (s *TravelStandardRelateAddRequestAddList) GetEntityType() *string {
	return s.EntityType
}

func (s *TravelStandardRelateAddRequestAddList) SetEntityId(v string) *TravelStandardRelateAddRequestAddList {
	s.EntityId = &v
	return s
}

func (s *TravelStandardRelateAddRequestAddList) SetEntityType(v string) *TravelStandardRelateAddRequestAddList {
	s.EntityType = &v
	return s
}

func (s *TravelStandardRelateAddRequestAddList) Validate() error {
	return dara.Validate(s)
}
