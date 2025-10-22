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
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *CreateAirflowResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
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
	// example:
	//
	// af-****
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
	// airflow
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// example:
	//
	// vpc not found
	DeployErrorMsg *string `json:"DeployErrorMsg,omitempty" xml:"DeployErrorMsg,omitempty"`
	// example:
	//
	// order scheduler
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2025-08-12T05:46:01.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// oss-test
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
	// sg-2ze1nak7h0alg1xxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// example:
	//
	// DEPLOYING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *CreateAirflowResponseBodyRoot) GetAppSpec() *string {
	return s.AppSpec
}

func (s *CreateAirflowResponseBodyRoot) GetAppType() *string {
	return s.AppType
}

func (s *CreateAirflowResponseBodyRoot) GetDagsDir() *string {
	return s.DagsDir
}

func (s *CreateAirflowResponseBodyRoot) GetDeployErrorMsg() *string {
	return s.DeployErrorMsg
}

func (s *CreateAirflowResponseBodyRoot) GetDescription() *string {
	return s.Description
}

func (s *CreateAirflowResponseBodyRoot) GetGmtCreated() *string {
	return s.GmtCreated
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

func (s *CreateAirflowResponseBodyRoot) SetAppSpec(v string) *CreateAirflowResponseBodyRoot {
	s.AppSpec = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetAppType(v string) *CreateAirflowResponseBodyRoot {
	s.AppType = &v
	return s
}

func (s *CreateAirflowResponseBodyRoot) SetDagsDir(v string) *CreateAirflowResponseBodyRoot {
	s.DagsDir = &v
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

func (s *CreateAirflowResponseBodyRoot) SetGmtCreated(v string) *CreateAirflowResponseBodyRoot {
	s.GmtCreated = &v
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
	return dara.Validate(s)
}
