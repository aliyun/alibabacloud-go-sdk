// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAirflowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowId(v string) *UpdateAirflowShrinkRequest
	GetAirflowId() *string
	SetAirflowName(v string) *UpdateAirflowShrinkRequest
	GetAirflowName() *string
	SetAppSpec(v string) *UpdateAirflowShrinkRequest
	GetAppSpec() *string
	SetClientToken(v string) *UpdateAirflowShrinkRequest
	GetClientToken() *string
	SetDagsDir(v string) *UpdateAirflowShrinkRequest
	GetDagsDir() *string
	SetDataMountInfoListShrink(v string) *UpdateAirflowShrinkRequest
	GetDataMountInfoListShrink() *string
	SetDescription(v string) *UpdateAirflowShrinkRequest
	GetDescription() *string
	SetEnableServerless(v bool) *UpdateAirflowShrinkRequest
	GetEnableServerless() *bool
	SetGracefulShutdownTimeout(v int32) *UpdateAirflowShrinkRequest
	GetGracefulShutdownTimeout() *int32
	SetPluginsDir(v string) *UpdateAirflowShrinkRequest
	GetPluginsDir() *string
	SetRequirementFile(v string) *UpdateAirflowShrinkRequest
	GetRequirementFile() *string
	SetStartupFile(v string) *UpdateAirflowShrinkRequest
	GetStartupFile() *string
	SetWorkerServerlessReplicas(v int32) *UpdateAirflowShrinkRequest
	GetWorkerServerlessReplicas() *int32
	SetWorkspaceId(v string) *UpdateAirflowShrinkRequest
	GetWorkspaceId() *string
}

type UpdateAirflowShrinkRequest struct {
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
	DagsDir                 *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	DataMountInfoListShrink *string `json:"DataMountInfoList,omitempty" xml:"DataMountInfoList,omitempty"`
	// example:
	//
	// test airflow
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableServerless *bool   `json:"EnableServerless,omitempty" xml:"EnableServerless,omitempty"`
	// example:
	//
	// 60
	GracefulShutdownTimeout *int32 `json:"GracefulShutdownTimeout,omitempty" xml:"GracefulShutdownTimeout,omitempty"`
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

func (s UpdateAirflowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAirflowShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAirflowShrinkRequest) GetAirflowId() *string {
	return s.AirflowId
}

func (s *UpdateAirflowShrinkRequest) GetAirflowName() *string {
	return s.AirflowName
}

func (s *UpdateAirflowShrinkRequest) GetAppSpec() *string {
	return s.AppSpec
}

func (s *UpdateAirflowShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAirflowShrinkRequest) GetDagsDir() *string {
	return s.DagsDir
}

func (s *UpdateAirflowShrinkRequest) GetDataMountInfoListShrink() *string {
	return s.DataMountInfoListShrink
}

func (s *UpdateAirflowShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAirflowShrinkRequest) GetEnableServerless() *bool {
	return s.EnableServerless
}

func (s *UpdateAirflowShrinkRequest) GetGracefulShutdownTimeout() *int32 {
	return s.GracefulShutdownTimeout
}

func (s *UpdateAirflowShrinkRequest) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *UpdateAirflowShrinkRequest) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *UpdateAirflowShrinkRequest) GetStartupFile() *string {
	return s.StartupFile
}

func (s *UpdateAirflowShrinkRequest) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *UpdateAirflowShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateAirflowShrinkRequest) SetAirflowId(v string) *UpdateAirflowShrinkRequest {
	s.AirflowId = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetAirflowName(v string) *UpdateAirflowShrinkRequest {
	s.AirflowName = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetAppSpec(v string) *UpdateAirflowShrinkRequest {
	s.AppSpec = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetClientToken(v string) *UpdateAirflowShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetDagsDir(v string) *UpdateAirflowShrinkRequest {
	s.DagsDir = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetDataMountInfoListShrink(v string) *UpdateAirflowShrinkRequest {
	s.DataMountInfoListShrink = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetDescription(v string) *UpdateAirflowShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetEnableServerless(v bool) *UpdateAirflowShrinkRequest {
	s.EnableServerless = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetGracefulShutdownTimeout(v int32) *UpdateAirflowShrinkRequest {
	s.GracefulShutdownTimeout = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetPluginsDir(v string) *UpdateAirflowShrinkRequest {
	s.PluginsDir = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetRequirementFile(v string) *UpdateAirflowShrinkRequest {
	s.RequirementFile = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetStartupFile(v string) *UpdateAirflowShrinkRequest {
	s.StartupFile = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetWorkerServerlessReplicas(v int32) *UpdateAirflowShrinkRequest {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetWorkspaceId(v string) *UpdateAirflowShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
