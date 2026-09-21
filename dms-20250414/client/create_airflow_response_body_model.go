// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAirflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAirflowResponseBody
	GetAccessDeniedDetail() *string
	SetErrorCode(v string) *CreateAirflowResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *CreateAirflowResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateAirflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAirflowResponseBody
	GetRequestId() *string
	SetRoot(v *CreateAirflowResponseBodyRoot) *CreateAirflowResponseBody
	GetRoot() *CreateAirflowResponseBodyRoot
	SetSuccess(v bool) *CreateAirflowResponseBody
	GetSuccess() *bool
}

type CreateAirflowResponseBody struct {
	// The details of the access denial.
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
	// The log information.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Airflow data dictionary.
	Root *CreateAirflowResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAirflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAirflowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAirflowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAirflowResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateAirflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAirflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAirflowResponseBody) GetRoot() *CreateAirflowResponseBodyRoot {
	return s.Root
}

func (s *CreateAirflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAirflowResponseBody) SetAccessDeniedDetail(v string) *CreateAirflowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAirflowResponseBody) SetErrorCode(v string) *CreateAirflowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAirflowResponseBody) SetHttpStatusCode(v int64) *CreateAirflowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAirflowResponseBody) SetMessage(v string) *CreateAirflowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAirflowResponseBody) SetRequestId(v string) *CreateAirflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAirflowResponseBody) SetRoot(v *CreateAirflowResponseBodyRoot) *CreateAirflowResponseBody {
	s.Root = v
	return s
}

func (s *CreateAirflowResponseBody) SetSuccess(v bool) *CreateAirflowResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAirflowResponseBody) Validate() error {
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAirflowResponseBodyRoot struct {
	// The instance ID of the Airflow instance.
	//
	// example:
	//
	// af-****
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// The name of the Airflow instance.
	//
	// example:
	//
	// testairflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
	// The Airflow version.
	//
	// example:
	//
	// 3.1
	AirflowVersion *string `json:"AirflowVersion,omitempty" xml:"AirflowVersion,omitempty"`
	// The specifications of the Airflow instance.
	//
	// example:
	//
	// SMALL
	AppSpec *string `json:"AppSpec,omitempty" xml:"AppSpec,omitempty"`
	// The type.
	//
	// example:
	//
	// airflow
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The custom Airflow configuration.
	CustomAirflowCfg []*string `json:"CustomAirflowCfg,omitempty" xml:"CustomAirflowCfg,omitempty" type:"Repeated"`
	// The directory that Airflow scans for DAGs.
	//
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// The list of data mount information.
	DataMountInfoList []*DataMountInfo `json:"DataMountInfoList,omitempty" xml:"DataMountInfoList,omitempty" type:"Repeated"`
	// The error log information.
	//
	// example:
	//
	// vpc not found
	DeployErrorMsg *string `json:"DeployErrorMsg,omitempty" xml:"DeployErrorMsg,omitempty"`
	// The description of the Airflow instance.
	//
	// example:
	//
	// order scheduler
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable Worker elasticity.
	EnableServerless *bool `json:"EnableServerless,omitempty" xml:"EnableServerless,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-08-12T05:46:01.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The timeout period for the Worker to gracefully shut down.
	//
	// example:
	//
	// 60
	GracefulShutdownTimeout *int32 `json:"GracefulShutdownTimeout,omitempty" xml:"GracefulShutdownTimeout,omitempty"`
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
	// The directory that Airflow scans for plugins.
	//
	// example:
	//
	// default/plugins
	PluginsDir *string `json:"PluginsDir,omitempty" xml:"PluginsDir,omitempty"`
	// The path of the installable package.
	//
	// example:
	//
	// default/requirements.txt
	RequirementFile *string `json:"RequirementFile,omitempty" xml:"RequirementFile,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze1nak7h0alg1xxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The file loaded when the Airflow instance starts. You can use this file to set environment variables and other configurations.
	//
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// The instance status of the Airflow instance.
	//
	// example:
	//
	// DEPLOYING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-8vbaf073jawozfpbg****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-uf63r6coyiw9o5gf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of worker scale-out nodes.
	//
	// example:
	//
	// 0
	WorkerServerlessReplicas *int32 `json:"WorkerServerlessReplicas,omitempty" xml:"WorkerServerlessReplicas,omitempty"`
	// The workspace ID.
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

func (s CreateAirflowResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *CreateAirflowResponseBodyRoot) GetAirflowId() *string {
	return s.AirflowId
}

func (s *CreateAirflowResponseBodyRoot) GetAirflowName() *string {
	return s.AirflowName
}

func (s *CreateAirflowResponseBodyRoot) GetAirflowVersion() *string {
	return s.AirflowVersion
}

func (s *CreateAirflowResponseBodyRoot) GetAppSpec() *string {
	return s.AppSpec
}

func (s *CreateAirflowResponseBodyRoot) GetAppType() *string {
	return s.AppType
}

func (s *CreateAirflowResponseBodyRoot) GetCustomAirflowCfg() []*string {
	return s.CustomAirflowCfg
}

func (s *CreateAirflowResponseBodyRoot) GetDagsDir() *string {
	return s.DagsDir
}

func (s *CreateAirflowResponseBodyRoot) GetDataMountInfoList() []*DataMountInfo {
	return s.DataMountInfoList
}

func (s *CreateAirflowResponseBodyRoot) GetDeployErrorMsg() *string {
	return s.DeployErrorMsg
}

func (s *CreateAirflowResponseBodyRoot) GetDescription() *string {
	return s.Description
}

func (s *CreateAirflowResponseBodyRoot) GetEnableServerless() *bool {
	return s.EnableServerless
}

func (s *CreateAirflowResponseBodyRoot) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateAirflowResponseBodyRoot) GetGracefulShutdownTimeout() *int32 {
	return s.GracefulShutdownTimeout
}

func (s *CreateAirflowResponseBodyRoot) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateAirflowResponseBodyRoot) GetOssPath() *string {
	return s.OssPath
}

func (s *CreateAirflowResponseBodyRoot) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *CreateAirflowResponseBodyRoot) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *CreateAirflowResponseBodyRoot) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateAirflowResponseBodyRoot) GetStartupFile() *string {
	return s.StartupFile
}

func (s *CreateAirflowResponseBodyRoot) GetStatus() *string {
	return s.Status
}

func (s *CreateAirflowResponseBodyRoot) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAirflowResponseBodyRoot) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateAirflowResponseBodyRoot) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *CreateAirflowResponseBodyRoot) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAirflowResponseBodyRoot) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateAirflowResponseBodyRoot) SetAirflowId(v string) *CreateAirflowResponseBodyRoot {
	s.AirflowId = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetAirflowName(v string) *CreateAirflowResponseBodyRoot {
	s.AirflowName = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetAirflowVersion(v string) *CreateAirflowResponseBodyRoot {
	s.AirflowVersion = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetAppSpec(v string) *CreateAirflowResponseBodyRoot {
	s.AppSpec = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetAppType(v string) *CreateAirflowResponseBodyRoot {
	s.AppType = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetCustomAirflowCfg(v []*string) *CreateAirflowResponseBodyRoot {
	s.CustomAirflowCfg = v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetDagsDir(v string) *CreateAirflowResponseBodyRoot {
	s.DagsDir = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetDataMountInfoList(v []*DataMountInfo) *CreateAirflowResponseBodyRoot {
	s.DataMountInfoList = v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetDeployErrorMsg(v string) *CreateAirflowResponseBodyRoot {
	s.DeployErrorMsg = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetDescription(v string) *CreateAirflowResponseBodyRoot {
	s.Description = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetEnableServerless(v bool) *CreateAirflowResponseBodyRoot {
	s.EnableServerless = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetGmtCreated(v string) *CreateAirflowResponseBodyRoot {
	s.GmtCreated = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetGracefulShutdownTimeout(v int32) *CreateAirflowResponseBodyRoot {
	s.GracefulShutdownTimeout = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetOssBucketName(v string) *CreateAirflowResponseBodyRoot {
	s.OssBucketName = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetOssPath(v string) *CreateAirflowResponseBodyRoot {
	s.OssPath = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetPluginsDir(v string) *CreateAirflowResponseBodyRoot {
	s.PluginsDir = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetRequirementFile(v string) *CreateAirflowResponseBodyRoot {
	s.RequirementFile = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetSecurityGroupId(v string) *CreateAirflowResponseBodyRoot {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetStartupFile(v string) *CreateAirflowResponseBodyRoot {
	s.StartupFile = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetStatus(v string) *CreateAirflowResponseBodyRoot {
	s.Status = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetVSwitchId(v string) *CreateAirflowResponseBodyRoot {
	s.VSwitchId = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetVpcId(v string) *CreateAirflowResponseBodyRoot {
	s.VpcId = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetWorkerServerlessReplicas(v int32) *CreateAirflowResponseBodyRoot {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetWorkspaceId(v string) *CreateAirflowResponseBodyRoot {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetZoneId(v string) *CreateAirflowResponseBodyRoot {
	s.ZoneId = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) Validate() error {
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
