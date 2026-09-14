// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrossProjectPipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentEnvironmentId(v int64) *CreateCrossProjectPipelineRunRequest
	GetDeploymentEnvironmentId() *int64
	SetDescription(v string) *CreateCrossProjectPipelineRunRequest
	GetDescription() *string
	SetObjectIds(v []*string) *CreateCrossProjectPipelineRunRequest
	GetObjectIds() []*string
	SetProjectId(v int64) *CreateCrossProjectPipelineRunRequest
	GetProjectId() *int64
	SetType(v string) *CreateCrossProjectPipelineRunRequest
	GetType() *string
}

type CreateCrossProjectPipelineRunRequest struct {
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
	ObjectIds []*string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty" type:"Repeated"`
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

func (s CreateCrossProjectPipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossProjectPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *CreateCrossProjectPipelineRunRequest) GetDeploymentEnvironmentId() *int64 {
	return s.DeploymentEnvironmentId
}

func (s *CreateCrossProjectPipelineRunRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCrossProjectPipelineRunRequest) GetObjectIds() []*string {
	return s.ObjectIds
}

func (s *CreateCrossProjectPipelineRunRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateCrossProjectPipelineRunRequest) GetType() *string {
	return s.Type
}

func (s *CreateCrossProjectPipelineRunRequest) SetDeploymentEnvironmentId(v int64) *CreateCrossProjectPipelineRunRequest {
	s.DeploymentEnvironmentId = &v
	return s
}

func (s *CreateCrossProjectPipelineRunRequest) SetDescription(v string) *CreateCrossProjectPipelineRunRequest {
	s.Description = &v
	return s
}

func (s *CreateCrossProjectPipelineRunRequest) SetObjectIds(v []*string) *CreateCrossProjectPipelineRunRequest {
	s.ObjectIds = v
	return s
}

func (s *CreateCrossProjectPipelineRunRequest) SetProjectId(v int64) *CreateCrossProjectPipelineRunRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCrossProjectPipelineRunRequest) SetType(v string) *CreateCrossProjectPipelineRunRequest {
	s.Type = &v
	return s
}

func (s *CreateCrossProjectPipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
