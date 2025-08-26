// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAirflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowName(v string) *CreateAirflowRequest
	GetAirflowName() *string
	SetAppSpec(v string) *CreateAirflowRequest
	GetAppSpec() *string
	SetClientToken(v string) *CreateAirflowRequest
	GetClientToken() *string
	SetDagsDir(v string) *CreateAirflowRequest
	GetDagsDir() *string
	SetDescription(v string) *CreateAirflowRequest
	GetDescription() *string
	SetOssBucketName(v string) *CreateAirflowRequest
	GetOssBucketName() *string
	SetOssPath(v string) *CreateAirflowRequest
	GetOssPath() *string
	SetPluginsDir(v string) *CreateAirflowRequest
	GetPluginsDir() *string
	SetRequirementFile(v string) *CreateAirflowRequest
	GetRequirementFile() *string
	SetSecurityGroupId(v string) *CreateAirflowRequest
	GetSecurityGroupId() *string
	SetStartupFile(v string) *CreateAirflowRequest
	GetStartupFile() *string
	SetVSwitchId(v string) *CreateAirflowRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateAirflowRequest
	GetVpcId() *string
	SetWorkerServerlessReplicas(v int32) *CreateAirflowRequest
	GetWorkerServerlessReplicas() *int32
	SetWorkspaceId(v string) *CreateAirflowRequest
	GetWorkspaceId() *string
	SetZoneId(v string) *CreateAirflowRequest
	GetZoneId() *string
}

type CreateAirflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testairflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
	// This parameter is required.
	//
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
	// This parameter is required.
	//
	// example:
	//
	// order scheduler
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss-test
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /airflow
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// example:
	//
	// default/plugins
	PluginsDir *string `json:"PluginsDir,omitempty" xml:"PluginsDir,omitempty"`
	// example:
	//
	// default/requirements.txt
	RequirementFile *string `json:"RequirementFile,omitempty" xml:"RequirementFile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-bp108t8ldzeyk1****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-8vbaf073jawozfp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf63r6coyiw9o5****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	WorkerServerlessReplicas *int32 `json:"WorkerServerlessReplicas,omitempty" xml:"WorkerServerlessReplicas,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateAirflowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowRequest) GoString() string {
	return s.String()
}

func (s *CreateAirflowRequest) GetAirflowName() *string {
	return s.AirflowName
}

func (s *CreateAirflowRequest) GetAppSpec() *string {
	return s.AppSpec
}

func (s *CreateAirflowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAirflowRequest) GetDagsDir() *string {
	return s.DagsDir
}

func (s *CreateAirflowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAirflowRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateAirflowRequest) GetOssPath() *string {
	return s.OssPath
}

func (s *CreateAirflowRequest) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *CreateAirflowRequest) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *CreateAirflowRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateAirflowRequest) GetStartupFile() *string {
	return s.StartupFile
}

func (s *CreateAirflowRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAirflowRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateAirflowRequest) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *CreateAirflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAirflowRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateAirflowRequest) SetAirflowName(v string) *CreateAirflowRequest {
	s.AirflowName = &v
	return s
}

func (s *CreateAirflowRequest) SetAppSpec(v string) *CreateAirflowRequest {
	s.AppSpec = &v
	return s
}

func (s *CreateAirflowRequest) SetClientToken(v string) *CreateAirflowRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAirflowRequest) SetDagsDir(v string) *CreateAirflowRequest {
	s.DagsDir = &v
	return s
}

func (s *CreateAirflowRequest) SetDescription(v string) *CreateAirflowRequest {
	s.Description = &v
	return s
}

func (s *CreateAirflowRequest) SetOssBucketName(v string) *CreateAirflowRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateAirflowRequest) SetOssPath(v string) *CreateAirflowRequest {
	s.OssPath = &v
	return s
}

func (s *CreateAirflowRequest) SetPluginsDir(v string) *CreateAirflowRequest {
	s.PluginsDir = &v
	return s
}

func (s *CreateAirflowRequest) SetRequirementFile(v string) *CreateAirflowRequest {
	s.RequirementFile = &v
	return s
}

func (s *CreateAirflowRequest) SetSecurityGroupId(v string) *CreateAirflowRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateAirflowRequest) SetStartupFile(v string) *CreateAirflowRequest {
	s.StartupFile = &v
	return s
}

func (s *CreateAirflowRequest) SetVSwitchId(v string) *CreateAirflowRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAirflowRequest) SetVpcId(v string) *CreateAirflowRequest {
	s.VpcId = &v
	return s
}

func (s *CreateAirflowRequest) SetWorkerServerlessReplicas(v int32) *CreateAirflowRequest {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *CreateAirflowRequest) SetWorkspaceId(v string) *CreateAirflowRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAirflowRequest) SetZoneId(v string) *CreateAirflowRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateAirflowRequest) Validate() error {
	return dara.Validate(s)
}
