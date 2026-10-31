// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamLogContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentInstanceId(v string) *GetStreamLogContentRequest
	GetDeploymentInstanceId() *string
	SetLogName(v string) *GetStreamLogContentRequest
	GetLogName() *string
	SetOpTenantId(v int64) *GetStreamLogContentRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetStreamLogContentRequest
	GetOpUserId() *string
	SetProjectId(v int64) *GetStreamLogContentRequest
	GetProjectId() *int64
	SetTaskmanagerId(v string) *GetStreamLogContentRequest
	GetTaskmanagerId() *string
}

type GetStreamLogContentRequest struct {
	// The deployment instance ID. Obtain the value from the deploymentInstanceId parameter returned in the operation record or historical running instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7db678957dc449aab0086b931c9f8218
	DeploymentInstanceId *string `json:"DeploymentInstanceId,omitempty" xml:"DeploymentInstanceId,omitempty"`
	// The original name of the log file. Obtain the value from the logName parameter in the corresponding JobManager or TaskManager log file list. Pass the value as is.
	//
	// This parameter is required.
	//
	// example:
	//
	// flink--standalonesession-0-dataphin-session-7db678957dc449aab0086b931c9f8218-jobmanager.log
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
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
	// The optional ID of the TaskManager. Obtain the value from the taskmanagerId parameter returned by the ListStreamTaskManagers operation. If this parameter is not specified or is empty, the JobManager logs are retrieved.
	//
	// example:
	//
	// dataphin-session-7db678957dc449aab0086b931c9f8218-taskmanazjkxq
	TaskmanagerId *string `json:"TaskmanagerId,omitempty" xml:"TaskmanagerId,omitempty"`
}

func (s GetStreamLogContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStreamLogContentRequest) GoString() string {
	return s.String()
}

func (s *GetStreamLogContentRequest) GetDeploymentInstanceId() *string {
	return s.DeploymentInstanceId
}

func (s *GetStreamLogContentRequest) GetLogName() *string {
	return s.LogName
}

func (s *GetStreamLogContentRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStreamLogContentRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetStreamLogContentRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetStreamLogContentRequest) GetTaskmanagerId() *string {
	return s.TaskmanagerId
}

func (s *GetStreamLogContentRequest) SetDeploymentInstanceId(v string) *GetStreamLogContentRequest {
	s.DeploymentInstanceId = &v
	return s
}

func (s *GetStreamLogContentRequest) SetLogName(v string) *GetStreamLogContentRequest {
	s.LogName = &v
	return s
}

func (s *GetStreamLogContentRequest) SetOpTenantId(v int64) *GetStreamLogContentRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStreamLogContentRequest) SetOpUserId(v string) *GetStreamLogContentRequest {
	s.OpUserId = &v
	return s
}

func (s *GetStreamLogContentRequest) SetProjectId(v int64) *GetStreamLogContentRequest {
	s.ProjectId = &v
	return s
}

func (s *GetStreamLogContentRequest) SetTaskmanagerId(v string) *GetStreamLogContentRequest {
	s.TaskmanagerId = &v
	return s
}

func (s *GetStreamLogContentRequest) Validate() error {
	return dara.Validate(s)
}
