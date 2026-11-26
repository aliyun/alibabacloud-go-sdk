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
	SetEnableLineage(v bool) *UpdateAirflowShrinkRequest
	GetEnableLineage() *bool
	SetEnableServerless(v bool) *UpdateAirflowShrinkRequest
	GetEnableServerless() *bool
	SetGracefulShutdownTimeout(v int32) *UpdateAirflowShrinkRequest
	GetGracefulShutdownTimeout() *int32
	SetLineageReplicas(v int32) *UpdateAirflowShrinkRequest
	GetLineageReplicas() *int32
	SetLineageSpec(v string) *UpdateAirflowShrinkRequest
	GetLineageSpec() *string
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
	// The unique key ID of Airflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// af-test****
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// The name of the Airflow instance.
	//
	// example:
	//
	// testairflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
	// The specification description of Airflow.
	//
	// example:
	//
	// SMALL
	AppSpec *string `json:"AppSpec,omitempty" xml:"AppSpec,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// token-****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The DAG directory scanned by Airflow.
	//
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// The list of data mount information.
	DataMountInfoListShrink *string `json:"DataMountInfoList,omitempty" xml:"DataMountInfoList,omitempty"`
	// The description of Airflow.
	//
	// example:
	//
	// test airflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable data lineage. For Airflow 2.x, this feature is always disabled, and the lineage specifications and number of replicas are ignored.
	EnableLineage *bool `json:"EnableLineage,omitempty" xml:"EnableLineage,omitempty"`
	// Specifies whether to enable serverless workers.
	EnableServerless *bool `json:"EnableServerless,omitempty" xml:"EnableServerless,omitempty"`
	// The timeout period for worker exit.
	//
	// example:
	//
	// 60
	GracefulShutdownTimeout *int32 `json:"GracefulShutdownTimeout,omitempty" xml:"GracefulShutdownTimeout,omitempty"`
	// The number of replicas for the lineage service. The default value is 1 when enabled. If omitted during an update, the original value is retained.
	//
	// example:
	//
	// 1
	LineageReplicas *int32 `json:"LineageReplicas,omitempty" xml:"LineageReplicas,omitempty"`
	// The merged lineage service specification, which uses the same tier as AppSpec.
	//
	// example:
	//
	// SMALL
	LineageSpec *string `json:"LineageSpec,omitempty" xml:"LineageSpec,omitempty"`
	// The plugin directory scanned by Airflow.
	//
	// example:
	//
	// default/plugins
	PluginsDir *string `json:"PluginsDir,omitempty" xml:"PluginsDir,omitempty"`
	// The installation file for package dependencies.
	//
	// example:
	//
	// default/requirements.txt
	RequirementFile *string `json:"RequirementFile,omitempty" xml:"RequirementFile,omitempty"`
	// The script loaded when the Airflow instance starts.
	//
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// The number of worker nodes.
	//
	// example:
	//
	// 0
	WorkerServerlessReplicas *int32 `json:"WorkerServerlessReplicas,omitempty" xml:"WorkerServerlessReplicas,omitempty"`
	// The ID of the Data Management (DMS) workspace.
	//
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

func (s *UpdateAirflowShrinkRequest) GetEnableLineage() *bool {
	return s.EnableLineage
}

func (s *UpdateAirflowShrinkRequest) GetEnableServerless() *bool {
	return s.EnableServerless
}

func (s *UpdateAirflowShrinkRequest) GetGracefulShutdownTimeout() *int32 {
	return s.GracefulShutdownTimeout
}

func (s *UpdateAirflowShrinkRequest) GetLineageReplicas() *int32 {
	return s.LineageReplicas
}

func (s *UpdateAirflowShrinkRequest) GetLineageSpec() *string {
	return s.LineageSpec
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

func (s *UpdateAirflowShrinkRequest) SetEnableLineage(v bool) *UpdateAirflowShrinkRequest {
	s.EnableLineage = &v
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

func (s *UpdateAirflowShrinkRequest) SetLineageReplicas(v int32) *UpdateAirflowShrinkRequest {
	s.LineageReplicas = &v
	return s
}

func (s *UpdateAirflowShrinkRequest) SetLineageSpec(v string) *UpdateAirflowShrinkRequest {
	s.LineageSpec = &v
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
