// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type OtsDetail struct {
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s OtsDetail) String() string {
	return tea.Prettify(s)
}

func (s OtsDetail) GoString() string {
	return s.String()
}

func (s *OtsDetail) SetTableNames(v []*string) *OtsDetail {
	s.TableNames = v
	return s
}

type OtsTableRestoreDetail struct {
	BatchChannelCount         *int32  `json:"BatchChannelCount,omitempty" xml:"BatchChannelCount,omitempty"`
	IndexNameSuffix           *string `json:"IndexNameSuffix,omitempty" xml:"IndexNameSuffix,omitempty"`
	OverwriteExisting         *bool   `json:"OverwriteExisting,omitempty" xml:"OverwriteExisting,omitempty"`
	ReGenerateAutoIncrementPK *bool   `json:"ReGenerateAutoIncrementPK,omitempty" xml:"ReGenerateAutoIncrementPK,omitempty"`
	RestoreIndex              *bool   `json:"RestoreIndex,omitempty" xml:"RestoreIndex,omitempty"`
	RestoreSearchIndex        *bool   `json:"RestoreSearchIndex,omitempty" xml:"RestoreSearchIndex,omitempty"`
	SearchIndexNameSuffix     *string `json:"SearchIndexNameSuffix,omitempty" xml:"SearchIndexNameSuffix,omitempty"`
}

func (s OtsTableRestoreDetail) String() string {
	return tea.Prettify(s)
}

func (s OtsTableRestoreDetail) GoString() string {
	return s.String()
}

func (s *OtsTableRestoreDetail) SetBatchChannelCount(v int32) *OtsTableRestoreDetail {
	s.BatchChannelCount = &v
	return s
}

func (s *OtsTableRestoreDetail) SetIndexNameSuffix(v string) *OtsTableRestoreDetail {
	s.IndexNameSuffix = &v
	return s
}

func (s *OtsTableRestoreDetail) SetOverwriteExisting(v bool) *OtsTableRestoreDetail {
	s.OverwriteExisting = &v
	return s
}

func (s *OtsTableRestoreDetail) SetReGenerateAutoIncrementPK(v bool) *OtsTableRestoreDetail {
	s.ReGenerateAutoIncrementPK = &v
	return s
}

func (s *OtsTableRestoreDetail) SetRestoreIndex(v bool) *OtsTableRestoreDetail {
	s.RestoreIndex = &v
	return s
}

func (s *OtsTableRestoreDetail) SetRestoreSearchIndex(v bool) *OtsTableRestoreDetail {
	s.RestoreSearchIndex = &v
	return s
}

func (s *OtsTableRestoreDetail) SetSearchIndexNameSuffix(v string) *OtsTableRestoreDetail {
	s.SearchIndexNameSuffix = &v
	return s
}

type Report struct {
	FailedFiles  *string `json:"FailedFiles,omitempty" xml:"FailedFiles,omitempty"`
	SkippedFiles *string `json:"SkippedFiles,omitempty" xml:"SkippedFiles,omitempty"`
	SuccessFiles *string `json:"SuccessFiles,omitempty" xml:"SuccessFiles,omitempty"`
	TotalFiles   *string `json:"TotalFiles,omitempty" xml:"TotalFiles,omitempty"`
}

func (s Report) String() string {
	return tea.Prettify(s)
}

func (s Report) GoString() string {
	return s.String()
}

func (s *Report) SetFailedFiles(v string) *Report {
	s.FailedFiles = &v
	return s
}

func (s *Report) SetSkippedFiles(v string) *Report {
	s.SkippedFiles = &v
	return s
}

func (s *Report) SetSuccessFiles(v string) *Report {
	s.SuccessFiles = &v
	return s
}

func (s *Report) SetTotalFiles(v string) *Report {
	s.TotalFiles = &v
	return s
}

type Rule struct {
	BackupType           *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DestinationRegionId  *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationRetention *int64  `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	Disabled             *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DoCopy               *bool   `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	Retention            *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Schedule             *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s Rule) String() string {
	return tea.Prettify(s)
}

func (s Rule) GoString() string {
	return s.String()
}

func (s *Rule) SetBackupType(v string) *Rule {
	s.BackupType = &v
	return s
}

func (s *Rule) SetDestinationRegionId(v string) *Rule {
	s.DestinationRegionId = &v
	return s
}

func (s *Rule) SetDestinationRetention(v int64) *Rule {
	s.DestinationRetention = &v
	return s
}

func (s *Rule) SetDisabled(v bool) *Rule {
	s.Disabled = &v
	return s
}

func (s *Rule) SetDoCopy(v bool) *Rule {
	s.DoCopy = &v
	return s
}

func (s *Rule) SetRetention(v int64) *Rule {
	s.Retention = &v
	return s
}

func (s *Rule) SetRuleName(v string) *Rule {
	s.RuleName = &v
	return s
}

func (s *Rule) SetSchedule(v string) *Rule {
	s.Schedule = &v
	return s
}

type AddContainerClusterRequest struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Identifier  *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s AddContainerClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s AddContainerClusterRequest) GoString() string {
	return s.String()
}

func (s *AddContainerClusterRequest) SetClusterType(v string) *AddContainerClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *AddContainerClusterRequest) SetDescription(v string) *AddContainerClusterRequest {
	s.Description = &v
	return s
}

func (s *AddContainerClusterRequest) SetIdentifier(v string) *AddContainerClusterRequest {
	s.Identifier = &v
	return s
}

func (s *AddContainerClusterRequest) SetName(v string) *AddContainerClusterRequest {
	s.Name = &v
	return s
}

func (s *AddContainerClusterRequest) SetNetworkType(v string) *AddContainerClusterRequest {
	s.NetworkType = &v
	return s
}

type AddContainerClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s AddContainerClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddContainerClusterResponseBody) GoString() string {
	return s.String()
}

func (s *AddContainerClusterResponseBody) SetClusterId(v string) *AddContainerClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetCode(v string) *AddContainerClusterResponseBody {
	s.Code = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetMessage(v string) *AddContainerClusterResponseBody {
	s.Message = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetRequestId(v string) *AddContainerClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetSuccess(v bool) *AddContainerClusterResponseBody {
	s.Success = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetToken(v string) *AddContainerClusterResponseBody {
	s.Token = &v
	return s
}

type AddContainerClusterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddContainerClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddContainerClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s AddContainerClusterResponse) GoString() string {
	return s.String()
}

func (s *AddContainerClusterResponse) SetHeaders(v map[string]*string) *AddContainerClusterResponse {
	s.Headers = v
	return s
}

func (s *AddContainerClusterResponse) SetStatusCode(v int32) *AddContainerClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContainerClusterResponse) SetBody(v *AddContainerClusterResponseBody) *AddContainerClusterResponse {
	s.Body = v
	return s
}

type CancelBackupJobRequest struct {
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CancelBackupJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelBackupJobRequest) GoString() string {
	return s.String()
}

func (s *CancelBackupJobRequest) SetJobId(v string) *CancelBackupJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelBackupJobRequest) SetVaultId(v string) *CancelBackupJobRequest {
	s.VaultId = &v
	return s
}

type CancelBackupJobResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelBackupJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelBackupJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelBackupJobResponseBody) SetCode(v string) *CancelBackupJobResponseBody {
	s.Code = &v
	return s
}

func (s *CancelBackupJobResponseBody) SetMessage(v string) *CancelBackupJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelBackupJobResponseBody) SetRequestId(v string) *CancelBackupJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelBackupJobResponseBody) SetSuccess(v bool) *CancelBackupJobResponseBody {
	s.Success = &v
	return s
}

type CancelBackupJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelBackupJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelBackupJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelBackupJobResponse) GoString() string {
	return s.String()
}

func (s *CancelBackupJobResponse) SetHeaders(v map[string]*string) *CancelBackupJobResponse {
	s.Headers = v
	return s
}

func (s *CancelBackupJobResponse) SetStatusCode(v int32) *CancelBackupJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelBackupJobResponse) SetBody(v *CancelBackupJobResponseBody) *CancelBackupJobResponse {
	s.Body = v
	return s
}

type CancelRestoreJobRequest struct {
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	VaultId   *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CancelRestoreJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CancelRestoreJobRequest) SetRestoreId(v string) *CancelRestoreJobRequest {
	s.RestoreId = &v
	return s
}

func (s *CancelRestoreJobRequest) SetVaultId(v string) *CancelRestoreJobRequest {
	s.VaultId = &v
	return s
}

type CancelRestoreJobResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelRestoreJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRestoreJobResponseBody) SetCode(v string) *CancelRestoreJobResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRestoreJobResponseBody) SetMessage(v string) *CancelRestoreJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelRestoreJobResponseBody) SetRequestId(v string) *CancelRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRestoreJobResponseBody) SetSuccess(v bool) *CancelRestoreJobResponseBody {
	s.Success = &v
	return s
}

type CancelRestoreJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelRestoreJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelRestoreJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRestoreJobResponse) GoString() string {
	return s.String()
}

func (s *CancelRestoreJobResponse) SetHeaders(v map[string]*string) *CancelRestoreJobResponse {
	s.Headers = v
	return s
}

func (s *CancelRestoreJobResponse) SetStatusCode(v int32) *CancelRestoreJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRestoreJobResponse) SetBody(v *CancelRestoreJobResponseBody) *CancelRestoreJobResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetCode(v string) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetSuccess(v bool) *ChangeResourceGroupResponseBody {
	s.Success = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CreateBackupJobRequest struct {
	BackupType           *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerClusterId   *string `json:"ContainerClusterId,omitempty" xml:"ContainerClusterId,omitempty"`
	ContainerResources   *string `json:"ContainerResources,omitempty" xml:"ContainerResources,omitempty"`
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	Exclude              *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	Include              *string `json:"Include,omitempty" xml:"Include,omitempty"`
	InitiatedByAck       *bool   `json:"InitiatedByAck,omitempty" xml:"InitiatedByAck,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Options              *string `json:"Options,omitempty" xml:"Options,omitempty"`
	Retention            *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	SourceType           *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SpeedLimit           *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	VaultId              *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateBackupJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupJobRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupJobRequest) SetBackupType(v string) *CreateBackupJobRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupJobRequest) SetClusterId(v string) *CreateBackupJobRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateBackupJobRequest) SetContainerClusterId(v string) *CreateBackupJobRequest {
	s.ContainerClusterId = &v
	return s
}

func (s *CreateBackupJobRequest) SetContainerResources(v string) *CreateBackupJobRequest {
	s.ContainerResources = &v
	return s
}

func (s *CreateBackupJobRequest) SetCrossAccountRoleName(v string) *CreateBackupJobRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateBackupJobRequest) SetCrossAccountType(v string) *CreateBackupJobRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateBackupJobRequest) SetCrossAccountUserId(v int64) *CreateBackupJobRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateBackupJobRequest) SetExclude(v string) *CreateBackupJobRequest {
	s.Exclude = &v
	return s
}

func (s *CreateBackupJobRequest) SetInclude(v string) *CreateBackupJobRequest {
	s.Include = &v
	return s
}

func (s *CreateBackupJobRequest) SetInitiatedByAck(v bool) *CreateBackupJobRequest {
	s.InitiatedByAck = &v
	return s
}

func (s *CreateBackupJobRequest) SetInstanceId(v string) *CreateBackupJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBackupJobRequest) SetJobName(v string) *CreateBackupJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateBackupJobRequest) SetOptions(v string) *CreateBackupJobRequest {
	s.Options = &v
	return s
}

func (s *CreateBackupJobRequest) SetRetention(v int64) *CreateBackupJobRequest {
	s.Retention = &v
	return s
}

func (s *CreateBackupJobRequest) SetSourceType(v string) *CreateBackupJobRequest {
	s.SourceType = &v
	return s
}

func (s *CreateBackupJobRequest) SetSpeedLimit(v string) *CreateBackupJobRequest {
	s.SpeedLimit = &v
	return s
}

func (s *CreateBackupJobRequest) SetVaultId(v string) *CreateBackupJobRequest {
	s.VaultId = &v
	return s
}

type CreateBackupJobResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBackupJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupJobResponseBody) SetCode(v string) *CreateBackupJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBackupJobResponseBody) SetJobId(v string) *CreateBackupJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateBackupJobResponseBody) SetMessage(v string) *CreateBackupJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBackupJobResponseBody) SetRequestId(v string) *CreateBackupJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupJobResponseBody) SetSuccess(v bool) *CreateBackupJobResponseBody {
	s.Success = &v
	return s
}

type CreateBackupJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBackupJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackupJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupJobResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupJobResponse) SetHeaders(v map[string]*string) *CreateBackupJobResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupJobResponse) SetStatusCode(v int32) *CreateBackupJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupJobResponse) SetBody(v *CreateBackupJobResponseBody) *CreateBackupJobResponse {
	s.Body = v
	return s
}

type CreateBackupPlanRequest struct {
	BackupType           *string                        `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	Bucket               *string                        `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	CreateTime           *int64                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossAccountRoleName *string                        `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                        `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                         `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	Detail               map[string]interface{}         `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Exclude              *string                        `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	FileSystemId         *string                        `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Include              *string                        `json:"Include,omitempty" xml:"Include,omitempty"`
	InstanceId           *string                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Options              *string                        `json:"Options,omitempty" xml:"Options,omitempty"`
	OtsDetail            *OtsDetail                     `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	Path                 []*string                      `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	PlanName             *string                        `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	Prefix               *string                        `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Retention            *int64                         `json:"Retention,omitempty" xml:"Retention,omitempty"`
	Rule                 []*CreateBackupPlanRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	Schedule             *string                        `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	SourceType           *string                        `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SpeedLimit           *string                        `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	UdmRegionId          *string                        `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
	VaultId              *string                        `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequest) SetBackupType(v string) *CreateBackupPlanRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetBucket(v string) *CreateBackupPlanRequest {
	s.Bucket = &v
	return s
}

func (s *CreateBackupPlanRequest) SetCreateTime(v int64) *CreateBackupPlanRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateBackupPlanRequest) SetCrossAccountRoleName(v string) *CreateBackupPlanRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateBackupPlanRequest) SetCrossAccountType(v string) *CreateBackupPlanRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetCrossAccountUserId(v int64) *CreateBackupPlanRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDetail(v map[string]interface{}) *CreateBackupPlanRequest {
	s.Detail = v
	return s
}

func (s *CreateBackupPlanRequest) SetExclude(v string) *CreateBackupPlanRequest {
	s.Exclude = &v
	return s
}

func (s *CreateBackupPlanRequest) SetFileSystemId(v string) *CreateBackupPlanRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInclude(v string) *CreateBackupPlanRequest {
	s.Include = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInstanceId(v string) *CreateBackupPlanRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInstanceName(v string) *CreateBackupPlanRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateBackupPlanRequest) SetOptions(v string) *CreateBackupPlanRequest {
	s.Options = &v
	return s
}

func (s *CreateBackupPlanRequest) SetOtsDetail(v *OtsDetail) *CreateBackupPlanRequest {
	s.OtsDetail = v
	return s
}

func (s *CreateBackupPlanRequest) SetPath(v []*string) *CreateBackupPlanRequest {
	s.Path = v
	return s
}

func (s *CreateBackupPlanRequest) SetPlanName(v string) *CreateBackupPlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateBackupPlanRequest) SetPrefix(v string) *CreateBackupPlanRequest {
	s.Prefix = &v
	return s
}

func (s *CreateBackupPlanRequest) SetRetention(v int64) *CreateBackupPlanRequest {
	s.Retention = &v
	return s
}

func (s *CreateBackupPlanRequest) SetRule(v []*CreateBackupPlanRequestRule) *CreateBackupPlanRequest {
	s.Rule = v
	return s
}

func (s *CreateBackupPlanRequest) SetSchedule(v string) *CreateBackupPlanRequest {
	s.Schedule = &v
	return s
}

func (s *CreateBackupPlanRequest) SetSourceType(v string) *CreateBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetSpeedLimit(v string) *CreateBackupPlanRequest {
	s.SpeedLimit = &v
	return s
}

func (s *CreateBackupPlanRequest) SetUdmRegionId(v string) *CreateBackupPlanRequest {
	s.UdmRegionId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetVaultId(v string) *CreateBackupPlanRequest {
	s.VaultId = &v
	return s
}

type CreateBackupPlanRequestRule struct {
	BackupType           *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DestinationRegionId  *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationRetention *int64  `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	Disabled             *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DoCopy               *bool   `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	Retention            *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Schedule             *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s CreateBackupPlanRequestRule) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanRequestRule) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequestRule) SetBackupType(v string) *CreateBackupPlanRequestRule {
	s.BackupType = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetDestinationRegionId(v string) *CreateBackupPlanRequestRule {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetDestinationRetention(v int64) *CreateBackupPlanRequestRule {
	s.DestinationRetention = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetDisabled(v bool) *CreateBackupPlanRequestRule {
	s.Disabled = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetDoCopy(v bool) *CreateBackupPlanRequestRule {
	s.DoCopy = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetRetention(v int64) *CreateBackupPlanRequestRule {
	s.Retention = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetRuleName(v string) *CreateBackupPlanRequestRule {
	s.RuleName = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetSchedule(v string) *CreateBackupPlanRequestRule {
	s.Schedule = &v
	return s
}

type CreateBackupPlanShrinkRequest struct {
	BackupType           *string                              `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	Bucket               *string                              `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	CreateTime           *int64                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossAccountRoleName *string                              `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                              `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                               `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	DetailShrink         *string                              `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Exclude              *string                              `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	FileSystemId         *string                              `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Include              *string                              `json:"Include,omitempty" xml:"Include,omitempty"`
	InstanceId           *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string                              `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Options              *string                              `json:"Options,omitempty" xml:"Options,omitempty"`
	OtsDetailShrink      *string                              `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	Path                 []*string                            `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	PlanName             *string                              `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	Prefix               *string                              `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Retention            *int64                               `json:"Retention,omitempty" xml:"Retention,omitempty"`
	Rule                 []*CreateBackupPlanShrinkRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	Schedule             *string                              `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	SourceType           *string                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SpeedLimit           *string                              `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	UdmRegionId          *string                              `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
	VaultId              *string                              `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateBackupPlanShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanShrinkRequest) SetBackupType(v string) *CreateBackupPlanShrinkRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetBucket(v string) *CreateBackupPlanShrinkRequest {
	s.Bucket = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetCreateTime(v int64) *CreateBackupPlanShrinkRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetCrossAccountRoleName(v string) *CreateBackupPlanShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetCrossAccountType(v string) *CreateBackupPlanShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetCrossAccountUserId(v int64) *CreateBackupPlanShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetDetailShrink(v string) *CreateBackupPlanShrinkRequest {
	s.DetailShrink = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetExclude(v string) *CreateBackupPlanShrinkRequest {
	s.Exclude = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetFileSystemId(v string) *CreateBackupPlanShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetInclude(v string) *CreateBackupPlanShrinkRequest {
	s.Include = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetInstanceId(v string) *CreateBackupPlanShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetInstanceName(v string) *CreateBackupPlanShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetOptions(v string) *CreateBackupPlanShrinkRequest {
	s.Options = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetOtsDetailShrink(v string) *CreateBackupPlanShrinkRequest {
	s.OtsDetailShrink = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetPath(v []*string) *CreateBackupPlanShrinkRequest {
	s.Path = v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetPlanName(v string) *CreateBackupPlanShrinkRequest {
	s.PlanName = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetPrefix(v string) *CreateBackupPlanShrinkRequest {
	s.Prefix = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetRetention(v int64) *CreateBackupPlanShrinkRequest {
	s.Retention = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetRule(v []*CreateBackupPlanShrinkRequestRule) *CreateBackupPlanShrinkRequest {
	s.Rule = v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetSchedule(v string) *CreateBackupPlanShrinkRequest {
	s.Schedule = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetSourceType(v string) *CreateBackupPlanShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetSpeedLimit(v string) *CreateBackupPlanShrinkRequest {
	s.SpeedLimit = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetUdmRegionId(v string) *CreateBackupPlanShrinkRequest {
	s.UdmRegionId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetVaultId(v string) *CreateBackupPlanShrinkRequest {
	s.VaultId = &v
	return s
}

type CreateBackupPlanShrinkRequestRule struct {
	BackupType           *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DestinationRegionId  *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationRetention *int64  `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	Disabled             *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DoCopy               *bool   `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	Retention            *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Schedule             *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s CreateBackupPlanShrinkRequestRule) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanShrinkRequestRule) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanShrinkRequestRule) SetBackupType(v string) *CreateBackupPlanShrinkRequestRule {
	s.BackupType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetDestinationRegionId(v string) *CreateBackupPlanShrinkRequestRule {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetDestinationRetention(v int64) *CreateBackupPlanShrinkRequestRule {
	s.DestinationRetention = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetDisabled(v bool) *CreateBackupPlanShrinkRequestRule {
	s.Disabled = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetDoCopy(v bool) *CreateBackupPlanShrinkRequestRule {
	s.DoCopy = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetRetention(v int64) *CreateBackupPlanShrinkRequestRule {
	s.Retention = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetRuleName(v string) *CreateBackupPlanShrinkRequestRule {
	s.RuleName = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetSchedule(v string) *CreateBackupPlanShrinkRequestRule {
	s.Schedule = &v
	return s
}

type CreateBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PlanId    *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponseBody) SetCode(v string) *CreateBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetMessage(v string) *CreateBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetPlanId(v string) *CreateBackupPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetRequestId(v string) *CreateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetSuccess(v bool) *CreateBackupPlanResponseBody {
	s.Success = &v
	return s
}

type CreateBackupPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponse) SetHeaders(v map[string]*string) *CreateBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupPlanResponse) SetStatusCode(v int32) *CreateBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupPlanResponse) SetBody(v *CreateBackupPlanResponseBody) *CreateBackupPlanResponse {
	s.Body = v
	return s
}

type CreateClientsRequest struct {
	AlertSetting    *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	ClientInfo      *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	UseHttps        *bool   `json:"UseHttps,omitempty" xml:"UseHttps,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClientsRequest) GoString() string {
	return s.String()
}

func (s *CreateClientsRequest) SetAlertSetting(v string) *CreateClientsRequest {
	s.AlertSetting = &v
	return s
}

func (s *CreateClientsRequest) SetClientInfo(v string) *CreateClientsRequest {
	s.ClientInfo = &v
	return s
}

func (s *CreateClientsRequest) SetResourceGroupId(v string) *CreateClientsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClientsRequest) SetUseHttps(v bool) *CreateClientsRequest {
	s.UseHttps = &v
	return s
}

func (s *CreateClientsRequest) SetVaultId(v string) *CreateClientsRequest {
	s.VaultId = &v
	return s
}

type CreateClientsResponseBody struct {
	Code             *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceStatuses *CreateClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Struct"`
	Message          *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId           *string                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClientsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientsResponseBody) SetCode(v string) *CreateClientsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateClientsResponseBody) SetInstanceStatuses(v *CreateClientsResponseBodyInstanceStatuses) *CreateClientsResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *CreateClientsResponseBody) SetMessage(v string) *CreateClientsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateClientsResponseBody) SetRequestId(v string) *CreateClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientsResponseBody) SetSuccess(v bool) *CreateClientsResponseBody {
	s.Success = &v
	return s
}

func (s *CreateClientsResponseBody) SetTaskId(v string) *CreateClientsResponseBody {
	s.TaskId = &v
	return s
}

type CreateClientsResponseBodyInstanceStatuses struct {
	InstanceStatus []*CreateClientsResponseBodyInstanceStatusesInstanceStatus `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty" type:"Repeated"`
}

func (s CreateClientsResponseBodyInstanceStatuses) String() string {
	return tea.Prettify(s)
}

func (s CreateClientsResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *CreateClientsResponseBodyInstanceStatuses) SetInstanceStatus(v []*CreateClientsResponseBodyInstanceStatusesInstanceStatus) *CreateClientsResponseBodyInstanceStatuses {
	s.InstanceStatus = v
	return s
}

type CreateClientsResponseBodyInstanceStatusesInstanceStatus struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ValidInstance *bool   `json:"ValidInstance,omitempty" xml:"ValidInstance,omitempty"`
}

func (s CreateClientsResponseBodyInstanceStatusesInstanceStatus) String() string {
	return tea.Prettify(s)
}

func (s CreateClientsResponseBodyInstanceStatusesInstanceStatus) GoString() string {
	return s.String()
}

func (s *CreateClientsResponseBodyInstanceStatusesInstanceStatus) SetInstanceId(v string) *CreateClientsResponseBodyInstanceStatusesInstanceStatus {
	s.InstanceId = &v
	return s
}

func (s *CreateClientsResponseBodyInstanceStatusesInstanceStatus) SetValidInstance(v bool) *CreateClientsResponseBodyInstanceStatusesInstanceStatus {
	s.ValidInstance = &v
	return s
}

type CreateClientsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClientsResponse) GoString() string {
	return s.String()
}

func (s *CreateClientsResponse) SetHeaders(v map[string]*string) *CreateClientsResponse {
	s.Headers = v
	return s
}

func (s *CreateClientsResponse) SetStatusCode(v int32) *CreateClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientsResponse) SetBody(v *CreateClientsResponseBody) *CreateClientsResponse {
	s.Body = v
	return s
}

type CreateHanaBackupPlanRequest struct {
	BackupPrefix    *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	BackupType      *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PlanName        *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Schedule        *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateHanaBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateHanaBackupPlanRequest) SetBackupPrefix(v string) *CreateHanaBackupPlanRequest {
	s.BackupPrefix = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetBackupType(v string) *CreateHanaBackupPlanRequest {
	s.BackupType = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetClusterId(v string) *CreateHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetDatabaseName(v string) *CreateHanaBackupPlanRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetPlanName(v string) *CreateHanaBackupPlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetResourceGroupId(v string) *CreateHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetSchedule(v string) *CreateHanaBackupPlanRequest {
	s.Schedule = &v
	return s
}

func (s *CreateHanaBackupPlanRequest) SetVaultId(v string) *CreateHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

type CreateHanaBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PlanId    *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHanaBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHanaBackupPlanResponseBody) SetCode(v string) *CreateHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) SetMessage(v string) *CreateHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) SetPlanId(v string) *CreateHanaBackupPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) SetRequestId(v string) *CreateHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHanaBackupPlanResponseBody) SetSuccess(v bool) *CreateHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

type CreateHanaBackupPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHanaBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateHanaBackupPlanResponse) SetHeaders(v map[string]*string) *CreateHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateHanaBackupPlanResponse) SetStatusCode(v int32) *CreateHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHanaBackupPlanResponse) SetBody(v *CreateHanaBackupPlanResponseBody) *CreateHanaBackupPlanResponse {
	s.Body = v
	return s
}

type CreateHanaInstanceRequest struct {
	AlertSetting        *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	EcsInstanceId       *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	HanaName            *string `json:"HanaName,omitempty" xml:"HanaName,omitempty"`
	Host                *string `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceNumber      *int32  `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	Password            *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sid                 *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	UseSsl              *bool   `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	UserName            *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ValidateCertificate *bool   `json:"ValidateCertificate,omitempty" xml:"ValidateCertificate,omitempty"`
	VaultId             *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateHanaInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateHanaInstanceRequest) SetAlertSetting(v string) *CreateHanaInstanceRequest {
	s.AlertSetting = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetEcsInstanceId(v string) *CreateHanaInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetHanaName(v string) *CreateHanaInstanceRequest {
	s.HanaName = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetHost(v string) *CreateHanaInstanceRequest {
	s.Host = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetInstanceNumber(v int32) *CreateHanaInstanceRequest {
	s.InstanceNumber = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetPassword(v string) *CreateHanaInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetResourceGroupId(v string) *CreateHanaInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetSid(v string) *CreateHanaInstanceRequest {
	s.Sid = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetUseSsl(v bool) *CreateHanaInstanceRequest {
	s.UseSsl = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetUserName(v string) *CreateHanaInstanceRequest {
	s.UserName = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetValidateCertificate(v bool) *CreateHanaInstanceRequest {
	s.ValidateCertificate = &v
	return s
}

func (s *CreateHanaInstanceRequest) SetVaultId(v string) *CreateHanaInstanceRequest {
	s.VaultId = &v
	return s
}

type CreateHanaInstanceResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHanaInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHanaInstanceResponseBody) SetClusterId(v string) *CreateHanaInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) SetCode(v string) *CreateHanaInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) SetMessage(v string) *CreateHanaInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) SetRequestId(v string) *CreateHanaInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) SetSuccess(v bool) *CreateHanaInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateHanaInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHanaInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHanaInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateHanaInstanceResponse) SetHeaders(v map[string]*string) *CreateHanaInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateHanaInstanceResponse) SetStatusCode(v int32) *CreateHanaInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHanaInstanceResponse) SetBody(v *CreateHanaInstanceResponseBody) *CreateHanaInstanceResponse {
	s.Body = v
	return s
}

type CreateHanaRestoreRequest struct {
	BackupId            *int64  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupPrefix        *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	CheckAccess         *bool   `json:"CheckAccess,omitempty" xml:"CheckAccess,omitempty"`
	ClearLog            *bool   `json:"ClearLog,omitempty" xml:"ClearLog,omitempty"`
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName        *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	LogPosition         *int64  `json:"LogPosition,omitempty" xml:"LogPosition,omitempty"`
	Mode                *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RecoveryPointInTime *int64  `json:"RecoveryPointInTime,omitempty" xml:"RecoveryPointInTime,omitempty"`
	SidAdmin            *string `json:"SidAdmin,omitempty" xml:"SidAdmin,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceClusterId     *string `json:"SourceClusterId,omitempty" xml:"SourceClusterId,omitempty"`
	SystemCopy          *bool   `json:"SystemCopy,omitempty" xml:"SystemCopy,omitempty"`
	UseCatalog          *bool   `json:"UseCatalog,omitempty" xml:"UseCatalog,omitempty"`
	UseDelta            *bool   `json:"UseDelta,omitempty" xml:"UseDelta,omitempty"`
	VolumeId            *int32  `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
}

func (s CreateHanaRestoreRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaRestoreRequest) GoString() string {
	return s.String()
}

func (s *CreateHanaRestoreRequest) SetBackupId(v int64) *CreateHanaRestoreRequest {
	s.BackupId = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetBackupPrefix(v string) *CreateHanaRestoreRequest {
	s.BackupPrefix = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetCheckAccess(v bool) *CreateHanaRestoreRequest {
	s.CheckAccess = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetClearLog(v bool) *CreateHanaRestoreRequest {
	s.ClearLog = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetClusterId(v string) *CreateHanaRestoreRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetDatabaseName(v string) *CreateHanaRestoreRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetLogPosition(v int64) *CreateHanaRestoreRequest {
	s.LogPosition = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetMode(v string) *CreateHanaRestoreRequest {
	s.Mode = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetRecoveryPointInTime(v int64) *CreateHanaRestoreRequest {
	s.RecoveryPointInTime = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetSidAdmin(v string) *CreateHanaRestoreRequest {
	s.SidAdmin = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetSource(v string) *CreateHanaRestoreRequest {
	s.Source = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetSourceClusterId(v string) *CreateHanaRestoreRequest {
	s.SourceClusterId = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetSystemCopy(v bool) *CreateHanaRestoreRequest {
	s.SystemCopy = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetUseCatalog(v bool) *CreateHanaRestoreRequest {
	s.UseCatalog = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetUseDelta(v bool) *CreateHanaRestoreRequest {
	s.UseDelta = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetVolumeId(v int32) *CreateHanaRestoreRequest {
	s.VolumeId = &v
	return s
}

type CreateHanaRestoreResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHanaRestoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaRestoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHanaRestoreResponseBody) SetCode(v string) *CreateHanaRestoreResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) SetMessage(v string) *CreateHanaRestoreResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) SetRequestId(v string) *CreateHanaRestoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) SetRestoreId(v string) *CreateHanaRestoreResponseBody {
	s.RestoreId = &v
	return s
}

func (s *CreateHanaRestoreResponseBody) SetSuccess(v bool) *CreateHanaRestoreResponseBody {
	s.Success = &v
	return s
}

type CreateHanaRestoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHanaRestoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHanaRestoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHanaRestoreResponse) GoString() string {
	return s.String()
}

func (s *CreateHanaRestoreResponse) SetHeaders(v map[string]*string) *CreateHanaRestoreResponse {
	s.Headers = v
	return s
}

func (s *CreateHanaRestoreResponse) SetStatusCode(v int32) *CreateHanaRestoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHanaRestoreResponse) SetBody(v *CreateHanaRestoreResponseBody) *CreateHanaRestoreResponse {
	s.Body = v
	return s
}

type CreateReplicationVaultRequest struct {
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RedundancyType            *string `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	ReplicationSourceRegionId *string `json:"ReplicationSourceRegionId,omitempty" xml:"ReplicationSourceRegionId,omitempty"`
	ReplicationSourceVaultId  *string `json:"ReplicationSourceVaultId,omitempty" xml:"ReplicationSourceVaultId,omitempty"`
	VaultName                 *string `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
	VaultRegionId             *string `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
	VaultStorageClass         *string `json:"VaultStorageClass,omitempty" xml:"VaultStorageClass,omitempty"`
}

func (s CreateReplicationVaultRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationVaultRequest) GoString() string {
	return s.String()
}

func (s *CreateReplicationVaultRequest) SetDescription(v string) *CreateReplicationVaultRequest {
	s.Description = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetRedundancyType(v string) *CreateReplicationVaultRequest {
	s.RedundancyType = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetReplicationSourceRegionId(v string) *CreateReplicationVaultRequest {
	s.ReplicationSourceRegionId = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetReplicationSourceVaultId(v string) *CreateReplicationVaultRequest {
	s.ReplicationSourceVaultId = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetVaultName(v string) *CreateReplicationVaultRequest {
	s.VaultName = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetVaultRegionId(v string) *CreateReplicationVaultRequest {
	s.VaultRegionId = &v
	return s
}

func (s *CreateReplicationVaultRequest) SetVaultStorageClass(v string) *CreateReplicationVaultRequest {
	s.VaultStorageClass = &v
	return s
}

type CreateReplicationVaultResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VaultId   *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateReplicationVaultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationVaultResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReplicationVaultResponseBody) SetCode(v string) *CreateReplicationVaultResponseBody {
	s.Code = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetMessage(v string) *CreateReplicationVaultResponseBody {
	s.Message = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetRequestId(v string) *CreateReplicationVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetSuccess(v bool) *CreateReplicationVaultResponseBody {
	s.Success = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetTaskId(v string) *CreateReplicationVaultResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateReplicationVaultResponseBody) SetVaultId(v string) *CreateReplicationVaultResponseBody {
	s.VaultId = &v
	return s
}

type CreateReplicationVaultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateReplicationVaultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateReplicationVaultResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationVaultResponse) GoString() string {
	return s.String()
}

func (s *CreateReplicationVaultResponse) SetHeaders(v map[string]*string) *CreateReplicationVaultResponse {
	s.Headers = v
	return s
}

func (s *CreateReplicationVaultResponse) SetStatusCode(v int32) *CreateReplicationVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReplicationVaultResponse) SetBody(v *CreateReplicationVaultResponseBody) *CreateReplicationVaultResponse {
	s.Body = v
	return s
}

type CreateRestoreJobRequest struct {
	CrossAccountRoleName     *string                `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType         *string                `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId       *int64                 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	Exclude                  *string                `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	Include                  *string                `json:"Include,omitempty" xml:"Include,omitempty"`
	InitiatedByAck           *bool                  `json:"InitiatedByAck,omitempty" xml:"InitiatedByAck,omitempty"`
	OtsDetail                *OtsTableRestoreDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	RestoreType              *string                `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SnapshotHash             *string                `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	SnapshotId               *string                `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SourceType               *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TargetBucket             *string                `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	TargetContainer          *string                `json:"TargetContainer,omitempty" xml:"TargetContainer,omitempty"`
	TargetContainerClusterId *string                `json:"TargetContainerClusterId,omitempty" xml:"TargetContainerClusterId,omitempty"`
	TargetCreateTime         *int64                 `json:"TargetCreateTime,omitempty" xml:"TargetCreateTime,omitempty"`
	TargetFileSystemId       *string                `json:"TargetFileSystemId,omitempty" xml:"TargetFileSystemId,omitempty"`
	TargetInstanceId         *string                `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	TargetInstanceName       *string                `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	TargetPath               *string                `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPrefix             *string                `json:"TargetPrefix,omitempty" xml:"TargetPrefix,omitempty"`
	TargetTableName          *string                `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	TargetTime               *int64                 `json:"TargetTime,omitempty" xml:"TargetTime,omitempty"`
	UdmDetail                map[string]interface{} `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty"`
	VaultId                  *string                `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateRestoreJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobRequest) SetCrossAccountRoleName(v string) *CreateRestoreJobRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateRestoreJobRequest) SetCrossAccountType(v string) *CreateRestoreJobRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateRestoreJobRequest) SetCrossAccountUserId(v int64) *CreateRestoreJobRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetExclude(v string) *CreateRestoreJobRequest {
	s.Exclude = &v
	return s
}

func (s *CreateRestoreJobRequest) SetInclude(v string) *CreateRestoreJobRequest {
	s.Include = &v
	return s
}

func (s *CreateRestoreJobRequest) SetInitiatedByAck(v bool) *CreateRestoreJobRequest {
	s.InitiatedByAck = &v
	return s
}

func (s *CreateRestoreJobRequest) SetOtsDetail(v *OtsTableRestoreDetail) *CreateRestoreJobRequest {
	s.OtsDetail = v
	return s
}

func (s *CreateRestoreJobRequest) SetRestoreType(v string) *CreateRestoreJobRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSnapshotHash(v string) *CreateRestoreJobRequest {
	s.SnapshotHash = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSnapshotId(v string) *CreateRestoreJobRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSourceType(v string) *CreateRestoreJobRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetBucket(v string) *CreateRestoreJobRequest {
	s.TargetBucket = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetContainer(v string) *CreateRestoreJobRequest {
	s.TargetContainer = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetContainerClusterId(v string) *CreateRestoreJobRequest {
	s.TargetContainerClusterId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetCreateTime(v int64) *CreateRestoreJobRequest {
	s.TargetCreateTime = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetFileSystemId(v string) *CreateRestoreJobRequest {
	s.TargetFileSystemId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetInstanceId(v string) *CreateRestoreJobRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetInstanceName(v string) *CreateRestoreJobRequest {
	s.TargetInstanceName = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetPath(v string) *CreateRestoreJobRequest {
	s.TargetPath = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetPrefix(v string) *CreateRestoreJobRequest {
	s.TargetPrefix = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetTableName(v string) *CreateRestoreJobRequest {
	s.TargetTableName = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetTime(v int64) *CreateRestoreJobRequest {
	s.TargetTime = &v
	return s
}

func (s *CreateRestoreJobRequest) SetUdmDetail(v map[string]interface{}) *CreateRestoreJobRequest {
	s.UdmDetail = v
	return s
}

func (s *CreateRestoreJobRequest) SetVaultId(v string) *CreateRestoreJobRequest {
	s.VaultId = &v
	return s
}

type CreateRestoreJobShrinkRequest struct {
	CrossAccountRoleName     *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType         *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId       *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	Exclude                  *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	Include                  *string `json:"Include,omitempty" xml:"Include,omitempty"`
	InitiatedByAck           *bool   `json:"InitiatedByAck,omitempty" xml:"InitiatedByAck,omitempty"`
	OtsDetailShrink          *string `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	RestoreType              *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SnapshotHash             *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	SnapshotId               *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SourceType               *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TargetBucket             *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	TargetContainer          *string `json:"TargetContainer,omitempty" xml:"TargetContainer,omitempty"`
	TargetContainerClusterId *string `json:"TargetContainerClusterId,omitempty" xml:"TargetContainerClusterId,omitempty"`
	TargetCreateTime         *int64  `json:"TargetCreateTime,omitempty" xml:"TargetCreateTime,omitempty"`
	TargetFileSystemId       *string `json:"TargetFileSystemId,omitempty" xml:"TargetFileSystemId,omitempty"`
	TargetInstanceId         *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	TargetInstanceName       *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	TargetPath               *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPrefix             *string `json:"TargetPrefix,omitempty" xml:"TargetPrefix,omitempty"`
	TargetTableName          *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	TargetTime               *int64  `json:"TargetTime,omitempty" xml:"TargetTime,omitempty"`
	UdmDetailShrink          *string `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty"`
	VaultId                  *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateRestoreJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRestoreJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobShrinkRequest) SetCrossAccountRoleName(v string) *CreateRestoreJobShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetCrossAccountType(v string) *CreateRestoreJobShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetCrossAccountUserId(v int64) *CreateRestoreJobShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetExclude(v string) *CreateRestoreJobShrinkRequest {
	s.Exclude = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetInclude(v string) *CreateRestoreJobShrinkRequest {
	s.Include = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetInitiatedByAck(v bool) *CreateRestoreJobShrinkRequest {
	s.InitiatedByAck = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetOtsDetailShrink(v string) *CreateRestoreJobShrinkRequest {
	s.OtsDetailShrink = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetRestoreType(v string) *CreateRestoreJobShrinkRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetSnapshotHash(v string) *CreateRestoreJobShrinkRequest {
	s.SnapshotHash = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetSnapshotId(v string) *CreateRestoreJobShrinkRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetSourceType(v string) *CreateRestoreJobShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetBucket(v string) *CreateRestoreJobShrinkRequest {
	s.TargetBucket = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetContainer(v string) *CreateRestoreJobShrinkRequest {
	s.TargetContainer = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetContainerClusterId(v string) *CreateRestoreJobShrinkRequest {
	s.TargetContainerClusterId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetCreateTime(v int64) *CreateRestoreJobShrinkRequest {
	s.TargetCreateTime = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetFileSystemId(v string) *CreateRestoreJobShrinkRequest {
	s.TargetFileSystemId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetInstanceId(v string) *CreateRestoreJobShrinkRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetInstanceName(v string) *CreateRestoreJobShrinkRequest {
	s.TargetInstanceName = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetPath(v string) *CreateRestoreJobShrinkRequest {
	s.TargetPath = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetPrefix(v string) *CreateRestoreJobShrinkRequest {
	s.TargetPrefix = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetTableName(v string) *CreateRestoreJobShrinkRequest {
	s.TargetTableName = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetTime(v int64) *CreateRestoreJobShrinkRequest {
	s.TargetTime = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetUdmDetailShrink(v string) *CreateRestoreJobShrinkRequest {
	s.UdmDetailShrink = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetVaultId(v string) *CreateRestoreJobShrinkRequest {
	s.VaultId = &v
	return s
}

type CreateRestoreJobResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRestoreJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobResponseBody) SetCode(v string) *CreateRestoreJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRestoreJobResponseBody) SetMessage(v string) *CreateRestoreJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRestoreJobResponseBody) SetRequestId(v string) *CreateRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRestoreJobResponseBody) SetRestoreId(v string) *CreateRestoreJobResponseBody {
	s.RestoreId = &v
	return s
}

func (s *CreateRestoreJobResponseBody) SetSuccess(v bool) *CreateRestoreJobResponseBody {
	s.Success = &v
	return s
}

type CreateRestoreJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRestoreJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRestoreJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRestoreJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobResponse) SetHeaders(v map[string]*string) *CreateRestoreJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRestoreJobResponse) SetStatusCode(v int32) *CreateRestoreJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRestoreJobResponse) SetBody(v *CreateRestoreJobResponseBody) *CreateRestoreJobResponse {
	s.Body = v
	return s
}

type CreateTempFileUploadUrlRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s CreateTempFileUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTempFileUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateTempFileUploadUrlRequest) SetFileName(v string) *CreateTempFileUploadUrlRequest {
	s.FileName = &v
	return s
}

type CreateTempFileUploadUrlResponseBody struct {
	BucketName     *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Endpoint       *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ExpireTime     *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OssAccessKeyId *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature      *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TempFileKey    *string `json:"TempFileKey,omitempty" xml:"TempFileKey,omitempty"`
}

func (s CreateTempFileUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTempFileUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTempFileUploadUrlResponseBody) SetBucketName(v string) *CreateTempFileUploadUrlResponseBody {
	s.BucketName = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetCode(v string) *CreateTempFileUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetEndpoint(v string) *CreateTempFileUploadUrlResponseBody {
	s.Endpoint = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetExpireTime(v int64) *CreateTempFileUploadUrlResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetMessage(v string) *CreateTempFileUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetOssAccessKeyId(v string) *CreateTempFileUploadUrlResponseBody {
	s.OssAccessKeyId = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetPolicy(v string) *CreateTempFileUploadUrlResponseBody {
	s.Policy = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetRequestId(v string) *CreateTempFileUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetSignature(v string) *CreateTempFileUploadUrlResponseBody {
	s.Signature = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetSuccess(v bool) *CreateTempFileUploadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTempFileUploadUrlResponseBody) SetTempFileKey(v string) *CreateTempFileUploadUrlResponseBody {
	s.TempFileKey = &v
	return s
}

type CreateTempFileUploadUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTempFileUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTempFileUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTempFileUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateTempFileUploadUrlResponse) SetHeaders(v map[string]*string) *CreateTempFileUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateTempFileUploadUrlResponse) SetStatusCode(v int32) *CreateTempFileUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTempFileUploadUrlResponse) SetBody(v *CreateTempFileUploadUrlResponseBody) *CreateTempFileUploadUrlResponse {
	s.Body = v
	return s
}

type CreateVaultRequest struct {
	BucketName           *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CompressionAlgorithm *string `json:"CompressionAlgorithm,omitempty" xml:"CompressionAlgorithm,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptType          *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	Endpoint             *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	KmsKeyId             *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	RedundancyType       *string `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	VaultAccessKeyId     *string `json:"VaultAccessKeyId,omitempty" xml:"VaultAccessKeyId,omitempty"`
	VaultAccessKeySecret *string `json:"VaultAccessKeySecret,omitempty" xml:"VaultAccessKeySecret,omitempty"`
	VaultName            *string `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
	VaultRegionId        *string `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
	VaultStorageClass    *string `json:"VaultStorageClass,omitempty" xml:"VaultStorageClass,omitempty"`
	VaultType            *string `json:"VaultType,omitempty" xml:"VaultType,omitempty"`
}

func (s CreateVaultRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVaultRequest) GoString() string {
	return s.String()
}

func (s *CreateVaultRequest) SetBucketName(v string) *CreateVaultRequest {
	s.BucketName = &v
	return s
}

func (s *CreateVaultRequest) SetCompressionAlgorithm(v string) *CreateVaultRequest {
	s.CompressionAlgorithm = &v
	return s
}

func (s *CreateVaultRequest) SetDescription(v string) *CreateVaultRequest {
	s.Description = &v
	return s
}

func (s *CreateVaultRequest) SetEncryptType(v string) *CreateVaultRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateVaultRequest) SetEndpoint(v string) *CreateVaultRequest {
	s.Endpoint = &v
	return s
}

func (s *CreateVaultRequest) SetKmsKeyId(v string) *CreateVaultRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateVaultRequest) SetRedundancyType(v string) *CreateVaultRequest {
	s.RedundancyType = &v
	return s
}

func (s *CreateVaultRequest) SetVaultAccessKeyId(v string) *CreateVaultRequest {
	s.VaultAccessKeyId = &v
	return s
}

func (s *CreateVaultRequest) SetVaultAccessKeySecret(v string) *CreateVaultRequest {
	s.VaultAccessKeySecret = &v
	return s
}

func (s *CreateVaultRequest) SetVaultName(v string) *CreateVaultRequest {
	s.VaultName = &v
	return s
}

func (s *CreateVaultRequest) SetVaultRegionId(v string) *CreateVaultRequest {
	s.VaultRegionId = &v
	return s
}

func (s *CreateVaultRequest) SetVaultStorageClass(v string) *CreateVaultRequest {
	s.VaultStorageClass = &v
	return s
}

func (s *CreateVaultRequest) SetVaultType(v string) *CreateVaultRequest {
	s.VaultType = &v
	return s
}

type CreateVaultResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VaultId   *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateVaultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVaultResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVaultResponseBody) SetCode(v string) *CreateVaultResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVaultResponseBody) SetMessage(v string) *CreateVaultResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVaultResponseBody) SetRequestId(v string) *CreateVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVaultResponseBody) SetSuccess(v bool) *CreateVaultResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVaultResponseBody) SetTaskId(v string) *CreateVaultResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVaultResponseBody) SetVaultId(v string) *CreateVaultResponseBody {
	s.VaultId = &v
	return s
}

type CreateVaultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVaultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVaultResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVaultResponse) GoString() string {
	return s.String()
}

func (s *CreateVaultResponse) SetHeaders(v map[string]*string) *CreateVaultResponse {
	s.Headers = v
	return s
}

func (s *CreateVaultResponse) SetStatusCode(v int32) *CreateVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVaultResponse) SetBody(v *CreateVaultResponseBody) *CreateVaultResponse {
	s.Body = v
	return s
}

type DeleteBackupClientRequest struct {
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s DeleteBackupClientRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupClientRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientRequest) SetClientId(v string) *DeleteBackupClientRequest {
	s.ClientId = &v
	return s
}

type DeleteBackupClientResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBackupClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupClientResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResponseBody) SetCode(v string) *DeleteBackupClientResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBackupClientResponseBody) SetMessage(v string) *DeleteBackupClientResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBackupClientResponseBody) SetRequestId(v string) *DeleteBackupClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupClientResponseBody) SetSuccess(v bool) *DeleteBackupClientResponseBody {
	s.Success = &v
	return s
}

type DeleteBackupClientResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBackupClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBackupClientResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupClientResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResponse) SetHeaders(v map[string]*string) *DeleteBackupClientResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupClientResponse) SetStatusCode(v int32) *DeleteBackupClientResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupClientResponse) SetBody(v *DeleteBackupClientResponseBody) *DeleteBackupClientResponse {
	s.Body = v
	return s
}

type DeleteBackupClientResourceRequest struct {
	ClientIds map[string]interface{} `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
}

func (s DeleteBackupClientResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupClientResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResourceRequest) SetClientIds(v map[string]interface{}) *DeleteBackupClientResourceRequest {
	s.ClientIds = v
	return s
}

type DeleteBackupClientResourceShrinkRequest struct {
	ClientIdsShrink *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
}

func (s DeleteBackupClientResourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupClientResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResourceShrinkRequest) SetClientIdsShrink(v string) *DeleteBackupClientResourceShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

type DeleteBackupClientResourceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBackupClientResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupClientResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResourceResponseBody) SetCode(v string) *DeleteBackupClientResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBackupClientResourceResponseBody) SetMessage(v string) *DeleteBackupClientResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBackupClientResourceResponseBody) SetRequestId(v string) *DeleteBackupClientResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupClientResourceResponseBody) SetSuccess(v bool) *DeleteBackupClientResourceResponseBody {
	s.Success = &v
	return s
}

type DeleteBackupClientResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBackupClientResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBackupClientResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupClientResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResourceResponse) SetHeaders(v map[string]*string) *DeleteBackupClientResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupClientResourceResponse) SetStatusCode(v int32) *DeleteBackupClientResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupClientResourceResponse) SetBody(v *DeleteBackupClientResourceResponseBody) *DeleteBackupClientResourceResponse {
	s.Body = v
	return s
}

type DeleteBackupPlanRequest struct {
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VaultId    *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanRequest) SetPlanId(v string) *DeleteBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteBackupPlanRequest) SetSourceType(v string) *DeleteBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *DeleteBackupPlanRequest) SetVaultId(v string) *DeleteBackupPlanRequest {
	s.VaultId = &v
	return s
}

type DeleteBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanResponseBody) SetCode(v string) *DeleteBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBackupPlanResponseBody) SetMessage(v string) *DeleteBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBackupPlanResponseBody) SetRequestId(v string) *DeleteBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupPlanResponseBody) SetSuccess(v bool) *DeleteBackupPlanResponseBody {
	s.Success = &v
	return s
}

type DeleteBackupPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanResponse) SetHeaders(v map[string]*string) *DeleteBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupPlanResponse) SetStatusCode(v int32) *DeleteBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupPlanResponse) SetBody(v *DeleteBackupPlanResponseBody) *DeleteBackupPlanResponse {
	s.Body = v
	return s
}

type DeleteHanaBackupPlanRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PlanId          *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteHanaBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteHanaBackupPlanRequest) SetClusterId(v string) *DeleteHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHanaBackupPlanRequest) SetPlanId(v string) *DeleteHanaBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteHanaBackupPlanRequest) SetResourceGroupId(v string) *DeleteHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteHanaBackupPlanRequest) SetVaultId(v string) *DeleteHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

type DeleteHanaBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHanaBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHanaBackupPlanResponseBody) SetCode(v string) *DeleteHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHanaBackupPlanResponseBody) SetMessage(v string) *DeleteHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHanaBackupPlanResponseBody) SetRequestId(v string) *DeleteHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHanaBackupPlanResponseBody) SetSuccess(v bool) *DeleteHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

type DeleteHanaBackupPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHanaBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteHanaBackupPlanResponse) SetHeaders(v map[string]*string) *DeleteHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteHanaBackupPlanResponse) SetStatusCode(v int32) *DeleteHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHanaBackupPlanResponse) SetBody(v *DeleteHanaBackupPlanResponseBody) *DeleteHanaBackupPlanResponse {
	s.Body = v
	return s
}

type DeleteHanaInstanceRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sid             *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteHanaInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHanaInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteHanaInstanceRequest) SetClusterId(v string) *DeleteHanaInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHanaInstanceRequest) SetResourceGroupId(v string) *DeleteHanaInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteHanaInstanceRequest) SetSid(v string) *DeleteHanaInstanceRequest {
	s.Sid = &v
	return s
}

func (s *DeleteHanaInstanceRequest) SetVaultId(v string) *DeleteHanaInstanceRequest {
	s.VaultId = &v
	return s
}

type DeleteHanaInstanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHanaInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHanaInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHanaInstanceResponseBody) SetCode(v string) *DeleteHanaInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHanaInstanceResponseBody) SetMessage(v string) *DeleteHanaInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHanaInstanceResponseBody) SetRequestId(v string) *DeleteHanaInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHanaInstanceResponseBody) SetSuccess(v bool) *DeleteHanaInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteHanaInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHanaInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHanaInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHanaInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteHanaInstanceResponse) SetHeaders(v map[string]*string) *DeleteHanaInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteHanaInstanceResponse) SetStatusCode(v int32) *DeleteHanaInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHanaInstanceResponse) SetBody(v *DeleteHanaInstanceResponseBody) *DeleteHanaInstanceResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Force      *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	VaultId    *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetClientId(v string) *DeleteSnapshotRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetForce(v bool) *DeleteSnapshotRequest {
	s.Force = &v
	return s
}

func (s *DeleteSnapshotRequest) SetInstanceId(v string) *DeleteSnapshotRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSourceType(v string) *DeleteSnapshotRequest {
	s.SourceType = &v
	return s
}

func (s *DeleteSnapshotRequest) SetToken(v string) *DeleteSnapshotRequest {
	s.Token = &v
	return s
}

func (s *DeleteSnapshotRequest) SetVaultId(v string) *DeleteSnapshotRequest {
	s.VaultId = &v
	return s
}

type DeleteSnapshotResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetCode(v string) *DeleteSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSnapshotResponseBody) SetMessage(v string) *DeleteSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotResponseBody) SetSuccess(v bool) *DeleteSnapshotResponseBody {
	s.Success = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetStatusCode(v int32) *DeleteSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DeleteVaultRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteVaultRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVaultRequest) GoString() string {
	return s.String()
}

func (s *DeleteVaultRequest) SetResourceGroupId(v string) *DeleteVaultRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteVaultRequest) SetToken(v string) *DeleteVaultRequest {
	s.Token = &v
	return s
}

func (s *DeleteVaultRequest) SetVaultId(v string) *DeleteVaultRequest {
	s.VaultId = &v
	return s
}

type DeleteVaultResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteVaultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVaultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVaultResponseBody) SetCode(v string) *DeleteVaultResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVaultResponseBody) SetMessage(v string) *DeleteVaultResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVaultResponseBody) SetRequestId(v string) *DeleteVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVaultResponseBody) SetSuccess(v bool) *DeleteVaultResponseBody {
	s.Success = &v
	return s
}

type DeleteVaultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVaultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVaultResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVaultResponse) GoString() string {
	return s.String()
}

func (s *DeleteVaultResponse) SetHeaders(v map[string]*string) *DeleteVaultResponse {
	s.Headers = v
	return s
}

func (s *DeleteVaultResponse) SetStatusCode(v int32) *DeleteVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVaultResponse) SetBody(v *DeleteVaultResponseBody) *DeleteVaultResponse {
	s.Body = v
	return s
}

type DescribeBackupClientsRequest struct {
	ClientIds            map[string]interface{}             `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	ClientType           *string                            `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	CrossAccountRoleName *string                            `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                            `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                             `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	InstanceIds          map[string]interface{}             `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PageNumber           *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tag                  []*DescribeBackupClientsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeBackupClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsRequest) SetClientIds(v map[string]interface{}) *DescribeBackupClientsRequest {
	s.ClientIds = v
	return s
}

func (s *DescribeBackupClientsRequest) SetClientType(v string) *DescribeBackupClientsRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetCrossAccountRoleName(v string) *DescribeBackupClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetCrossAccountType(v string) *DescribeBackupClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetCrossAccountUserId(v int64) *DescribeBackupClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetInstanceIds(v map[string]interface{}) *DescribeBackupClientsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeBackupClientsRequest) SetPageNumber(v int32) *DescribeBackupClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetPageSize(v int32) *DescribeBackupClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetTag(v []*DescribeBackupClientsRequestTag) *DescribeBackupClientsRequest {
	s.Tag = v
	return s
}

type DescribeBackupClientsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackupClientsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsRequestTag) SetKey(v string) *DescribeBackupClientsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeBackupClientsRequestTag) SetValue(v string) *DescribeBackupClientsRequestTag {
	s.Value = &v
	return s
}

type DescribeBackupClientsShrinkRequest struct {
	ClientIdsShrink      *string                                  `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	ClientType           *string                                  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	CrossAccountRoleName *string                                  `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                                  `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                                   `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	InstanceIdsShrink    *string                                  `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PageNumber           *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tag                  []*DescribeBackupClientsShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeBackupClientsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsShrinkRequest) SetClientIdsShrink(v string) *DescribeBackupClientsShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetClientType(v string) *DescribeBackupClientsShrinkRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetCrossAccountRoleName(v string) *DescribeBackupClientsShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetCrossAccountType(v string) *DescribeBackupClientsShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetCrossAccountUserId(v int64) *DescribeBackupClientsShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetInstanceIdsShrink(v string) *DescribeBackupClientsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetPageNumber(v int32) *DescribeBackupClientsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetPageSize(v int32) *DescribeBackupClientsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetTag(v []*DescribeBackupClientsShrinkRequestTag) *DescribeBackupClientsShrinkRequest {
	s.Tag = v
	return s
}

type DescribeBackupClientsShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackupClientsShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsShrinkRequestTag) SetKey(v string) *DescribeBackupClientsShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequestTag) SetValue(v string) *DescribeBackupClientsShrinkRequestTag {
	s.Value = &v
	return s
}

type DescribeBackupClientsResponseBody struct {
	Clients    []*DescribeBackupClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBody) SetClients(v []*DescribeBackupClientsResponseBodyClients) *DescribeBackupClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetCode(v string) *DescribeBackupClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetMessage(v string) *DescribeBackupClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetPageNumber(v int32) *DescribeBackupClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetPageSize(v int32) *DescribeBackupClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetRequestId(v string) *DescribeBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetSuccess(v bool) *DescribeBackupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupClientsResponseBody) SetTotalCount(v int64) *DescribeBackupClientsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBackupClientsResponseBodyClients struct {
	Appliance            *bool                                             `json:"Appliance,omitempty" xml:"Appliance,omitempty"`
	ArchType             *string                                           `json:"ArchType,omitempty" xml:"ArchType,omitempty"`
	BackupStatus         *string                                           `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	ClientId             *string                                           `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType           *string                                           `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientVersion        *string                                           `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CreatedTime          *int64                                            `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CrossAccountRoleName *string                                           `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                                           `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                                            `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	Hostname             *string                                           `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	InstanceId           *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string                                           `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LastHeartBeatTime    *int64                                            `json:"LastHeartBeatTime,omitempty" xml:"LastHeartBeatTime,omitempty"`
	MaxClientVersion     *string                                           `json:"MaxClientVersion,omitempty" xml:"MaxClientVersion,omitempty"`
	OsType               *string                                           `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PrivateIpV4          *string                                           `json:"PrivateIpV4,omitempty" xml:"PrivateIpV4,omitempty"`
	Settings             *DescribeBackupClientsResponseBodyClientsSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Struct"`
	Status               *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                 []*DescribeBackupClientsResponseBodyClientsTags   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UpdatedTime          *int64                                            `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	ZoneId               *string                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeBackupClientsResponseBodyClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBodyClients) SetAppliance(v bool) *DescribeBackupClientsResponseBodyClients {
	s.Appliance = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetArchType(v string) *DescribeBackupClientsResponseBodyClients {
	s.ArchType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetBackupStatus(v string) *DescribeBackupClientsResponseBodyClients {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientId(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientId = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientType(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetClientVersion(v string) *DescribeBackupClientsResponseBodyClients {
	s.ClientVersion = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetCreatedTime(v int64) *DescribeBackupClientsResponseBodyClients {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetCrossAccountRoleName(v string) *DescribeBackupClientsResponseBodyClients {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetCrossAccountType(v string) *DescribeBackupClientsResponseBodyClients {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetCrossAccountUserId(v int64) *DescribeBackupClientsResponseBodyClients {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetHostname(v string) *DescribeBackupClientsResponseBodyClients {
	s.Hostname = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetInstanceId(v string) *DescribeBackupClientsResponseBodyClients {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetInstanceName(v string) *DescribeBackupClientsResponseBodyClients {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetLastHeartBeatTime(v int64) *DescribeBackupClientsResponseBodyClients {
	s.LastHeartBeatTime = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetMaxClientVersion(v string) *DescribeBackupClientsResponseBodyClients {
	s.MaxClientVersion = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetOsType(v string) *DescribeBackupClientsResponseBodyClients {
	s.OsType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetPrivateIpV4(v string) *DescribeBackupClientsResponseBodyClients {
	s.PrivateIpV4 = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetSettings(v *DescribeBackupClientsResponseBodyClientsSettings) *DescribeBackupClientsResponseBodyClients {
	s.Settings = v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetStatus(v string) *DescribeBackupClientsResponseBodyClients {
	s.Status = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetTags(v []*DescribeBackupClientsResponseBodyClientsTags) *DescribeBackupClientsResponseBodyClients {
	s.Tags = v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetUpdatedTime(v int64) *DescribeBackupClientsResponseBodyClients {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClients) SetZoneId(v string) *DescribeBackupClientsResponseBodyClients {
	s.ZoneId = &v
	return s
}

type DescribeBackupClientsResponseBodyClientsSettings struct {
	DataNetworkType  *string `json:"DataNetworkType,omitempty" xml:"DataNetworkType,omitempty"`
	DataProxySetting *string `json:"DataProxySetting,omitempty" xml:"DataProxySetting,omitempty"`
	MaxCpuCore       *string `json:"MaxCpuCore,omitempty" xml:"MaxCpuCore,omitempty"`
	MaxWorker        *string `json:"MaxWorker,omitempty" xml:"MaxWorker,omitempty"`
	ProxyHost        *string `json:"ProxyHost,omitempty" xml:"ProxyHost,omitempty"`
	ProxyPassword    *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	ProxyPort        *int32  `json:"ProxyPort,omitempty" xml:"ProxyPort,omitempty"`
	ProxyUser        *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	UseHttps         *string `json:"UseHttps,omitempty" xml:"UseHttps,omitempty"`
}

func (s DescribeBackupClientsResponseBodyClientsSettings) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsResponseBodyClientsSettings) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetDataNetworkType(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.DataNetworkType = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetDataProxySetting(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.DataProxySetting = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetMaxCpuCore(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.MaxCpuCore = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetMaxWorker(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.MaxWorker = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetProxyHost(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.ProxyHost = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetProxyPassword(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.ProxyPassword = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetProxyPort(v int32) *DescribeBackupClientsResponseBodyClientsSettings {
	s.ProxyPort = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetProxyUser(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.ProxyUser = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsSettings) SetUseHttps(v string) *DescribeBackupClientsResponseBodyClientsSettings {
	s.UseHttps = &v
	return s
}

type DescribeBackupClientsResponseBodyClientsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackupClientsResponseBodyClientsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsResponseBodyClientsTags) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponseBodyClientsTags) SetKey(v string) *DescribeBackupClientsResponseBodyClientsTags {
	s.Key = &v
	return s
}

func (s *DescribeBackupClientsResponseBodyClientsTags) SetValue(v string) *DescribeBackupClientsResponseBodyClientsTags {
	s.Value = &v
	return s
}

type DescribeBackupClientsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponse) SetHeaders(v map[string]*string) *DescribeBackupClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupClientsResponse) SetStatusCode(v int32) *DescribeBackupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupClientsResponse) SetBody(v *DescribeBackupClientsResponseBody) *DescribeBackupClientsResponse {
	s.Body = v
	return s
}

type DescribeBackupJobs2Request struct {
	Filters       []*DescribeBackupJobs2RequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	PageNumber    *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortDirection *string                              `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	SourceType    *string                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeBackupJobs2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2Request) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2Request) SetFilters(v []*DescribeBackupJobs2RequestFilters) *DescribeBackupJobs2Request {
	s.Filters = v
	return s
}

func (s *DescribeBackupJobs2Request) SetPageNumber(v int32) *DescribeBackupJobs2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupJobs2Request) SetPageSize(v int32) *DescribeBackupJobs2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupJobs2Request) SetSortDirection(v string) *DescribeBackupJobs2Request {
	s.SortDirection = &v
	return s
}

func (s *DescribeBackupJobs2Request) SetSourceType(v string) *DescribeBackupJobs2Request {
	s.SourceType = &v
	return s
}

type DescribeBackupJobs2RequestFilters struct {
	Key      *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Values   []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2RequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2RequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2RequestFilters) SetKey(v string) *DescribeBackupJobs2RequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeBackupJobs2RequestFilters) SetOperator(v string) *DescribeBackupJobs2RequestFilters {
	s.Operator = &v
	return s
}

func (s *DescribeBackupJobs2RequestFilters) SetValues(v []*string) *DescribeBackupJobs2RequestFilters {
	s.Values = v
	return s
}

type DescribeBackupJobs2ResponseBody struct {
	BackupJobs *DescribeBackupJobs2ResponseBodyBackupJobs `json:"BackupJobs,omitempty" xml:"BackupJobs,omitempty" type:"Struct"`
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupJobs2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBody) SetBackupJobs(v *DescribeBackupJobs2ResponseBodyBackupJobs) *DescribeBackupJobs2ResponseBody {
	s.BackupJobs = v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetCode(v string) *DescribeBackupJobs2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetMessage(v string) *DescribeBackupJobs2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetPageNumber(v int32) *DescribeBackupJobs2ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetPageSize(v int32) *DescribeBackupJobs2ResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetRequestId(v string) *DescribeBackupJobs2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetSuccess(v bool) *DescribeBackupJobs2ResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetTotalCount(v int64) *DescribeBackupJobs2ResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBackupJobs2ResponseBodyBackupJobs struct {
	BackupJob []*DescribeBackupJobs2ResponseBodyBackupJobsBackupJob `json:"BackupJob,omitempty" xml:"BackupJob,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobs) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobs) SetBackupJob(v []*DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) *DescribeBackupJobs2ResponseBodyBackupJobs {
	s.BackupJob = v
	return s
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJob struct {
	ActualBytes          *int64                                                       `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	ActualItems          *int64                                                       `json:"ActualItems,omitempty" xml:"ActualItems,omitempty"`
	BackupType           *string                                                      `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	Bucket               *string                                                      `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	BytesDone            *int64                                                       `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	BytesTotal           *int64                                                       `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	ClientId             *string                                                      `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CompleteTime         *int64                                                       `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreateTime           *int64                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedTime          *int64                                                       `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CrossAccountRoleName *string                                                      `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                                                      `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                                                       `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	ErrorMessage         *string                                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Exclude              *string                                                      `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	FileSystemId         *string                                                      `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Include              *string                                                      `json:"Include,omitempty" xml:"Include,omitempty"`
	InstanceId           *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string                                                      `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ItemsDone            *int64                                                       `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	ItemsTotal           *int64                                                       `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	JobId                *string                                                      `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName              *string                                                      `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Options              *string                                                      `json:"Options,omitempty" xml:"Options,omitempty"`
	OtsDetail            *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty" type:"Struct"`
	Paths                *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths     `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Struct"`
	PlanId               *string                                                      `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	Prefix               *string                                                      `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Progress             *int32                                                       `json:"Progress,omitempty" xml:"Progress,omitempty"`
	SourceType           *string                                                      `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Speed                *int64                                                       `json:"Speed,omitempty" xml:"Speed,omitempty"`
	SpeedLimit           *string                                                      `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	StartTime            *int64                                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName            *string                                                      `json:"TableName,omitempty" xml:"TableName,omitempty"`
	UpdatedTime          *int64                                                       `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	VaultId              *string                                                      `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetActualBytes(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ActualBytes = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetActualItems(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ActualItems = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetBackupType(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetBucket(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Bucket = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetBytesDone(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.BytesDone = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetBytesTotal(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.BytesTotal = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetClientId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ClientId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCompleteTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CompleteTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCreateTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCreatedTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCrossAccountRoleName(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCrossAccountType(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCrossAccountUserId(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetErrorMessage(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetExclude(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Exclude = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetFileSystemId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.FileSystemId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetInclude(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Include = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetInstanceId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetInstanceName(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetItemsDone(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ItemsDone = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetItemsTotal(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ItemsTotal = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetJobId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.JobId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetJobName(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.JobName = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetOptions(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Options = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetOtsDetail(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.OtsDetail = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetPaths(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Paths = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetPlanId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.PlanId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetPrefix(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Prefix = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetProgress(v int32) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Progress = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetSourceType(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.SourceType = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetSpeed(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Speed = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetSpeedLimit(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.SpeedLimit = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetStartTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetStatus(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Status = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetTableName(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.TableName = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetUpdatedTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetVaultId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.VaultId = &v
	return s
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail struct {
	TableNames *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Struct"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) SetTableNames(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail {
	s.TableNames = v
	return s
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames struct {
	TableName []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) SetTableName(v []*string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames {
	s.TableName = v
	return s
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths struct {
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) SetPath(v []*string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths {
	s.Path = v
	return s
}

type DescribeBackupJobs2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupJobs2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupJobs2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupJobs2Response) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2Response) SetHeaders(v map[string]*string) *DescribeBackupJobs2Response {
	s.Headers = v
	return s
}

func (s *DescribeBackupJobs2Response) SetStatusCode(v int32) *DescribeBackupJobs2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupJobs2Response) SetBody(v *DescribeBackupJobs2ResponseBody) *DescribeBackupJobs2Response {
	s.Body = v
	return s
}

type DescribeBackupPlansRequest struct {
	Filters    []*DescribeBackupPlansRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceType *string                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeBackupPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansRequest) SetFilters(v []*DescribeBackupPlansRequestFilters) *DescribeBackupPlansRequest {
	s.Filters = v
	return s
}

func (s *DescribeBackupPlansRequest) SetPageNumber(v int32) *DescribeBackupPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupPlansRequest) SetPageSize(v int32) *DescribeBackupPlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPlansRequest) SetSourceType(v string) *DescribeBackupPlansRequest {
	s.SourceType = &v
	return s
}

type DescribeBackupPlansRequestFilters struct {
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansRequestFilters) SetKey(v string) *DescribeBackupPlansRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeBackupPlansRequestFilters) SetValues(v []*string) *DescribeBackupPlansRequestFilters {
	s.Values = v
	return s
}

type DescribeBackupPlansResponseBody struct {
	BackupPlans *DescribeBackupPlansResponseBodyBackupPlans `json:"BackupPlans,omitempty" xml:"BackupPlans,omitempty" type:"Struct"`
	Code        *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBody) SetBackupPlans(v *DescribeBackupPlansResponseBodyBackupPlans) *DescribeBackupPlansResponseBody {
	s.BackupPlans = v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetCode(v string) *DescribeBackupPlansResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetMessage(v string) *DescribeBackupPlansResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetPageNumber(v int32) *DescribeBackupPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetPageSize(v int32) *DescribeBackupPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetRequestId(v string) *DescribeBackupPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetSuccess(v bool) *DescribeBackupPlansResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetTotalCount(v int64) *DescribeBackupPlansResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlans struct {
	BackupPlan []*DescribeBackupPlansResponseBodyBackupPlansBackupPlan `json:"BackupPlan,omitempty" xml:"BackupPlan,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlans) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlans) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlans) SetBackupPlan(v []*DescribeBackupPlansResponseBodyBackupPlansBackupPlan) *DescribeBackupPlansResponseBodyBackupPlans {
	s.BackupPlan = v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlan struct {
	BackupSourceGroupId  *string                                                        `json:"BackupSourceGroupId,omitempty" xml:"BackupSourceGroupId,omitempty"`
	BackupType           *string                                                        `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	Bucket               *string                                                        `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	ClientId             *string                                                        `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClusterId            *string                                                        `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime           *int64                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedTime          *int64                                                         `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CrossAccountRoleName *string                                                        `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                                                        `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                                                         `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	DataSourceId         *string                                                        `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	Detail               *string                                                        `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Disabled             *bool                                                          `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	Exclude              *string                                                        `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	FileSystemId         *string                                                        `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Include              *string                                                        `json:"Include,omitempty" xml:"Include,omitempty"`
	InstanceGroupId      *string                                                        `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	InstanceId           *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string                                                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Options              *string                                                        `json:"Options,omitempty" xml:"Options,omitempty"`
	OtsDetail            *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty" type:"Struct"`
	Paths                *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths     `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Struct"`
	PlanId               *string                                                        `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName             *string                                                        `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	Prefix               *string                                                        `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Resources            *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Retention            *int64                                                         `json:"Retention,omitempty" xml:"Retention,omitempty"`
	Rules                *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules     `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	Schedule             *string                                                        `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	SourceType           *string                                                        `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SpeedLimit           *string                                                        `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	TrialInfo            *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo `json:"TrialInfo,omitempty" xml:"TrialInfo,omitempty" type:"Struct"`
	UpdatedTime          *int64                                                         `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	VaultId              *string                                                        `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlan) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBackupSourceGroupId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.BackupSourceGroupId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBackupType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBucket(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Bucket = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetClientId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.ClientId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetClusterId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCreateTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CreateTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCreatedTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCrossAccountRoleName(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCrossAccountType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCrossAccountUserId(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDataSourceId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.DataSourceId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDetail(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Detail = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDisabled(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Disabled = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetExclude(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Exclude = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetFileSystemId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.FileSystemId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetInclude(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Include = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetInstanceGroupId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetInstanceId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetInstanceName(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetOptions(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Options = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetOtsDetail(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.OtsDetail = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetPaths(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Paths = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetPlanId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.PlanId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetPlanName(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.PlanName = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetPrefix(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Prefix = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetResources(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Resources = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetRetention(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Retention = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetRules(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Rules = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetSchedule(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Schedule = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetSourceType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.SourceType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetSpeedLimit(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.SpeedLimit = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetTrialInfo(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.TrialInfo = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetUpdatedTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetVaultId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.VaultId = &v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail struct {
	TableNames *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Struct"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) SetTableNames(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail {
	s.TableNames = v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames struct {
	TableName []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) SetTableName(v []*string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames {
	s.TableName = v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths struct {
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) SetPath(v []*string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths {
	s.Path = v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources struct {
	Resource []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) SetResource(v []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources {
	s.Resource = v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource struct {
	Extra      *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) SetExtra(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource {
	s.Extra = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) SetResourceId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) SetSourceType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource {
	s.SourceType = &v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules struct {
	Rule []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) SetRule(v []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules {
	s.Rule = v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule struct {
	BackupType           *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DestinationRegionId  *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationRetention *int64  `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	Disabled             *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DoCopy               *bool   `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	Retention            *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RuleId               *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Schedule             *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetBackupType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetDestinationRegionId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.DestinationRegionId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetDestinationRetention(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.DestinationRetention = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetDisabled(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.Disabled = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetDoCopy(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.DoCopy = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetRetention(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.Retention = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetRuleId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetRuleName(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetSchedule(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.Schedule = &v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo struct {
	KeepAfterTrialExpiration *bool  `json:"KeepAfterTrialExpiration,omitempty" xml:"KeepAfterTrialExpiration,omitempty"`
	TrialExpireTime          *int64 `json:"TrialExpireTime,omitempty" xml:"TrialExpireTime,omitempty"`
	TrialStartTime           *int64 `json:"TrialStartTime,omitempty" xml:"TrialStartTime,omitempty"`
	TrialVaultReleaseTime    *int64 `json:"TrialVaultReleaseTime,omitempty" xml:"TrialVaultReleaseTime,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) SetKeepAfterTrialExpiration(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	s.KeepAfterTrialExpiration = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) SetTrialExpireTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	s.TrialExpireTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) SetTrialStartTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	s.TrialStartTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) SetTrialVaultReleaseTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	s.TrialVaultReleaseTime = &v
	return s
}

type DescribeBackupPlansResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponse) SetHeaders(v map[string]*string) *DescribeBackupPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlansResponse) SetStatusCode(v int32) *DescribeBackupPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPlansResponse) SetBody(v *DescribeBackupPlansResponseBody) *DescribeBackupPlansResponse {
	s.Body = v
	return s
}

type DescribeClientsRequest struct {
	ClientId        *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType      *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientsRequest) SetClientId(v string) *DescribeClientsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeClientsRequest) SetClientType(v string) *DescribeClientsRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeClientsRequest) SetClusterId(v string) *DescribeClientsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClientsRequest) SetPageNumber(v int32) *DescribeClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClientsRequest) SetPageSize(v int32) *DescribeClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClientsRequest) SetResourceGroupId(v string) *DescribeClientsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClientsRequest) SetSourceType(v string) *DescribeClientsRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeClientsRequest) SetVaultId(v string) *DescribeClientsRequest {
	s.VaultId = &v
	return s
}

type DescribeClientsResponseBody struct {
	Clients    *DescribeClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Struct"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBody) SetClients(v *DescribeClientsResponseBodyClients) *DescribeClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeClientsResponseBody) SetCode(v string) *DescribeClientsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClientsResponseBody) SetMessage(v string) *DescribeClientsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeClientsResponseBody) SetPageNumber(v int32) *DescribeClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClientsResponseBody) SetPageSize(v int32) *DescribeClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeClientsResponseBody) SetRequestId(v string) *DescribeClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientsResponseBody) SetSuccess(v bool) *DescribeClientsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeClientsResponseBody) SetTotalCount(v int32) *DescribeClientsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeClientsResponseBodyClients struct {
	Client []*DescribeClientsResponseBodyClientsClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Repeated"`
}

func (s DescribeClientsResponseBodyClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBodyClients) SetClient(v []*DescribeClientsResponseBodyClientsClient) *DescribeClientsResponseBodyClients {
	s.Client = v
	return s
}

type DescribeClientsResponseBodyClientsClient struct {
	AlertSetting  *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientName    *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	ClientType    *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreatedTime   *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MaxVersion    *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	NetworkType   *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	UpdatedTime   *int64  `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	UseHttps      *bool   `json:"UseHttps,omitempty" xml:"UseHttps,omitempty"`
	VaultId       *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeClientsResponseBodyClientsClient) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientsResponseBodyClientsClient) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponseBodyClientsClient) SetAlertSetting(v string) *DescribeClientsResponseBodyClientsClient {
	s.AlertSetting = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClientId(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClientId = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClientName(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClientName = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClientType(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClientType = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClientVersion(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetClusterId(v string) *DescribeClientsResponseBodyClientsClient {
	s.ClusterId = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetCreatedTime(v int64) *DescribeClientsResponseBodyClientsClient {
	s.CreatedTime = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetInstanceId(v string) *DescribeClientsResponseBodyClientsClient {
	s.InstanceId = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetInstanceName(v string) *DescribeClientsResponseBodyClientsClient {
	s.InstanceName = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetMaxVersion(v string) *DescribeClientsResponseBodyClientsClient {
	s.MaxVersion = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetNetworkType(v string) *DescribeClientsResponseBodyClientsClient {
	s.NetworkType = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetStatus(v string) *DescribeClientsResponseBodyClientsClient {
	s.Status = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetStatusMessage(v string) *DescribeClientsResponseBodyClientsClient {
	s.StatusMessage = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetUpdatedTime(v int64) *DescribeClientsResponseBodyClientsClient {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetUseHttps(v bool) *DescribeClientsResponseBodyClientsClient {
	s.UseHttps = &v
	return s
}

func (s *DescribeClientsResponseBodyClientsClient) SetVaultId(v string) *DescribeClientsResponseBodyClientsClient {
	s.VaultId = &v
	return s
}

type DescribeClientsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponse) SetHeaders(v map[string]*string) *DescribeClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientsResponse) SetStatusCode(v int32) *DescribeClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientsResponse) SetBody(v *DescribeClientsResponseBody) *DescribeClientsResponse {
	s.Body = v
	return s
}

type DescribeContainerClusterRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeContainerClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerClusterRequest) SetClusterId(v string) *DescribeContainerClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerClusterRequest) SetPageNumber(v int32) *DescribeContainerClusterRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeContainerClusterRequest) SetPageSize(v int32) *DescribeContainerClusterRequest {
	s.PageSize = &v
	return s
}

type DescribeContainerClusterResponseBody struct {
	Clusters   []*DescribeContainerClusterResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	Code       *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerClusterResponseBody) SetClusters(v []*DescribeContainerClusterResponseBodyClusters) *DescribeContainerClusterResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetCode(v string) *DescribeContainerClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetMessage(v string) *DescribeContainerClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetPageNumber(v int32) *DescribeContainerClusterResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetPageSize(v int32) *DescribeContainerClusterResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetRequestId(v string) *DescribeContainerClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetSuccess(v bool) *DescribeContainerClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeContainerClusterResponseBody) SetTotalCount(v int64) *DescribeContainerClusterResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeContainerClusterResponseBodyClusters struct {
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Identifier  *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeContainerClusterResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerClusterResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeContainerClusterResponseBodyClusters) SetAgentStatus(v string) *DescribeContainerClusterResponseBodyClusters {
	s.AgentStatus = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetClusterId(v string) *DescribeContainerClusterResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetClusterType(v string) *DescribeContainerClusterResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetDescription(v string) *DescribeContainerClusterResponseBodyClusters {
	s.Description = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetIdentifier(v string) *DescribeContainerClusterResponseBodyClusters {
	s.Identifier = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetName(v string) *DescribeContainerClusterResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetNetworkType(v string) *DescribeContainerClusterResponseBodyClusters {
	s.NetworkType = &v
	return s
}

func (s *DescribeContainerClusterResponseBodyClusters) SetToken(v string) *DescribeContainerClusterResponseBodyClusters {
	s.Token = &v
	return s
}

type DescribeContainerClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerClusterResponse) SetHeaders(v map[string]*string) *DescribeContainerClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerClusterResponse) SetStatusCode(v int32) *DescribeContainerClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerClusterResponse) SetBody(v *DescribeContainerClusterResponseBody) *DescribeContainerClusterResponse {
	s.Body = v
	return s
}

type DescribeHanaBackupPlansRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaBackupPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansRequest) SetClusterId(v string) *DescribeHanaBackupPlansRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetDatabaseName(v string) *DescribeHanaBackupPlansRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetPageNumber(v int32) *DescribeHanaBackupPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetPageSize(v int32) *DescribeHanaBackupPlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetResourceGroupId(v string) *DescribeHanaBackupPlansRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaBackupPlansRequest) SetVaultId(v string) *DescribeHanaBackupPlansRequest {
	s.VaultId = &v
	return s
}

type DescribeHanaBackupPlansResponseBody struct {
	Code            *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	HanaBackupPlans *DescribeHanaBackupPlansResponseBodyHanaBackupPlans `json:"HanaBackupPlans,omitempty" xml:"HanaBackupPlans,omitempty" type:"Struct"`
	Message         *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber      *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount      *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaBackupPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBody) SetCode(v string) *DescribeHanaBackupPlansResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetHanaBackupPlans(v *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) *DescribeHanaBackupPlansResponseBody {
	s.HanaBackupPlans = v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetMessage(v string) *DescribeHanaBackupPlansResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetPageNumber(v int32) *DescribeHanaBackupPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetPageSize(v int32) *DescribeHanaBackupPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetRequestId(v string) *DescribeHanaBackupPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetSuccess(v bool) *DescribeHanaBackupPlansResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetTotalCount(v int64) *DescribeHanaBackupPlansResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeHanaBackupPlansResponseBodyHanaBackupPlans struct {
	HanaBackupPlan []*DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan `json:"HanaBackupPlan,omitempty" xml:"HanaBackupPlan,omitempty" type:"Repeated"`
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlans) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlans) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) SetHanaBackupPlan(v []*DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) *DescribeHanaBackupPlansResponseBodyHanaBackupPlans {
	s.HanaBackupPlan = v
	return s
}

type DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan struct {
	BackupPrefix *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	BackupType   *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Disabled     *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	PlanId       *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName     *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	Schedule     *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	VaultId      *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetBackupPrefix(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.BackupPrefix = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetBackupType(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.BackupType = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetClusterId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetDatabaseName(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetDisabled(v bool) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.Disabled = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetPlanId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.PlanId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetPlanName(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.PlanName = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetSchedule(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.Schedule = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetVaultId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.VaultId = &v
	return s
}

type DescribeHanaBackupPlansResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHanaBackupPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHanaBackupPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponse) SetHeaders(v map[string]*string) *DescribeHanaBackupPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaBackupPlansResponse) SetStatusCode(v int32) *DescribeHanaBackupPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaBackupPlansResponse) SetBody(v *DescribeHanaBackupPlansResponseBody) *DescribeHanaBackupPlansResponse {
	s.Body = v
	return s
}

type DescribeHanaBackupSettingRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	VaultId      *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaBackupSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupSettingRequest) SetClusterId(v string) *DescribeHanaBackupSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupSettingRequest) SetDatabaseName(v string) *DescribeHanaBackupSettingRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupSettingRequest) SetVaultId(v string) *DescribeHanaBackupSettingRequest {
	s.VaultId = &v
	return s
}

type DescribeHanaBackupSettingResponseBody struct {
	Code              *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	HanaBackupSetting *DescribeHanaBackupSettingResponseBodyHanaBackupSetting `json:"HanaBackupSetting,omitempty" xml:"HanaBackupSetting,omitempty" type:"Struct"`
	Message           *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHanaBackupSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupSettingResponseBody) SetCode(v string) *DescribeHanaBackupSettingResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) SetHanaBackupSetting(v *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) *DescribeHanaBackupSettingResponseBody {
	s.HanaBackupSetting = v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) SetMessage(v string) *DescribeHanaBackupSettingResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) SetRequestId(v string) *DescribeHanaBackupSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBody) SetSuccess(v bool) *DescribeHanaBackupSettingResponseBody {
	s.Success = &v
	return s
}

type DescribeHanaBackupSettingResponseBodyHanaBackupSetting struct {
	CatalogBackupParameterFile *string `json:"CatalogBackupParameterFile,omitempty" xml:"CatalogBackupParameterFile,omitempty"`
	CatalogBackupUsingBackint  *bool   `json:"CatalogBackupUsingBackint,omitempty" xml:"CatalogBackupUsingBackint,omitempty"`
	DataBackupParameterFile    *string `json:"DataBackupParameterFile,omitempty" xml:"DataBackupParameterFile,omitempty"`
	DatabaseName               *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EnableAutoLogBackup        *bool   `json:"EnableAutoLogBackup,omitempty" xml:"EnableAutoLogBackup,omitempty"`
	LogBackupParameterFile     *string `json:"LogBackupParameterFile,omitempty" xml:"LogBackupParameterFile,omitempty"`
	LogBackupTimeout           *int64  `json:"LogBackupTimeout,omitempty" xml:"LogBackupTimeout,omitempty"`
	LogBackupUsingBackint      *bool   `json:"LogBackupUsingBackint,omitempty" xml:"LogBackupUsingBackint,omitempty"`
}

func (s DescribeHanaBackupSettingResponseBodyHanaBackupSetting) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupSettingResponseBodyHanaBackupSetting) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetCatalogBackupParameterFile(v string) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.CatalogBackupParameterFile = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetCatalogBackupUsingBackint(v bool) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.CatalogBackupUsingBackint = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetDataBackupParameterFile(v string) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.DataBackupParameterFile = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetDatabaseName(v string) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetEnableAutoLogBackup(v bool) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.EnableAutoLogBackup = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetLogBackupParameterFile(v string) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.LogBackupParameterFile = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetLogBackupTimeout(v int64) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.LogBackupTimeout = &v
	return s
}

func (s *DescribeHanaBackupSettingResponseBodyHanaBackupSetting) SetLogBackupUsingBackint(v bool) *DescribeHanaBackupSettingResponseBodyHanaBackupSetting {
	s.LogBackupUsingBackint = &v
	return s
}

type DescribeHanaBackupSettingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHanaBackupSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHanaBackupSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupSettingResponse) SetHeaders(v map[string]*string) *DescribeHanaBackupSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaBackupSettingResponse) SetStatusCode(v int32) *DescribeHanaBackupSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaBackupSettingResponse) SetBody(v *DescribeHanaBackupSettingResponseBody) *DescribeHanaBackupSettingResponse {
	s.Body = v
	return s
}

type DescribeHanaBackupsAsyncRequest struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName        *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IncludeDifferential *bool   `json:"IncludeDifferential,omitempty" xml:"IncludeDifferential,omitempty"`
	IncludeIncremental  *bool   `json:"IncludeIncremental,omitempty" xml:"IncludeIncremental,omitempty"`
	IncludeLog          *bool   `json:"IncludeLog,omitempty" xml:"IncludeLog,omitempty"`
	LogPosition         *int64  `json:"LogPosition,omitempty" xml:"LogPosition,omitempty"`
	Mode                *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RecoveryPointInTime *int64  `json:"RecoveryPointInTime,omitempty" xml:"RecoveryPointInTime,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceClusterId     *string `json:"SourceClusterId,omitempty" xml:"SourceClusterId,omitempty"`
	SystemCopy          *bool   `json:"SystemCopy,omitempty" xml:"SystemCopy,omitempty"`
	UseBackint          *bool   `json:"UseBackint,omitempty" xml:"UseBackint,omitempty"`
	VaultId             *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	VolumeId            *int32  `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
}

func (s DescribeHanaBackupsAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupsAsyncRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupsAsyncRequest) SetClusterId(v string) *DescribeHanaBackupsAsyncRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetDatabaseName(v string) *DescribeHanaBackupsAsyncRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetIncludeDifferential(v bool) *DescribeHanaBackupsAsyncRequest {
	s.IncludeDifferential = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetIncludeIncremental(v bool) *DescribeHanaBackupsAsyncRequest {
	s.IncludeIncremental = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetIncludeLog(v bool) *DescribeHanaBackupsAsyncRequest {
	s.IncludeLog = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetLogPosition(v int64) *DescribeHanaBackupsAsyncRequest {
	s.LogPosition = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetMode(v string) *DescribeHanaBackupsAsyncRequest {
	s.Mode = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetPageNumber(v int32) *DescribeHanaBackupsAsyncRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetPageSize(v int32) *DescribeHanaBackupsAsyncRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetRecoveryPointInTime(v int64) *DescribeHanaBackupsAsyncRequest {
	s.RecoveryPointInTime = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetResourceGroupId(v string) *DescribeHanaBackupsAsyncRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetSource(v string) *DescribeHanaBackupsAsyncRequest {
	s.Source = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetSourceClusterId(v string) *DescribeHanaBackupsAsyncRequest {
	s.SourceClusterId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetSystemCopy(v bool) *DescribeHanaBackupsAsyncRequest {
	s.SystemCopy = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetUseBackint(v bool) *DescribeHanaBackupsAsyncRequest {
	s.UseBackint = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetVaultId(v string) *DescribeHanaBackupsAsyncRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetVolumeId(v int32) *DescribeHanaBackupsAsyncRequest {
	s.VolumeId = &v
	return s
}

type DescribeHanaBackupsAsyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeHanaBackupsAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupsAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetCode(v string) *DescribeHanaBackupsAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetMessage(v string) *DescribeHanaBackupsAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetRequestId(v string) *DescribeHanaBackupsAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetSuccess(v bool) *DescribeHanaBackupsAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetTaskId(v string) *DescribeHanaBackupsAsyncResponseBody {
	s.TaskId = &v
	return s
}

type DescribeHanaBackupsAsyncResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHanaBackupsAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHanaBackupsAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaBackupsAsyncResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupsAsyncResponse) SetHeaders(v map[string]*string) *DescribeHanaBackupsAsyncResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaBackupsAsyncResponse) SetStatusCode(v int32) *DescribeHanaBackupsAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponse) SetBody(v *DescribeHanaBackupsAsyncResponseBody) *DescribeHanaBackupsAsyncResponse {
	s.Body = v
	return s
}

type DescribeHanaDatabasesRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesRequest) SetClusterId(v string) *DescribeHanaDatabasesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) SetPageNumber(v int32) *DescribeHanaDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) SetPageSize(v int32) *DescribeHanaDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) SetResourceGroupId(v string) *DescribeHanaDatabasesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaDatabasesRequest) SetVaultId(v string) *DescribeHanaDatabasesRequest {
	s.VaultId = &v
	return s
}

type DescribeHanaDatabasesResponseBody struct {
	Code          *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	HanaDatabases *DescribeHanaDatabasesResponseBodyHanaDatabases `json:"HanaDatabases,omitempty" xml:"HanaDatabases,omitempty" type:"Struct"`
	Message       *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount    *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesResponseBody) SetCode(v string) *DescribeHanaDatabasesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetHanaDatabases(v *DescribeHanaDatabasesResponseBodyHanaDatabases) *DescribeHanaDatabasesResponseBody {
	s.HanaDatabases = v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetMessage(v string) *DescribeHanaDatabasesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetPageNumber(v int32) *DescribeHanaDatabasesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetPageSize(v int32) *DescribeHanaDatabasesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetRequestId(v string) *DescribeHanaDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetSuccess(v bool) *DescribeHanaDatabasesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetTotalCount(v int64) *DescribeHanaDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeHanaDatabasesResponseBodyHanaDatabases struct {
	HanaDatabase []*DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase `json:"HanaDatabase,omitempty" xml:"HanaDatabase,omitempty" type:"Repeated"`
}

func (s DescribeHanaDatabasesResponseBodyHanaDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaDatabasesResponseBodyHanaDatabases) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabases) SetHanaDatabase(v []*DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) *DescribeHanaDatabasesResponseBodyHanaDatabases {
	s.HanaDatabase = v
	return s
}

type DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase struct {
	ActiveStatus *string `json:"ActiveStatus,omitempty" xml:"ActiveStatus,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Detail       *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SqlPort      *int32  `json:"SqlPort,omitempty" xml:"SqlPort,omitempty"`
}

func (s DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetActiveStatus(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.ActiveStatus = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetDatabaseName(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetDetail(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.Detail = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetHost(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.Host = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetServiceName(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.ServiceName = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetSqlPort(v int32) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.SqlPort = &v
	return s
}

type DescribeHanaDatabasesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHanaDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHanaDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesResponse) SetHeaders(v map[string]*string) *DescribeHanaDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaDatabasesResponse) SetStatusCode(v int32) *DescribeHanaDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaDatabasesResponse) SetBody(v *DescribeHanaDatabasesResponseBody) *DescribeHanaDatabasesResponse {
	s.Body = v
	return s
}

type DescribeHanaInstancesRequest struct {
	ClusterId       *string                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber      *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeHanaInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VaultId         *string                            `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesRequest) SetClusterId(v string) *DescribeHanaInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaInstancesRequest) SetPageNumber(v int32) *DescribeHanaInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaInstancesRequest) SetPageSize(v int32) *DescribeHanaInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaInstancesRequest) SetResourceGroupId(v string) *DescribeHanaInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaInstancesRequest) SetTag(v []*DescribeHanaInstancesRequestTag) *DescribeHanaInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeHanaInstancesRequest) SetVaultId(v string) *DescribeHanaInstancesRequest {
	s.VaultId = &v
	return s
}

type DescribeHanaInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHanaInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesRequestTag) SetKey(v string) *DescribeHanaInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeHanaInstancesRequestTag) SetValue(v string) *DescribeHanaInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeHanaInstancesResponseBody struct {
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Hanas      *DescribeHanaInstancesResponseBodyHanas `json:"Hanas,omitempty" xml:"Hanas,omitempty" type:"Struct"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBody) SetCode(v string) *DescribeHanaInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetHanas(v *DescribeHanaInstancesResponseBodyHanas) *DescribeHanaInstancesResponseBody {
	s.Hanas = v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetMessage(v string) *DescribeHanaInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetPageNumber(v int32) *DescribeHanaInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetPageSize(v int32) *DescribeHanaInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetRequestId(v string) *DescribeHanaInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetSuccess(v bool) *DescribeHanaInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaInstancesResponseBody) SetTotalCount(v int32) *DescribeHanaInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeHanaInstancesResponseBodyHanas struct {
	Hana []*DescribeHanaInstancesResponseBodyHanasHana `json:"Hana,omitempty" xml:"Hana,omitempty" type:"Repeated"`
}

func (s DescribeHanaInstancesResponseBodyHanas) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaInstancesResponseBodyHanas) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBodyHanas) SetHana(v []*DescribeHanaInstancesResponseBodyHanasHana) *DescribeHanaInstancesResponseBodyHanas {
	s.Hana = v
	return s
}

type DescribeHanaInstancesResponseBodyHanasHana struct {
	AlertSetting        *string                                         `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	ClusterId           *string                                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HanaName            *string                                         `json:"HanaName,omitempty" xml:"HanaName,omitempty"`
	Host                *string                                         `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceNumber      *int32                                          `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	ResourceGroupId     *string                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status              *int64                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage       *string                                         `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Tags                *DescribeHanaInstancesResponseBodyHanasHanaTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UseSsl              *bool                                           `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	UserName            *string                                         `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ValidateCertificate *bool                                           `json:"ValidateCertificate,omitempty" xml:"ValidateCertificate,omitempty"`
	VaultId             *string                                         `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaInstancesResponseBodyHanasHana) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaInstancesResponseBodyHanasHana) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetAlertSetting(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.AlertSetting = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetClusterId(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetHanaName(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.HanaName = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetHost(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.Host = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetInstanceNumber(v int32) *DescribeHanaInstancesResponseBodyHanasHana {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetResourceGroupId(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetStatus(v int64) *DescribeHanaInstancesResponseBodyHanasHana {
	s.Status = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetStatusMessage(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.StatusMessage = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetTags(v *DescribeHanaInstancesResponseBodyHanasHanaTags) *DescribeHanaInstancesResponseBodyHanasHana {
	s.Tags = v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetUseSsl(v bool) *DescribeHanaInstancesResponseBodyHanasHana {
	s.UseSsl = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetUserName(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.UserName = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetValidateCertificate(v bool) *DescribeHanaInstancesResponseBodyHanasHana {
	s.ValidateCertificate = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHana) SetVaultId(v string) *DescribeHanaInstancesResponseBodyHanasHana {
	s.VaultId = &v
	return s
}

type DescribeHanaInstancesResponseBodyHanasHanaTags struct {
	Tag []*DescribeHanaInstancesResponseBodyHanasHanaTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeHanaInstancesResponseBodyHanasHanaTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaInstancesResponseBodyHanasHanaTags) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTags) SetTag(v []*DescribeHanaInstancesResponseBodyHanasHanaTagsTag) *DescribeHanaInstancesResponseBodyHanasHanaTags {
	s.Tag = v
	return s
}

type DescribeHanaInstancesResponseBodyHanasHanaTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHanaInstancesResponseBodyHanasHanaTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaInstancesResponseBodyHanasHanaTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTagsTag) SetKey(v string) *DescribeHanaInstancesResponseBodyHanasHanaTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeHanaInstancesResponseBodyHanasHanaTagsTag) SetValue(v string) *DescribeHanaInstancesResponseBodyHanasHanaTagsTag {
	s.Value = &v
	return s
}

type DescribeHanaInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHanaInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHanaInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesResponse) SetHeaders(v map[string]*string) *DescribeHanaInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaInstancesResponse) SetStatusCode(v int32) *DescribeHanaInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaInstancesResponse) SetBody(v *DescribeHanaInstancesResponseBody) *DescribeHanaInstancesResponse {
	s.Body = v
	return s
}

type DescribeHanaRestoresRequest struct {
	BackupId        *int64  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RestoreId       *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	RestoreStatus   *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaRestoresRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaRestoresRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresRequest) SetBackupId(v int64) *DescribeHanaRestoresRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetClusterId(v string) *DescribeHanaRestoresRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetDatabaseName(v string) *DescribeHanaRestoresRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetPageNumber(v int32) *DescribeHanaRestoresRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetPageSize(v int32) *DescribeHanaRestoresRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetResourceGroupId(v string) *DescribeHanaRestoresRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetRestoreId(v string) *DescribeHanaRestoresRequest {
	s.RestoreId = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetRestoreStatus(v string) *DescribeHanaRestoresRequest {
	s.RestoreStatus = &v
	return s
}

func (s *DescribeHanaRestoresRequest) SetVaultId(v string) *DescribeHanaRestoresRequest {
	s.VaultId = &v
	return s
}

type DescribeHanaRestoresResponseBody struct {
	Code        *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HanaRestore *DescribeHanaRestoresResponseBodyHanaRestore `json:"HanaRestore,omitempty" xml:"HanaRestore,omitempty" type:"Struct"`
	Message     *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaRestoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaRestoresResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresResponseBody) SetCode(v string) *DescribeHanaRestoresResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetHanaRestore(v *DescribeHanaRestoresResponseBodyHanaRestore) *DescribeHanaRestoresResponseBody {
	s.HanaRestore = v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetMessage(v string) *DescribeHanaRestoresResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetPageNumber(v int32) *DescribeHanaRestoresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetPageSize(v int32) *DescribeHanaRestoresResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetRequestId(v string) *DescribeHanaRestoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetSuccess(v bool) *DescribeHanaRestoresResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetTotalCount(v int32) *DescribeHanaRestoresResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeHanaRestoresResponseBodyHanaRestore struct {
	HanaRestores []*DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores `json:"HanaRestores,omitempty" xml:"HanaRestores,omitempty" type:"Repeated"`
}

func (s DescribeHanaRestoresResponseBodyHanaRestore) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaRestoresResponseBodyHanaRestore) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresResponseBodyHanaRestore) SetHanaRestores(v []*DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) *DescribeHanaRestoresResponseBodyHanaRestore {
	s.HanaRestores = v
	return s
}

type DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores struct {
	BackupID            *int64  `json:"BackupID,omitempty" xml:"BackupID,omitempty"`
	BackupPrefix        *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	CheckAccess         *bool   `json:"CheckAccess,omitempty" xml:"CheckAccess,omitempty"`
	ClearLog            *bool   `json:"ClearLog,omitempty" xml:"ClearLog,omitempty"`
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CurrentPhase        *int32  `json:"CurrentPhase,omitempty" xml:"CurrentPhase,omitempty"`
	CurrentProgress     *int64  `json:"CurrentProgress,omitempty" xml:"CurrentProgress,omitempty"`
	DatabaseName        *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DatabaseRestoreId   *int64  `json:"DatabaseRestoreId,omitempty" xml:"DatabaseRestoreId,omitempty"`
	EndTime             *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogPosition         *int64  `json:"LogPosition,omitempty" xml:"LogPosition,omitempty"`
	MaxPhase            *int32  `json:"MaxPhase,omitempty" xml:"MaxPhase,omitempty"`
	MaxProgress         *int64  `json:"MaxProgress,omitempty" xml:"MaxProgress,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Mode                *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Phase               *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	ReachedTime         *int64  `json:"ReachedTime,omitempty" xml:"ReachedTime,omitempty"`
	RecoveryPointInTime *int64  `json:"RecoveryPointInTime,omitempty" xml:"RecoveryPointInTime,omitempty"`
	RestoreId           *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceClusterId     *string `json:"SourceClusterId,omitempty" xml:"SourceClusterId,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemCopy          *bool   `json:"SystemCopy,omitempty" xml:"SystemCopy,omitempty"`
	UseCatalog          *bool   `json:"UseCatalog,omitempty" xml:"UseCatalog,omitempty"`
	UseDelta            *bool   `json:"UseDelta,omitempty" xml:"UseDelta,omitempty"`
	VaultId             *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	VolumeId            *int32  `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
}

func (s DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetBackupID(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.BackupID = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetBackupPrefix(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.BackupPrefix = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetCheckAccess(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.CheckAccess = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetClearLog(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.ClearLog = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetClusterId(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetCurrentPhase(v int32) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.CurrentPhase = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetCurrentProgress(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.CurrentProgress = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetDatabaseName(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetDatabaseRestoreId(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.DatabaseRestoreId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetEndTime(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.EndTime = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetLogPosition(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.LogPosition = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetMaxPhase(v int32) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.MaxPhase = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetMaxProgress(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.MaxProgress = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetMessage(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Message = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetMode(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Mode = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetPhase(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Phase = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetReachedTime(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.ReachedTime = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetRecoveryPointInTime(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.RecoveryPointInTime = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetRestoreId(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.RestoreId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetSource(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Source = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetSourceClusterId(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.SourceClusterId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetStartTime(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.StartTime = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetState(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.State = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetStatus(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Status = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetSystemCopy(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.SystemCopy = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetUseCatalog(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.UseCatalog = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetUseDelta(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.UseDelta = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetVaultId(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetVolumeId(v int32) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.VolumeId = &v
	return s
}

type DescribeHanaRestoresResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHanaRestoresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHanaRestoresResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaRestoresResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresResponse) SetHeaders(v map[string]*string) *DescribeHanaRestoresResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaRestoresResponse) SetStatusCode(v int32) *DescribeHanaRestoresResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaRestoresResponse) SetBody(v *DescribeHanaRestoresResponseBody) *DescribeHanaRestoresResponse {
	s.Body = v
	return s
}

type DescribeHanaRetentionSettingRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	VaultId      *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaRetentionSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaRetentionSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaRetentionSettingRequest) SetClusterId(v string) *DescribeHanaRetentionSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaRetentionSettingRequest) SetDatabaseName(v string) *DescribeHanaRetentionSettingRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaRetentionSettingRequest) SetVaultId(v string) *DescribeHanaRetentionSettingRequest {
	s.VaultId = &v
	return s
}

type DescribeHanaRetentionSettingResponseBody struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DatabaseName  *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Disabled      *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetentionDays *int64  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	Schedule      *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	VaultId       *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaRetentionSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaRetentionSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaRetentionSettingResponseBody) SetClusterId(v string) *DescribeHanaRetentionSettingResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetCode(v string) *DescribeHanaRetentionSettingResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetDatabaseName(v string) *DescribeHanaRetentionSettingResponseBody {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetDisabled(v bool) *DescribeHanaRetentionSettingResponseBody {
	s.Disabled = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetMessage(v string) *DescribeHanaRetentionSettingResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetRequestId(v string) *DescribeHanaRetentionSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetRetentionDays(v int64) *DescribeHanaRetentionSettingResponseBody {
	s.RetentionDays = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetSchedule(v string) *DescribeHanaRetentionSettingResponseBody {
	s.Schedule = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetSuccess(v bool) *DescribeHanaRetentionSettingResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponseBody) SetVaultId(v string) *DescribeHanaRetentionSettingResponseBody {
	s.VaultId = &v
	return s
}

type DescribeHanaRetentionSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHanaRetentionSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHanaRetentionSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHanaRetentionSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaRetentionSettingResponse) SetHeaders(v map[string]*string) *DescribeHanaRetentionSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaRetentionSettingResponse) SetStatusCode(v int32) *DescribeHanaRetentionSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaRetentionSettingResponse) SetBody(v *DescribeHanaRetentionSettingResponseBody) *DescribeHanaRetentionSettingResponse {
	s.Body = v
	return s
}

type DescribeOtsTableSnapshotsRequest struct {
	CrossAccountRoleName *string                                         `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                                         `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                                          `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	EndTime              *int64                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit                *int32                                          `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken            *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OtsInstances         []*DescribeOtsTableSnapshotsRequestOtsInstances `json:"OtsInstances,omitempty" xml:"OtsInstances,omitempty" type:"Repeated"`
	SnapshotIds          []*string                                       `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty" type:"Repeated"`
	StartTime            *int64                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOtsTableSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOtsTableSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsRequest) SetCrossAccountRoleName(v string) *DescribeOtsTableSnapshotsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetCrossAccountType(v string) *DescribeOtsTableSnapshotsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetCrossAccountUserId(v int64) *DescribeOtsTableSnapshotsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetEndTime(v int64) *DescribeOtsTableSnapshotsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetLimit(v int32) *DescribeOtsTableSnapshotsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetNextToken(v string) *DescribeOtsTableSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetOtsInstances(v []*DescribeOtsTableSnapshotsRequestOtsInstances) *DescribeOtsTableSnapshotsRequest {
	s.OtsInstances = v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetSnapshotIds(v []*string) *DescribeOtsTableSnapshotsRequest {
	s.SnapshotIds = v
	return s
}

func (s *DescribeOtsTableSnapshotsRequest) SetStartTime(v int64) *DescribeOtsTableSnapshotsRequest {
	s.StartTime = &v
	return s
}

type DescribeOtsTableSnapshotsRequestOtsInstances struct {
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	TableNames   []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s DescribeOtsTableSnapshotsRequestOtsInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeOtsTableSnapshotsRequestOtsInstances) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsRequestOtsInstances) SetInstanceName(v string) *DescribeOtsTableSnapshotsRequestOtsInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeOtsTableSnapshotsRequestOtsInstances) SetTableNames(v []*string) *DescribeOtsTableSnapshotsRequestOtsInstances {
	s.TableNames = v
	return s
}

type DescribeOtsTableSnapshotsResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Limit     *int32                                            `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots []*DescribeOtsTableSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeOtsTableSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOtsTableSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetCode(v string) *DescribeOtsTableSnapshotsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetLimit(v int32) *DescribeOtsTableSnapshotsResponseBody {
	s.Limit = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetMessage(v string) *DescribeOtsTableSnapshotsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetNextToken(v string) *DescribeOtsTableSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetRequestId(v string) *DescribeOtsTableSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetSnapshots(v []*DescribeOtsTableSnapshotsResponseBodySnapshots) *DescribeOtsTableSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBody) SetSuccess(v bool) *DescribeOtsTableSnapshotsResponseBody {
	s.Success = &v
	return s
}

type DescribeOtsTableSnapshotsResponseBodySnapshots struct {
	ActualBytes        *string `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	BackupType         *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BytesTotal         *int64  `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	CompleteTime       *int64  `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedTime        *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	JobId              *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ParentSnapshotHash *string `json:"ParentSnapshotHash,omitempty" xml:"ParentSnapshotHash,omitempty"`
	RangeEnd           *int64  `json:"RangeEnd,omitempty" xml:"RangeEnd,omitempty"`
	RangeStart         *int64  `json:"RangeStart,omitempty" xml:"RangeStart,omitempty"`
	Retention          *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	SnapshotHash       *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	SnapshotId         *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SourceType         *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartTime          *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName          *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	UpdatedTime        *int64  `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	VaultId            *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeOtsTableSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeOtsTableSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetActualBytes(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.ActualBytes = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetBackupType(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.BackupType = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetBytesTotal(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.BytesTotal = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetCompleteTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.CompleteTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetCreateTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.CreateTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetCreatedTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.CreatedTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetInstanceName(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.InstanceName = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetJobId(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.JobId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetParentSnapshotHash(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.ParentSnapshotHash = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetRangeEnd(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.RangeEnd = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetRangeStart(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.RangeStart = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetRetention(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.Retention = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetSnapshotHash(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetSourceType(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.SourceType = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetStartTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.StartTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetTableName(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.TableName = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetUpdatedTime(v int64) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponseBodySnapshots) SetVaultId(v string) *DescribeOtsTableSnapshotsResponseBodySnapshots {
	s.VaultId = &v
	return s
}

type DescribeOtsTableSnapshotsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOtsTableSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOtsTableSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOtsTableSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOtsTableSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeOtsTableSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOtsTableSnapshotsResponse) SetStatusCode(v int32) *DescribeOtsTableSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOtsTableSnapshotsResponse) SetBody(v *DescribeOtsTableSnapshotsResponseBody) *DescribeOtsTableSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeRecoverableOtsInstancesRequest struct {
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s DescribeRecoverableOtsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableOtsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableOtsInstancesRequest) SetCrossAccountRoleName(v string) *DescribeRecoverableOtsInstancesRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesRequest) SetCrossAccountType(v string) *DescribeRecoverableOtsInstancesRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesRequest) SetCrossAccountUserId(v int64) *DescribeRecoverableOtsInstancesRequest {
	s.CrossAccountUserId = &v
	return s
}

type DescribeRecoverableOtsInstancesResponseBody struct {
	Code         *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	OtsInstances []*DescribeRecoverableOtsInstancesResponseBodyOtsInstances `json:"OtsInstances,omitempty" xml:"OtsInstances,omitempty" type:"Repeated"`
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRecoverableOtsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableOtsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetCode(v string) *DescribeRecoverableOtsInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetMessage(v string) *DescribeRecoverableOtsInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetOtsInstances(v []*DescribeRecoverableOtsInstancesResponseBodyOtsInstances) *DescribeRecoverableOtsInstancesResponseBody {
	s.OtsInstances = v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetRequestId(v string) *DescribeRecoverableOtsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBody) SetSuccess(v bool) *DescribeRecoverableOtsInstancesResponseBody {
	s.Success = &v
	return s
}

type DescribeRecoverableOtsInstancesResponseBodyOtsInstances struct {
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	TableNames   []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s DescribeRecoverableOtsInstancesResponseBodyOtsInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableOtsInstancesResponseBodyOtsInstances) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableOtsInstancesResponseBodyOtsInstances) SetInstanceName(v string) *DescribeRecoverableOtsInstancesResponseBodyOtsInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponseBodyOtsInstances) SetTableNames(v []*string) *DescribeRecoverableOtsInstancesResponseBodyOtsInstances {
	s.TableNames = v
	return s
}

type DescribeRecoverableOtsInstancesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRecoverableOtsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecoverableOtsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableOtsInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableOtsInstancesResponse) SetHeaders(v map[string]*string) *DescribeRecoverableOtsInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponse) SetStatusCode(v int32) *DescribeRecoverableOtsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponse) SetBody(v *DescribeRecoverableOtsInstancesResponseBody) *DescribeRecoverableOtsInstancesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	NeedVaultCount *bool `json:"NeedVaultCount,omitempty" xml:"NeedVaultCount,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetNeedVaultCount(v bool) *DescribeRegionsRequest {
	s.NeedVaultCount = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName  *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VaultCount *int32  `json:"VaultCount,omitempty" xml:"VaultCount,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetVaultCount(v int32) *DescribeRegionsResponseBodyRegionsRegion {
	s.VaultCount = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRestoreJobs2Request struct {
	Filters     []*DescribeRestoreJobs2RequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	PageNumber  *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreType *string                               `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
}

func (s DescribeRestoreJobs2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreJobs2Request) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2Request) SetFilters(v []*DescribeRestoreJobs2RequestFilters) *DescribeRestoreJobs2Request {
	s.Filters = v
	return s
}

func (s *DescribeRestoreJobs2Request) SetPageNumber(v int32) *DescribeRestoreJobs2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreJobs2Request) SetPageSize(v int32) *DescribeRestoreJobs2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreJobs2Request) SetRestoreType(v string) *DescribeRestoreJobs2Request {
	s.RestoreType = &v
	return s
}

type DescribeRestoreJobs2RequestFilters struct {
	Key      *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Values   []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeRestoreJobs2RequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreJobs2RequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2RequestFilters) SetKey(v string) *DescribeRestoreJobs2RequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeRestoreJobs2RequestFilters) SetOperator(v string) *DescribeRestoreJobs2RequestFilters {
	s.Operator = &v
	return s
}

func (s *DescribeRestoreJobs2RequestFilters) SetValues(v []*string) *DescribeRestoreJobs2RequestFilters {
	s.Values = v
	return s
}

type DescribeRestoreJobs2ResponseBody struct {
	Code        *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreJobs *DescribeRestoreJobs2ResponseBodyRestoreJobs `json:"RestoreJobs,omitempty" xml:"RestoreJobs,omitempty" type:"Struct"`
	Success     *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRestoreJobs2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBody) SetCode(v string) *DescribeRestoreJobs2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetMessage(v string) *DescribeRestoreJobs2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetPageNumber(v int32) *DescribeRestoreJobs2ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetPageSize(v int32) *DescribeRestoreJobs2ResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetRequestId(v string) *DescribeRestoreJobs2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetRestoreJobs(v *DescribeRestoreJobs2ResponseBodyRestoreJobs) *DescribeRestoreJobs2ResponseBody {
	s.RestoreJobs = v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetSuccess(v bool) *DescribeRestoreJobs2ResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetTotalCount(v int32) *DescribeRestoreJobs2ResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRestoreJobs2ResponseBodyRestoreJobs struct {
	RestoreJob []*DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob `json:"RestoreJob,omitempty" xml:"RestoreJob,omitempty" type:"Repeated"`
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobs) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobs) SetRestoreJob(v []*DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) *DescribeRestoreJobs2ResponseBodyRestoreJobs {
	s.RestoreJob = v
	return s
}

type DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob struct {
	ActualBytes          *int64                                                          `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	ActualItems          *int64                                                          `json:"ActualItems,omitempty" xml:"ActualItems,omitempty"`
	BytesDone            *int64                                                          `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	BytesTotal           *int64                                                          `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	ClusterId            *string                                                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CompleteTime         *int64                                                          `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreatedTime          *int64                                                          `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CrossAccountRoleName *string                                                         `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                                                         `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                                                          `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	ErrorFile            *string                                                         `json:"ErrorFile,omitempty" xml:"ErrorFile,omitempty"`
	ErrorMessage         *string                                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Exclude              *string                                                         `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	ExpireTime           *int64                                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Include              *string                                                         `json:"Include,omitempty" xml:"Include,omitempty"`
	ItemsDone            *int64                                                          `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	ItemsTotal           *int64                                                          `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	Options              *string                                                         `json:"Options,omitempty" xml:"Options,omitempty"`
	OtsDetail            *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty" type:"Struct"`
	ParentId             *string                                                         `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Progress             *int32                                                          `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Report               *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport    `json:"Report,omitempty" xml:"Report,omitempty" type:"Struct"`
	RestoreId            *string                                                         `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	RestoreType          *string                                                         `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SnapshotHash         *string                                                         `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	SnapshotId           *string                                                         `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SourceType           *string                                                         `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Speed                *int64                                                          `json:"Speed,omitempty" xml:"Speed,omitempty"`
	StartTime            *int64                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetBucket         *string                                                         `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	TargetClientId       *string                                                         `json:"TargetClientId,omitempty" xml:"TargetClientId,omitempty"`
	TargetCreateTime     *int64                                                          `json:"TargetCreateTime,omitempty" xml:"TargetCreateTime,omitempty"`
	TargetDataSourceId   *string                                                         `json:"TargetDataSourceId,omitempty" xml:"TargetDataSourceId,omitempty"`
	TargetFileSystemId   *string                                                         `json:"TargetFileSystemId,omitempty" xml:"TargetFileSystemId,omitempty"`
	TargetInstanceId     *string                                                         `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	TargetInstanceName   *string                                                         `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	TargetPath           *string                                                         `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetPrefix         *string                                                         `json:"TargetPrefix,omitempty" xml:"TargetPrefix,omitempty"`
	TargetTableName      *string                                                         `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	TargetTime           *int64                                                          `json:"TargetTime,omitempty" xml:"TargetTime,omitempty"`
	UdmDetail            *string                                                         `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty"`
	UpdatedTime          *int64                                                          `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	VaultId              *string                                                         `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetActualBytes(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ActualBytes = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetActualItems(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ActualItems = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetBytesDone(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.BytesDone = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetBytesTotal(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.BytesTotal = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetClusterId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCompleteTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CompleteTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCreatedTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CreatedTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCrossAccountRoleName(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCrossAccountType(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCrossAccountUserId(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetErrorFile(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ErrorFile = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetErrorMessage(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetExclude(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Exclude = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetExpireTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetInclude(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Include = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetItemsDone(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ItemsDone = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetItemsTotal(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ItemsTotal = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetOptions(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Options = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetOtsDetail(v *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.OtsDetail = v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetParentId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ParentId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetProgress(v int32) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Progress = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetReport(v *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Report = v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetRestoreId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.RestoreId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetRestoreType(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.RestoreType = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSnapshotHash(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSnapshotId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.SnapshotId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSourceType(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.SourceType = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSpeed(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetStartTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetStatus(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Status = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetBucket(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetBucket = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetClientId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetClientId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetCreateTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetCreateTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetDataSourceId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetDataSourceId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetFileSystemId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetFileSystemId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetInstanceId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetInstanceId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetInstanceName(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetInstanceName = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetPath(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetPath = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetPrefix(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetPrefix = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetTableName(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetTableName = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetUdmDetail(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.UdmDetail = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetUpdatedTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetVaultId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.VaultId = &v
	return s
}

type DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail struct {
	BatchChannelCount *int32 `json:"BatchChannelCount,omitempty" xml:"BatchChannelCount,omitempty"`
	OverwriteExisting *bool  `json:"OverwriteExisting,omitempty" xml:"OverwriteExisting,omitempty"`
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) SetBatchChannelCount(v int32) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail {
	s.BatchChannelCount = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) SetOverwriteExisting(v bool) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail {
	s.OverwriteExisting = &v
	return s
}

type DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport struct {
	TotalFiles *string `json:"TotalFiles,omitempty" xml:"TotalFiles,omitempty"`
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) SetTotalFiles(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport {
	s.TotalFiles = &v
	return s
}

type DescribeRestoreJobs2Response struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRestoreJobs2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreJobs2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreJobs2Response) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2Response) SetHeaders(v map[string]*string) *DescribeRestoreJobs2Response {
	s.Headers = v
	return s
}

func (s *DescribeRestoreJobs2Response) SetStatusCode(v int32) *DescribeRestoreJobs2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreJobs2Response) SetBody(v *DescribeRestoreJobs2ResponseBody) *DescribeRestoreJobs2Response {
	s.Body = v
	return s
}

type DescribeTaskRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskRequest) SetResourceGroupId(v string) *DescribeTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTaskRequest) SetTaskId(v string) *DescribeTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskRequest) SetToken(v string) *DescribeTaskRequest {
	s.Token = &v
	return s
}

type DescribeTaskResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CompletedTime *int64  `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	CreatedTime   *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress      *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result        *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	UpdatedTime   *int64  `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) SetCode(v string) *DescribeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTaskResponseBody) SetCompletedTime(v int64) *DescribeTaskResponseBody {
	s.CompletedTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetCreatedTime(v int64) *DescribeTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetDescription(v string) *DescribeTaskResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeTaskResponseBody) SetMessage(v string) *DescribeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBody) SetName(v string) *DescribeTaskResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeTaskResponseBody) SetProgress(v int32) *DescribeTaskResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetResult(v string) *DescribeTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeTaskResponseBody) SetSuccess(v bool) *DescribeTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTaskResponseBody) SetUpdatedTime(v int64) *DescribeTaskResponseBody {
	s.UpdatedTime = &v
	return s
}

type DescribeTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponse) SetHeaders(v map[string]*string) *DescribeTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskResponse) SetStatusCode(v int32) *DescribeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskResponse) SetBody(v *DescribeTaskResponseBody) *DescribeTaskResponse {
	s.Body = v
	return s
}

type DescribeUdmSnapshotsRequest struct {
	DiskId      *string                `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	EndTime     *int64                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId  *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId       *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SnapshotIds map[string]interface{} `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	SourceType  *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartTime   *int64                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UdmRegionId *string                `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
}

func (s DescribeUdmSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUdmSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsRequest) SetDiskId(v string) *DescribeUdmSnapshotsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetEndTime(v int64) *DescribeUdmSnapshotsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetInstanceId(v string) *DescribeUdmSnapshotsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetJobId(v string) *DescribeUdmSnapshotsRequest {
	s.JobId = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetSnapshotIds(v map[string]interface{}) *DescribeUdmSnapshotsRequest {
	s.SnapshotIds = v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetSourceType(v string) *DescribeUdmSnapshotsRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetStartTime(v int64) *DescribeUdmSnapshotsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetUdmRegionId(v string) *DescribeUdmSnapshotsRequest {
	s.UdmRegionId = &v
	return s
}

type DescribeUdmSnapshotsShrinkRequest struct {
	DiskId            *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	EndTime           *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SnapshotIdsShrink *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	SourceType        *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartTime         *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UdmRegionId       *string `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
}

func (s DescribeUdmSnapshotsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUdmSnapshotsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetDiskId(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetEndTime(v int64) *DescribeUdmSnapshotsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetInstanceId(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetJobId(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.JobId = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetSnapshotIdsShrink(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.SnapshotIdsShrink = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetSourceType(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetStartTime(v int64) *DescribeUdmSnapshotsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetUdmRegionId(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.UdmRegionId = &v
	return s
}

type DescribeUdmSnapshotsResponseBody struct {
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots  []*DescribeUdmSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUdmSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUdmSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsResponseBody) SetCode(v string) *DescribeUdmSnapshotsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetMessage(v string) *DescribeUdmSnapshotsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetRequestId(v string) *DescribeUdmSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetSnapshots(v []*DescribeUdmSnapshotsResponseBodySnapshots) *DescribeUdmSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetSuccess(v bool) *DescribeUdmSnapshotsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetTotalCount(v int64) *DescribeUdmSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeUdmSnapshotsResponseBodySnapshots struct {
	ActualBytes        *string                                          `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	BackupType         *string                                          `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BytesTotal         *int64                                           `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	CompleteTime       *int64                                           `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreateTime         *int64                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedTime        *int64                                           `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Detail             *DescribeUdmSnapshotsResponseBodySnapshotsDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	DiskId             *string                                          `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	InstanceId         *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId              *string                                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	NativeSnapshotId   *string                                          `json:"NativeSnapshotId,omitempty" xml:"NativeSnapshotId,omitempty"`
	NativeSnapshotInfo *string                                          `json:"NativeSnapshotInfo,omitempty" xml:"NativeSnapshotInfo,omitempty"`
	ParentSnapshotHash *string                                          `json:"ParentSnapshotHash,omitempty" xml:"ParentSnapshotHash,omitempty"`
	Prefix             *string                                          `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	RealSnapshotTime   *int64                                           `json:"RealSnapshotTime,omitempty" xml:"RealSnapshotTime,omitempty"`
	Retention          *int64                                           `json:"Retention,omitempty" xml:"Retention,omitempty"`
	SnapshotHash       *string                                          `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	SnapshotId         *string                                          `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SourceType         *string                                          `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartTime          *int64                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status             *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdatedTime        *int64                                           `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeUdmSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeUdmSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetActualBytes(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.ActualBytes = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetBackupType(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.BackupType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetBytesTotal(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.BytesTotal = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetCompleteTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.CompleteTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetCreateTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.CreateTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetCreatedTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.CreatedTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetDetail(v *DescribeUdmSnapshotsResponseBodySnapshotsDetail) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.Detail = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetDiskId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.DiskId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetInstanceId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetJobId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.JobId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetNativeSnapshotId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.NativeSnapshotId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetNativeSnapshotInfo(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.NativeSnapshotInfo = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetParentSnapshotHash(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.ParentSnapshotHash = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetPrefix(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.Prefix = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetRealSnapshotTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.RealSnapshotTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetRetention(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.Retention = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetSnapshotHash(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetSourceType(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.SourceType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetStartTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.StartTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetUpdatedTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.UpdatedTime = &v
	return s
}

type DescribeUdmSnapshotsResponseBodySnapshotsDetail struct {
	ConsistentLevel                *string                `json:"ConsistentLevel,omitempty" xml:"ConsistentLevel,omitempty"`
	ContainOsDisk                  *bool                  `json:"ContainOsDisk,omitempty" xml:"ContainOsDisk,omitempty"`
	DiskCategory                   *string                `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	DiskDevName                    *string                `json:"DiskDevName,omitempty" xml:"DiskDevName,omitempty"`
	DiskHbrSnapshotIdWithDeviceMap map[string]interface{} `json:"DiskHbrSnapshotIdWithDeviceMap,omitempty" xml:"DiskHbrSnapshotIdWithDeviceMap,omitempty"`
	DiskIdList                     []*string              `json:"DiskIdList,omitempty" xml:"DiskIdList,omitempty" type:"Repeated"`
	DowngradeReason                *string                `json:"DowngradeReason,omitempty" xml:"DowngradeReason,omitempty"`
	InstanceIdWithDiskIdListMap    map[string]interface{} `json:"InstanceIdWithDiskIdListMap,omitempty" xml:"InstanceIdWithDiskIdListMap,omitempty"`
	InstanceName                   *string                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType                   *string                `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NativeSnapshotIdList           []*string              `json:"NativeSnapshotIdList,omitempty" xml:"NativeSnapshotIdList,omitempty" type:"Repeated"`
	OsDiskId                       *string                `json:"OsDiskId,omitempty" xml:"OsDiskId,omitempty"`
	OsName                         *string                `json:"OsName,omitempty" xml:"OsName,omitempty"`
	OsNameEn                       *string                `json:"OsNameEn,omitempty" xml:"OsNameEn,omitempty"`
	OsType                         *string                `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PerformanceLevel               *string                `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Platform                       *string                `json:"Platform,omitempty" xml:"Platform,omitempty"`
	SnapshotGroupId                *string                `json:"SnapshotGroupId,omitempty" xml:"SnapshotGroupId,omitempty"`
	SystemDisk                     *bool                  `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	VmName                         *string                `json:"VmName,omitempty" xml:"VmName,omitempty"`
}

func (s DescribeUdmSnapshotsResponseBodySnapshotsDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeUdmSnapshotsResponseBodySnapshotsDetail) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetConsistentLevel(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.ConsistentLevel = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetContainOsDisk(v bool) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.ContainOsDisk = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDiskCategory(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DiskCategory = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDiskDevName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DiskDevName = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDiskHbrSnapshotIdWithDeviceMap(v map[string]interface{}) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DiskHbrSnapshotIdWithDeviceMap = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDiskIdList(v []*string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DiskIdList = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDowngradeReason(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DowngradeReason = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetInstanceIdWithDiskIdListMap(v map[string]interface{}) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.InstanceIdWithDiskIdListMap = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetInstanceName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.InstanceName = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetInstanceType(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.InstanceType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetNativeSnapshotIdList(v []*string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.NativeSnapshotIdList = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetOsDiskId(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.OsDiskId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetOsName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.OsName = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetOsNameEn(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.OsNameEn = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetOsType(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.OsType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetPerformanceLevel(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetPlatform(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.Platform = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetSnapshotGroupId(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.SnapshotGroupId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetSystemDisk(v bool) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.SystemDisk = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetVmName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.VmName = &v
	return s
}

type DescribeUdmSnapshotsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUdmSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUdmSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUdmSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeUdmSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUdmSnapshotsResponse) SetStatusCode(v int32) *DescribeUdmSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUdmSnapshotsResponse) SetBody(v *DescribeUdmSnapshotsResponseBody) *DescribeUdmSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeVaultReplicationRegionsRequest struct {
	Token   *string `json:"Token,omitempty" xml:"Token,omitempty"`
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeVaultReplicationRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultReplicationRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVaultReplicationRegionsRequest) SetToken(v string) *DescribeVaultReplicationRegionsRequest {
	s.Token = &v
	return s
}

func (s *DescribeVaultReplicationRegionsRequest) SetVaultId(v string) *DescribeVaultReplicationRegionsRequest {
	s.VaultId = &v
	return s
}

type DescribeVaultReplicationRegionsResponseBody struct {
	Code      *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions   *DescribeVaultReplicationRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeVaultReplicationRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultReplicationRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetCode(v string) *DescribeVaultReplicationRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetMessage(v string) *DescribeVaultReplicationRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetRegions(v *DescribeVaultReplicationRegionsResponseBodyRegions) *DescribeVaultReplicationRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetRequestId(v string) *DescribeVaultReplicationRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetSuccess(v bool) *DescribeVaultReplicationRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeVaultReplicationRegionsResponseBodyRegions struct {
	RegionId []*string `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeVaultReplicationRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultReplicationRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeVaultReplicationRegionsResponseBodyRegions) SetRegionId(v []*string) *DescribeVaultReplicationRegionsResponseBodyRegions {
	s.RegionId = v
	return s
}

type DescribeVaultReplicationRegionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVaultReplicationRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVaultReplicationRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultReplicationRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVaultReplicationRegionsResponse) SetHeaders(v map[string]*string) *DescribeVaultReplicationRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVaultReplicationRegionsResponse) SetStatusCode(v int32) *DescribeVaultReplicationRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponse) SetBody(v *DescribeVaultReplicationRegionsResponseBody) *DescribeVaultReplicationRegionsResponse {
	s.Body = v
	return s
}

type DescribeVaultsRequest struct {
	PageNumber      *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag             []*DescribeVaultsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VaultId         *string                     `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	VaultRegionId   *string                     `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
	VaultType       *string                     `json:"VaultType,omitempty" xml:"VaultType,omitempty"`
}

func (s DescribeVaultsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVaultsRequest) SetPageNumber(v int32) *DescribeVaultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVaultsRequest) SetPageSize(v int32) *DescribeVaultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVaultsRequest) SetResourceGroupId(v string) *DescribeVaultsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVaultsRequest) SetStatus(v string) *DescribeVaultsRequest {
	s.Status = &v
	return s
}

func (s *DescribeVaultsRequest) SetTag(v []*DescribeVaultsRequestTag) *DescribeVaultsRequest {
	s.Tag = v
	return s
}

func (s *DescribeVaultsRequest) SetVaultId(v string) *DescribeVaultsRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeVaultsRequest) SetVaultRegionId(v string) *DescribeVaultsRequest {
	s.VaultRegionId = &v
	return s
}

func (s *DescribeVaultsRequest) SetVaultType(v string) *DescribeVaultsRequest {
	s.VaultType = &v
	return s
}

type DescribeVaultsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVaultsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVaultsRequestTag) SetKey(v string) *DescribeVaultsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVaultsRequestTag) SetValue(v string) *DescribeVaultsRequestTag {
	s.Value = &v
	return s
}

type DescribeVaultsResponseBody struct {
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vaults     *DescribeVaultsResponseBodyVaults `json:"Vaults,omitempty" xml:"Vaults,omitempty" type:"Struct"`
}

func (s DescribeVaultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBody) SetCode(v string) *DescribeVaultsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetMessage(v string) *DescribeVaultsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetPageNumber(v int32) *DescribeVaultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetPageSize(v int32) *DescribeVaultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetRequestId(v string) *DescribeVaultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetSuccess(v bool) *DescribeVaultsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetTotalCount(v int32) *DescribeVaultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetVaults(v *DescribeVaultsResponseBodyVaults) *DescribeVaultsResponseBody {
	s.Vaults = v
	return s
}

type DescribeVaultsResponseBodyVaults struct {
	Vault []*DescribeVaultsResponseBodyVaultsVault `json:"Vault,omitempty" xml:"Vault,omitempty" type:"Repeated"`
}

func (s DescribeVaultsResponseBodyVaults) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaults) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaults) SetVault(v []*DescribeVaultsResponseBodyVaultsVault) *DescribeVaultsResponseBodyVaults {
	s.Vault = v
	return s
}

type DescribeVaultsResponseBodyVaultsVault struct {
	BackupPlanStatistics      *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics `json:"BackupPlanStatistics,omitempty" xml:"BackupPlanStatistics,omitempty" type:"Struct"`
	BucketName                *string                                                    `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BytesDone                 *int64                                                     `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	ChargeType                *string                                                    `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ChargedStorageSize        *int64                                                     `json:"ChargedStorageSize,omitempty" xml:"ChargedStorageSize,omitempty"`
	CompressionAlgorithm      *string                                                    `json:"CompressionAlgorithm,omitempty" xml:"CompressionAlgorithm,omitempty"`
	CreatedTime               *int64                                                     `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Dedup                     *bool                                                      `json:"Dedup,omitempty" xml:"Dedup,omitempty"`
	Description               *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptType               *string                                                    `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	IndexAvailable            *bool                                                      `json:"IndexAvailable,omitempty" xml:"IndexAvailable,omitempty"`
	IndexLevel                *string                                                    `json:"IndexLevel,omitempty" xml:"IndexLevel,omitempty"`
	IndexUpdateTime           *int64                                                     `json:"IndexUpdateTime,omitempty" xml:"IndexUpdateTime,omitempty"`
	KmsKeyId                  *string                                                    `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	LatestReplicationTime     *int64                                                     `json:"LatestReplicationTime,omitempty" xml:"LatestReplicationTime,omitempty"`
	RedundancyType            *string                                                    `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	Replication               *bool                                                      `json:"Replication,omitempty" xml:"Replication,omitempty"`
	ReplicationProgress       *DescribeVaultsResponseBodyVaultsVaultReplicationProgress  `json:"ReplicationProgress,omitempty" xml:"ReplicationProgress,omitempty" type:"Struct"`
	ReplicationSourceRegionId *string                                                    `json:"ReplicationSourceRegionId,omitempty" xml:"ReplicationSourceRegionId,omitempty"`
	ReplicationSourceVaultId  *string                                                    `json:"ReplicationSourceVaultId,omitempty" xml:"ReplicationSourceVaultId,omitempty"`
	ResourceGroupId           *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Retention                 *int64                                                     `json:"Retention,omitempty" xml:"Retention,omitempty"`
	SearchEnabled             *bool                                                      `json:"SearchEnabled,omitempty" xml:"SearchEnabled,omitempty"`
	SnapshotCount             *int64                                                     `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	SourceTypes               *DescribeVaultsResponseBodyVaultsVaultSourceTypes          `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty" type:"Struct"`
	Status                    *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageSize               *int64                                                     `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	Tags                      *DescribeVaultsResponseBodyVaultsVaultTags                 `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TrialInfo                 *DescribeVaultsResponseBodyVaultsVaultTrialInfo            `json:"TrialInfo,omitempty" xml:"TrialInfo,omitempty" type:"Struct"`
	UpdatedTime               *int64                                                     `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	VaultId                   *string                                                    `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	VaultName                 *string                                                    `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
	VaultRegionId             *string                                                    `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
	VaultStatusMessage        *string                                                    `json:"VaultStatusMessage,omitempty" xml:"VaultStatusMessage,omitempty"`
	VaultStorageClass         *string                                                    `json:"VaultStorageClass,omitempty" xml:"VaultStorageClass,omitempty"`
	VaultType                 *string                                                    `json:"VaultType,omitempty" xml:"VaultType,omitempty"`
	WormEnabled               *bool                                                      `json:"WormEnabled,omitempty" xml:"WormEnabled,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVault) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVault) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetBackupPlanStatistics(v *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) *DescribeVaultsResponseBodyVaultsVault {
	s.BackupPlanStatistics = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetBucketName(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.BucketName = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetBytesDone(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.BytesDone = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetChargeType(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ChargeType = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetChargedStorageSize(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.ChargedStorageSize = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetCompressionAlgorithm(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.CompressionAlgorithm = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetCreatedTime(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.CreatedTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetDedup(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.Dedup = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetDescription(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.Description = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetEncryptType(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.EncryptType = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetIndexAvailable(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.IndexAvailable = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetIndexLevel(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.IndexLevel = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetIndexUpdateTime(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.IndexUpdateTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetKmsKeyId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.KmsKeyId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetLatestReplicationTime(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.LatestReplicationTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetRedundancyType(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.RedundancyType = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplication(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.Replication = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationProgress(v *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationProgress = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationSourceRegionId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationSourceRegionId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationSourceVaultId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationSourceVaultId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetResourceGroupId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetRetention(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.Retention = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetSearchEnabled(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.SearchEnabled = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetSnapshotCount(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.SnapshotCount = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetSourceTypes(v *DescribeVaultsResponseBodyVaultsVaultSourceTypes) *DescribeVaultsResponseBodyVaultsVault {
	s.SourceTypes = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetStatus(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.Status = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetStorageSize(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.StorageSize = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetTags(v *DescribeVaultsResponseBodyVaultsVaultTags) *DescribeVaultsResponseBodyVaultsVault {
	s.Tags = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetTrialInfo(v *DescribeVaultsResponseBodyVaultsVaultTrialInfo) *DescribeVaultsResponseBodyVaultsVault {
	s.TrialInfo = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetUpdatedTime(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultName(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultName = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultRegionId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultRegionId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultStatusMessage(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultStatusMessage = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultStorageClass(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultStorageClass = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultType(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultType = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetWormEnabled(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.WormEnabled = &v
	return s
}

type DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics struct {
	CommonNas *int32 `json:"CommonNas,omitempty" xml:"CommonNas,omitempty"`
	Csg       *int32 `json:"Csg,omitempty" xml:"Csg,omitempty"`
	EcsFile   *int32 `json:"EcsFile,omitempty" xml:"EcsFile,omitempty"`
	EcsHana   *int32 `json:"EcsHana,omitempty" xml:"EcsHana,omitempty"`
	Isilon    *int32 `json:"Isilon,omitempty" xml:"Isilon,omitempty"`
	LocalFile *int32 `json:"LocalFile,omitempty" xml:"LocalFile,omitempty"`
	LocalVm   *int32 `json:"LocalVm,omitempty" xml:"LocalVm,omitempty"`
	MySql     *int32 `json:"MySql,omitempty" xml:"MySql,omitempty"`
	Nas       *int32 `json:"Nas,omitempty" xml:"Nas,omitempty"`
	Oracle    *int32 `json:"Oracle,omitempty" xml:"Oracle,omitempty"`
	Oss       *int32 `json:"Oss,omitempty" xml:"Oss,omitempty"`
	Ots       *int32 `json:"Ots,omitempty" xml:"Ots,omitempty"`
	SqlServer *int32 `json:"SqlServer,omitempty" xml:"SqlServer,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetCommonNas(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.CommonNas = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetCsg(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Csg = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetEcsFile(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.EcsFile = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetEcsHana(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.EcsHana = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetIsilon(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Isilon = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetLocalFile(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.LocalFile = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetLocalVm(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.LocalVm = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetMySql(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.MySql = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetNas(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Nas = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetOracle(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Oracle = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetOss(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Oss = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetOts(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Ots = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetSqlServer(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.SqlServer = &v
	return s
}

type DescribeVaultsResponseBodyVaultsVaultReplicationProgress struct {
	HistoricalReplicationProgress *int32 `json:"HistoricalReplicationProgress,omitempty" xml:"HistoricalReplicationProgress,omitempty"`
	NewReplicationProgress        *int64 `json:"NewReplicationProgress,omitempty" xml:"NewReplicationProgress,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVaultReplicationProgress) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultReplicationProgress) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) SetHistoricalReplicationProgress(v int32) *DescribeVaultsResponseBodyVaultsVaultReplicationProgress {
	s.HistoricalReplicationProgress = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) SetNewReplicationProgress(v int64) *DescribeVaultsResponseBodyVaultsVaultReplicationProgress {
	s.NewReplicationProgress = &v
	return s
}

type DescribeVaultsResponseBodyVaultsVaultSourceTypes struct {
	SourceType []*string `json:"SourceType,omitempty" xml:"SourceType,omitempty" type:"Repeated"`
}

func (s DescribeVaultsResponseBodyVaultsVaultSourceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultSourceTypes) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultSourceTypes) SetSourceType(v []*string) *DescribeVaultsResponseBodyVaultsVaultSourceTypes {
	s.SourceType = v
	return s
}

type DescribeVaultsResponseBodyVaultsVaultTags struct {
	Tag []*DescribeVaultsResponseBodyVaultsVaultTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVaultsResponseBodyVaultsVaultTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultTags) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultTags) SetTag(v []*DescribeVaultsResponseBodyVaultsVaultTagsTag) *DescribeVaultsResponseBodyVaultsVaultTags {
	s.Tag = v
	return s
}

type DescribeVaultsResponseBodyVaultsVaultTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVaultTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultTagsTag) SetKey(v string) *DescribeVaultsResponseBodyVaultsVaultTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTagsTag) SetValue(v string) *DescribeVaultsResponseBodyVaultsVaultTagsTag {
	s.Value = &v
	return s
}

type DescribeVaultsResponseBodyVaultsVaultTrialInfo struct {
	KeepAfterTrialExpiration *bool  `json:"KeepAfterTrialExpiration,omitempty" xml:"KeepAfterTrialExpiration,omitempty"`
	TrialExpireTime          *int64 `json:"TrialExpireTime,omitempty" xml:"TrialExpireTime,omitempty"`
	TrialStartTime           *int64 `json:"TrialStartTime,omitempty" xml:"TrialStartTime,omitempty"`
	TrialVaultReleaseTime    *int64 `json:"TrialVaultReleaseTime,omitempty" xml:"TrialVaultReleaseTime,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVaultTrialInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultTrialInfo) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) SetKeepAfterTrialExpiration(v bool) *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	s.KeepAfterTrialExpiration = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) SetTrialExpireTime(v int64) *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	s.TrialExpireTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) SetTrialStartTime(v int64) *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	s.TrialStartTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) SetTrialVaultReleaseTime(v int64) *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	s.TrialVaultReleaseTime = &v
	return s
}

type DescribeVaultsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVaultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVaultsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVaultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponse) SetHeaders(v map[string]*string) *DescribeVaultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVaultsResponse) SetStatusCode(v int32) *DescribeVaultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVaultsResponse) SetBody(v *DescribeVaultsResponseBody) *DescribeVaultsResponse {
	s.Body = v
	return s
}

type DetachNasFileSystemRequest struct {
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	FileSystemId         *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DetachNasFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachNasFileSystemRequest) GoString() string {
	return s.String()
}

func (s *DetachNasFileSystemRequest) SetCreateTime(v string) *DetachNasFileSystemRequest {
	s.CreateTime = &v
	return s
}

func (s *DetachNasFileSystemRequest) SetCrossAccountRoleName(v string) *DetachNasFileSystemRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DetachNasFileSystemRequest) SetCrossAccountType(v string) *DetachNasFileSystemRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DetachNasFileSystemRequest) SetCrossAccountUserId(v int64) *DetachNasFileSystemRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DetachNasFileSystemRequest) SetFileSystemId(v string) *DetachNasFileSystemRequest {
	s.FileSystemId = &v
	return s
}

type DetachNasFileSystemResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetachNasFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachNasFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *DetachNasFileSystemResponseBody) SetCode(v string) *DetachNasFileSystemResponseBody {
	s.Code = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) SetMessage(v string) *DetachNasFileSystemResponseBody {
	s.Message = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) SetRequestId(v string) *DetachNasFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) SetSuccess(v bool) *DetachNasFileSystemResponseBody {
	s.Success = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) SetTaskId(v string) *DetachNasFileSystemResponseBody {
	s.TaskId = &v
	return s
}

type DetachNasFileSystemResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachNasFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachNasFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachNasFileSystemResponse) GoString() string {
	return s.String()
}

func (s *DetachNasFileSystemResponse) SetHeaders(v map[string]*string) *DetachNasFileSystemResponse {
	s.Headers = v
	return s
}

func (s *DetachNasFileSystemResponse) SetStatusCode(v int32) *DetachNasFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachNasFileSystemResponse) SetBody(v *DetachNasFileSystemResponseBody) *DetachNasFileSystemResponse {
	s.Body = v
	return s
}

type DisableBackupPlanRequest struct {
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VaultId    *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DisableBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DisableBackupPlanRequest) SetPlanId(v string) *DisableBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DisableBackupPlanRequest) SetSourceType(v string) *DisableBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *DisableBackupPlanRequest) SetVaultId(v string) *DisableBackupPlanRequest {
	s.VaultId = &v
	return s
}

type DisableBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DisableBackupPlanResponseBody) SetCode(v string) *DisableBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DisableBackupPlanResponseBody) SetMessage(v string) *DisableBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DisableBackupPlanResponseBody) SetRequestId(v string) *DisableBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableBackupPlanResponseBody) SetSuccess(v bool) *DisableBackupPlanResponseBody {
	s.Success = &v
	return s
}

type DisableBackupPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DisableBackupPlanResponse) SetHeaders(v map[string]*string) *DisableBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DisableBackupPlanResponse) SetStatusCode(v int32) *DisableBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableBackupPlanResponse) SetBody(v *DisableBackupPlanResponseBody) *DisableBackupPlanResponse {
	s.Body = v
	return s
}

type DisableHanaBackupPlanRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PlanId          *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DisableHanaBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DisableHanaBackupPlanRequest) SetClusterId(v string) *DisableHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *DisableHanaBackupPlanRequest) SetPlanId(v string) *DisableHanaBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DisableHanaBackupPlanRequest) SetResourceGroupId(v string) *DisableHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableHanaBackupPlanRequest) SetVaultId(v string) *DisableHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

type DisableHanaBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableHanaBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DisableHanaBackupPlanResponseBody) SetCode(v string) *DisableHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DisableHanaBackupPlanResponseBody) SetMessage(v string) *DisableHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DisableHanaBackupPlanResponseBody) SetRequestId(v string) *DisableHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableHanaBackupPlanResponseBody) SetSuccess(v bool) *DisableHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

type DisableHanaBackupPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableHanaBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DisableHanaBackupPlanResponse) SetHeaders(v map[string]*string) *DisableHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DisableHanaBackupPlanResponse) SetStatusCode(v int32) *DisableHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableHanaBackupPlanResponse) SetBody(v *DisableHanaBackupPlanResponseBody) *DisableHanaBackupPlanResponse {
	s.Body = v
	return s
}

type EnableBackupPlanRequest struct {
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VaultId    *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s EnableBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *EnableBackupPlanRequest) SetPlanId(v string) *EnableBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *EnableBackupPlanRequest) SetSourceType(v string) *EnableBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *EnableBackupPlanRequest) SetVaultId(v string) *EnableBackupPlanRequest {
	s.VaultId = &v
	return s
}

type EnableBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *EnableBackupPlanResponseBody) SetCode(v string) *EnableBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *EnableBackupPlanResponseBody) SetMessage(v string) *EnableBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *EnableBackupPlanResponseBody) SetRequestId(v string) *EnableBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableBackupPlanResponseBody) SetSuccess(v bool) *EnableBackupPlanResponseBody {
	s.Success = &v
	return s
}

type EnableBackupPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *EnableBackupPlanResponse) SetHeaders(v map[string]*string) *EnableBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *EnableBackupPlanResponse) SetStatusCode(v int32) *EnableBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableBackupPlanResponse) SetBody(v *EnableBackupPlanResponseBody) *EnableBackupPlanResponse {
	s.Body = v
	return s
}

type EnableHanaBackupPlanRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PlanId          *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s EnableHanaBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *EnableHanaBackupPlanRequest) SetClusterId(v string) *EnableHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *EnableHanaBackupPlanRequest) SetPlanId(v string) *EnableHanaBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *EnableHanaBackupPlanRequest) SetResourceGroupId(v string) *EnableHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableHanaBackupPlanRequest) SetVaultId(v string) *EnableHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

type EnableHanaBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableHanaBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *EnableHanaBackupPlanResponseBody) SetCode(v string) *EnableHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *EnableHanaBackupPlanResponseBody) SetMessage(v string) *EnableHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *EnableHanaBackupPlanResponseBody) SetRequestId(v string) *EnableHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableHanaBackupPlanResponseBody) SetSuccess(v bool) *EnableHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

type EnableHanaBackupPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableHanaBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *EnableHanaBackupPlanResponse) SetHeaders(v map[string]*string) *EnableHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *EnableHanaBackupPlanResponse) SetStatusCode(v int32) *EnableHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableHanaBackupPlanResponse) SetBody(v *EnableHanaBackupPlanResponseBody) *EnableHanaBackupPlanResponse {
	s.Body = v
	return s
}

type ExecuteBackupPlanRequest struct {
	PlanId     *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RuleId     *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VaultId    *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s ExecuteBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *ExecuteBackupPlanRequest) SetPlanId(v string) *ExecuteBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *ExecuteBackupPlanRequest) SetRuleId(v string) *ExecuteBackupPlanRequest {
	s.RuleId = &v
	return s
}

func (s *ExecuteBackupPlanRequest) SetSourceType(v string) *ExecuteBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *ExecuteBackupPlanRequest) SetVaultId(v string) *ExecuteBackupPlanRequest {
	s.VaultId = &v
	return s
}

type ExecuteBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteBackupPlanResponseBody) SetCode(v string) *ExecuteBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteBackupPlanResponseBody) SetJobId(v string) *ExecuteBackupPlanResponseBody {
	s.JobId = &v
	return s
}

func (s *ExecuteBackupPlanResponseBody) SetMessage(v string) *ExecuteBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteBackupPlanResponseBody) SetRequestId(v string) *ExecuteBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteBackupPlanResponseBody) SetSuccess(v bool) *ExecuteBackupPlanResponseBody {
	s.Success = &v
	return s
}

type ExecuteBackupPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *ExecuteBackupPlanResponse) SetHeaders(v map[string]*string) *ExecuteBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *ExecuteBackupPlanResponse) SetStatusCode(v int32) *ExecuteBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteBackupPlanResponse) SetBody(v *ExecuteBackupPlanResponseBody) *ExecuteBackupPlanResponse {
	s.Body = v
	return s
}

type GenerateRamPolicyRequest struct {
	ActionType        *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	RequireBasePolicy *bool   `json:"RequireBasePolicy,omitempty" xml:"RequireBasePolicy,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId           *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s GenerateRamPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateRamPolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateRamPolicyRequest) SetActionType(v string) *GenerateRamPolicyRequest {
	s.ActionType = &v
	return s
}

func (s *GenerateRamPolicyRequest) SetRequireBasePolicy(v bool) *GenerateRamPolicyRequest {
	s.RequireBasePolicy = &v
	return s
}

func (s *GenerateRamPolicyRequest) SetResourceGroupId(v string) *GenerateRamPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GenerateRamPolicyRequest) SetVaultId(v string) *GenerateRamPolicyRequest {
	s.VaultId = &v
	return s
}

type GenerateRamPolicyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateRamPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateRamPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateRamPolicyResponseBody) SetCode(v string) *GenerateRamPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) SetMessage(v string) *GenerateRamPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) SetPolicyDocument(v string) *GenerateRamPolicyResponseBody {
	s.PolicyDocument = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) SetRequestId(v string) *GenerateRamPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateRamPolicyResponseBody) SetSuccess(v bool) *GenerateRamPolicyResponseBody {
	s.Success = &v
	return s
}

type GenerateRamPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateRamPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateRamPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateRamPolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateRamPolicyResponse) SetHeaders(v map[string]*string) *GenerateRamPolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateRamPolicyResponse) SetStatusCode(v int32) *GenerateRamPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateRamPolicyResponse) SetBody(v *GenerateRamPolicyResponseBody) *GenerateRamPolicyResponse {
	s.Body = v
	return s
}

type GetTempFileDownloadLinkRequest struct {
	TempFileKey *string `json:"TempFileKey,omitempty" xml:"TempFileKey,omitempty"`
}

func (s GetTempFileDownloadLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTempFileDownloadLinkRequest) GoString() string {
	return s.String()
}

func (s *GetTempFileDownloadLinkRequest) SetTempFileKey(v string) *GetTempFileDownloadLinkRequest {
	s.TempFileKey = &v
	return s
}

type GetTempFileDownloadLinkResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTempFileDownloadLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTempFileDownloadLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetTempFileDownloadLinkResponseBody) SetCode(v string) *GetTempFileDownloadLinkResponseBody {
	s.Code = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) SetMessage(v string) *GetTempFileDownloadLinkResponseBody {
	s.Message = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) SetRequestId(v string) *GetTempFileDownloadLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) SetSuccess(v bool) *GetTempFileDownloadLinkResponseBody {
	s.Success = &v
	return s
}

func (s *GetTempFileDownloadLinkResponseBody) SetUrl(v string) *GetTempFileDownloadLinkResponseBody {
	s.Url = &v
	return s
}

type GetTempFileDownloadLinkResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTempFileDownloadLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTempFileDownloadLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTempFileDownloadLinkResponse) GoString() string {
	return s.String()
}

func (s *GetTempFileDownloadLinkResponse) SetHeaders(v map[string]*string) *GetTempFileDownloadLinkResponse {
	s.Headers = v
	return s
}

func (s *GetTempFileDownloadLinkResponse) SetStatusCode(v int32) *GetTempFileDownloadLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTempFileDownloadLinkResponse) SetBody(v *GetTempFileDownloadLinkResponseBody) *GetTempFileDownloadLinkResponse {
	s.Body = v
	return s
}

type InstallBackupClientsRequest struct {
	CrossAccountRoleName *string                `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	InstanceIds          map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s InstallBackupClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsRequest) SetCrossAccountRoleName(v string) *InstallBackupClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *InstallBackupClientsRequest) SetCrossAccountType(v string) *InstallBackupClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *InstallBackupClientsRequest) SetCrossAccountUserId(v int64) *InstallBackupClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *InstallBackupClientsRequest) SetInstanceIds(v map[string]interface{}) *InstallBackupClientsRequest {
	s.InstanceIds = v
	return s
}

type InstallBackupClientsShrinkRequest struct {
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	InstanceIdsShrink    *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s InstallBackupClientsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallBackupClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsShrinkRequest) SetCrossAccountRoleName(v string) *InstallBackupClientsShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *InstallBackupClientsShrinkRequest) SetCrossAccountType(v string) *InstallBackupClientsShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *InstallBackupClientsShrinkRequest) SetCrossAccountUserId(v int64) *InstallBackupClientsShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *InstallBackupClientsShrinkRequest) SetInstanceIdsShrink(v string) *InstallBackupClientsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

type InstallBackupClientsResponseBody struct {
	Code             *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceStatuses []*InstallBackupClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
	Message          *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId           *string                                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InstallBackupClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsResponseBody) SetCode(v string) *InstallBackupClientsResponseBody {
	s.Code = &v
	return s
}

func (s *InstallBackupClientsResponseBody) SetInstanceStatuses(v []*InstallBackupClientsResponseBodyInstanceStatuses) *InstallBackupClientsResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *InstallBackupClientsResponseBody) SetMessage(v string) *InstallBackupClientsResponseBody {
	s.Message = &v
	return s
}

func (s *InstallBackupClientsResponseBody) SetRequestId(v string) *InstallBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallBackupClientsResponseBody) SetSuccess(v bool) *InstallBackupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *InstallBackupClientsResponseBody) SetTaskId(v string) *InstallBackupClientsResponseBody {
	s.TaskId = &v
	return s
}

type InstallBackupClientsResponseBodyInstanceStatuses struct {
	ErrorCode     *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ValidInstance *bool   `json:"ValidInstance,omitempty" xml:"ValidInstance,omitempty"`
}

func (s InstallBackupClientsResponseBodyInstanceStatuses) String() string {
	return tea.Prettify(s)
}

func (s InstallBackupClientsResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) SetErrorCode(v string) *InstallBackupClientsResponseBodyInstanceStatuses {
	s.ErrorCode = &v
	return s
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) SetInstanceId(v string) *InstallBackupClientsResponseBodyInstanceStatuses {
	s.InstanceId = &v
	return s
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) SetValidInstance(v bool) *InstallBackupClientsResponseBodyInstanceStatuses {
	s.ValidInstance = &v
	return s
}

type InstallBackupClientsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallBackupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallBackupClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallBackupClientsResponse) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsResponse) SetHeaders(v map[string]*string) *InstallBackupClientsResponse {
	s.Headers = v
	return s
}

func (s *InstallBackupClientsResponse) SetStatusCode(v int32) *InstallBackupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallBackupClientsResponse) SetBody(v *InstallBackupClientsResponseBody) *InstallBackupClientsResponse {
	s.Body = v
	return s
}

type OpenHbrServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenHbrServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenHbrServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenHbrServiceResponseBody) SetOrderId(v string) *OpenHbrServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenHbrServiceResponseBody) SetRequestId(v string) *OpenHbrServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenHbrServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenHbrServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenHbrServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenHbrServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenHbrServiceResponse) SetHeaders(v map[string]*string) *OpenHbrServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenHbrServiceResponse) SetStatusCode(v int32) *OpenHbrServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenHbrServiceResponse) SetBody(v *OpenHbrServiceResponseBody) *OpenHbrServiceResponse {
	s.Body = v
	return s
}

type SearchHistoricalSnapshotsRequest struct {
	Limit      *int32                 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken  *string                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Query      map[string]interface{} `json:"Query,omitempty" xml:"Query,omitempty"`
	SourceType *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SearchHistoricalSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchHistoricalSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsRequest) SetLimit(v int32) *SearchHistoricalSnapshotsRequest {
	s.Limit = &v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetNextToken(v string) *SearchHistoricalSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetQuery(v map[string]interface{}) *SearchHistoricalSnapshotsRequest {
	s.Query = v
	return s
}

func (s *SearchHistoricalSnapshotsRequest) SetSourceType(v string) *SearchHistoricalSnapshotsRequest {
	s.SourceType = &v
	return s
}

type SearchHistoricalSnapshotsShrinkRequest struct {
	Limit       *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SourceType  *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SearchHistoricalSnapshotsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchHistoricalSnapshotsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsShrinkRequest) SetLimit(v int32) *SearchHistoricalSnapshotsShrinkRequest {
	s.Limit = &v
	return s
}

func (s *SearchHistoricalSnapshotsShrinkRequest) SetNextToken(v string) *SearchHistoricalSnapshotsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *SearchHistoricalSnapshotsShrinkRequest) SetQueryShrink(v string) *SearchHistoricalSnapshotsShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *SearchHistoricalSnapshotsShrinkRequest) SetSourceType(v string) *SearchHistoricalSnapshotsShrinkRequest {
	s.SourceType = &v
	return s
}

type SearchHistoricalSnapshotsResponseBody struct {
	Code       *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Limit      *int32                                          `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots  *SearchHistoricalSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Success    *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchHistoricalSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponseBody) SetCode(v string) *SearchHistoricalSnapshotsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetLimit(v int32) *SearchHistoricalSnapshotsResponseBody {
	s.Limit = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetMessage(v string) *SearchHistoricalSnapshotsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetNextToken(v string) *SearchHistoricalSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetRequestId(v string) *SearchHistoricalSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetSnapshots(v *SearchHistoricalSnapshotsResponseBodySnapshots) *SearchHistoricalSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetSuccess(v bool) *SearchHistoricalSnapshotsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetTotalCount(v int32) *SearchHistoricalSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

type SearchHistoricalSnapshotsResponseBodySnapshots struct {
	Snapshot []*SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s SearchHistoricalSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshots) SetSnapshot(v []*SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) *SearchHistoricalSnapshotsResponseBodySnapshots {
	s.Snapshot = v
	return s
}

type SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot struct {
	ActualBytes        *int64                                                       `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	ActualItems        *int64                                                       `json:"ActualItems,omitempty" xml:"ActualItems,omitempty"`
	BackupType         *string                                                      `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	Bucket             *string                                                      `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	BytesDone          *int64                                                       `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	BytesTotal         *int64                                                       `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	ClientId           *string                                                      `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CompleteTime       *int64                                                       `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreateTime         *int64                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedTime        *int64                                                       `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ErrorFile          *string                                                      `json:"ErrorFile,omitempty" xml:"ErrorFile,omitempty"`
	ExpireTime         *int64                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileSystemId       *string                                                      `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InstanceId         *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName       *string                                                      `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ItemsDone          *int64                                                       `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	ItemsTotal         *int64                                                       `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	JobId              *string                                                      `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ParentSnapshotHash *string                                                      `json:"ParentSnapshotHash,omitempty" xml:"ParentSnapshotHash,omitempty"`
	Path               *string                                                      `json:"Path,omitempty" xml:"Path,omitempty"`
	Paths              *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Struct"`
	Prefix             *string                                                      `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	RangeEnd           *int64                                                       `json:"RangeEnd,omitempty" xml:"RangeEnd,omitempty"`
	RangeStart         *int64                                                       `json:"RangeStart,omitempty" xml:"RangeStart,omitempty"`
	Retention          *int64                                                       `json:"Retention,omitempty" xml:"Retention,omitempty"`
	SnapshotHash       *string                                                      `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	SnapshotId         *string                                                      `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SourceType         *string                                                      `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartTime          *int64                                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status             *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName          *string                                                      `json:"TableName,omitempty" xml:"TableName,omitempty"`
	UpdatedTime        *int64                                                       `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	VaultId            *string                                                      `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) String() string {
	return tea.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetActualBytes(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ActualBytes = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetActualItems(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ActualItems = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetBackupType(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.BackupType = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetBucket(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Bucket = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetBytesDone(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.BytesDone = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetBytesTotal(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.BytesTotal = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetClientId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ClientId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetCompleteTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.CompleteTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetCreateTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.CreateTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetCreatedTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.CreatedTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetErrorFile(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ErrorFile = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetExpireTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ExpireTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetFileSystemId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.FileSystemId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetInstanceId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.InstanceId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetInstanceName(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.InstanceName = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetItemsDone(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ItemsDone = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetItemsTotal(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ItemsTotal = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetJobId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.JobId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetParentSnapshotHash(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ParentSnapshotHash = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetPath(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Path = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetPaths(v *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Paths = v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetPrefix(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Prefix = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetRangeEnd(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.RangeEnd = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetRangeStart(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.RangeStart = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetRetention(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Retention = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotHash(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotHash = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetSourceType(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceType = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetStartTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.StartTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetStatus(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Status = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetTableName(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.TableName = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetUpdatedTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.UpdatedTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetVaultId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.VaultId = &v
	return s
}

type SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths struct {
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) String() string {
	return tea.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) SetPath(v []*string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths {
	s.Path = v
	return s
}

type SearchHistoricalSnapshotsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchHistoricalSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchHistoricalSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponse) SetHeaders(v map[string]*string) *SearchHistoricalSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *SearchHistoricalSnapshotsResponse) SetStatusCode(v int32) *SearchHistoricalSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponse) SetBody(v *SearchHistoricalSnapshotsResponseBody) *SearchHistoricalSnapshotsResponse {
	s.Body = v
	return s
}

type StartHanaDatabaseAsyncRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	VaultId      *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s StartHanaDatabaseAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s StartHanaDatabaseAsyncRequest) GoString() string {
	return s.String()
}

func (s *StartHanaDatabaseAsyncRequest) SetClusterId(v string) *StartHanaDatabaseAsyncRequest {
	s.ClusterId = &v
	return s
}

func (s *StartHanaDatabaseAsyncRequest) SetDatabaseName(v string) *StartHanaDatabaseAsyncRequest {
	s.DatabaseName = &v
	return s
}

func (s *StartHanaDatabaseAsyncRequest) SetVaultId(v string) *StartHanaDatabaseAsyncRequest {
	s.VaultId = &v
	return s
}

type StartHanaDatabaseAsyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartHanaDatabaseAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartHanaDatabaseAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *StartHanaDatabaseAsyncResponseBody) SetCode(v string) *StartHanaDatabaseAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) SetMessage(v string) *StartHanaDatabaseAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) SetRequestId(v string) *StartHanaDatabaseAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) SetSuccess(v bool) *StartHanaDatabaseAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) SetTaskId(v string) *StartHanaDatabaseAsyncResponseBody {
	s.TaskId = &v
	return s
}

type StartHanaDatabaseAsyncResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartHanaDatabaseAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartHanaDatabaseAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s StartHanaDatabaseAsyncResponse) GoString() string {
	return s.String()
}

func (s *StartHanaDatabaseAsyncResponse) SetHeaders(v map[string]*string) *StartHanaDatabaseAsyncResponse {
	s.Headers = v
	return s
}

func (s *StartHanaDatabaseAsyncResponse) SetStatusCode(v int32) *StartHanaDatabaseAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponse) SetBody(v *StartHanaDatabaseAsyncResponseBody) *StartHanaDatabaseAsyncResponse {
	s.Body = v
	return s
}

type StopHanaDatabaseAsyncRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	VaultId      *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s StopHanaDatabaseAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s StopHanaDatabaseAsyncRequest) GoString() string {
	return s.String()
}

func (s *StopHanaDatabaseAsyncRequest) SetClusterId(v string) *StopHanaDatabaseAsyncRequest {
	s.ClusterId = &v
	return s
}

func (s *StopHanaDatabaseAsyncRequest) SetDatabaseName(v string) *StopHanaDatabaseAsyncRequest {
	s.DatabaseName = &v
	return s
}

func (s *StopHanaDatabaseAsyncRequest) SetVaultId(v string) *StopHanaDatabaseAsyncRequest {
	s.VaultId = &v
	return s
}

type StopHanaDatabaseAsyncResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopHanaDatabaseAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopHanaDatabaseAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *StopHanaDatabaseAsyncResponseBody) SetCode(v string) *StopHanaDatabaseAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) SetMessage(v string) *StopHanaDatabaseAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) SetRequestId(v string) *StopHanaDatabaseAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) SetSuccess(v bool) *StopHanaDatabaseAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) SetTaskId(v string) *StopHanaDatabaseAsyncResponseBody {
	s.TaskId = &v
	return s
}

type StopHanaDatabaseAsyncResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopHanaDatabaseAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopHanaDatabaseAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s StopHanaDatabaseAsyncResponse) GoString() string {
	return s.String()
}

func (s *StopHanaDatabaseAsyncResponse) SetHeaders(v map[string]*string) *StopHanaDatabaseAsyncResponse {
	s.Headers = v
	return s
}

func (s *StopHanaDatabaseAsyncResponse) SetStatusCode(v int32) *StopHanaDatabaseAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponse) SetBody(v *StopHanaDatabaseAsyncResponseBody) *StopHanaDatabaseAsyncResponse {
	s.Body = v
	return s
}

type UninstallBackupClientsRequest struct {
	ClientIds            map[string]interface{} `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	CrossAccountRoleName *string                `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	InstanceIds          map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s UninstallBackupClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsRequest) SetClientIds(v map[string]interface{}) *UninstallBackupClientsRequest {
	s.ClientIds = v
	return s
}

func (s *UninstallBackupClientsRequest) SetCrossAccountRoleName(v string) *UninstallBackupClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UninstallBackupClientsRequest) SetCrossAccountType(v string) *UninstallBackupClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *UninstallBackupClientsRequest) SetCrossAccountUserId(v int64) *UninstallBackupClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *UninstallBackupClientsRequest) SetInstanceIds(v map[string]interface{}) *UninstallBackupClientsRequest {
	s.InstanceIds = v
	return s
}

type UninstallBackupClientsShrinkRequest struct {
	ClientIdsShrink      *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	InstanceIdsShrink    *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s UninstallBackupClientsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallBackupClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsShrinkRequest) SetClientIdsShrink(v string) *UninstallBackupClientsShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) SetCrossAccountRoleName(v string) *UninstallBackupClientsShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) SetCrossAccountType(v string) *UninstallBackupClientsShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) SetCrossAccountUserId(v int64) *UninstallBackupClientsShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) SetInstanceIdsShrink(v string) *UninstallBackupClientsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

type UninstallBackupClientsResponseBody struct {
	Code             *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceStatuses []*UninstallBackupClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
	Message          *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId           *string                                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UninstallBackupClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsResponseBody) SetCode(v string) *UninstallBackupClientsResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetInstanceStatuses(v []*UninstallBackupClientsResponseBodyInstanceStatuses) *UninstallBackupClientsResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetMessage(v string) *UninstallBackupClientsResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetRequestId(v string) *UninstallBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetSuccess(v bool) *UninstallBackupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetTaskId(v string) *UninstallBackupClientsResponseBody {
	s.TaskId = &v
	return s
}

type UninstallBackupClientsResponseBodyInstanceStatuses struct {
	ErrorCode     *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ValidInstance *bool   `json:"ValidInstance,omitempty" xml:"ValidInstance,omitempty"`
}

func (s UninstallBackupClientsResponseBodyInstanceStatuses) String() string {
	return tea.Prettify(s)
}

func (s UninstallBackupClientsResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) SetErrorCode(v string) *UninstallBackupClientsResponseBodyInstanceStatuses {
	s.ErrorCode = &v
	return s
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) SetInstanceId(v string) *UninstallBackupClientsResponseBodyInstanceStatuses {
	s.InstanceId = &v
	return s
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) SetValidInstance(v bool) *UninstallBackupClientsResponseBodyInstanceStatuses {
	s.ValidInstance = &v
	return s
}

type UninstallBackupClientsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UninstallBackupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallBackupClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallBackupClientsResponse) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsResponse) SetHeaders(v map[string]*string) *UninstallBackupClientsResponse {
	s.Headers = v
	return s
}

func (s *UninstallBackupClientsResponse) SetStatusCode(v int32) *UninstallBackupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallBackupClientsResponse) SetBody(v *UninstallBackupClientsResponseBody) *UninstallBackupClientsResponse {
	s.Body = v
	return s
}

type UninstallClientRequest struct {
	ClientId        *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UninstallClientRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallClientRequest) GoString() string {
	return s.String()
}

func (s *UninstallClientRequest) SetClientId(v string) *UninstallClientRequest {
	s.ClientId = &v
	return s
}

func (s *UninstallClientRequest) SetResourceGroupId(v string) *UninstallClientRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UninstallClientRequest) SetVaultId(v string) *UninstallClientRequest {
	s.VaultId = &v
	return s
}

type UninstallClientResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UninstallClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallClientResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallClientResponseBody) SetCode(v string) *UninstallClientResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallClientResponseBody) SetMessage(v string) *UninstallClientResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallClientResponseBody) SetRequestId(v string) *UninstallClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallClientResponseBody) SetSuccess(v bool) *UninstallClientResponseBody {
	s.Success = &v
	return s
}

func (s *UninstallClientResponseBody) SetTaskId(v string) *UninstallClientResponseBody {
	s.TaskId = &v
	return s
}

type UninstallClientResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UninstallClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallClientResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallClientResponse) GoString() string {
	return s.String()
}

func (s *UninstallClientResponse) SetHeaders(v map[string]*string) *UninstallClientResponse {
	s.Headers = v
	return s
}

func (s *UninstallClientResponse) SetStatusCode(v int32) *UninstallClientResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallClientResponse) SetBody(v *UninstallClientResponseBody) *UninstallClientResponse {
	s.Body = v
	return s
}

type UpdateBackupPlanRequest struct {
	Detail      map[string]interface{}         `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Exclude     *string                        `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	Include     *string                        `json:"Include,omitempty" xml:"Include,omitempty"`
	Options     *string                        `json:"Options,omitempty" xml:"Options,omitempty"`
	OtsDetail   *OtsDetail                     `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	Path        []*string                      `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	PlanId      *string                        `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName    *string                        `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	Prefix      *string                        `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Retention   *int64                         `json:"Retention,omitempty" xml:"Retention,omitempty"`
	Rule        []*UpdateBackupPlanRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	Schedule    *string                        `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	SourceType  *string                        `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SpeedLimit  *string                        `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	UpdatePaths *bool                          `json:"UpdatePaths,omitempty" xml:"UpdatePaths,omitempty"`
	VaultId     *string                        `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanRequest) SetDetail(v map[string]interface{}) *UpdateBackupPlanRequest {
	s.Detail = v
	return s
}

func (s *UpdateBackupPlanRequest) SetExclude(v string) *UpdateBackupPlanRequest {
	s.Exclude = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetInclude(v string) *UpdateBackupPlanRequest {
	s.Include = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetOptions(v string) *UpdateBackupPlanRequest {
	s.Options = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetOtsDetail(v *OtsDetail) *UpdateBackupPlanRequest {
	s.OtsDetail = v
	return s
}

func (s *UpdateBackupPlanRequest) SetPath(v []*string) *UpdateBackupPlanRequest {
	s.Path = v
	return s
}

func (s *UpdateBackupPlanRequest) SetPlanId(v string) *UpdateBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetPlanName(v string) *UpdateBackupPlanRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetPrefix(v string) *UpdateBackupPlanRequest {
	s.Prefix = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetRetention(v int64) *UpdateBackupPlanRequest {
	s.Retention = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetRule(v []*UpdateBackupPlanRequestRule) *UpdateBackupPlanRequest {
	s.Rule = v
	return s
}

func (s *UpdateBackupPlanRequest) SetSchedule(v string) *UpdateBackupPlanRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetSourceType(v string) *UpdateBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetSpeedLimit(v string) *UpdateBackupPlanRequest {
	s.SpeedLimit = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetUpdatePaths(v bool) *UpdateBackupPlanRequest {
	s.UpdatePaths = &v
	return s
}

func (s *UpdateBackupPlanRequest) SetVaultId(v string) *UpdateBackupPlanRequest {
	s.VaultId = &v
	return s
}

type UpdateBackupPlanRequestRule struct {
	BackupType           *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DestinationRegionId  *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationRetention *int64  `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	Disabled             *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DoCopy               *bool   `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	Retention            *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Schedule             *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s UpdateBackupPlanRequestRule) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPlanRequestRule) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanRequestRule) SetBackupType(v string) *UpdateBackupPlanRequestRule {
	s.BackupType = &v
	return s
}

func (s *UpdateBackupPlanRequestRule) SetDestinationRegionId(v string) *UpdateBackupPlanRequestRule {
	s.DestinationRegionId = &v
	return s
}

func (s *UpdateBackupPlanRequestRule) SetDestinationRetention(v int64) *UpdateBackupPlanRequestRule {
	s.DestinationRetention = &v
	return s
}

func (s *UpdateBackupPlanRequestRule) SetDisabled(v bool) *UpdateBackupPlanRequestRule {
	s.Disabled = &v
	return s
}

func (s *UpdateBackupPlanRequestRule) SetDoCopy(v bool) *UpdateBackupPlanRequestRule {
	s.DoCopy = &v
	return s
}

func (s *UpdateBackupPlanRequestRule) SetRetention(v int64) *UpdateBackupPlanRequestRule {
	s.Retention = &v
	return s
}

func (s *UpdateBackupPlanRequestRule) SetRuleName(v string) *UpdateBackupPlanRequestRule {
	s.RuleName = &v
	return s
}

func (s *UpdateBackupPlanRequestRule) SetSchedule(v string) *UpdateBackupPlanRequestRule {
	s.Schedule = &v
	return s
}

type UpdateBackupPlanShrinkRequest struct {
	DetailShrink    *string                              `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Exclude         *string                              `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	Include         *string                              `json:"Include,omitempty" xml:"Include,omitempty"`
	Options         *string                              `json:"Options,omitempty" xml:"Options,omitempty"`
	OtsDetailShrink *string                              `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	Path            []*string                            `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	PlanId          *string                              `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName        *string                              `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	Prefix          *string                              `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Retention       *int64                               `json:"Retention,omitempty" xml:"Retention,omitempty"`
	Rule            []*UpdateBackupPlanShrinkRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	Schedule        *string                              `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	SourceType      *string                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SpeedLimit      *string                              `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	UpdatePaths     *bool                                `json:"UpdatePaths,omitempty" xml:"UpdatePaths,omitempty"`
	VaultId         *string                              `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateBackupPlanShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanShrinkRequest) SetDetailShrink(v string) *UpdateBackupPlanShrinkRequest {
	s.DetailShrink = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetExclude(v string) *UpdateBackupPlanShrinkRequest {
	s.Exclude = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetInclude(v string) *UpdateBackupPlanShrinkRequest {
	s.Include = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetOptions(v string) *UpdateBackupPlanShrinkRequest {
	s.Options = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetOtsDetailShrink(v string) *UpdateBackupPlanShrinkRequest {
	s.OtsDetailShrink = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetPath(v []*string) *UpdateBackupPlanShrinkRequest {
	s.Path = v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetPlanId(v string) *UpdateBackupPlanShrinkRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetPlanName(v string) *UpdateBackupPlanShrinkRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetPrefix(v string) *UpdateBackupPlanShrinkRequest {
	s.Prefix = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetRetention(v int64) *UpdateBackupPlanShrinkRequest {
	s.Retention = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetRule(v []*UpdateBackupPlanShrinkRequestRule) *UpdateBackupPlanShrinkRequest {
	s.Rule = v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetSchedule(v string) *UpdateBackupPlanShrinkRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetSourceType(v string) *UpdateBackupPlanShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetSpeedLimit(v string) *UpdateBackupPlanShrinkRequest {
	s.SpeedLimit = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetUpdatePaths(v bool) *UpdateBackupPlanShrinkRequest {
	s.UpdatePaths = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetVaultId(v string) *UpdateBackupPlanShrinkRequest {
	s.VaultId = &v
	return s
}

type UpdateBackupPlanShrinkRequestRule struct {
	BackupType           *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DestinationRegionId  *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationRetention *int64  `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	Disabled             *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DoCopy               *bool   `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	Retention            *int64  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Schedule             *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s UpdateBackupPlanShrinkRequestRule) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPlanShrinkRequestRule) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanShrinkRequestRule) SetBackupType(v string) *UpdateBackupPlanShrinkRequestRule {
	s.BackupType = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetDestinationRegionId(v string) *UpdateBackupPlanShrinkRequestRule {
	s.DestinationRegionId = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetDestinationRetention(v int64) *UpdateBackupPlanShrinkRequestRule {
	s.DestinationRetention = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetDisabled(v bool) *UpdateBackupPlanShrinkRequestRule {
	s.Disabled = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetDoCopy(v bool) *UpdateBackupPlanShrinkRequestRule {
	s.DoCopy = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetRetention(v int64) *UpdateBackupPlanShrinkRequestRule {
	s.Retention = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetRuleName(v string) *UpdateBackupPlanShrinkRequestRule {
	s.RuleName = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetSchedule(v string) *UpdateBackupPlanShrinkRequestRule {
	s.Schedule = &v
	return s
}

type UpdateBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanResponseBody) SetCode(v string) *UpdateBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBackupPlanResponseBody) SetMessage(v string) *UpdateBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBackupPlanResponseBody) SetRequestId(v string) *UpdateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBackupPlanResponseBody) SetSuccess(v bool) *UpdateBackupPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateBackupPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanResponse) SetHeaders(v map[string]*string) *UpdateBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateBackupPlanResponse) SetStatusCode(v int32) *UpdateBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBackupPlanResponse) SetBody(v *UpdateBackupPlanResponseBody) *UpdateBackupPlanResponse {
	s.Body = v
	return s
}

type UpdateClientSettingsRequest struct {
	ClientId         *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DataNetworkType  *string `json:"DataNetworkType,omitempty" xml:"DataNetworkType,omitempty"`
	DataProxySetting *string `json:"DataProxySetting,omitempty" xml:"DataProxySetting,omitempty"`
	MaxCpuCore       *int32  `json:"MaxCpuCore,omitempty" xml:"MaxCpuCore,omitempty"`
	MaxWorker        *int32  `json:"MaxWorker,omitempty" xml:"MaxWorker,omitempty"`
	ProxyHost        *string `json:"ProxyHost,omitempty" xml:"ProxyHost,omitempty"`
	ProxyPassword    *string `json:"ProxyPassword,omitempty" xml:"ProxyPassword,omitempty"`
	ProxyPort        *int32  `json:"ProxyPort,omitempty" xml:"ProxyPort,omitempty"`
	ProxyUser        *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	UseHttps         *bool   `json:"UseHttps,omitempty" xml:"UseHttps,omitempty"`
	VaultId          *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateClientSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientSettingsRequest) SetClientId(v string) *UpdateClientSettingsRequest {
	s.ClientId = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetDataNetworkType(v string) *UpdateClientSettingsRequest {
	s.DataNetworkType = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetDataProxySetting(v string) *UpdateClientSettingsRequest {
	s.DataProxySetting = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetMaxCpuCore(v int32) *UpdateClientSettingsRequest {
	s.MaxCpuCore = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetMaxWorker(v int32) *UpdateClientSettingsRequest {
	s.MaxWorker = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetProxyHost(v string) *UpdateClientSettingsRequest {
	s.ProxyHost = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetProxyPassword(v string) *UpdateClientSettingsRequest {
	s.ProxyPassword = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetProxyPort(v int32) *UpdateClientSettingsRequest {
	s.ProxyPort = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetProxyUser(v string) *UpdateClientSettingsRequest {
	s.ProxyUser = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetResourceGroupId(v string) *UpdateClientSettingsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetUseHttps(v bool) *UpdateClientSettingsRequest {
	s.UseHttps = &v
	return s
}

func (s *UpdateClientSettingsRequest) SetVaultId(v string) *UpdateClientSettingsRequest {
	s.VaultId = &v
	return s
}

type UpdateClientSettingsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClientSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientSettingsResponseBody) SetCode(v string) *UpdateClientSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateClientSettingsResponseBody) SetMessage(v string) *UpdateClientSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClientSettingsResponseBody) SetRequestId(v string) *UpdateClientSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClientSettingsResponseBody) SetSuccess(v bool) *UpdateClientSettingsResponseBody {
	s.Success = &v
	return s
}

type UpdateClientSettingsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateClientSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClientSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientSettingsResponse) SetHeaders(v map[string]*string) *UpdateClientSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientSettingsResponse) SetStatusCode(v int32) *UpdateClientSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientSettingsResponse) SetBody(v *UpdateClientSettingsResponseBody) *UpdateClientSettingsResponse {
	s.Body = v
	return s
}

type UpdateHanaBackupPlanRequest struct {
	BackupPrefix    *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PlanId          *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName        *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Schedule        *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateHanaBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupPlanRequest) SetBackupPrefix(v string) *UpdateHanaBackupPlanRequest {
	s.BackupPrefix = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetClusterId(v string) *UpdateHanaBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetPlanId(v string) *UpdateHanaBackupPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetPlanName(v string) *UpdateHanaBackupPlanRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetResourceGroupId(v string) *UpdateHanaBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetSchedule(v string) *UpdateHanaBackupPlanRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateHanaBackupPlanRequest) SetVaultId(v string) *UpdateHanaBackupPlanRequest {
	s.VaultId = &v
	return s
}

type UpdateHanaBackupPlanResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHanaBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupPlanResponseBody) SetCode(v string) *UpdateHanaBackupPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHanaBackupPlanResponseBody) SetMessage(v string) *UpdateHanaBackupPlanResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHanaBackupPlanResponseBody) SetRequestId(v string) *UpdateHanaBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHanaBackupPlanResponseBody) SetSuccess(v bool) *UpdateHanaBackupPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateHanaBackupPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHanaBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupPlanResponse) SetHeaders(v map[string]*string) *UpdateHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateHanaBackupPlanResponse) SetStatusCode(v int32) *UpdateHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHanaBackupPlanResponse) SetBody(v *UpdateHanaBackupPlanResponseBody) *UpdateHanaBackupPlanResponse {
	s.Body = v
	return s
}

type UpdateHanaBackupSettingRequest struct {
	CatalogBackupParameterFile *string `json:"CatalogBackupParameterFile,omitempty" xml:"CatalogBackupParameterFile,omitempty"`
	CatalogBackupUsingBackint  *bool   `json:"CatalogBackupUsingBackint,omitempty" xml:"CatalogBackupUsingBackint,omitempty"`
	ClusterId                  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataBackupParameterFile    *string `json:"DataBackupParameterFile,omitempty" xml:"DataBackupParameterFile,omitempty"`
	DatabaseName               *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EnableAutoLogBackup        *bool   `json:"EnableAutoLogBackup,omitempty" xml:"EnableAutoLogBackup,omitempty"`
	LogBackupParameterFile     *string `json:"LogBackupParameterFile,omitempty" xml:"LogBackupParameterFile,omitempty"`
	LogBackupTimeout           *int64  `json:"LogBackupTimeout,omitempty" xml:"LogBackupTimeout,omitempty"`
	LogBackupUsingBackint      *bool   `json:"LogBackupUsingBackint,omitempty" xml:"LogBackupUsingBackint,omitempty"`
	VaultId                    *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateHanaBackupSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaBackupSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupSettingRequest) SetCatalogBackupParameterFile(v string) *UpdateHanaBackupSettingRequest {
	s.CatalogBackupParameterFile = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetCatalogBackupUsingBackint(v bool) *UpdateHanaBackupSettingRequest {
	s.CatalogBackupUsingBackint = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetClusterId(v string) *UpdateHanaBackupSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetDataBackupParameterFile(v string) *UpdateHanaBackupSettingRequest {
	s.DataBackupParameterFile = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetDatabaseName(v string) *UpdateHanaBackupSettingRequest {
	s.DatabaseName = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetEnableAutoLogBackup(v bool) *UpdateHanaBackupSettingRequest {
	s.EnableAutoLogBackup = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetLogBackupParameterFile(v string) *UpdateHanaBackupSettingRequest {
	s.LogBackupParameterFile = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetLogBackupTimeout(v int64) *UpdateHanaBackupSettingRequest {
	s.LogBackupTimeout = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetLogBackupUsingBackint(v bool) *UpdateHanaBackupSettingRequest {
	s.LogBackupUsingBackint = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetVaultId(v string) *UpdateHanaBackupSettingRequest {
	s.VaultId = &v
	return s
}

type UpdateHanaBackupSettingResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHanaBackupSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaBackupSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupSettingResponseBody) SetCode(v string) *UpdateHanaBackupSettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHanaBackupSettingResponseBody) SetMessage(v string) *UpdateHanaBackupSettingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHanaBackupSettingResponseBody) SetRequestId(v string) *UpdateHanaBackupSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHanaBackupSettingResponseBody) SetSuccess(v bool) *UpdateHanaBackupSettingResponseBody {
	s.Success = &v
	return s
}

type UpdateHanaBackupSettingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateHanaBackupSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHanaBackupSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaBackupSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupSettingResponse) SetHeaders(v map[string]*string) *UpdateHanaBackupSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateHanaBackupSettingResponse) SetStatusCode(v int32) *UpdateHanaBackupSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHanaBackupSettingResponse) SetBody(v *UpdateHanaBackupSettingResponseBody) *UpdateHanaBackupSettingResponse {
	s.Body = v
	return s
}

type UpdateHanaInstanceRequest struct {
	AlertSetting        *string `json:"AlertSetting,omitempty" xml:"AlertSetting,omitempty"`
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HanaName            *string `json:"HanaName,omitempty" xml:"HanaName,omitempty"`
	Host                *string `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceNumber      *int32  `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	Password            *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	UseSsl              *bool   `json:"UseSsl,omitempty" xml:"UseSsl,omitempty"`
	UserName            *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ValidateCertificate *bool   `json:"ValidateCertificate,omitempty" xml:"ValidateCertificate,omitempty"`
	VaultId             *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateHanaInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateHanaInstanceRequest) SetAlertSetting(v string) *UpdateHanaInstanceRequest {
	s.AlertSetting = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetClusterId(v string) *UpdateHanaInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetHanaName(v string) *UpdateHanaInstanceRequest {
	s.HanaName = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetHost(v string) *UpdateHanaInstanceRequest {
	s.Host = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetInstanceNumber(v int32) *UpdateHanaInstanceRequest {
	s.InstanceNumber = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetPassword(v string) *UpdateHanaInstanceRequest {
	s.Password = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetResourceGroupId(v string) *UpdateHanaInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetUseSsl(v bool) *UpdateHanaInstanceRequest {
	s.UseSsl = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetUserName(v string) *UpdateHanaInstanceRequest {
	s.UserName = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetValidateCertificate(v bool) *UpdateHanaInstanceRequest {
	s.ValidateCertificate = &v
	return s
}

func (s *UpdateHanaInstanceRequest) SetVaultId(v string) *UpdateHanaInstanceRequest {
	s.VaultId = &v
	return s
}

type UpdateHanaInstanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHanaInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHanaInstanceResponseBody) SetCode(v string) *UpdateHanaInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHanaInstanceResponseBody) SetMessage(v string) *UpdateHanaInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHanaInstanceResponseBody) SetRequestId(v string) *UpdateHanaInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHanaInstanceResponseBody) SetSuccess(v bool) *UpdateHanaInstanceResponseBody {
	s.Success = &v
	return s
}

type UpdateHanaInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateHanaInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHanaInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateHanaInstanceResponse) SetHeaders(v map[string]*string) *UpdateHanaInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateHanaInstanceResponse) SetStatusCode(v int32) *UpdateHanaInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHanaInstanceResponse) SetBody(v *UpdateHanaInstanceResponseBody) *UpdateHanaInstanceResponse {
	s.Body = v
	return s
}

type UpdateHanaRetentionSettingRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName  *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Disabled      *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	RetentionDays *int64  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	Schedule      *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	VaultId       *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateHanaRetentionSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaRetentionSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateHanaRetentionSettingRequest) SetClusterId(v string) *UpdateHanaRetentionSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetDatabaseName(v string) *UpdateHanaRetentionSettingRequest {
	s.DatabaseName = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetDisabled(v bool) *UpdateHanaRetentionSettingRequest {
	s.Disabled = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetRetentionDays(v int64) *UpdateHanaRetentionSettingRequest {
	s.RetentionDays = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetSchedule(v string) *UpdateHanaRetentionSettingRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateHanaRetentionSettingRequest) SetVaultId(v string) *UpdateHanaRetentionSettingRequest {
	s.VaultId = &v
	return s
}

type UpdateHanaRetentionSettingResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHanaRetentionSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaRetentionSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHanaRetentionSettingResponseBody) SetCode(v string) *UpdateHanaRetentionSettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponseBody) SetMessage(v string) *UpdateHanaRetentionSettingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponseBody) SetRequestId(v string) *UpdateHanaRetentionSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponseBody) SetSuccess(v bool) *UpdateHanaRetentionSettingResponseBody {
	s.Success = &v
	return s
}

type UpdateHanaRetentionSettingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateHanaRetentionSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHanaRetentionSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHanaRetentionSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateHanaRetentionSettingResponse) SetHeaders(v map[string]*string) *UpdateHanaRetentionSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateHanaRetentionSettingResponse) SetStatusCode(v int32) *UpdateHanaRetentionSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponse) SetBody(v *UpdateHanaRetentionSettingResponseBody) *UpdateHanaRetentionSettingResponse {
	s.Body = v
	return s
}

type UpdateVaultRequest struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	VaultName       *string `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
}

func (s UpdateVaultRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVaultRequest) GoString() string {
	return s.String()
}

func (s *UpdateVaultRequest) SetDescription(v string) *UpdateVaultRequest {
	s.Description = &v
	return s
}

func (s *UpdateVaultRequest) SetResourceGroupId(v string) *UpdateVaultRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateVaultRequest) SetVaultId(v string) *UpdateVaultRequest {
	s.VaultId = &v
	return s
}

func (s *UpdateVaultRequest) SetVaultName(v string) *UpdateVaultRequest {
	s.VaultName = &v
	return s
}

type UpdateVaultResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateVaultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVaultResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVaultResponseBody) SetCode(v string) *UpdateVaultResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVaultResponseBody) SetMessage(v string) *UpdateVaultResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVaultResponseBody) SetRequestId(v string) *UpdateVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVaultResponseBody) SetSuccess(v bool) *UpdateVaultResponseBody {
	s.Success = &v
	return s
}

type UpdateVaultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVaultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVaultResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVaultResponse) GoString() string {
	return s.String()
}

func (s *UpdateVaultResponse) SetHeaders(v map[string]*string) *UpdateVaultResponse {
	s.Headers = v
	return s
}

func (s *UpdateVaultResponse) SetStatusCode(v int32) *UpdateVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVaultResponse) SetBody(v *UpdateVaultResponseBody) *UpdateVaultResponse {
	s.Body = v
	return s
}

type UpgradeBackupClientsRequest struct {
	ClientIds            map[string]interface{} `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	CrossAccountRoleName *string                `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string                `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64                 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	InstanceIds          map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s UpgradeBackupClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsRequest) SetClientIds(v map[string]interface{}) *UpgradeBackupClientsRequest {
	s.ClientIds = v
	return s
}

func (s *UpgradeBackupClientsRequest) SetCrossAccountRoleName(v string) *UpgradeBackupClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UpgradeBackupClientsRequest) SetCrossAccountType(v string) *UpgradeBackupClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *UpgradeBackupClientsRequest) SetCrossAccountUserId(v int64) *UpgradeBackupClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *UpgradeBackupClientsRequest) SetInstanceIds(v map[string]interface{}) *UpgradeBackupClientsRequest {
	s.InstanceIds = v
	return s
}

type UpgradeBackupClientsShrinkRequest struct {
	ClientIdsShrink      *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	InstanceIdsShrink    *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s UpgradeBackupClientsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeBackupClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsShrinkRequest) SetClientIdsShrink(v string) *UpgradeBackupClientsShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) SetCrossAccountRoleName(v string) *UpgradeBackupClientsShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) SetCrossAccountType(v string) *UpgradeBackupClientsShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) SetCrossAccountUserId(v int64) *UpgradeBackupClientsShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) SetInstanceIdsShrink(v string) *UpgradeBackupClientsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

type UpgradeBackupClientsResponseBody struct {
	Code             *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceStatuses []*UpgradeBackupClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
	Message          *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId           *string                                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeBackupClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsResponseBody) SetCode(v string) *UpgradeBackupClientsResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetInstanceStatuses(v []*UpgradeBackupClientsResponseBodyInstanceStatuses) *UpgradeBackupClientsResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetMessage(v string) *UpgradeBackupClientsResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetRequestId(v string) *UpgradeBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetSuccess(v bool) *UpgradeBackupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetTaskId(v string) *UpgradeBackupClientsResponseBody {
	s.TaskId = &v
	return s
}

type UpgradeBackupClientsResponseBodyInstanceStatuses struct {
	ErrorCode     *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ValidInstance *bool   `json:"ValidInstance,omitempty" xml:"ValidInstance,omitempty"`
}

func (s UpgradeBackupClientsResponseBodyInstanceStatuses) String() string {
	return tea.Prettify(s)
}

func (s UpgradeBackupClientsResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) SetErrorCode(v string) *UpgradeBackupClientsResponseBodyInstanceStatuses {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) SetInstanceId(v string) *UpgradeBackupClientsResponseBodyInstanceStatuses {
	s.InstanceId = &v
	return s
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) SetValidInstance(v bool) *UpgradeBackupClientsResponseBodyInstanceStatuses {
	s.ValidInstance = &v
	return s
}

type UpgradeBackupClientsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeBackupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeBackupClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeBackupClientsResponse) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsResponse) SetHeaders(v map[string]*string) *UpgradeBackupClientsResponse {
	s.Headers = v
	return s
}

func (s *UpgradeBackupClientsResponse) SetStatusCode(v int32) *UpgradeBackupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeBackupClientsResponse) SetBody(v *UpgradeBackupClientsResponseBody) *UpgradeBackupClientsResponse {
	s.Body = v
	return s
}

type UpgradeClientRequest struct {
	ClientId        *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	VaultId         *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpgradeClientRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClientRequest) SetClientId(v string) *UpgradeClientRequest {
	s.ClientId = &v
	return s
}

func (s *UpgradeClientRequest) SetResourceGroupId(v string) *UpgradeClientRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpgradeClientRequest) SetVaultId(v string) *UpgradeClientRequest {
	s.VaultId = &v
	return s
}

type UpgradeClientResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponseBody) SetCode(v string) *UpgradeClientResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeClientResponseBody) SetMessage(v string) *UpgradeClientResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeClientResponseBody) SetRequestId(v string) *UpgradeClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeClientResponseBody) SetSuccess(v bool) *UpgradeClientResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeClientResponseBody) SetTaskId(v string) *UpgradeClientResponseBody {
	s.TaskId = &v
	return s
}

type UpgradeClientResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeClientResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponse) SetHeaders(v map[string]*string) *UpgradeClientResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClientResponse) SetStatusCode(v int32) *UpgradeClientResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeClientResponse) SetBody(v *UpgradeClientResponseBody) *UpgradeClientResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("hbr.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("hbr.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("hbr.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("hbr.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("hbr.aliyuncs.com"),
		"cn-edge-1":                   tea.String("hbr.aliyuncs.com"),
		"cn-fujian":                   tea.String("hbr.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("hbr.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("hbr.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("hbr.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("hbr.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("hbr.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("hbr.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("hbr.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("hbr.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("hbr.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("hbr.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("hbr.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("hbr.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("hbr.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("hbr.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("hbr.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("hbr.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("hbr.aliyuncs.com"),
		"cn-wuhan":                    tea.String("hbr.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("hbr.aliyuncs.com"),
		"cn-yushanfang":               tea.String("hbr.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("hbr.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("hbr.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("hbr.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("hbr.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("hbr.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("hbr.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("hbr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddContainerClusterWithOptions(request *AddContainerClusterRequest, runtime *util.RuntimeOptions) (_result *AddContainerClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddContainerCluster"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddContainerClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddContainerCluster(request *AddContainerClusterRequest) (_result *AddContainerClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddContainerClusterResponse{}
	_body, _err := client.AddContainerClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelBackupJobWithOptions(request *CancelBackupJobRequest, runtime *util.RuntimeOptions) (_result *CancelBackupJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelBackupJob"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelBackupJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelBackupJob(request *CancelBackupJobRequest) (_result *CancelBackupJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelBackupJobResponse{}
	_body, _err := client.CancelBackupJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelRestoreJobWithOptions(request *CancelRestoreJobRequest, runtime *util.RuntimeOptions) (_result *CancelRestoreJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RestoreId)) {
		query["RestoreId"] = request.RestoreId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRestoreJob"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRestoreJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelRestoreJob(request *CancelRestoreJobRequest) (_result *CancelRestoreJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelRestoreJobResponse{}
	_body, _err := client.CancelRestoreJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupJobWithOptions(request *CreateBackupJobRequest, runtime *util.RuntimeOptions) (_result *CreateBackupJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupType)) {
		query["BackupType"] = request.BackupType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerClusterId)) {
		query["ContainerClusterId"] = request.ContainerClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerResources)) {
		query["ContainerResources"] = request.ContainerResources
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Exclude)) {
		query["Exclude"] = request.Exclude
	}

	if !tea.BoolValue(util.IsUnset(request.Include)) {
		query["Include"] = request.Include
	}

	if !tea.BoolValue(util.IsUnset(request.InitiatedByAck)) {
		query["InitiatedByAck"] = request.InitiatedByAck
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		query["JobName"] = request.JobName
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		query["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Retention)) {
		query["Retention"] = request.Retention
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SpeedLimit)) {
		query["SpeedLimit"] = request.SpeedLimit
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBackupJob"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBackupJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackupJob(request *CreateBackupJobRequest) (_result *CreateBackupJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupJobResponse{}
	_body, _err := client.CreateBackupJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupPlanWithOptions(tmpReq *CreateBackupPlanRequest, runtime *util.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateBackupPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Detail)) {
		request.DetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Detail, tea.String("Detail"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OtsDetail))) {
		request.OtsDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OtsDetail), tea.String("OtsDetail"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupType)) {
		query["BackupType"] = request.BackupType
	}

	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		query["CreateTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DetailShrink)) {
		query["Detail"] = request.DetailShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		query["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["Prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.Retention)) {
		query["Retention"] = request.Retention
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		query["Schedule"] = request.Schedule
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.UdmRegionId)) {
		query["UdmRegionId"] = request.UdmRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Exclude)) {
		body["Exclude"] = request.Exclude
	}

	if !tea.BoolValue(util.IsUnset(request.Include)) {
		body["Include"] = request.Include
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.OtsDetailShrink)) {
		body["OtsDetail"] = request.OtsDetailShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Rule)) {
		body["Rule"] = request.Rule
	}

	if !tea.BoolValue(util.IsUnset(request.SpeedLimit)) {
		body["SpeedLimit"] = request.SpeedLimit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackupPlan(request *CreateBackupPlanRequest) (_result *CreateBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CreateBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClientsWithOptions(request *CreateClientsRequest, runtime *util.RuntimeOptions) (_result *CreateClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertSetting)) {
		query["AlertSetting"] = request.AlertSetting
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UseHttps)) {
		query["UseHttps"] = request.UseHttps
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClients"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClients(request *CreateClientsRequest) (_result *CreateClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClientsResponse{}
	_body, _err := client.CreateClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHanaBackupPlanWithOptions(request *CreateHanaBackupPlanRequest, runtime *util.RuntimeOptions) (_result *CreateHanaBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPrefix)) {
		query["BackupPrefix"] = request.BackupPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.BackupType)) {
		query["BackupType"] = request.BackupType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		query["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		query["Schedule"] = request.Schedule
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHanaBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHanaBackupPlan(request *CreateHanaBackupPlanRequest) (_result *CreateHanaBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHanaBackupPlanResponse{}
	_body, _err := client.CreateHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHanaInstanceWithOptions(request *CreateHanaInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateHanaInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertSetting)) {
		query["AlertSetting"] = request.AlertSetting
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceId)) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.HanaName)) {
		query["HanaName"] = request.HanaName
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		query["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNumber)) {
		query["InstanceNumber"] = request.InstanceNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Sid)) {
		query["Sid"] = request.Sid
	}

	if !tea.BoolValue(util.IsUnset(request.UseSsl)) {
		query["UseSsl"] = request.UseSsl
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.ValidateCertificate)) {
		query["ValidateCertificate"] = request.ValidateCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHanaInstance"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHanaInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHanaInstance(request *CreateHanaInstanceRequest) (_result *CreateHanaInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHanaInstanceResponse{}
	_body, _err := client.CreateHanaInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHanaRestoreWithOptions(request *CreateHanaRestoreRequest, runtime *util.RuntimeOptions) (_result *CreateHanaRestoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupPrefix)) {
		query["BackupPrefix"] = request.BackupPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.CheckAccess)) {
		query["CheckAccess"] = request.CheckAccess
	}

	if !tea.BoolValue(util.IsUnset(request.ClearLog)) {
		query["ClearLog"] = request.ClearLog
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.LogPosition)) {
		query["LogPosition"] = request.LogPosition
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryPointInTime)) {
		query["RecoveryPointInTime"] = request.RecoveryPointInTime
	}

	if !tea.BoolValue(util.IsUnset(request.SidAdmin)) {
		query["SidAdmin"] = request.SidAdmin
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceClusterId)) {
		query["SourceClusterId"] = request.SourceClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemCopy)) {
		query["SystemCopy"] = request.SystemCopy
	}

	if !tea.BoolValue(util.IsUnset(request.UseCatalog)) {
		query["UseCatalog"] = request.UseCatalog
	}

	if !tea.BoolValue(util.IsUnset(request.UseDelta)) {
		query["UseDelta"] = request.UseDelta
	}

	if !tea.BoolValue(util.IsUnset(request.VolumeId)) {
		query["VolumeId"] = request.VolumeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHanaRestore"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHanaRestoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHanaRestore(request *CreateHanaRestoreRequest) (_result *CreateHanaRestoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHanaRestoreResponse{}
	_body, _err := client.CreateHanaRestoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateReplicationVaultWithOptions(request *CreateReplicationVaultRequest, runtime *util.RuntimeOptions) (_result *CreateReplicationVaultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RedundancyType)) {
		query["RedundancyType"] = request.RedundancyType
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicationSourceRegionId)) {
		query["ReplicationSourceRegionId"] = request.ReplicationSourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicationSourceVaultId)) {
		query["ReplicationSourceVaultId"] = request.ReplicationSourceVaultId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultName)) {
		query["VaultName"] = request.VaultName
	}

	if !tea.BoolValue(util.IsUnset(request.VaultRegionId)) {
		query["VaultRegionId"] = request.VaultRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultStorageClass)) {
		query["VaultStorageClass"] = request.VaultStorageClass
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateReplicationVault"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateReplicationVaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateReplicationVault(request *CreateReplicationVaultRequest) (_result *CreateReplicationVaultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateReplicationVaultResponse{}
	_body, _err := client.CreateReplicationVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRestoreJobWithOptions(tmpReq *CreateRestoreJobRequest, runtime *util.RuntimeOptions) (_result *CreateRestoreJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRestoreJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OtsDetail))) {
		request.OtsDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OtsDetail), tea.String("OtsDetail"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UdmDetail)) {
		request.UdmDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UdmDetail, tea.String("UdmDetail"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InitiatedByAck)) {
		query["InitiatedByAck"] = request.InitiatedByAck
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreType)) {
		query["RestoreType"] = request.RestoreType
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotHash)) {
		query["SnapshotHash"] = request.SnapshotHash
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetBucket)) {
		query["TargetBucket"] = request.TargetBucket
	}

	if !tea.BoolValue(util.IsUnset(request.TargetContainer)) {
		query["TargetContainer"] = request.TargetContainer
	}

	if !tea.BoolValue(util.IsUnset(request.TargetContainerClusterId)) {
		query["TargetContainerClusterId"] = request.TargetContainerClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCreateTime)) {
		query["TargetCreateTime"] = request.TargetCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFileSystemId)) {
		query["TargetFileSystemId"] = request.TargetFileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInstanceName)) {
		query["TargetInstanceName"] = request.TargetInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetPrefix)) {
		query["TargetPrefix"] = request.TargetPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTableName)) {
		query["TargetTableName"] = request.TargetTableName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTime)) {
		query["TargetTime"] = request.TargetTime
	}

	if !tea.BoolValue(util.IsUnset(request.UdmDetailShrink)) {
		query["UdmDetail"] = request.UdmDetailShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Exclude)) {
		body["Exclude"] = request.Exclude
	}

	if !tea.BoolValue(util.IsUnset(request.Include)) {
		body["Include"] = request.Include
	}

	if !tea.BoolValue(util.IsUnset(request.OtsDetailShrink)) {
		body["OtsDetail"] = request.OtsDetailShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInstanceId)) {
		body["TargetInstanceId"] = request.TargetInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetPath)) {
		body["TargetPath"] = request.TargetPath
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRestoreJob"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRestoreJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRestoreJob(request *CreateRestoreJobRequest) (_result *CreateRestoreJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRestoreJobResponse{}
	_body, _err := client.CreateRestoreJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTempFileUploadUrlWithOptions(request *CreateTempFileUploadUrlRequest, runtime *util.RuntimeOptions) (_result *CreateTempFileUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTempFileUploadUrl"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTempFileUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTempFileUploadUrl(request *CreateTempFileUploadUrlRequest) (_result *CreateTempFileUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTempFileUploadUrlResponse{}
	_body, _err := client.CreateTempFileUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVaultWithOptions(request *CreateVaultRequest, runtime *util.RuntimeOptions) (_result *CreateVaultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.CompressionAlgorithm)) {
		query["CompressionAlgorithm"] = request.CompressionAlgorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.KmsKeyId)) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.RedundancyType)) {
		query["RedundancyType"] = request.RedundancyType
	}

	if !tea.BoolValue(util.IsUnset(request.VaultAccessKeyId)) {
		query["VaultAccessKeyId"] = request.VaultAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultAccessKeySecret)) {
		query["VaultAccessKeySecret"] = request.VaultAccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.VaultName)) {
		query["VaultName"] = request.VaultName
	}

	if !tea.BoolValue(util.IsUnset(request.VaultRegionId)) {
		query["VaultRegionId"] = request.VaultRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultStorageClass)) {
		query["VaultStorageClass"] = request.VaultStorageClass
	}

	if !tea.BoolValue(util.IsUnset(request.VaultType)) {
		query["VaultType"] = request.VaultType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVault"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVault(request *CreateVaultRequest) (_result *CreateVaultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVaultResponse{}
	_body, _err := client.CreateVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBackupClientWithOptions(request *DeleteBackupClientRequest, runtime *util.RuntimeOptions) (_result *DeleteBackupClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBackupClient"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBackupClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBackupClient(request *DeleteBackupClientRequest) (_result *DeleteBackupClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBackupClientResponse{}
	_body, _err := client.DeleteBackupClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBackupClientResourceWithOptions(tmpReq *DeleteBackupClientResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteBackupClientResourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteBackupClientResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClientIds)) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, tea.String("ClientIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIdsShrink)) {
		query["ClientIds"] = request.ClientIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBackupClientResource"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBackupClientResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBackupClientResource(request *DeleteBackupClientResourceRequest) (_result *DeleteBackupClientResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBackupClientResourceResponse{}
	_body, _err := client.DeleteBackupClientResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBackupPlanWithOptions(request *DeleteBackupPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBackupPlan(request *DeleteBackupPlanRequest) (_result *DeleteBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBackupPlanResponse{}
	_body, _err := client.DeleteBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHanaBackupPlanWithOptions(request *DeleteHanaBackupPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteHanaBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHanaBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHanaBackupPlan(request *DeleteHanaBackupPlanRequest) (_result *DeleteHanaBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHanaBackupPlanResponse{}
	_body, _err := client.DeleteHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHanaInstanceWithOptions(request *DeleteHanaInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteHanaInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Sid)) {
		query["Sid"] = request.Sid
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHanaInstance"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHanaInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHanaInstance(request *DeleteHanaInstanceRequest) (_result *DeleteHanaInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHanaInstanceResponse{}
	_body, _err := client.DeleteHanaInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshot"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVaultWithOptions(request *DeleteVaultRequest, runtime *util.RuntimeOptions) (_result *DeleteVaultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVault"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVault(request *DeleteVaultRequest) (_result *DeleteVaultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVaultResponse{}
	_body, _err := client.DeleteVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupClientsWithOptions(tmpReq *DescribeBackupClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupClientsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeBackupClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClientIds)) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, tea.String("ClientIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIdsShrink)) {
		body["ClientIds"] = request.ClientIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		body["InstanceIds"] = request.InstanceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupClients"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupClients(request *DescribeBackupClientsRequest) (_result *DescribeBackupClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupClientsResponse{}
	_body, _err := client.DescribeBackupClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupJobs2WithOptions(request *DescribeBackupJobs2Request, runtime *util.RuntimeOptions) (_result *DescribeBackupJobs2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirection)) {
		query["SortDirection"] = request.SortDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupJobs2"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupJobs2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupJobs2(request *DescribeBackupJobs2Request) (_result *DescribeBackupJobs2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupJobs2Response{}
	_body, _err := client.DescribeBackupJobs2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPlansWithOptions(request *DescribeBackupPlansRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPlans"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPlans(request *DescribeBackupPlansRequest) (_result *DescribeBackupPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPlansResponse{}
	_body, _err := client.DescribeBackupPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClientsWithOptions(request *DescribeClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClients"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClients(request *DescribeClientsRequest) (_result *DescribeClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientsResponse{}
	_body, _err := client.DescribeClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerClusterWithOptions(request *DescribeContainerClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerCluster"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerCluster(request *DescribeContainerClusterRequest) (_result *DescribeContainerClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerClusterResponse{}
	_body, _err := client.DescribeContainerClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHanaBackupPlansWithOptions(request *DescribeHanaBackupPlansRequest, runtime *util.RuntimeOptions) (_result *DescribeHanaBackupPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHanaBackupPlans"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHanaBackupPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHanaBackupPlans(request *DescribeHanaBackupPlansRequest) (_result *DescribeHanaBackupPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHanaBackupPlansResponse{}
	_body, _err := client.DescribeHanaBackupPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHanaBackupSettingWithOptions(request *DescribeHanaBackupSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeHanaBackupSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHanaBackupSetting"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHanaBackupSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHanaBackupSetting(request *DescribeHanaBackupSettingRequest) (_result *DescribeHanaBackupSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHanaBackupSettingResponse{}
	_body, _err := client.DescribeHanaBackupSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHanaBackupsAsyncWithOptions(request *DescribeHanaBackupsAsyncRequest, runtime *util.RuntimeOptions) (_result *DescribeHanaBackupsAsyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDifferential)) {
		query["IncludeDifferential"] = request.IncludeDifferential
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeIncremental)) {
		query["IncludeIncremental"] = request.IncludeIncremental
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeLog)) {
		query["IncludeLog"] = request.IncludeLog
	}

	if !tea.BoolValue(util.IsUnset(request.LogPosition)) {
		query["LogPosition"] = request.LogPosition
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RecoveryPointInTime)) {
		query["RecoveryPointInTime"] = request.RecoveryPointInTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceClusterId)) {
		query["SourceClusterId"] = request.SourceClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemCopy)) {
		query["SystemCopy"] = request.SystemCopy
	}

	if !tea.BoolValue(util.IsUnset(request.UseBackint)) {
		query["UseBackint"] = request.UseBackint
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	if !tea.BoolValue(util.IsUnset(request.VolumeId)) {
		query["VolumeId"] = request.VolumeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHanaBackupsAsync"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHanaBackupsAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHanaBackupsAsync(request *DescribeHanaBackupsAsyncRequest) (_result *DescribeHanaBackupsAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHanaBackupsAsyncResponse{}
	_body, _err := client.DescribeHanaBackupsAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHanaDatabasesWithOptions(request *DescribeHanaDatabasesRequest, runtime *util.RuntimeOptions) (_result *DescribeHanaDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHanaDatabases"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHanaDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHanaDatabases(request *DescribeHanaDatabasesRequest) (_result *DescribeHanaDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHanaDatabasesResponse{}
	_body, _err := client.DescribeHanaDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHanaInstancesWithOptions(request *DescribeHanaInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeHanaInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHanaInstances"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHanaInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHanaInstances(request *DescribeHanaInstancesRequest) (_result *DescribeHanaInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHanaInstancesResponse{}
	_body, _err := client.DescribeHanaInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHanaRestoresWithOptions(request *DescribeHanaRestoresRequest, runtime *util.RuntimeOptions) (_result *DescribeHanaRestoresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreId)) {
		query["RestoreId"] = request.RestoreId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreStatus)) {
		query["RestoreStatus"] = request.RestoreStatus
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHanaRestores"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHanaRestoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHanaRestores(request *DescribeHanaRestoresRequest) (_result *DescribeHanaRestoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHanaRestoresResponse{}
	_body, _err := client.DescribeHanaRestoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHanaRetentionSettingWithOptions(request *DescribeHanaRetentionSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeHanaRetentionSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHanaRetentionSetting"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHanaRetentionSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHanaRetentionSetting(request *DescribeHanaRetentionSettingRequest) (_result *DescribeHanaRetentionSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHanaRetentionSettingResponse{}
	_body, _err := client.DescribeHanaRetentionSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOtsTableSnapshotsWithOptions(request *DescribeOtsTableSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeOtsTableSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OtsInstances)) {
		bodyFlat["OtsInstances"] = request.OtsInstances
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		bodyFlat["SnapshotIds"] = request.SnapshotIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOtsTableSnapshots"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOtsTableSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOtsTableSnapshots(request *DescribeOtsTableSnapshotsRequest) (_result *DescribeOtsTableSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOtsTableSnapshotsResponse{}
	_body, _err := client.DescribeOtsTableSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecoverableOtsInstancesWithOptions(request *DescribeRecoverableOtsInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeRecoverableOtsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecoverableOtsInstances"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecoverableOtsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecoverableOtsInstances(request *DescribeRecoverableOtsInstancesRequest) (_result *DescribeRecoverableOtsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecoverableOtsInstancesResponse{}
	_body, _err := client.DescribeRecoverableOtsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedVaultCount)) {
		query["NeedVaultCount"] = request.NeedVaultCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreJobs2WithOptions(request *DescribeRestoreJobs2Request, runtime *util.RuntimeOptions) (_result *DescribeRestoreJobs2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreType)) {
		query["RestoreType"] = request.RestoreType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRestoreJobs2"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRestoreJobs2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreJobs2(request *DescribeRestoreJobs2Request) (_result *DescribeRestoreJobs2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreJobs2Response{}
	_body, _err := client.DescribeRestoreJobs2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTaskWithOptions(request *DescribeTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTask"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTask(request *DescribeTaskRequest) (_result *DescribeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskResponse{}
	_body, _err := client.DescribeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUdmSnapshotsWithOptions(tmpReq *DescribeUdmSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeUdmSnapshotsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeUdmSnapshotsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SnapshotIds)) {
		request.SnapshotIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SnapshotIds, tea.String("SnapshotIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UdmRegionId)) {
		query["UdmRegionId"] = request.UdmRegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SnapshotIdsShrink)) {
		body["SnapshotIds"] = request.SnapshotIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUdmSnapshots"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUdmSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUdmSnapshots(request *DescribeUdmSnapshotsRequest) (_result *DescribeUdmSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUdmSnapshotsResponse{}
	_body, _err := client.DescribeUdmSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVaultReplicationRegionsWithOptions(request *DescribeVaultReplicationRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeVaultReplicationRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVaultReplicationRegions"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVaultReplicationRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVaultReplicationRegions(request *DescribeVaultReplicationRegionsRequest) (_result *DescribeVaultReplicationRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVaultReplicationRegionsResponse{}
	_body, _err := client.DescribeVaultReplicationRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVaultsWithOptions(request *DescribeVaultsRequest, runtime *util.RuntimeOptions) (_result *DescribeVaultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultRegionId)) {
		query["VaultRegionId"] = request.VaultRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultType)) {
		query["VaultType"] = request.VaultType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVaults"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVaultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVaults(request *DescribeVaultsRequest) (_result *DescribeVaultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVaultsResponse{}
	_body, _err := client.DescribeVaultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachNasFileSystemWithOptions(request *DetachNasFileSystemRequest, runtime *util.RuntimeOptions) (_result *DetachNasFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		query["CreateTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachNasFileSystem"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachNasFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachNasFileSystem(request *DetachNasFileSystemRequest) (_result *DetachNasFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachNasFileSystemResponse{}
	_body, _err := client.DetachNasFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableBackupPlanWithOptions(request *DisableBackupPlanRequest, runtime *util.RuntimeOptions) (_result *DisableBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableBackupPlan(request *DisableBackupPlanRequest) (_result *DisableBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableBackupPlanResponse{}
	_body, _err := client.DisableBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableHanaBackupPlanWithOptions(request *DisableHanaBackupPlanRequest, runtime *util.RuntimeOptions) (_result *DisableHanaBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableHanaBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableHanaBackupPlan(request *DisableHanaBackupPlanRequest) (_result *DisableHanaBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableHanaBackupPlanResponse{}
	_body, _err := client.DisableHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableBackupPlanWithOptions(request *EnableBackupPlanRequest, runtime *util.RuntimeOptions) (_result *EnableBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableBackupPlan(request *EnableBackupPlanRequest) (_result *EnableBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableBackupPlanResponse{}
	_body, _err := client.EnableBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableHanaBackupPlanWithOptions(request *EnableHanaBackupPlanRequest, runtime *util.RuntimeOptions) (_result *EnableHanaBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableHanaBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableHanaBackupPlan(request *EnableHanaBackupPlanRequest) (_result *EnableHanaBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableHanaBackupPlanResponse{}
	_body, _err := client.EnableHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteBackupPlanWithOptions(request *ExecuteBackupPlanRequest, runtime *util.RuntimeOptions) (_result *ExecuteBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteBackupPlan(request *ExecuteBackupPlanRequest) (_result *ExecuteBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteBackupPlanResponse{}
	_body, _err := client.ExecuteBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateRamPolicyWithOptions(request *GenerateRamPolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateRamPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.RequireBasePolicy)) {
		query["RequireBasePolicy"] = request.RequireBasePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateRamPolicy"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateRamPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateRamPolicy(request *GenerateRamPolicyRequest) (_result *GenerateRamPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateRamPolicyResponse{}
	_body, _err := client.GenerateRamPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTempFileDownloadLinkWithOptions(request *GetTempFileDownloadLinkRequest, runtime *util.RuntimeOptions) (_result *GetTempFileDownloadLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TempFileKey)) {
		query["TempFileKey"] = request.TempFileKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTempFileDownloadLink"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTempFileDownloadLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTempFileDownloadLink(request *GetTempFileDownloadLinkRequest) (_result *GetTempFileDownloadLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTempFileDownloadLinkResponse{}
	_body, _err := client.GetTempFileDownloadLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallBackupClientsWithOptions(tmpReq *InstallBackupClientsRequest, runtime *util.RuntimeOptions) (_result *InstallBackupClientsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InstallBackupClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallBackupClients"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallBackupClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallBackupClients(request *InstallBackupClientsRequest) (_result *InstallBackupClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallBackupClientsResponse{}
	_body, _err := client.InstallBackupClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenHbrServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenHbrServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenHbrService"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenHbrServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenHbrService() (_result *OpenHbrServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenHbrServiceResponse{}
	_body, _err := client.OpenHbrServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchHistoricalSnapshotsWithOptions(tmpReq *SearchHistoricalSnapshotsRequest, runtime *util.RuntimeOptions) (_result *SearchHistoricalSnapshotsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchHistoricalSnapshotsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Query)) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, tea.String("Query"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.QueryShrink)) {
		query["Query"] = request.QueryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchHistoricalSnapshots"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchHistoricalSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchHistoricalSnapshots(request *SearchHistoricalSnapshotsRequest) (_result *SearchHistoricalSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchHistoricalSnapshotsResponse{}
	_body, _err := client.SearchHistoricalSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartHanaDatabaseAsyncWithOptions(request *StartHanaDatabaseAsyncRequest, runtime *util.RuntimeOptions) (_result *StartHanaDatabaseAsyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartHanaDatabaseAsync"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartHanaDatabaseAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartHanaDatabaseAsync(request *StartHanaDatabaseAsyncRequest) (_result *StartHanaDatabaseAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartHanaDatabaseAsyncResponse{}
	_body, _err := client.StartHanaDatabaseAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopHanaDatabaseAsyncWithOptions(request *StopHanaDatabaseAsyncRequest, runtime *util.RuntimeOptions) (_result *StopHanaDatabaseAsyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopHanaDatabaseAsync"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopHanaDatabaseAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopHanaDatabaseAsync(request *StopHanaDatabaseAsyncRequest) (_result *StopHanaDatabaseAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopHanaDatabaseAsyncResponse{}
	_body, _err := client.StopHanaDatabaseAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallBackupClientsWithOptions(tmpReq *UninstallBackupClientsRequest, runtime *util.RuntimeOptions) (_result *UninstallBackupClientsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UninstallBackupClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClientIds)) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, tea.String("ClientIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIdsShrink)) {
		query["ClientIds"] = request.ClientIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallBackupClients"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallBackupClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallBackupClients(request *UninstallBackupClientsRequest) (_result *UninstallBackupClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallBackupClientsResponse{}
	_body, _err := client.UninstallBackupClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallClientWithOptions(request *UninstallClientRequest, runtime *util.RuntimeOptions) (_result *UninstallClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallClient"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallClient(request *UninstallClientRequest) (_result *UninstallClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallClientResponse{}
	_body, _err := client.UninstallClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBackupPlanWithOptions(tmpReq *UpdateBackupPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateBackupPlanResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateBackupPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Detail)) {
		request.DetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Detail, tea.String("Detail"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OtsDetail))) {
		request.OtsDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OtsDetail), tea.String("OtsDetail"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DetailShrink)) {
		query["Detail"] = request.DetailShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		query["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["Prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.Retention)) {
		query["Retention"] = request.Retention
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		query["Schedule"] = request.Schedule
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SpeedLimit)) {
		query["SpeedLimit"] = request.SpeedLimit
	}

	if !tea.BoolValue(util.IsUnset(request.UpdatePaths)) {
		query["UpdatePaths"] = request.UpdatePaths
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Exclude)) {
		body["Exclude"] = request.Exclude
	}

	if !tea.BoolValue(util.IsUnset(request.Include)) {
		body["Include"] = request.Include
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.OtsDetailShrink)) {
		body["OtsDetail"] = request.OtsDetailShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Rule)) {
		body["Rule"] = request.Rule
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBackupPlan(request *UpdateBackupPlanRequest) (_result *UpdateBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBackupPlanResponse{}
	_body, _err := client.UpdateBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClientSettingsWithOptions(request *UpdateClientSettingsRequest, runtime *util.RuntimeOptions) (_result *UpdateClientSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DataNetworkType)) {
		query["DataNetworkType"] = request.DataNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.DataProxySetting)) {
		query["DataProxySetting"] = request.DataProxySetting
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCpuCore)) {
		query["MaxCpuCore"] = request.MaxCpuCore
	}

	if !tea.BoolValue(util.IsUnset(request.MaxWorker)) {
		query["MaxWorker"] = request.MaxWorker
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyHost)) {
		query["ProxyHost"] = request.ProxyHost
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyPassword)) {
		query["ProxyPassword"] = request.ProxyPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyPort)) {
		query["ProxyPort"] = request.ProxyPort
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyUser)) {
		query["ProxyUser"] = request.ProxyUser
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UseHttps)) {
		query["UseHttps"] = request.UseHttps
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClientSettings"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClientSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateClientSettings(request *UpdateClientSettingsRequest) (_result *UpdateClientSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClientSettingsResponse{}
	_body, _err := client.UpdateClientSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHanaBackupPlanWithOptions(request *UpdateHanaBackupPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateHanaBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPrefix)) {
		query["BackupPrefix"] = request.BackupPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		query["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		query["Schedule"] = request.Schedule
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHanaBackupPlan"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHanaBackupPlan(request *UpdateHanaBackupPlanRequest) (_result *UpdateHanaBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHanaBackupPlanResponse{}
	_body, _err := client.UpdateHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHanaBackupSettingWithOptions(request *UpdateHanaBackupSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateHanaBackupSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogBackupParameterFile)) {
		query["CatalogBackupParameterFile"] = request.CatalogBackupParameterFile
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogBackupUsingBackint)) {
		query["CatalogBackupUsingBackint"] = request.CatalogBackupUsingBackint
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DataBackupParameterFile)) {
		query["DataBackupParameterFile"] = request.DataBackupParameterFile
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAutoLogBackup)) {
		query["EnableAutoLogBackup"] = request.EnableAutoLogBackup
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupParameterFile)) {
		query["LogBackupParameterFile"] = request.LogBackupParameterFile
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupTimeout)) {
		query["LogBackupTimeout"] = request.LogBackupTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupUsingBackint)) {
		query["LogBackupUsingBackint"] = request.LogBackupUsingBackint
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHanaBackupSetting"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHanaBackupSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHanaBackupSetting(request *UpdateHanaBackupSettingRequest) (_result *UpdateHanaBackupSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHanaBackupSettingResponse{}
	_body, _err := client.UpdateHanaBackupSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHanaInstanceWithOptions(request *UpdateHanaInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateHanaInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertSetting)) {
		query["AlertSetting"] = request.AlertSetting
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.HanaName)) {
		query["HanaName"] = request.HanaName
	}

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		query["Host"] = request.Host
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceNumber)) {
		query["InstanceNumber"] = request.InstanceNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UseSsl)) {
		query["UseSsl"] = request.UseSsl
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.ValidateCertificate)) {
		query["ValidateCertificate"] = request.ValidateCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHanaInstance"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHanaInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHanaInstance(request *UpdateHanaInstanceRequest) (_result *UpdateHanaInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHanaInstanceResponse{}
	_body, _err := client.UpdateHanaInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHanaRetentionSettingWithOptions(request *UpdateHanaRetentionSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateHanaRetentionSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Disabled)) {
		query["Disabled"] = request.Disabled
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.Schedule)) {
		query["Schedule"] = request.Schedule
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHanaRetentionSetting"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHanaRetentionSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHanaRetentionSetting(request *UpdateHanaRetentionSettingRequest) (_result *UpdateHanaRetentionSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHanaRetentionSettingResponse{}
	_body, _err := client.UpdateHanaRetentionSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVaultWithOptions(request *UpdateVaultRequest, runtime *util.RuntimeOptions) (_result *UpdateVaultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultName)) {
		query["VaultName"] = request.VaultName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVault"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVault(request *UpdateVaultRequest) (_result *UpdateVaultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVaultResponse{}
	_body, _err := client.UpdateVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeBackupClientsWithOptions(tmpReq *UpgradeBackupClientsRequest, runtime *util.RuntimeOptions) (_result *UpgradeBackupClientsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpgradeBackupClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClientIds)) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, tea.String("ClientIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIdsShrink)) {
		query["ClientIds"] = request.ClientIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountRoleName)) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountType)) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CrossAccountUserId)) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeBackupClients"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeBackupClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeBackupClients(request *UpgradeBackupClientsRequest) (_result *UpgradeBackupClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeBackupClientsResponse{}
	_body, _err := client.UpgradeBackupClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClientWithOptions(request *UpgradeClientRequest, runtime *util.RuntimeOptions) (_result *UpgradeClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VaultId)) {
		query["VaultId"] = request.VaultId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeClient"),
		Version:     tea.String("2017-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeClient(request *UpgradeClientRequest) (_result *UpgradeClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClientResponse{}
	_body, _err := client.UpgradeClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
