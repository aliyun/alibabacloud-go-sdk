// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListWorkflowsResponseBodyPagingInfo) *ListWorkflowsResponseBody
	GetPagingInfo() *ListWorkflowsResponseBodyPagingInfo
	SetRequestId(v string) *ListWorkflowsResponseBody
	GetRequestId() *string
}

type ListWorkflowsResponseBody struct {
	// The pagination information.
	PagingInfo *ListWorkflowsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBody) GetPagingInfo() *ListWorkflowsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListWorkflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowsResponseBody) SetPagingInfo(v *ListWorkflowsResponseBodyPagingInfo) *ListWorkflowsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListWorkflowsResponseBody) SetRequestId(v string) *ListWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowsResponseBodyPagingInfo struct {
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
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The workflows.
	Workflows []*ListWorkflowsResponseBodyPagingInfoWorkflows `json:"Workflows,omitempty" xml:"Workflows,omitempty" type:"Repeated"`
}

func (s ListWorkflowsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkflowsResponseBodyPagingInfo) GetWorkflows() []*ListWorkflowsResponseBodyPagingInfoWorkflows {
	return s.Workflows
}

func (s *ListWorkflowsResponseBodyPagingInfo) SetPageNumber(v int32) *ListWorkflowsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfo) SetPageSize(v int32) *ListWorkflowsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfo) SetTotalCount(v int32) *ListWorkflowsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfo) SetWorkflows(v []*ListWorkflowsResponseBodyPagingInfoWorkflows) *ListWorkflowsResponseBodyPagingInfo {
	s.Workflows = v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfo) Validate() error {
	if s.Workflows != nil {
		for _, item := range s.Workflows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowsResponseBodyPagingInfoWorkflows struct {
	// The unique code of the client. This parameter is used to create a workflow asynchronously and implement the idempotence of the workflow. If you do not specify this parameter when you create the workflow, the system automatically generates a unique code. The unique code is uniquely associated with the workflow ID. If you specify this parameter when you update or delete the workflow, the value of this parameter must be the unique code that is used to create the workflow.
	//
	// example:
	//
	// Workflow_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the creator.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment of the workspace. Valid values:
	//
	// 	- Prod
	//
	// 	- Dev
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account ID of the modifier.
	//
	// example:
	//
	// 1000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name.
	//
	// example:
	//
	// Workflow1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameters.
	//
	// example:
	//
	// para1=$bizdate para2=$[yyyymmdd]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The task tag.
	Tags []*ListWorkflowsResponseBodyPagingInfoWorkflowsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The trigger method.
	Trigger *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListWorkflowsResponseBodyPagingInfoWorkflows) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBodyPagingInfoWorkflows) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetDescription() *string {
	return s.Description
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetEnvType() *string {
	return s.EnvType
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetName() *string {
	return s.Name
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetOwner() *string {
	return s.Owner
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetParameters() *string {
	return s.Parameters
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetTags() []*ListWorkflowsResponseBodyPagingInfoWorkflowsTags {
	return s.Tags
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) GetTrigger() *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger {
	return s.Trigger
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetClientUniqueCode(v string) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.ClientUniqueCode = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetCreateTime(v int64) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetCreateUser(v string) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.CreateUser = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetDescription(v string) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.Description = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetEnvType(v string) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.EnvType = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetId(v int64) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.Id = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetModifyTime(v int64) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.ModifyTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetModifyUser(v string) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.ModifyUser = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetName(v string) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.Name = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetOwner(v string) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.Owner = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetParameters(v string) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.Parameters = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetProjectId(v int64) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetTags(v []*ListWorkflowsResponseBodyPagingInfoWorkflowsTags) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.Tags = v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) SetTrigger(v *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) *ListWorkflowsResponseBodyPagingInfoWorkflows {
	s.Trigger = v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflows) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowsResponseBodyPagingInfoWorkflowsTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListWorkflowsResponseBodyPagingInfoWorkflowsTags) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBodyPagingInfoWorkflowsTags) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTags) GetKey() *string {
	return s.Key
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTags) GetValue() *string {
	return s.Value
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTags) SetKey(v string) *ListWorkflowsResponseBodyPagingInfoWorkflowsTags {
	s.Key = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTags) SetValue(v string) *ListWorkflowsResponseBodyPagingInfoWorkflowsTags {
	s.Value = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTags) Validate() error {
	return dara.Validate(s)
}

type ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger struct {
	// The CRON expression. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the time range during which the workflow is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The running mode of the workflow after it is triggered. This parameter takes effect only if the Type parameter is set to Scheduler. Valid values:
	//
	// 	- Pause
	//
	// 	- Skip
	//
	// 	- Normal
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The start time of the time range during which the workflow is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trigger type. Valid values:
	//
	// 	- Scheduler: scheduling cycle-based trigger
	//
	// 	- Manual: manual trigger
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) GoString() string {
	return s.String()
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) GetType() *string {
	return s.Type
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) SetCron(v string) *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger {
	s.Cron = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) SetEndTime(v string) *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) SetRecurrence(v string) *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger {
	s.Recurrence = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) SetStartTime(v string) *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) SetType(v string) *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger {
	s.Type = &v
	return s
}

func (s *ListWorkflowsResponseBodyPagingInfoWorkflowsTrigger) Validate() error {
	return dara.Validate(s)
}
