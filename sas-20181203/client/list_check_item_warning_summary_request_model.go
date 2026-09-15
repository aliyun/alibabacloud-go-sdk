// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemWarningSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckItemFuzzy(v string) *ListCheckItemWarningSummaryRequest
	GetCheckItemFuzzy() *string
	SetCheckLevel(v string) *ListCheckItemWarningSummaryRequest
	GetCheckLevel() *string
	SetCheckType(v string) *ListCheckItemWarningSummaryRequest
	GetCheckType() *string
	SetCheckWarningStatus(v int32) *ListCheckItemWarningSummaryRequest
	GetCheckWarningStatus() *int32
	SetCheckWarningStatusList(v []*int32) *ListCheckItemWarningSummaryRequest
	GetCheckWarningStatusList() []*int32
	SetContainerFieldName(v string) *ListCheckItemWarningSummaryRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *ListCheckItemWarningSummaryRequest
	GetContainerFieldValue() *string
	SetCurrentPage(v int32) *ListCheckItemWarningSummaryRequest
	GetCurrentPage() *int32
	SetGroupId(v int64) *ListCheckItemWarningSummaryRequest
	GetGroupId() *int64
	SetLang(v string) *ListCheckItemWarningSummaryRequest
	GetLang() *string
	SetPageSize(v int32) *ListCheckItemWarningSummaryRequest
	GetPageSize() *int32
	SetResourceDirectoryAccountId(v int64) *ListCheckItemWarningSummaryRequest
	GetResourceDirectoryAccountId() *int64
	SetRiskType(v string) *ListCheckItemWarningSummaryRequest
	GetRiskType() *string
	SetSource(v string) *ListCheckItemWarningSummaryRequest
	GetSource() *string
	SetStartTime(v int64) *ListCheckItemWarningSummaryRequest
	GetStartTime() *int64
	SetUuidList(v []*string) *ListCheckItemWarningSummaryRequest
	GetUuidList() []*string
}

type ListCheckItemWarningSummaryRequest struct {
	// The fuzzy match for the check item name.
	//
	// example:
	//
	// password
	CheckItemFuzzy *string `json:"CheckItemFuzzy,omitempty" xml:"CheckItemFuzzy,omitempty"`
	// The risk level. Default value: null, which indicates that all levels are queried. Valid values:
	//
	// - **high**: High.
	//
	// - **medium**: Medium.
	//
	// - **low**: Low.
	//
	// example:
	//
	// medium
	CheckLevel *string `json:"CheckLevel,omitempty" xml:"CheckLevel,omitempty"`
	// The check item category name.
	//
	// example:
	//
	// hc.check.type.attack_defense
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The risk status. Default value: null, which indicates that all statuses are queried. Valid values:
	//
	// - **1**: Failed.
	//
	// - **3**: Passed.
	//
	// - **6**: Whitelisted.
	//
	// - **8**: Fixed.
	//
	// example:
	//
	// 3
	CheckWarningStatus *int32 `json:"CheckWarningStatus,omitempty" xml:"CheckWarningStatus,omitempty"`
	// The list of risk statuses. If both this parameter and CheckWarningStatus are specified, only CheckWarningStatus takes effect.
	CheckWarningStatusList []*int32 `json:"CheckWarningStatusList,omitempty" xml:"CheckWarningStatusList,omitempty" type:"Repeated"`
	// The container security query parameter name.
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The container security query parameter value.
	//
	// example:
	//
	// c471f0f61b9c04f8380556e922cf1****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The page number of the page to return. Default value: **1**, which indicates that query results are displayed starting from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the asset group to query.
	//
	// > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1161****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page when paging. Default value: 20. If the PageSize parameter is left empty, 20 entries are returned per page.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the member accounts in the resource directory (Alibaba Cloud account).
	//
	// > You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The baseline category name.
	//
	// example:
	//
	// weak_password
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// The data source. Default value: **default**. Valid values:
	//
	// - **agentless**: agentless detection.
	//
	// - **default**: host baseline.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start time for filtering alerts. This parameter takes effect only when you query historical processed alerts. Specify a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1732793158366
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of server UUIDs to query.
	//
	// > You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to obtain the UUID of a server.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s ListCheckItemWarningSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningSummaryRequest) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningSummaryRequest) GetCheckItemFuzzy() *string {
	return s.CheckItemFuzzy
}

func (s *ListCheckItemWarningSummaryRequest) GetCheckLevel() *string {
	return s.CheckLevel
}

func (s *ListCheckItemWarningSummaryRequest) GetCheckType() *string {
	return s.CheckType
}

func (s *ListCheckItemWarningSummaryRequest) GetCheckWarningStatus() *int32 {
	return s.CheckWarningStatus
}

func (s *ListCheckItemWarningSummaryRequest) GetCheckWarningStatusList() []*int32 {
	return s.CheckWarningStatusList
}

func (s *ListCheckItemWarningSummaryRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *ListCheckItemWarningSummaryRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *ListCheckItemWarningSummaryRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckItemWarningSummaryRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListCheckItemWarningSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckItemWarningSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckItemWarningSummaryRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ListCheckItemWarningSummaryRequest) GetRiskType() *string {
	return s.RiskType
}

func (s *ListCheckItemWarningSummaryRequest) GetSource() *string {
	return s.Source
}

func (s *ListCheckItemWarningSummaryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCheckItemWarningSummaryRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *ListCheckItemWarningSummaryRequest) SetCheckItemFuzzy(v string) *ListCheckItemWarningSummaryRequest {
	s.CheckItemFuzzy = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetCheckLevel(v string) *ListCheckItemWarningSummaryRequest {
	s.CheckLevel = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetCheckType(v string) *ListCheckItemWarningSummaryRequest {
	s.CheckType = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetCheckWarningStatus(v int32) *ListCheckItemWarningSummaryRequest {
	s.CheckWarningStatus = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetCheckWarningStatusList(v []*int32) *ListCheckItemWarningSummaryRequest {
	s.CheckWarningStatusList = v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetContainerFieldName(v string) *ListCheckItemWarningSummaryRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetContainerFieldValue(v string) *ListCheckItemWarningSummaryRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetCurrentPage(v int32) *ListCheckItemWarningSummaryRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetGroupId(v int64) *ListCheckItemWarningSummaryRequest {
	s.GroupId = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetLang(v string) *ListCheckItemWarningSummaryRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetPageSize(v int32) *ListCheckItemWarningSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetResourceDirectoryAccountId(v int64) *ListCheckItemWarningSummaryRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetRiskType(v string) *ListCheckItemWarningSummaryRequest {
	s.RiskType = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetSource(v string) *ListCheckItemWarningSummaryRequest {
	s.Source = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetStartTime(v int64) *ListCheckItemWarningSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) SetUuidList(v []*string) *ListCheckItemWarningSummaryRequest {
	s.UuidList = v
	return s
}

func (s *ListCheckItemWarningSummaryRequest) Validate() error {
	return dara.Validate(s)
}
