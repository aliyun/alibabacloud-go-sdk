// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAirflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAirflowResponseBody
	GetAccessDeniedDetail() *string
	SetErrorCode(v string) *GetAirflowResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *GetAirflowResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetAirflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAirflowResponseBody
	GetRequestId() *string
	SetRoot(v *GetAirflowResponseBodyRoot) *GetAirflowResponseBody
	GetRoot() *GetAirflowResponseBodyRoot
	SetSuccess(v bool) *GetAirflowResponseBody
	GetSuccess() *bool
}

type GetAirflowResponseBody struct {
	// The access denial details.
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
	// E0D21075-CD3E-4D98-8264-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data details.
	Root *GetAirflowResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAirflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAirflowResponseBody) GoString() string {
	return s.String()
}

func (s *GetAirflowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAirflowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAirflowResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetAirflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAirflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAirflowResponseBody) GetRoot() *GetAirflowResponseBodyRoot {
	return s.Root
}

func (s *GetAirflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAirflowResponseBody) SetAccessDeniedDetail(v string) *GetAirflowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAirflowResponseBody) SetErrorCode(v string) *GetAirflowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAirflowResponseBody) SetHttpStatusCode(v int64) *GetAirflowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAirflowResponseBody) SetMessage(v string) *GetAirflowResponseBody {
	s.Message = &v
	return s
}

func (s *GetAirflowResponseBody) SetRequestId(v string) *GetAirflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAirflowResponseBody) SetRoot(v *GetAirflowResponseBodyRoot) *GetAirflowResponseBody {
	s.Root = v
	return s
}

func (s *GetAirflowResponseBody) SetSuccess(v bool) *GetAirflowResponseBody {
	s.Success = &v
	return s
}

func (s *GetAirflowResponseBody) Validate() error {
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAirflowResponseBodyRoot struct {
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
	// The specifications of the Airflow instance.
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
	// The custom configurations.
	CustomAirflowCfg []*string `json:"CustomAirflowCfg,omitempty" xml:"CustomAirflowCfg,omitempty" type:"Repeated"`
	// The DAG directory scanned by Airflow.
	//
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// The deployment error message.
	//
	// example:
	//
	// Deployed
	DeployErrorMsg *string `json:"DeployErrorMsg,omitempty" xml:"DeployErrorMsg,omitempty"`
	// The description of the Airflow instance.
	//
	// example:
	//
	// test airflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether data lineage is enabled.
	EnableLineage *bool `json:"EnableLineage,omitempty" xml:"EnableLineage,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-08-12T05:46:01.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The number of replicas for the lineage service. The default value is 1 when enabled. If omitted during an update, the original value is retained.
	//
	// example:
	//
	// 1
	LineageReplicas *int32 `json:"LineageReplicas,omitempty" xml:"LineageReplicas,omitempty"`
	// The merged lineage service specifications, which use the same tier as the application specifications.
	//
	// example:
	//
	// SMALL
	LineageSpec *string `json:"LineageSpec,omitempty" xml:"LineageSpec,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// osstest
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// /airflow
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// The plugin directory scanned by Airflow.
	//
	// example:
	//
	// default/plugins
	PluginsDir *string `json:"PluginsDir,omitempty" xml:"PluginsDir,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configuration file for dependency packages.
	//
	// example:
	//
	// default/requirements.txt
	RequirementFile *string `json:"RequirementFile,omitempty" xml:"RequirementFile,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2ze9gj646bkv****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The script loaded before Airflow starts.
	//
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// The Airflow instance status.
	//
	// example:
	//
	// DEPLOYED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1931trfxkvf74v****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zevqv4obraqd5p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of worker nodes.
	//
	// example:
	//
	// 0
	WorkerServerlessReplicas *int32 `json:"WorkerServerlessReplicas,omitempty" xml:"WorkerServerlessReplicas,omitempty"`
	// The Data Management (DMS) workspace ID.
	//
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The zone ID in the region.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetAirflowResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s GetAirflowResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *GetAirflowResponseBodyRoot) GetAirflowId() *string {
	return s.AirflowId
}

func (s *GetAirflowResponseBodyRoot) GetAirflowName() *string {
	return s.AirflowName
}

func (s *GetAirflowResponseBodyRoot) GetAppSpec() *string {
	return s.AppSpec
}

func (s *GetAirflowResponseBodyRoot) GetAppType() *string {
	return s.AppType
}

func (s *GetAirflowResponseBodyRoot) GetCustomAirflowCfg() []*string {
	return s.CustomAirflowCfg
}

func (s *GetAirflowResponseBodyRoot) GetDagsDir() *string {
	return s.DagsDir
}

func (s *GetAirflowResponseBodyRoot) GetDeployErrorMsg() *string {
	return s.DeployErrorMsg
}

func (s *GetAirflowResponseBodyRoot) GetDescription() *string {
	return s.Description
}

func (s *GetAirflowResponseBodyRoot) GetEnableLineage() *bool {
	return s.EnableLineage
}

func (s *GetAirflowResponseBodyRoot) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *GetAirflowResponseBodyRoot) GetLineageReplicas() *int32 {
	return s.LineageReplicas
}

func (s *GetAirflowResponseBodyRoot) GetLineageSpec() *string {
	return s.LineageSpec
}

func (s *GetAirflowResponseBodyRoot) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetAirflowResponseBodyRoot) GetOssPath() *string {
	return s.OssPath
}

func (s *GetAirflowResponseBodyRoot) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *GetAirflowResponseBodyRoot) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAirflowResponseBodyRoot) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *GetAirflowResponseBodyRoot) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetAirflowResponseBodyRoot) GetStartupFile() *string {
	return s.StartupFile
}

func (s *GetAirflowResponseBodyRoot) GetStatus() *string {
	return s.Status
}

func (s *GetAirflowResponseBodyRoot) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetAirflowResponseBodyRoot) GetVpcId() *string {
	return s.VpcId
}

func (s *GetAirflowResponseBodyRoot) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *GetAirflowResponseBodyRoot) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAirflowResponseBodyRoot) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetAirflowResponseBodyRoot) SetAirflowId(v string) *GetAirflowResponseBodyRoot {
	s.AirflowId = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetAirflowName(v string) *GetAirflowResponseBodyRoot {
	s.AirflowName = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetAppSpec(v string) *GetAirflowResponseBodyRoot {
	s.AppSpec = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetAppType(v string) *GetAirflowResponseBodyRoot {
	s.AppType = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetCustomAirflowCfg(v []*string) *GetAirflowResponseBodyRoot {
	s.CustomAirflowCfg = v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetDagsDir(v string) *GetAirflowResponseBodyRoot {
	s.DagsDir = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetDeployErrorMsg(v string) *GetAirflowResponseBodyRoot {
	s.DeployErrorMsg = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetDescription(v string) *GetAirflowResponseBodyRoot {
	s.Description = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetEnableLineage(v bool) *GetAirflowResponseBodyRoot {
	s.EnableLineage = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetGmtCreated(v string) *GetAirflowResponseBodyRoot {
	s.GmtCreated = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetLineageReplicas(v int32) *GetAirflowResponseBodyRoot {
	s.LineageReplicas = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetLineageSpec(v string) *GetAirflowResponseBodyRoot {
	s.LineageSpec = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetOssBucketName(v string) *GetAirflowResponseBodyRoot {
	s.OssBucketName = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetOssPath(v string) *GetAirflowResponseBodyRoot {
	s.OssPath = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetPluginsDir(v string) *GetAirflowResponseBodyRoot {
	s.PluginsDir = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetRegionId(v string) *GetAirflowResponseBodyRoot {
	s.RegionId = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetRequirementFile(v string) *GetAirflowResponseBodyRoot {
	s.RequirementFile = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetSecurityGroupId(v string) *GetAirflowResponseBodyRoot {
	s.SecurityGroupId = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetStartupFile(v string) *GetAirflowResponseBodyRoot {
	s.StartupFile = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetStatus(v string) *GetAirflowResponseBodyRoot {
	s.Status = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetVSwitchId(v string) *GetAirflowResponseBodyRoot {
	s.VSwitchId = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetVpcId(v string) *GetAirflowResponseBodyRoot {
	s.VpcId = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetWorkerServerlessReplicas(v int32) *GetAirflowResponseBodyRoot {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetWorkspaceId(v string) *GetAirflowResponseBodyRoot {
	s.WorkspaceId = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) SetZoneId(v string) *GetAirflowResponseBodyRoot {
	s.ZoneId = &v
	return s
}

func (s *GetAirflowResponseBodyRoot) Validate() error {
	return dara.Validate(s)
}
