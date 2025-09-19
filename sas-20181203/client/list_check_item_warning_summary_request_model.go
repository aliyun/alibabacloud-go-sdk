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
	// The name of the check item. Fuzzy match is supported.
	//
	// example:
	//
	// password
	CheckItemFuzzy *string `json:"CheckItemFuzzy,omitempty" xml:"CheckItemFuzzy,omitempty"`
	// The risk level. Default value: null, which indicates that check items at all risk levels are queried.Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	CheckLevel *string `json:"CheckLevel,omitempty" xml:"CheckLevel,omitempty"`
	// The type of the check item.
	//
	// example:
	//
	// hc.check.type.attack_defense
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The risk status. Default value is null, meaning check items in all states are queried. Valid values:
	//
	// 	- **1**: failed
	//
	// 	- **3**: passed
	//
	// 	- **6**: whitelisted
	//
	// example:
	//
	// 3
	CheckWarningStatus *int32 `json:"CheckWarningStatus,omitempty" xml:"CheckWarningStatus,omitempty"`
	// The list of risk levels. If the CheckWarningStatus parameter is specified, only it takes effect.
	CheckWarningStatusList []*int32 `json:"CheckWarningStatusList,omitempty" xml:"CheckWarningStatusList,omitempty" type:"Repeated"`
	// The name of the field that is used to query containers.
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the field that is used to query containers.
	//
	// example:
	//
	// c471f0f61b9c04f8380556e922cf1****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the asset group.
	//
	// > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of asset groups.
	//
	// example:
	//
	// 1161****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the baseline.
	//
	// example:
	//
	// weak_password
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// The data source. Default value: **default**. Valid value:
	//
	// 	- **agentless**: The check items of baselines for agentless detection.
	//
	// 	- **default**: The check items of baselines for hosts.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Start of time range for filtering alerts, effective only for querying historically handled alerts.
	//
	// example:
	//
	// 1732793158366
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The UUIDs of the servers.
	//
	// >  You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of the servers.
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
