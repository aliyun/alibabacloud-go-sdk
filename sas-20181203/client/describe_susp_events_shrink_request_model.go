// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmUniqueInfo(v string) *DescribeSuspEventsShrinkRequest
	GetAlarmUniqueInfo() *string
	SetAssetsTypeList(v []*string) *DescribeSuspEventsShrinkRequest
	GetAssetsTypeList() []*string
	SetClusterId(v string) *DescribeSuspEventsShrinkRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeSuspEventsShrinkRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeSuspEventsShrinkRequest
	GetContainerFieldValue() *string
	SetCurrentPage(v string) *DescribeSuspEventsShrinkRequest
	GetCurrentPage() *string
	SetDealed(v string) *DescribeSuspEventsShrinkRequest
	GetDealed() *string
	SetDetectSource(v string) *DescribeSuspEventsShrinkRequest
	GetDetectSource() *string
	SetEventNames(v string) *DescribeSuspEventsShrinkRequest
	GetEventNames() *string
	SetFrom(v string) *DescribeSuspEventsShrinkRequest
	GetFrom() *string
	SetGroupId(v int64) *DescribeSuspEventsShrinkRequest
	GetGroupId() *int64
	SetId(v int64) *DescribeSuspEventsShrinkRequest
	GetId() *int64
	SetLang(v string) *DescribeSuspEventsShrinkRequest
	GetLang() *string
	SetLevels(v string) *DescribeSuspEventsShrinkRequest
	GetLevels() *string
	SetMultiAccountActionType(v int32) *DescribeSuspEventsShrinkRequest
	GetMultiAccountActionType() *int32
	SetName(v string) *DescribeSuspEventsShrinkRequest
	GetName() *string
	SetOperateErrorCodeList(v []*string) *DescribeSuspEventsShrinkRequest
	GetOperateErrorCodeList() []*string
	SetOperateTimeEnd(v string) *DescribeSuspEventsShrinkRequest
	GetOperateTimeEnd() *string
	SetOperateTimeStart(v string) *DescribeSuspEventsShrinkRequest
	GetOperateTimeStart() *string
	SetPageSize(v string) *DescribeSuspEventsShrinkRequest
	GetPageSize() *string
	SetParentEventTypes(v string) *DescribeSuspEventsShrinkRequest
	GetParentEventTypes() *string
	SetRemark(v string) *DescribeSuspEventsShrinkRequest
	GetRemark() *string
	SetResourceDirectoryAccountId(v int64) *DescribeSuspEventsShrinkRequest
	GetResourceDirectoryAccountId() *int64
	SetSortColumn(v string) *DescribeSuspEventsShrinkRequest
	GetSortColumn() *string
	SetSortType(v string) *DescribeSuspEventsShrinkRequest
	GetSortType() *string
	SetSource(v string) *DescribeSuspEventsShrinkRequest
	GetSource() *string
	SetSourceAliUidsShrink(v string) *DescribeSuspEventsShrinkRequest
	GetSourceAliUidsShrink() *string
	SetSourceIp(v string) *DescribeSuspEventsShrinkRequest
	GetSourceIp() *string
	SetStatus(v string) *DescribeSuspEventsShrinkRequest
	GetStatus() *string
	SetStrictMode(v string) *DescribeSuspEventsShrinkRequest
	GetStrictMode() *string
	SetSupportOperateCodeList(v []*string) *DescribeSuspEventsShrinkRequest
	GetSupportOperateCodeList() []*string
	SetTacticId(v string) *DescribeSuspEventsShrinkRequest
	GetTacticId() *string
	SetTargetType(v string) *DescribeSuspEventsShrinkRequest
	GetTargetType() *string
	SetTimeEnd(v string) *DescribeSuspEventsShrinkRequest
	GetTimeEnd() *string
	SetTimeStart(v string) *DescribeSuspEventsShrinkRequest
	GetTimeStart() *string
	SetUniqueInfo(v string) *DescribeSuspEventsShrinkRequest
	GetUniqueInfo() *string
	SetUuids(v string) *DescribeSuspEventsShrinkRequest
	GetUuids() *string
}

type DescribeSuspEventsShrinkRequest struct {
	// The unique ID of the alert event.
	//
	// > To query the exception information of a single alert event, provide the unique ID of the alert event. You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to obtain the ID.
	//
	// example:
	//
	// 8df914418f4211fb****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// The collection of asset types.
	AssetsTypeList []*string `json:"AssetsTypeList,omitempty" xml:"AssetsTypeList,omitempty" type:"Repeated"`
	// The ID of the cluster for which you want to query alert events.
	//
	// example:
	//
	// c4af4fdf38a98496a9b63c2be5dae****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The container search field. Valid values:
	//
	// - **instanceId**: instance ID
	//
	// - **appName**: application name
	//
	// - **clusterId**: cluster ID
	//
	// - **regionId**: region
	//
	// - **nodeName**: node name
	//
	// - **namespace**: namespace
	//
	// - **clusterName**: cluster name
	//
	// - **image**: image name
	//
	// - **imageRepoName**: image repository name
	//
	// - **imageRepoNamespace**: image repository namespace
	//
	// - **imageRepoTag**: image tag
	//
	// - **imageDigest**: image digest
	//
	// example:
	//
	// instanceId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the container search field.
	//
	// example:
	//
	// ccf9769c22b844ff9b8d57417683b****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The page number of the results to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the alert events to query have been handled. Valid values:
	//
	// example:
	//
	// N
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The discovery source. This parameter is invalid.
	//
	// example:
	//
	// linux
	DetectSource *string `json:"DetectSource,omitempty" xml:"DetectSource,omitempty"`
	// The subtypes of the alert events. Separate multiple subtypes with commas (,).
	//
	// example:
	//
	// WEBSHELL
	EventNames *string `json:"EventNames,omitempty" xml:"EventNames,omitempty"`
	// The data source identifier of the alert event. The value is fixed as sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The group ID of the asset affected by the alert event.
	//
	// example:
	//
	// 18768
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The unique ID that identifies the alert event record.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The severity levels of the security alerts that you want to query. Separate multiple severity levels with commas (,). The severity levels are listed in descending order. Valid values:
	//
	// - **serious**: Critical.
	//
	// - **suspicious**: Suspicious.
	//
	// - **remind**: Informational.
	//
	// example:
	//
	// serious
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// The multi-account query type. Default value: **0**. Valid values:
	//
	// - **0**: Queries data of the current account.
	//
	// - **1**: Queries data of all accounts.
	//
	// example:
	//
	// 0
	MultiAccountActionType *int32 `json:"MultiAccountActionType,omitempty" xml:"MultiAccountActionType,omitempty"`
	// The name of the asset affected by the alert event.
	//
	// example:
	//
	// ecs-xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The collection of alert event handling result codes.
	OperateErrorCodeList []*string `json:"OperateErrorCodeList,omitempty" xml:"OperateErrorCodeList,omitempty" type:"Repeated"`
	// The end timestamp of the handling time.
	//
	// example:
	//
	// 2022-07-06 13:50:38
	OperateTimeEnd *string `json:"OperateTimeEnd,omitempty" xml:"OperateTimeEnd,omitempty"`
	// The start timestamp of the handling time.
	//
	// example:
	//
	// 2022-07-05 13:50:38
	OperateTimeStart *string `json:"OperateTimeStart,omitempty" xml:"OperateTimeStart,omitempty"`
	// The number of alert events to display on each page in a paged query. Default value: **20**. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alarm Metric of the alerting events to query. Valid values:
	//
	// example:
	//
	// other
	ParentEventTypes *string `json:"ParentEventTypes,omitempty" xml:"ParentEventTypes,omitempty"`
	// The alert name or asset information to query.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The China site (Chinese mainland) account ID of the member account in the resource directory.
	//
	// >Call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The custom sort field. Default value: **operateTime**. Valid values:
	//
	// - **lastTime**: the most recent occurrence time.
	//
	// - **operateTime**: the processing time.
	//
	// > This field takes effect only when **Dealed*	- is set to Y.
	//
	// example:
	//
	// operateTime
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// The custom sort type. Default value: **desc**. Valid values:
	//
	// - **asc**: ascending order.
	//
	// - **desc**: descending order.
	//
	// > This parameter takes effect only when **Dealed*	- is set to Y.
	//
	// example:
	//
	// desc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The alert source.
	//
	// example:
	//
	// aegis_suspicious_file_v2
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of Alibaba Cloud account IDs that generated the alerts.
	SourceAliUidsShrink *string `json:"SourceAliUids,omitempty" xml:"SourceAliUids,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The status of the alert events to query. Valid values:
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether the alert is identified in strict mode.
	//
	// example:
	//
	// Y
	StrictMode *string `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// The list of operation types supported by the alert.
	SupportOperateCodeList []*string `json:"SupportOperateCodeList,omitempty" xml:"SupportOperateCodeList,omitempty" type:"Repeated"`
	// The tactic ID in ATT&CK.
	//
	// example:
	//
	// TA0001
	TacticId *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	// The type of the container search target. Valid values:
	//
	// - **containerId**: container ID.
	//
	// - **uuid**: server UUID.
	//
	// - **imageUuid**: image UUID.
	//
	// example:
	//
	// containerId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The end time of the latest occurrence time range.
	//
	// example:
	//
	// 2022-07-06 13:50:38
	TimeEnd *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	// The start time of the latest occurrence time range.
	//
	// example:
	//
	// 2022-07-05 13:50:38
	TimeStart *string `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
	// The unique key of the security alert.
	//
	// example:
	//
	// 73fc06fb175a7405697e402f52864****
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	// The UUIDs of the servers for which you want to query alerts. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// bb5d2484-f10e-450d-8917-3e79667e****,0e7c2fcd-7100-42c7-a21a-db6e4f32****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeSuspEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsShrinkRequest) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *DescribeSuspEventsShrinkRequest) GetAssetsTypeList() []*string {
	return s.AssetsTypeList
}

func (s *DescribeSuspEventsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSuspEventsShrinkRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeSuspEventsShrinkRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeSuspEventsShrinkRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSuspEventsShrinkRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeSuspEventsShrinkRequest) GetDetectSource() *string {
	return s.DetectSource
}

func (s *DescribeSuspEventsShrinkRequest) GetEventNames() *string {
	return s.EventNames
}

func (s *DescribeSuspEventsShrinkRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeSuspEventsShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeSuspEventsShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeSuspEventsShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSuspEventsShrinkRequest) GetLevels() *string {
	return s.Levels
}

func (s *DescribeSuspEventsShrinkRequest) GetMultiAccountActionType() *int32 {
	return s.MultiAccountActionType
}

func (s *DescribeSuspEventsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSuspEventsShrinkRequest) GetOperateErrorCodeList() []*string {
	return s.OperateErrorCodeList
}

func (s *DescribeSuspEventsShrinkRequest) GetOperateTimeEnd() *string {
	return s.OperateTimeEnd
}

func (s *DescribeSuspEventsShrinkRequest) GetOperateTimeStart() *string {
	return s.OperateTimeStart
}

func (s *DescribeSuspEventsShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSuspEventsShrinkRequest) GetParentEventTypes() *string {
	return s.ParentEventTypes
}

func (s *DescribeSuspEventsShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeSuspEventsShrinkRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeSuspEventsShrinkRequest) GetSortColumn() *string {
	return s.SortColumn
}

func (s *DescribeSuspEventsShrinkRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeSuspEventsShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeSuspEventsShrinkRequest) GetSourceAliUidsShrink() *string {
	return s.SourceAliUidsShrink
}

func (s *DescribeSuspEventsShrinkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSuspEventsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSuspEventsShrinkRequest) GetStrictMode() *string {
	return s.StrictMode
}

func (s *DescribeSuspEventsShrinkRequest) GetSupportOperateCodeList() []*string {
	return s.SupportOperateCodeList
}

func (s *DescribeSuspEventsShrinkRequest) GetTacticId() *string {
	return s.TacticId
}

func (s *DescribeSuspEventsShrinkRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeSuspEventsShrinkRequest) GetTimeEnd() *string {
	return s.TimeEnd
}

func (s *DescribeSuspEventsShrinkRequest) GetTimeStart() *string {
	return s.TimeStart
}

func (s *DescribeSuspEventsShrinkRequest) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *DescribeSuspEventsShrinkRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeSuspEventsShrinkRequest) SetAlarmUniqueInfo(v string) *DescribeSuspEventsShrinkRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetAssetsTypeList(v []*string) *DescribeSuspEventsShrinkRequest {
	s.AssetsTypeList = v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetClusterId(v string) *DescribeSuspEventsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetContainerFieldName(v string) *DescribeSuspEventsShrinkRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetContainerFieldValue(v string) *DescribeSuspEventsShrinkRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetCurrentPage(v string) *DescribeSuspEventsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetDealed(v string) *DescribeSuspEventsShrinkRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetDetectSource(v string) *DescribeSuspEventsShrinkRequest {
	s.DetectSource = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetEventNames(v string) *DescribeSuspEventsShrinkRequest {
	s.EventNames = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetFrom(v string) *DescribeSuspEventsShrinkRequest {
	s.From = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetGroupId(v int64) *DescribeSuspEventsShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetId(v int64) *DescribeSuspEventsShrinkRequest {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetLang(v string) *DescribeSuspEventsShrinkRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetLevels(v string) *DescribeSuspEventsShrinkRequest {
	s.Levels = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetMultiAccountActionType(v int32) *DescribeSuspEventsShrinkRequest {
	s.MultiAccountActionType = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetName(v string) *DescribeSuspEventsShrinkRequest {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetOperateErrorCodeList(v []*string) *DescribeSuspEventsShrinkRequest {
	s.OperateErrorCodeList = v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetOperateTimeEnd(v string) *DescribeSuspEventsShrinkRequest {
	s.OperateTimeEnd = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetOperateTimeStart(v string) *DescribeSuspEventsShrinkRequest {
	s.OperateTimeStart = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetPageSize(v string) *DescribeSuspEventsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetParentEventTypes(v string) *DescribeSuspEventsShrinkRequest {
	s.ParentEventTypes = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetRemark(v string) *DescribeSuspEventsShrinkRequest {
	s.Remark = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetResourceDirectoryAccountId(v int64) *DescribeSuspEventsShrinkRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetSortColumn(v string) *DescribeSuspEventsShrinkRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetSortType(v string) *DescribeSuspEventsShrinkRequest {
	s.SortType = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetSource(v string) *DescribeSuspEventsShrinkRequest {
	s.Source = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetSourceAliUidsShrink(v string) *DescribeSuspEventsShrinkRequest {
	s.SourceAliUidsShrink = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetSourceIp(v string) *DescribeSuspEventsShrinkRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetStatus(v string) *DescribeSuspEventsShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetStrictMode(v string) *DescribeSuspEventsShrinkRequest {
	s.StrictMode = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetSupportOperateCodeList(v []*string) *DescribeSuspEventsShrinkRequest {
	s.SupportOperateCodeList = v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetTacticId(v string) *DescribeSuspEventsShrinkRequest {
	s.TacticId = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetTargetType(v string) *DescribeSuspEventsShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetTimeEnd(v string) *DescribeSuspEventsShrinkRequest {
	s.TimeEnd = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetTimeStart(v string) *DescribeSuspEventsShrinkRequest {
	s.TimeStart = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetUniqueInfo(v string) *DescribeSuspEventsShrinkRequest {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetUuids(v string) *DescribeSuspEventsShrinkRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
