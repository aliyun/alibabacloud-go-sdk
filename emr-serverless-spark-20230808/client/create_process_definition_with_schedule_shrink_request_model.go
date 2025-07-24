// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcessDefinitionWithScheduleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEmailAddress(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetAlertEmailAddress() *string
	SetDescription(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetDescription() *string
	SetExecutionType(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetExecutionType() *string
	SetGlobalParamsShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetGlobalParamsShrink() *string
	SetName(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetName() *string
	SetProductNamespace(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetProductNamespace() *string
	SetPublish(v bool) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetPublish() *bool
	SetRegionId(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetRegionId() *string
	SetResourceQueue(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetResourceQueue() *string
	SetRetryTimes(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetRetryTimes() *int32
	SetRunAs(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetRunAs() *string
	SetScheduleShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetScheduleShrink() *string
	SetTagsShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetTagsShrink() *string
	SetTaskDefinitionJsonShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetTaskDefinitionJsonShrink() *string
	SetTaskParallelism(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetTaskParallelism() *int32
	SetTaskRelationJsonShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetTaskRelationJsonShrink() *string
	SetTimeout(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest
	GetTimeout() *int32
}

type CreateProcessDefinitionWithScheduleShrinkRequest struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The description of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods batch workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy
	//
	// This parameter is required.
	//
	// example:
	//
	// PARALLEL
	ExecutionType      *string `json:"executionType,omitempty" xml:"executionType,omitempty"`
	GlobalParamsShrink *string `json:"globalParams,omitempty" xml:"globalParams,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// Specifies whether to publish the workflow.
	//
	// example:
	//
	// true
	Publish *bool `json:"publish,omitempty" xml:"publish,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource queue.
	//
	// example:
	//
	// root_queue
	ResourceQueue *string `json:"resourceQueue,omitempty" xml:"resourceQueue,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 1
	RetryTimes *int32 `json:"retryTimes,omitempty" xml:"retryTimes,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who creates the workflow.
	//
	// example:
	//
	// 113***************
	RunAs *string `json:"runAs,omitempty" xml:"runAs,omitempty"`
	// The scheduling settings.
	ScheduleShrink *string `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// The tags.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The descriptions of all nodes in the workflow.
	//
	// This parameter is required.
	TaskDefinitionJsonShrink *string `json:"taskDefinitionJson,omitempty" xml:"taskDefinitionJson,omitempty"`
	// The node parallelism.
	//
	// example:
	//
	// 1
	TaskParallelism *int32 `json:"taskParallelism,omitempty" xml:"taskParallelism,omitempty"`
	// The dependencies of all nodes in the workflow. preTaskCode specifies the ID of an upstream node, and postTaskCode specifies the ID of a downstream node. The ID of each node is unique. If a node does not have an upstream node, set preTaskCode to 0.
	//
	// This parameter is required.
	TaskRelationJsonShrink *string `json:"taskRelationJson,omitempty" xml:"taskRelationJson,omitempty"`
	// The default timeout period of the workflow.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateProcessDefinitionWithScheduleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionWithScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetAlertEmailAddress() *string {
	return s.AlertEmailAddress
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetGlobalParamsShrink() *string {
	return s.GlobalParamsShrink
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetProductNamespace() *string {
	return s.ProductNamespace
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetPublish() *bool {
	return s.Publish
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetResourceQueue() *string {
	return s.ResourceQueue
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetRunAs() *string {
	return s.RunAs
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetScheduleShrink() *string {
	return s.ScheduleShrink
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetTaskDefinitionJsonShrink() *string {
	return s.TaskDefinitionJsonShrink
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetTaskParallelism() *int32 {
	return s.TaskParallelism
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetTaskRelationJsonShrink() *string {
	return s.TaskRelationJsonShrink
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetAlertEmailAddress(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.AlertEmailAddress = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetDescription(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetExecutionType(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.ExecutionType = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetGlobalParamsShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.GlobalParamsShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetName(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetProductNamespace(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.ProductNamespace = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetPublish(v bool) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.Publish = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetRegionId(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetResourceQueue(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.ResourceQueue = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetRetryTimes(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.RetryTimes = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetRunAs(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.RunAs = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetScheduleShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTagsShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTaskDefinitionJsonShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskDefinitionJsonShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTaskParallelism(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskParallelism = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTaskRelationJsonShrink(v string) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskRelationJsonShrink = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) SetTimeout(v int32) *CreateProcessDefinitionWithScheduleShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *CreateProcessDefinitionWithScheduleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
