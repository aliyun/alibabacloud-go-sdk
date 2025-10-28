// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceQueuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspaceQueuesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspaceQueuesResponseBody
	GetNextToken() *string
	SetQueues(v []*ListWorkspaceQueuesResponseBodyQueues) *ListWorkspaceQueuesResponseBody
	GetQueues() []*ListWorkspaceQueuesResponseBodyQueues
	SetRequestId(v string) *ListWorkspaceQueuesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListWorkspaceQueuesResponseBody
	GetTotalCount() *int32
}

type ListWorkspaceQueuesResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of queues.
	Queues []*ListWorkspaceQueuesResponseBodyQueues `json:"queues,omitempty" xml:"queues,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListWorkspaceQueuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspaceQueuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspaceQueuesResponseBody) GetQueues() []*ListWorkspaceQueuesResponseBodyQueues {
	return s.Queues
}

func (s *ListWorkspaceQueuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspaceQueuesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkspaceQueuesResponseBody) SetMaxResults(v int32) *ListWorkspaceQueuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) SetNextToken(v string) *ListWorkspaceQueuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) SetQueues(v []*ListWorkspaceQueuesResponseBodyQueues) *ListWorkspaceQueuesResponseBody {
	s.Queues = v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) SetRequestId(v string) *ListWorkspaceQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) SetTotalCount(v int32) *ListWorkspaceQueuesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBody) Validate() error {
	if s.Queues != nil {
		for _, item := range s.Queues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceQueuesResponseBodyQueues struct {
	// The operations allowed for the queue.
	AllowActions []*ListWorkspaceQueuesResponseBodyQueuesAllowActions `json:"allowActions,omitempty" xml:"allowActions,omitempty" type:"Repeated"`
	// The time when the workspace was created.
	//
	// example:
	//
	// 1684115879955
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the user who created the queue.
	//
	// example:
	//
	// 237109
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The environment types of the queue.
	Environments []*string `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// The maximum capacity of resources that can be used in the queue.
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	MaxResource *string `json:"maxResource,omitempty" xml:"maxResource,omitempty"`
	// The minimum capacity of resources that can be used in the queue.
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	MinResource *string `json:"minResource,omitempty" xml:"minResource,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Pre
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	Preheat     *bool   `json:"preheat,omitempty" xml:"preheat,omitempty"`
	// The queue label.
	//
	// example:
	//
	// dev_queue
	Properties *string `json:"properties,omitempty" xml:"properties,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// dev_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The queue architecture.
	//
	// example:
	//
	// {"arch": "x86"}
	QueueScope *string `json:"queueScope,omitempty" xml:"queueScope,omitempty"`
	// The status of the queue.
	//
	// example:
	//
	// RUNNING
	QueueStatus *string `json:"queueStatus,omitempty" xml:"queueStatus,omitempty"`
	// The type of the queue. Valid values:
	//
	// 	- instance
	//
	// 	- instanceChildren
	//
	// example:
	//
	// instance, instanceChildren
	QueueType *string `json:"queueType,omitempty" xml:"queueType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The capacity of resources that are used in the queue.
	//
	// example:
	//
	// {"cpu": "2","memory": "2Gi"}
	UsedResource *string `json:"usedResource,omitempty" xml:"usedResource,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkspaceQueuesResponseBodyQueues) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceQueuesResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetAllowActions() []*ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	return s.AllowActions
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetEnvironments() []*string {
	return s.Environments
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetMaxResource() *string {
	return s.MaxResource
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetMinResource() *string {
	return s.MinResource
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetPreheat() *bool {
	return s.Preheat
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetProperties() *string {
	return s.Properties
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetQueueName() *string {
	return s.QueueName
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetQueueScope() *string {
	return s.QueueScope
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetQueueStatus() *string {
	return s.QueueStatus
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetQueueType() *string {
	return s.QueueType
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetUsedResource() *string {
	return s.UsedResource
}

func (s *ListWorkspaceQueuesResponseBodyQueues) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetAllowActions(v []*ListWorkspaceQueuesResponseBodyQueuesAllowActions) *ListWorkspaceQueuesResponseBodyQueues {
	s.AllowActions = v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetCreateTime(v int64) *ListWorkspaceQueuesResponseBodyQueues {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetCreator(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.Creator = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetEnvironments(v []*string) *ListWorkspaceQueuesResponseBodyQueues {
	s.Environments = v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetMaxResource(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.MaxResource = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetMinResource(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.MinResource = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetPaymentType(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.PaymentType = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetPreheat(v bool) *ListWorkspaceQueuesResponseBodyQueues {
	s.Preheat = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetProperties(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.Properties = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetQueueName(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.QueueName = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetQueueScope(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.QueueScope = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetQueueStatus(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.QueueStatus = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetQueueType(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.QueueType = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetRegionId(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.RegionId = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetUsedResource(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.UsedResource = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) SetWorkspaceId(v string) *ListWorkspaceQueuesResponseBodyQueues {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueues) Validate() error {
	if s.AllowActions != nil {
		for _, item := range s.AllowActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceQueuesResponseBodyQueuesAllowActions struct {
	// The Alibaba Cloud Resource Name (ARN) of a behavior.
	//
	// example:
	//
	// acs:emr::workspaceId:action/create_queue
	ActionArn *string `json:"actionArn,omitempty" xml:"actionArn,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// view
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// The dependencies of the operation.
	//
	// example:
	//
	// ["view"]
	Dependencies []*string `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	// The description of the operation.
	//
	// example:
	//
	// 文件目录遍历、文件浏览
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the permission.
	//
	// example:
	//
	// 文件目录遍历、文件浏览
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s ListWorkspaceQueuesResponseBodyQueuesAllowActions) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceQueuesResponseBodyQueuesAllowActions) GoString() string {
	return s.String()
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) GetActionArn() *string {
	return s.ActionArn
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) GetActionName() *string {
	return s.ActionName
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) GetDependencies() []*string {
	return s.Dependencies
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) GetDescription() *string {
	return s.Description
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetActionArn(v string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.ActionArn = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetActionName(v string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.ActionName = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetDependencies(v []*string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.Dependencies = v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetDescription(v string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.Description = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) SetDisplayName(v string) *ListWorkspaceQueuesResponseBodyQueuesAllowActions {
	s.DisplayName = &v
	return s
}

func (s *ListWorkspaceQueuesResponseBodyQueuesAllowActions) Validate() error {
	return dara.Validate(s)
}
