// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessTasksRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The sort order. Valid values:
	//
	// 	- **desc*	- (default).
	//
	// 	- **asc**.
	//
	// example:
	//
	// desc
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The name of the handling entity.
	//
	// example:
	//
	// 127.0.0.1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The type of the handling entity. Valid values:
	//
	// 	- **ip**.
	//
	// 	- **file**.
	//
	// 	- **process**.
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The UUID of the handling entity.
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
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The field that you use to sort the result.
	//
	// >  You can obtain the field from the response result.
	//
	// example:
	//
	// gmtCreate
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. If you do not specify the PageSize parameter, 10 entries are returned by default.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The handling entity, handling scenario, or handling parameter that is used for fuzzy match.
	//
	// example:
	//
	// 12.x.x.x
	ParamContent *string `json:"ParamContent,omitempty" xml:"ParamContent,omitempty"`
	// The end of the time range for a handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessActionEnd *int64 `json:"ProcessActionEnd,omitempty" xml:"ProcessActionEnd,omitempty"`
	// The beginning of the time range for a handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessActionStart *int64 `json:"ProcessActionStart,omitempty" xml:"ProcessActionStart,omitempty"`
	// The end of the time range for an unblocking task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessRemoveEnd *int64 `json:"ProcessRemoveEnd,omitempty" xml:"ProcessRemoveEnd,omitempty"`
	// The beginning of the time range for an unblocking task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessRemoveStart *int64 `json:"ProcessRemoveStart,omitempty" xml:"ProcessRemoveStart,omitempty"`
	// The UUID of the handling policy.
	//
	// >  You can call the [ListDisposeStrategy](https://help.aliyun.com/document_detail/2584440.html) operation to query the UUID of the handling policy.
	//
	// example:
	//
	// 92af3c79-1754-4646-9366-9ddbd1e45536_xxxx
	ProcessStrategyUuid *string `json:"ProcessStrategyUuid,omitempty" xml:"ProcessStrategyUuid,omitempty"`
	ReqUuid             *string `json:"ReqUuid,omitempty" xml:"ReqUuid,omitempty"`
	// The scenario code of the handling task.
	//
	// >  You can call the [DescribeEnumItems](~~DescribeEnumItems~~) operation to query the scenario code of the handling task. This parameter is available when you set **EnumType*	- to **process**.
	//
	// example:
	//
	// event_xxx_whole_process
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The ID of the Alibaba Cloud account that is specified in the handling task.
	//
	// example:
	//
	// 125xxxxx9870
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The triggering source of the handling task. The value is a string array. Valid values:
	//
	// 	- **system**: triggered when you manually handle an event.
	//
	// 	- **custom**: triggered by an event based on an automatic response rule.
	//
	// 	- **custom_alert**: triggered by an alert based on an automatic response rule.
	//
	// 	- **soar-manual**: triggered when you use SOAR to manually run a playbook.
	//
	// 	- **soar-mdr**: triggered by Managed Security Service.
	//
	// example:
	//
	// ["system"]
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique identifier of the handling task.
	//
	// >  This parameter is used to query a specific task. You can obtain the value from the response result.
	//
	// example:
	//
	// 150xxxxxxxxx95066
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The status of the handling task. The value is a string. Valid values:
	//
	// 	- **11**: being handled.
	//
	// 	- **21**: being blocked.
	//
	// 	- **22**: being quarantined.
	//
	// 	- **23**: completed.
	//
	// 	- **24**: added to the whitelist.
	//
	// 	- **20**: successful.
	//
	// 	- **90**: failed.
	//
	// 	- **91**: unblocking failed.
	//
	// 	- **92**: restoring quarantined files failed
	//
	// example:
	//
	// ["11","21"]
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The triggering source of the handling task. Valid values:
	//
	// 	- **system**: triggered when you manually handle an event.
	//
	// 	- **custom**: triggered by an event based on an automatic response rule.
	//
	// 	- **custom_alert**: triggered by an alert based on an automatic response rule.
	//
	// 	- **soar-manual**: triggered when you use SOAR to manually run a playbook.
	//
	// 	- **soar-mdr**: triggered by Managed Security Service.
	//
	// example:
	//
	// system
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
	// The cloud service that is associated with the handling task. The value is a string. Valid values:
	//
	// 	- **WAF**: Web Application Firewall (WAF).
	//
	// 	- **CFW**: Cloud Firewall.
	//
	// 	- **Aegis**: Security Center.
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
