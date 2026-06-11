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
	// Details about the access denial.
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
	// Unknown error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the updated Airflow instance.
	Root *UpdateAirflowResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
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
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAirflowResponseBodyRoot struct {
	// The name of the Airflow instance.
	//
	// example:
	//
	// test airflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
	// The version of Airflow.
	//
	// example:
	//
	// 3.1
	AirflowVersion *string `json:"AirflowVersion,omitempty" xml:"AirflowVersion,omitempty"`
	// The instance specification.
	//
	// example:
	//
	// SMALL
	AppSpec *string `json:"AppSpec,omitempty" xml:"AppSpec,omitempty"`
	// The instance type.
	//
	// example:
	//
	// AIRFLOW
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// A list of custom configuration items.
	CustomAirflowCfg []*string `json:"CustomAirflowCfg,omitempty" xml:"CustomAirflowCfg,omitempty" type:"Repeated"`
	// The directory where DAGs are scanned.
	//
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// A data mount item.
	DataMountInfoList []*DataMountInfo `json:"DataMountInfoList,omitempty" xml:"DataMountInfoList,omitempty" type:"Repeated"`
	// The error message returned upon deployment failure.
	//
	// example:
	//
	// quota exists
	DeployErrorMsg *string `json:"DeployErrorMsg,omitempty" xml:"DeployErrorMsg,omitempty"`
	// The description of the Airflow instance.
	//
	// example:
	//
	// order schedule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether serverless mode is enabled.
	EnableServerless *bool `json:"EnableServerless,omitempty" xml:"EnableServerless,omitempty"`
	// The environment variables.
	//
	// example:
	//
	// k=v
	Environments *string `json:"Environments,omitempty" xml:"Environments,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2025-01-07T15:10:32+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The graceful shutdown timeout, in seconds.
	//
	// example:
	//
	// 60
	GracefulShutdownTimeout *int32 `json:"GracefulShutdownTimeout,omitempty" xml:"GracefulShutdownTimeout,omitempty"`
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// osstest
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The path in OSS where logs are stored.
	//
	// example:
	//
	// /airflow
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// The directory where Airflow plugins are scanned.
	//
	// example:
	//
	// default/plugins
	PluginsDir *string `json:"PluginsDir,omitempty" xml:"PluginsDir,omitempty"`
	// The path to the Python requirements file.
	//
	// example:
	//
	// default/requirements.txt
	RequirementFile *string `json:"RequirementFile,omitempty" xml:"RequirementFile,omitempty"`
	// The Python package requirements.
	//
	// example:
	//
	// jieba==0.42
	Requirements *string `json:"Requirements,omitempty" xml:"Requirements,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze1nak7h0alg1w5****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The startup script for the Airflow container.
	//
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// The instance status.
	//
	// example:
	//
	// DEPLOYED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Airflow instance ID.
	//
	// example:
	//
	// af-xxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-hp3hyga33aur8tj36****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp16ko44pgciwv0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of worker node replicas.
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

func (s *UpdateAirflowResponseBodyRoot) GetAirflowVersion() *string {
	return s.AirflowVersion
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

func (s *UpdateAirflowResponseBodyRoot) GetDataMountInfoList() []*DataMountInfo {
	return s.DataMountInfoList
}

func (s *UpdateAirflowResponseBodyRoot) GetDeployErrorMsg() *string {
	return s.DeployErrorMsg
}

func (s *UpdateAirflowResponseBodyRoot) GetDescription() *string {
	return s.Description
}

func (s *UpdateAirflowResponseBodyRoot) GetEnableServerless() *bool {
	return s.EnableServerless
}

func (s *UpdateAirflowResponseBodyRoot) GetEnvironments() *string {
	return s.Environments
}

func (s *UpdateAirflowResponseBodyRoot) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *UpdateAirflowResponseBodyRoot) GetGracefulShutdownTimeout() *int32 {
	return s.GracefulShutdownTimeout
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

func (s *UpdateAirflowResponseBodyRoot) SetAirflowVersion(v string) *UpdateAirflowResponseBodyRoot {
	s.AirflowVersion = &v
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

func (s *UpdateAirflowResponseBodyRoot) SetDataMountInfoList(v []*DataMountInfo) *UpdateAirflowResponseBodyRoot {
	s.DataMountInfoList = v
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

func (s *UpdateAirflowResponseBodyRoot) SetEnableServerless(v bool) *UpdateAirflowResponseBodyRoot {
	s.EnableServerless = &v
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

func (s *UpdateAirflowResponseBodyRoot) SetGracefulShutdownTimeout(v int32) *UpdateAirflowResponseBodyRoot {
	s.GracefulShutdownTimeout = &v
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
