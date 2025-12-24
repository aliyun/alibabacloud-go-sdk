// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v *DescribeProcessTasksResponseBodyPage) *DescribeProcessTasksResponseBody
	GetPage() *DescribeProcessTasksResponseBodyPage
	SetProcessTasks(v []*DescribeProcessTasksResponseBodyProcessTasks) *DescribeProcessTasksResponseBody
	GetProcessTasks() []*DescribeProcessTasksResponseBodyProcessTasks
	SetRequestId(v string) *DescribeProcessTasksResponseBody
	GetRequestId() *string
}

type DescribeProcessTasksResponseBody struct {
	// The pagination information.
	Page *DescribeProcessTasksResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The handling tasks.
	ProcessTasks []*DescribeProcessTasksResponseBodyProcessTasks `json:"ProcessTasks,omitempty" xml:"ProcessTasks,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E7698CFB-4E1C-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBody) GetPage() *DescribeProcessTasksResponseBodyPage {
	return s.Page
}

func (s *DescribeProcessTasksResponseBody) GetProcessTasks() []*DescribeProcessTasksResponseBodyProcessTasks {
	return s.ProcessTasks
}

func (s *DescribeProcessTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProcessTasksResponseBody) SetPage(v *DescribeProcessTasksResponseBodyPage) *DescribeProcessTasksResponseBody {
	s.Page = v
	return s
}

func (s *DescribeProcessTasksResponseBody) SetProcessTasks(v []*DescribeProcessTasksResponseBodyProcessTasks) *DescribeProcessTasksResponseBody {
	s.ProcessTasks = v
	return s
}

func (s *DescribeProcessTasksResponseBody) SetRequestId(v string) *DescribeProcessTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProcessTasksResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.ProcessTasks != nil {
		for _, item := range s.ProcessTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProcessTasksResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProcessTasksResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessTasksResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeProcessTasksResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProcessTasksResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeProcessTasksResponseBodyPage) SetPageNumber(v int32) *DescribeProcessTasksResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyPage) SetPageSize(v int32) *DescribeProcessTasksResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyPage) SetTotalCount(v int32) *DescribeProcessTasksResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type DescribeProcessTasksResponseBodyProcessTasks struct {
	// The ID of the Alibaba Cloud account that is used to submit the handling task.
	//
	// example:
	//
	// 123xxxx355
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The name of the handling entity.
	//
	// example:
	//
	// 1.1.1.x
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The type of the handling entity.
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
	// The error code returned if the call failed.
	//
	// example:
	//
	// sts_openapi.Info.DefenseSceneNotSupported
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// ParamError : The parameters of your request are invalid
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The error tip returned if the call failed.
	//
	// example:
	//
	// Verify that the input parameters of the components are correct
	ErrTip *string `json:"ErrTip,omitempty" xml:"ErrTip,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// c1020ce1-d6a5-11e8-8298-00163e10****
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The creation time of the handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	GmtCreateMillis *int64 `json:"GmtCreateMillis,omitempty" xml:"GmtCreateMillis,omitempty"`
	// The modification time of the handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	GmtModifiedMillis *int64 `json:"GmtModifiedMillis,omitempty" xml:"GmtModifiedMillis,omitempty"`
	// The input parameter of the handling task.
	//
	// example:
	//
	// {"groupuuid":"c6a9b1df-f4ac-4078-bef4-99xxxxxx"}
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The ID of the associated policy.
	//
	// example:
	//
	// 92af3c79-1754-4646-9366-9ddbd1e45536_xxxx
	ProcessStrategyUuid *string `json:"ProcessStrategyUuid,omitempty" xml:"ProcessStrategyUuid,omitempty"`
	// The delivery time of the handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessTime *int64 `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty"`
	// The unblocking time of the handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	RemoveTime *int64 `json:"RemoveTime,omitempty" xml:"RemoveTime,omitempty"`
	// The UUID of the playbook execution record.
	//
	// example:
	//
	// 93e5df20-3d03-42e1-b44b-58197c71****
	ReqUuid *string `json:"ReqUuid,omitempty" xml:"ReqUuid,omitempty"`
	// The scenario code of the handling task.
	//
	// example:
	//
	// event_xxx_whole_process
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The scenario name of the handling task.
	//
	// example:
	//
	// waf_whole_process
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The ID of the Alibaba Cloud account that is specified in the handling task.
	//
	// example:
	//
	// 123xxxxx234
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The submission source of the handling task.
	//
	// example:
	//
	// system
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique identifier of the handling task.
	//
	// example:
	//
	// 150xxxxxxxxx95066
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The status of the handling task.
	//
	// example:
	//
	// 11
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
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
	// The code of the cloud service that is associated with the handling task.
	//
	// example:
	//
	// WAF
	YunCode *string `json:"YunCode,omitempty" xml:"YunCode,omitempty"`
}

func (s DescribeProcessTasksResponseBodyProcessTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessTasksResponseBodyProcessTasks) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetCreator() *string {
	return s.Creator
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetEntityName() *string {
	return s.EntityName
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetErrTip() *string {
	return s.ErrTip
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetEventUuid() *string {
	return s.EventUuid
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetGmtCreateMillis() *int64 {
	return s.GmtCreateMillis
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetGmtModifiedMillis() *int64 {
	return s.GmtModifiedMillis
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetInputParams() *string {
	return s.InputParams
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetProcessStrategyUuid() *string {
	return s.ProcessStrategyUuid
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetProcessTime() *int64 {
	return s.ProcessTime
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetRemoveTime() *int64 {
	return s.RemoveTime
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetReqUuid() *string {
	return s.ReqUuid
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetScope() *string {
	return s.Scope
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetSource() *string {
	return s.Source
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) GetYunCode() *string {
	return s.YunCode
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetCreator(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Creator = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEntityName(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EntityName = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEntityType(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EntityType = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEntityUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EntityUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetErrCode(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ErrCode = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetErrMsg(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ErrMsg = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetErrTip(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ErrTip = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEventUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EventUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetGmtCreateMillis(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.GmtCreateMillis = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetGmtModifiedMillis(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.GmtModifiedMillis = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetInputParams(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.InputParams = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetProcessStrategyUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ProcessStrategyUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetProcessTime(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ProcessTime = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetRemoveTime(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.RemoveTime = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetReqUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ReqUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSceneCode(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.SceneCode = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSceneName(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.SceneName = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetScope(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Scope = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSource(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Source = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetTaskId(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetTaskStatus(v int32) *DescribeProcessTasksResponseBodyProcessTasks {
	s.TaskStatus = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetTriggerSource(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.TriggerSource = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetYunCode(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.YunCode = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) Validate() error {
	return dara.Validate(s)
}
