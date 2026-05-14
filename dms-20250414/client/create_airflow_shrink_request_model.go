// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAirflowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowName(v string) *CreateAirflowShrinkRequest
	GetAirflowName() *string
	SetAirflowVersion(v string) *CreateAirflowShrinkRequest
	GetAirflowVersion() *string
	SetAppSpec(v string) *CreateAirflowShrinkRequest
	GetAppSpec() *string
	SetClientToken(v string) *CreateAirflowShrinkRequest
	GetClientToken() *string
	SetDagsDir(v string) *CreateAirflowShrinkRequest
	GetDagsDir() *string
	SetDataMountInfoListShrink(v string) *CreateAirflowShrinkRequest
	GetDataMountInfoListShrink() *string
	SetDescription(v string) *CreateAirflowShrinkRequest
	GetDescription() *string
	SetEnableServerless(v bool) *CreateAirflowShrinkRequest
	GetEnableServerless() *bool
	SetGracefulShutdownTimeout(v int32) *CreateAirflowShrinkRequest
	GetGracefulShutdownTimeout() *int32
	SetOssBucketName(v string) *CreateAirflowShrinkRequest
	GetOssBucketName() *string
	SetOssPath(v string) *CreateAirflowShrinkRequest
	GetOssPath() *string
	SetPluginsDir(v string) *CreateAirflowShrinkRequest
	GetPluginsDir() *string
	SetRequirementFile(v string) *CreateAirflowShrinkRequest
	GetRequirementFile() *string
	SetSecurityGroupId(v string) *CreateAirflowShrinkRequest
	GetSecurityGroupId() *string
	SetStartupFile(v string) *CreateAirflowShrinkRequest
	GetStartupFile() *string
	SetVSwitchId(v string) *CreateAirflowShrinkRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateAirflowShrinkRequest
	GetVpcId() *string
	SetWorkerServerlessReplicas(v int32) *CreateAirflowShrinkRequest
	GetWorkerServerlessReplicas() *int32
	SetWorkspaceId(v string) *CreateAirflowShrinkRequest
	GetWorkspaceId() *string
	SetZoneId(v string) *CreateAirflowShrinkRequest
	GetZoneId() *string
}

type CreateAirflowShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testairflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
	// example:
	//
	// 3.1
	AirflowVersion *string `json:"AirflowVersion,omitempty" xml:"AirflowVersion,omitempty"`
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
	DagsDir                 *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	DataMountInfoListShrink *string `json:"DataMountInfoList,omitempty" xml:"DataMountInfoList,omitempty"`
	// example:
	//
	// order scheduler
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableServerless *bool   `json:"EnableServerless,omitempty" xml:"EnableServerless,omitempty"`
	// example:
	//
	// 60
	GracefulShutdownTimeout *int32 `json:"GracefulShutdownTimeout,omitempty" xml:"GracefulShutdownTimeout,omitempty"`
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
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateAirflowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAirflowShrinkRequest) GetAirflowName() *string {
	return s.AirflowName
}

func (s *CreateAirflowShrinkRequest) GetAirflowVersion() *string {
	return s.AirflowVersion
}

func (s *CreateAirflowShrinkRequest) GetAppSpec() *string {
	return s.AppSpec
}

func (s *CreateAirflowShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAirflowShrinkRequest) GetDagsDir() *string {
	return s.DagsDir
}

func (s *CreateAirflowShrinkRequest) GetDataMountInfoListShrink() *string {
	return s.DataMountInfoListShrink
}

func (s *CreateAirflowShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAirflowShrinkRequest) GetEnableServerless() *bool {
	return s.EnableServerless
}

func (s *CreateAirflowShrinkRequest) GetGracefulShutdownTimeout() *int32 {
	return s.GracefulShutdownTimeout
}

func (s *CreateAirflowShrinkRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateAirflowShrinkRequest) GetOssPath() *string {
	return s.OssPath
}

func (s *CreateAirflowShrinkRequest) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *CreateAirflowShrinkRequest) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *CreateAirflowShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateAirflowShrinkRequest) GetStartupFile() *string {
	return s.StartupFile
}

func (s *CreateAirflowShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAirflowShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateAirflowShrinkRequest) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *CreateAirflowShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAirflowShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateAirflowShrinkRequest) SetAirflowName(v string) *CreateAirflowShrinkRequest {
	s.AirflowName = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetAirflowVersion(v string) *CreateAirflowShrinkRequest {
	s.AirflowVersion = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetAppSpec(v string) *CreateAirflowShrinkRequest {
	s.AppSpec = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetClientToken(v string) *CreateAirflowShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetDagsDir(v string) *CreateAirflowShrinkRequest {
	s.DagsDir = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetDataMountInfoListShrink(v string) *CreateAirflowShrinkRequest {
	s.DataMountInfoListShrink = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetDescription(v string) *CreateAirflowShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetEnableServerless(v bool) *CreateAirflowShrinkRequest {
	s.EnableServerless = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetGracefulShutdownTimeout(v int32) *CreateAirflowShrinkRequest {
	s.GracefulShutdownTimeout = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetOssBucketName(v string) *CreateAirflowShrinkRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetOssPath(v string) *CreateAirflowShrinkRequest {
	s.OssPath = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetPluginsDir(v string) *CreateAirflowShrinkRequest {
	s.PluginsDir = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetRequirementFile(v string) *CreateAirflowShrinkRequest {
	s.RequirementFile = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetSecurityGroupId(v string) *CreateAirflowShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetStartupFile(v string) *CreateAirflowShrinkRequest {
	s.StartupFile = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetVSwitchId(v string) *CreateAirflowShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetVpcId(v string) *CreateAirflowShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetWorkerServerlessReplicas(v int32) *CreateAirflowShrinkRequest {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetWorkspaceId(v string) *CreateAirflowShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAirflowShrinkRequest) SetZoneId(v string) *CreateAirflowShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateAirflowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
