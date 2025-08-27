// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromGroup(v bool) *TravelStandardQueryRequest
	GetFromGroup() *bool
	SetRuleCode(v int64) *TravelStandardQueryRequest
	GetRuleCode() *int64
	SetServiceTypeList(v []*string) *TravelStandardQueryRequest
	GetServiceTypeList() []*string
}

type TravelStandardQueryRequest struct {
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
	ServiceTypeList []*string `json:"service_type_list,omitempty" xml:"service_type_list,omitempty" type:"Repeated"`
}

func (s TravelStandardQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardQueryRequest) GetRuleCode() *int64 {
	return s.RuleCode
}

func (s *TravelStandardQueryRequest) GetServiceTypeList() []*string {
	return s.ServiceTypeList
}

func (s *TravelStandardQueryRequest) SetFromGroup(v bool) *TravelStandardQueryRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardQueryRequest) SetRuleCode(v int64) *TravelStandardQueryRequest {
	s.RuleCode = &v
	return s
}

func (s *TravelStandardQueryRequest) SetServiceTypeList(v []*string) *TravelStandardQueryRequest {
	s.ServiceTypeList = v
	return s
}

func (s *TravelStandardQueryRequest) Validate() error {
	return dara.Validate(s)
}
