// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListWorkflowInstancesResponseBodyPagingInfo) *ListWorkflowInstancesResponseBody
	GetPagingInfo() *ListWorkflowInstancesResponseBodyPagingInfo
	SetRequestId(v string) *ListWorkflowInstancesResponseBody
	GetRequestId() *string
}

type ListWorkflowInstancesResponseBody struct {
	// The pagination information.
	PagingInfo *ListWorkflowInstancesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWorkflowInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBody) GetPagingInfo() *ListWorkflowInstancesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListWorkflowInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowInstancesResponseBody) SetPagingInfo(v *ListWorkflowInstancesResponseBodyPagingInfo) *ListWorkflowInstancesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetRequestId(v string) *ListWorkflowInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowInstancesResponseBodyPagingInfo struct {
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
	// The workflow instances.
	WorkflowInstances []*ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances `json:"WorkflowInstances,omitempty" xml:"WorkflowInstances,omitempty" type:"Repeated"`
}

func (s ListWorkflowInstancesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) GetWorkflowInstances() []*ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	return s.WorkflowInstances
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) SetPageNumber(v int32) *ListWorkflowInstancesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) SetPageSize(v int32) *ListWorkflowInstancesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) SetTotalCount(v int32) *ListWorkflowInstancesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) SetWorkflowInstances(v []*ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) *ListWorkflowInstancesResponseBodyPagingInfo {
	s.WorkflowInstances = v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfo) Validate() error {
	if s.WorkflowInstances != nil {
		for _, item := range s.WorkflowInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances struct {
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
	// 100
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
	// The workflow instance ID.
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
	// 100
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name of the workflow instance.
	//
	// example:
	//
	// WorkflowInstance1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1000
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
	Tags []*ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the workflow instance. Valid values:
	//
	// 	- Normal: Scheduled execution
	//
	// 	- Manual: Manually triggered node
	//
	// 	- SmokeTest: Smoke test
	//
	// 	- SupplementData: Data backfill
	//
	// 	- ManualWorkflow: Manually triggered workflow
	//
	// 	- TriggerWorkflow: Triggered Workflow
	//
	// example:
	//
	// Normal
	Type                      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UnifiedWorkflowInstanceId *int64  `json:"UnifiedWorkflowInstanceId,omitempty" xml:"UnifiedWorkflowInstanceId,omitempty"`
	// The ID of the workflow to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow parameters.
	WorkflowParameters *string `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty"`
	// example:
	//
	// 1234
	WorkflowTaskInstanceId *int64 `json:"WorkflowTaskInstanceId,omitempty" xml:"WorkflowTaskInstanceId,omitempty"`
}

func (s ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetBizDate() *int64 {
	return s.BizDate
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetEnvType() *string {
	return s.EnvType
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetName() *string {
	return s.Name
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetOwner() *string {
	return s.Owner
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetTags() []*ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags {
	return s.Tags
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetType() *string {
	return s.Type
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetUnifiedWorkflowInstanceId() *int64 {
	return s.UnifiedWorkflowInstanceId
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetWorkflowParameters() *string {
	return s.WorkflowParameters
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) GetWorkflowTaskInstanceId() *int64 {
	return s.WorkflowTaskInstanceId
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetBizDate(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.BizDate = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetCreateTime(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetCreateUser(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.CreateUser = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetEnvType(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.EnvType = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetFinishedTime(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.FinishedTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetId(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.Id = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetModifyTime(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetModifyUser(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.ModifyUser = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetName(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.Name = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetOwner(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.Owner = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetProjectId(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetStartedTime(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.StartedTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetStatus(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetTags(v []*ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.Tags = v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetType(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.Type = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetUnifiedWorkflowInstanceId(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.UnifiedWorkflowInstanceId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetWorkflowId(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetWorkflowParameters(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.WorkflowParameters = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) SetWorkflowTaskInstanceId(v int64) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances {
	s.WorkflowTaskInstanceId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstances) Validate() error {
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

type ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags struct {
	// The key of a tag.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags) GetKey() *string {
	return s.Key
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags) GetValue() *string {
	return s.Value
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags) SetKey(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags {
	s.Key = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags) SetValue(v string) *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags {
	s.Value = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags) Validate() error {
	return dara.Validate(s)
}
