// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAirflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAirflowResponseBody
	GetAccessDeniedDetail() *string
	SetErrorCode(v string) *UpdateAirflowResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *UpdateAirflowResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateAirflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAirflowResponseBody
	GetRequestId() *string
	SetRoot(v *UpdateAirflowResponseBodyRoot) *UpdateAirflowResponseBody
	GetRoot() *UpdateAirflowResponseBodyRoot
	SetSuccess(v bool) *UpdateAirflowResponseBody
	GetSuccess() *bool
}

type UpdateAirflowResponseBody struct {
	// example:
	//
	// NOT_FOUND
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Unknown error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E0D21075-CD3E-4D98-8264-F****
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *UpdateAirflowResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAirflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAirflowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAirflowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAirflowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAirflowResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateAirflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAirflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAirflowResponseBody) GetRoot() *UpdateAirflowResponseBodyRoot {
	return s.Root
}

func (s *UpdateAirflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAirflowResponseBody) SetAccessDeniedDetail(v string) *UpdateAirflowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAirflowResponseBody) SetErrorCode(v string) *UpdateAirflowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAirflowResponseBody) SetHttpStatusCode(v int64) *UpdateAirflowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAirflowResponseBody) SetMessage(v string) *UpdateAirflowResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAirflowResponseBody) SetRequestId(v string) *UpdateAirflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAirflowResponseBody) SetRoot(v *UpdateAirflowResponseBodyRoot) *UpdateAirflowResponseBody {
	s.Root = v
	return s
}

func (s *UpdateAirflowResponseBody) SetSuccess(v bool) *UpdateAirflowResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAirflowResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateAirflowResponseBodyRoot struct {
	// example:
	//
	// test airflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
	// SMALL。
	//
	// example:
	//
	// SMALL
	AppSpec *string `json:"AppSpec,omitempty" xml:"AppSpec,omitempty"`
	// example:
	//
	// AIRFLOW
	AppType          *string   `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CustomAirflowCfg []*string `json:"CustomAirflowCfg,omitempty" xml:"CustomAirflowCfg,omitempty" type:"Repeated"`
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// example:
	//
	// quota exists
	DeployErrorMsg *string `json:"DeployErrorMsg,omitempty" xml:"DeployErrorMsg,omitempty"`
	// example:
	//
	// order schedule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// k=v
	Environments *string `json:"Environments,omitempty" xml:"Environments,omitempty"`
	// example:
	//
	// 2025-01-07T15:10:32+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// osstest
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
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
	// example:
	//
	// jieba==0.42
	Requirements *string `json:"Requirements,omitempty" xml:"Requirements,omitempty"`
	// example:
	//
	// sg-2ze1nak7h0alg1w5****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// example:
	//
	// DEPLOYED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// af-xxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// vsw-hp3hyga33aur8tj36****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-bp16ko44pgciwv0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 0
	WorkerServerlessReplicas *int32 `json:"WorkerServerlessReplicas,omitempty" xml:"WorkerServerlessReplicas,omitempty"`
	// example:
	//
	// 86302423828****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateAirflowResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s UpdateAirflowResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *UpdateAirflowResponseBodyRoot) GetAirflowName() *string {
	return s.AirflowName
}

func (s *UpdateAirflowResponseBodyRoot) GetAppSpec() *string {
	return s.AppSpec
}

func (s *UpdateAirflowResponseBodyRoot) GetAppType() *string {
	return s.AppType
}

func (s *UpdateAirflowResponseBodyRoot) GetCustomAirflowCfg() []*string {
	return s.CustomAirflowCfg
}

func (s *UpdateAirflowResponseBodyRoot) GetDagsDir() *string {
	return s.DagsDir
}

func (s *UpdateAirflowResponseBodyRoot) GetDeployErrorMsg() *string {
	return s.DeployErrorMsg
}

func (s *UpdateAirflowResponseBodyRoot) GetDescription() *string {
	return s.Description
}

func (s *UpdateAirflowResponseBodyRoot) GetEnvironments() *string {
	return s.Environments
}

func (s *UpdateAirflowResponseBodyRoot) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *UpdateAirflowResponseBodyRoot) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *UpdateAirflowResponseBodyRoot) GetOssPath() *string {
	return s.OssPath
}

func (s *UpdateAirflowResponseBodyRoot) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *UpdateAirflowResponseBodyRoot) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *UpdateAirflowResponseBodyRoot) GetRequirements() *string {
	return s.Requirements
}

func (s *UpdateAirflowResponseBodyRoot) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateAirflowResponseBodyRoot) GetStartupFile() *string {
	return s.StartupFile
}

func (s *UpdateAirflowResponseBodyRoot) GetStatus() *string {
	return s.Status
}

func (s *UpdateAirflowResponseBodyRoot) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateAirflowResponseBodyRoot) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateAirflowResponseBodyRoot) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateAirflowResponseBodyRoot) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *UpdateAirflowResponseBodyRoot) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateAirflowResponseBodyRoot) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateAirflowResponseBodyRoot) SetAirflowName(v string) *UpdateAirflowResponseBodyRoot {
	s.AirflowName = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetAppSpec(v string) *UpdateAirflowResponseBodyRoot {
	s.AppSpec = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetAppType(v string) *UpdateAirflowResponseBodyRoot {
	s.AppType = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetCustomAirflowCfg(v []*string) *UpdateAirflowResponseBodyRoot {
	s.CustomAirflowCfg = v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetDagsDir(v string) *UpdateAirflowResponseBodyRoot {
	s.DagsDir = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetDeployErrorMsg(v string) *UpdateAirflowResponseBodyRoot {
	s.DeployErrorMsg = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetDescription(v string) *UpdateAirflowResponseBodyRoot {
	s.Description = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetEnvironments(v string) *UpdateAirflowResponseBodyRoot {
	s.Environments = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetGmtCreated(v string) *UpdateAirflowResponseBodyRoot {
	s.GmtCreated = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetOssBucketName(v string) *UpdateAirflowResponseBodyRoot {
	s.OssBucketName = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetOssPath(v string) *UpdateAirflowResponseBodyRoot {
	s.OssPath = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetPluginsDir(v string) *UpdateAirflowResponseBodyRoot {
	s.PluginsDir = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetRequirementFile(v string) *UpdateAirflowResponseBodyRoot {
	s.RequirementFile = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetRequirements(v string) *UpdateAirflowResponseBodyRoot {
	s.Requirements = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetSecurityGroupId(v string) *UpdateAirflowResponseBodyRoot {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetStartupFile(v string) *UpdateAirflowResponseBodyRoot {
	s.StartupFile = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetStatus(v string) *UpdateAirflowResponseBodyRoot {
	s.Status = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetUuid(v string) *UpdateAirflowResponseBodyRoot {
	s.Uuid = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetVSwitchId(v string) *UpdateAirflowResponseBodyRoot {
	s.VSwitchId = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetVpcId(v string) *UpdateAirflowResponseBodyRoot {
	s.VpcId = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetWorkerServerlessReplicas(v int32) *UpdateAirflowResponseBodyRoot {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetWorkspaceId(v string) *UpdateAirflowResponseBodyRoot {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) SetZoneId(v string) *UpdateAirflowResponseBodyRoot {
	s.ZoneId = &v
	return s
}

func (s *UpdateAirflowResponseBodyRoot) Validate() error {
	return dara.Validate(s)
}
