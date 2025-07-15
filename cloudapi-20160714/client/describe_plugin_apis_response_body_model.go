// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiSummarys(v *DescribePluginApisResponseBodyApiSummarys) *DescribePluginApisResponseBody
	GetApiSummarys() *DescribePluginApisResponseBodyApiSummarys
	SetPageNumber(v int32) *DescribePluginApisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginApisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePluginApisResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePluginApisResponseBody
	GetTotalCount() *int32
}

type DescribePluginApisResponseBody struct {
	// The information about APIs.
	ApiSummarys *DescribePluginApisResponseBodyApiSummarys `json:"ApiSummarys,omitempty" xml:"ApiSummarys,omitempty" type:"Struct"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F9C5C4A5-BC6C-57A3-839F-AB08********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePluginApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginApisResponseBody) GetApiSummarys() *DescribePluginApisResponseBodyApiSummarys {
	return s.ApiSummarys
}

func (s *DescribePluginApisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginApisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePluginApisResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePluginApisResponseBody) SetApiSummarys(v *DescribePluginApisResponseBodyApiSummarys) *DescribePluginApisResponseBody {
	s.ApiSummarys = v
	return s
}

func (s *DescribePluginApisResponseBody) SetPageNumber(v int32) *DescribePluginApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginApisResponseBody) SetPageSize(v int32) *DescribePluginApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePluginApisResponseBody) SetRequestId(v string) *DescribePluginApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginApisResponseBody) SetTotalCount(v int32) *DescribePluginApisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePluginApisResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePluginApisResponseBodyApiSummarys struct {
	ApiPluginSummary []*DescribePluginApisResponseBodyApiSummarysApiPluginSummary `json:"ApiPluginSummary,omitempty" xml:"ApiPluginSummary,omitempty" type:"Repeated"`
}

func (s DescribePluginApisResponseBodyApiSummarys) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginApisResponseBodyApiSummarys) GoString() string {
	return s.String()
}

func (s *DescribePluginApisResponseBodyApiSummarys) GetApiPluginSummary() []*DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	return s.ApiPluginSummary
}

func (s *DescribePluginApisResponseBodyApiSummarys) SetApiPluginSummary(v []*DescribePluginApisResponseBodyApiSummarysApiPluginSummary) *DescribePluginApisResponseBodyApiSummarys {
	s.ApiPluginSummary = v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarys) Validate() error {
	return dara.Validate(s)
}

type DescribePluginApisResponseBodyApiSummarysApiPluginSummary struct {
	// The API ID.
	//
	// example:
	//
	// accc8c68b7294b1cb4928741********
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// fhosQueryDayOfStock_V2
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The API description.
	//
	// example:
	//
	// API description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 5f51f89261854fd9ad5116be********
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The API group to which the API belongs.
	//
	// example:
	//
	// myGroup2
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The HTTP method of the API.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /mqTest
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The ID of the region in which the API resides.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The environment alias.
	//
	// example:
	//
	// Production
	StageAlias *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	// The environment to which the API is published. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the pre-release environment
	//
	// 	- **TEST**: the test environment
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribePluginApisResponseBodyApiSummarysApiPluginSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GoString() string {
	return s.String()
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetApiId() *string {
	return s.ApiId
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetApiName() *string {
	return s.ApiName
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetMethod() *string {
	return s.Method
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetPath() *string {
	return s.Path
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetStageAlias() *string {
	return s.StageAlias
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) GetStageName() *string {
	return s.StageName
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetApiId(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.ApiId = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetApiName(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.ApiName = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetDescription(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.Description = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetGroupId(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.GroupId = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetGroupName(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.GroupName = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetMethod(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.Method = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetPath(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.Path = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetRegionId(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.RegionId = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetStageAlias(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.StageAlias = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) SetStageName(v string) *DescribePluginApisResponseBodyApiSummarysApiPluginSummary {
	s.StageName = &v
	return s
}

func (s *DescribePluginApisResponseBodyApiSummarysApiPluginSummary) Validate() error {
	return dara.Validate(s)
}
