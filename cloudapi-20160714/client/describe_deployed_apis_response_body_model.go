// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeployedApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployedApis(v *DescribeDeployedApisResponseBodyDeployedApis) *DescribeDeployedApisResponseBody
	GetDeployedApis() *DescribeDeployedApisResponseBodyDeployedApis
	SetPageNumber(v int32) *DescribeDeployedApisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeployedApisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDeployedApisResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDeployedApisResponseBody
	GetTotalCount() *int32
}

type DescribeDeployedApisResponseBody struct {
	DeployedApis *DescribeDeployedApisResponseBodyDeployedApis `json:"DeployedApis,omitempty" xml:"DeployedApis,omitempty" type:"Struct"`
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
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeployedApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBody) GetDeployedApis() *DescribeDeployedApisResponseBodyDeployedApis {
	return s.DeployedApis
}

func (s *DescribeDeployedApisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeployedApisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeployedApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeployedApisResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDeployedApisResponseBody) SetDeployedApis(v *DescribeDeployedApisResponseBodyDeployedApis) *DescribeDeployedApisResponseBody {
	s.DeployedApis = v
	return s
}

func (s *DescribeDeployedApisResponseBody) SetPageNumber(v int32) *DescribeDeployedApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeployedApisResponseBody) SetPageSize(v int32) *DescribeDeployedApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeployedApisResponseBody) SetRequestId(v string) *DescribeDeployedApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeployedApisResponseBody) SetTotalCount(v int32) *DescribeDeployedApisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeployedApisResponseBody) Validate() error {
	if s.DeployedApis != nil {
		if err := s.DeployedApis.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeployedApisResponseBodyDeployedApis struct {
	DeployedApiItem []*DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem `json:"DeployedApiItem,omitempty" xml:"DeployedApiItem,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApisResponseBodyDeployedApis) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApisResponseBodyDeployedApis) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBodyDeployedApis) GetDeployedApiItem() []*DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	return s.DeployedApiItem
}

func (s *DescribeDeployedApisResponseBodyDeployedApis) SetDeployedApiItem(v []*DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) *DescribeDeployedApisResponseBodyDeployedApis {
	s.DeployedApiItem = v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApis) Validate() error {
	if s.DeployedApiItem != nil {
		for _, item := range s.DeployedApiItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem struct {
	ApiId        *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiMethod    *string `json:"ApiMethod,omitempty" xml:"ApiMethod,omitempty"`
	ApiName      *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiPath      *string `json:"ApiPath,omitempty" xml:"ApiPath,omitempty"`
	DeployedTime *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName    *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility   *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetApiMethod() *string {
	return s.ApiMethod
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetApiPath() *string {
	return s.ApiPath
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetDeployedTime() *string {
	return s.DeployedTime
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetStageName() *string {
	return s.StageName
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetApiId(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.ApiId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetApiMethod(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.ApiMethod = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetApiName(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.ApiName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetApiPath(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.ApiPath = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetDeployedTime(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.DeployedTime = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetDescription(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetGroupId(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetGroupName(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.GroupName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetRegionId(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.RegionId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetStageName(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.StageName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetVisibility(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.Visibility = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) Validate() error {
	return dara.Validate(s)
}
