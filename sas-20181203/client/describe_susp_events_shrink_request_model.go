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
	// The ID of the alert event.
	//
	// > To query the details of an alert event, you must specify the ID of the alert event. You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to query the IDs of alert events.
	//
	// example:
	//
	// 8df914418f4211fb****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// The types of the assets.
	AssetsTypeList []*string `json:"AssetsTypeList,omitempty" xml:"AssetsTypeList,omitempty" type:"Repeated"`
	// The ID of the cluster of whose alert events you want to query.
	//
	// example:
	//
	// c4af4fdf38a98496a9b63c2be5dae****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The key of the condition that is used to query alert events on containers. Valid values:
	//
	// 	- **instanceId**: the ID of the asset
	//
	// 	- **appName**: the name of the application
	//
	// 	- **clusterId**: the ID of the cluster
	//
	// 	- **regionId**: the ID of the region
	//
	// 	- **nodeName**: the name of the node
	//
	// 	- **namespace**: the namespace
	//
	// 	- **clusterName**: the name of the cluster
	//
	// 	- **image**: the name of the image
	//
	// 	- **imageRepoName**: the name of the image repository
	//
	// 	- **imageRepoNamespace**: the namespace to which the image repository belongs
	//
	// 	- **imageRepoTag**: the tag that is added to the image
	//
	// 	- **imageDigest**: the digest of the image
	//
	// example:
	//
	// instanceId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the condition that is used to query alert events on containers.
	//
	// example:
	//
	// ccf9769c22b844ff9b8d57417683b****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the alert event is handled. Valid values:
	//
	// 	- **N**: unhandled
	//
	// 	- **Y**: handled
	//
	// example:
	//
	// N
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The subtype of the alert event. Separate multiple subtypes with commas (,).
	//
	// example:
	//
	// WEBSHELL
	EventNames *string `json:"EventNames,omitempty" xml:"EventNames,omitempty"`
	// The data source of the alert event. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the asset group to which the affected asset belongs.
	//
	// example:
	//
	// 18768
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the alert event.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The severity of the alert event. Separate multiple severities with commas (,). Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// The type of the accounts that you want to query. Default value: **0**. Valid values:
	//
	// 	- **0**: the current account.
	//
	// 	- **1**: all accounts.
	//
	// example:
	//
	// 0
	MultiAccountActionType *int32 `json:"MultiAccountActionType,omitempty" xml:"MultiAccountActionType,omitempty"`
	// The name of the asset that is affected by the alert event.
	//
	// example:
	//
	// ecs-xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array that consists of the handling result codes of alert events.
	OperateErrorCodeList []*string `json:"OperateErrorCodeList,omitempty" xml:"OperateErrorCodeList,omitempty" type:"Repeated"`
	// The timestamp when the handling operation ends.
	//
	// example:
	//
	// 2022-07-06 13:50:38
	OperateTimeEnd *string `json:"OperateTimeEnd,omitempty" xml:"OperateTimeEnd,omitempty"`
	// The timestamp when the handling operation starts.
	//
	// example:
	//
	// 2022-07-05 13:50:38
	OperateTimeStart *string `json:"OperateTimeStart,omitempty" xml:"OperateTimeStart,omitempty"`
	// The number of entries per page. Default value: **20**. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The alert type of the alert event. Valid values:
	//
	// 	- **Suspicious process**
	//
	// 	- **Webshell**
	//
	// 	- **Unusual logon**
	//
	// 	- **Exception**
	//
	// 	- **Sensitive file tampering**
	//
	// 	- **Malicious process (cloud threat detection)**
	//
	// 	- **Suspicious network connection**
	//
	// 	- **Suspicious account**
	//
	// 	- **Application intrusion event**
	//
	// 	- **Cloud threat detection**
	//
	// 	- **Precise defense**
	//
	// 	- **Application whitelist**
	//
	// 	- **Persistent webshell**
	//
	// 	- **Web application threat detection**
	//
	// 	- **Malicious script**
	//
	// 	- **Threat intelligence**
	//
	// 	- **Malicious network activity**
	//
	// 	- **Cluster exception**
	//
	// 	- **Webshell (on-premises threat detection)**
	//
	// 	- **Vulnerability exploitation**
	//
	// 	- **Malicious process (on-premises threat detection)**
	//
	// 	- **Trusted exception**
	//
	// 	- **Others**
	//
	// example:
	//
	// Webshell
	ParentEventTypes *string `json:"ParentEventTypes,omitempty" xml:"ParentEventTypes,omitempty"`
	// The name of the alert or the information about the asset.
	//
	// >  Fuzzy search is supported. The asset information includes the name, public IP address, and private IP address of an asset.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the ID.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The custom sorting field. Default value: **operateTime**. Valid values:
	//
	// 	- **lastTime**: the latest occurrence time.
	//
	// 	- **operateTime**: the handling time.
	//
	// >  This parameter takes effect if you set the **Dealed*	- parameter to Y.
	//
	// example:
	//
	// operateTime
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// The custom sorting order. Default value: **desc**. Valid values:
	//
	// 	- **asc**: the ascending order
	//
	// 	- **desc**: the descending order
	//
	// >  This parameter takes effect if you set the **Dealed*	- parameter to Y.
	//
	// example:
	//
	// desc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The source of the alert.
	//
	// example:
	//
	// aegis_suspicious_file_v2
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The IDs of the Alibaba Cloud accounts within which alerts are generated.
	SourceAliUidsShrink *string `json:"SourceAliUids,omitempty" xml:"SourceAliUids,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The status of the alert event. Valid values:
	//
	// 	- **0**: all
	//
	// 	- **1**: pending handling
	//
	// 	- **2**: ignored
	//
	// 	- **4**: confirmed
	//
	// 	- **8**: marked as a false positive
	//
	// 	- **16**: handling
	//
	// 	- **32**: handled
	//
	// 	- **64**: expired
	//
	// 	- **128**: deleted
	//
	// 	- **512**: automatically blocking
	//
	// 	- **513**: automatically blocked
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether to enable the strict alerting mode.
	//
	// 	- N: no
	//
	// 	- Y: Yes
	//
	// example:
	//
	// Y
	StrictMode *string `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// 告警支持的操作类型列表。
	SupportOperateCodeList []*string `json:"SupportOperateCodeList,omitempty" xml:"SupportOperateCodeList,omitempty" type:"Repeated"`
	// The tactic ID of ATT\\&CK.
	//
	// example:
	//
	// TA0001
	TacticId *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	// The item that is used to search for the container. Valid values:
	//
	// 	- **containerId**: the ID of the container
	//
	// 	- **uuid**: the UUID of the server
	//
	// 	- **imageUuid**: the UUID of the image
	//
	// example:
	//
	// containerId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The end time when the alert event was last detected.
	//
	// example:
	//
	// 2022-07-06 13:50:38
	TimeEnd *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	// The start time when the alert event was last detected.
	//
	// example:
	//
	// 2022-07-05 13:50:38
	TimeStart *string `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
	// The unique key of the alert.
	//
	// example:
	//
	// 73fc06fb175a7405697e402f52864****
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	// The UUID of the server on which the alert is detected. Separate multiple UUIDs with commas (,).
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
