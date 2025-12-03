// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiHisItems(v *DescribeApiHistoriesResponseBodyApiHisItems) *DescribeApiHistoriesResponseBody
	GetApiHisItems() *DescribeApiHistoriesResponseBodyApiHisItems
	SetPageNumber(v int32) *DescribeApiHistoriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiHistoriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApiHistoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApiHistoriesResponseBody
	GetTotalCount() *int32
}

type DescribeApiHistoriesResponseBody struct {
	// The returned API information. It is an array consisting of ApiHisItem data.
	ApiHisItems *DescribeApiHistoriesResponseBodyApiHisItems `json:"ApiHisItems,omitempty" xml:"ApiHisItems,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 15
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ003
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesResponseBody) GetApiHisItems() *DescribeApiHistoriesResponseBodyApiHisItems {
	return s.ApiHisItems
}

func (s *DescribeApiHistoriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiHistoriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiHistoriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApiHistoriesResponseBody) SetApiHisItems(v *DescribeApiHistoriesResponseBodyApiHisItems) *DescribeApiHistoriesResponseBody {
	s.ApiHisItems = v
	return s
}

func (s *DescribeApiHistoriesResponseBody) SetPageNumber(v int32) *DescribeApiHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiHistoriesResponseBody) SetPageSize(v int32) *DescribeApiHistoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiHistoriesResponseBody) SetRequestId(v string) *DescribeApiHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiHistoriesResponseBody) SetTotalCount(v int32) *DescribeApiHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApiHistoriesResponseBody) Validate() error {
	if s.ApiHisItems != nil {
		if err := s.ApiHisItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiHistoriesResponseBodyApiHisItems struct {
	ApiHisItem []*DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem `json:"ApiHisItem,omitempty" xml:"ApiHisItem,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoriesResponseBodyApiHisItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoriesResponseBodyApiHisItems) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesResponseBodyApiHisItems) GetApiHisItem() []*DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	return s.ApiHisItem
}

func (s *DescribeApiHistoriesResponseBodyApiHisItems) SetApiHisItem(v []*DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) *DescribeApiHistoriesResponseBodyApiHisItems {
	s.ApiHisItem = v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItems) Validate() error {
	if s.ApiHisItem != nil {
		for _, item := range s.ApiHisItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem struct {
	// The ID of the API.
	//
	// example:
	//
	// 5af418828f0344a3b588c0cc1331a3bc
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// CreateObject
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The publishing time (UTC) of the API.
	//
	// example:
	//
	// 2016-07-20T08:28:48Z
	DeployedTime *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// Creates an object
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 1084f9034c744137901057206b39d2b6
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// myGroup2
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The historical version of the API.
	//
	// example:
	//
	// 20160705104552393
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
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
	// 	- **TEST**
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// Indicates whether an API version is effective. Valid values: **ONLINE*	- and **OFFLINE**.
	//
	// example:
	//
	// ONLINE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetDeployedTime() *string {
	return s.DeployedTime
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetHistoryVersion() *string {
	return s.HistoryVersion
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GetStatus() *string {
	return s.Status
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetApiId(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.ApiId = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetApiName(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.ApiName = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetDeployedTime(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.DeployedTime = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetDescription(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetGroupId(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.GroupId = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetGroupName(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.GroupName = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetHistoryVersion(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetRegionId(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.RegionId = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetStageName(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.StageName = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetStatus(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.Status = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) Validate() error {
	return dara.Validate(s)
}
