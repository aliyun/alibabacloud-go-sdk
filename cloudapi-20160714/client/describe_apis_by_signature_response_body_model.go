// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisBySignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiInfos(v *DescribeApisBySignatureResponseBodyApiInfos) *DescribeApisBySignatureResponseBody
	GetApiInfos() *DescribeApisBySignatureResponseBodyApiInfos
	SetPageNumber(v int32) *DescribeApisBySignatureResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisBySignatureResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApisBySignatureResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisBySignatureResponseBody
	GetTotalCount() *int32
}

type DescribeApisBySignatureResponseBody struct {
	// The returned API information. It is an array consisting of ApiInfo data.
	ApiInfos *DescribeApisBySignatureResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
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

func (s DescribeApisBySignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisBySignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureResponseBody) GetApiInfos() *DescribeApisBySignatureResponseBodyApiInfos {
	return s.ApiInfos
}

func (s *DescribeApisBySignatureResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisBySignatureResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisBySignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisBySignatureResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisBySignatureResponseBody) SetApiInfos(v *DescribeApisBySignatureResponseBodyApiInfos) *DescribeApisBySignatureResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisBySignatureResponseBody) SetPageNumber(v int32) *DescribeApisBySignatureResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisBySignatureResponseBody) SetPageSize(v int32) *DescribeApisBySignatureResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisBySignatureResponseBody) SetRequestId(v string) *DescribeApisBySignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisBySignatureResponseBody) SetTotalCount(v int32) *DescribeApisBySignatureResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisBySignatureResponseBody) Validate() error {
	if s.ApiInfos != nil {
		if err := s.ApiInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApisBySignatureResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisBySignatureResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisBySignatureResponseBodyApiInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisBySignatureResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureResponseBodyApiInfos) GetApiInfo() []*DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	return s.ApiInfo
}

func (s *DescribeApisBySignatureResponseBodyApiInfos) SetApiInfo(v []*DescribeApisBySignatureResponseBodyApiInfosApiInfo) *DescribeApisBySignatureResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfos) Validate() error {
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

type DescribeApisBySignatureResponseBodyApiInfosApiInfo struct {
	// The ID of the API.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
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
	// The name of the group to which the API belongs.
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

func (s DescribeApisBySignatureResponseBodyApiInfosApiInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisBySignatureResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetBoundTime() *string {
	return s.BoundTime
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetBoundTime(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.BoundTime = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) Validate() error {
	return dara.Validate(s)
}
