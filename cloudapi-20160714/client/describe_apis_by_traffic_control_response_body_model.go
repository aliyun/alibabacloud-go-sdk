// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByTrafficControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiInfos(v *DescribeApisByTrafficControlResponseBodyApiInfos) *DescribeApisByTrafficControlResponseBody
	GetApiInfos() *DescribeApisByTrafficControlResponseBodyApiInfos
	SetPageNumber(v int32) *DescribeApisByTrafficControlResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByTrafficControlResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApisByTrafficControlResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisByTrafficControlResponseBody
	GetTotalCount() *int32
}

type DescribeApisByTrafficControlResponseBody struct {
	// The returned API information. It is an array consisting of ApiInfo data.
	ApiInfos *DescribeApisByTrafficControlResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	// The page number of the returned page.
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
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisByTrafficControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByTrafficControlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlResponseBody) GetApiInfos() *DescribeApisByTrafficControlResponseBodyApiInfos {
	return s.ApiInfos
}

func (s *DescribeApisByTrafficControlResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByTrafficControlResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByTrafficControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisByTrafficControlResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisByTrafficControlResponseBody) SetApiInfos(v *DescribeApisByTrafficControlResponseBodyApiInfos) *DescribeApisByTrafficControlResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) SetPageNumber(v int32) *DescribeApisByTrafficControlResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) SetPageSize(v int32) *DescribeApisByTrafficControlResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) SetRequestId(v string) *DescribeApisByTrafficControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) SetTotalCount(v int32) *DescribeApisByTrafficControlResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApisByTrafficControlResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisByTrafficControlResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByTrafficControlResponseBodyApiInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByTrafficControlResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfos) GetApiInfo() []*DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	return s.ApiInfo
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfos) SetApiInfo(v []*DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) *DescribeApisByTrafficControlResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeApisByTrafficControlResponseBodyApiInfosApiInfo struct {
	// The ID of the API.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API
	//
	// example:
	//
	// testapi
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The binding time of the API.
	//
	// example:
	//
	// 2016-07-23T08:28:48Z
	BoundTime *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group to which an API belongs.
	//
	// example:
	//
	// mygroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The region where the API is located.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// Indicates whether the API is public. Valid values:
	//
	// 	- **PUBLIC**
	//
	// 	- **PRIVATE**
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetBoundTime() *string {
	return s.BoundTime
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetBoundTime(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.BoundTime = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) Validate() error {
	return dara.Validate(s)
}
