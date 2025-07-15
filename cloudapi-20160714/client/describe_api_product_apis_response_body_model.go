// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiProductApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiInfoList(v *DescribeApiProductApisResponseBodyApiInfoList) *DescribeApiProductApisResponseBody
	GetApiInfoList() *DescribeApiProductApisResponseBodyApiInfoList
	SetPageNumber(v int32) *DescribeApiProductApisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiProductApisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApiProductApisResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApiProductApisResponseBody
	GetTotalCount() *int32
}

type DescribeApiProductApisResponseBody struct {
	// The information about the returned APIs.
	ApiInfoList *DescribeApiProductApisResponseBodyApiInfoList `json:"ApiInfoList,omitempty" xml:"ApiInfoList,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 03442A3D-3B7D-434C-8A95-A5FEB989B519
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiProductApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiProductApisResponseBody) GetApiInfoList() *DescribeApiProductApisResponseBodyApiInfoList {
	return s.ApiInfoList
}

func (s *DescribeApiProductApisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiProductApisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiProductApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiProductApisResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApiProductApisResponseBody) SetApiInfoList(v *DescribeApiProductApisResponseBodyApiInfoList) *DescribeApiProductApisResponseBody {
	s.ApiInfoList = v
	return s
}

func (s *DescribeApiProductApisResponseBody) SetPageNumber(v int32) *DescribeApiProductApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiProductApisResponseBody) SetPageSize(v int32) *DescribeApiProductApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiProductApisResponseBody) SetRequestId(v string) *DescribeApiProductApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiProductApisResponseBody) SetTotalCount(v int32) *DescribeApiProductApisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApiProductApisResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiProductApisResponseBodyApiInfoList struct {
	ApiInfo []*DescribeApiProductApisResponseBodyApiInfoListApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiProductApisResponseBodyApiInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductApisResponseBodyApiInfoList) GoString() string {
	return s.String()
}

func (s *DescribeApiProductApisResponseBodyApiInfoList) GetApiInfo() []*DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	return s.ApiInfo
}

func (s *DescribeApiProductApisResponseBodyApiInfoList) SetApiInfo(v []*DescribeApiProductApisResponseBodyApiInfoListApiInfo) *DescribeApiProductApisResponseBodyApiInfoList {
	s.ApiInfo = v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeApiProductApisResponseBodyApiInfoListApiInfo struct {
	// The API ID.
	//
	// example:
	//
	// dd46297680014a7e8e318308f3345951
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// testApi
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The API description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 1e377f18142345dfb700cd8911c2463a
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group to which the API belongs.
	//
	// example:
	//
	// testApiGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The request method of the API.
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The ID of the region where the API is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The environment to which the API is published. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the staging environment
	//
	// 	- **TEST**: the test environment
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiProductApisResponseBodyApiInfoListApiInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiProductApisResponseBodyApiInfoListApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetMethod() *string {
	return s.Method
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetPath() *string {
	return s.Path
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetApiId(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetApiName(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetDescription(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetGroupId(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetGroupName(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetMethod(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.Method = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetPath(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.Path = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetRegionId(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) SetStageName(v string) *DescribeApiProductApisResponseBodyApiInfoListApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApiProductApisResponseBodyApiInfoListApiInfo) Validate() error {
	return dara.Validate(s)
}
