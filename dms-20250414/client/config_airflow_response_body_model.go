// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAirflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ConfigAirflowResponseBody
	GetAccessDeniedDetail() *string
	SetErrorCode(v string) *ConfigAirflowResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *ConfigAirflowResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ConfigAirflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigAirflowResponseBody
	GetRequestId() *string
	SetRoot(v *ConfigAirflowResponseBodyRoot) *ConfigAirflowResponseBody
	GetRoot() *ConfigAirflowResponseBodyRoot
	SetSuccess(v bool) *ConfigAirflowResponseBody
	GetSuccess() *bool
}

type ConfigAirflowResponseBody struct {
	// Details of the access denial.
	//
	// example:
	//
	// NOT_FOUND
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Information about the Airflow instance.
	Root *ConfigAirflowResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigAirflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigAirflowResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigAirflowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ConfigAirflowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ConfigAirflowResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ConfigAirflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigAirflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigAirflowResponseBody) GetRoot() *ConfigAirflowResponseBodyRoot {
	return s.Root
}

func (s *ConfigAirflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigAirflowResponseBody) SetAccessDeniedDetail(v string) *ConfigAirflowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ConfigAirflowResponseBody) SetErrorCode(v string) *ConfigAirflowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConfigAirflowResponseBody) SetHttpStatusCode(v int64) *ConfigAirflowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigAirflowResponseBody) SetMessage(v string) *ConfigAirflowResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigAirflowResponseBody) SetRequestId(v string) *ConfigAirflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigAirflowResponseBody) SetRoot(v *ConfigAirflowResponseBodyRoot) *ConfigAirflowResponseBody {
	s.Root = v
	return s
}

func (s *ConfigAirflowResponseBody) SetSuccess(v bool) *ConfigAirflowResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigAirflowResponseBody) Validate() error {
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigAirflowResponseBodyRoot struct {
	// The ID of the Airflow instance.
	//
	// example:
	//
	// af-7a6ygsh80dx1jn****
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// The name of the Airflow instance.
	//
	// example:
	//
	// testairflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
	// The Airflow instance specifications.
	//
	// example:
	//
	// SMALL
	AppSpec *string `json:"AppSpec,omitempty" xml:"AppSpec,omitempty"`
	// The application type.
	//
	// example:
	//
	// AIRFLOW
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The custom configuration.
	CustomAirflowCfg []*string `json:"CustomAirflowCfg,omitempty" xml:"CustomAirflowCfg,omitempty" type:"Repeated"`
	// The DAG directory that Airflow scans.
	//
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// The deployment error message.
	//
	// example:
	//
	// vpc not found
	DeployErrorMsg *string `json:"DeployErrorMsg,omitempty" xml:"DeployErrorMsg,omitempty"`
	// The description of the Airflow instance.
	//
	// example:
	//
	// test airflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp when the instance was created.
	//
	// example:
	//
	// 2025-08-12T05:46:01.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// oss-test
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// /airflow
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// The plugin directory that Airflow scans.
	//
	// example:
	//
	// default/plugins
	PluginsDir *string `json:"PluginsDir,omitempty" xml:"PluginsDir,omitempty"`
	// The configuration file for dependencies.
	//
	// example:
	//
	// default/requirements.txt
	RequirementFile *string `json:"RequirementFile,omitempty" xml:"RequirementFile,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2ze1nak7h0alg1xxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The script that runs before Airflow starts.
	//
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// The status of the Airflow instance.
	//
	// example:
	//
	// DEPLOYING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VSwitch ID.
	//
	// example:
	//
	// vsw-bp1931trfxkvf74v****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf63r6coyiw9o5gf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of worker nodes.
	//
	// example:
	//
	// 0
	WorkerServerlessReplicas *int32 `json:"WorkerServerlessReplicas,omitempty" xml:"WorkerServerlessReplicas,omitempty"`
	// The DMS workspace ID.
	//
	// example:
	//
	// 86302423828****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ConfigAirflowResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s ConfigAirflowResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *ConfigAirflowResponseBodyRoot) GetAirflowId() *string {
	return s.AirflowId
}

func (s *ConfigAirflowResponseBodyRoot) GetAirflowName() *string {
	return s.AirflowName
}

func (s *ConfigAirflowResponseBodyRoot) GetAppSpec() *string {
	return s.AppSpec
}

func (s *ConfigAirflowResponseBodyRoot) GetAppType() *string {
	return s.AppType
}

func (s *ConfigAirflowResponseBodyRoot) GetCustomAirflowCfg() []*string {
	return s.CustomAirflowCfg
}

func (s *ConfigAirflowResponseBodyRoot) GetDagsDir() *string {
	return s.DagsDir
}

func (s *ConfigAirflowResponseBodyRoot) GetDeployErrorMsg() *string {
	return s.DeployErrorMsg
}

func (s *ConfigAirflowResponseBodyRoot) GetDescription() *string {
	return s.Description
}

func (s *ConfigAirflowResponseBodyRoot) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ConfigAirflowResponseBodyRoot) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *ConfigAirflowResponseBodyRoot) GetOssPath() *string {
	return s.OssPath
}

func (s *ConfigAirflowResponseBodyRoot) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *ConfigAirflowResponseBodyRoot) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *ConfigAirflowResponseBodyRoot) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ConfigAirflowResponseBodyRoot) GetStartupFile() *string {
	return s.StartupFile
}

func (s *ConfigAirflowResponseBodyRoot) GetStatus() *string {
	return s.Status
}

func (s *ConfigAirflowResponseBodyRoot) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ConfigAirflowResponseBodyRoot) GetVpcId() *string {
	return s.VpcId
}

func (s *ConfigAirflowResponseBodyRoot) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *ConfigAirflowResponseBodyRoot) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ConfigAirflowResponseBodyRoot) GetZoneId() *string {
	return s.ZoneId
}

func (s *ConfigAirflowResponseBodyRoot) SetAirflowId(v string) *ConfigAirflowResponseBodyRoot {
	s.AirflowId = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetAirflowName(v string) *ConfigAirflowResponseBodyRoot {
	s.AirflowName = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetAppSpec(v string) *ConfigAirflowResponseBodyRoot {
	s.AppSpec = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetAppType(v string) *ConfigAirflowResponseBodyRoot {
	s.AppType = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetCustomAirflowCfg(v []*string) *ConfigAirflowResponseBodyRoot {
	s.CustomAirflowCfg = v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetDagsDir(v string) *ConfigAirflowResponseBodyRoot {
	s.DagsDir = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetDeployErrorMsg(v string) *ConfigAirflowResponseBodyRoot {
	s.DeployErrorMsg = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetDescription(v string) *ConfigAirflowResponseBodyRoot {
	s.Description = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetGmtCreated(v string) *ConfigAirflowResponseBodyRoot {
	s.GmtCreated = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetOssBucketName(v string) *ConfigAirflowResponseBodyRoot {
	s.OssBucketName = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetOssPath(v string) *ConfigAirflowResponseBodyRoot {
	s.OssPath = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetPluginsDir(v string) *ConfigAirflowResponseBodyRoot {
	s.PluginsDir = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetRequirementFile(v string) *ConfigAirflowResponseBodyRoot {
	s.RequirementFile = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetSecurityGroupId(v string) *ConfigAirflowResponseBodyRoot {
	s.SecurityGroupId = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetStartupFile(v string) *ConfigAirflowResponseBodyRoot {
	s.StartupFile = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetStatus(v string) *ConfigAirflowResponseBodyRoot {
	s.Status = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetVSwitchId(v string) *ConfigAirflowResponseBodyRoot {
	s.VSwitchId = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetVpcId(v string) *ConfigAirflowResponseBodyRoot {
	s.VpcId = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetWorkerServerlessReplicas(v int32) *ConfigAirflowResponseBodyRoot {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetWorkspaceId(v string) *ConfigAirflowResponseBodyRoot {
	s.WorkspaceId = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) SetZoneId(v string) *ConfigAirflowResponseBodyRoot {
	s.ZoneId = &v
	return s
}

func (s *ConfigAirflowResponseBodyRoot) Validate() error {
	return dara.Validate(s)
}
