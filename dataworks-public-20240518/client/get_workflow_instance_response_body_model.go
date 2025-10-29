// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkflowInstanceResponseBody
	GetRequestId() *string
	SetWorkflowInstance(v *GetWorkflowInstanceResponseBodyWorkflowInstance) *GetWorkflowInstanceResponseBody
	GetWorkflowInstance() *GetWorkflowInstanceResponseBodyWorkflowInstance
}

type GetWorkflowInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the workflow instance.
	WorkflowInstance *GetWorkflowInstanceResponseBodyWorkflowInstance `json:"WorkflowInstance,omitempty" xml:"WorkflowInstance,omitempty" type:"Struct"`
}

func (s GetWorkflowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowInstanceResponseBody) GetWorkflowInstance() *GetWorkflowInstanceResponseBodyWorkflowInstance {
	return s.WorkflowInstance
}

func (s *GetWorkflowInstanceResponseBody) SetRequestId(v string) *GetWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetWorkflowInstance(v *GetWorkflowInstanceResponseBodyWorkflowInstance) *GetWorkflowInstanceResponseBody {
	s.WorkflowInstance = v
	return s
}

func (s *GetWorkflowInstanceResponseBody) Validate() error {
	if s.WorkflowInstance != nil {
		if err := s.WorkflowInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowInstanceResponseBodyWorkflowInstance struct {
	// The data timestamp.
	//
	// example:
	//
	// 1710239005403
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
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
	// The time when the instance finished running.
	//
	// example:
	//
	// 1710239005403
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The ID of the workflow instance.
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
	// The name of the workflow instance.
	//
	// example:
	//
	// WorkInstance1
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The time when the instance started to run.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the workflow instance. Valid values:
	//
	// 	- NotRun: The instance is not run.
	//
	// 	- Running: The instance is running.
	//
	// 	- WaitTime: The instance is waiting for the scheduling time to arrive.
	//
	// 	- CheckingCondition: Branch conditions are being checked for the instance.
	//
	// 	- WaitResource: The instance is waiting for resources.
	//
	// 	- Failure: The instance fails to be run.
	//
	// 	- Success: The instance is successfully run.
	//
	// 	- Checking: Data quality is being checked for the instance.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task tag.
	Tags []*GetWorkflowInstanceResponseBodyWorkflowInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the workflow instance. Valid values:
	//
	// 	- Normal
	//
	// 	- Manual
	//
	// 	- SmokeTest
	//
	// 	- SupplementData
	//
	// 	- ManualWorkflow
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workflow to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow parameters.
	WorkflowParameters     *string `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty"`
	WorkflowTaskInstanceId *int64  `json:"WorkflowTaskInstanceId,omitempty" xml:"WorkflowTaskInstanceId,omitempty"`
}

func (s GetWorkflowInstanceResponseBodyWorkflowInstance) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyWorkflowInstance) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetBizDate() *int64 {
	return s.BizDate
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetEnvType() *string {
	return s.EnvType
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetId() *int64 {
	return s.Id
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetName() *string {
	return s.Name
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetOwner() *string {
	return s.Owner
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetStatus() *string {
	return s.Status
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetTags() []*GetWorkflowInstanceResponseBodyWorkflowInstanceTags {
	return s.Tags
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetType() *string {
	return s.Type
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetWorkflowParameters() *string {
	return s.WorkflowParameters
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) GetWorkflowTaskInstanceId() *int64 {
	return s.WorkflowTaskInstanceId
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetBizDate(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.BizDate = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetCreateTime(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetCreateUser(v string) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.CreateUser = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetEnvType(v string) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.EnvType = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetFinishedTime(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.FinishedTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetId(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.Id = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetModifyTime(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.ModifyTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetModifyUser(v string) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.ModifyUser = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetName(v string) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.Name = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetOwner(v string) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.Owner = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetProjectId(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.ProjectId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetStartedTime(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.StartedTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetStatus(v string) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.Status = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetTags(v []*GetWorkflowInstanceResponseBodyWorkflowInstanceTags) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.Tags = v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetType(v string) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.Type = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetWorkflowId(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetWorkflowParameters(v string) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.WorkflowParameters = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) SetWorkflowTaskInstanceId(v int64) *GetWorkflowInstanceResponseBodyWorkflowInstance {
	s.WorkflowTaskInstanceId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstance) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkflowInstanceResponseBodyWorkflowInstanceTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetWorkflowInstanceResponseBodyWorkflowInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyWorkflowInstanceTags) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstanceTags) GetKey() *string {
	return s.Key
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstanceTags) GetValue() *string {
	return s.Value
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstanceTags) SetKey(v string) *GetWorkflowInstanceResponseBodyWorkflowInstanceTags {
	s.Key = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstanceTags) SetValue(v string) *GetWorkflowInstanceResponseBodyWorkflowInstanceTags {
	s.Value = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyWorkflowInstanceTags) Validate() error {
	return dara.Validate(s)
}
