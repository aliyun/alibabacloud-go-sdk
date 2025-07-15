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
	// The returned API information. It is an array consisting of ApiHisItems.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeHistoryApisResponseBodyApiHisItemsApiHisItem struct {
	// The API ID.
	//
	// example:
	//
	// 5af418828f0344a3b588c0cc1331a3bc
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// v2_role_assign
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The time when the API was published. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-07-20T08:28:48Z
	DeployedTime *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	// The API description.
	//
	// example:
	//
	// Creates an object
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The API group ID.
	//
	// example:
	//
	// 1084f9034c744137901057206b39d2b6
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group to which the API belongs.
	//
	// example:
	//
	// myGroup2
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The historical version of the API definition.
	//
	// example:
	//
	// 20210915101416294
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The environment alias.
	//
	// example:
	//
	// Online
	StageAlias *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	// The environment name. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// Indicates whether an API version is effective. Valid values: ONLINE and OFFLINE.
	//
	// example:
	//
	// ONLINE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
