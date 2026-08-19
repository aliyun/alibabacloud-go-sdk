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
	// The request ID. Used for locating logs and troubleshooting issues.
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
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of workflow instances.
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
	// The business date.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The creation time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the user who created the instance.
	//
	// example:
	//
	// 100
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod (production)
	//
	// - Dev (development)
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The finish time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The unique identifier of the workflow instance.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account ID of the user who last modified the instance.
	//
	// example:
	//
	// 100
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name.
	//
	// example:
	//
	// WorkflowInstance1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the workflow owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The start time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The running status of the workflow instance. Valid values:
	//
	// - NotRun: not run
	//
	// - Running: running
	//
	// - WaitTime: waiting for TriggerTime
	//
	// - CheckingCondition: checking branch conditions
	//
	// - WaitResource: waiting for resources
	//
	// - Failure: failed
	//
	// - Success: succeeded
	//
	// - Checking: submitted for data quality check
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The node tags.
	Tags []*ListWorkflowInstancesResponseBodyPagingInfoWorkflowInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the workflow instance. Valid values:
	//
	// - Normal: periodic scheduling
	//
	// - Manual: manual task
	//
	// - SmokeTest: test
	//
	// - SupplementData: data backfill
	//
	// - ManualWorkflow: manual workflow
	//
	// - TriggerWorkflow: trigger-based workflow
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unified workflow instance ID. All workflow instances within the same business date under a specific trigger share the same value for this field.
	//
	// example:
	//
	// 1234
	UnifiedWorkflowInstanceId *int64 `json:"UnifiedWorkflowInstanceId,omitempty" xml:"UnifiedWorkflowInstanceId,omitempty"`
	// The ID of the workflow to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow parameters.
	//
	// example:
	//
	// Periodic workflow:
	//
	// key1=value1 key2=value2
	//
	// Manual workflow:
	//
	// {"key1":"value1", "key2": "value2"}
	WorkflowParameters *string `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty"`
	// The task instance ID corresponding to the workflow instance.
	//
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
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
