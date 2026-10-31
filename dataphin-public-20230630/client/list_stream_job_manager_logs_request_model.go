// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamJobManagerLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentInstanceId(v string) *ListStreamJobManagerLogsRequest
	GetDeploymentInstanceId() *string
	SetOpTenantId(v int64) *ListStreamJobManagerLogsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListStreamJobManagerLogsRequest
	GetOpUserId() *string
	SetProjectId(v int64) *ListStreamJobManagerLogsRequest
	GetProjectId() *int64
}

type ListStreamJobManagerLogsRequest struct {
	// The deployment instance ID. Use the deploymentInstanceId returned in operation records or historical run instances, rather than the taskId or deploymentId.
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
	// The operator ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID of the Dataphin project to which the task belongs, used for permission verification in the production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7299592831043008
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListStreamJobManagerLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStreamJobManagerLogsRequest) GoString() string {
	return s.String()
}

func (s *ListStreamJobManagerLogsRequest) GetDeploymentInstanceId() *string {
	return s.DeploymentInstanceId
}

func (s *ListStreamJobManagerLogsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListStreamJobManagerLogsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListStreamJobManagerLogsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListStreamJobManagerLogsRequest) SetDeploymentInstanceId(v string) *ListStreamJobManagerLogsRequest {
	s.DeploymentInstanceId = &v
	return s
}

func (s *ListStreamJobManagerLogsRequest) SetOpTenantId(v int64) *ListStreamJobManagerLogsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListStreamJobManagerLogsRequest) SetOpUserId(v string) *ListStreamJobManagerLogsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListStreamJobManagerLogsRequest) SetProjectId(v int64) *ListStreamJobManagerLogsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListStreamJobManagerLogsRequest) Validate() error {
	return dara.Validate(s)
}
