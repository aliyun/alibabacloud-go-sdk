// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmEventName(v string) *DescribeSuspEventsRequest
	GetAlarmEventName() *string
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
	SetCurrentPage(v int32) *DescribeSuspEventsRequest
	GetCurrentPage() *int32
	SetDealed(v string) *DescribeSuspEventsRequest
	GetDealed() *string
	SetDepartment(v string) *DescribeSuspEventsRequest
	GetDepartment() *string
	SetEventNames(v string) *DescribeSuspEventsRequest
	GetEventNames() *string
	SetGroupId(v int64) *DescribeSuspEventsRequest
	GetGroupId() *int64
	SetHostName(v string) *DescribeSuspEventsRequest
	GetHostName() *string
	SetId(v int64) *DescribeSuspEventsRequest
	GetId() *int64
	SetInnerIp(v string) *DescribeSuspEventsRequest
	GetInnerIp() *string
	SetLang(v string) *DescribeSuspEventsRequest
	GetLang() *string
	SetLevels(v string) *DescribeSuspEventsRequest
	GetLevels() *string
	SetName(v string) *DescribeSuspEventsRequest
	GetName() *string
	SetOperateErrorCodeList(v []*string) *DescribeSuspEventsRequest
	GetOperateErrorCodeList() []*string
	SetOperateTimeEnd(v string) *DescribeSuspEventsRequest
	GetOperateTimeEnd() *string
	SetOperateTimeStart(v string) *DescribeSuspEventsRequest
	GetOperateTimeStart() *string
	SetPageSize(v int32) *DescribeSuspEventsRequest
	GetPageSize() *int32
	SetParentEventTypes(v string) *DescribeSuspEventsRequest
	GetParentEventTypes() *string
	SetRemark(v string) *DescribeSuspEventsRequest
	GetRemark() *string
	SetSortColumn(v string) *DescribeSuspEventsRequest
	GetSortColumn() *string
	SetSortType(v string) *DescribeSuspEventsRequest
	GetSortType() *string
	SetSource(v string) *DescribeSuspEventsRequest
	GetSource() *string
	SetSourceIp(v string) *DescribeSuspEventsRequest
	GetSourceIp() *string
	SetSrcIp(v string) *DescribeSuspEventsRequest
	GetSrcIp() *string
	SetStatus(v string) *DescribeSuspEventsRequest
	GetStatus() *string
	SetStrictMode(v string) *DescribeSuspEventsRequest
	GetStrictMode() *string
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
	SetUserName(v string) *DescribeSuspEventsRequest
	GetUserName() *string
	SetUuids(v string) *DescribeSuspEventsRequest
	GetUuids() *string
}

type DescribeSuspEventsRequest struct {
	// The name of the alert event.
	//
	// example:
	//
	// Trojan
	AlarmEventName *string `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	// The unique ID of the alert event.
	//
	// To query the exception event information of a single alert event, you must provide the unique ID of the alert event. You can call the DescribeSuspEvents operation to obtain the ID.
	//
	// example:
	//
	// 8df914418f4211fb****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// The collection of asset types. Not supported by SASE-EDR.
	AssetsTypeList []*string `json:"AssetsTypeList,omitempty" xml:"AssetsTypeList,omitempty" type:"Repeated"`
	// The ID of the cluster for which you want to query alert events.
	//
	// example:
	//
	// c4af4fdf38a98496a9b63c2be5dae****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The container search field. Not supported by SASE-EDR.
	//
	// example:
	//
	// instanceId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the container search field. Not supported by SASE-EDR.
	//
	// example:
	//
	// ccf9769c22b844ff9b8d57417683b****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The page number of the returned page. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Filters alerts by processing group. Valid values:
	//
	// - `N`: Pending.
	//
	// - `Y`: Processed.
	//
	// - `A`: All.
	//
	// If this parameter is not specified, all groups are queried. The Status parameter further filters alerts within the selected group. If the specified status does not match the group, an empty result is returned.
	//
	// example:
	//
	// N
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The department to which the user belongs.
	//
	// example:
	//
	// Supplier Operations Team
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The internal subtype identifiers of the alerts. Separate multiple subtypes with commas (,). This parameter is used to match event names. The valid values must belong to the alert configurations visible to the current account and are not the alert names displayed to users.
	//
	// example:
	//
	// SASE non-whitelisted startup interception_PC
	EventNames *string `json:"EventNames,omitempty" xml:"EventNames,omitempty"`
	// The ID of the group to which the asset affected by the alert event belongs.
	//
	// example:
	//
	// 18768
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the endpoint device.
	//
	// example:
	//
	// qin-ubuntu
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The unique ID of the recorded alert event.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The private IP address of the endpoint device.
	//
	// example:
	//
	// 192.168.XXX.XXX
	InnerIp *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The urgency levels of the alert events to query. Separate multiple urgency levels with commas (,). The urgency levels are listed in descending order. Valid values:
	//
	// - **serious**: Urgent.
	//
	// - **suspicious**: Suspicious.
	//
	// - **remind**: Reminder.
	//
	// example:
	//
	// serious
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	// The name of the asset affected by the alert event.
	//
	// example:
	//
	// ecs-xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The collection of alert event processing result codes.
	OperateErrorCodeList []*string `json:"OperateErrorCodeList,omitempty" xml:"OperateErrorCodeList,omitempty" type:"Repeated"`
	// The end time of the alert processing time. Specify a datetime string in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-07-06 13:50:38
	OperateTimeEnd *string `json:"OperateTimeEnd,omitempty" xml:"OperateTimeEnd,omitempty"`
	// The start time of the alert processing time. Specify a datetime string in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-07-05 13:50:38
	OperateTimeStart *string `json:"OperateTimeStart,omitempty" xml:"OperateTimeStart,omitempty"`
	// The number of alert events to return on each page in a paging query. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The internal type identifiers of the alerts. Separate multiple types with commas (,). The valid values are derived from the current alert configurations. You can call the DescribeNsasSuspEventType operation to query the Type field in EventTypes. The Name field is the display name and is not used as a fixed enumeration list.
	//
	// example:
	//
	// Precision defense
	ParentEventTypes *string `json:"ParentEventTypes,omitempty" xml:"ParentEventTypes,omitempty"`
	// The alert name or asset information to query.
	//
	// Note: Fuzzy match is supported. Asset information includes the asset name, public IP address, and private IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The custom sort field. Default value: **operateTime**. Valid values:
	//
	// - **lastTime**: Latest occurrence time.
	//
	// - **operateTime**: Processing time.
	//
	// > This field takes effect only when Dealed is set to Y.
	//
	// example:
	//
	// operateTime
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// The custom sort order. Default value: **desc**. Valid values:
	//
	// - **asc**: Ascending order.
	//
	// - **desc**: Descending order.
	//
	// > This field takes effect only when Dealed is set to Y.
	//
	// example:
	//
	// desc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The compatibility source classification parameter. This parameter is not used to specify the alert data source. You do not need to specify this parameter for common alert list queries.
	//
	// example:
	//
	// virus
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The IP address of the endpoint.
	//
	// example:
	//
	// 111.55.XXX.XXX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// Further filters alerts by individual alert status within the processing group specified by the Dealed parameter. Status descriptions:
	//
	// - `0`: No further status restriction. The alerts are still subject to the Dealed group.
	//
	// - `1`: Pending.
	//
	// - `2`: Ignored.
	//
	// - `4`: Manually processed.
	//
	// - `8`: Whitelisted (marked as false positive).
	//
	// - `32`: Processed. For example, automatic defense interception succeeded.
	//
	// - `64`: Expired.
	//
	// Unprocessed alerts are automatically marked as expired if more than 31 days have passed since the last occurrence.
	//
	// Deleted alerts cannot be queried. If the specified status does not belong to the group specified by Dealed, an empty result is returned.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether the alert is in strict mode. Not supported by SASE-EDR.
	//
	// example:
	//
	// Y
	StrictMode *string `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// The ATT&CK tactic ID.
	//
	// example:
	//
	// TA0001
	TacticId *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	// The container search target type. Not supported by SASE-EDR.
	//
	// example:
	//
	// containerId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The end time of the latest occurrence. Specify a datetime string in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-07-06 13:50:38
	TimeEnd *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	// The start time of the latest occurrence. Specify a datetime string in the YYYY-MM-DD HH:mm:ss format.
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
	// The username of the endpoint.
	//
	// example:
	//
	// User
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *DescribeSuspEventsRequest) GetAlarmEventName() *string {
	return s.AlarmEventName
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

func (s *DescribeSuspEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSuspEventsRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeSuspEventsRequest) GetDepartment() *string {
	return s.Department
}

func (s *DescribeSuspEventsRequest) GetEventNames() *string {
	return s.EventNames
}

func (s *DescribeSuspEventsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeSuspEventsRequest) GetHostName() *string {
	return s.HostName
}

func (s *DescribeSuspEventsRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeSuspEventsRequest) GetInnerIp() *string {
	return s.InnerIp
}

func (s *DescribeSuspEventsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSuspEventsRequest) GetLevels() *string {
	return s.Levels
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

func (s *DescribeSuspEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSuspEventsRequest) GetParentEventTypes() *string {
	return s.ParentEventTypes
}

func (s *DescribeSuspEventsRequest) GetRemark() *string {
	return s.Remark
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

func (s *DescribeSuspEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSuspEventsRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeSuspEventsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSuspEventsRequest) GetStrictMode() *string {
	return s.StrictMode
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

func (s *DescribeSuspEventsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSuspEventsRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeSuspEventsRequest) SetAlarmEventName(v string) *DescribeSuspEventsRequest {
	s.AlarmEventName = &v
	return s
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

func (s *DescribeSuspEventsRequest) SetCurrentPage(v int32) *DescribeSuspEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetDealed(v string) *DescribeSuspEventsRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetDepartment(v string) *DescribeSuspEventsRequest {
	s.Department = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetEventNames(v string) *DescribeSuspEventsRequest {
	s.EventNames = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetGroupId(v int64) *DescribeSuspEventsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetHostName(v string) *DescribeSuspEventsRequest {
	s.HostName = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetId(v int64) *DescribeSuspEventsRequest {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetInnerIp(v string) *DescribeSuspEventsRequest {
	s.InnerIp = &v
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

func (s *DescribeSuspEventsRequest) SetPageSize(v int32) *DescribeSuspEventsRequest {
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

func (s *DescribeSuspEventsRequest) SetSourceIp(v string) *DescribeSuspEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetSrcIp(v string) *DescribeSuspEventsRequest {
	s.SrcIp = &v
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

func (s *DescribeSuspEventsRequest) SetUserName(v string) *DescribeSuspEventsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetUuids(v string) *DescribeSuspEventsRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeSuspEventsRequest) Validate() error {
	return dara.Validate(s)
}
