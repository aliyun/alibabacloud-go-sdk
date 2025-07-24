// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProcessDefinitionWithScheduleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetAlertEmailAddress() *string
	SetDescription(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetDescription() *string
	SetExecutionType(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetExecutionType() *string
	SetGlobalParamsShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetGlobalParamsShrink() *string
	SetName(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetName() *string
	SetProductNamespace(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetProductNamespace() *string
	SetPublish(v bool) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetPublish() *bool
	SetRegionId(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetRegionId() *string
	SetReleaseState(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetReleaseState() *string
	SetResourceQueue(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetResourceQueue() *string
	SetRetryTimes(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetRetryTimes() *int32
	SetRunAs(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetRunAs() *string
	SetScheduleShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetScheduleShrink() *string
	SetTagsShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetTagsShrink() *string
	SetTaskDefinitionJsonShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetTaskDefinitionJsonShrink() *string
	SetTaskParallelism(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetTaskParallelism() *int32
	SetTaskRelationJsonShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetTaskRelationJsonShrink() *string
	SetTimeout(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest
	GetTimeout() *int32
}

type UpdateProcessDefinitionWithScheduleShrinkRequest struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// ods batch workflow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution policy.
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
	// The status of the workflow.
	//
	// example:
	//
	// ONLINE
	ReleaseState *string `json:"releaseState,omitempty" xml:"releaseState,omitempty"`
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
	// The execution user.
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
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetAlertEmailAddress() *string {
	return s.AlertEmailAddress
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetGlobalParamsShrink() *string {
	return s.GlobalParamsShrink
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetProductNamespace() *string {
	return s.ProductNamespace
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetPublish() *bool {
	return s.Publish
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetReleaseState() *string {
	return s.ReleaseState
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetResourceQueue() *string {
	return s.ResourceQueue
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetRunAs() *string {
	return s.RunAs
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetScheduleShrink() *string {
	return s.ScheduleShrink
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetTaskDefinitionJsonShrink() *string {
	return s.TaskDefinitionJsonShrink
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetTaskParallelism() *int32 {
	return s.TaskParallelism
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetTaskRelationJsonShrink() *string {
	return s.TaskRelationJsonShrink
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.AlertEmailAddress = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetDescription(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetExecutionType(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ExecutionType = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetGlobalParamsShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.GlobalParamsShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetName(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetProductNamespace(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ProductNamespace = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetPublish(v bool) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.Publish = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetRegionId(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetReleaseState(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ReleaseState = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetResourceQueue(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ResourceQueue = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetRetryTimes(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.RetryTimes = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetRunAs(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.RunAs = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetScheduleShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTagsShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTaskDefinitionJsonShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskDefinitionJsonShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTaskParallelism(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskParallelism = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTaskRelationJsonShrink(v string) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.TaskRelationJsonShrink = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) SetTimeout(v int32) *UpdateProcessDefinitionWithScheduleShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
