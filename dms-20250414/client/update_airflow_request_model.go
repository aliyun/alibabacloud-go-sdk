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
	SetDataMountInfoList(v []*DataMountInfo) *UpdateAirflowRequest
	GetDataMountInfoList() []*DataMountInfo
	SetDescription(v string) *UpdateAirflowRequest
	GetDescription() *string
	SetEnableLineage(v bool) *UpdateAirflowRequest
	GetEnableLineage() *bool
	SetEnableServerless(v bool) *UpdateAirflowRequest
	GetEnableServerless() *bool
	SetGracefulShutdownTimeout(v int32) *UpdateAirflowRequest
	GetGracefulShutdownTimeout() *int32
	SetLineageReplicas(v int32) *UpdateAirflowRequest
	GetLineageReplicas() *int32
	SetLineageSpec(v string) *UpdateAirflowRequest
	GetLineageSpec() *string
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
	DataMountInfoList []*DataMountInfo `json:"DataMountInfoList,omitempty" xml:"DataMountInfoList,omitempty" type:"Repeated"`
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

func (s *UpdateAirflowRequest) GetDataMountInfoList() []*DataMountInfo {
	return s.DataMountInfoList
}

func (s *UpdateAirflowRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAirflowRequest) GetEnableLineage() *bool {
	return s.EnableLineage
}

func (s *UpdateAirflowRequest) GetEnableServerless() *bool {
	return s.EnableServerless
}

func (s *UpdateAirflowRequest) GetGracefulShutdownTimeout() *int32 {
	return s.GracefulShutdownTimeout
}

func (s *UpdateAirflowRequest) GetLineageReplicas() *int32 {
	return s.LineageReplicas
}

func (s *UpdateAirflowRequest) GetLineageSpec() *string {
	return s.LineageSpec
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

func (s *UpdateAirflowRequest) SetDataMountInfoList(v []*DataMountInfo) *UpdateAirflowRequest {
	s.DataMountInfoList = v
	return s
}

func (s *UpdateAirflowRequest) SetDescription(v string) *UpdateAirflowRequest {
	s.Description = &v
	return s
}

func (s *UpdateAirflowRequest) SetEnableLineage(v bool) *UpdateAirflowRequest {
	s.EnableLineage = &v
	return s
}

func (s *UpdateAirflowRequest) SetEnableServerless(v bool) *UpdateAirflowRequest {
	s.EnableServerless = &v
	return s
}

func (s *UpdateAirflowRequest) SetGracefulShutdownTimeout(v int32) *UpdateAirflowRequest {
	s.GracefulShutdownTimeout = &v
	return s
}

func (s *UpdateAirflowRequest) SetLineageReplicas(v int32) *UpdateAirflowRequest {
	s.LineageReplicas = &v
	return s
}

func (s *UpdateAirflowRequest) SetLineageSpec(v string) *UpdateAirflowRequest {
	s.LineageSpec = &v
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
	if s.DataMountInfoList != nil {
		for _, item := range s.DataMountInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
