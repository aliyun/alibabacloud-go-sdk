// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiHisItems(v *DescribeHistoryApisResponseBodyApiHisItems) *DescribeHistoryApisResponseBody
	GetApiHisItems() *DescribeHistoryApisResponseBodyApiHisItems
	SetPageNumber(v int32) *DescribeHistoryApisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryApisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHistoryApisResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHistoryApisResponseBody
	GetTotalCount() *int32
}

type DescribeHistoryApisResponseBody struct {
	ApiHisItems *DescribeHistoryApisResponseBodyApiHisItems `json:"ApiHisItems,omitempty" xml:"ApiHisItems,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6C87A26A-6A18-4B8E-8099-705278381A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBody) GetApiHisItems() *DescribeHistoryApisResponseBodyApiHisItems {
	return s.ApiHisItems
}

func (s *DescribeHistoryApisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryApisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryApisResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHistoryApisResponseBody) SetApiHisItems(v *DescribeHistoryApisResponseBodyApiHisItems) *DescribeHistoryApisResponseBody {
	s.ApiHisItems = v
	return s
}

func (s *DescribeHistoryApisResponseBody) SetPageNumber(v int32) *DescribeHistoryApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryApisResponseBody) SetPageSize(v int32) *DescribeHistoryApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryApisResponseBody) SetRequestId(v string) *DescribeHistoryApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryApisResponseBody) SetTotalCount(v int32) *DescribeHistoryApisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHistoryApisResponseBody) Validate() error {
	if s.ApiHisItems != nil {
		if err := s.ApiHisItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHistoryApisResponseBodyApiHisItems struct {
	ApiHisItem []*DescribeHistoryApisResponseBodyApiHisItemsApiHisItem `json:"ApiHisItem,omitempty" xml:"ApiHisItem,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApisResponseBodyApiHisItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryApisResponseBodyApiHisItems) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBodyApiHisItems) GetApiHisItem() []*DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	return s.ApiHisItem
}

func (s *DescribeHistoryApisResponseBodyApiHisItems) SetApiHisItem(v []*DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) *DescribeHistoryApisResponseBodyApiHisItems {
	s.ApiHisItem = v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItems) Validate() error {
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

type DescribeHistoryApisResponseBodyApiHisItemsApiHisItem struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName        *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	DeployedTime   *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageAlias     *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetDeployedTime() *string {
	return s.DeployedTime
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetDescription() *string {
	return s.Description
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetHistoryVersion() *string {
	return s.HistoryVersion
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetStageAlias() *string {
	return s.StageAlias
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetStageName() *string {
	return s.StageName
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GetStatus() *string {
	return s.Status
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetApiId(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.ApiId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetApiName(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.ApiName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetDeployedTime(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.DeployedTime = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetDescription(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetGroupId(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.GroupId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetGroupName(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.GroupName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetHistoryVersion(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetRegionId(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetStageAlias(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.StageAlias = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetStageName(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.StageName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetStatus(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.Status = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) Validate() error {
	return dara.Validate(s)
}
