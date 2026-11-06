// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DescribePolarClawSkillsRequest
	GetAgentId() *string
	SetApplicationId(v string) *DescribePolarClawSkillsRequest
	GetApplicationId() *string
	SetPageNumber(v int32) *DescribePolarClawSkillsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolarClawSkillsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribePolarClawSkillsRequest
	GetRegionId() *string
}

type DescribePolarClawSkillsRequest struct {
	// Specifies the agent workspace perspective. If this parameter is specified, the Disabled and AvailableToAgent fields are returned based on the specified agent. If this parameter is not specified, the default agent is used.
	//
	// example:
	//
	// test1
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The page number. The value must be an integer greater than 0 and less than or equal to the maximum value of the Integer data type.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 30, 50, and 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all available regions for the current account, including region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePolarClawSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribePolarClawSkillsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawSkillsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarClawSkillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolarClawSkillsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarClawSkillsRequest) SetAgentId(v string) *DescribePolarClawSkillsRequest {
	s.AgentId = &v
	return s
}

func (s *DescribePolarClawSkillsRequest) SetApplicationId(v string) *DescribePolarClawSkillsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillsRequest) SetPageNumber(v int32) *DescribePolarClawSkillsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarClawSkillsRequest) SetPageSize(v int32) *DescribePolarClawSkillsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolarClawSkillsRequest) SetRegionId(v string) *DescribePolarClawSkillsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarClawSkillsRequest) Validate() error {
	return dara.Validate(s)
}
