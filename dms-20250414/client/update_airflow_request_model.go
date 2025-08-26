// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAirflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowId(v string) *UpdateAirflowRequest
	GetAirflowId() *string
	SetAirflowName(v string) *UpdateAirflowRequest
	GetAirflowName() *string
	SetAppSpec(v string) *UpdateAirflowRequest
	GetAppSpec() *string
	SetClientToken(v string) *UpdateAirflowRequest
	GetClientToken() *string
	SetDagsDir(v string) *UpdateAirflowRequest
	GetDagsDir() *string
	SetDescription(v string) *UpdateAirflowRequest
	GetDescription() *string
	SetPluginsDir(v string) *UpdateAirflowRequest
	GetPluginsDir() *string
	SetRequirementFile(v string) *UpdateAirflowRequest
	GetRequirementFile() *string
	SetStartupFile(v string) *UpdateAirflowRequest
	GetStartupFile() *string
	SetWorkerServerlessReplicas(v int32) *UpdateAirflowRequest
	GetWorkerServerlessReplicas() *int32
	SetWorkspaceId(v string) *UpdateAirflowRequest
	GetWorkspaceId() *string
}

type UpdateAirflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af-test****
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// example:
	//
	// testairflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
	// example:
	//
	// SMALL
	AppSpec *string `json:"AppSpec,omitempty" xml:"AppSpec,omitempty"`
	// example:
	//
	// token-****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// example:
	//
	// test airflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// default/plugins
	PluginsDir *string `json:"PluginsDir,omitempty" xml:"PluginsDir,omitempty"`
	// example:
	//
	// default/requirements.txt
	RequirementFile *string `json:"RequirementFile,omitempty" xml:"RequirementFile,omitempty"`
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// example:
	//
	// 0
	WorkerServerlessReplicas *int32 `json:"WorkerServerlessReplicas,omitempty" xml:"WorkerServerlessReplicas,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 863024238280****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateAirflowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAirflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateAirflowRequest) GetAirflowId() *string {
	return s.AirflowId
}

func (s *UpdateAirflowRequest) GetAirflowName() *string {
	return s.AirflowName
}

func (s *UpdateAirflowRequest) GetAppSpec() *string {
	return s.AppSpec
}

func (s *UpdateAirflowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAirflowRequest) GetDagsDir() *string {
	return s.DagsDir
}

func (s *UpdateAirflowRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAirflowRequest) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *UpdateAirflowRequest) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *UpdateAirflowRequest) GetStartupFile() *string {
	return s.StartupFile
}

func (s *UpdateAirflowRequest) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *UpdateAirflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateAirflowRequest) SetAirflowId(v string) *UpdateAirflowRequest {
	s.AirflowId = &v
	return s
}

func (s *UpdateAirflowRequest) SetAirflowName(v string) *UpdateAirflowRequest {
	s.AirflowName = &v
	return s
}

func (s *UpdateAirflowRequest) SetAppSpec(v string) *UpdateAirflowRequest {
	s.AppSpec = &v
	return s
}

func (s *UpdateAirflowRequest) SetClientToken(v string) *UpdateAirflowRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAirflowRequest) SetDagsDir(v string) *UpdateAirflowRequest {
	s.DagsDir = &v
	return s
}

func (s *UpdateAirflowRequest) SetDescription(v string) *UpdateAirflowRequest {
	s.Description = &v
	return s
}

func (s *UpdateAirflowRequest) SetPluginsDir(v string) *UpdateAirflowRequest {
	s.PluginsDir = &v
	return s
}

func (s *UpdateAirflowRequest) SetRequirementFile(v string) *UpdateAirflowRequest {
	s.RequirementFile = &v
	return s
}

func (s *UpdateAirflowRequest) SetStartupFile(v string) *UpdateAirflowRequest {
	s.StartupFile = &v
	return s
}

func (s *UpdateAirflowRequest) SetWorkerServerlessReplicas(v int32) *UpdateAirflowRequest {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *UpdateAirflowRequest) SetWorkspaceId(v string) *UpdateAirflowRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateAirflowRequest) Validate() error {
	return dara.Validate(s)
}
