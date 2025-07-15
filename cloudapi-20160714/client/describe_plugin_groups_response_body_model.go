// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupSummarys(v *DescribePluginGroupsResponseBodyGroupSummarys) *DescribePluginGroupsResponseBody
	GetGroupSummarys() *DescribePluginGroupsResponseBodyGroupSummarys
	SetPageNumber(v int32) *DescribePluginGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePluginGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePluginGroupsResponseBody
	GetTotalCount() *int32
}

type DescribePluginGroupsResponseBody struct {
	// Collection of group information
	GroupSummarys *DescribePluginGroupsResponseBodyGroupSummarys `json:"GroupSummarys,omitempty" xml:"GroupSummarys,omitempty" type:"Struct"`
	// Pagination parameter: current page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Pagination parameter: number of items per page, default value 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 765BC99E-F583-5A80-9A42-42AC125C2CDC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of returned results
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePluginGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginGroupsResponseBody) GetGroupSummarys() *DescribePluginGroupsResponseBodyGroupSummarys {
	return s.GroupSummarys
}

func (s *DescribePluginGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePluginGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePluginGroupsResponseBody) SetGroupSummarys(v *DescribePluginGroupsResponseBodyGroupSummarys) *DescribePluginGroupsResponseBody {
	s.GroupSummarys = v
	return s
}

func (s *DescribePluginGroupsResponseBody) SetPageNumber(v int32) *DescribePluginGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginGroupsResponseBody) SetPageSize(v int32) *DescribePluginGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePluginGroupsResponseBody) SetRequestId(v string) *DescribePluginGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginGroupsResponseBody) SetTotalCount(v int32) *DescribePluginGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePluginGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePluginGroupsResponseBodyGroupSummarys struct {
	GroupPluginSummary []*DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary `json:"GroupPluginSummary,omitempty" xml:"GroupPluginSummary,omitempty" type:"Repeated"`
}

func (s DescribePluginGroupsResponseBodyGroupSummarys) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginGroupsResponseBodyGroupSummarys) GoString() string {
	return s.String()
}

func (s *DescribePluginGroupsResponseBodyGroupSummarys) GetGroupPluginSummary() []*DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary {
	return s.GroupPluginSummary
}

func (s *DescribePluginGroupsResponseBodyGroupSummarys) SetGroupPluginSummary(v []*DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) *DescribePluginGroupsResponseBodyGroupSummarys {
	s.GroupPluginSummary = v
	return s
}

func (s *DescribePluginGroupsResponseBodyGroupSummarys) Validate() error {
	return dara.Validate(s)
}

type DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary struct {
	// API root path
	//
	// example:
	//
	// /rpew
	BasePath *string `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	// Description
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// API group ID
	//
	// example:
	//
	// 4ed31575e2de43de8c51eb1217a1f56b
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// API group name
	//
	// example:
	//
	// ECP_API
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Region ID where the API group is located
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Stage name Alias
	//
	// example:
	//
	// DEV
	StageAlias *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	// Environment name, possible values:
	//
	// - **RELEASE**: Production
	//
	// - **PRE**: Pre-release
	//
	// - **TEST**: Testing
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) GoString() string {
	return s.String()
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) GetBasePath() *string {
	return s.BasePath
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) GetStageAlias() *string {
	return s.StageAlias
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) GetStageName() *string {
	return s.StageName
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) SetBasePath(v string) *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary {
	s.BasePath = &v
	return s
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) SetDescription(v string) *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary {
	s.Description = &v
	return s
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) SetGroupId(v string) *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary {
	s.GroupId = &v
	return s
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) SetGroupName(v string) *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary {
	s.GroupName = &v
	return s
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) SetRegionId(v string) *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary {
	s.RegionId = &v
	return s
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) SetStageAlias(v string) *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary {
	s.StageAlias = &v
	return s
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) SetStageName(v string) *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary {
	s.StageName = &v
	return s
}

func (s *DescribePluginGroupsResponseBodyGroupSummarysGroupPluginSummary) Validate() error {
	return dara.Validate(s)
}
