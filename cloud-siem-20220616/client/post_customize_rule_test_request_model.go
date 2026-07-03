// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostCustomizeRuleTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *PostCustomizeRuleTestRequest
	GetId() *int64
	SetRegionId(v string) *PostCustomizeRuleTestRequest
	GetRegionId() *string
	SetRoleFor(v int64) *PostCustomizeRuleTestRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *PostCustomizeRuleTestRequest
	GetRoleType() *int32
	SetSimulatedData(v string) *PostCustomizeRuleTestRequest
	GetSimulatedData() *string
	SetTestType(v string) *PostCustomizeRuleTestRequest
	GetTestType() *string
}

type PostCustomizeRuleTestRequest struct {
	// The ID of the custom rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region where the Data Management center of threat analysis is located. Select the region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member that the administrator wants to impersonate to view data.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The simulated data for the test. This parameter is required only when TestType is set to simulate.
	//
	// example:
	//
	// [{"key1":"value1","key2":"value2","key3":"value3","key4":"value4","key5":"value5"}]
	SimulatedData *string `json:"SimulatedData,omitempty" xml:"SimulatedData,omitempty"`
	// The test type. Valid values:
	//
	// - simulate: tests the rule with simulated data
	//
	// - business: tests the rule with business data
	//
	// example:
	//
	// simulate
	TestType *string `json:"TestType,omitempty" xml:"TestType,omitempty"`
}

func (s PostCustomizeRuleTestRequest) String() string {
	return dara.Prettify(s)
}

func (s PostCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestRequest) GetId() *int64 {
	return s.Id
}

func (s *PostCustomizeRuleTestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PostCustomizeRuleTestRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *PostCustomizeRuleTestRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *PostCustomizeRuleTestRequest) GetSimulatedData() *string {
	return s.SimulatedData
}

func (s *PostCustomizeRuleTestRequest) GetTestType() *string {
	return s.TestType
}

func (s *PostCustomizeRuleTestRequest) SetId(v int64) *PostCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetRegionId(v string) *PostCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetRoleFor(v int64) *PostCustomizeRuleTestRequest {
	s.RoleFor = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetRoleType(v int32) *PostCustomizeRuleTestRequest {
	s.RoleType = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetSimulatedData(v string) *PostCustomizeRuleTestRequest {
	s.SimulatedData = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetTestType(v string) *PostCustomizeRuleTestRequest {
	s.TestType = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) Validate() error {
	return dara.Validate(s)
}
