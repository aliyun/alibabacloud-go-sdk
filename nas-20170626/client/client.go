// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddClientToBlackListRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddClientToBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddClientToBlackListRequest) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListRequest) SetRegionId(v string) *AddClientToBlackListRequest {
	s.RegionId = &v
	return s
}

func (s *AddClientToBlackListRequest) SetFileSystemId(v string) *AddClientToBlackListRequest {
	s.FileSystemId = &v
	return s
}

func (s *AddClientToBlackListRequest) SetClientIP(v string) *AddClientToBlackListRequest {
	s.ClientIP = &v
	return s
}

func (s *AddClientToBlackListRequest) SetClientToken(v string) *AddClientToBlackListRequest {
	s.ClientToken = &v
	return s
}

type AddClientToBlackListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddClientToBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddClientToBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListResponseBody) SetRequestId(v string) *AddClientToBlackListResponseBody {
	s.RequestId = &v
	return s
}

type AddClientToBlackListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddClientToBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddClientToBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddClientToBlackListResponse) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListResponse) SetHeaders(v map[string]*string) *AddClientToBlackListResponse {
	s.Headers = v
	return s
}

func (s *AddClientToBlackListResponse) SetBody(v *AddClientToBlackListResponseBody) *AddClientToBlackListResponse {
	s.Body = v
	return s
}

type AddTagsRequest struct {
	FileSystemId *string              `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Tag          []*AddTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AddTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsRequest) GoString() string {
	return s.String()
}

func (s *AddTagsRequest) SetFileSystemId(v string) *AddTagsRequest {
	s.FileSystemId = &v
	return s
}

func (s *AddTagsRequest) SetTag(v []*AddTagsRequestTag) *AddTagsRequest {
	s.Tag = v
	return s
}

type AddTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s AddTagsRequestTag) GoString() string {
	return s.String()
}

func (s *AddTagsRequestTag) SetKey(v string) *AddTagsRequestTag {
	s.Key = &v
	return s
}

func (s *AddTagsRequestTag) SetValue(v string) *AddTagsRequestTag {
	s.Value = &v
	return s
}

type AddTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTagsResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsResponseBody) SetRequestId(v string) *AddTagsResponseBody {
	s.RequestId = &v
	return s
}

type AddTagsResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTagsResponse) GoString() string {
	return s.String()
}

func (s *AddTagsResponse) SetHeaders(v map[string]*string) *AddTagsResponse {
	s.Headers = v
	return s
}

func (s *AddTagsResponse) SetBody(v *AddTagsResponseBody) *AddTagsResponse {
	s.Body = v
	return s
}

type ApplyAutoSnapshotPolicyRequest struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	FileSystemIds        *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
}

func (s ApplyAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetFileSystemIds(v string) *ApplyAutoSnapshotPolicyRequest {
	s.FileSystemIds = &v
	return s
}

type ApplyAutoSnapshotPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyResponseBody) SetRequestId(v string) *ApplyAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ApplyAutoSnapshotPolicyResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *ApplyAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) SetBody(v *ApplyAutoSnapshotPolicyResponseBody) *ApplyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CancelAutoSnapshotPolicyRequest struct {
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
}

func (s CancelAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyRequest) SetFileSystemIds(v string) *CancelAutoSnapshotPolicyRequest {
	s.FileSystemIds = &v
	return s
}

type CancelAutoSnapshotPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CancelAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CancelAutoSnapshotPolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CancelAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CancelAutoSnapshotPolicyResponse) SetBody(v *CancelAutoSnapshotPolicyResponseBody) *CancelAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CancelDirQuotaRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UserType     *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelDirQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDirQuotaRequest) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaRequest) SetFileSystemId(v string) *CancelDirQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDirQuotaRequest) SetPath(v string) *CancelDirQuotaRequest {
	s.Path = &v
	return s
}

func (s *CancelDirQuotaRequest) SetUserType(v string) *CancelDirQuotaRequest {
	s.UserType = &v
	return s
}

func (s *CancelDirQuotaRequest) SetUserId(v string) *CancelDirQuotaRequest {
	s.UserId = &v
	return s
}

type CancelDirQuotaResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDirQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDirQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaResponseBody) SetSuccess(v bool) *CancelDirQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *CancelDirQuotaResponseBody) SetRequestId(v string) *CancelDirQuotaResponseBody {
	s.RequestId = &v
	return s
}

type CancelDirQuotaResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelDirQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelDirQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDirQuotaResponse) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaResponse) SetHeaders(v map[string]*string) *CancelDirQuotaResponse {
	s.Headers = v
	return s
}

func (s *CancelDirQuotaResponse) SetBody(v *CancelDirQuotaResponseBody) *CancelDirQuotaResponse {
	s.Body = v
	return s
}

type CancelLifecycleRetrieveJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelLifecycleRetrieveJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobRequest) SetJobId(v string) *CancelLifecycleRetrieveJobRequest {
	s.JobId = &v
	return s
}

type CancelLifecycleRetrieveJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelLifecycleRetrieveJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobResponseBody) SetRequestId(v string) *CancelLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

type CancelLifecycleRetrieveJobResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelLifecycleRetrieveJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *CancelLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *CancelLifecycleRetrieveJobResponse) SetBody(v *CancelLifecycleRetrieveJobResponseBody) *CancelLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

type CancelRecycleBinJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelRecycleBinJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRecycleBinJobRequest) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobRequest) SetJobId(v string) *CancelRecycleBinJobRequest {
	s.JobId = &v
	return s
}

type CancelRecycleBinJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRecycleBinJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRecycleBinJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobResponseBody) SetRequestId(v string) *CancelRecycleBinJobResponseBody {
	s.RequestId = &v
	return s
}

type CancelRecycleBinJobResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelRecycleBinJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelRecycleBinJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRecycleBinJobResponse) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobResponse) SetHeaders(v map[string]*string) *CancelRecycleBinJobResponse {
	s.Headers = v
	return s
}

func (s *CancelRecycleBinJobResponse) SetBody(v *CancelRecycleBinJobResponseBody) *CancelRecycleBinJobResponse {
	s.Body = v
	return s
}

type CreateAccessGroupRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	AccessGroupType *string `json:"AccessGroupType,omitempty" xml:"AccessGroupType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemType  *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s CreateAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupRequest) SetAccessGroupName(v string) *CreateAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessGroupRequest) SetAccessGroupType(v string) *CreateAccessGroupRequest {
	s.AccessGroupType = &v
	return s
}

func (s *CreateAccessGroupRequest) SetDescription(v string) *CreateAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessGroupRequest) SetFileSystemType(v string) *CreateAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

type CreateAccessGroupResponseBody struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponseBody) SetAccessGroupName(v string) *CreateAccessGroupResponseBody {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessGroupResponseBody) SetRequestId(v string) *CreateAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponse) SetHeaders(v map[string]*string) *CreateAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessGroupResponse) SetBody(v *CreateAccessGroupResponseBody) *CreateAccessGroupResponse {
	s.Body = v
	return s
}

type CreateAccessRuleRequest struct {
	AccessGroupName  *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	SourceCidrIp     *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	RWAccessType     *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	UserAccessType   *string `json:"UserAccessType,omitempty" xml:"UserAccessType,omitempty"`
	Priority         *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	FileSystemType   *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
}

func (s CreateAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleRequest) SetAccessGroupName(v string) *CreateAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessRuleRequest) SetSourceCidrIp(v string) *CreateAccessRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateAccessRuleRequest) SetRWAccessType(v string) *CreateAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetUserAccessType(v string) *CreateAccessRuleRequest {
	s.UserAccessType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetPriority(v int32) *CreateAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateAccessRuleRequest) SetFileSystemType(v string) *CreateAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetIpv6SourceCidrIp(v string) *CreateAccessRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

type CreateAccessRuleResponseBody struct {
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleResponseBody) SetAccessRuleId(v string) *CreateAccessRuleResponseBody {
	s.AccessRuleId = &v
	return s
}

func (s *CreateAccessRuleResponseBody) SetRequestId(v string) *CreateAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessRuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleResponse) SetHeaders(v map[string]*string) *CreateAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessRuleResponse) SetBody(v *CreateAccessRuleResponseBody) *CreateAccessRuleResponse {
	s.Body = v
	return s
}

type CreateAutoSnapshotPolicyRequest struct {
	RepeatWeekdays         *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	TimePoints             *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
	RetentionDays          *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	FileSystemType         *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetFileSystemType(v string) *CreateAutoSnapshotPolicyRequest {
	s.FileSystemType = &v
	return s
}

type CreateAutoSnapshotPolicyResponseBody struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetAutoSnapshotPolicyId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateAutoSnapshotPolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CreateAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) SetBody(v *CreateAutoSnapshotPolicyResponseBody) *CreateAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CreateFileSystemRequest struct {
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Duration       *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Capacity       *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Bandwidth      *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	StorageType    *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ProtocolType   *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	EncryptType    *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId      *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	KmsKeyId       *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s CreateFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequest) SetFileSystemType(v string) *CreateFileSystemRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateFileSystemRequest) SetChargeType(v string) *CreateFileSystemRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateFileSystemRequest) SetDuration(v int32) *CreateFileSystemRequest {
	s.Duration = &v
	return s
}

func (s *CreateFileSystemRequest) SetCapacity(v int64) *CreateFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *CreateFileSystemRequest) SetBandwidth(v int64) *CreateFileSystemRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageType(v string) *CreateFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateFileSystemRequest) SetZoneId(v string) *CreateFileSystemRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateFileSystemRequest) SetProtocolType(v string) *CreateFileSystemRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateFileSystemRequest) SetEncryptType(v int32) *CreateFileSystemRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateFileSystemRequest) SetSnapshotId(v string) *CreateFileSystemRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateFileSystemRequest) SetVpcId(v string) *CreateFileSystemRequest {
	s.VpcId = &v
	return s
}

func (s *CreateFileSystemRequest) SetVSwitchId(v string) *CreateFileSystemRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateFileSystemRequest) SetDescription(v string) *CreateFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateFileSystemRequest) SetClientToken(v string) *CreateFileSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFileSystemRequest) SetKmsKeyId(v string) *CreateFileSystemRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateFileSystemRequest) SetDryRun(v bool) *CreateFileSystemRequest {
	s.DryRun = &v
	return s
}

type CreateFileSystemResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s CreateFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponseBody) SetRequestId(v string) *CreateFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileSystemResponseBody) SetFileSystemId(v string) *CreateFileSystemResponseBody {
	s.FileSystemId = &v
	return s
}

type CreateFileSystemResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponse) SetHeaders(v map[string]*string) *CreateFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CreateFileSystemResponse) SetBody(v *CreateFileSystemResponseBody) *CreateFileSystemResponse {
	s.Body = v
	return s
}

type CreateLDAPConfigRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	URI          *string `json:"URI,omitempty" xml:"URI,omitempty"`
	BindDN       *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	SearchBase   *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
}

func (s CreateLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigRequest) SetFileSystemId(v string) *CreateLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetURI(v string) *CreateLDAPConfigRequest {
	s.URI = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetBindDN(v string) *CreateLDAPConfigRequest {
	s.BindDN = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetSearchBase(v string) *CreateLDAPConfigRequest {
	s.SearchBase = &v
	return s
}

type CreateLDAPConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigResponseBody) SetRequestId(v string) *CreateLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateLDAPConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLDAPConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigResponse) SetHeaders(v map[string]*string) *CreateLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLDAPConfigResponse) SetBody(v *CreateLDAPConfigResponseBody) *CreateLDAPConfigResponse {
	s.Body = v
	return s
}

type CreateLifecyclePolicyRequest struct {
	FileSystemId        *string   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	LifecyclePolicyName *string   `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	Path                *string   `json:"Path,omitempty" xml:"Path,omitempty"`
	LifecycleRuleName   *string   `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	StorageType         *string   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Paths               []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
}

func (s CreateLifecyclePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyRequest) SetFileSystemId(v string) *CreateLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *CreateLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPath(v string) *CreateLifecyclePolicyRequest {
	s.Path = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecycleRuleName(v string) *CreateLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetStorageType(v string) *CreateLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPaths(v []*string) *CreateLifecyclePolicyRequest {
	s.Paths = v
	return s
}

type CreateLifecyclePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLifecyclePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyResponseBody) SetRequestId(v string) *CreateLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLifecyclePolicyResponseBody) SetSuccess(v bool) *CreateLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

type CreateLifecyclePolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLifecyclePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyResponse) SetHeaders(v map[string]*string) *CreateLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateLifecyclePolicyResponse) SetBody(v *CreateLifecyclePolicyResponseBody) *CreateLifecyclePolicyResponse {
	s.Body = v
	return s
}

type CreateLifecycleRetrieveJobRequest struct {
	FileSystemId *string   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Paths        []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
}

func (s CreateLifecycleRetrieveJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobRequest) SetFileSystemId(v string) *CreateLifecycleRetrieveJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLifecycleRetrieveJobRequest) SetPaths(v []*string) *CreateLifecycleRetrieveJobRequest {
	s.Paths = v
	return s
}

type CreateLifecycleRetrieveJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateLifecycleRetrieveJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobResponseBody) SetRequestId(v string) *CreateLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLifecycleRetrieveJobResponseBody) SetJobId(v string) *CreateLifecycleRetrieveJobResponseBody {
	s.JobId = &v
	return s
}

type CreateLifecycleRetrieveJobResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLifecycleRetrieveJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *CreateLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLifecycleRetrieveJobResponse) SetBody(v *CreateLifecycleRetrieveJobResponseBody) *CreateLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

type CreateMountTargetRequest struct {
	FileSystemId    *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	EnableIpv6      *bool   `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	DryRun          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s CreateMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateMountTargetRequest) SetFileSystemId(v string) *CreateMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateMountTargetRequest) SetAccessGroupName(v string) *CreateMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateMountTargetRequest) SetNetworkType(v string) *CreateMountTargetRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateMountTargetRequest) SetVpcId(v string) *CreateMountTargetRequest {
	s.VpcId = &v
	return s
}

func (s *CreateMountTargetRequest) SetVSwitchId(v string) *CreateMountTargetRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMountTargetRequest) SetSecurityGroupId(v string) *CreateMountTargetRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMountTargetRequest) SetEnableIpv6(v bool) *CreateMountTargetRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateMountTargetRequest) SetDryRun(v bool) *CreateMountTargetRequest {
	s.DryRun = &v
	return s
}

type CreateMountTargetResponseBody struct {
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MountTargetDomain *string                                        `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	MountTargetExtra  *CreateMountTargetResponseBodyMountTargetExtra `json:"MountTargetExtra,omitempty" xml:"MountTargetExtra,omitempty" type:"Struct"`
}

func (s CreateMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponseBody) SetRequestId(v string) *CreateMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMountTargetResponseBody) SetMountTargetDomain(v string) *CreateMountTargetResponseBody {
	s.MountTargetDomain = &v
	return s
}

func (s *CreateMountTargetResponseBody) SetMountTargetExtra(v *CreateMountTargetResponseBodyMountTargetExtra) *CreateMountTargetResponseBody {
	s.MountTargetExtra = v
	return s
}

type CreateMountTargetResponseBodyMountTargetExtra struct {
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
}

func (s CreateMountTargetResponseBodyMountTargetExtra) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetResponseBodyMountTargetExtra) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponseBodyMountTargetExtra) SetDualStackMountTargetDomain(v string) *CreateMountTargetResponseBodyMountTargetExtra {
	s.DualStackMountTargetDomain = &v
	return s
}

type CreateMountTargetResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponse) SetHeaders(v map[string]*string) *CreateMountTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateMountTargetResponse) SetBody(v *CreateMountTargetResponseBody) *CreateMountTargetResponse {
	s.Body = v
	return s
}

type CreateRecycleBinDeleteJobRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateRecycleBinDeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinDeleteJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobRequest) SetFileSystemId(v string) *CreateRecycleBinDeleteJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) SetFileId(v string) *CreateRecycleBinDeleteJobRequest {
	s.FileId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) SetClientToken(v string) *CreateRecycleBinDeleteJobRequest {
	s.ClientToken = &v
	return s
}

type CreateRecycleBinDeleteJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateRecycleBinDeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinDeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobResponseBody) SetRequestId(v string) *CreateRecycleBinDeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobResponseBody) SetJobId(v string) *CreateRecycleBinDeleteJobResponseBody {
	s.JobId = &v
	return s
}

type CreateRecycleBinDeleteJobResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRecycleBinDeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRecycleBinDeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinDeleteJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobResponse) SetHeaders(v map[string]*string) *CreateRecycleBinDeleteJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRecycleBinDeleteJobResponse) SetBody(v *CreateRecycleBinDeleteJobResponseBody) *CreateRecycleBinDeleteJobResponse {
	s.Body = v
	return s
}

type CreateRecycleBinRestoreJobRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	TargetFileId *string `json:"TargetFileId,omitempty" xml:"TargetFileId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateRecycleBinRestoreJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobRequest) SetFileSystemId(v string) *CreateRecycleBinRestoreJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetFileId(v string) *CreateRecycleBinRestoreJobRequest {
	s.FileId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetTargetFileId(v string) *CreateRecycleBinRestoreJobRequest {
	s.TargetFileId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetClientToken(v string) *CreateRecycleBinRestoreJobRequest {
	s.ClientToken = &v
	return s
}

type CreateRecycleBinRestoreJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateRecycleBinRestoreJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobResponseBody) SetRequestId(v string) *CreateRecycleBinRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobResponseBody) SetJobId(v string) *CreateRecycleBinRestoreJobResponseBody {
	s.JobId = &v
	return s
}

type CreateRecycleBinRestoreJobResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRecycleBinRestoreJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRecycleBinRestoreJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinRestoreJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobResponse) SetHeaders(v map[string]*string) *CreateRecycleBinRestoreJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRecycleBinRestoreJobResponse) SetBody(v *CreateRecycleBinRestoreJobResponseBody) *CreateRecycleBinRestoreJobResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	SnapshotName  *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RetentionDays *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetFileSystemId(v string) *CreateSnapshotRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetRetentionDays(v int32) *CreateSnapshotRequest {
	s.RetentionDays = &v
	return s
}

type CreateSnapshotResponseBody struct {
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type DeleteAccessGroupRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	FileSystemType  *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DeleteAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupRequest) SetAccessGroupName(v string) *DeleteAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DeleteAccessGroupRequest) SetFileSystemType(v string) *DeleteAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

type DeleteAccessGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupResponseBody) SetRequestId(v string) *DeleteAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupResponse) SetHeaders(v map[string]*string) *DeleteAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessGroupResponse) SetBody(v *DeleteAccessGroupResponseBody) *DeleteAccessGroupResponse {
	s.Body = v
	return s
}

type DeleteAccessRuleRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	AccessRuleId    *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	FileSystemType  *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DeleteAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleRequest) SetAccessGroupName(v string) *DeleteAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetAccessRuleId(v string) *DeleteAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetFileSystemType(v string) *DeleteAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

type DeleteAccessRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleResponseBody) SetRequestId(v string) *DeleteAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessRuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleResponse) SetHeaders(v map[string]*string) *DeleteAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessRuleResponse) SetBody(v *DeleteAccessRuleResponseBody) *DeleteAccessRuleResponse {
	s.Body = v
	return s
}

type DeleteAutoSnapshotPolicyRequest struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *DeleteAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

type DeleteAutoSnapshotPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyResponseBody) SetRequestId(v string) *DeleteAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAutoSnapshotPolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *DeleteAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoSnapshotPolicyResponse) SetBody(v *DeleteAutoSnapshotPolicyResponseBody) *DeleteAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type DeleteFileSystemRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemRequest) SetFileSystemId(v string) *DeleteFileSystemRequest {
	s.FileSystemId = &v
	return s
}

type DeleteFileSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemResponseBody) SetRequestId(v string) *DeleteFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFileSystemResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemResponse) SetHeaders(v map[string]*string) *DeleteFileSystemResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileSystemResponse) SetBody(v *DeleteFileSystemResponseBody) *DeleteFileSystemResponse {
	s.Body = v
	return s
}

type DeleteLDAPConfigRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigRequest) SetFileSystemId(v string) *DeleteLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

type DeleteLDAPConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigResponseBody) SetRequestId(v string) *DeleteLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLDAPConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLDAPConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigResponse) SetHeaders(v map[string]*string) *DeleteLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLDAPConfigResponse) SetBody(v *DeleteLDAPConfigResponseBody) *DeleteLDAPConfigResponse {
	s.Body = v
	return s
}

type DeleteLifecyclePolicyRequest struct {
	FileSystemId        *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
}

func (s DeleteLifecyclePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyRequest) SetFileSystemId(v string) *DeleteLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *DeleteLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

type DeleteLifecyclePolicyResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLifecyclePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyResponseBody) SetSuccess(v bool) *DeleteLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLifecyclePolicyResponseBody) SetRequestId(v string) *DeleteLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLifecyclePolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLifecyclePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyResponse) SetHeaders(v map[string]*string) *DeleteLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteLifecyclePolicyResponse) SetBody(v *DeleteLifecyclePolicyResponseBody) *DeleteLifecyclePolicyResponse {
	s.Body = v
	return s
}

type DeleteMountTargetRequest struct {
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
}

func (s DeleteMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetRequest) SetFileSystemId(v string) *DeleteMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteMountTargetRequest) SetMountTargetDomain(v string) *DeleteMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

type DeleteMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetResponseBody) SetRequestId(v string) *DeleteMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMountTargetResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetResponse) SetHeaders(v map[string]*string) *DeleteMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteMountTargetResponse) SetBody(v *DeleteMountTargetResponseBody) *DeleteMountTargetResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type DeleteSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DescribeAccessGroupsRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	UseUTCDateTime  *bool   `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
	FileSystemType  *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DescribeAccessGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsRequest) SetAccessGroupName(v string) *DescribeAccessGroupsRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetPageSize(v int32) *DescribeAccessGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetPageNumber(v int32) *DescribeAccessGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetUseUTCDateTime(v bool) *DescribeAccessGroupsRequest {
	s.UseUTCDateTime = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetFileSystemType(v string) *DescribeAccessGroupsRequest {
	s.FileSystemType = &v
	return s
}

type DescribeAccessGroupsResponseBody struct {
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	AccessGroups *DescribeAccessGroupsResponseBodyAccessGroups `json:"AccessGroups,omitempty" xml:"AccessGroups,omitempty" type:"Struct"`
}

func (s DescribeAccessGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBody) SetPageNumber(v int32) *DescribeAccessGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetRequestId(v string) *DescribeAccessGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetPageSize(v int32) *DescribeAccessGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetTotalCount(v int32) *DescribeAccessGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetAccessGroups(v *DescribeAccessGroupsResponseBodyAccessGroups) *DescribeAccessGroupsResponseBody {
	s.AccessGroups = v
	return s
}

type DescribeAccessGroupsResponseBodyAccessGroups struct {
	AccessGroup []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty" type:"Repeated"`
}

func (s DescribeAccessGroupsResponseBodyAccessGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBodyAccessGroups) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBodyAccessGroups) SetAccessGroup(v []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) *DescribeAccessGroupsResponseBodyAccessGroups {
	s.AccessGroup = v
	return s
}

type DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup struct {
	AccessGroupName  *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AccessGroupType  *string `json:"AccessGroupType,omitempty" xml:"AccessGroupType,omitempty"`
	RuleCount        *int32  `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	MountTargetCount *int32  `json:"MountTargetCount,omitempty" xml:"MountTargetCount,omitempty"`
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetAccessGroupName(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetDescription(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.Description = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetAccessGroupType(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.AccessGroupType = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetRuleCount(v int32) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.RuleCount = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetMountTargetCount(v int32) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.MountTargetCount = &v
	return s
}

type DescribeAccessGroupsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccessGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponse) SetHeaders(v map[string]*string) *DescribeAccessGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessGroupsResponse) SetBody(v *DescribeAccessGroupsResponseBody) *DescribeAccessGroupsResponse {
	s.Body = v
	return s
}

type DescribeAccessRulesRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	AccessRuleId    *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FileSystemType  *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DescribeAccessRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesRequest) SetAccessGroupName(v string) *DescribeAccessRulesRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetAccessRuleId(v string) *DescribeAccessRulesRequest {
	s.AccessRuleId = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetPageSize(v int32) *DescribeAccessRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetPageNumber(v int32) *DescribeAccessRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetFileSystemType(v string) *DescribeAccessRulesRequest {
	s.FileSystemType = &v
	return s
}

type DescribeAccessRulesResponseBody struct {
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	AccessRules *DescribeAccessRulesResponseBodyAccessRules `json:"AccessRules,omitempty" xml:"AccessRules,omitempty" type:"Struct"`
}

func (s DescribeAccessRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBody) SetPageNumber(v int32) *DescribeAccessRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetRequestId(v string) *DescribeAccessRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetPageSize(v int32) *DescribeAccessRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetTotalCount(v int32) *DescribeAccessRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetAccessRules(v *DescribeAccessRulesResponseBodyAccessRules) *DescribeAccessRulesResponseBody {
	s.AccessRules = v
	return s
}

type DescribeAccessRulesResponseBodyAccessRules struct {
	AccessRule []*DescribeAccessRulesResponseBodyAccessRulesAccessRule `json:"AccessRule,omitempty" xml:"AccessRule,omitempty" type:"Repeated"`
}

func (s DescribeAccessRulesResponseBodyAccessRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponseBodyAccessRules) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBodyAccessRules) SetAccessRule(v []*DescribeAccessRulesResponseBodyAccessRulesAccessRule) *DescribeAccessRulesResponseBodyAccessRules {
	s.AccessRule = v
	return s
}

type DescribeAccessRulesResponseBodyAccessRulesAccessRule struct {
	AccessRuleId     *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	SourceCidrIp     *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	RWAccess         *string `json:"RWAccess,omitempty" xml:"RWAccess,omitempty"`
	UserAccess       *string `json:"UserAccess,omitempty" xml:"UserAccess,omitempty"`
	Priority         *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s DescribeAccessRulesResponseBodyAccessRulesAccessRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponseBodyAccessRulesAccessRule) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetAccessRuleId(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.AccessRuleId = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetSourceCidrIp(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetIpv6SourceCidrIp(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetRWAccess(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.RWAccess = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetUserAccess(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.UserAccess = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetPriority(v int32) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.Priority = &v
	return s
}

type DescribeAccessRulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccessRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccessRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponse) SetHeaders(v map[string]*string) *DescribeAccessRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessRulesResponse) SetBody(v *DescribeAccessRulesResponseBody) *DescribeAccessRulesResponse {
	s.Body = v
	return s
}

type DescribeAutoSnapshotPoliciesRequest struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FileSystemType       *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPoliciesRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetPageSize(v int32) *DescribeAutoSnapshotPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetFileSystemType(v string) *DescribeAutoSnapshotPoliciesRequest {
	s.FileSystemType = &v
	return s
}

type DescribeAutoSnapshotPoliciesResponseBody struct {
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	AutoSnapshotPolicies *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies `json:"AutoSnapshotPolicies,omitempty" xml:"AutoSnapshotPolicies,omitempty" type:"Struct"`
}

func (s DescribeAutoSnapshotPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetRequestId(v string) *DescribeAutoSnapshotPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetAutoSnapshotPolicies(v *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPoliciesResponseBody {
	s.AutoSnapshotPolicies = v
	return s
}

type DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies struct {
	AutoSnapshotPolicy []*DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy `json:"AutoSnapshotPolicy,omitempty" xml:"AutoSnapshotPolicy,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) SetAutoSnapshotPolicy(v []*DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies {
	s.AutoSnapshotPolicy = v
	return s
}

type DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy struct {
	TimePoints             *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RepeatWeekdays         *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AutoSnapshotPolicyId   *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	RetentionDays          *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	FileSystemNums         *int32  `json:"FileSystemNums,omitempty" xml:"FileSystemNums,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetTimePoints(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.TimePoints = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetStatus(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.Status = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRepeatWeekdays(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RepeatWeekdays = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyName(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetCreateTime(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.CreateTime = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRetentionDays(v int32) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RetentionDays = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetFileSystemNums(v int32) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.FileSystemNums = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRegionId(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RegionId = &v
	return s
}

type DescribeAutoSnapshotPoliciesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoSnapshotPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoSnapshotPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetBody(v *DescribeAutoSnapshotPoliciesResponseBody) *DescribeAutoSnapshotPoliciesResponse {
	s.Body = v
	return s
}

type DescribeAutoSnapshotTasksRequest struct {
	FileSystemIds         *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
	AutoSnapshotPolicyIds *string `json:"AutoSnapshotPolicyIds,omitempty" xml:"AutoSnapshotPolicyIds,omitempty"`
	FileSystemType        *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeAutoSnapshotTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksRequest) SetFileSystemIds(v string) *DescribeAutoSnapshotTasksRequest {
	s.FileSystemIds = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetAutoSnapshotPolicyIds(v string) *DescribeAutoSnapshotTasksRequest {
	s.AutoSnapshotPolicyIds = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetFileSystemType(v string) *DescribeAutoSnapshotTasksRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetPageSize(v int32) *DescribeAutoSnapshotTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetPageNumber(v int32) *DescribeAutoSnapshotTasksRequest {
	s.PageNumber = &v
	return s
}

type DescribeAutoSnapshotTasksResponseBody struct {
	PageNumber        *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize          *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount        *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	AutoSnapshotTasks *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks `json:"AutoSnapshotTasks,omitempty" xml:"AutoSnapshotTasks,omitempty" type:"Struct"`
}

func (s DescribeAutoSnapshotTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetRequestId(v string) *DescribeAutoSnapshotTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetAutoSnapshotTasks(v *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) *DescribeAutoSnapshotTasksResponseBody {
	s.AutoSnapshotTasks = v
	return s
}

type DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks struct {
	AutoSnapshotTask []*DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask `json:"AutoSnapshotTask,omitempty" xml:"AutoSnapshotTask,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) SetAutoSnapshotTask(v []*DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks {
	s.AutoSnapshotTask = v
	return s
}

type DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	SourceFileSystemId   *string `json:"SourceFileSystemId,omitempty" xml:"SourceFileSystemId,omitempty"`
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask) SetSourceFileSystemId(v string) *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasksAutoSnapshotTask {
	s.SourceFileSystemId = &v
	return s
}

type DescribeAutoSnapshotTasksResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoSnapshotTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoSnapshotTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponse) SetHeaders(v map[string]*string) *DescribeAutoSnapshotTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoSnapshotTasksResponse) SetBody(v *DescribeAutoSnapshotTasksResponseBody) *DescribeAutoSnapshotTasksResponse {
	s.Body = v
	return s
}

type DescribeBlackListClientsRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
}

func (s DescribeBlackListClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsRequest) SetRegionId(v string) *DescribeBlackListClientsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBlackListClientsRequest) SetFileSystemId(v string) *DescribeBlackListClientsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeBlackListClientsRequest) SetClientIP(v string) *DescribeBlackListClientsRequest {
	s.ClientIP = &v
	return s
}

type DescribeBlackListClientsResponseBody struct {
	Clients   *string `json:"Clients,omitempty" xml:"Clients,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBlackListClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsResponseBody) SetClients(v string) *DescribeBlackListClientsResponseBody {
	s.Clients = &v
	return s
}

func (s *DescribeBlackListClientsResponseBody) SetRequestId(v string) *DescribeBlackListClientsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBlackListClientsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBlackListClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBlackListClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsResponse) SetHeaders(v map[string]*string) *DescribeBlackListClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlackListClientsResponse) SetBody(v *DescribeBlackListClientsResponseBody) *DescribeBlackListClientsResponse {
	s.Body = v
	return s
}

type DescribeDirQuotasRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDirQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasRequest) SetFileSystemId(v string) *DescribeDirQuotasRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPath(v string) *DescribeDirQuotasRequest {
	s.Path = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPageSize(v int32) *DescribeDirQuotasRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPageNumber(v int32) *DescribeDirQuotasRequest {
	s.PageNumber = &v
	return s
}

type DescribeDirQuotasResponseBody struct {
	PageNumber    *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize      *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DirQuotaInfos []*DescribeDirQuotasResponseBodyDirQuotaInfos `json:"DirQuotaInfos,omitempty" xml:"DirQuotaInfos,omitempty" type:"Repeated"`
}

func (s DescribeDirQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBody) SetPageNumber(v int32) *DescribeDirQuotasResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetRequestId(v string) *DescribeDirQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetPageSize(v int32) *DescribeDirQuotasResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetTotalCount(v int32) *DescribeDirQuotasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetDirQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfos) *DescribeDirQuotasResponseBody {
	s.DirQuotaInfos = v
	return s
}

type DescribeDirQuotasResponseBodyDirQuotaInfos struct {
	Path           *string                                                     `json:"Path,omitempty" xml:"Path,omitempty"`
	Status         *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	DirInode       *string                                                     `json:"DirInode,omitempty" xml:"DirInode,omitempty"`
	UserQuotaInfos []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos `json:"UserQuotaInfos,omitempty" xml:"UserQuotaInfos,omitempty" type:"Repeated"`
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfos) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetPath(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.Path = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetStatus(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.Status = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetDirInode(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.DirInode = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetUserQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.UserQuotaInfos = v
	return s
}

type DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos struct {
	FileCountReal  *int64  `json:"FileCountReal,omitempty" xml:"FileCountReal,omitempty"`
	UserType       *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	FileCountLimit *int64  `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SizeLimit      *int64  `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	QuotaType      *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	SizeReal       *int64  `json:"SizeReal,omitempty" xml:"SizeReal,omitempty"`
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetFileCountReal(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.FileCountReal = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetUserType(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.UserType = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetFileCountLimit(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.FileCountLimit = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetUserId(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.UserId = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeLimit(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeLimit = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetQuotaType(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.QuotaType = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeReal(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeReal = &v
	return s
}

type DescribeDirQuotasResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDirQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDirQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponse) SetHeaders(v map[string]*string) *DescribeDirQuotasResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirQuotasResponse) SetBody(v *DescribeDirQuotasResponseBody) *DescribeDirQuotasResponse {
	s.Body = v
	return s
}

type DescribeFileSystemsRequest struct {
	FileSystemId   *string                          `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType *string                          `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	VpcId          *string                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	PageSize       *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Tag            []*DescribeFileSystemsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequest) SetFileSystemId(v string) *DescribeFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetFileSystemType(v string) *DescribeFileSystemsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetVpcId(v string) *DescribeFileSystemsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPageSize(v int32) *DescribeFileSystemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPageNumber(v int32) *DescribeFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetTag(v []*DescribeFileSystemsRequestTag) *DescribeFileSystemsRequest {
	s.Tag = v
	return s
}

type DescribeFileSystemsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsRequestTag) SetKey(v string) *DescribeFileSystemsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsRequestTag) SetValue(v string) *DescribeFileSystemsRequestTag {
	s.Value = &v
	return s
}

type DescribeFileSystemsResponseBody struct {
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	FileSystems *DescribeFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
}

func (s DescribeFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBody) SetPageNumber(v int32) *DescribeFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetRequestId(v string) *DescribeFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageSize(v int32) *DescribeFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetTotalCount(v int32) *DescribeFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetFileSystems(v *DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystems struct {
	FileSystem []*DescribeFileSystemsResponseBodyFileSystemsFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetFileSystem(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystem) *DescribeFileSystemsResponseBodyFileSystems {
	s.FileSystem = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystem struct {
	Status            *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Capacity          *int64                                                                 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	MeteredIASize     *int64                                                                 `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	CreateTime        *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType        *string                                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	StorageType       *string                                                                `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	RegionId          *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemType    *string                                                                `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	FileSystemId      *string                                                                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MeteredSize       *int64                                                                 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	EncryptType       *int32                                                                 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	Bandwidth         *int64                                                                 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description       *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Version           *string                                                                `json:"Version,omitempty" xml:"Version,omitempty"`
	ExpiredTime       *string                                                                `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ZoneId            *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ProtocolType      *string                                                                `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	KMSKeyId          *string                                                                `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	MountTargets      *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets      `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Struct"`
	Packages          *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages          `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	Tags              *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	SupportedFeatures *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures `json:"SupportedFeatures,omitempty" xml:"SupportedFeatures,omitempty" type:"Struct"`
	Ldap              *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap              `json:"Ldap,omitempty" xml:"Ldap,omitempty" type:"Struct"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetChargeType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetEncryptType(v int32) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.EncryptType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetBandwidth(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Bandwidth = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetDescription(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVersion(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Version = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetExpiredTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetKMSKeyId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMountTargets(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MountTargets = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetPackages(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Packages = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetTags(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Tags = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetSupportedFeatures(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.SupportedFeatures = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetLdap(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Ldap = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets struct {
	MountTarget []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget `json:"MountTarget,omitempty" xml:"MountTarget,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) SetMountTarget(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets {
	s.MountTarget = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget struct {
	VpcId                      *string                                                                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status                     *string                                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	MountTargetDomain          *string                                                                                       `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	AccessGroupName            *string                                                                                       `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	DualStackMountTargetDomain *string                                                                                       `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	VswId                      *string                                                                                       `json:"VswId,omitempty" xml:"VswId,omitempty"`
	NetworkType                *string                                                                                       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	ClientMasterNodes          *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
	Tags                       *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetVpcId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetAccessGroupName(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetDualStackMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetVswId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.VswId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetNetworkType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.NetworkType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetClientMasterNodes(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.ClientMasterNodes = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetTags(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.Tags = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes struct {
	ClientMasterNode []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode `json:"ClientMasterNode,omitempty" xml:"ClientMasterNode,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) SetClientMasterNode(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes {
	s.ClientMasterNode = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode struct {
	EcsIp         *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
	EcsId         *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsIp(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsIp = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetDefaultPasswd(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.DefaultPasswd = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags struct {
	Tag []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) SetTag(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags {
	s.Tag = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) SetKey(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag) SetValue(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTagsTag {
	s.Value = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages struct {
	Package []*DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) SetPackage(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages {
	s.Package = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage struct {
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetStartTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetExpiredTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageType = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemTags struct {
	Tag []*DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) SetTag(v []*DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags {
	s.Tag = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) SetKey(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag) SetValue(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemTagsTag {
	s.Value = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures struct {
	SupportedFeature []*string `json:"SupportedFeature,omitempty" xml:"SupportedFeature,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) SetSupportedFeature(v []*string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures {
	s.SupportedFeature = v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap struct {
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI        *string `json:"URI,omitempty" xml:"URI,omitempty"`
	BindDN     *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetSearchBase(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.SearchBase = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetURI(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.URI = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetBindDN(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.BindDN = &v
	return s
}

type DescribeFileSystemsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponse) SetHeaders(v map[string]*string) *DescribeFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileSystemsResponse) SetBody(v *DescribeFileSystemsResponseBody) *DescribeFileSystemsResponse {
	s.Body = v
	return s
}

type DescribeFileSystemStatisticsRequest struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeFileSystemStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsRequest) SetPageSize(v int32) *DescribeFileSystemStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsRequest) SetPageNumber(v int32) *DescribeFileSystemStatisticsRequest {
	s.PageNumber = &v
	return s
}

type DescribeFileSystemStatisticsResponseBody struct {
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	FileSystemStatistics *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics `json:"FileSystemStatistics,omitempty" xml:"FileSystemStatistics,omitempty" type:"Struct"`
	FileSystems          *DescribeFileSystemStatisticsResponseBodyFileSystems          `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
}

func (s DescribeFileSystemStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBody) SetRequestId(v string) *DescribeFileSystemStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetPageSize(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetPageNumber(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetTotalCount(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetFileSystemStatistics(v *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) *DescribeFileSystemStatisticsResponseBody {
	s.FileSystemStatistics = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetFileSystems(v *DescribeFileSystemStatisticsResponseBodyFileSystems) *DescribeFileSystemStatisticsResponseBody {
	s.FileSystems = v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemStatistics struct {
	FileSystemStatistic []*DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic `json:"FileSystemStatistic,omitempty" xml:"FileSystemStatistic,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) SetFileSystemStatistic(v []*DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics {
	s.FileSystemStatistic = v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic struct {
	ExpiringCount  *int32  `json:"ExpiringCount,omitempty" xml:"ExpiringCount,omitempty"`
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredSize    *int64  `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	TotalCount     *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ExpiredCount   *int32  `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetExpiringCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.ExpiringCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetFileSystemType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetMeteredSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetTotalCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.TotalCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetExpiredCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.ExpiredCount = &v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystems struct {
	FileSystem []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystems) SetFileSystem(v []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) *DescribeFileSystemStatisticsResponseBodyFileSystems {
	s.FileSystem = v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem struct {
	Status         *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	MeteredIASize  *int64                                                                 `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	Capacity       *int64                                                                 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime     *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType     *string                                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	StorageType    *string                                                                `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	RegionId       *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemId   *string                                                                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType *string                                                                `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredSize    *int64                                                                 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	Description    *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime    *string                                                                `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ZoneId         *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ProtocolType   *string                                                                `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Packages       *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetChargeType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetDescription(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetExpiredTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetPackages(v *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Packages = v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages struct {
	Package []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) SetPackage(v []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages {
	s.Package = v
	return s
}

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage struct {
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetStartTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetExpiredTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.Size = &v
	return s
}

type DescribeFileSystemStatisticsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFileSystemStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFileSystemStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponse) SetHeaders(v map[string]*string) *DescribeFileSystemStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileSystemStatisticsResponse) SetBody(v *DescribeFileSystemStatisticsResponseBody) *DescribeFileSystemStatisticsResponse {
	s.Body = v
	return s
}

type DescribeLDAPConfigRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigRequest) SetFileSystemId(v string) *DescribeLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

type DescribeLDAPConfigResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ldap      *DescribeLDAPConfigResponseBodyLdap `json:"Ldap,omitempty" xml:"Ldap,omitempty" type:"Struct"`
}

func (s DescribeLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigResponseBody) SetRequestId(v string) *DescribeLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLDAPConfigResponseBody) SetLdap(v *DescribeLDAPConfigResponseBodyLdap) *DescribeLDAPConfigResponseBody {
	s.Ldap = v
	return s
}

type DescribeLDAPConfigResponseBodyLdap struct {
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI        *string `json:"URI,omitempty" xml:"URI,omitempty"`
	BindDN     *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
}

func (s DescribeLDAPConfigResponseBodyLdap) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigResponseBodyLdap) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetSearchBase(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.SearchBase = &v
	return s
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetURI(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.URI = &v
	return s
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetBindDN(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.BindDN = &v
	return s
}

type DescribeLDAPConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLDAPConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigResponse) SetHeaders(v map[string]*string) *DescribeLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLDAPConfigResponse) SetBody(v *DescribeLDAPConfigResponseBody) *DescribeLDAPConfigResponse {
	s.Body = v
	return s
}

type DescribeLifecyclePoliciesRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeLifecyclePoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesRequest) SetFileSystemId(v string) *DescribeLifecyclePoliciesRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageSize(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageNumber(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageNumber = &v
	return s
}

type DescribeLifecyclePoliciesResponseBody struct {
	TotalCount        *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize          *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber        *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	LifecyclePolicies []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies `json:"LifecyclePolicies,omitempty" xml:"LifecyclePolicies,omitempty" type:"Repeated"`
}

func (s DescribeLifecyclePoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBody) SetTotalCount(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetRequestId(v string) *DescribeLifecyclePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetPageSize(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetPageNumber(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetLifecyclePolicies(v []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) *DescribeLifecyclePoliciesResponseBody {
	s.LifecyclePolicies = v
	return s
}

type DescribeLifecyclePoliciesResponseBodyLifecyclePolicies struct {
	FileSystemId        *string   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	LifecycleRuleName   *string   `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	CreateTime          *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Path                *string   `json:"Path,omitempty" xml:"Path,omitempty"`
	StorageType         *string   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	LifecyclePolicyName *string   `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	Paths               []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetFileSystemId(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecycleRuleName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecycleRuleName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetCreateTime(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.CreateTime = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetPath(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Path = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetStorageType(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.StorageType = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetPaths(v []*string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Paths = v
	return s
}

type DescribeLifecyclePoliciesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLifecyclePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLifecyclePoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponse) SetHeaders(v map[string]*string) *DescribeLifecyclePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLifecyclePoliciesResponse) SetBody(v *DescribeLifecyclePoliciesResponseBody) *DescribeLifecyclePoliciesResponse {
	s.Body = v
	return s
}

type DescribeLogAnalysisRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeLogAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisRequest) SetRegionId(v string) *DescribeLogAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetPageSize(v int32) *DescribeLogAnalysisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetPageNumber(v int32) *DescribeLogAnalysisRequest {
	s.PageNumber = &v
	return s
}

type DescribeLogAnalysisResponseBody struct {
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Analyses   *DescribeLogAnalysisResponseBodyAnalyses `json:"Analyses,omitempty" xml:"Analyses,omitempty" type:"Struct"`
}

func (s DescribeLogAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBody) SetCode(v string) *DescribeLogAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetPageNumber(v int32) *DescribeLogAnalysisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetPageSize(v int32) *DescribeLogAnalysisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetRequestId(v string) *DescribeLogAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetTotalCount(v int32) *DescribeLogAnalysisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLogAnalysisResponseBody) SetAnalyses(v *DescribeLogAnalysisResponseBodyAnalyses) *DescribeLogAnalysisResponseBody {
	s.Analyses = v
	return s
}

type DescribeLogAnalysisResponseBodyAnalyses struct {
	Analysis []*DescribeLogAnalysisResponseBodyAnalysesAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Repeated"`
}

func (s DescribeLogAnalysisResponseBodyAnalyses) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalyses) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalyses) SetAnalysis(v []*DescribeLogAnalysisResponseBodyAnalysesAnalysis) *DescribeLogAnalysisResponseBodyAnalyses {
	s.Analysis = v
	return s
}

type DescribeLogAnalysisResponseBodyAnalysesAnalysis struct {
	MetaKey   *string                                                   `json:"MetaKey,omitempty" xml:"MetaKey,omitempty"`
	MetaValue *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue `json:"MetaValue,omitempty" xml:"MetaValue,omitempty" type:"Struct"`
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysis) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysis) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) SetMetaKey(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysis {
	s.MetaKey = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysis) SetMetaValue(v *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) *DescribeLogAnalysisResponseBodyAnalysesAnalysis {
	s.MetaValue = v
	return s
}

type DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue struct {
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetLogstore(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Logstore = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetRoleArn(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.RoleArn = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetProject(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Project = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetRegion(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Region = &v
	return s
}

type DescribeLogAnalysisResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponse) SetHeaders(v map[string]*string) *DescribeLogAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogAnalysisResponse) SetBody(v *DescribeLogAnalysisResponseBody) *DescribeLogAnalysisResponse {
	s.Body = v
	return s
}

type DescribeMountedClientsRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ClientIP          *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeMountedClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsRequest) SetRegionId(v string) *DescribeMountedClientsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetPageSize(v int32) *DescribeMountedClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetFileSystemId(v string) *DescribeMountedClientsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetClientIP(v string) *DescribeMountedClientsRequest {
	s.ClientIP = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetMountTargetDomain(v string) *DescribeMountedClientsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetPageNumber(v int32) *DescribeMountedClientsRequest {
	s.PageNumber = &v
	return s
}

type DescribeMountedClientsResponseBody struct {
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Clients    *DescribeMountedClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Struct"`
}

func (s DescribeMountedClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBody) SetPageNumber(v int32) *DescribeMountedClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetRequestId(v string) *DescribeMountedClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetPageSize(v int32) *DescribeMountedClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetTotalCount(v int32) *DescribeMountedClientsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetClients(v *DescribeMountedClientsResponseBodyClients) *DescribeMountedClientsResponseBody {
	s.Clients = v
	return s
}

type DescribeMountedClientsResponseBodyClients struct {
	Client []*DescribeMountedClientsResponseBodyClientsClient `json:"Client,omitempty" xml:"Client,omitempty" type:"Repeated"`
}

func (s DescribeMountedClientsResponseBodyClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponseBodyClients) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBodyClients) SetClient(v []*DescribeMountedClientsResponseBodyClientsClient) *DescribeMountedClientsResponseBodyClients {
	s.Client = v
	return s
}

type DescribeMountedClientsResponseBodyClientsClient struct {
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
}

func (s DescribeMountedClientsResponseBodyClientsClient) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponseBodyClientsClient) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBodyClientsClient) SetClientIP(v string) *DescribeMountedClientsResponseBodyClientsClient {
	s.ClientIP = &v
	return s
}

type DescribeMountedClientsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMountedClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMountedClientsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponse) SetHeaders(v map[string]*string) *DescribeMountedClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMountedClientsResponse) SetBody(v *DescribeMountedClientsResponseBody) *DescribeMountedClientsResponse {
	s.Body = v
	return s
}

type DescribeMountTargetsRequest struct {
	FileSystemId               *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain          *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	PageSize                   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber                 *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
}

func (s DescribeMountTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsRequest) SetFileSystemId(v string) *DescribeMountTargetsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetMountTargetDomain(v string) *DescribeMountTargetsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageSize(v int32) *DescribeMountTargetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageNumber(v int32) *DescribeMountTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetDualStackMountTargetDomain(v string) *DescribeMountTargetsRequest {
	s.DualStackMountTargetDomain = &v
	return s
}

type DescribeMountTargetsResponseBody struct {
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MountTargets *DescribeMountTargetsResponseBodyMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Struct"`
}

func (s DescribeMountTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBody) SetPageNumber(v int32) *DescribeMountTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetRequestId(v string) *DescribeMountTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageSize(v int32) *DescribeMountTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetTotalCount(v int32) *DescribeMountTargetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetMountTargets(v *DescribeMountTargetsResponseBodyMountTargets) *DescribeMountTargetsResponseBody {
	s.MountTargets = v
	return s
}

type DescribeMountTargetsResponseBodyMountTargets struct {
	MountTarget []*DescribeMountTargetsResponseBodyMountTargetsMountTarget `json:"MountTarget,omitempty" xml:"MountTarget,omitempty" type:"Repeated"`
}

func (s DescribeMountTargetsResponseBodyMountTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetMountTarget(v []*DescribeMountTargetsResponseBodyMountTargetsMountTarget) *DescribeMountTargetsResponseBodyMountTargets {
	s.MountTarget = v
	return s
}

type DescribeMountTargetsResponseBodyMountTargetsMountTarget struct {
	VpcId                      *string                                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status                     *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	MountTargetDomain          *string                                                                   `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	AccessGroup                *string                                                                   `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	DualStackMountTargetDomain *string                                                                   `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	VswId                      *string                                                                   `json:"VswId,omitempty" xml:"VswId,omitempty"`
	NetworkType                *string                                                                   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	ClientMasterNodes          *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTarget) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetVpcId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.VpcId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetStatus(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.Status = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetAccessGroup(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.AccessGroup = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetDualStackMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetVswId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.VswId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetNetworkType(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.NetworkType = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetClientMasterNodes(v *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.ClientMasterNodes = v
	return s
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes struct {
	ClientMasterNode []*DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode `json:"ClientMasterNode,omitempty" xml:"ClientMasterNode,omitempty" type:"Repeated"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) SetClientMasterNode(v []*DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes {
	s.ClientMasterNode = v
	return s
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode struct {
	EcsIp         *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
	EcsId         *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsIp(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsIp = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetDefaultPasswd(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.DefaultPasswd = &v
	return s
}

type DescribeMountTargetsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMountTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMountTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponse) SetHeaders(v map[string]*string) *DescribeMountTargetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMountTargetsResponse) SetBody(v *DescribeMountTargetsResponseBody) *DescribeMountTargetsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetPageSize(v int32) *DescribeRegionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRegionsRequest) SetPageNumber(v int32) *DescribeRegionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRegionsRequest) SetFileSystemType(v string) *DescribeRegionsRequest {
	s.FileSystemType = &v
	return s
}

type DescribeRegionsResponseBody struct {
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Regions    *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetPageNumber(v int32) *DescribeRegionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetPageSize(v int32) *DescribeRegionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetTotalCount(v int32) *DescribeRegionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSnapshotsRequest struct {
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	FileSystemId   *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	SnapshotIds    *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotType   *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetFileSystemType(v string) *DescribeSnapshotsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetFileSystemId(v string) *DescribeSnapshotsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotIds(v string) *DescribeSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotName(v string) *DescribeSnapshotsRequest {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotType(v string) *DescribeSnapshotsRequest {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetStatus(v string) *DescribeSnapshotsRequest {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageSize(v int32) *DescribeSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageNumber(v int32) *DescribeSnapshotsRequest {
	s.PageNumber = &v
	return s
}

type DescribeSnapshotsResponseBody struct {
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Snapshots  *DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) SetPageNumber(v int32) *DescribeSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetPageSize(v int32) *DescribeSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetTotalCount(v int32) *DescribeSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v *DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

type DescribeSnapshotsResponseBodySnapshots struct {
	Snapshot []*DescribeSnapshotsResponseBodySnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshot(v []*DescribeSnapshotsResponseBodySnapshotsSnapshot) *DescribeSnapshotsResponseBodySnapshots {
	s.Snapshot = v
	return s
}

type DescribeSnapshotsResponseBodySnapshotsSnapshot struct {
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Progress                *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SourceFileSystemId      *string `json:"SourceFileSystemId,omitempty" xml:"SourceFileSystemId,omitempty"`
	RemainTime              *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	RetentionDays           *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SourceFileSystemSize    *int64  `json:"SourceFileSystemSize,omitempty" xml:"SourceFileSystemSize,omitempty"`
	SnapshotName            *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SourceFileSystemVersion *string `json:"SourceFileSystemVersion,omitempty" xml:"SourceFileSystemVersion,omitempty"`
	EncryptType             *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotId              *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetCreateTime(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.CreateTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRetentionDays(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RetentionDays = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemSize(v int64) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemVersion(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemVersion = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetEncryptType(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.EncryptType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotId = &v
	return s
}

type DescribeSnapshotsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotsResponse) SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeStoragePackagesRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UseUTCDateTime *bool   `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeStoragePackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesRequest) SetRegionId(v string) *DescribeStoragePackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetPageSize(v int32) *DescribeStoragePackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetUseUTCDateTime(v bool) *DescribeStoragePackagesRequest {
	s.UseUTCDateTime = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetPageNumber(v int32) *DescribeStoragePackagesRequest {
	s.PageNumber = &v
	return s
}

type DescribeStoragePackagesResponseBody struct {
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Packages   *DescribeStoragePackagesResponseBodyPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
}

func (s DescribeStoragePackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBody) SetPageNumber(v int32) *DescribeStoragePackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetRequestId(v string) *DescribeStoragePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetPageSize(v int32) *DescribeStoragePackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetTotalCount(v int32) *DescribeStoragePackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetPackages(v *DescribeStoragePackagesResponseBodyPackages) *DescribeStoragePackagesResponseBody {
	s.Packages = v
	return s
}

type DescribeStoragePackagesResponseBodyPackages struct {
	Package []*DescribeStoragePackagesResponseBodyPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeStoragePackagesResponseBodyPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponseBodyPackages) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBodyPackages) SetPackage(v []*DescribeStoragePackagesResponseBodyPackagesPackage) *DescribeStoragePackagesResponseBodyPackages {
	s.Package = v
	return s
}

type DescribeStoragePackagesResponseBodyPackagesPackage struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ExpiredTime  *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	PackageId    *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
}

func (s DescribeStoragePackagesResponseBodyPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponseBodyPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStatus(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.Status = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetFileSystemId(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.FileSystemId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStartTime(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetExpiredTime(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetSize(v int64) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStorageType(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.StorageType = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetPackageId(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.PackageId = &v
	return s
}

type DescribeStoragePackagesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStoragePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStoragePackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponse) SetHeaders(v map[string]*string) *DescribeStoragePackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoragePackagesResponse) SetBody(v *DescribeStoragePackagesResponseBody) *DescribeStoragePackagesResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	FileSystemId *string                   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageSize     *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Tag          []*DescribeTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetFileSystemId(v string) *DescribeTagsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int32) *DescribeTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsRequest) SetPageNumber(v int32) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetTag(v []*DescribeTagsRequestTag) *DescribeTagsRequest {
	s.Tag = v
	return s
}

type DescribeTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequestTag) SetKey(v string) *DescribeTagsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTagsRequestTag) SetValue(v string) *DescribeTagsRequestTag {
	s.Value = &v
	return s
}

type DescribeTagsResponseBody struct {
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Tags       *DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) SetPageNumber(v int32) *DescribeTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageSize(v int32) *DescribeTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int32) *DescribeTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v *DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

type DescribeTagsResponseBodyTags struct {
	Tag []*DescribeTagsResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) SetTag(v []*DescribeTagsResponseBodyTagsTag) *DescribeTagsResponseBodyTags {
	s.Tag = v
	return s
}

type DescribeTagsResponseBodyTagsTag struct {
	Key           *string                                       `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string                                       `json:"Value,omitempty" xml:"Value,omitempty"`
	FileSystemIds *DescribeTagsResponseBodyTagsTagFileSystemIds `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Struct"`
}

func (s DescribeTagsResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagsTag) SetKey(v string) *DescribeTagsResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) SetValue(v string) *DescribeTagsResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) SetFileSystemIds(v *DescribeTagsResponseBodyTagsTagFileSystemIds) *DescribeTagsResponseBodyTagsTag {
	s.FileSystemIds = v
	return s
}

type DescribeTagsResponseBodyTagsTagFileSystemIds struct {
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTagsTagFileSystemIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTagsTagFileSystemIds) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagsTagFileSystemIds) SetFileSystemId(v []*string) *DescribeTagsResponseBodyTagsTagFileSystemIds {
	s.FileSystemId = v
	return s
}

type DescribeTagsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) SetHeaders(v map[string]*string) *DescribeTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetFileSystemType(v string) *DescribeZonesRequest {
	s.FileSystemType = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZonesResponseBodyZonesZone struct {
	Performance   *DescribeZonesResponseBodyZonesZonePerformance   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	Capacity      *DescribeZonesResponseBodyZonesZoneCapacity      `json:"Capacity,omitempty" xml:"Capacity,omitempty" type:"Struct"`
	ZoneId        *string                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceTypes *DescribeZonesResponseBodyZonesZoneInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetPerformance(v *DescribeZonesResponseBodyZonesZonePerformance) *DescribeZonesResponseBodyZonesZone {
	s.Performance = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCapacity(v *DescribeZonesResponseBodyZonesZoneCapacity) *DescribeZonesResponseBodyZonesZone {
	s.Capacity = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetInstanceTypes(v *DescribeZonesResponseBodyZonesZoneInstanceTypes) *DescribeZonesResponseBodyZonesZone {
	s.InstanceTypes = v
	return s
}

type DescribeZonesResponseBodyZonesZonePerformance struct {
	Protocol []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZonePerformance) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZonePerformance) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZonePerformance) SetProtocol(v []*string) *DescribeZonesResponseBodyZonesZonePerformance {
	s.Protocol = v
	return s
}

type DescribeZonesResponseBodyZonesZoneCapacity struct {
	Protocol []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneCapacity) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneCapacity) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneCapacity) SetProtocol(v []*string) *DescribeZonesResponseBodyZonesZoneCapacity {
	s.Protocol = v
	return s
}

type DescribeZonesResponseBodyZonesZoneInstanceTypes struct {
	InstanceType []*DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypes) SetInstanceType(v []*DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) *DescribeZonesResponseBodyZonesZoneInstanceTypes {
	s.InstanceType = v
	return s
}

type DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType struct {
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) SetStorageType(v string) *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	s.StorageType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) SetProtocolType(v string) *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	s.ProtocolType = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DisableAndCleanRecycleBinRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableAndCleanRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAndCleanRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinRequest) SetFileSystemId(v string) *DisableAndCleanRecycleBinRequest {
	s.FileSystemId = &v
	return s
}

type DisableAndCleanRecycleBinResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAndCleanRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAndCleanRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinResponseBody) SetRequestId(v string) *DisableAndCleanRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type DisableAndCleanRecycleBinResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableAndCleanRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAndCleanRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAndCleanRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinResponse) SetHeaders(v map[string]*string) *DisableAndCleanRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *DisableAndCleanRecycleBinResponse) SetBody(v *DisableAndCleanRecycleBinResponseBody) *DisableAndCleanRecycleBinResponse {
	s.Body = v
	return s
}

type EnableRecycleBinRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ReservedDays *int64  `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
}

func (s EnableRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *EnableRecycleBinRequest) SetFileSystemId(v string) *EnableRecycleBinRequest {
	s.FileSystemId = &v
	return s
}

func (s *EnableRecycleBinRequest) SetReservedDays(v int64) *EnableRecycleBinRequest {
	s.ReservedDays = &v
	return s
}

type EnableRecycleBinResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRecycleBinResponseBody) SetRequestId(v string) *EnableRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type EnableRecycleBinResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *EnableRecycleBinResponse) SetHeaders(v map[string]*string) *EnableRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *EnableRecycleBinResponse) SetBody(v *EnableRecycleBinResponseBody) *EnableRecycleBinResponse {
	s.Body = v
	return s
}

type GetDirectoryOrFilePropertiesRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetDirectoryOrFilePropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesRequest) SetFileSystemId(v string) *GetDirectoryOrFilePropertiesRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesRequest) SetPath(v string) *GetDirectoryOrFilePropertiesRequest {
	s.Path = &v
	return s
}

type GetDirectoryOrFilePropertiesResponseBody struct {
	Entry     *GetDirectoryOrFilePropertiesResponseBodyEntry `json:"Entry,omitempty" xml:"Entry,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDirectoryOrFilePropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponseBody) SetEntry(v *GetDirectoryOrFilePropertiesResponseBodyEntry) *GetDirectoryOrFilePropertiesResponseBody {
	s.Entry = v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBody) SetRequestId(v string) *GetDirectoryOrFilePropertiesResponseBody {
	s.RequestId = &v
	return s
}

type GetDirectoryOrFilePropertiesResponseBodyEntry struct {
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	HasInfrequentAccessFile *bool   `json:"HasInfrequentAccessFile,omitempty" xml:"HasInfrequentAccessFile,omitempty"`
	MTime                   *string `json:"MTime,omitempty" xml:"MTime,omitempty"`
	ATime                   *string `json:"ATime,omitempty" xml:"ATime,omitempty"`
	Size                    *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	CTime                   *string `json:"CTime,omitempty" xml:"CTime,omitempty"`
	StorageType             *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RetrieveTime            *string `json:"RetrieveTime,omitempty" xml:"RetrieveTime,omitempty"`
	Inode                   *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
}

func (s GetDirectoryOrFilePropertiesResponseBodyEntry) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponseBodyEntry) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetType(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Type = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetHasInfrequentAccessFile(v bool) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.HasInfrequentAccessFile = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetMTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.MTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetATime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.ATime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetSize(v int64) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Size = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetCTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.CTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetStorageType(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.StorageType = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetName(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Name = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetRetrieveTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.RetrieveTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetInode(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Inode = &v
	return s
}

type GetDirectoryOrFilePropertiesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDirectoryOrFilePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDirectoryOrFilePropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponse) SetHeaders(v map[string]*string) *GetDirectoryOrFilePropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponse) SetBody(v *GetDirectoryOrFilePropertiesResponseBody) *GetDirectoryOrFilePropertiesResponse {
	s.Body = v
	return s
}

type GetRecycleBinAttributeRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s GetRecycleBinAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeRequest) SetFileSystemId(v string) *GetRecycleBinAttributeRequest {
	s.FileSystemId = &v
	return s
}

type GetRecycleBinAttributeResponseBody struct {
	// Id of the request
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecycleBinAttribute *GetRecycleBinAttributeResponseBodyRecycleBinAttribute `json:"RecycleBinAttribute,omitempty" xml:"RecycleBinAttribute,omitempty" type:"Struct"`
}

func (s GetRecycleBinAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponseBody) SetRequestId(v string) *GetRecycleBinAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBody) SetRecycleBinAttribute(v *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) *GetRecycleBinAttributeResponseBody {
	s.RecycleBinAttribute = v
	return s
}

type GetRecycleBinAttributeResponseBodyRecycleBinAttribute struct {
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ReservedDays *int64  `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
	EnableTime   *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
}

func (s GetRecycleBinAttributeResponseBodyRecycleBinAttribute) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetSize(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.Size = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetStatus(v string) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.Status = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetReservedDays(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.ReservedDays = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetEnableTime(v string) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.EnableTime = &v
	return s
}

type GetRecycleBinAttributeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRecycleBinAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRecycleBinAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponse) SetHeaders(v map[string]*string) *GetRecycleBinAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetRecycleBinAttributeResponse) SetBody(v *GetRecycleBinAttributeResponseBody) *GetRecycleBinAttributeResponse {
	s.Body = v
	return s
}

type ListDirectoriesAndFilesRequest struct {
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Path          *string `json:"Path,omitempty" xml:"Path,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	StorageType   *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	DirectoryOnly *bool   `json:"DirectoryOnly,omitempty" xml:"DirectoryOnly,omitempty"`
	MaxResults    *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDirectoriesAndFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesRequest) SetFileSystemId(v string) *ListDirectoriesAndFilesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetPath(v string) *ListDirectoriesAndFilesRequest {
	s.Path = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetNextToken(v string) *ListDirectoriesAndFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetStorageType(v string) *ListDirectoriesAndFilesRequest {
	s.StorageType = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetDirectoryOnly(v bool) *ListDirectoriesAndFilesRequest {
	s.DirectoryOnly = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetMaxResults(v int64) *ListDirectoriesAndFilesRequest {
	s.MaxResults = &v
	return s
}

type ListDirectoriesAndFilesResponseBody struct {
	NextToken *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Entries   []*ListDirectoriesAndFilesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
}

func (s ListDirectoriesAndFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponseBody) SetNextToken(v string) *ListDirectoriesAndFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) SetRequestId(v string) *ListDirectoriesAndFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) SetEntries(v []*ListDirectoriesAndFilesResponseBodyEntries) *ListDirectoriesAndFilesResponseBody {
	s.Entries = v
	return s
}

type ListDirectoriesAndFilesResponseBodyEntries struct {
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	HasInfrequentAccessFile *bool   `json:"HasInfrequentAccessFile,omitempty" xml:"HasInfrequentAccessFile,omitempty"`
	Ctime                   *string `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime                   *string `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Size                    *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageType             *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Atime                   *string `json:"Atime,omitempty" xml:"Atime,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RetrieveTime            *string `json:"RetrieveTime,omitempty" xml:"RetrieveTime,omitempty"`
	Inode                   *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	FileId                  *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s ListDirectoriesAndFilesResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetType(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Type = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetHasInfrequentAccessFile(v bool) *ListDirectoriesAndFilesResponseBodyEntries {
	s.HasInfrequentAccessFile = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetCtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Ctime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetMtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Mtime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetSize(v int64) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Size = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetStorageType(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.StorageType = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetAtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Atime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetName(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetRetrieveTime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.RetrieveTime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetInode(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Inode = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetFileId(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.FileId = &v
	return s
}

type ListDirectoriesAndFilesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDirectoriesAndFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDirectoriesAndFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponse) SetHeaders(v map[string]*string) *ListDirectoriesAndFilesResponse {
	s.Headers = v
	return s
}

func (s *ListDirectoriesAndFilesResponse) SetBody(v *ListDirectoriesAndFilesResponseBody) *ListDirectoriesAndFilesResponse {
	s.Body = v
	return s
}

type ListLifecycleRetrieveJobsRequest struct {
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLifecycleRetrieveJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsRequest) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsRequest) SetPageSize(v int32) *ListLifecycleRetrieveJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetPageNumber(v int32) *ListLifecycleRetrieveJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetFileSystemId(v string) *ListLifecycleRetrieveJobsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetStatus(v string) *ListLifecycleRetrieveJobsRequest {
	s.Status = &v
	return s
}

type ListLifecycleRetrieveJobsResponseBody struct {
	TotalCount            *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize              *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber            *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	LifecycleRetrieveJobs []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs `json:"LifecycleRetrieveJobs,omitempty" xml:"LifecycleRetrieveJobs,omitempty" type:"Repeated"`
}

func (s ListLifecycleRetrieveJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetTotalCount(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetRequestId(v string) *ListLifecycleRetrieveJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetPageSize(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetPageNumber(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetLifecycleRetrieveJobs(v []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) *ListLifecycleRetrieveJobsResponseBody {
	s.LifecycleRetrieveJobs = v
	return s
}

type ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs struct {
	FileSystemId        *string   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Status              *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	DiscoveredFileCount *int64    `json:"DiscoveredFileCount,omitempty" xml:"DiscoveredFileCount,omitempty"`
	UpdateTime          *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Paths               []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	RetrievedFileCount  *int64    `json:"RetrievedFileCount,omitempty" xml:"RetrievedFileCount,omitempty"`
	JobId               *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	CreateTime          *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetFileSystemId(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.FileSystemId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetStatus(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.Status = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetDiscoveredFileCount(v int64) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.DiscoveredFileCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetUpdateTime(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.UpdateTime = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetPaths(v []*string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.Paths = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetRetrievedFileCount(v int64) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.RetrievedFileCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetJobId(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.JobId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetCreateTime(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.CreateTime = &v
	return s
}

type ListLifecycleRetrieveJobsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLifecycleRetrieveJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLifecycleRetrieveJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponse) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponse) SetHeaders(v map[string]*string) *ListLifecycleRetrieveJobsResponse {
	s.Headers = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponse) SetBody(v *ListLifecycleRetrieveJobsResponseBody) *ListLifecycleRetrieveJobsResponse {
	s.Body = v
	return s
}

type ListRecentlyRecycledDirectoriesRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetFileSystemId(v string) *ListRecentlyRecycledDirectoriesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetNextToken(v string) *ListRecentlyRecycledDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetMaxResults(v int64) *ListRecentlyRecycledDirectoriesRequest {
	s.MaxResults = &v
	return s
}

type ListRecentlyRecycledDirectoriesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Id of the request
	NextToken *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Entries   []*ListRecentlyRecycledDirectoriesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
}

func (s ListRecentlyRecycledDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetRequestId(v string) *ListRecentlyRecycledDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetNextToken(v string) *ListRecentlyRecycledDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetEntries(v []*ListRecentlyRecycledDirectoriesResponseBodyEntries) *ListRecentlyRecycledDirectoriesResponseBody {
	s.Entries = v
	return s
}

type ListRecentlyRecycledDirectoriesResponseBodyEntries struct {
	FileId         *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	LastDeleteTime *string `json:"LastDeleteTime,omitempty" xml:"LastDeleteTime,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetFileId(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetPath(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.Path = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetName(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetLastDeleteTime(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.LastDeleteTime = &v
	return s
}

type ListRecentlyRecycledDirectoriesResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRecentlyRecycledDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecentlyRecycledDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetHeaders(v map[string]*string) *ListRecentlyRecycledDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetBody(v *ListRecentlyRecycledDirectoriesResponseBody) *ListRecentlyRecycledDirectoriesResponse {
	s.Body = v
	return s
}

type ListRecycleBinJobsRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRecycleBinJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsRequest) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsRequest) SetFileSystemId(v string) *ListRecycleBinJobsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetJobId(v string) *ListRecycleBinJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetPageSize(v int64) *ListRecycleBinJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetPageNumber(v int64) *ListRecycleBinJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetStatus(v string) *ListRecycleBinJobsRequest {
	s.Status = &v
	return s
}

type ListRecycleBinJobsResponseBody struct {
	// Id of the request
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Jobs       []*ListRecycleBinJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
}

func (s ListRecycleBinJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponseBody) SetRequestId(v string) *ListRecycleBinJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetTotalCount(v int64) *ListRecycleBinJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetPageNumber(v int64) *ListRecycleBinJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetPageSize(v int64) *ListRecycleBinJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetJobs(v []*ListRecycleBinJobsResponseBodyJobs) *ListRecycleBinJobsResponseBody {
	s.Jobs = v
	return s
}

type ListRecycleBinJobsResponseBodyJobs struct {
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ListRecycleBinJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetId(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetType(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Type = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetFileId(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.FileId = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetStatus(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetErrorCode(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.ErrorCode = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetProgress(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Progress = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetCreateTime(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetFileName(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.FileName = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetErrorMessage(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.ErrorMessage = &v
	return s
}

type ListRecycleBinJobsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRecycleBinJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecycleBinJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsResponse) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponse) SetHeaders(v map[string]*string) *ListRecycleBinJobsResponse {
	s.Headers = v
	return s
}

func (s *ListRecycleBinJobsResponse) SetBody(v *ListRecycleBinJobsResponseBody) *ListRecycleBinJobsResponse {
	s.Body = v
	return s
}

type ListRecycledDirectoriesAndFilesRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesRequest) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetFileSystemId(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetFileId(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.FileId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetNextToken(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetMaxResults(v int64) *ListRecycledDirectoriesAndFilesRequest {
	s.MaxResults = &v
	return s
}

type ListRecycledDirectoriesAndFilesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Id of the request
	NextToken *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Entries   []*ListRecycledDirectoriesAndFilesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
}

func (s ListRecycledDirectoriesAndFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetRequestId(v string) *ListRecycledDirectoriesAndFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetNextToken(v string) *ListRecycledDirectoriesAndFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetEntries(v []*ListRecycledDirectoriesAndFilesResponseBodyEntries) *ListRecycledDirectoriesAndFilesResponseBody {
	s.Entries = v
	return s
}

type ListRecycledDirectoriesAndFilesResponseBodyEntries struct {
	FileId     *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	Inode      *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	ATime      *string `json:"ATime,omitempty" xml:"ATime,omitempty"`
	MTime      *string `json:"MTime,omitempty" xml:"MTime,omitempty"`
	CTime      *string `json:"CTime,omitempty" xml:"CTime,omitempty"`
	Size       *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetFileId(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetType(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Type = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetName(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetDeleteTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.DeleteTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetInode(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Inode = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetATime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.ATime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetMTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.MTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetCTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.CTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetSize(v int64) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Size = &v
	return s
}

type ListRecycledDirectoriesAndFilesResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRecycledDirectoriesAndFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecycledDirectoriesAndFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponse) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetHeaders(v map[string]*string) *ListRecycledDirectoriesAndFilesResponse {
	s.Headers = v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetBody(v *ListRecycledDirectoriesAndFilesResponseBody) *ListRecycledDirectoriesAndFilesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyAccessGroupRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemType  *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s ModifyAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupRequest) SetAccessGroupName(v string) *ModifyAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetDescription(v string) *ModifyAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetFileSystemType(v string) *ModifyAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

type ModifyAccessGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupResponseBody) SetRequestId(v string) *ModifyAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccessGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupResponse) SetHeaders(v map[string]*string) *ModifyAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessGroupResponse) SetBody(v *ModifyAccessGroupResponseBody) *ModifyAccessGroupResponse {
	s.Body = v
	return s
}

type ModifyAccessRuleRequest struct {
	AccessGroupName  *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	AccessRuleId     *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	SourceCidrIp     *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	RWAccessType     *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	UserAccessType   *string `json:"UserAccessType,omitempty" xml:"UserAccessType,omitempty"`
	Priority         *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	FileSystemType   *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
}

func (s ModifyAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleRequest) SetAccessGroupName(v string) *ModifyAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetAccessRuleId(v string) *ModifyAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetSourceCidrIp(v string) *ModifyAccessRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetRWAccessType(v string) *ModifyAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetUserAccessType(v string) *ModifyAccessRuleRequest {
	s.UserAccessType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetPriority(v int32) *ModifyAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetFileSystemType(v string) *ModifyAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetIpv6SourceCidrIp(v string) *ModifyAccessRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

type ModifyAccessRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleResponseBody) SetRequestId(v string) *ModifyAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccessRuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleResponse) SetHeaders(v map[string]*string) *ModifyAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessRuleResponse) SetBody(v *ModifyAccessRuleResponseBody) *ModifyAccessRuleResponse {
	s.Body = v
	return s
}

type ModifyAutoSnapshotPolicyRequest struct {
	AutoSnapshotPolicyId   *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	RepeatWeekdays         *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	RetentionDays          *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	TimePoints             *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s ModifyAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *ModifyAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetTimePoints(v string) *ModifyAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

type ModifyAutoSnapshotPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyResponseBody) SetRequestId(v string) *ModifyAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAutoSnapshotPolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *ModifyAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoSnapshotPolicyResponse) SetBody(v *ModifyAutoSnapshotPolicyResponseBody) *ModifyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type ModifyFileSystemRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequest) SetFileSystemId(v string) *ModifyFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFileSystemRequest) SetDescription(v string) *ModifyFileSystemRequest {
	s.Description = &v
	return s
}

type ModifyFileSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemResponseBody) SetRequestId(v string) *ModifyFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFileSystemResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemResponse) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemResponse) SetHeaders(v map[string]*string) *ModifyFileSystemResponse {
	s.Headers = v
	return s
}

func (s *ModifyFileSystemResponse) SetBody(v *ModifyFileSystemResponseBody) *ModifyFileSystemResponse {
	s.Body = v
	return s
}

type ModifyLDAPConfigRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	URI          *string `json:"URI,omitempty" xml:"URI,omitempty"`
	BindDN       *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	SearchBase   *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
}

func (s ModifyLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigRequest) SetFileSystemId(v string) *ModifyLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetURI(v string) *ModifyLDAPConfigRequest {
	s.URI = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetBindDN(v string) *ModifyLDAPConfigRequest {
	s.BindDN = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetSearchBase(v string) *ModifyLDAPConfigRequest {
	s.SearchBase = &v
	return s
}

type ModifyLDAPConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigResponseBody) SetRequestId(v string) *ModifyLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLDAPConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLDAPConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigResponse) SetHeaders(v map[string]*string) *ModifyLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLDAPConfigResponse) SetBody(v *ModifyLDAPConfigResponseBody) *ModifyLDAPConfigResponse {
	s.Body = v
	return s
}

type ModifyLifecyclePolicyRequest struct {
	FileSystemId        *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
	LifecycleRuleName   *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	StorageType         *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ModifyLifecyclePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyRequest) SetFileSystemId(v string) *ModifyLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *ModifyLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetPath(v string) *ModifyLifecyclePolicyRequest {
	s.Path = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetLifecycleRuleName(v string) *ModifyLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetStorageType(v string) *ModifyLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

type ModifyLifecyclePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyLifecyclePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyResponseBody) SetRequestId(v string) *ModifyLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLifecyclePolicyResponseBody) SetSuccess(v bool) *ModifyLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

type ModifyLifecyclePolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLifecyclePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLifecyclePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyResponse) SetHeaders(v map[string]*string) *ModifyLifecyclePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLifecyclePolicyResponse) SetBody(v *ModifyLifecyclePolicyResponseBody) *ModifyLifecyclePolicyResponse {
	s.Body = v
	return s
}

type ModifyMountTargetRequest struct {
	FileSystemId               *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain          *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	AccessGroupName            *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	Status                     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
}

func (s ModifyMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetRequest) SetFileSystemId(v string) *ModifyMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyMountTargetRequest) SetMountTargetDomain(v string) *ModifyMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *ModifyMountTargetRequest) SetAccessGroupName(v string) *ModifyMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyMountTargetRequest) SetStatus(v string) *ModifyMountTargetRequest {
	s.Status = &v
	return s
}

func (s *ModifyMountTargetRequest) SetDualStackMountTargetDomain(v string) *ModifyMountTargetRequest {
	s.DualStackMountTargetDomain = &v
	return s
}

type ModifyMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetResponseBody) SetRequestId(v string) *ModifyMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMountTargetResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetResponse) SetHeaders(v map[string]*string) *ModifyMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyMountTargetResponse) SetBody(v *ModifyMountTargetResponseBody) *ModifyMountTargetResponse {
	s.Body = v
	return s
}

type OpenNASServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenNASServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenNASServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenNASServiceResponseBody) SetOrderId(v string) *OpenNASServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenNASServiceResponseBody) SetRequestId(v string) *OpenNASServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenNASServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenNASServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenNASServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenNASServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenNASServiceResponse) SetHeaders(v map[string]*string) *OpenNASServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenNASServiceResponse) SetBody(v *OpenNASServiceResponseBody) *OpenNASServiceResponse {
	s.Body = v
	return s
}

type RemoveClientFromBlackListRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RemoveClientFromBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientFromBlackListRequest) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListRequest) SetRegionId(v string) *RemoveClientFromBlackListRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetFileSystemId(v string) *RemoveClientFromBlackListRequest {
	s.FileSystemId = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetClientIP(v string) *RemoveClientFromBlackListRequest {
	s.ClientIP = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetClientToken(v string) *RemoveClientFromBlackListRequest {
	s.ClientToken = &v
	return s
}

type RemoveClientFromBlackListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveClientFromBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientFromBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListResponseBody) SetRequestId(v string) *RemoveClientFromBlackListResponseBody {
	s.RequestId = &v
	return s
}

type RemoveClientFromBlackListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveClientFromBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveClientFromBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientFromBlackListResponse) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListResponse) SetHeaders(v map[string]*string) *RemoveClientFromBlackListResponse {
	s.Headers = v
	return s
}

func (s *RemoveClientFromBlackListResponse) SetBody(v *RemoveClientFromBlackListResponseBody) *RemoveClientFromBlackListResponse {
	s.Body = v
	return s
}

type RemoveTagsRequest struct {
	FileSystemId *string                 `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Tag          []*RemoveTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s RemoveTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequest) SetFileSystemId(v string) *RemoveTagsRequest {
	s.FileSystemId = &v
	return s
}

func (s *RemoveTagsRequest) SetTag(v []*RemoveTagsRequestTag) *RemoveTagsRequest {
	s.Tag = v
	return s
}

type RemoveTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RemoveTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsRequestTag) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequestTag) SetKey(v string) *RemoveTagsRequestTag {
	s.Key = &v
	return s
}

func (s *RemoveTagsRequestTag) SetValue(v string) *RemoveTagsRequestTag {
	s.Value = &v
	return s
}

type RemoveTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponseBody) SetRequestId(v string) *RemoveTagsResponseBody {
	s.RequestId = &v
	return s
}

type RemoveTagsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponse) SetHeaders(v map[string]*string) *RemoveTagsResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagsResponse) SetBody(v *RemoveTagsResponseBody) *RemoveTagsResponse {
	s.Body = v
	return s
}

type ResetFileSystemRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	SnapshotId   *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ResetFileSystemRequest) SetFileSystemId(v string) *ResetFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ResetFileSystemRequest) SetSnapshotId(v string) *ResetFileSystemRequest {
	s.SnapshotId = &v
	return s
}

type ResetFileSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ResetFileSystemResponseBody) SetRequestId(v string) *ResetFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type ResetFileSystemResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetFileSystemResponse) GoString() string {
	return s.String()
}

func (s *ResetFileSystemResponse) SetHeaders(v map[string]*string) *ResetFileSystemResponse {
	s.Headers = v
	return s
}

func (s *ResetFileSystemResponse) SetBody(v *ResetFileSystemResponseBody) *ResetFileSystemResponse {
	s.Body = v
	return s
}

type RetryLifecycleRetrieveJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s RetryLifecycleRetrieveJobRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobRequest) SetJobId(v string) *RetryLifecycleRetrieveJobRequest {
	s.JobId = &v
	return s
}

type RetryLifecycleRetrieveJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryLifecycleRetrieveJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobResponseBody) SetRequestId(v string) *RetryLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

type RetryLifecycleRetrieveJobResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetryLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetryLifecycleRetrieveJobResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *RetryLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *RetryLifecycleRetrieveJobResponse) SetBody(v *RetryLifecycleRetrieveJobResponseBody) *RetryLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

type SetDirQuotaRequest struct {
	FileSystemId   *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	QuotaType      *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	UserType       *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SizeLimit      *int64  `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	FileCountLimit *int64  `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
}

func (s SetDirQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDirQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetDirQuotaRequest) SetFileSystemId(v string) *SetDirQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *SetDirQuotaRequest) SetPath(v string) *SetDirQuotaRequest {
	s.Path = &v
	return s
}

func (s *SetDirQuotaRequest) SetQuotaType(v string) *SetDirQuotaRequest {
	s.QuotaType = &v
	return s
}

func (s *SetDirQuotaRequest) SetUserType(v string) *SetDirQuotaRequest {
	s.UserType = &v
	return s
}

func (s *SetDirQuotaRequest) SetUserId(v string) *SetDirQuotaRequest {
	s.UserId = &v
	return s
}

func (s *SetDirQuotaRequest) SetSizeLimit(v int64) *SetDirQuotaRequest {
	s.SizeLimit = &v
	return s
}

func (s *SetDirQuotaRequest) SetFileCountLimit(v int64) *SetDirQuotaRequest {
	s.FileCountLimit = &v
	return s
}

type SetDirQuotaResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDirQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDirQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetDirQuotaResponseBody) SetSuccess(v bool) *SetDirQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *SetDirQuotaResponseBody) SetRequestId(v string) *SetDirQuotaResponseBody {
	s.RequestId = &v
	return s
}

type SetDirQuotaResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDirQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDirQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDirQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetDirQuotaResponse) SetHeaders(v map[string]*string) *SetDirQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetDirQuotaResponse) SetBody(v *SetDirQuotaResponseBody) *SetDirQuotaResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateRecycleBinAttributeRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ReservedDays *int64  `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
}

func (s UpdateRecycleBinAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecycleBinAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeRequest) SetFileSystemId(v string) *UpdateRecycleBinAttributeRequest {
	s.FileSystemId = &v
	return s
}

func (s *UpdateRecycleBinAttributeRequest) SetReservedDays(v int64) *UpdateRecycleBinAttributeRequest {
	s.ReservedDays = &v
	return s
}

type UpdateRecycleBinAttributeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecycleBinAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecycleBinAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeResponseBody) SetRequestId(v string) *UpdateRecycleBinAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRecycleBinAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRecycleBinAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRecycleBinAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecycleBinAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeResponse) SetHeaders(v map[string]*string) *UpdateRecycleBinAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecycleBinAttributeResponse) SetBody(v *UpdateRecycleBinAttributeResponseBody) *UpdateRecycleBinAttributeResponse {
	s.Body = v
	return s
}

type UpgradeFileSystemRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Capacity     *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpgradeFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFileSystemRequest) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemRequest) SetFileSystemId(v string) *UpgradeFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetCapacity(v int64) *UpgradeFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetDryRun(v bool) *UpgradeFileSystemRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetClientToken(v string) *UpgradeFileSystemRequest {
	s.ClientToken = &v
	return s
}

type UpgradeFileSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemResponseBody) SetRequestId(v string) *UpgradeFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeFileSystemResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFileSystemResponse) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemResponse) SetHeaders(v map[string]*string) *UpgradeFileSystemResponse {
	s.Headers = v
	return s
}

func (s *UpgradeFileSystemResponse) SetBody(v *UpgradeFileSystemResponseBody) *UpgradeFileSystemResponse {
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
		"cn-chengdu":          tea.String("nas.aliyuncs.com"),
		"me-east-1":           tea.String("nas.ap-northeast-1.aliyuncs.com"),
		"cn-hangzhou-finance": tea.String("nas.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddClientToBlackListWithOptions(request *AddClientToBlackListRequest, runtime *util.RuntimeOptions) (_result *AddClientToBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddClientToBlackListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddClientToBlackList"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddClientToBlackList(request *AddClientToBlackListRequest) (_result *AddClientToBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddClientToBlackListResponse{}
	_body, _err := client.AddClientToBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTagsWithOptions(request *AddTagsRequest, runtime *util.RuntimeOptions) (_result *AddTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddTags"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTags(request *AddTagsRequest) (_result *AddTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTagsResponse{}
	_body, _err := client.AddTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyAutoSnapshotPolicyWithOptions(request *ApplyAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyAutoSnapshotPolicy"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyAutoSnapshotPolicy(request *ApplyAutoSnapshotPolicyRequest) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.ApplyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelAutoSnapshotPolicyWithOptions(request *CancelAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelAutoSnapshotPolicy"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelAutoSnapshotPolicy(request *CancelAutoSnapshotPolicyRequest) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CancelAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelDirQuotaWithOptions(request *CancelDirQuotaRequest, runtime *util.RuntimeOptions) (_result *CancelDirQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelDirQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelDirQuota"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelDirQuota(request *CancelDirQuotaRequest) (_result *CancelDirQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelDirQuotaResponse{}
	_body, _err := client.CancelDirQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelLifecycleRetrieveJobWithOptions(request *CancelLifecycleRetrieveJobRequest, runtime *util.RuntimeOptions) (_result *CancelLifecycleRetrieveJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelLifecycleRetrieveJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelLifecycleRetrieveJob"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelLifecycleRetrieveJob(request *CancelLifecycleRetrieveJobRequest) (_result *CancelLifecycleRetrieveJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelLifecycleRetrieveJobResponse{}
	_body, _err := client.CancelLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelRecycleBinJobWithOptions(request *CancelRecycleBinJobRequest, runtime *util.RuntimeOptions) (_result *CancelRecycleBinJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CancelRecycleBinJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelRecycleBinJob"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelRecycleBinJob(request *CancelRecycleBinJobRequest) (_result *CancelRecycleBinJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelRecycleBinJobResponse{}
	_body, _err := client.CancelRecycleBinJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccessGroupWithOptions(request *CreateAccessGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAccessGroup"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (_result *CreateAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CreateAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccessRuleWithOptions(request *CreateAccessRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAccessRule"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessRule(request *CreateAccessRuleRequest) (_result *CreateAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CreateAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAutoSnapshotPolicyWithOptions(request *CreateAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAutoSnapshotPolicy"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CreateAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFileSystemWithOptions(request *CreateFileSystemRequest, runtime *util.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFileSystem"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFileSystem(request *CreateFileSystemRequest) (_result *CreateFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CreateFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLDAPConfigWithOptions(request *CreateLDAPConfigRequest, runtime *util.RuntimeOptions) (_result *CreateLDAPConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLDAPConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLDAPConfig"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLDAPConfig(request *CreateLDAPConfigRequest) (_result *CreateLDAPConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLDAPConfigResponse{}
	_body, _err := client.CreateLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLifecyclePolicyWithOptions(request *CreateLifecyclePolicyRequest, runtime *util.RuntimeOptions) (_result *CreateLifecyclePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLifecyclePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLifecyclePolicy"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLifecyclePolicy(request *CreateLifecyclePolicyRequest) (_result *CreateLifecyclePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLifecyclePolicyResponse{}
	_body, _err := client.CreateLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLifecycleRetrieveJobWithOptions(request *CreateLifecycleRetrieveJobRequest, runtime *util.RuntimeOptions) (_result *CreateLifecycleRetrieveJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLifecycleRetrieveJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLifecycleRetrieveJob"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLifecycleRetrieveJob(request *CreateLifecycleRetrieveJobRequest) (_result *CreateLifecycleRetrieveJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLifecycleRetrieveJobResponse{}
	_body, _err := client.CreateLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMountTargetWithOptions(request *CreateMountTargetRequest, runtime *util.RuntimeOptions) (_result *CreateMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMountTarget"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMountTarget(request *CreateMountTargetRequest) (_result *CreateMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.CreateMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRecycleBinDeleteJobWithOptions(request *CreateRecycleBinDeleteJobRequest, runtime *util.RuntimeOptions) (_result *CreateRecycleBinDeleteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateRecycleBinDeleteJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRecycleBinDeleteJob"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRecycleBinDeleteJob(request *CreateRecycleBinDeleteJobRequest) (_result *CreateRecycleBinDeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecycleBinDeleteJobResponse{}
	_body, _err := client.CreateRecycleBinDeleteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRecycleBinRestoreJobWithOptions(request *CreateRecycleBinRestoreJobRequest, runtime *util.RuntimeOptions) (_result *CreateRecycleBinRestoreJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateRecycleBinRestoreJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRecycleBinRestoreJob"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRecycleBinRestoreJob(request *CreateRecycleBinRestoreJobRequest) (_result *CreateRecycleBinRestoreJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecycleBinRestoreJobResponse{}
	_body, _err := client.CreateRecycleBinRestoreJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSnapshot"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessGroupWithOptions(request *DeleteAccessGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAccessGroup"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessGroup(request *DeleteAccessGroupRequest) (_result *DeleteAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.DeleteAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessRuleWithOptions(request *DeleteAccessRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAccessRule"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessRule(request *DeleteAccessRuleRequest) (_result *DeleteAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.DeleteAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAutoSnapshotPolicyWithOptions(request *DeleteAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAutoSnapshotPolicy"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAutoSnapshotPolicy(request *DeleteAutoSnapshotPolicyRequest) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.DeleteAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileSystemWithOptions(request *DeleteFileSystemRequest, runtime *util.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFileSystem"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (_result *DeleteFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.DeleteFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLDAPConfigWithOptions(request *DeleteLDAPConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLDAPConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLDAPConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLDAPConfig"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLDAPConfig(request *DeleteLDAPConfigRequest) (_result *DeleteLDAPConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLDAPConfigResponse{}
	_body, _err := client.DeleteLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLifecyclePolicyWithOptions(request *DeleteLifecyclePolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteLifecyclePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLifecyclePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLifecyclePolicy"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLifecyclePolicy(request *DeleteLifecyclePolicyRequest) (_result *DeleteLifecyclePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLifecyclePolicyResponse{}
	_body, _err := client.DeleteLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMountTargetWithOptions(request *DeleteMountTargetRequest, runtime *util.RuntimeOptions) (_result *DeleteMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMountTarget"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMountTarget(request *DeleteMountTargetRequest) (_result *DeleteMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.DeleteMountTargetWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSnapshot"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAccessGroupsWithOptions(request *DescribeAccessGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccessGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccessGroups"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessGroups(request *DescribeAccessGroupsRequest) (_result *DescribeAccessGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessGroupsResponse{}
	_body, _err := client.DescribeAccessGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccessRulesWithOptions(request *DescribeAccessRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccessRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccessRules"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccessRules(request *DescribeAccessRulesRequest) (_result *DescribeAccessRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccessRulesResponse{}
	_body, _err := client.DescribeAccessRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoSnapshotPoliciesWithOptions(request *DescribeAutoSnapshotPoliciesRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoSnapshotPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoSnapshotPoliciesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoSnapshotPolicies"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (_result *DescribeAutoSnapshotPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoSnapshotPoliciesResponse{}
	_body, _err := client.DescribeAutoSnapshotPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoSnapshotTasksWithOptions(request *DescribeAutoSnapshotTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoSnapshotTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoSnapshotTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoSnapshotTasks"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoSnapshotTasks(request *DescribeAutoSnapshotTasksRequest) (_result *DescribeAutoSnapshotTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoSnapshotTasksResponse{}
	_body, _err := client.DescribeAutoSnapshotTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBlackListClientsWithOptions(request *DescribeBlackListClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeBlackListClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBlackListClientsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBlackListClients"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBlackListClients(request *DescribeBlackListClientsRequest) (_result *DescribeBlackListClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBlackListClientsResponse{}
	_body, _err := client.DescribeBlackListClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDirQuotasWithOptions(request *DescribeDirQuotasRequest, runtime *util.RuntimeOptions) (_result *DescribeDirQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDirQuotasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDirQuotas"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDirQuotas(request *DescribeDirQuotasRequest) (_result *DescribeDirQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirQuotasResponse{}
	_body, _err := client.DescribeDirQuotasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileSystemsWithOptions(request *DescribeFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DescribeFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFileSystems"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (_result *DescribeFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.DescribeFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileSystemStatisticsWithOptions(request *DescribeFileSystemStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeFileSystemStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFileSystemStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFileSystemStatistics"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFileSystemStatistics(request *DescribeFileSystemStatisticsRequest) (_result *DescribeFileSystemStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileSystemStatisticsResponse{}
	_body, _err := client.DescribeFileSystemStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLDAPConfigWithOptions(request *DescribeLDAPConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLDAPConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLDAPConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLDAPConfig"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLDAPConfig(request *DescribeLDAPConfigRequest) (_result *DescribeLDAPConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLDAPConfigResponse{}
	_body, _err := client.DescribeLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLifecyclePoliciesWithOptions(request *DescribeLifecyclePoliciesRequest, runtime *util.RuntimeOptions) (_result *DescribeLifecyclePoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeLifecyclePoliciesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLifecyclePolicies"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLifecyclePolicies(request *DescribeLifecyclePoliciesRequest) (_result *DescribeLifecyclePoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLifecyclePoliciesResponse{}
	_body, _err := client.DescribeLifecyclePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogAnalysisWithOptions(request *DescribeLogAnalysisRequest, runtime *util.RuntimeOptions) (_result *DescribeLogAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLogAnalysisResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLogAnalysis"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogAnalysis(request *DescribeLogAnalysisRequest) (_result *DescribeLogAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogAnalysisResponse{}
	_body, _err := client.DescribeLogAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMountedClientsWithOptions(request *DescribeMountedClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeMountedClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMountedClientsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMountedClients"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMountedClients(request *DescribeMountedClientsRequest) (_result *DescribeMountedClientsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMountedClientsResponse{}
	_body, _err := client.DescribeMountedClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMountTargetsWithOptions(request *DescribeMountTargetsRequest, runtime *util.RuntimeOptions) (_result *DescribeMountTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMountTargets"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (_result *DescribeMountTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.DescribeMountTargetsWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSnapshots"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStoragePackagesWithOptions(request *DescribeStoragePackagesRequest, runtime *util.RuntimeOptions) (_result *DescribeStoragePackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStoragePackagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStoragePackages"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStoragePackages(request *DescribeStoragePackagesRequest) (_result *DescribeStoragePackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStoragePackagesResponse{}
	_body, _err := client.DescribeStoragePackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTags"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZones"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAndCleanRecycleBinWithOptions(request *DisableAndCleanRecycleBinRequest, runtime *util.RuntimeOptions) (_result *DisableAndCleanRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DisableAndCleanRecycleBinResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableAndCleanRecycleBin"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAndCleanRecycleBin(request *DisableAndCleanRecycleBinRequest) (_result *DisableAndCleanRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAndCleanRecycleBinResponse{}
	_body, _err := client.DisableAndCleanRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableRecycleBinWithOptions(request *EnableRecycleBinRequest, runtime *util.RuntimeOptions) (_result *EnableRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableRecycleBinResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableRecycleBin"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableRecycleBin(request *EnableRecycleBinRequest) (_result *EnableRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableRecycleBinResponse{}
	_body, _err := client.EnableRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDirectoryOrFilePropertiesWithOptions(request *GetDirectoryOrFilePropertiesRequest, runtime *util.RuntimeOptions) (_result *GetDirectoryOrFilePropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDirectoryOrFilePropertiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDirectoryOrFileProperties"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDirectoryOrFileProperties(request *GetDirectoryOrFilePropertiesRequest) (_result *GetDirectoryOrFilePropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectoryOrFilePropertiesResponse{}
	_body, _err := client.GetDirectoryOrFilePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecycleBinAttributeWithOptions(request *GetRecycleBinAttributeRequest, runtime *util.RuntimeOptions) (_result *GetRecycleBinAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetRecycleBinAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRecycleBinAttribute"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecycleBinAttribute(request *GetRecycleBinAttributeRequest) (_result *GetRecycleBinAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecycleBinAttributeResponse{}
	_body, _err := client.GetRecycleBinAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDirectoriesAndFilesWithOptions(request *ListDirectoriesAndFilesRequest, runtime *util.RuntimeOptions) (_result *ListDirectoriesAndFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDirectoriesAndFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDirectoriesAndFiles"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDirectoriesAndFiles(request *ListDirectoriesAndFilesRequest) (_result *ListDirectoriesAndFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectoriesAndFilesResponse{}
	_body, _err := client.ListDirectoriesAndFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLifecycleRetrieveJobsWithOptions(request *ListLifecycleRetrieveJobsRequest, runtime *util.RuntimeOptions) (_result *ListLifecycleRetrieveJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLifecycleRetrieveJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLifecycleRetrieveJobs"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLifecycleRetrieveJobs(request *ListLifecycleRetrieveJobsRequest) (_result *ListLifecycleRetrieveJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLifecycleRetrieveJobsResponse{}
	_body, _err := client.ListLifecycleRetrieveJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecentlyRecycledDirectoriesWithOptions(request *ListRecentlyRecycledDirectoriesRequest, runtime *util.RuntimeOptions) (_result *ListRecentlyRecycledDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListRecentlyRecycledDirectoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRecentlyRecycledDirectories"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecentlyRecycledDirectories(request *ListRecentlyRecycledDirectoriesRequest) (_result *ListRecentlyRecycledDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecentlyRecycledDirectoriesResponse{}
	_body, _err := client.ListRecentlyRecycledDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecycleBinJobsWithOptions(request *ListRecycleBinJobsRequest, runtime *util.RuntimeOptions) (_result *ListRecycleBinJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListRecycleBinJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRecycleBinJobs"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecycleBinJobs(request *ListRecycleBinJobsRequest) (_result *ListRecycleBinJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecycleBinJobsResponse{}
	_body, _err := client.ListRecycleBinJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecycledDirectoriesAndFilesWithOptions(request *ListRecycledDirectoriesAndFilesRequest, runtime *util.RuntimeOptions) (_result *ListRecycledDirectoriesAndFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListRecycledDirectoriesAndFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRecycledDirectoriesAndFiles"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecycledDirectoriesAndFiles(request *ListRecycledDirectoriesAndFilesRequest) (_result *ListRecycledDirectoriesAndFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecycledDirectoriesAndFilesResponse{}
	_body, _err := client.ListRecycledDirectoriesAndFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccessGroupWithOptions(request *ModifyAccessGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccessGroup"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (_result *ModifyAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.ModifyAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccessRuleWithOptions(request *ModifyAccessRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccessRule"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccessRule(request *ModifyAccessRuleRequest) (_result *ModifyAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.ModifyAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAutoSnapshotPolicyWithOptions(request *ModifyAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAutoSnapshotPolicy"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAutoSnapshotPolicy(request *ModifyAutoSnapshotPolicyRequest) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.ModifyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFileSystemWithOptions(request *ModifyFileSystemRequest, runtime *util.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFileSystem"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (_result *ModifyFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.ModifyFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLDAPConfigWithOptions(request *ModifyLDAPConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyLDAPConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLDAPConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLDAPConfig"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLDAPConfig(request *ModifyLDAPConfigRequest) (_result *ModifyLDAPConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLDAPConfigResponse{}
	_body, _err := client.ModifyLDAPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLifecyclePolicyWithOptions(request *ModifyLifecyclePolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyLifecyclePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLifecyclePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLifecyclePolicy"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLifecyclePolicy(request *ModifyLifecyclePolicyRequest) (_result *ModifyLifecyclePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLifecyclePolicyResponse{}
	_body, _err := client.ModifyLifecyclePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMountTargetWithOptions(request *ModifyMountTargetRequest, runtime *util.RuntimeOptions) (_result *ModifyMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMountTargetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMountTarget"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMountTarget(request *ModifyMountTargetRequest) (_result *ModifyMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMountTargetResponse{}
	_body, _err := client.ModifyMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenNASServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenNASServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &OpenNASServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenNASService"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenNASService() (_result *OpenNASServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenNASServiceResponse{}
	_body, _err := client.OpenNASServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveClientFromBlackListWithOptions(request *RemoveClientFromBlackListRequest, runtime *util.RuntimeOptions) (_result *RemoveClientFromBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveClientFromBlackListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveClientFromBlackList"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveClientFromBlackList(request *RemoveClientFromBlackListRequest) (_result *RemoveClientFromBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveClientFromBlackListResponse{}
	_body, _err := client.RemoveClientFromBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTagsWithOptions(request *RemoveTagsRequest, runtime *util.RuntimeOptions) (_result *RemoveTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveTags"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTags(request *RemoveTagsRequest) (_result *RemoveTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTagsResponse{}
	_body, _err := client.RemoveTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetFileSystemWithOptions(request *ResetFileSystemRequest, runtime *util.RuntimeOptions) (_result *ResetFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetFileSystemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetFileSystem"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetFileSystem(request *ResetFileSystemRequest) (_result *ResetFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetFileSystemResponse{}
	_body, _err := client.ResetFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetryLifecycleRetrieveJobWithOptions(request *RetryLifecycleRetrieveJobRequest, runtime *util.RuntimeOptions) (_result *RetryLifecycleRetrieveJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RetryLifecycleRetrieveJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RetryLifecycleRetrieveJob"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryLifecycleRetrieveJob(request *RetryLifecycleRetrieveJobRequest) (_result *RetryLifecycleRetrieveJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryLifecycleRetrieveJobResponse{}
	_body, _err := client.RetryLifecycleRetrieveJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDirQuotaWithOptions(request *SetDirQuotaRequest, runtime *util.RuntimeOptions) (_result *SetDirQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDirQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDirQuota"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDirQuota(request *SetDirQuotaRequest) (_result *SetDirQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDirQuotaResponse{}
	_body, _err := client.SetDirQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRecycleBinAttributeWithOptions(request *UpdateRecycleBinAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateRecycleBinAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UpdateRecycleBinAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRecycleBinAttribute"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRecycleBinAttribute(request *UpdateRecycleBinAttributeRequest) (_result *UpdateRecycleBinAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRecycleBinAttributeResponse{}
	_body, _err := client.UpdateRecycleBinAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeFileSystemWithOptions(request *UpgradeFileSystemRequest, runtime *util.RuntimeOptions) (_result *UpgradeFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeFileSystemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeFileSystem"), tea.String("2017-06-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeFileSystem(request *UpgradeFileSystemRequest) (_result *UpgradeFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeFileSystemResponse{}
	_body, _err := client.UpgradeFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
