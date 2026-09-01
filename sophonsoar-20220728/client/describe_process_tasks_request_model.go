// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v string) *DescribeProcessTasksRequest
	GetAlertId() *string
	SetDirection(v string) *DescribeProcessTasksRequest
	GetDirection() *string
	SetEntityName(v string) *DescribeProcessTasksRequest
	GetEntityName() *string
	SetEntityType(v string) *DescribeProcessTasksRequest
	GetEntityType() *string
	SetEntityUuid(v string) *DescribeProcessTasksRequest
	GetEntityUuid() *string
	SetEventUuid(v string) *DescribeProcessTasksRequest
	GetEventUuid() *string
	SetExecuteUuid(v string) *DescribeProcessTasksRequest
	GetExecuteUuid() *string
	SetOrderField(v string) *DescribeProcessTasksRequest
	GetOrderField() *string
	SetPageNumber(v int64) *DescribeProcessTasksRequest
	GetPageNumber() *int64
	SetPageSize(v int32) *DescribeProcessTasksRequest
	GetPageSize() *int32
	SetParamContent(v string) *DescribeProcessTasksRequest
	GetParamContent() *string
	SetProcessActionEnd(v int64) *DescribeProcessTasksRequest
	GetProcessActionEnd() *int64
	SetProcessActionStart(v int64) *DescribeProcessTasksRequest
	GetProcessActionStart() *int64
	SetProcessRemoveEnd(v int64) *DescribeProcessTasksRequest
	GetProcessRemoveEnd() *int64
	SetProcessRemoveStart(v int64) *DescribeProcessTasksRequest
	GetProcessRemoveStart() *int64
	SetProcessStrategyUuid(v string) *DescribeProcessTasksRequest
	GetProcessStrategyUuid() *string
	SetReqUuid(v string) *DescribeProcessTasksRequest
	GetReqUuid() *string
	SetResponseRuleId(v string) *DescribeProcessTasksRequest
	GetResponseRuleId() *string
	SetSceneCode(v string) *DescribeProcessTasksRequest
	GetSceneCode() *string
	SetScope(v string) *DescribeProcessTasksRequest
	GetScope() *string
	SetSource(v string) *DescribeProcessTasksRequest
	GetSource() *string
	SetTaskId(v string) *DescribeProcessTasksRequest
	GetTaskId() *string
	SetTaskStatus(v string) *DescribeProcessTasksRequest
	GetTaskStatus() *string
	SetTriggerSource(v string) *DescribeProcessTasksRequest
	GetTriggerSource() *string
	SetYunCode(v string) *DescribeProcessTasksRequest
	GetYunCode() *string
}

type DescribeProcessTasksRequest struct {
	AlertId *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The sort direction. Valid values:
	//
	// - **desc**: Descending (default).
	//
	// - **asc**: Ascending.
	//
	// example:
	//
	// desc
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The name of the entity to be disposed.
	//
	// example:
	//
	// 127.0.0.1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The type of the entity to be disposed. Valid values:
	//
	// - **ip**: IP address entity.
	//
	// - **file**: File entity.
	//
	// - **process**: Process entity.
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The UUID of the entity.
	//
	// example:
	//
	// 69d189e2-ec17-4676-a2fe-02969234****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// c1020ce1-d6a5-11e8-8298-00163e10****
	EventUuid   *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	ExecuteUuid *string `json:"ExecuteUuid,omitempty" xml:"ExecuteUuid,omitempty"`
	// The field used to sort the results.
	//
	// > You can obtain the sort field from the response of this operation.
	//
	// example:
	//
	// gmtCreate
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The page number of the page to return. Default value: 1, which indicates the first page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for paging queries. Default value: 20. If the PageSize parameter is left empty, 10 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The fuzzy match content. This parameter queries the entity, disposal scene, and disposal parameter fields.
	//
	// example:
	//
	// 12.x.x.x
	ParamContent *string `json:"ParamContent,omitempty" xml:"ParamContent,omitempty"`
	// The end time of the query range for the disposal time. Format: 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessActionEnd *int64 `json:"ProcessActionEnd,omitempty" xml:"ProcessActionEnd,omitempty"`
	// The start time of the query range for the disposal time. Format: 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessActionStart *int64 `json:"ProcessActionStart,omitempty" xml:"ProcessActionStart,omitempty"`
	// The end time of the query range for the unblocking time. Format: 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessRemoveEnd *int64 `json:"ProcessRemoveEnd,omitempty" xml:"ProcessRemoveEnd,omitempty"`
	// The start time of the query range for the unblocking time. Format: 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessRemoveStart *int64 `json:"ProcessRemoveStart,omitempty" xml:"ProcessRemoveStart,omitempty"`
	// The UUID of the disposal strategy.
	//
	// >You can call the [ListDisposeStrategy](https://help.aliyun.com/document_detail/2584440.html) operation to obtain this parameter.
	//
	// example:
	//
	// 92af3c79-1754-4646-9366-9ddbd1e45536_****
	ProcessStrategyUuid *string `json:"ProcessStrategyUuid,omitempty" xml:"ProcessStrategyUuid,omitempty"`
	// The trigger ID of the playbook.
	//
	// example:
	//
	// b73d0b08-f1bd-4e8f-967a-8e2982c9****
	ReqUuid        *string `json:"ReqUuid,omitempty" xml:"ReqUuid,omitempty"`
	ResponseRuleId *string `json:"ResponseRuleId,omitempty" xml:"ResponseRuleId,omitempty"`
	// The scene code of the disposal task.
	//
	// >You can call the [DescribeEnumItems](~~DescribeEnumItems~~) operation to obtain this parameter.
	//
	// example:
	//
	// event_xxx_whole_process
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The Alibaba Cloud account ID for the disposal.
	//
	// example:
	//
	// 125xxxxx9870
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The trigger source of the disposal task, in array string format. Valid values:
	//
	// - **system**: Triggered by manual event disposal.
	//
	// - **custom**: Triggered by an automatic response rule based on an event.
	//
	// - **custom_alert**: Triggered by an automatic response rule based on an alert.
	//
	// - **soar-manual**: Triggered by manually invoking a SOAR playbook.
	//
	// - **soar-mdr**: Triggered by the Managed Security Service.
	//
	// example:
	//
	// ["system"]
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique identifier of the disposal task.
	//
	// > This parameter is used to query a specific task. You can obtain the value from the response of this operation.
	//
	// example:
	//
	// 150xxxxxxxxx95066
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The status list of the disposal task, in data string format. Valid values:
	//
	// - **11**: Disposing.
	//
	// - **21**: Blocking.
	//
	// - **22**: Isolating.
	//
	// - **23**: Ended.
	//
	// - **24**: Whitelisted.
	//
	// - **20**: Succeeded.
	//
	// - **90**: Failed.
	//
	// - **91**: Unblocking failed.
	//
	// - **92**: Unisolation failed.
	//
	// example:
	//
	// ["11","21"]
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The trigger source of the disposal task. Valid values:
	//
	// - **system**: Triggered by manual event disposal.
	//
	// - **custom**: Triggered by an automatic response rule based on an event.
	//
	// - **custom_alert**: Triggered by an automatic response rule based on an alert.
	//
	// - **soar-manual**: Triggered by manually invoking a SOAR playbook.
	//
	// - **soar-mdr**: Triggered by the Managed Security Service.
	//
	// example:
	//
	// system
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
	// The cloud product associated with the disposal task, in data string format. Valid values:
	//
	// - **WAF**: Web Application Firewall.
	//
	// - **CFW**: Cloud Firewall.
	//
	// - **Aegis**: Security Center.
	//
	// example:
	//
	// ["WAF"]
	YunCode *string `json:"YunCode,omitempty" xml:"YunCode,omitempty"`
}

func (s DescribeProcessTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksRequest) GetAlertId() *string {
	return s.AlertId
}

func (s *DescribeProcessTasksRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeProcessTasksRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *DescribeProcessTasksRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeProcessTasksRequest) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *DescribeProcessTasksRequest) GetEventUuid() *string {
	return s.EventUuid
}

func (s *DescribeProcessTasksRequest) GetExecuteUuid() *string {
	return s.ExecuteUuid
}

func (s *DescribeProcessTasksRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *DescribeProcessTasksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeProcessTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProcessTasksRequest) GetParamContent() *string {
	return s.ParamContent
}

func (s *DescribeProcessTasksRequest) GetProcessActionEnd() *int64 {
	return s.ProcessActionEnd
}

func (s *DescribeProcessTasksRequest) GetProcessActionStart() *int64 {
	return s.ProcessActionStart
}

func (s *DescribeProcessTasksRequest) GetProcessRemoveEnd() *int64 {
	return s.ProcessRemoveEnd
}

func (s *DescribeProcessTasksRequest) GetProcessRemoveStart() *int64 {
	return s.ProcessRemoveStart
}

func (s *DescribeProcessTasksRequest) GetProcessStrategyUuid() *string {
	return s.ProcessStrategyUuid
}

func (s *DescribeProcessTasksRequest) GetReqUuid() *string {
	return s.ReqUuid
}

func (s *DescribeProcessTasksRequest) GetResponseRuleId() *string {
	return s.ResponseRuleId
}

func (s *DescribeProcessTasksRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DescribeProcessTasksRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeProcessTasksRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeProcessTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeProcessTasksRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeProcessTasksRequest) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *DescribeProcessTasksRequest) GetYunCode() *string {
	return s.YunCode
}

func (s *DescribeProcessTasksRequest) SetAlertId(v string) *DescribeProcessTasksRequest {
	s.AlertId = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetDirection(v string) *DescribeProcessTasksRequest {
	s.Direction = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEntityName(v string) *DescribeProcessTasksRequest {
	s.EntityName = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEntityType(v string) *DescribeProcessTasksRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEntityUuid(v string) *DescribeProcessTasksRequest {
	s.EntityUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEventUuid(v string) *DescribeProcessTasksRequest {
	s.EventUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetExecuteUuid(v string) *DescribeProcessTasksRequest {
	s.ExecuteUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetOrderField(v string) *DescribeProcessTasksRequest {
	s.OrderField = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetPageNumber(v int64) *DescribeProcessTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetPageSize(v int32) *DescribeProcessTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetParamContent(v string) *DescribeProcessTasksRequest {
	s.ParamContent = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessActionEnd(v int64) *DescribeProcessTasksRequest {
	s.ProcessActionEnd = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessActionStart(v int64) *DescribeProcessTasksRequest {
	s.ProcessActionStart = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessRemoveEnd(v int64) *DescribeProcessTasksRequest {
	s.ProcessRemoveEnd = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessRemoveStart(v int64) *DescribeProcessTasksRequest {
	s.ProcessRemoveStart = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessStrategyUuid(v string) *DescribeProcessTasksRequest {
	s.ProcessStrategyUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetReqUuid(v string) *DescribeProcessTasksRequest {
	s.ReqUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetResponseRuleId(v string) *DescribeProcessTasksRequest {
	s.ResponseRuleId = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetSceneCode(v string) *DescribeProcessTasksRequest {
	s.SceneCode = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetScope(v string) *DescribeProcessTasksRequest {
	s.Scope = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetSource(v string) *DescribeProcessTasksRequest {
	s.Source = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetTaskId(v string) *DescribeProcessTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetTaskStatus(v string) *DescribeProcessTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetTriggerSource(v string) *DescribeProcessTasksRequest {
	s.TriggerSource = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetYunCode(v string) *DescribeProcessTasksRequest {
	s.YunCode = &v
	return s
}

func (s *DescribeProcessTasksRequest) Validate() error {
	return dara.Validate(s)
}
