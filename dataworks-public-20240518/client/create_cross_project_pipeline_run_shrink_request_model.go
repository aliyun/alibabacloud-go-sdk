// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrossProjectPipelineRunShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentEnvironmentId(v int64) *CreateCrossProjectPipelineRunShrinkRequest
	GetDeploymentEnvironmentId() *int64
	SetDescription(v string) *CreateCrossProjectPipelineRunShrinkRequest
	GetDescription() *string
	SetObjectIdsShrink(v string) *CreateCrossProjectPipelineRunShrinkRequest
	GetObjectIdsShrink() *string
	SetProjectId(v int64) *CreateCrossProjectPipelineRunShrinkRequest
	GetProjectId() *int64
	SetType(v string) *CreateCrossProjectPipelineRunShrinkRequest
	GetType() *string
}

type CreateCrossProjectPipelineRunShrinkRequest struct {
	// The cross-workspace deployment environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 101
	DeploymentEnvironmentId *int64 `json:"DeploymentEnvironmentId,omitempty" xml:"DeploymentEnvironmentId,omitempty"`
	// The deployment description.
	//
	// example:
	//
	// This is a business process created through the API
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of top-level object IDs from the source project to deploy. The list must contain exactly one object. Child objects of composite objects such as workflows are automatically included by the system.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1"]
	ObjectIdsShrink *string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The deployment type. Valid values:
	//
	// - Offline: Offline deployment.
	//
	// - Online: Online deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCrossProjectPipelineRunShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossProjectPipelineRunShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) GetDeploymentEnvironmentId() *int64 {
	return s.DeploymentEnvironmentId
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) GetObjectIdsShrink() *string {
	return s.ObjectIdsShrink
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) SetDeploymentEnvironmentId(v int64) *CreateCrossProjectPipelineRunShrinkRequest {
	s.DeploymentEnvironmentId = &v
	return s
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) SetDescription(v string) *CreateCrossProjectPipelineRunShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) SetObjectIdsShrink(v string) *CreateCrossProjectPipelineRunShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) SetProjectId(v int64) *CreateCrossProjectPipelineRunShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) SetType(v string) *CreateCrossProjectPipelineRunShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateCrossProjectPipelineRunShrinkRequest) Validate() error {
	return dara.Validate(s)
}
