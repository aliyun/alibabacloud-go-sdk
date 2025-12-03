// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByIpControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiInfos(v *DescribeApisByIpControlResponseBodyApiInfos) *DescribeApisByIpControlResponseBody
	GetApiInfos() *DescribeApisByIpControlResponseBodyApiInfos
	SetPageNumber(v int32) *DescribeApisByIpControlResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByIpControlResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApisByIpControlResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisByIpControlResponseBody
	GetTotalCount() *int32
}

type DescribeApisByIpControlResponseBody struct {
	// The returned API information. It is an array of ApiInfo data.
	ApiInfos *DescribeApisByIpControlResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
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

func (s DescribeApisByIpControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByIpControlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlResponseBody) GetApiInfos() *DescribeApisByIpControlResponseBodyApiInfos {
	return s.ApiInfos
}

func (s *DescribeApisByIpControlResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByIpControlResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByIpControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisByIpControlResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisByIpControlResponseBody) SetApiInfos(v *DescribeApisByIpControlResponseBodyApiInfos) *DescribeApisByIpControlResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisByIpControlResponseBody) SetPageNumber(v int32) *DescribeApisByIpControlResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByIpControlResponseBody) SetPageSize(v int32) *DescribeApisByIpControlResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByIpControlResponseBody) SetRequestId(v string) *DescribeApisByIpControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByIpControlResponseBody) SetTotalCount(v int32) *DescribeApisByIpControlResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisByIpControlResponseBody) Validate() error {
	if s.ApiInfos != nil {
		if err := s.ApiInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApisByIpControlResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisByIpControlResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByIpControlResponseBodyApiInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByIpControlResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlResponseBodyApiInfos) GetApiInfo() []*DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	return s.ApiInfo
}

func (s *DescribeApisByIpControlResponseBodyApiInfos) SetApiInfo(v []*DescribeApisByIpControlResponseBodyApiInfosApiInfo) *DescribeApisByIpControlResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfos) Validate() error {
	if s.ApiInfo != nil {
		for _, item := range s.ApiInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisByIpControlResponseBodyApiInfosApiInfo struct {
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
	// The time of API binding.
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
	// The name of the API group.
	//
	// example:
	//
	// mygroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The region in which the API is located.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST.
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The visibility of the API. Valid values:
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

func (s DescribeApisByIpControlResponseBodyApiInfosApiInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByIpControlResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetBoundTime() *string {
	return s.BoundTime
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetBoundTime(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.BoundTime = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) Validate() error {
	return dara.Validate(s)
}
