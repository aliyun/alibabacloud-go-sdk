// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamTaskManagersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentInstanceId(v string) *ListStreamTaskManagersRequest
	GetDeploymentInstanceId() *string
	SetOpTenantId(v int64) *ListStreamTaskManagersRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListStreamTaskManagersRequest
	GetOpUserId() *string
	SetProjectId(v int64) *ListStreamTaskManagersRequest
	GetProjectId() *int64
}

type ListStreamTaskManagersRequest struct {
	// The deployment instance ID. You can obtain the value from the deploymentInstanceId returned in the operation record or historical run instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7db678957dc449aab0086b931c9f8218
	DeploymentInstanceId *string `json:"DeploymentInstanceId,omitempty" xml:"DeploymentInstanceId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID of the Dataphin project to which the task belongs. This parameter is used for permission verification in the production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7299592831043008
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListStreamTaskManagersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStreamTaskManagersRequest) GoString() string {
	return s.String()
}

func (s *ListStreamTaskManagersRequest) GetDeploymentInstanceId() *string {
	return s.DeploymentInstanceId
}

func (s *ListStreamTaskManagersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListStreamTaskManagersRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListStreamTaskManagersRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListStreamTaskManagersRequest) SetDeploymentInstanceId(v string) *ListStreamTaskManagersRequest {
	s.DeploymentInstanceId = &v
	return s
}

func (s *ListStreamTaskManagersRequest) SetOpTenantId(v int64) *ListStreamTaskManagersRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListStreamTaskManagersRequest) SetOpUserId(v string) *ListStreamTaskManagersRequest {
	s.OpUserId = &v
	return s
}

func (s *ListStreamTaskManagersRequest) SetProjectId(v int64) *ListStreamTaskManagersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListStreamTaskManagersRequest) Validate() error {
	return dara.Validate(s)
}
