// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v string) *GetTaskRequest
	GetEntityId() *string
	SetEntityType(v string) *GetTaskRequest
	GetEntityType() *string
	SetEntityVersionId(v string) *GetTaskRequest
	GetEntityVersionId() *string
	SetTaskCategory(v string) *GetTaskRequest
	GetTaskCategory() *string
	SetTaskId(v string) *GetTaskRequest
	GetTaskId() *string
	SetTaskType(v string) *GetTaskRequest
	GetTaskType() *string
}

type GetTaskRequest struct {
	// The associated resource ID. **If TaskId is not specified, you must provide EntityId, EntityType, and TaskCategory at the same time.*	- The resource ID must match the selected resource type.
	//
	// example:
	//
	// imgc-bp1example****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The associated resource type. Required if TaskId is not specified.
	//
	// Valid values:
	//
	// - **IMAGE**: Image resource.
	//
	// - **APP**: Application resource.
	//
	// - **SNAPSHOT**: Snapshot resource.
	//
	// You must also select the applicable TaskCategory. The resource type itself does not mean that all task categories can be queried by resource.
	//
	// example:
	//
	// IMAGE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The associated resource version ID. Used to specify the version when querying by resource. For image resources, enter the image version ID.
	//
	// example:
	//
	// iv-bp1example****
	EntityVersionId *string `json:"EntityVersionId,omitempty" xml:"EntityVersionId,omitempty"`
	// The task category when querying by resource. Required if TaskId is not specified.
	//
	// Valid values:
	//
	// - **IMAGE_BUILDER**: Image creation. Queries tasks for manual creation, automatic creation, or creation based on an instance.
	//
	// - **IMAGE_DISTRIBUTE**: Image region distribution task.
	//
	// example:
	//
	// IMAGE_BUILDER
	TaskCategory *string `json:"TaskCategory,omitempty" xml:"TaskCategory,omitempty"`
	// The task ID. If specified, the query prioritizes this ID, and resource information query conditions are no longer used for locating the task.
	//
	// example:
	//
	// tid-bp1example****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// **Not enabled. Do not pass this parameter.*	- When querying a task, locate the task by using TaskId or resource information and TaskCategory. This parameter is not used for task filtering.
	//
	// The example "Not applicable (not enabled)" is only a description of the field status, not a valid parameter value.
	//
	// example:
	//
	// Not applicable (not enabled)
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *GetTaskRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *GetTaskRequest) GetEntityVersionId() *string {
	return s.EntityVersionId
}

func (s *GetTaskRequest) GetTaskCategory() *string {
	return s.TaskCategory
}

func (s *GetTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskRequest) SetEntityId(v string) *GetTaskRequest {
	s.EntityId = &v
	return s
}

func (s *GetTaskRequest) SetEntityType(v string) *GetTaskRequest {
	s.EntityType = &v
	return s
}

func (s *GetTaskRequest) SetEntityVersionId(v string) *GetTaskRequest {
	s.EntityVersionId = &v
	return s
}

func (s *GetTaskRequest) SetTaskCategory(v string) *GetTaskRequest {
	s.TaskCategory = &v
	return s
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskRequest) SetTaskType(v string) *GetTaskRequest {
	s.TaskType = &v
	return s
}

func (s *GetTaskRequest) Validate() error {
	return dara.Validate(s)
}
