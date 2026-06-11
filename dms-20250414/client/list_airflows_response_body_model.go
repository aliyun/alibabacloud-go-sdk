// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAirflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAirflowsResponseBody
	GetAccessDeniedDetail() *string
	SetErrorCode(v string) *ListAirflowsResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int64) *ListAirflowsResponseBody
	GetHttpStatusCode() *int64
	SetMaxResults(v int32) *ListAirflowsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAirflowsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAirflowsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAirflowsResponseBody
	GetRequestId() *string
	SetRoot(v *ListAirflowsResponseBodyRoot) *ListAirflowsResponseBody
	GetRoot() *ListAirflowsResponseBodyRoot
	SetSuccess(v bool) *ListAirflowsResponseBody
	GetSuccess() *bool
}

type ListAirflowsResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// NOT_FOUND
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The error code. This parameter is returned when an error occurs.
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
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The error message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the data.
	Root *ListAirflowsResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
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

func (s ListAirflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAirflowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAirflowsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAirflowsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAirflowsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListAirflowsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAirflowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAirflowsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAirflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAirflowsResponseBody) GetRoot() *ListAirflowsResponseBodyRoot {
	return s.Root
}

func (s *ListAirflowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAirflowsResponseBody) SetAccessDeniedDetail(v string) *ListAirflowsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAirflowsResponseBody) SetErrorCode(v string) *ListAirflowsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAirflowsResponseBody) SetHttpStatusCode(v int64) *ListAirflowsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAirflowsResponseBody) SetMaxResults(v int32) *ListAirflowsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAirflowsResponseBody) SetMessage(v string) *ListAirflowsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAirflowsResponseBody) SetNextToken(v string) *ListAirflowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAirflowsResponseBody) SetRequestId(v string) *ListAirflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAirflowsResponseBody) SetRoot(v *ListAirflowsResponseBodyRoot) *ListAirflowsResponseBody {
	s.Root = v
	return s
}

func (s *ListAirflowsResponseBody) SetSuccess(v bool) *ListAirflowsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAirflowsResponseBody) Validate() error {
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAirflowsResponseBodyRoot struct {
	// The list of data.
	List []*ListAirflowsResponseBodyRootList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of matching Airflow instances.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAirflowsResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s ListAirflowsResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *ListAirflowsResponseBodyRoot) GetList() []*ListAirflowsResponseBodyRootList {
	return s.List
}

func (s *ListAirflowsResponseBodyRoot) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAirflowsResponseBodyRoot) SetList(v []*ListAirflowsResponseBodyRootList) *ListAirflowsResponseBodyRoot {
	s.List = v
	return s
}

func (s *ListAirflowsResponseBodyRoot) SetTotalCount(v int32) *ListAirflowsResponseBodyRoot {
	s.TotalCount = &v
	return s
}

func (s *ListAirflowsResponseBodyRoot) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAirflowsResponseBodyRootList struct {
	// The ID of the Airflow instance.
	//
	// example:
	//
	// af-7a6ygsh80d****
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// The name of the Airflow instance.
	//
	// example:
	//
	// test-airflow
	AirflowName *string `json:"AirflowName,omitempty" xml:"AirflowName,omitempty"`
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
	// AIRFLOW
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The custom configuration.
	CustomAirflowCfg []*string `json:"CustomAirflowCfg,omitempty" xml:"CustomAirflowCfg,omitempty" type:"Repeated"`
	// The DAGs folder that Airflow scans.
	//
	// example:
	//
	// default/dags
	DagsDir *string `json:"DagsDir,omitempty" xml:"DagsDir,omitempty"`
	// The error message returned when the deployment fails.
	//
	// example:
	//
	// quota exists
	DeployErrorMsg *string `json:"DeployErrorMsg,omitempty" xml:"DeployErrorMsg,omitempty"`
	// The description of the Airflow instance.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the Airflow instance was created.
	//
	// example:
	//
	// 2025-08-12T05:46:01.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// osstest
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The path in OSS.
	//
	// example:
	//
	// /airflow
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// The plugins folder that Airflow scans.
	//
	// example:
	//
	// default/plugins
	PluginsDir *string `json:"PluginsDir,omitempty" xml:"PluginsDir,omitempty"`
	// The path to the file that contains the dependencies.
	//
	// example:
	//
	// default/requirements.txt
	RequirementFile *string `json:"RequirementFile,omitempty" xml:"RequirementFile,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2ze1nak7h0alg1w****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The path to the script that is loaded when Airflow starts.
	//
	// example:
	//
	// default/startup.sh
	StartupFile *string `json:"StartupFile,omitempty" xml:"StartupFile,omitempty"`
	// The status of the Airflow instance.
	//
	// example:
	//
	// DEPLOYED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6sxdc22x7sbdb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-8vbbfm33dy0y1pek****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of worker nodes.
	//
	// example:
	//
	// 0
	WorkerServerlessReplicas *int32 `json:"WorkerServerlessReplicas,omitempty" xml:"WorkerServerlessReplicas,omitempty"`
	// The ID of the DMS workspace.
	//
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The ID of the zone in the region.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListAirflowsResponseBodyRootList) String() string {
	return dara.Prettify(s)
}

func (s ListAirflowsResponseBodyRootList) GoString() string {
	return s.String()
}

func (s *ListAirflowsResponseBodyRootList) GetAirflowId() *string {
	return s.AirflowId
}

func (s *ListAirflowsResponseBodyRootList) GetAirflowName() *string {
	return s.AirflowName
}

func (s *ListAirflowsResponseBodyRootList) GetAppSpec() *string {
	return s.AppSpec
}

func (s *ListAirflowsResponseBodyRootList) GetAppType() *string {
	return s.AppType
}

func (s *ListAirflowsResponseBodyRootList) GetCustomAirflowCfg() []*string {
	return s.CustomAirflowCfg
}

func (s *ListAirflowsResponseBodyRootList) GetDagsDir() *string {
	return s.DagsDir
}

func (s *ListAirflowsResponseBodyRootList) GetDeployErrorMsg() *string {
	return s.DeployErrorMsg
}

func (s *ListAirflowsResponseBodyRootList) GetDescription() *string {
	return s.Description
}

func (s *ListAirflowsResponseBodyRootList) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListAirflowsResponseBodyRootList) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *ListAirflowsResponseBodyRootList) GetOssPath() *string {
	return s.OssPath
}

func (s *ListAirflowsResponseBodyRootList) GetPluginsDir() *string {
	return s.PluginsDir
}

func (s *ListAirflowsResponseBodyRootList) GetRequirementFile() *string {
	return s.RequirementFile
}

func (s *ListAirflowsResponseBodyRootList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListAirflowsResponseBodyRootList) GetStartupFile() *string {
	return s.StartupFile
}

func (s *ListAirflowsResponseBodyRootList) GetStatus() *string {
	return s.Status
}

func (s *ListAirflowsResponseBodyRootList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListAirflowsResponseBodyRootList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAirflowsResponseBodyRootList) GetWorkerServerlessReplicas() *int32 {
	return s.WorkerServerlessReplicas
}

func (s *ListAirflowsResponseBodyRootList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAirflowsResponseBodyRootList) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListAirflowsResponseBodyRootList) SetAirflowId(v string) *ListAirflowsResponseBodyRootList {
	s.AirflowId = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetAirflowName(v string) *ListAirflowsResponseBodyRootList {
	s.AirflowName = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetAppSpec(v string) *ListAirflowsResponseBodyRootList {
	s.AppSpec = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetAppType(v string) *ListAirflowsResponseBodyRootList {
	s.AppType = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetCustomAirflowCfg(v []*string) *ListAirflowsResponseBodyRootList {
	s.CustomAirflowCfg = v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetDagsDir(v string) *ListAirflowsResponseBodyRootList {
	s.DagsDir = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetDeployErrorMsg(v string) *ListAirflowsResponseBodyRootList {
	s.DeployErrorMsg = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetDescription(v string) *ListAirflowsResponseBodyRootList {
	s.Description = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetGmtCreated(v string) *ListAirflowsResponseBodyRootList {
	s.GmtCreated = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetOssBucketName(v string) *ListAirflowsResponseBodyRootList {
	s.OssBucketName = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetOssPath(v string) *ListAirflowsResponseBodyRootList {
	s.OssPath = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetPluginsDir(v string) *ListAirflowsResponseBodyRootList {
	s.PluginsDir = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetRequirementFile(v string) *ListAirflowsResponseBodyRootList {
	s.RequirementFile = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetSecurityGroupId(v string) *ListAirflowsResponseBodyRootList {
	s.SecurityGroupId = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetStartupFile(v string) *ListAirflowsResponseBodyRootList {
	s.StartupFile = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetStatus(v string) *ListAirflowsResponseBodyRootList {
	s.Status = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetVSwitchId(v string) *ListAirflowsResponseBodyRootList {
	s.VSwitchId = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetVpcId(v string) *ListAirflowsResponseBodyRootList {
	s.VpcId = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetWorkerServerlessReplicas(v int32) *ListAirflowsResponseBodyRootList {
	s.WorkerServerlessReplicas = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetWorkspaceId(v string) *ListAirflowsResponseBodyRootList {
	s.WorkspaceId = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) SetZoneId(v string) *ListAirflowsResponseBodyRootList {
	s.ZoneId = &v
	return s
}

func (s *ListAirflowsResponseBodyRootList) Validate() error {
	return dara.Validate(s)
}
