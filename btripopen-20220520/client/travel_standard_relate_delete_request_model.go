// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromGroup(v bool) *TravelStandardRelateDeleteRequest
	GetFromGroup() *bool
	SetRemoveList(v []*TravelStandardRelateDeleteRequestRemoveList) *TravelStandardRelateDeleteRequest
	GetRemoveList() []*TravelStandardRelateDeleteRequestRemoveList
	SetRuleId(v int64) *TravelStandardRelateDeleteRequest
	GetRuleId() *int64
}

type TravelStandardRelateDeleteRequest struct {
	// example:
	//
	// false
	FromGroup  *bool                                          `json:"from_group,omitempty" xml:"from_group,omitempty"`
	RemoveList []*TravelStandardRelateDeleteRequestRemoveList `json:"remove_list,omitempty" xml:"remove_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 6523763
	RuleId *int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
}

func (s TravelStandardRelateDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateDeleteRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateDeleteRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardRelateDeleteRequest) GetRemoveList() []*TravelStandardRelateDeleteRequestRemoveList {
	return s.RemoveList
}

func (s *TravelStandardRelateDeleteRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *TravelStandardRelateDeleteRequest) SetFromGroup(v bool) *TravelStandardRelateDeleteRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardRelateDeleteRequest) SetRemoveList(v []*TravelStandardRelateDeleteRequestRemoveList) *TravelStandardRelateDeleteRequest {
	s.RemoveList = v
	return s
}

func (s *TravelStandardRelateDeleteRequest) SetRuleId(v int64) *TravelStandardRelateDeleteRequest {
	s.RuleId = &v
	return s
}

func (s *TravelStandardRelateDeleteRequest) Validate() error {
	return dara.Validate(s)
}

type TravelStandardRelateDeleteRequestRemoveList struct {
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

func (s TravelStandardRelateDeleteRequestRemoveList) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateDeleteRequestRemoveList) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateDeleteRequestRemoveList) GetEntityId() *string {
	return s.EntityId
}

func (s *TravelStandardRelateDeleteRequestRemoveList) GetEntityType() *string {
	return s.EntityType
}

func (s *TravelStandardRelateDeleteRequestRemoveList) SetEntityId(v string) *TravelStandardRelateDeleteRequestRemoveList {
	s.EntityId = &v
	return s
}

func (s *TravelStandardRelateDeleteRequestRemoveList) SetEntityType(v string) *TravelStandardRelateDeleteRequestRemoveList {
	s.EntityType = &v
	return s
}

func (s *TravelStandardRelateDeleteRequestRemoveList) Validate() error {
	return dara.Validate(s)
}
