// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamTaskManagerLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentInstanceId(v string) *ListStreamTaskManagerLogsRequest
	GetDeploymentInstanceId() *string
	SetOpTenantId(v int64) *ListStreamTaskManagerLogsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListStreamTaskManagerLogsRequest
	GetOpUserId() *string
	SetProjectId(v int64) *ListStreamTaskManagerLogsRequest
	GetProjectId() *int64
	SetTaskmanagerId(v string) *ListStreamTaskManagerLogsRequest
	GetTaskmanagerId() *string
}

type ListStreamTaskManagerLogsRequest struct {
	// The deployment instance ID. Set this parameter to the deploymentInstanceId value returned in the operation record or historical running instance.
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
	// The user ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID of the Dataphin project to which the node belongs. This parameter is used for permission verification in the production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7299592831043008
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The identifier of the TaskManager. Set this parameter to the taskmanagerId value returned by the ListStreamTaskManagers operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataphin-session-7db678957dc449aab0086b931c9f8218-taskmanazjkxq
	TaskmanagerId *string `json:"TaskmanagerId,omitempty" xml:"TaskmanagerId,omitempty"`
}

func (s ListStreamTaskManagerLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStreamTaskManagerLogsRequest) GoString() string {
	return s.String()
}

func (s *ListStreamTaskManagerLogsRequest) GetDeploymentInstanceId() *string {
	return s.DeploymentInstanceId
}

func (s *ListStreamTaskManagerLogsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListStreamTaskManagerLogsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListStreamTaskManagerLogsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListStreamTaskManagerLogsRequest) GetTaskmanagerId() *string {
	return s.TaskmanagerId
}

func (s *ListStreamTaskManagerLogsRequest) SetDeploymentInstanceId(v string) *ListStreamTaskManagerLogsRequest {
	s.DeploymentInstanceId = &v
	return s
}

func (s *ListStreamTaskManagerLogsRequest) SetOpTenantId(v int64) *ListStreamTaskManagerLogsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListStreamTaskManagerLogsRequest) SetOpUserId(v string) *ListStreamTaskManagerLogsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListStreamTaskManagerLogsRequest) SetProjectId(v int64) *ListStreamTaskManagerLogsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListStreamTaskManagerLogsRequest) SetTaskmanagerId(v string) *ListStreamTaskManagerLogsRequest {
	s.TaskmanagerId = &v
	return s
}

func (s *ListStreamTaskManagerLogsRequest) Validate() error {
	return dara.Validate(s)
}
