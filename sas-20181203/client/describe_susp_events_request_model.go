// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmUniqueInfo(v string) *DescribeSuspEventsRequest
	GetAlarmUniqueInfo() *string
	SetAssetsTypeList(v []*string) *DescribeSuspEventsRequest
	GetAssetsTypeList() []*string
	SetClusterId(v string) *DescribeSuspEventsRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeSuspEventsRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeSuspEventsRequest
	GetContainerFieldValue() *string
	SetCurrentPage(v string) *DescribeSuspEventsRequest
	GetCurrentPage() *string
	SetDealed(v string) *DescribeSuspEventsRequest
	GetDealed() *string
	SetDetectSource(v string) *DescribeSuspEventsRequest
	GetDetectSource() *string
	SetEventNames(v string) *DescribeSuspEventsRequest
	GetEventNames() *string
	SetFrom(v string) *DescribeSuspEventsRequest
	GetFrom() *string
	SetGroupId(v int64) *DescribeSuspEventsRequest
	GetGroupId() *int64
	SetId(v int64) *DescribeSuspEventsRequest
	GetId() *int64
	SetLang(v string) *DescribeSuspEventsRequest
	GetLang() *string
	SetLevels(v string) *DescribeSuspEventsRequest
	GetLevels() *string
	SetMultiAccountActionType(v int32) *DescribeSuspEventsRequest
	GetMultiAccountActionType() *int32
	SetName(v string) *DescribeSuspEventsRequest
	GetName() *string
	SetOperateErrorCodeList(v []*string) *DescribeSuspEventsRequest
	GetOperateErrorCodeList() []*string
	SetOperateTimeEnd(v string) *DescribeSuspEventsRequest
	GetOperateTimeEnd() *string
	SetOperateTimeStart(v string) *DescribeSuspEventsRequest
	GetOperateTimeStart() *string
	SetPageSize(v string) *DescribeSuspEventsRequest
	GetPageSize() *string
	SetParentEventTypes(v string) *DescribeSuspEventsRequest
	GetParentEventTypes() *string
	SetRemark(v string) *DescribeSuspEventsRequest
	GetRemark() *string
	SetResourceDirectoryAccountId(v int64) *DescribeSuspEventsRequest
	GetResourceDirectoryAccountId() *int64
	SetSortColumn(v string) *DescribeSuspEventsRequest
	GetSortColumn() *string
	SetSortType(v string) *DescribeSuspEventsRequest
	GetSortType() *string
	SetSource(v string) *DescribeSuspEventsRequest
	GetSource() *string
	SetSourceAliUids(v []*int64) *DescribeSuspEventsRequest
	GetSourceAliUids() []*int64
	SetSourceIp(v string) *DescribeSuspEventsRequest
	GetSourceIp() *string
	SetStatus(v string) *DescribeSuspEventsRequest
	GetStatus() *string
	SetStrictMode(v string) *DescribeSuspEventsRequest
	GetStrictMode() *string
	SetSupportOperateCodeList(v []*string) *DescribeSuspEventsRequest
	GetSupportOperateCodeList() []*string
	SetTacticId(v string) *DescribeSuspEventsRequest
	GetTacticId() *string
	SetTargetType(v string) *DescribeSuspEventsRequest
	GetTargetType() *string
	SetTimeEnd(v string) *DescribeSuspEventsRequest
	GetTimeEnd() *string
	SetTimeStart(v string) *DescribeSuspEventsRequest
	GetTimeStart() *string
	SetUniqueInfo(v string) *DescribeSuspEventsRequest
	GetUniqueInfo() *string
	SetUuids(v string) *DescribeSuspEventsRequest
	GetUuids() *string
}

type DescribeSuspEventsRequest struct {
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
	SourceAliUids []*int64 `json:"SourceAliUids,omitempty" xml:"SourceAliUids,omitempty" type:"Repeated"`
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

func (s DescribeSuspEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsRequest) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *DescribeSuspEventsRequest) GetAssetsTypeList() []*string {
	return s.AssetsTypeList
}

func (s *DescribeSuspEventsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSuspEventsRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeSuspEventsRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeSuspEventsRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSuspEventsRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeSuspEventsRequest) GetDetectSource() *string {
	return s.DetectSource
}

func (s *DescribeSuspEventsRequest) GetEventNames() *string {
	return s.EventNames
}

func (s *DescribeSuspEventsRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeSuspEventsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeSuspEventsRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeSuspEventsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSuspEventsRequest) GetLevels() *string {
	return s.Levels
}

func (s *DescribeSuspEventsRequest) GetMultiAccountActionType() *int32 {
	return s.MultiAccountActionType
}

func (s *DescribeSuspEventsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSuspEventsRequest) GetOperateErrorCodeList() []*string {
	return s.OperateErrorCodeList
}

func (s *DescribeSuspEventsRequest) GetOperateTimeEnd() *string {
	return s.OperateTimeEnd
}

func (s *DescribeSuspEventsRequest) GetOperateTimeStart() *string {
	return s.OperateTimeStart
}

func (s *DescribeSuspEventsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSuspEventsRequest) GetParentEventTypes() *string {
	return s.ParentEventTypes
}

func (s *DescribeSuspEventsRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeSuspEventsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeSuspEventsRequest) GetSortColumn() *string {
	return s.SortColumn
}

func (s *DescribeSuspEventsRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeSuspEventsRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeSuspEventsRequest) GetSourceAliUids() []*int64 {
	return s.SourceAliUids
}

func (s *DescribeSuspEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSuspEventsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSuspEventsRequest) GetStrictMode() *string {
	return s.StrictMode
}

func (s *DescribeSuspEventsRequest) GetSupportOperateCodeList() []*string {
	return s.SupportOperateCodeList
}

func (s *DescribeSuspEventsRequest) GetTacticId() *string {
	return s.TacticId
}

func (s *DescribeSuspEventsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeSuspEventsRequest) GetTimeEnd() *string {
	return s.TimeEnd
}

func (s *DescribeSuspEventsRequest) GetTimeStart() *string {
	return s.TimeStart
}

func (s *DescribeSuspEventsRequest) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *DescribeSuspEventsRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeSuspEventsRequest) SetAlarmUniqueInfo(v string) *DescribeSuspEventsRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetAssetsTypeList(v []*string) *DescribeSuspEventsRequest {
	s.AssetsTypeList = v
	return s
}

func (s *DescribeSuspEventsRequest) SetClusterId(v string) *DescribeSuspEventsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetContainerFieldName(v string) *DescribeSuspEventsRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetContainerFieldValue(v string) *DescribeSuspEventsRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetCurrentPage(v string) *DescribeSuspEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetDealed(v string) *DescribeSuspEventsRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetDetectSource(v string) *DescribeSuspEventsRequest {
	s.DetectSource = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetEventNames(v string) *DescribeSuspEventsRequest {
	s.EventNames = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetFrom(v string) *DescribeSuspEventsRequest {
	s.From = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetGroupId(v int64) *DescribeSuspEventsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetId(v int64) *DescribeSuspEventsRequest {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLang(v string) *DescribeSuspEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLevels(v string) *DescribeSuspEventsRequest {
	s.Levels = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetMultiAccountActionType(v int32) *DescribeSuspEventsRequest {
	s.MultiAccountActionType = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetName(v string) *DescribeSuspEventsRequest {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetOperateErrorCodeList(v []*string) *DescribeSuspEventsRequest {
	s.OperateErrorCodeList = v
	return s
}

func (s *DescribeSuspEventsRequest) SetOperateTimeEnd(v string) *DescribeSuspEventsRequest {
	s.OperateTimeEnd = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetOperateTimeStart(v string) *DescribeSuspEventsRequest {
	s.OperateTimeStart = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetPageSize(v string) *DescribeSuspEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetParentEventTypes(v string) *DescribeSuspEventsRequest {
	s.ParentEventTypes = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetRemark(v string) *DescribeSuspEventsRequest {
	s.Remark = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetResourceDirectoryAccountId(v int64) *DescribeSuspEventsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetSortColumn(v string) *DescribeSuspEventsRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetSortType(v string) *DescribeSuspEventsRequest {
	s.SortType = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetSource(v string) *DescribeSuspEventsRequest {
	s.Source = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetSourceAliUids(v []*int64) *DescribeSuspEventsRequest {
	s.SourceAliUids = v
	return s
}

func (s *DescribeSuspEventsRequest) SetSourceIp(v string) *DescribeSuspEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetStatus(v string) *DescribeSuspEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetStrictMode(v string) *DescribeSuspEventsRequest {
	s.StrictMode = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetSupportOperateCodeList(v []*string) *DescribeSuspEventsRequest {
	s.SupportOperateCodeList = v
	return s
}

func (s *DescribeSuspEventsRequest) SetTacticId(v string) *DescribeSuspEventsRequest {
	s.TacticId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetTargetType(v string) *DescribeSuspEventsRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetTimeEnd(v string) *DescribeSuspEventsRequest {
	s.TimeEnd = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetTimeStart(v string) *DescribeSuspEventsRequest {
	s.TimeStart = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetUniqueInfo(v string) *DescribeSuspEventsRequest {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetUuids(v string) *DescribeSuspEventsRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeSuspEventsRequest) Validate() error {
	return dara.Validate(s)
}
