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

type AddClientToBlackListRequest struct {
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddClientToBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddClientToBlackListRequest) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListRequest) SetClientIP(v string) *AddClientToBlackListRequest {
	s.ClientIP = &v
	return s
}

func (s *AddClientToBlackListRequest) SetClientToken(v string) *AddClientToBlackListRequest {
	s.ClientToken = &v
	return s
}

func (s *AddClientToBlackListRequest) SetFileSystemId(v string) *AddClientToBlackListRequest {
	s.FileSystemId = &v
	return s
}

func (s *AddClientToBlackListRequest) SetRegionId(v string) *AddClientToBlackListRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddClientToBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddClientToBlackListResponse) SetStatusCode(v int32) *AddClientToBlackListResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddTagsResponse) SetStatusCode(v int32) *AddTagsResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ApplyAutoSnapshotPolicyResponse) SetStatusCode(v int32) *ApplyAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) SetBody(v *ApplyAutoSnapshotPolicyResponseBody) *ApplyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type ApplyDataFlowAutoRefreshRequest struct {
	AutoRefreshInterval *int64                                         `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	AutoRefreshPolicy   *string                                        `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	AutoRefreshs        []*ApplyDataFlowAutoRefreshRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
	ClientToken         *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId          *string                                        `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun              *bool                                          `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId        *string                                        `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ApplyDataFlowAutoRefreshRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshInterval(v int64) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshPolicy(v string) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshs(v []*ApplyDataFlowAutoRefreshRequestAutoRefreshs) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshs = v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetClientToken(v string) *ApplyDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetDataFlowId(v string) *ApplyDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetDryRun(v bool) *ApplyDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetFileSystemId(v string) *ApplyDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

type ApplyDataFlowAutoRefreshRequestAutoRefreshs struct {
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s ApplyDataFlowAutoRefreshRequestAutoRefreshs) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshRequestAutoRefreshs) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshRequestAutoRefreshs) SetRefreshPath(v string) *ApplyDataFlowAutoRefreshRequestAutoRefreshs {
	s.RefreshPath = &v
	return s
}

type ApplyDataFlowAutoRefreshResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyDataFlowAutoRefreshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshResponseBody) SetRequestId(v string) *ApplyDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

type ApplyDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyDataFlowAutoRefreshResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *ApplyDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *ApplyDataFlowAutoRefreshResponse) SetStatusCode(v int32) *ApplyDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshResponse) SetBody(v *ApplyDataFlowAutoRefreshResponseBody) *ApplyDataFlowAutoRefreshResponse {
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CancelAutoSnapshotPolicyResponse) SetStatusCode(v int32) *CancelAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAutoSnapshotPolicyResponse) SetBody(v *CancelAutoSnapshotPolicyResponseBody) *CancelAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CancelDataFlowAutoRefreshRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId   *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	RefreshPath  *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s CancelDataFlowAutoRefreshRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshRequest) SetClientToken(v string) *CancelDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetDataFlowId(v string) *CancelDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetDryRun(v bool) *CancelDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetFileSystemId(v string) *CancelDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetRefreshPath(v string) *CancelDataFlowAutoRefreshRequest {
	s.RefreshPath = &v
	return s
}

type CancelDataFlowAutoRefreshResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDataFlowAutoRefreshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshResponseBody) SetRequestId(v string) *CancelDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

type CancelDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelDataFlowAutoRefreshResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *CancelDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *CancelDataFlowAutoRefreshResponse) SetStatusCode(v int32) *CancelDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDataFlowAutoRefreshResponse) SetBody(v *CancelDataFlowAutoRefreshResponseBody) *CancelDataFlowAutoRefreshResponse {
	s.Body = v
	return s
}

type CancelDataFlowTaskRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId   *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelDataFlowTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskRequest) SetClientToken(v string) *CancelDataFlowTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetDataFlowId(v string) *CancelDataFlowTaskRequest {
	s.DataFlowId = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetDryRun(v bool) *CancelDataFlowTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetFileSystemId(v string) *CancelDataFlowTaskRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetTaskId(v string) *CancelDataFlowTaskRequest {
	s.TaskId = &v
	return s
}

type CancelDataFlowTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDataFlowTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskResponseBody) SetRequestId(v string) *CancelDataFlowTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelDataFlowTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelDataFlowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelDataFlowTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDataFlowTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskResponse) SetHeaders(v map[string]*string) *CancelDataFlowTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelDataFlowTaskResponse) SetStatusCode(v int32) *CancelDataFlowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDataFlowTaskResponse) SetBody(v *CancelDataFlowTaskResponseBody) *CancelDataFlowTaskResponse {
	s.Body = v
	return s
}

type CancelDirQuotaRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType     *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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

func (s *CancelDirQuotaRequest) SetUserId(v string) *CancelDirQuotaRequest {
	s.UserId = &v
	return s
}

func (s *CancelDirQuotaRequest) SetUserType(v string) *CancelDirQuotaRequest {
	s.UserType = &v
	return s
}

type CancelDirQuotaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelDirQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDirQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaResponseBody) SetRequestId(v string) *CancelDirQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDirQuotaResponseBody) SetSuccess(v bool) *CancelDirQuotaResponseBody {
	s.Success = &v
	return s
}

type CancelDirQuotaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelDirQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CancelDirQuotaResponse) SetStatusCode(v int32) *CancelDirQuotaResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CancelLifecycleRetrieveJobResponse) SetStatusCode(v int32) *CancelLifecycleRetrieveJobResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelRecycleBinJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CancelRecycleBinJobResponse) SetStatusCode(v int32) *CancelRecycleBinJobResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAccessGroupResponse) SetStatusCode(v int32) *CreateAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessGroupResponse) SetBody(v *CreateAccessGroupResponseBody) *CreateAccessGroupResponse {
	s.Body = v
	return s
}

type CreateAccessRuleRequest struct {
	AccessGroupName  *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	FileSystemType   *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	Priority         *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RWAccessType     *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	SourceCidrIp     *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	UserAccessType   *string `json:"UserAccessType,omitempty" xml:"UserAccessType,omitempty"`
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

func (s *CreateAccessRuleRequest) SetFileSystemType(v string) *CreateAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetIpv6SourceCidrIp(v string) *CreateAccessRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *CreateAccessRuleRequest) SetPriority(v int32) *CreateAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateAccessRuleRequest) SetRWAccessType(v string) *CreateAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *CreateAccessRuleRequest) SetSourceCidrIp(v string) *CreateAccessRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateAccessRuleRequest) SetUserAccessType(v string) *CreateAccessRuleRequest {
	s.UserAccessType = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAccessRuleResponse) SetStatusCode(v int32) *CreateAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessRuleResponse) SetBody(v *CreateAccessRuleResponseBody) *CreateAccessRuleResponse {
	s.Body = v
	return s
}

type CreateAutoSnapshotPolicyRequest struct {
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	FileSystemType         *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	RepeatWeekdays         *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	RetentionDays          *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	TimePoints             *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetFileSystemType(v string) *CreateAutoSnapshotPolicyRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest {
	s.TimePoints = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAutoSnapshotPolicyResponse) SetStatusCode(v int32) *CreateAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) SetBody(v *CreateAutoSnapshotPolicyResponseBody) *CreateAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CreateDataFlowRequest struct {
	AutoRefreshInterval *int64                               `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	AutoRefreshPolicy   *string                              `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	AutoRefreshs        []*CreateDataFlowRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
	ClientToken         *string                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description         *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun              *bool                                `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId        *string                              `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FsetId              *string                              `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	SourceSecurityType  *string                              `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	SourceStorage       *string                              `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	Throughput          *int64                               `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s CreateDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateDataFlowRequest) SetAutoRefreshInterval(v int64) *CreateDataFlowRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *CreateDataFlowRequest) SetAutoRefreshPolicy(v string) *CreateDataFlowRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *CreateDataFlowRequest) SetAutoRefreshs(v []*CreateDataFlowRequestAutoRefreshs) *CreateDataFlowRequest {
	s.AutoRefreshs = v
	return s
}

func (s *CreateDataFlowRequest) SetClientToken(v string) *CreateDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataFlowRequest) SetDescription(v string) *CreateDataFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateDataFlowRequest) SetDryRun(v bool) *CreateDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataFlowRequest) SetFileSystemId(v string) *CreateDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataFlowRequest) SetFsetId(v string) *CreateDataFlowRequest {
	s.FsetId = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceSecurityType(v string) *CreateDataFlowRequest {
	s.SourceSecurityType = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceStorage(v string) *CreateDataFlowRequest {
	s.SourceStorage = &v
	return s
}

func (s *CreateDataFlowRequest) SetThroughput(v int64) *CreateDataFlowRequest {
	s.Throughput = &v
	return s
}

type CreateDataFlowRequestAutoRefreshs struct {
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s CreateDataFlowRequestAutoRefreshs) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowRequestAutoRefreshs) GoString() string {
	return s.String()
}

func (s *CreateDataFlowRequestAutoRefreshs) SetRefreshPath(v string) *CreateDataFlowRequestAutoRefreshs {
	s.RefreshPath = &v
	return s
}

type CreateDataFlowResponseBody struct {
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataFlowResponseBody) SetDataFlowId(v string) *CreateDataFlowResponseBody {
	s.DataFlowId = &v
	return s
}

func (s *CreateDataFlowResponseBody) SetRequestId(v string) *CreateDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateDataFlowResponse) SetHeaders(v map[string]*string) *CreateDataFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateDataFlowResponse) SetStatusCode(v int32) *CreateDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataFlowResponse) SetBody(v *CreateDataFlowResponseBody) *CreateDataFlowResponse {
	s.Body = v
	return s
}

type CreateDataFlowTaskRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId   *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DataType     *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Directory    *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EntryList    *string `json:"EntryList,omitempty" xml:"EntryList,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	SrcTaskId    *string `json:"SrcTaskId,omitempty" xml:"SrcTaskId,omitempty"`
	TaskAction   *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s CreateDataFlowTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskRequest) SetClientToken(v string) *CreateDataFlowTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDataFlowId(v string) *CreateDataFlowTaskRequest {
	s.DataFlowId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDataType(v string) *CreateDataFlowTaskRequest {
	s.DataType = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDirectory(v string) *CreateDataFlowTaskRequest {
	s.Directory = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDryRun(v bool) *CreateDataFlowTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetEntryList(v string) *CreateDataFlowTaskRequest {
	s.EntryList = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetFileSystemId(v string) *CreateDataFlowTaskRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetSrcTaskId(v string) *CreateDataFlowTaskRequest {
	s.SrcTaskId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetTaskAction(v string) *CreateDataFlowTaskRequest {
	s.TaskAction = &v
	return s
}

type CreateDataFlowTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDataFlowTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskResponseBody) SetRequestId(v string) *CreateDataFlowTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataFlowTaskResponseBody) SetTaskId(v string) *CreateDataFlowTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDataFlowTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDataFlowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataFlowTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataFlowTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskResponse) SetHeaders(v map[string]*string) *CreateDataFlowTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDataFlowTaskResponse) SetStatusCode(v int32) *CreateDataFlowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataFlowTaskResponse) SetBody(v *CreateDataFlowTaskResponseBody) *CreateDataFlowTaskResponse {
	s.Body = v
	return s
}

type CreateFileRequest struct {
	FileSystemId           *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Owner                  *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerAccessInheritable *bool   `json:"OwnerAccessInheritable,omitempty" xml:"OwnerAccessInheritable,omitempty"`
	Path                   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type                   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) SetFileSystemId(v string) *CreateFileRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateFileRequest) SetOwner(v string) *CreateFileRequest {
	s.Owner = &v
	return s
}

func (s *CreateFileRequest) SetOwnerAccessInheritable(v bool) *CreateFileRequest {
	s.OwnerAccessInheritable = &v
	return s
}

func (s *CreateFileRequest) SetPath(v string) *CreateFileRequest {
	s.Path = &v
	return s
}

func (s *CreateFileRequest) SetType(v string) *CreateFileRequest {
	s.Type = &v
	return s
}

type CreateFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) SetRequestId(v string) *CreateFileResponseBody {
	s.RequestId = &v
	return s
}

type CreateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponse) GoString() string {
	return s.String()
}

func (s *CreateFileResponse) SetHeaders(v map[string]*string) *CreateFileResponse {
	s.Headers = v
	return s
}

func (s *CreateFileResponse) SetStatusCode(v int32) *CreateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileResponse) SetBody(v *CreateFileResponseBody) *CreateFileResponse {
	s.Body = v
	return s
}

type CreateFileSystemRequest struct {
	Bandwidth      *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Capacity       *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Duration       *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EncryptType    *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	KmsKeyId       *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	ProtocolType   *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	StorageType    *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	VSwitchId      *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequest) SetBandwidth(v int64) *CreateFileSystemRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateFileSystemRequest) SetCapacity(v int64) *CreateFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *CreateFileSystemRequest) SetChargeType(v string) *CreateFileSystemRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateFileSystemRequest) SetClientToken(v string) *CreateFileSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFileSystemRequest) SetDescription(v string) *CreateFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateFileSystemRequest) SetDryRun(v bool) *CreateFileSystemRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFileSystemRequest) SetDuration(v int32) *CreateFileSystemRequest {
	s.Duration = &v
	return s
}

func (s *CreateFileSystemRequest) SetEncryptType(v int32) *CreateFileSystemRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateFileSystemRequest) SetFileSystemType(v string) *CreateFileSystemRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateFileSystemRequest) SetKmsKeyId(v string) *CreateFileSystemRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateFileSystemRequest) SetProtocolType(v string) *CreateFileSystemRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateFileSystemRequest) SetSnapshotId(v string) *CreateFileSystemRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageType(v string) *CreateFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateFileSystemRequest) SetVSwitchId(v string) *CreateFileSystemRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateFileSystemRequest) SetVpcId(v string) *CreateFileSystemRequest {
	s.VpcId = &v
	return s
}

func (s *CreateFileSystemRequest) SetZoneId(v string) *CreateFileSystemRequest {
	s.ZoneId = &v
	return s
}

type CreateFileSystemResponseBody struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponseBody) SetFileSystemId(v string) *CreateFileSystemResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *CreateFileSystemResponseBody) SetRequestId(v string) *CreateFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type CreateFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateFileSystemResponse) SetStatusCode(v int32) *CreateFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileSystemResponse) SetBody(v *CreateFileSystemResponseBody) *CreateFileSystemResponse {
	s.Body = v
	return s
}

type CreateFilesetRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId   *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
}

func (s CreateFilesetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFilesetRequest) GoString() string {
	return s.String()
}

func (s *CreateFilesetRequest) SetClientToken(v string) *CreateFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFilesetRequest) SetDescription(v string) *CreateFilesetRequest {
	s.Description = &v
	return s
}

func (s *CreateFilesetRequest) SetDryRun(v bool) *CreateFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFilesetRequest) SetFileSystemId(v string) *CreateFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateFilesetRequest) SetFileSystemPath(v string) *CreateFilesetRequest {
	s.FileSystemPath = &v
	return s
}

type CreateFilesetResponseBody struct {
	FsetId    *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFilesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFilesetResponseBody) SetFsetId(v string) *CreateFilesetResponseBody {
	s.FsetId = &v
	return s
}

func (s *CreateFilesetResponseBody) SetRequestId(v string) *CreateFilesetResponseBody {
	s.RequestId = &v
	return s
}

type CreateFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFilesetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFilesetResponse) GoString() string {
	return s.String()
}

func (s *CreateFilesetResponse) SetHeaders(v map[string]*string) *CreateFilesetResponse {
	s.Headers = v
	return s
}

func (s *CreateFilesetResponse) SetStatusCode(v int32) *CreateFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFilesetResponse) SetBody(v *CreateFilesetResponseBody) *CreateFilesetResponse {
	s.Body = v
	return s
}

type CreateLDAPConfigRequest struct {
	BindDN       *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	SearchBase   *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI          *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigRequest) SetBindDN(v string) *CreateLDAPConfigRequest {
	s.BindDN = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetFileSystemId(v string) *CreateLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetSearchBase(v string) *CreateLDAPConfigRequest {
	s.SearchBase = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetURI(v string) *CreateLDAPConfigRequest {
	s.URI = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateLDAPConfigResponse) SetStatusCode(v int32) *CreateLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLDAPConfigResponse) SetBody(v *CreateLDAPConfigResponseBody) *CreateLDAPConfigResponse {
	s.Body = v
	return s
}

type CreateLifecyclePolicyRequest struct {
	FileSystemId        *string   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	LifecyclePolicyName *string   `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	LifecycleRuleName   *string   `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	Path                *string   `json:"Path,omitempty" xml:"Path,omitempty"`
	Paths               []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	StorageType         *string   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
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

func (s *CreateLifecyclePolicyRequest) SetLifecycleRuleName(v string) *CreateLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPath(v string) *CreateLifecyclePolicyRequest {
	s.Path = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPaths(v []*string) *CreateLifecyclePolicyRequest {
	s.Paths = v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetStorageType(v string) *CreateLifecyclePolicyRequest {
	s.StorageType = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateLifecyclePolicyResponse) SetStatusCode(v int32) *CreateLifecyclePolicyResponse {
	s.StatusCode = &v
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
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLifecycleRetrieveJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLifecycleRetrieveJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobResponseBody) SetJobId(v string) *CreateLifecycleRetrieveJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateLifecycleRetrieveJobResponseBody) SetRequestId(v string) *CreateLifecycleRetrieveJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateLifecycleRetrieveJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateLifecycleRetrieveJobResponse) SetStatusCode(v int32) *CreateLifecycleRetrieveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLifecycleRetrieveJobResponse) SetBody(v *CreateLifecycleRetrieveJobResponseBody) *CreateLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

type CreateMountTargetRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	DryRun          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EnableIpv6      *bool   `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	FileSystemId    *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateMountTargetRequest) SetAccessGroupName(v string) *CreateMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateMountTargetRequest) SetDryRun(v bool) *CreateMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateMountTargetRequest) SetEnableIpv6(v bool) *CreateMountTargetRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateMountTargetRequest) SetFileSystemId(v string) *CreateMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateMountTargetRequest) SetNetworkType(v string) *CreateMountTargetRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateMountTargetRequest) SetSecurityGroupId(v string) *CreateMountTargetRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMountTargetRequest) SetVSwitchId(v string) *CreateMountTargetRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMountTargetRequest) SetVpcId(v string) *CreateMountTargetRequest {
	s.VpcId = &v
	return s
}

type CreateMountTargetResponseBody struct {
	MountTargetDomain *string                                        `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	MountTargetExtra  *CreateMountTargetResponseBodyMountTargetExtra `json:"MountTargetExtra,omitempty" xml:"MountTargetExtra,omitempty" type:"Struct"`
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponseBody) SetMountTargetDomain(v string) *CreateMountTargetResponseBody {
	s.MountTargetDomain = &v
	return s
}

func (s *CreateMountTargetResponseBody) SetMountTargetExtra(v *CreateMountTargetResponseBodyMountTargetExtra) *CreateMountTargetResponseBody {
	s.MountTargetExtra = v
	return s
}

func (s *CreateMountTargetResponseBody) SetRequestId(v string) *CreateMountTargetResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateMountTargetResponse) SetStatusCode(v int32) *CreateMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMountTargetResponse) SetBody(v *CreateMountTargetResponseBody) *CreateMountTargetResponse {
	s.Body = v
	return s
}

type CreateProtocolMountTargetRequest struct {
	AccessGroupName   *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FsetId            *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	VSwitchId         *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateProtocolMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetRequest) SetAccessGroupName(v string) *CreateProtocolMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetClientToken(v string) *CreateProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetDescription(v string) *CreateProtocolMountTargetRequest {
	s.Description = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetDryRun(v bool) *CreateProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetFileSystemId(v string) *CreateProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetFsetId(v string) *CreateProtocolMountTargetRequest {
	s.FsetId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetPath(v string) *CreateProtocolMountTargetRequest {
	s.Path = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetProtocolServiceId(v string) *CreateProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetVSwitchId(v string) *CreateProtocolMountTargetRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetVpcId(v string) *CreateProtocolMountTargetRequest {
	s.VpcId = &v
	return s
}

type CreateProtocolMountTargetResponseBody struct {
	ExportId  *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProtocolMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetResponseBody) SetExportId(v string) *CreateProtocolMountTargetResponseBody {
	s.ExportId = &v
	return s
}

func (s *CreateProtocolMountTargetResponseBody) SetRequestId(v string) *CreateProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type CreateProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProtocolMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetResponse) SetHeaders(v map[string]*string) *CreateProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateProtocolMountTargetResponse) SetStatusCode(v int32) *CreateProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtocolMountTargetResponse) SetBody(v *CreateProtocolMountTargetResponseBody) *CreateProtocolMountTargetResponse {
	s.Body = v
	return s
}

type CreateProtocolServiceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProtocolSpec *string `json:"ProtocolSpec,omitempty" xml:"ProtocolSpec,omitempty"`
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Throughput   *int32  `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	VSwitchId    *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateProtocolServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceRequest) SetClientToken(v string) *CreateProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetDescription(v string) *CreateProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetDryRun(v bool) *CreateProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetFileSystemId(v string) *CreateProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetProtocolSpec(v string) *CreateProtocolServiceRequest {
	s.ProtocolSpec = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetProtocolType(v string) *CreateProtocolServiceRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetThroughput(v int32) *CreateProtocolServiceRequest {
	s.Throughput = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetVSwitchId(v string) *CreateProtocolServiceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetVpcId(v string) *CreateProtocolServiceRequest {
	s.VpcId = &v
	return s
}

type CreateProtocolServiceResponseBody struct {
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProtocolServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceResponseBody) SetProtocolServiceId(v string) *CreateProtocolServiceResponseBody {
	s.ProtocolServiceId = &v
	return s
}

func (s *CreateProtocolServiceResponseBody) SetRequestId(v string) *CreateProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProtocolServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceResponse) SetHeaders(v map[string]*string) *CreateProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateProtocolServiceResponse) SetStatusCode(v int32) *CreateProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtocolServiceResponse) SetBody(v *CreateProtocolServiceResponseBody) *CreateProtocolServiceResponse {
	s.Body = v
	return s
}

type CreateRecycleBinDeleteJobRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s CreateRecycleBinDeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinDeleteJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobRequest) SetClientToken(v string) *CreateRecycleBinDeleteJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) SetFileId(v string) *CreateRecycleBinDeleteJobRequest {
	s.FileId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) SetFileSystemId(v string) *CreateRecycleBinDeleteJobRequest {
	s.FileSystemId = &v
	return s
}

type CreateRecycleBinDeleteJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecycleBinDeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinDeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobResponseBody) SetJobId(v string) *CreateRecycleBinDeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobResponseBody) SetRequestId(v string) *CreateRecycleBinDeleteJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateRecycleBinDeleteJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRecycleBinDeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRecycleBinDeleteJobResponse) SetStatusCode(v int32) *CreateRecycleBinDeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecycleBinDeleteJobResponse) SetBody(v *CreateRecycleBinDeleteJobResponseBody) *CreateRecycleBinDeleteJobResponse {
	s.Body = v
	return s
}

type CreateRecycleBinRestoreJobRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	TargetFileId *string `json:"TargetFileId,omitempty" xml:"TargetFileId,omitempty"`
}

func (s CreateRecycleBinRestoreJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobRequest) SetClientToken(v string) *CreateRecycleBinRestoreJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetFileId(v string) *CreateRecycleBinRestoreJobRequest {
	s.FileId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetFileSystemId(v string) *CreateRecycleBinRestoreJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetTargetFileId(v string) *CreateRecycleBinRestoreJobRequest {
	s.TargetFileId = &v
	return s
}

type CreateRecycleBinRestoreJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecycleBinRestoreJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecycleBinRestoreJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobResponseBody) SetJobId(v string) *CreateRecycleBinRestoreJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobResponseBody) SetRequestId(v string) *CreateRecycleBinRestoreJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateRecycleBinRestoreJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRecycleBinRestoreJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRecycleBinRestoreJobResponse) SetStatusCode(v int32) *CreateRecycleBinRestoreJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecycleBinRestoreJobResponse) SetBody(v *CreateRecycleBinRestoreJobResponseBody) *CreateRecycleBinRestoreJobResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	RetentionDays *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SnapshotName  *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetFileSystemId(v string) *CreateSnapshotRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRetentionDays(v int32) *CreateSnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

type CreateSnapshotResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSnapshotResponse) SetStatusCode(v int32) *CreateSnapshotResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAccessGroupResponse) SetStatusCode(v int32) *DeleteAccessGroupResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAccessRuleResponse) SetStatusCode(v int32) *DeleteAccessRuleResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAutoSnapshotPolicyResponse) SetStatusCode(v int32) *DeleteAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyResponse) SetBody(v *DeleteAutoSnapshotPolicyResponseBody) *DeleteAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type DeleteDataFlowRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId   *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowRequest) SetClientToken(v string) *DeleteDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDataFlowRequest) SetDataFlowId(v string) *DeleteDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *DeleteDataFlowRequest) SetDryRun(v bool) *DeleteDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteDataFlowRequest) SetFileSystemId(v string) *DeleteDataFlowRequest {
	s.FileSystemId = &v
	return s
}

type DeleteDataFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowResponseBody) SetRequestId(v string) *DeleteDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowResponse) SetHeaders(v map[string]*string) *DeleteDataFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataFlowResponse) SetStatusCode(v int32) *DeleteDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataFlowResponse) SetBody(v *DeleteDataFlowResponseBody) *DeleteDataFlowResponse {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFileSystemResponse) SetStatusCode(v int32) *DeleteFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileSystemResponse) SetBody(v *DeleteFileSystemResponseBody) *DeleteFileSystemResponse {
	s.Body = v
	return s
}

type DeleteFilesetRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FsetId       *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s DeleteFilesetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesetRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilesetRequest) SetClientToken(v string) *DeleteFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFilesetRequest) SetDryRun(v bool) *DeleteFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteFilesetRequest) SetFileSystemId(v string) *DeleteFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFilesetRequest) SetFsetId(v string) *DeleteFilesetRequest {
	s.FsetId = &v
	return s
}

type DeleteFilesetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFilesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilesetResponseBody) SetRequestId(v string) *DeleteFilesetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFilesetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilesetResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilesetResponse) SetHeaders(v map[string]*string) *DeleteFilesetResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilesetResponse) SetStatusCode(v int32) *DeleteFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFilesetResponse) SetBody(v *DeleteFilesetResponseBody) *DeleteFilesetResponse {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteLDAPConfigResponse) SetStatusCode(v int32) *DeleteLDAPConfigResponse {
	s.StatusCode = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLifecyclePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyResponseBody) SetRequestId(v string) *DeleteLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLifecyclePolicyResponseBody) SetSuccess(v bool) *DeleteLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

type DeleteLifecyclePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteLifecyclePolicyResponse) SetStatusCode(v int32) *DeleteLifecyclePolicyResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteMountTargetResponse) SetStatusCode(v int32) *DeleteMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMountTargetResponse) SetBody(v *DeleteMountTargetResponseBody) *DeleteMountTargetResponse {
	s.Body = v
	return s
}

type DeleteProtocolMountTargetRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ExportId          *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s DeleteProtocolMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetRequest) SetClientToken(v string) *DeleteProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetDryRun(v bool) *DeleteProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetExportId(v string) *DeleteProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetFileSystemId(v string) *DeleteProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetProtocolServiceId(v string) *DeleteProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

type DeleteProtocolMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProtocolMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetResponseBody) SetRequestId(v string) *DeleteProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProtocolMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetResponse) SetHeaders(v map[string]*string) *DeleteProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtocolMountTargetResponse) SetStatusCode(v int32) *DeleteProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtocolMountTargetResponse) SetBody(v *DeleteProtocolMountTargetResponseBody) *DeleteProtocolMountTargetResponse {
	s.Body = v
	return s
}

type DeleteProtocolServiceRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s DeleteProtocolServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceRequest) SetClientToken(v string) *DeleteProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetDryRun(v bool) *DeleteProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetFileSystemId(v string) *DeleteProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetProtocolServiceId(v string) *DeleteProtocolServiceRequest {
	s.ProtocolServiceId = &v
	return s
}

type DeleteProtocolServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProtocolServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceResponseBody) SetRequestId(v string) *DeleteProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProtocolServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceResponse) SetHeaders(v map[string]*string) *DeleteProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtocolServiceResponse) SetStatusCode(v int32) *DeleteProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtocolServiceResponse) SetBody(v *DeleteProtocolServiceResponseBody) *DeleteProtocolServiceResponse {
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

type DescribeAccessGroupsRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	FileSystemType  *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UseUTCDateTime  *bool   `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
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

func (s *DescribeAccessGroupsRequest) SetFileSystemType(v string) *DescribeAccessGroupsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetPageNumber(v int32) *DescribeAccessGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetPageSize(v int32) *DescribeAccessGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessGroupsRequest) SetUseUTCDateTime(v bool) *DescribeAccessGroupsRequest {
	s.UseUTCDateTime = &v
	return s
}

type DescribeAccessGroupsResponseBody struct {
	AccessGroups *DescribeAccessGroupsResponseBodyAccessGroups `json:"AccessGroups,omitempty" xml:"AccessGroups,omitempty" type:"Struct"`
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBody) SetAccessGroups(v *DescribeAccessGroupsResponseBodyAccessGroups) *DescribeAccessGroupsResponseBody {
	s.AccessGroups = v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetPageNumber(v int32) *DescribeAccessGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetPageSize(v int32) *DescribeAccessGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetRequestId(v string) *DescribeAccessGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetTotalCount(v int32) *DescribeAccessGroupsResponseBody {
	s.TotalCount = &v
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
	AccessGroupType  *string `json:"AccessGroupType,omitempty" xml:"AccessGroupType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MountTargetCount *int32  `json:"MountTargetCount,omitempty" xml:"MountTargetCount,omitempty"`
	RuleCount        *int32  `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
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

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetAccessGroupType(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.AccessGroupType = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetDescription(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.Description = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetMountTargetCount(v int32) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.MountTargetCount = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetRuleCount(v int32) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.RuleCount = &v
	return s
}

type DescribeAccessGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAccessGroupsResponse) SetStatusCode(v int32) *DescribeAccessGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessGroupsResponse) SetBody(v *DescribeAccessGroupsResponseBody) *DescribeAccessGroupsResponse {
	s.Body = v
	return s
}

type DescribeAccessRulesRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	AccessRuleId    *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	FileSystemType  *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeAccessRulesRequest) SetFileSystemType(v string) *DescribeAccessRulesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetPageNumber(v int32) *DescribeAccessRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessRulesRequest) SetPageSize(v int32) *DescribeAccessRulesRequest {
	s.PageSize = &v
	return s
}

type DescribeAccessRulesResponseBody struct {
	AccessRules *DescribeAccessRulesResponseBodyAccessRules `json:"AccessRules,omitempty" xml:"AccessRules,omitempty" type:"Struct"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessRulesResponseBody) SetAccessRules(v *DescribeAccessRulesResponseBodyAccessRules) *DescribeAccessRulesResponseBody {
	s.AccessRules = v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetPageNumber(v int32) *DescribeAccessRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetPageSize(v int32) *DescribeAccessRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetRequestId(v string) *DescribeAccessRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessRulesResponseBody) SetTotalCount(v int32) *DescribeAccessRulesResponseBody {
	s.TotalCount = &v
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
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	Priority         *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RWAccess         *string `json:"RWAccess,omitempty" xml:"RWAccess,omitempty"`
	SourceCidrIp     *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	UserAccess       *string `json:"UserAccess,omitempty" xml:"UserAccess,omitempty"`
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

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetIpv6SourceCidrIp(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetPriority(v int32) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.Priority = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetRWAccess(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.RWAccess = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetSourceCidrIp(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeAccessRulesResponseBodyAccessRulesAccessRule) SetUserAccess(v string) *DescribeAccessRulesResponseBodyAccessRulesAccessRule {
	s.UserAccess = &v
	return s
}

type DescribeAccessRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccessRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAccessRulesResponse) SetStatusCode(v int32) *DescribeAccessRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessRulesResponse) SetBody(v *DescribeAccessRulesResponseBody) *DescribeAccessRulesResponse {
	s.Body = v
	return s
}

type DescribeAutoSnapshotPoliciesRequest struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	FileSystemType       *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeAutoSnapshotPoliciesRequest) SetFileSystemType(v string) *DescribeAutoSnapshotPoliciesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetPageSize(v int32) *DescribeAutoSnapshotPoliciesRequest {
	s.PageSize = &v
	return s
}

type DescribeAutoSnapshotPoliciesResponseBody struct {
	AutoSnapshotPolicies *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies `json:"AutoSnapshotPolicies,omitempty" xml:"AutoSnapshotPolicies,omitempty" type:"Struct"`
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetAutoSnapshotPolicies(v *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPoliciesResponseBody {
	s.AutoSnapshotPolicies = v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetRequestId(v string) *DescribeAutoSnapshotPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.TotalCount = &v
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
	AutoSnapshotPolicyId   *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileSystemNums         *int32  `json:"FileSystemNums,omitempty" xml:"FileSystemNums,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepeatWeekdays         *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	RetentionDays          *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TimePoints             *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyId = &v
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

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetFileSystemNums(v int32) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.FileSystemNums = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRegionId(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRepeatWeekdays(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RepeatWeekdays = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRetentionDays(v int32) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RetentionDays = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetStatus(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.Status = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetTimePoints(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.TimePoints = &v
	return s
}

type DescribeAutoSnapshotPoliciesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutoSnapshotPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAutoSnapshotPoliciesResponse) SetStatusCode(v int32) *DescribeAutoSnapshotPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponse) SetBody(v *DescribeAutoSnapshotPoliciesResponseBody) *DescribeAutoSnapshotPoliciesResponse {
	s.Body = v
	return s
}

type DescribeAutoSnapshotTasksRequest struct {
	AutoSnapshotPolicyIds *string `json:"AutoSnapshotPolicyIds,omitempty" xml:"AutoSnapshotPolicyIds,omitempty"`
	FileSystemIds         *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
	FileSystemType        *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAutoSnapshotTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksRequest) SetAutoSnapshotPolicyIds(v string) *DescribeAutoSnapshotTasksRequest {
	s.AutoSnapshotPolicyIds = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetFileSystemIds(v string) *DescribeAutoSnapshotTasksRequest {
	s.FileSystemIds = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetFileSystemType(v string) *DescribeAutoSnapshotTasksRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetPageNumber(v int32) *DescribeAutoSnapshotTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetPageSize(v int32) *DescribeAutoSnapshotTasksRequest {
	s.PageSize = &v
	return s
}

type DescribeAutoSnapshotTasksResponseBody struct {
	AutoSnapshotTasks *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks `json:"AutoSnapshotTasks,omitempty" xml:"AutoSnapshotTasks,omitempty" type:"Struct"`
	PageNumber        *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoSnapshotTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoSnapshotTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetAutoSnapshotTasks(v *DescribeAutoSnapshotTasksResponseBodyAutoSnapshotTasks) *DescribeAutoSnapshotTasksResponseBody {
	s.AutoSnapshotTasks = v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetRequestId(v string) *DescribeAutoSnapshotTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotTasksResponseBody {
	s.TotalCount = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutoSnapshotTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAutoSnapshotTasksResponse) SetStatusCode(v int32) *DescribeAutoSnapshotTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoSnapshotTasksResponse) SetBody(v *DescribeAutoSnapshotTasksResponseBody) *DescribeAutoSnapshotTasksResponse {
	s.Body = v
	return s
}

type DescribeBlackListClientsRequest struct {
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBlackListClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsRequest) SetClientIP(v string) *DescribeBlackListClientsRequest {
	s.ClientIP = &v
	return s
}

func (s *DescribeBlackListClientsRequest) SetFileSystemId(v string) *DescribeBlackListClientsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeBlackListClientsRequest) SetRegionId(v string) *DescribeBlackListClientsRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBlackListClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBlackListClientsResponse) SetStatusCode(v int32) *DescribeBlackListClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBlackListClientsResponse) SetBody(v *DescribeBlackListClientsResponseBody) *DescribeBlackListClientsResponse {
	s.Body = v
	return s
}

type DescribeDataFlowTasksRequest struct {
	FileSystemId *string                                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*DescribeDataFlowTasksRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults   *int64                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDataFlowTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksRequest) SetFileSystemId(v string) *DescribeDataFlowTasksRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetFilters(v []*DescribeDataFlowTasksRequestFilters) *DescribeDataFlowTasksRequest {
	s.Filters = v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetMaxResults(v int64) *DescribeDataFlowTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetNextToken(v string) *DescribeDataFlowTasksRequest {
	s.NextToken = &v
	return s
}

type DescribeDataFlowTasksRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataFlowTasksRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksRequestFilters) SetKey(v string) *DescribeDataFlowTasksRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeDataFlowTasksRequestFilters) SetValue(v string) *DescribeDataFlowTasksRequestFilters {
	s.Value = &v
	return s
}

type DescribeDataFlowTasksResponseBody struct {
	NextToken *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfo  *DescribeDataFlowTasksResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DescribeDataFlowTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBody) SetNextToken(v string) *DescribeDataFlowTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBody) SetRequestId(v string) *DescribeDataFlowTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBody) SetTaskInfo(v *DescribeDataFlowTasksResponseBodyTaskInfo) *DescribeDataFlowTasksResponseBody {
	s.TaskInfo = v
	return s
}

type DescribeDataFlowTasksResponseBodyTaskInfo struct {
	Task []*DescribeDataFlowTasksResponseBodyTaskInfoTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfo) SetTask(v []*DescribeDataFlowTasksResponseBodyTaskInfoTask) *DescribeDataFlowTasksResponseBodyTaskInfo {
	s.Task = v
	return s
}

type DescribeDataFlowTasksResponseBodyTaskInfoTask struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataFlowId     *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	FilesystemId   *string `json:"FilesystemId,omitempty" xml:"FilesystemId,omitempty"`
	FsPath         *string `json:"FsPath,omitempty" xml:"FsPath,omitempty"`
	Originator     *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	Progress       *int64  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ReportPath     *string `json:"ReportPath,omitempty" xml:"ReportPath,omitempty"`
	SourceStorage  *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskAction     *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTask) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetCreateTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.CreateTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetDataFlowId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.DataFlowId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetDataType(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.DataType = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetEndTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.EndTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFileSystemPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFilesystemId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FilesystemId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFsPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FsPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetOriginator(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Originator = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetProgress(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Progress = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetReportPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.ReportPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetSourceStorage(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.SourceStorage = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetStartTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.StartTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetStatus(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Status = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetTaskAction(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.TaskAction = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetTaskId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.TaskId = &v
	return s
}

type DescribeDataFlowTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataFlowTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataFlowTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponse) SetHeaders(v map[string]*string) *DescribeDataFlowTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataFlowTasksResponse) SetStatusCode(v int32) *DescribeDataFlowTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataFlowTasksResponse) SetBody(v *DescribeDataFlowTasksResponseBody) *DescribeDataFlowTasksResponse {
	s.Body = v
	return s
}

type DescribeDataFlowsRequest struct {
	FileSystemId *string                            `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*DescribeDataFlowsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults   *int64                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDataFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsRequest) SetFileSystemId(v string) *DescribeDataFlowsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowsRequest) SetFilters(v []*DescribeDataFlowsRequestFilters) *DescribeDataFlowsRequest {
	s.Filters = v
	return s
}

func (s *DescribeDataFlowsRequest) SetMaxResults(v int64) *DescribeDataFlowsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDataFlowsRequest) SetNextToken(v string) *DescribeDataFlowsRequest {
	s.NextToken = &v
	return s
}

type DescribeDataFlowsRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataFlowsRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsRequestFilters) SetKey(v string) *DescribeDataFlowsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeDataFlowsRequestFilters) SetValue(v string) *DescribeDataFlowsRequestFilters {
	s.Value = &v
	return s
}

type DescribeDataFlowsResponseBody struct {
	DataFlowInfo *DescribeDataFlowsResponseBodyDataFlowInfo `json:"DataFlowInfo,omitempty" xml:"DataFlowInfo,omitempty" type:"Struct"`
	NextToken    *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBody) SetDataFlowInfo(v *DescribeDataFlowsResponseBodyDataFlowInfo) *DescribeDataFlowsResponseBody {
	s.DataFlowInfo = v
	return s
}

func (s *DescribeDataFlowsResponseBody) SetNextToken(v string) *DescribeDataFlowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowsResponseBody) SetRequestId(v string) *DescribeDataFlowsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDataFlowsResponseBodyDataFlowInfo struct {
	DataFlow []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlow `json:"DataFlow,omitempty" xml:"DataFlow,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfo) SetDataFlow(v []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) *DescribeDataFlowsResponseBodyDataFlowInfo {
	s.DataFlow = v
	return s
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlow struct {
	AutoRefresh         *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh `json:"AutoRefresh,omitempty" xml:"AutoRefresh,omitempty" type:"Struct"`
	AutoRefreshInterval *int64                                                        `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	AutoRefreshPolicy   *string                                                       `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	CreateTime          *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataFlowId          *string                                                       `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	Description         *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorMessage        *string                                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FileSystemId        *string                                                       `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemPath      *string                                                       `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	FsetDescription     *string                                                       `json:"FsetDescription,omitempty" xml:"FsetDescription,omitempty"`
	FsetId              *string                                                       `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	SourceSecurityType  *string                                                       `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	SourceStorage       *string                                                       `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	Status              *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Throughput          *int64                                                        `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	UpdateTime          *string                                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetAutoRefresh(v *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.AutoRefresh = v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetAutoRefreshInterval(v int64) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.AutoRefreshInterval = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetAutoRefreshPolicy(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetCreateTime(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.CreateTime = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetDataFlowId(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.DataFlowId = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetDescription(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.Description = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetErrorMessage(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetFileSystemId(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetFileSystemPath(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetFsetDescription(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.FsetDescription = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetFsetId(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.FsetId = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetSourceSecurityType(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.SourceSecurityType = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetSourceStorage(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.SourceStorage = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetStatus(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.Status = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetThroughput(v int64) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.Throughput = &v
	return s
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow) SetUpdateTime(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlow {
	s.UpdateTime = &v
	return s
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh struct {
	AutoRefresh []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh `json:"AutoRefresh,omitempty" xml:"AutoRefresh,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh) SetAutoRefresh(v []*DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefresh {
	s.AutoRefresh = v
	return s
}

type DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh struct {
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh) SetRefreshPath(v string) *DescribeDataFlowsResponseBodyDataFlowInfoDataFlowAutoRefreshAutoRefresh {
	s.RefreshPath = &v
	return s
}

type DescribeDataFlowsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataFlowsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsResponse) SetHeaders(v map[string]*string) *DescribeDataFlowsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataFlowsResponse) SetStatusCode(v int32) *DescribeDataFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataFlowsResponse) SetBody(v *DescribeDataFlowsResponseBody) *DescribeDataFlowsResponse {
	s.Body = v
	return s
}

type DescribeDirQuotasRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
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

func (s *DescribeDirQuotasRequest) SetPageNumber(v int32) *DescribeDirQuotasRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPageSize(v int32) *DescribeDirQuotasRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPath(v string) *DescribeDirQuotasRequest {
	s.Path = &v
	return s
}

type DescribeDirQuotasResponseBody struct {
	DirQuotaInfos []*DescribeDirQuotasResponseBodyDirQuotaInfos `json:"DirQuotaInfos,omitempty" xml:"DirQuotaInfos,omitempty" type:"Repeated"`
	PageNumber    *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDirQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBody) SetDirQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfos) *DescribeDirQuotasResponseBody {
	s.DirQuotaInfos = v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetPageNumber(v int32) *DescribeDirQuotasResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetPageSize(v int32) *DescribeDirQuotasResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetRequestId(v string) *DescribeDirQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirQuotasResponseBody) SetTotalCount(v int32) *DescribeDirQuotasResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDirQuotasResponseBodyDirQuotaInfos struct {
	DirInode       *string                                                     `json:"DirInode,omitempty" xml:"DirInode,omitempty"`
	Path           *string                                                     `json:"Path,omitempty" xml:"Path,omitempty"`
	Status         *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	UserQuotaInfos []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos `json:"UserQuotaInfos,omitempty" xml:"UserQuotaInfos,omitempty" type:"Repeated"`
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfos) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetDirInode(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.DirInode = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetPath(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.Path = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetStatus(v string) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.Status = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfos) SetUserQuotaInfos(v []*DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) *DescribeDirQuotasResponseBodyDirQuotaInfos {
	s.UserQuotaInfos = v
	return s
}

type DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos struct {
	FileCountLimit *int64  `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	FileCountReal  *int64  `json:"FileCountReal,omitempty" xml:"FileCountReal,omitempty"`
	QuotaType      *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	SizeLimit      *int64  `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	SizeReal       *int64  `json:"SizeReal,omitempty" xml:"SizeReal,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType       *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetFileCountLimit(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.FileCountLimit = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetFileCountReal(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.FileCountReal = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetQuotaType(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.QuotaType = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeLimit(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeLimit = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetSizeReal(v int64) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.SizeReal = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetUserId(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.UserId = &v
	return s
}

func (s *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos) SetUserType(v string) *DescribeDirQuotasResponseBodyDirQuotaInfosUserQuotaInfos {
	s.UserType = &v
	return s
}

type DescribeDirQuotasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDirQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDirQuotasResponse) SetStatusCode(v int32) *DescribeDirQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDirQuotasResponse) SetBody(v *DescribeDirQuotasResponseBody) *DescribeDirQuotasResponse {
	s.Body = v
	return s
}

type DescribeFileSystemStatisticsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeFileSystemStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsRequest) SetPageNumber(v int32) *DescribeFileSystemStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemStatisticsRequest) SetPageSize(v int32) *DescribeFileSystemStatisticsRequest {
	s.PageSize = &v
	return s
}

type DescribeFileSystemStatisticsResponseBody struct {
	FileSystemStatistics *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics `json:"FileSystemStatistics,omitempty" xml:"FileSystemStatistics,omitempty" type:"Struct"`
	FileSystems          *DescribeFileSystemStatisticsResponseBodyFileSystems          `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBody) SetFileSystemStatistics(v *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) *DescribeFileSystemStatisticsResponseBody {
	s.FileSystemStatistics = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetFileSystems(v *DescribeFileSystemStatisticsResponseBodyFileSystems) *DescribeFileSystemStatisticsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetPageNumber(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetPageSize(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetRequestId(v string) *DescribeFileSystemStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetTotalCount(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.TotalCount = &v
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
	ExpiredCount   *int32  `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
	ExpiringCount  *int32  `json:"ExpiringCount,omitempty" xml:"ExpiringCount,omitempty"`
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredSize    *int64  `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	TotalCount     *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetExpiredCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.ExpiredCount = &v
	return s
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
	Capacity       *int64                                                                 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType     *string                                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime     *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime    *string                                                                `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FileSystemId   *string                                                                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType *string                                                                `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredIASize  *int64                                                                 `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	MeteredSize    *int64                                                                 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	Packages       *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	ProtocolType   *string                                                                `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId       *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status         *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType    *string                                                                `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ZoneId         *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetChargeType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
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

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetPackages(v *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Packages = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
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
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetExpiredTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetStartTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.StartTime = &v
	return s
}

type DescribeFileSystemStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFileSystemStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeFileSystemStatisticsResponse) SetStatusCode(v int32) *DescribeFileSystemStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponse) SetBody(v *DescribeFileSystemStatisticsResponseBody) *DescribeFileSystemStatisticsResponse {
	s.Body = v
	return s
}

type DescribeFileSystemsRequest struct {
	FileSystemId   *string                          `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType *string                          `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageNumber     *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tag            []*DescribeFileSystemsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VpcId          *string                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *DescribeFileSystemsRequest) SetPageNumber(v int32) *DescribeFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetPageSize(v int32) *DescribeFileSystemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsRequest) SetTag(v []*DescribeFileSystemsRequestTag) *DescribeFileSystemsRequest {
	s.Tag = v
	return s
}

func (s *DescribeFileSystemsRequest) SetVpcId(v string) *DescribeFileSystemsRequest {
	s.VpcId = &v
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
	FileSystems *DescribeFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBody) SetFileSystems(v *DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageNumber(v int32) *DescribeFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageSize(v int32) *DescribeFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetRequestId(v string) *DescribeFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetTotalCount(v int32) *DescribeFileSystemsResponseBody {
	s.TotalCount = &v
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
	Bandwidth         *int64                                                                 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Capacity          *int64                                                                 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType        *string                                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime        *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptType       *int32                                                                 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	ExpiredTime       *string                                                                `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FileSystemId      *string                                                                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType    *string                                                                `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	GuiInfo           *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo           `json:"GuiInfo,omitempty" xml:"GuiInfo,omitempty" type:"Struct"`
	KMSKeyId          *string                                                                `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	Ldap              *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap              `json:"Ldap,omitempty" xml:"Ldap,omitempty" type:"Struct"`
	MeteredIASize     *int64                                                                 `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	MeteredSize       *int64                                                                 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	MountTargets      *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets      `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Struct"`
	NodeNum           *int32                                                                 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Packages          *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages          `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	ProtocolType      *string                                                                `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId          *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status            *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType       *string                                                                `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SupportedFeatures *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures `json:"SupportedFeatures,omitempty" xml:"SupportedFeatures,omitempty" type:"Struct"`
	Tags              *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Version           *string                                                                `json:"Version,omitempty" xml:"Version,omitempty"`
	ZoneId            *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetBandwidth(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Bandwidth = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetChargeType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetDescription(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetEncryptType(v int32) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.EncryptType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetExpiredTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetGuiInfo(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.GuiInfo = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetKMSKeyId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetLdap(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Ldap = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetMountTargets(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargets) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.MountTargets = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetNodeNum(v int32) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.NodeNum = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetPackages(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackages) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Packages = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetSupportedFeatures(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemSupportedFeatures) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.SupportedFeatures = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetTags(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemTags) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Tags = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetVersion(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.Version = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	User     *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) SetEndpoint(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo {
	s.Endpoint = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) SetPassword(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo {
	s.Password = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo) SetUser(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemGuiInfo {
	s.User = &v
	return s
}

type DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap struct {
	BindDN     *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI        *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetBindDN(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.BindDN = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetSearchBase(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.SearchBase = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap) SetURI(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemLdap {
	s.URI = &v
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
	AccessGroupName            *string                                                                                       `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	ClientMasterNodes          *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
	DualStackMountTargetDomain *string                                                                                       `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	MountTargetDomain          *string                                                                                       `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NetworkType                *string                                                                                       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Status                     *string                                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                       *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcId                      *string                                                                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId                      *string                                                                                       `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetAccessGroupName(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetClientMasterNodes(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodes) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.ClientMasterNodes = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetDualStackMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetNetworkType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.NetworkType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetTags(v *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetTags) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.Tags = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetVpcId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.VpcId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget) SetVswId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTarget {
	s.VswId = &v
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
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
	EcsId         *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	EcsIp         *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetDefaultPasswd(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.DefaultPasswd = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsIp(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsIp = &v
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
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetExpiredTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageId(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageType(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetSize(v int64) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage) SetStartTime(v string) *DescribeFileSystemsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.StartTime = &v
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

type DescribeFileSystemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeFileSystemsResponse) SetStatusCode(v int32) *DescribeFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileSystemsResponse) SetBody(v *DescribeFileSystemsResponseBody) *DescribeFileSystemsResponse {
	s.Body = v
	return s
}

type DescribeFilesetsRequest struct {
	FileSystemId *string                           `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*DescribeFilesetsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults   *int64                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeFilesetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsRequest) SetFileSystemId(v string) *DescribeFilesetsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesetsRequest) SetFilters(v []*DescribeFilesetsRequestFilters) *DescribeFilesetsRequest {
	s.Filters = v
	return s
}

func (s *DescribeFilesetsRequest) SetMaxResults(v int64) *DescribeFilesetsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFilesetsRequest) SetNextToken(v string) *DescribeFilesetsRequest {
	s.NextToken = &v
	return s
}

type DescribeFilesetsRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFilesetsRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsRequestFilters) SetKey(v string) *DescribeFilesetsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeFilesetsRequestFilters) SetValue(v string) *DescribeFilesetsRequestFilters {
	s.Value = &v
	return s
}

type DescribeFilesetsResponseBody struct {
	Entries      *DescribeFilesetsResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Struct"`
	FileSystemId *string                              `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	NextToken    *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFilesetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBody) SetEntries(v *DescribeFilesetsResponseBodyEntries) *DescribeFilesetsResponseBody {
	s.Entries = v
	return s
}

func (s *DescribeFilesetsResponseBody) SetFileSystemId(v string) *DescribeFilesetsResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesetsResponseBody) SetNextToken(v string) *DescribeFilesetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFilesetsResponseBody) SetRequestId(v string) *DescribeFilesetsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFilesetsResponseBodyEntries struct {
	Entrie []*DescribeFilesetsResponseBodyEntriesEntrie `json:"Entrie,omitempty" xml:"Entrie,omitempty" type:"Repeated"`
}

func (s DescribeFilesetsResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBodyEntries) SetEntrie(v []*DescribeFilesetsResponseBodyEntriesEntrie) *DescribeFilesetsResponseBodyEntries {
	s.Entrie = v
	return s
}

type DescribeFilesetsResponseBodyEntriesEntrie struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	FsetId         *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeFilesetsResponseBodyEntriesEntrie) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsResponseBodyEntriesEntrie) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetCreateTime(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.CreateTime = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetDescription(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.Description = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetFileSystemPath(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetFsetId(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.FsetId = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetStatus(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.Status = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetUpdateTime(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.UpdateTime = &v
	return s
}

type DescribeFilesetsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFilesetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFilesetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFilesetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponse) SetHeaders(v map[string]*string) *DescribeFilesetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFilesetsResponse) SetStatusCode(v int32) *DescribeFilesetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFilesetsResponse) SetBody(v *DescribeFilesetsResponseBody) *DescribeFilesetsResponse {
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
	Ldap      *DescribeLDAPConfigResponseBodyLdap `json:"Ldap,omitempty" xml:"Ldap,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLDAPConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigResponseBody) SetLdap(v *DescribeLDAPConfigResponseBodyLdap) *DescribeLDAPConfigResponseBody {
	s.Ldap = v
	return s
}

func (s *DescribeLDAPConfigResponseBody) SetRequestId(v string) *DescribeLDAPConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLDAPConfigResponseBodyLdap struct {
	BindDN     *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI        *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s DescribeLDAPConfigResponseBodyLdap) String() string {
	return tea.Prettify(s)
}

func (s DescribeLDAPConfigResponseBodyLdap) GoString() string {
	return s.String()
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetBindDN(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.BindDN = &v
	return s
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetSearchBase(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.SearchBase = &v
	return s
}

func (s *DescribeLDAPConfigResponseBodyLdap) SetURI(v string) *DescribeLDAPConfigResponseBodyLdap {
	s.URI = &v
	return s
}

type DescribeLDAPConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeLDAPConfigResponse) SetStatusCode(v int32) *DescribeLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLDAPConfigResponse) SetBody(v *DescribeLDAPConfigResponseBody) *DescribeLDAPConfigResponse {
	s.Body = v
	return s
}

type DescribeLifecyclePoliciesRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeLifecyclePoliciesRequest) SetPageNumber(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageSize(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageSize = &v
	return s
}

type DescribeLifecyclePoliciesResponseBody struct {
	LifecyclePolicies []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies `json:"LifecyclePolicies,omitempty" xml:"LifecyclePolicies,omitempty" type:"Repeated"`
	PageNumber        *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLifecyclePoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBody) SetLifecyclePolicies(v []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) *DescribeLifecyclePoliciesResponseBody {
	s.LifecyclePolicies = v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetPageNumber(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetPageSize(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetRequestId(v string) *DescribeLifecyclePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetTotalCount(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLifecyclePoliciesResponseBodyLifecyclePolicies struct {
	CreateTime          *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileSystemId        *string   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	LifecyclePolicyName *string   `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	LifecycleRuleName   *string   `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	Path                *string   `json:"Path,omitempty" xml:"Path,omitempty"`
	Paths               []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	StorageType         *string   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetCreateTime(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.CreateTime = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetFileSystemId(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecycleRuleName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecycleRuleName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetPath(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Path = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetPaths(v []*string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Paths = v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetStorageType(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.StorageType = &v
	return s
}

type DescribeLifecyclePoliciesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLifecyclePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeLifecyclePoliciesResponse) SetStatusCode(v int32) *DescribeLifecyclePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponse) SetBody(v *DescribeLifecyclePoliciesResponseBody) *DescribeLifecyclePoliciesResponse {
	s.Body = v
	return s
}

type DescribeLogAnalysisRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisRequest) SetPageNumber(v int32) *DescribeLogAnalysisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetPageSize(v int32) *DescribeLogAnalysisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogAnalysisRequest) SetRegionId(v string) *DescribeLogAnalysisRequest {
	s.RegionId = &v
	return s
}

type DescribeLogAnalysisResponseBody struct {
	Analyses   *DescribeLogAnalysisResponseBodyAnalyses `json:"Analyses,omitempty" xml:"Analyses,omitempty" type:"Struct"`
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLogAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogAnalysisResponseBody) SetAnalyses(v *DescribeLogAnalysisResponseBodyAnalyses) *DescribeLogAnalysisResponseBody {
	s.Analyses = v
	return s
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
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetProject(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Project = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetRegion(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.Region = &v
	return s
}

func (s *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue) SetRoleArn(v string) *DescribeLogAnalysisResponseBodyAnalysesAnalysisMetaValue {
	s.RoleArn = &v
	return s
}

type DescribeLogAnalysisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeLogAnalysisResponse) SetStatusCode(v int32) *DescribeLogAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogAnalysisResponse) SetBody(v *DescribeLogAnalysisResponseBody) *DescribeLogAnalysisResponse {
	s.Body = v
	return s
}

type DescribeMountTargetsRequest struct {
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	FileSystemId               *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain          *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	PageNumber                 *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMountTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsRequest) SetDualStackMountTargetDomain(v string) *DescribeMountTargetsRequest {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetFileSystemId(v string) *DescribeMountTargetsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetMountTargetDomain(v string) *DescribeMountTargetsRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageNumber(v int32) *DescribeMountTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsRequest) SetPageSize(v int32) *DescribeMountTargetsRequest {
	s.PageSize = &v
	return s
}

type DescribeMountTargetsResponseBody struct {
	MountTargets *DescribeMountTargetsResponseBodyMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Struct"`
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMountTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBody) SetMountTargets(v *DescribeMountTargetsResponseBodyMountTargets) *DescribeMountTargetsResponseBody {
	s.MountTargets = v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageNumber(v int32) *DescribeMountTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageSize(v int32) *DescribeMountTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetRequestId(v string) *DescribeMountTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetTotalCount(v int32) *DescribeMountTargetsResponseBody {
	s.TotalCount = &v
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
	AccessGroup                *string                                                                   `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	ClientMasterNodes          *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
	DualStackMountTargetDomain *string                                                                   `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	IPVersion                  *string                                                                   `json:"IPVersion,omitempty" xml:"IPVersion,omitempty"`
	MountTargetDomain          *string                                                                   `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NetworkType                *string                                                                   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Status                     *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId                      *string                                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId                      *string                                                                   `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTarget) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetAccessGroup(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.AccessGroup = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetClientMasterNodes(v *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.ClientMasterNodes = v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetDualStackMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetIPVersion(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.IPVersion = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetNetworkType(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.NetworkType = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetStatus(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.Status = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetVpcId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.VpcId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetVswId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.VswId = &v
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
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
	EcsId         *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	EcsIp         *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetDefaultPasswd(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.DefaultPasswd = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsIp(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsIp = &v
	return s
}

type DescribeMountTargetsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMountTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeMountTargetsResponse) SetStatusCode(v int32) *DescribeMountTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMountTargetsResponse) SetBody(v *DescribeMountTargetsResponseBody) *DescribeMountTargetsResponse {
	s.Body = v
	return s
}

type DescribeMountedClientsRequest struct {
	ClientIP          *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMountedClientsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsRequest) SetClientIP(v string) *DescribeMountedClientsRequest {
	s.ClientIP = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetFileSystemId(v string) *DescribeMountedClientsRequest {
	s.FileSystemId = &v
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

func (s *DescribeMountedClientsRequest) SetPageSize(v int32) *DescribeMountedClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMountedClientsRequest) SetRegionId(v string) *DescribeMountedClientsRequest {
	s.RegionId = &v
	return s
}

type DescribeMountedClientsResponseBody struct {
	Clients    *DescribeMountedClientsResponseBodyClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Struct"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMountedClientsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMountedClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponseBody) SetClients(v *DescribeMountedClientsResponseBodyClients) *DescribeMountedClientsResponseBody {
	s.Clients = v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetPageNumber(v int32) *DescribeMountedClientsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetPageSize(v int32) *DescribeMountedClientsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetRequestId(v string) *DescribeMountedClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountedClientsResponseBody) SetTotalCount(v int32) *DescribeMountedClientsResponseBody {
	s.TotalCount = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMountedClientsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeMountedClientsResponse) SetStatusCode(v int32) *DescribeMountedClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMountedClientsResponse) SetBody(v *DescribeMountedClientsResponseBody) *DescribeMountedClientsResponse {
	s.Body = v
	return s
}

type DescribeProtocolMountTargetRequest struct {
	ClientToken  *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileSystemId *string                                      `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*DescribeProtocolMountTargetRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults   *int64                                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeProtocolMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetRequest) SetClientToken(v string) *DescribeProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetFileSystemId(v string) *DescribeProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetFilters(v []*DescribeProtocolMountTargetRequestFilters) *DescribeProtocolMountTargetRequest {
	s.Filters = v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetMaxResults(v int64) *DescribeProtocolMountTargetRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetNextToken(v string) *DescribeProtocolMountTargetRequest {
	s.NextToken = &v
	return s
}

type DescribeProtocolMountTargetRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeProtocolMountTargetRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetRequestFilters) SetKey(v string) *DescribeProtocolMountTargetRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeProtocolMountTargetRequestFilters) SetValue(v string) *DescribeProtocolMountTargetRequestFilters {
	s.Value = &v
	return s
}

type DescribeProtocolMountTargetResponseBody struct {
	NextToken            *string                                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProtocolMountTargets []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets `json:"ProtocolMountTargets,omitempty" xml:"ProtocolMountTargets,omitempty" type:"Repeated"`
	RequestId            *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtocolMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponseBody) SetNextToken(v string) *DescribeProtocolMountTargetResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBody) SetProtocolMountTargets(v []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets) *DescribeProtocolMountTargetResponseBody {
	s.ProtocolMountTargets = v
	return s
}

func (s *DescribeProtocolMountTargetResponseBody) SetRequestId(v string) *DescribeProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProtocolMountTargetResponseBodyProtocolMountTargets struct {
	AccessGroupName           *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExportId                  *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	FsetId                    *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	Path                      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ProtocolMountTargetDomain *string `json:"ProtocolMountTargetDomain,omitempty" xml:"ProtocolMountTargetDomain,omitempty"`
	ProtocolServiceId         *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	ProtocolType              *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId                 *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeProtocolMountTargetResponseBodyProtocolMountTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetAccessGroupName(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetCreateTime(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.CreateTime = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetDescription(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Description = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetExportId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ExportId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetFsetId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.FsetId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetPath(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Path = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolMountTargetDomain(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolMountTargetDomain = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolServiceId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolServiceId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolType(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolType = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetStatus(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Status = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetVSwitchId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.VSwitchId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetVpcId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.VpcId = &v
	return s
}

type DescribeProtocolMountTargetResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProtocolMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponse) SetHeaders(v map[string]*string) *DescribeProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtocolMountTargetResponse) SetStatusCode(v int32) *DescribeProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtocolMountTargetResponse) SetBody(v *DescribeProtocolMountTargetResponseBody) *DescribeProtocolMountTargetResponse {
	s.Body = v
	return s
}

type DescribeProtocolServiceRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId       *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MaxResults         *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProtocolServiceIds *string `json:"ProtocolServiceIds,omitempty" xml:"ProtocolServiceIds,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeProtocolServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceRequest) SetClientToken(v string) *DescribeProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetDescription(v string) *DescribeProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetFileSystemId(v string) *DescribeProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetMaxResults(v int64) *DescribeProtocolServiceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetNextToken(v string) *DescribeProtocolServiceRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetProtocolServiceIds(v string) *DescribeProtocolServiceRequest {
	s.ProtocolServiceIds = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetStatus(v string) *DescribeProtocolServiceRequest {
	s.Status = &v
	return s
}

type DescribeProtocolServiceResponseBody struct {
	NextToken        *string                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProtocolServices []*DescribeProtocolServiceResponseBodyProtocolServices `json:"ProtocolServices,omitempty" xml:"ProtocolServices,omitempty" type:"Repeated"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtocolServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponseBody) SetNextToken(v string) *DescribeProtocolServiceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolServiceResponseBody) SetProtocolServices(v []*DescribeProtocolServiceResponseBodyProtocolServices) *DescribeProtocolServiceResponseBody {
	s.ProtocolServices = v
	return s
}

func (s *DescribeProtocolServiceResponseBody) SetRequestId(v string) *DescribeProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProtocolServiceResponseBodyProtocolServices struct {
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId            *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InstanceBaseThroughput  *int32  `json:"InstanceBaseThroughput,omitempty" xml:"InstanceBaseThroughput,omitempty"`
	InstanceBurstThroughput *int32  `json:"InstanceBurstThroughput,omitempty" xml:"InstanceBurstThroughput,omitempty"`
	InstanceRAM             *int32  `json:"InstanceRAM,omitempty" xml:"InstanceRAM,omitempty"`
	ModifyTime              *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	MountTargetCount        *int32  `json:"MountTargetCount,omitempty" xml:"MountTargetCount,omitempty"`
	ProtocolServiceId       *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	ProtocolSpec            *string `json:"ProtocolSpec,omitempty" xml:"ProtocolSpec,omitempty"`
	ProtocolThroughput      *int32  `json:"ProtocolThroughput,omitempty" xml:"ProtocolThroughput,omitempty"`
	ProtocolType            *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeProtocolServiceResponseBodyProtocolServices) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolServiceResponseBodyProtocolServices) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetCreateTime(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.CreateTime = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetDescription(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.Description = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetFileSystemId(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceBaseThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceBaseThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceBurstThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceBurstThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceRAM(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceRAM = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetModifyTime(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ModifyTime = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetMountTargetCount(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.MountTargetCount = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolServiceId(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolServiceId = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolSpec(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolSpec = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolType(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolType = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetStatus(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.Status = &v
	return s
}

type DescribeProtocolServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProtocolServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponse) SetHeaders(v map[string]*string) *DescribeProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtocolServiceResponse) SetStatusCode(v int32) *DescribeProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtocolServiceResponse) SetBody(v *DescribeProtocolServiceResponseBody) *DescribeProtocolServiceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetFileSystemType(v string) *DescribeRegionsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeRegionsRequest) SetPageNumber(v int32) *DescribeRegionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRegionsRequest) SetPageSize(v int32) *DescribeRegionsRequest {
	s.PageSize = &v
	return s
}

type DescribeRegionsResponseBody struct {
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Regions    *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeRegionsResponseBody) SetPageSize(v int32) *DescribeRegionsResponseBody {
	s.PageSize = &v
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

func (s *DescribeRegionsResponseBody) SetTotalCount(v int32) *DescribeRegionsResponseBody {
	s.TotalCount = &v
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
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
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

type DescribeSmbAclRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeSmbAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmbAclRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclRequest) SetFileSystemId(v string) *DescribeSmbAclRequest {
	s.FileSystemId = &v
	return s
}

type DescribeSmbAclResponseBody struct {
	Acl       *DescribeSmbAclResponseBodyAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSmbAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponseBody) SetAcl(v *DescribeSmbAclResponseBodyAcl) *DescribeSmbAclResponseBody {
	s.Acl = v
	return s
}

func (s *DescribeSmbAclResponseBody) SetRequestId(v string) *DescribeSmbAclResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSmbAclResponseBodyAcl struct {
	EnableAnonymousAccess   *bool   `json:"EnableAnonymousAccess,omitempty" xml:"EnableAnonymousAccess,omitempty"`
	Enabled                 *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EncryptData             *bool   `json:"EncryptData,omitempty" xml:"EncryptData,omitempty"`
	HomeDirPath             *string `json:"HomeDirPath,omitempty" xml:"HomeDirPath,omitempty"`
	RejectUnencryptedAccess *bool   `json:"RejectUnencryptedAccess,omitempty" xml:"RejectUnencryptedAccess,omitempty"`
	SuperAdminSid           *string `json:"SuperAdminSid,omitempty" xml:"SuperAdminSid,omitempty"`
}

func (s DescribeSmbAclResponseBodyAcl) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmbAclResponseBodyAcl) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponseBodyAcl) SetEnableAnonymousAccess(v bool) *DescribeSmbAclResponseBodyAcl {
	s.EnableAnonymousAccess = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetEnabled(v bool) *DescribeSmbAclResponseBodyAcl {
	s.Enabled = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetEncryptData(v bool) *DescribeSmbAclResponseBodyAcl {
	s.EncryptData = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetHomeDirPath(v string) *DescribeSmbAclResponseBodyAcl {
	s.HomeDirPath = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetRejectUnencryptedAccess(v bool) *DescribeSmbAclResponseBodyAcl {
	s.RejectUnencryptedAccess = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetSuperAdminSid(v string) *DescribeSmbAclResponseBodyAcl {
	s.SuperAdminSid = &v
	return s
}

type DescribeSmbAclResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmbAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmbAclResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponse) SetHeaders(v map[string]*string) *DescribeSmbAclResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmbAclResponse) SetStatusCode(v int32) *DescribeSmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmbAclResponse) SetBody(v *DescribeSmbAclResponseBody) *DescribeSmbAclResponse {
	s.Body = v
	return s
}

type DescribeSnapshotsRequest struct {
	FileSystemId   *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SnapshotIds    *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotType   *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetFileSystemId(v string) *DescribeSnapshotsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetFileSystemType(v string) *DescribeSnapshotsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageNumber(v int32) *DescribeSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageSize(v int32) *DescribeSnapshotsRequest {
	s.PageSize = &v
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

type DescribeSnapshotsResponseBody struct {
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots  *DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeSnapshotsResponseBody) SetPageSize(v int32) *DescribeSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v *DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetTotalCount(v int32) *DescribeSnapshotsResponseBody {
	s.TotalCount = &v
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
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptType             *int32  `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	Progress                *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RemainTime              *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	RetentionDays           *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SnapshotId              *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName            *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SourceFileSystemId      *string `json:"SourceFileSystemId,omitempty" xml:"SourceFileSystemId,omitempty"`
	SourceFileSystemSize    *int64  `json:"SourceFileSystemSize,omitempty" xml:"SourceFileSystemSize,omitempty"`
	SourceFileSystemVersion *string `json:"SourceFileSystemVersion,omitempty" xml:"SourceFileSystemVersion,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetCreateTime(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.CreateTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetEncryptType(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.EncryptType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Progress = &v
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

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemSize(v int64) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemVersion(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemVersion = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Status = &v
	return s
}

type DescribeSnapshotsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSnapshotsResponse) SetStatusCode(v int32) *DescribeSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeStoragePackagesRequest struct {
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UseUTCDateTime *bool   `json:"UseUTCDateTime,omitempty" xml:"UseUTCDateTime,omitempty"`
}

func (s DescribeStoragePackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesRequest) SetPageNumber(v int32) *DescribeStoragePackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetPageSize(v int32) *DescribeStoragePackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetRegionId(v string) *DescribeStoragePackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStoragePackagesRequest) SetUseUTCDateTime(v bool) *DescribeStoragePackagesRequest {
	s.UseUTCDateTime = &v
	return s
}

type DescribeStoragePackagesResponseBody struct {
	Packages   *DescribeStoragePackagesResponseBodyPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStoragePackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBody) SetPackages(v *DescribeStoragePackagesResponseBodyPackages) *DescribeStoragePackagesResponseBody {
	s.Packages = v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetPageNumber(v int32) *DescribeStoragePackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetPageSize(v int32) *DescribeStoragePackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetRequestId(v string) *DescribeStoragePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetTotalCount(v int32) *DescribeStoragePackagesResponseBody {
	s.TotalCount = &v
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
	ExpiredTime  *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PackageId    *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeStoragePackagesResponseBodyPackagesPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoragePackagesResponseBodyPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetExpiredTime(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetFileSystemId(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.FileSystemId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetPackageId(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetSize(v int64) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStartTime(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStatus(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.Status = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStorageType(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.StorageType = &v
	return s
}

type DescribeStoragePackagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeStoragePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeStoragePackagesResponse) SetStatusCode(v int32) *DescribeStoragePackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoragePackagesResponse) SetBody(v *DescribeStoragePackagesResponseBody) *DescribeStoragePackagesResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	FileSystemId *string                   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageNumber   *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeTagsRequest) SetPageNumber(v int32) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int32) *DescribeTagsRequest {
	s.PageSize = &v
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
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags       *DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeTagsResponseBody) SetPageSize(v int32) *DescribeTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v *DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int32) *DescribeTagsResponseBody {
	s.TotalCount = &v
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
	FileSystemIds *DescribeTagsResponseBodyTagsTagFileSystemIds `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Struct"`
	Key           *string                                       `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string                                       `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagsResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagsTag) SetFileSystemIds(v *DescribeTagsResponseBodyTagsTagFileSystemIds) *DescribeTagsResponseBodyTagsTag {
	s.FileSystemIds = v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) SetKey(v string) *DescribeTagsResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeTagsResponseBodyTagsTag) SetValue(v string) *DescribeTagsResponseBodyTagsTag {
	s.Value = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTagsResponse) SetStatusCode(v int32) *DescribeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetFileSystemType(v string) *DescribeZonesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
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
	Capacity      *DescribeZonesResponseBodyZonesZoneCapacity      `json:"Capacity,omitempty" xml:"Capacity,omitempty" type:"Struct"`
	InstanceTypes *DescribeZonesResponseBodyZonesZoneInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	Performance   *DescribeZonesResponseBodyZonesZonePerformance   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	ZoneId        *string                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetCapacity(v *DescribeZonesResponseBodyZonesZoneCapacity) *DescribeZonesResponseBodyZonesZone {
	s.Capacity = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetInstanceTypes(v *DescribeZonesResponseBodyZonesZoneInstanceTypes) *DescribeZonesResponseBodyZonesZone {
	s.InstanceTypes = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetPerformance(v *DescribeZonesResponseBodyZonesZonePerformance) *DescribeZonesResponseBodyZonesZone {
	s.Performance = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
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
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) SetProtocolType(v string) *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	s.ProtocolType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType) SetStorageType(v string) *DescribeZonesResponseBodyZonesZoneInstanceTypesInstanceType {
	s.StorageType = &v
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

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableAndCleanRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DisableAndCleanRecycleBinResponse) SetStatusCode(v int32) *DisableAndCleanRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAndCleanRecycleBinResponse) SetBody(v *DisableAndCleanRecycleBinResponseBody) *DisableAndCleanRecycleBinResponse {
	s.Body = v
	return s
}

type DisableSmbAclRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableSmbAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSmbAclRequest) GoString() string {
	return s.String()
}

func (s *DisableSmbAclRequest) SetFileSystemId(v string) *DisableSmbAclRequest {
	s.FileSystemId = &v
	return s
}

type DisableSmbAclResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableSmbAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSmbAclResponseBody) SetRequestId(v string) *DisableSmbAclResponseBody {
	s.RequestId = &v
	return s
}

type DisableSmbAclResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSmbAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSmbAclResponse) GoString() string {
	return s.String()
}

func (s *DisableSmbAclResponse) SetHeaders(v map[string]*string) *DisableSmbAclResponse {
	s.Headers = v
	return s
}

func (s *DisableSmbAclResponse) SetStatusCode(v int32) *DisableSmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSmbAclResponse) SetBody(v *DisableSmbAclResponseBody) *DisableSmbAclResponse {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableRecycleBinResponse) SetStatusCode(v int32) *EnableRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableRecycleBinResponse) SetBody(v *EnableRecycleBinResponseBody) *EnableRecycleBinResponse {
	s.Body = v
	return s
}

type EnableSmbAclRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Keytab       *string `json:"Keytab,omitempty" xml:"Keytab,omitempty"`
	KeytabMd5    *string `json:"KeytabMd5,omitempty" xml:"KeytabMd5,omitempty"`
}

func (s EnableSmbAclRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSmbAclRequest) GoString() string {
	return s.String()
}

func (s *EnableSmbAclRequest) SetFileSystemId(v string) *EnableSmbAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *EnableSmbAclRequest) SetKeytab(v string) *EnableSmbAclRequest {
	s.Keytab = &v
	return s
}

func (s *EnableSmbAclRequest) SetKeytabMd5(v string) *EnableSmbAclRequest {
	s.KeytabMd5 = &v
	return s
}

type EnableSmbAclResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSmbAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSmbAclResponseBody) SetRequestId(v string) *EnableSmbAclResponseBody {
	s.RequestId = &v
	return s
}

type EnableSmbAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSmbAclResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSmbAclResponse) GoString() string {
	return s.String()
}

func (s *EnableSmbAclResponse) SetHeaders(v map[string]*string) *EnableSmbAclResponse {
	s.Headers = v
	return s
}

func (s *EnableSmbAclResponse) SetStatusCode(v int32) *EnableSmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSmbAclResponse) SetBody(v *EnableSmbAclResponseBody) *EnableSmbAclResponse {
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
	ATime                   *string `json:"ATime,omitempty" xml:"ATime,omitempty"`
	CTime                   *string `json:"CTime,omitempty" xml:"CTime,omitempty"`
	HasInfrequentAccessFile *bool   `json:"HasInfrequentAccessFile,omitempty" xml:"HasInfrequentAccessFile,omitempty"`
	Inode                   *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	MTime                   *string `json:"MTime,omitempty" xml:"MTime,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RetrieveTime            *string `json:"RetrieveTime,omitempty" xml:"RetrieveTime,omitempty"`
	Size                    *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageType             *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDirectoryOrFilePropertiesResponseBodyEntry) String() string {
	return tea.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponseBodyEntry) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetATime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.ATime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetCTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.CTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetHasInfrequentAccessFile(v bool) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.HasInfrequentAccessFile = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetInode(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Inode = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetMTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.MTime = &v
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

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetSize(v int64) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Size = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetStorageType(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.StorageType = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetType(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Type = &v
	return s
}

type GetDirectoryOrFilePropertiesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDirectoryOrFilePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDirectoryOrFilePropertiesResponse) SetStatusCode(v int32) *GetDirectoryOrFilePropertiesResponse {
	s.StatusCode = &v
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
	RecycleBinAttribute *GetRecycleBinAttributeResponseBodyRecycleBinAttribute `json:"RecycleBinAttribute,omitempty" xml:"RecycleBinAttribute,omitempty" type:"Struct"`
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecycleBinAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponseBody) SetRecycleBinAttribute(v *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) *GetRecycleBinAttributeResponseBody {
	s.RecycleBinAttribute = v
	return s
}

func (s *GetRecycleBinAttributeResponseBody) SetRequestId(v string) *GetRecycleBinAttributeResponseBody {
	s.RequestId = &v
	return s
}

type GetRecycleBinAttributeResponseBodyRecycleBinAttribute struct {
	EnableTime    *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	ReservedDays  *int64  `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
	SecondarySize *int64  `json:"SecondarySize,omitempty" xml:"SecondarySize,omitempty"`
	Size          *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRecycleBinAttributeResponseBodyRecycleBinAttribute) String() string {
	return tea.Prettify(s)
}

func (s GetRecycleBinAttributeResponseBodyRecycleBinAttribute) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetEnableTime(v string) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.EnableTime = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetReservedDays(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.ReservedDays = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetSecondarySize(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.SecondarySize = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetSize(v int64) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.Size = &v
	return s
}

func (s *GetRecycleBinAttributeResponseBodyRecycleBinAttribute) SetStatus(v string) *GetRecycleBinAttributeResponseBodyRecycleBinAttribute {
	s.Status = &v
	return s
}

type GetRecycleBinAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRecycleBinAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRecycleBinAttributeResponse) SetStatusCode(v int32) *GetRecycleBinAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecycleBinAttributeResponse) SetBody(v *GetRecycleBinAttributeResponseBody) *GetRecycleBinAttributeResponse {
	s.Body = v
	return s
}

type ListDirectoriesAndFilesRequest struct {
	DirectoryOnly *bool   `json:"DirectoryOnly,omitempty" xml:"DirectoryOnly,omitempty"`
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MaxResults    *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Path          *string `json:"Path,omitempty" xml:"Path,omitempty"`
	StorageType   *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ListDirectoriesAndFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesRequest) SetDirectoryOnly(v bool) *ListDirectoriesAndFilesRequest {
	s.DirectoryOnly = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetFileSystemId(v string) *ListDirectoriesAndFilesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetMaxResults(v int64) *ListDirectoriesAndFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetNextToken(v string) *ListDirectoriesAndFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetPath(v string) *ListDirectoriesAndFilesRequest {
	s.Path = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetStorageType(v string) *ListDirectoriesAndFilesRequest {
	s.StorageType = &v
	return s
}

type ListDirectoriesAndFilesResponseBody struct {
	Entries   []*ListDirectoriesAndFilesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	NextToken *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDirectoriesAndFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponseBody) SetEntries(v []*ListDirectoriesAndFilesResponseBodyEntries) *ListDirectoriesAndFilesResponseBody {
	s.Entries = v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) SetNextToken(v string) *ListDirectoriesAndFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) SetRequestId(v string) *ListDirectoriesAndFilesResponseBody {
	s.RequestId = &v
	return s
}

type ListDirectoriesAndFilesResponseBodyEntries struct {
	Atime                   *string `json:"Atime,omitempty" xml:"Atime,omitempty"`
	Ctime                   *string `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	FileId                  *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	HasInfrequentAccessFile *bool   `json:"HasInfrequentAccessFile,omitempty" xml:"HasInfrequentAccessFile,omitempty"`
	Inode                   *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	Mtime                   *string `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner                   *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RetrieveTime            *string `json:"RetrieveTime,omitempty" xml:"RetrieveTime,omitempty"`
	Size                    *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageType             *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDirectoriesAndFilesResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoriesAndFilesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetAtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Atime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetCtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Ctime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetFileId(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetHasInfrequentAccessFile(v bool) *ListDirectoriesAndFilesResponseBodyEntries {
	s.HasInfrequentAccessFile = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetInode(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Inode = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetMtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Mtime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetName(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetOwner(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Owner = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetRetrieveTime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.RetrieveTime = &v
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

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetType(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Type = &v
	return s
}

type ListDirectoriesAndFilesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDirectoriesAndFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDirectoriesAndFilesResponse) SetStatusCode(v int32) *ListDirectoriesAndFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDirectoriesAndFilesResponse) SetBody(v *ListDirectoriesAndFilesResponseBody) *ListDirectoriesAndFilesResponse {
	s.Body = v
	return s
}

type ListLifecycleRetrieveJobsRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLifecycleRetrieveJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsRequest) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsRequest) SetFileSystemId(v string) *ListLifecycleRetrieveJobsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetPageNumber(v int32) *ListLifecycleRetrieveJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetPageSize(v int32) *ListLifecycleRetrieveJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetStatus(v string) *ListLifecycleRetrieveJobsRequest {
	s.Status = &v
	return s
}

type ListLifecycleRetrieveJobsResponseBody struct {
	LifecycleRetrieveJobs []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs `json:"LifecycleRetrieveJobs,omitempty" xml:"LifecycleRetrieveJobs,omitempty" type:"Repeated"`
	PageNumber            *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLifecycleRetrieveJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetLifecycleRetrieveJobs(v []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) *ListLifecycleRetrieveJobsResponseBody {
	s.LifecycleRetrieveJobs = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetPageNumber(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetPageSize(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetRequestId(v string) *ListLifecycleRetrieveJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetTotalCount(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs struct {
	CreateTime          *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DiscoveredFileCount *int64    `json:"DiscoveredFileCount,omitempty" xml:"DiscoveredFileCount,omitempty"`
	FileSystemId        *string   `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	JobId               *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Paths               []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	RetrievedFileCount  *int64    `json:"RetrievedFileCount,omitempty" xml:"RetrievedFileCount,omitempty"`
	Status              *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime          *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) String() string {
	return tea.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetCreateTime(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.CreateTime = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetDiscoveredFileCount(v int64) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.DiscoveredFileCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetFileSystemId(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.FileSystemId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetJobId(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.JobId = &v
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

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetStatus(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.Status = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetUpdateTime(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.UpdateTime = &v
	return s
}

type ListLifecycleRetrieveJobsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLifecycleRetrieveJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListLifecycleRetrieveJobsResponse) SetStatusCode(v int32) *ListLifecycleRetrieveJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponse) SetBody(v *ListLifecycleRetrieveJobsResponseBody) *ListLifecycleRetrieveJobsResponse {
	s.Body = v
	return s
}

type ListRecentlyRecycledDirectoriesRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MaxResults   *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s *ListRecentlyRecycledDirectoriesRequest) SetMaxResults(v int64) *ListRecentlyRecycledDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetNextToken(v string) *ListRecentlyRecycledDirectoriesRequest {
	s.NextToken = &v
	return s
}

type ListRecentlyRecycledDirectoriesResponseBody struct {
	Entries   []*ListRecentlyRecycledDirectoriesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	NextToken *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetEntries(v []*ListRecentlyRecycledDirectoriesResponseBodyEntries) *ListRecentlyRecycledDirectoriesResponseBody {
	s.Entries = v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetNextToken(v string) *ListRecentlyRecycledDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBody) SetRequestId(v string) *ListRecentlyRecycledDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type ListRecentlyRecycledDirectoriesResponseBodyEntries struct {
	FileId         *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	LastDeleteTime *string `json:"LastDeleteTime,omitempty" xml:"LastDeleteTime,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
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

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetLastDeleteTime(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.LastDeleteTime = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetName(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponseBodyEntries) SetPath(v string) *ListRecentlyRecycledDirectoriesResponseBodyEntries {
	s.Path = &v
	return s
}

type ListRecentlyRecycledDirectoriesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecentlyRecycledDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRecentlyRecycledDirectoriesResponse) SetStatusCode(v int32) *ListRecentlyRecycledDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesResponse) SetBody(v *ListRecentlyRecycledDirectoriesResponseBody) *ListRecentlyRecycledDirectoriesResponse {
	s.Body = v
	return s
}

type ListRecycleBinJobsRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber   *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListRecycleBinJobsRequest) SetPageNumber(v int64) *ListRecycleBinJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetPageSize(v int64) *ListRecycleBinJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetStatus(v string) *ListRecycleBinJobsRequest {
	s.Status = &v
	return s
}

type ListRecycleBinJobsResponseBody struct {
	Jobs       []*ListRecycleBinJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecycleBinJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponseBody) SetJobs(v []*ListRecycleBinJobsResponseBodyJobs) *ListRecycleBinJobsResponseBody {
	s.Jobs = v
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

func (s *ListRecycleBinJobsResponseBody) SetRequestId(v string) *ListRecycleBinJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetTotalCount(v int64) *ListRecycleBinJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListRecycleBinJobsResponseBodyJobs struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecycleBinJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListRecycleBinJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetCreateTime(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetErrorCode(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.ErrorCode = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetErrorMessage(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.ErrorMessage = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetFileId(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.FileId = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetFileName(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.FileName = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetId(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetProgress(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Progress = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetStatus(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetType(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Type = &v
	return s
}

type ListRecycleBinJobsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecycleBinJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRecycleBinJobsResponse) SetStatusCode(v int32) *ListRecycleBinJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecycleBinJobsResponse) SetBody(v *ListRecycleBinJobsResponseBody) *ListRecycleBinJobsResponse {
	s.Body = v
	return s
}

type ListRecycledDirectoriesAndFilesRequest struct {
	FileId       *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MaxResults   *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesRequest) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetFileId(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.FileId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetFileSystemId(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetMaxResults(v int64) *ListRecycledDirectoriesAndFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetNextToken(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.NextToken = &v
	return s
}

type ListRecycledDirectoriesAndFilesResponseBody struct {
	Entries   []*ListRecycledDirectoriesAndFilesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	NextToken *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetEntries(v []*ListRecycledDirectoriesAndFilesResponseBodyEntries) *ListRecycledDirectoriesAndFilesResponseBody {
	s.Entries = v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetNextToken(v string) *ListRecycledDirectoriesAndFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetRequestId(v string) *ListRecycledDirectoriesAndFilesResponseBody {
	s.RequestId = &v
	return s
}

type ListRecycledDirectoriesAndFilesResponseBodyEntries struct {
	ATime      *string `json:"ATime,omitempty" xml:"ATime,omitempty"`
	CTime      *string `json:"CTime,omitempty" xml:"CTime,omitempty"`
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	FileId     *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Inode      *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	MTime      *string `json:"MTime,omitempty" xml:"MTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size       *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesResponseBodyEntries) String() string {
	return tea.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetATime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.ATime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetCTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.CTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetDeleteTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.DeleteTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetFileId(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetInode(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Inode = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetMTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.MTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetName(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetSize(v int64) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Size = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetType(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Type = &v
	return s
}

type ListRecycledDirectoriesAndFilesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecycledDirectoriesAndFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRecycledDirectoriesAndFilesResponse) SetStatusCode(v int32) *ListRecycledDirectoriesAndFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponse) SetBody(v *ListRecycledDirectoriesAndFilesResponseBody) *ListRecycledDirectoriesAndFilesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyAccessGroupResponse) SetStatusCode(v int32) *ModifyAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessGroupResponse) SetBody(v *ModifyAccessGroupResponseBody) *ModifyAccessGroupResponse {
	s.Body = v
	return s
}

type ModifyAccessRuleRequest struct {
	AccessGroupName  *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	AccessRuleId     *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	FileSystemType   *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	Priority         *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RWAccessType     *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	SourceCidrIp     *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	UserAccessType   *string `json:"UserAccessType,omitempty" xml:"UserAccessType,omitempty"`
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

func (s *ModifyAccessRuleRequest) SetFileSystemType(v string) *ModifyAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetIpv6SourceCidrIp(v string) *ModifyAccessRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetPriority(v int32) *ModifyAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetRWAccessType(v string) *ModifyAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetSourceCidrIp(v string) *ModifyAccessRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetUserAccessType(v string) *ModifyAccessRuleRequest {
	s.UserAccessType = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyAccessRuleResponse) SetStatusCode(v int32) *ModifyAccessRuleResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyAutoSnapshotPolicyResponse) SetStatusCode(v int32) *ModifyAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyResponse) SetBody(v *ModifyAutoSnapshotPolicyResponseBody) *ModifyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type ModifyDataFlowRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId   *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Throughput   *int64  `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s ModifyDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowRequest) SetClientToken(v string) *ModifyDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDataFlowId(v string) *ModifyDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDescription(v string) *ModifyDataFlowRequest {
	s.Description = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDryRun(v bool) *ModifyDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDataFlowRequest) SetFileSystemId(v string) *ModifyDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyDataFlowRequest) SetThroughput(v int64) *ModifyDataFlowRequest {
	s.Throughput = &v
	return s
}

type ModifyDataFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowResponseBody) SetRequestId(v string) *ModifyDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowResponse) SetHeaders(v map[string]*string) *ModifyDataFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataFlowResponse) SetStatusCode(v int32) *ModifyDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataFlowResponse) SetBody(v *ModifyDataFlowResponseBody) *ModifyDataFlowResponse {
	s.Body = v
	return s
}

type ModifyDataFlowAutoRefreshRequest struct {
	AutoRefreshInterval *int64  `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	AutoRefreshPolicy   *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId          *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId        *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyDataFlowAutoRefreshRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshRequest) SetAutoRefreshInterval(v int64) *ModifyDataFlowAutoRefreshRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetAutoRefreshPolicy(v string) *ModifyDataFlowAutoRefreshRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetClientToken(v string) *ModifyDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetDataFlowId(v string) *ModifyDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetDryRun(v bool) *ModifyDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetFileSystemId(v string) *ModifyDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

type ModifyDataFlowAutoRefreshResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataFlowAutoRefreshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshResponseBody) SetRequestId(v string) *ModifyDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataFlowAutoRefreshResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *ModifyDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataFlowAutoRefreshResponse) SetStatusCode(v int32) *ModifyDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshResponse) SetBody(v *ModifyDataFlowAutoRefreshResponseBody) *ModifyDataFlowAutoRefreshResponse {
	s.Body = v
	return s
}

type ModifyFileSystemRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequest) SetDescription(v string) *ModifyFileSystemRequest {
	s.Description = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemId(v string) *ModifyFileSystemRequest {
	s.FileSystemId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyFileSystemResponse) SetStatusCode(v int32) *ModifyFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFileSystemResponse) SetBody(v *ModifyFileSystemResponseBody) *ModifyFileSystemResponse {
	s.Body = v
	return s
}

type ModifyFilesetRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FsetId       *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s ModifyFilesetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFilesetRequest) GoString() string {
	return s.String()
}

func (s *ModifyFilesetRequest) SetClientToken(v string) *ModifyFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFilesetRequest) SetDescription(v string) *ModifyFilesetRequest {
	s.Description = &v
	return s
}

func (s *ModifyFilesetRequest) SetDryRun(v bool) *ModifyFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyFilesetRequest) SetFileSystemId(v string) *ModifyFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFilesetRequest) SetFsetId(v string) *ModifyFilesetRequest {
	s.FsetId = &v
	return s
}

type ModifyFilesetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFilesetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFilesetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFilesetResponseBody) SetRequestId(v string) *ModifyFilesetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFilesetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFilesetResponse) GoString() string {
	return s.String()
}

func (s *ModifyFilesetResponse) SetHeaders(v map[string]*string) *ModifyFilesetResponse {
	s.Headers = v
	return s
}

func (s *ModifyFilesetResponse) SetStatusCode(v int32) *ModifyFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFilesetResponse) SetBody(v *ModifyFilesetResponseBody) *ModifyFilesetResponse {
	s.Body = v
	return s
}

type ModifyLDAPConfigRequest struct {
	BindDN       *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	SearchBase   *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	URI          *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s ModifyLDAPConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigRequest) SetBindDN(v string) *ModifyLDAPConfigRequest {
	s.BindDN = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetFileSystemId(v string) *ModifyLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetSearchBase(v string) *ModifyLDAPConfigRequest {
	s.SearchBase = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetURI(v string) *ModifyLDAPConfigRequest {
	s.URI = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyLDAPConfigResponse) SetStatusCode(v int32) *ModifyLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLDAPConfigResponse) SetBody(v *ModifyLDAPConfigResponseBody) *ModifyLDAPConfigResponse {
	s.Body = v
	return s
}

type ModifyLifecyclePolicyRequest struct {
	FileSystemId        *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	LifecycleRuleName   *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
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

func (s *ModifyLifecyclePolicyRequest) SetLifecycleRuleName(v string) *ModifyLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetPath(v string) *ModifyLifecyclePolicyRequest {
	s.Path = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLifecyclePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyLifecyclePolicyResponse) SetStatusCode(v int32) *ModifyLifecyclePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLifecyclePolicyResponse) SetBody(v *ModifyLifecyclePolicyResponseBody) *ModifyLifecyclePolicyResponse {
	s.Body = v
	return s
}

type ModifyMountTargetRequest struct {
	AccessGroupName            *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	FileSystemId               *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain          *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	Status                     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetRequest) SetAccessGroupName(v string) *ModifyMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyMountTargetRequest) SetDualStackMountTargetDomain(v string) *ModifyMountTargetRequest {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *ModifyMountTargetRequest) SetFileSystemId(v string) *ModifyMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyMountTargetRequest) SetMountTargetDomain(v string) *ModifyMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *ModifyMountTargetRequest) SetStatus(v string) *ModifyMountTargetRequest {
	s.Status = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyMountTargetResponse) SetStatusCode(v int32) *ModifyMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMountTargetResponse) SetBody(v *ModifyMountTargetResponseBody) *ModifyMountTargetResponse {
	s.Body = v
	return s
}

type ModifyProtocolMountTargetRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ExportId          *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s ModifyProtocolMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtocolMountTargetRequest) SetClientToken(v string) *ModifyProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetDescription(v string) *ModifyProtocolMountTargetRequest {
	s.Description = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetDryRun(v bool) *ModifyProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetExportId(v string) *ModifyProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetFileSystemId(v string) *ModifyProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetProtocolServiceId(v string) *ModifyProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

type ModifyProtocolMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtocolMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtocolMountTargetResponseBody) SetRequestId(v string) *ModifyProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtocolMountTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyProtocolMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtocolMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtocolMountTargetResponse) SetHeaders(v map[string]*string) *ModifyProtocolMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtocolMountTargetResponse) SetStatusCode(v int32) *ModifyProtocolMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtocolMountTargetResponse) SetBody(v *ModifyProtocolMountTargetResponseBody) *ModifyProtocolMountTargetResponse {
	s.Body = v
	return s
}

type ModifyProtocolServiceRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s ModifyProtocolServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceRequest) SetClientToken(v string) *ModifyProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetDescription(v string) *ModifyProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetDryRun(v bool) *ModifyProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetFileSystemId(v string) *ModifyProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetProtocolServiceId(v string) *ModifyProtocolServiceRequest {
	s.ProtocolServiceId = &v
	return s
}

type ModifyProtocolServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtocolServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceResponseBody) SetRequestId(v string) *ModifyProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtocolServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceResponse) SetHeaders(v map[string]*string) *ModifyProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtocolServiceResponse) SetStatusCode(v int32) *ModifyProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtocolServiceResponse) SetBody(v *ModifyProtocolServiceResponseBody) *ModifyProtocolServiceResponse {
	s.Body = v
	return s
}

type ModifySmbAclRequest struct {
	EnableAnonymousAccess   *bool   `json:"EnableAnonymousAccess,omitempty" xml:"EnableAnonymousAccess,omitempty"`
	EncryptData             *bool   `json:"EncryptData,omitempty" xml:"EncryptData,omitempty"`
	FileSystemId            *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	HomeDirPath             *string `json:"HomeDirPath,omitempty" xml:"HomeDirPath,omitempty"`
	Keytab                  *string `json:"Keytab,omitempty" xml:"Keytab,omitempty"`
	KeytabMd5               *string `json:"KeytabMd5,omitempty" xml:"KeytabMd5,omitempty"`
	RejectUnencryptedAccess *bool   `json:"RejectUnencryptedAccess,omitempty" xml:"RejectUnencryptedAccess,omitempty"`
	SuperAdminSid           *string `json:"SuperAdminSid,omitempty" xml:"SuperAdminSid,omitempty"`
}

func (s ModifySmbAclRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySmbAclRequest) GoString() string {
	return s.String()
}

func (s *ModifySmbAclRequest) SetEnableAnonymousAccess(v bool) *ModifySmbAclRequest {
	s.EnableAnonymousAccess = &v
	return s
}

func (s *ModifySmbAclRequest) SetEncryptData(v bool) *ModifySmbAclRequest {
	s.EncryptData = &v
	return s
}

func (s *ModifySmbAclRequest) SetFileSystemId(v string) *ModifySmbAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifySmbAclRequest) SetHomeDirPath(v string) *ModifySmbAclRequest {
	s.HomeDirPath = &v
	return s
}

func (s *ModifySmbAclRequest) SetKeytab(v string) *ModifySmbAclRequest {
	s.Keytab = &v
	return s
}

func (s *ModifySmbAclRequest) SetKeytabMd5(v string) *ModifySmbAclRequest {
	s.KeytabMd5 = &v
	return s
}

func (s *ModifySmbAclRequest) SetRejectUnencryptedAccess(v bool) *ModifySmbAclRequest {
	s.RejectUnencryptedAccess = &v
	return s
}

func (s *ModifySmbAclRequest) SetSuperAdminSid(v string) *ModifySmbAclRequest {
	s.SuperAdminSid = &v
	return s
}

type ModifySmbAclResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySmbAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmbAclResponseBody) SetRequestId(v string) *ModifySmbAclResponseBody {
	s.RequestId = &v
	return s
}

type ModifySmbAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySmbAclResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySmbAclResponse) GoString() string {
	return s.String()
}

func (s *ModifySmbAclResponse) SetHeaders(v map[string]*string) *ModifySmbAclResponse {
	s.Headers = v
	return s
}

func (s *ModifySmbAclResponse) SetStatusCode(v int32) *ModifySmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySmbAclResponse) SetBody(v *ModifySmbAclResponseBody) *ModifySmbAclResponse {
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenNASServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OpenNASServiceResponse) SetStatusCode(v int32) *OpenNASServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenNASServiceResponse) SetBody(v *OpenNASServiceResponseBody) *OpenNASServiceResponse {
	s.Body = v
	return s
}

type RemoveClientFromBlackListRequest struct {
	ClientIP     *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveClientFromBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientFromBlackListRequest) GoString() string {
	return s.String()
}

func (s *RemoveClientFromBlackListRequest) SetClientIP(v string) *RemoveClientFromBlackListRequest {
	s.ClientIP = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetClientToken(v string) *RemoveClientFromBlackListRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetFileSystemId(v string) *RemoveClientFromBlackListRequest {
	s.FileSystemId = &v
	return s
}

func (s *RemoveClientFromBlackListRequest) SetRegionId(v string) *RemoveClientFromBlackListRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveClientFromBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveClientFromBlackListResponse) SetStatusCode(v int32) *RemoveClientFromBlackListResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveTagsResponse) SetStatusCode(v int32) *RemoveTagsResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetFileSystemResponse) SetStatusCode(v int32) *ResetFileSystemResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RetryLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RetryLifecycleRetrieveJobResponse) SetStatusCode(v int32) *RetryLifecycleRetrieveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryLifecycleRetrieveJobResponse) SetBody(v *RetryLifecycleRetrieveJobResponseBody) *RetryLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

type SetDirQuotaRequest struct {
	FileCountLimit *int64  `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	FileSystemId   *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	QuotaType      *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	SizeLimit      *int64  `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType       *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s SetDirQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDirQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetDirQuotaRequest) SetFileCountLimit(v int64) *SetDirQuotaRequest {
	s.FileCountLimit = &v
	return s
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

func (s *SetDirQuotaRequest) SetSizeLimit(v int64) *SetDirQuotaRequest {
	s.SizeLimit = &v
	return s
}

func (s *SetDirQuotaRequest) SetUserId(v string) *SetDirQuotaRequest {
	s.UserId = &v
	return s
}

func (s *SetDirQuotaRequest) SetUserType(v string) *SetDirQuotaRequest {
	s.UserType = &v
	return s
}

type SetDirQuotaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDirQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDirQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetDirQuotaResponseBody) SetRequestId(v string) *SetDirQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDirQuotaResponseBody) SetSuccess(v bool) *SetDirQuotaResponseBody {
	s.Success = &v
	return s
}

type SetDirQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDirQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetDirQuotaResponse) SetStatusCode(v int32) *SetDirQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDirQuotaResponse) SetBody(v *SetDirQuotaResponseBody) *SetDirQuotaResponse {
	s.Body = v
	return s
}

type StartDataFlowRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId   *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s StartDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDataFlowRequest) GoString() string {
	return s.String()
}

func (s *StartDataFlowRequest) SetClientToken(v string) *StartDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDataFlowRequest) SetDataFlowId(v string) *StartDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *StartDataFlowRequest) SetDryRun(v bool) *StartDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *StartDataFlowRequest) SetFileSystemId(v string) *StartDataFlowRequest {
	s.FileSystemId = &v
	return s
}

type StartDataFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StartDataFlowResponseBody) SetRequestId(v string) *StartDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type StartDataFlowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDataFlowResponse) GoString() string {
	return s.String()
}

func (s *StartDataFlowResponse) SetHeaders(v map[string]*string) *StartDataFlowResponse {
	s.Headers = v
	return s
}

func (s *StartDataFlowResponse) SetStatusCode(v int32) *StartDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDataFlowResponse) SetBody(v *StartDataFlowResponseBody) *StartDataFlowResponse {
	s.Body = v
	return s
}

type StopDataFlowRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataFlowId   *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s StopDataFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDataFlowRequest) GoString() string {
	return s.String()
}

func (s *StopDataFlowRequest) SetClientToken(v string) *StopDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDataFlowRequest) SetDataFlowId(v string) *StopDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *StopDataFlowRequest) SetDryRun(v bool) *StopDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *StopDataFlowRequest) SetFileSystemId(v string) *StopDataFlowRequest {
	s.FileSystemId = &v
	return s
}

type StopDataFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDataFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StopDataFlowResponseBody) SetRequestId(v string) *StopDataFlowResponseBody {
	s.RequestId = &v
	return s
}

type StopDataFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDataFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDataFlowResponse) GoString() string {
	return s.String()
}

func (s *StopDataFlowResponse) SetHeaders(v map[string]*string) *StopDataFlowResponse {
	s.Headers = v
	return s
}

func (s *StopDataFlowResponse) SetStatusCode(v int32) *StopDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDataFlowResponse) SetBody(v *StopDataFlowResponseBody) *StopDataFlowResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRecycleBinAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateRecycleBinAttributeResponse) SetStatusCode(v int32) *UpdateRecycleBinAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecycleBinAttributeResponse) SetBody(v *UpdateRecycleBinAttributeResponseBody) *UpdateRecycleBinAttributeResponse {
	s.Body = v
	return s
}

type UpgradeFileSystemRequest struct {
	Capacity     *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s UpgradeFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeFileSystemRequest) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemRequest) SetCapacity(v int64) *UpgradeFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetClientToken(v string) *UpgradeFileSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetDryRun(v bool) *UpgradeFileSystemRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeFileSystemRequest) SetFileSystemId(v string) *UpgradeFileSystemRequest {
	s.FileSystemId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeFileSystemResponse) SetStatusCode(v int32) *UpgradeFileSystemResponse {
	s.StatusCode = &v
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIP)) {
		query["ClientIP"] = request.ClientIP
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddClientToBlackList"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddClientToBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTags"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyId)) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemIds)) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ApplyDataFlowAutoRefreshWithOptions(request *ApplyDataFlowAutoRefreshRequest, runtime *util.RuntimeOptions) (_result *ApplyDataFlowAutoRefreshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRefreshInterval)) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshPolicy)) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshs)) {
		query["AutoRefreshs"] = request.AutoRefreshs
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyDataFlowAutoRefresh"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyDataFlowAutoRefresh(request *ApplyDataFlowAutoRefreshRequest) (_result *ApplyDataFlowAutoRefreshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyDataFlowAutoRefreshResponse{}
	_body, _err := client.ApplyDataFlowAutoRefreshWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemIds)) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CancelDataFlowAutoRefreshWithOptions(request *CancelDataFlowAutoRefreshRequest, runtime *util.RuntimeOptions) (_result *CancelDataFlowAutoRefreshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshPath)) {
		query["RefreshPath"] = request.RefreshPath
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDataFlowAutoRefresh"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelDataFlowAutoRefresh(request *CancelDataFlowAutoRefreshRequest) (_result *CancelDataFlowAutoRefreshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelDataFlowAutoRefreshResponse{}
	_body, _err := client.CancelDataFlowAutoRefreshWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelDataFlowTaskWithOptions(request *CancelDataFlowTaskRequest, runtime *util.RuntimeOptions) (_result *CancelDataFlowTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDataFlowTask"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDataFlowTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelDataFlowTask(request *CancelDataFlowTaskRequest) (_result *CancelDataFlowTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelDataFlowTaskResponse{}
	_body, _err := client.CancelDataFlowTaskWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDirQuota"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDirQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelLifecycleRetrieveJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRecycleBinJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRecycleBinJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AccessGroupType)) {
		query["AccessGroupType"] = request.AccessGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessGroup"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6SourceCidrIp)) {
		query["Ipv6SourceCidrIp"] = request.Ipv6SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RWAccessType)) {
		query["RWAccessType"] = request.RWAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrIp)) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessType)) {
		query["UserAccessType"] = request.UserAccessType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessRule"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyName)) {
		query["AutoSnapshotPolicyName"] = request.AutoSnapshotPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatWeekdays)) {
		query["RepeatWeekdays"] = request.RepeatWeekdays
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.TimePoints)) {
		query["TimePoints"] = request.TimePoints
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateDataFlowWithOptions(request *CreateDataFlowRequest, runtime *util.RuntimeOptions) (_result *CreateDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRefreshInterval)) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshPolicy)) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshs)) {
		query["AutoRefreshs"] = request.AutoRefreshs
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FsetId)) {
		query["FsetId"] = request.FsetId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceSecurityType)) {
		query["SourceSecurityType"] = request.SourceSecurityType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceStorage)) {
		query["SourceStorage"] = request.SourceStorage
	}

	if !tea.BoolValue(util.IsUnset(request.Throughput)) {
		query["Throughput"] = request.Throughput
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataFlow(request *CreateDataFlowRequest) (_result *CreateDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataFlowResponse{}
	_body, _err := client.CreateDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataFlowTaskWithOptions(request *CreateDataFlowTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDataFlowTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.Directory)) {
		query["Directory"] = request.Directory
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EntryList)) {
		query["EntryList"] = request.EntryList
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.SrcTaskId)) {
		query["SrcTaskId"] = request.SrcTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskAction)) {
		query["TaskAction"] = request.TaskAction
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataFlowTask"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataFlowTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataFlowTask(request *CreateDataFlowTaskRequest) (_result *CreateDataFlowTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataFlowTaskResponse{}
	_body, _err := client.CreateDataFlowTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFileWithOptions(request *CreateFileRequest, runtime *util.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccessInheritable)) {
		query["OwnerAccessInheritable"] = request.OwnerAccessInheritable
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFile"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFile(request *CreateFileRequest) (_result *CreateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFileResponse{}
	_body, _err := client.CreateFileWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		query["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptType)) {
		query["EncryptType"] = request.EncryptType
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.KmsKeyId)) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateFilesetWithOptions(request *CreateFilesetRequest, runtime *util.RuntimeOptions) (_result *CreateFilesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemPath)) {
		query["FileSystemPath"] = request.FileSystemPath
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFileset"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFileset(request *CreateFilesetRequest) (_result *CreateFilesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFilesetResponse{}
	_body, _err := client.CreateFilesetWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindDN)) {
		query["BindDN"] = request.BindDN
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchBase)) {
		query["SearchBase"] = request.SearchBase
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLDAPConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.LifecyclePolicyName)) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleRuleName)) {
		query["LifecycleRuleName"] = request.LifecycleRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Paths)) {
		query["Paths"] = request.Paths
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLifecyclePolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Paths)) {
		query["Paths"] = request.Paths
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLifecycleRetrieveJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableIpv6)) {
		query["EnableIpv6"] = request.EnableIpv6
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateProtocolMountTargetWithOptions(request *CreateProtocolMountTargetRequest, runtime *util.RuntimeOptions) (_result *CreateProtocolMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FsetId)) {
		query["FsetId"] = request.FsetId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProtocolMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProtocolMountTarget(request *CreateProtocolMountTargetRequest) (_result *CreateProtocolMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProtocolMountTargetResponse{}
	_body, _err := client.CreateProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProtocolServiceWithOptions(request *CreateProtocolServiceRequest, runtime *util.RuntimeOptions) (_result *CreateProtocolServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolSpec)) {
		query["ProtocolSpec"] = request.ProtocolSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.Throughput)) {
		query["Throughput"] = request.Throughput
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProtocolService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProtocolService(request *CreateProtocolServiceRequest) (_result *CreateProtocolServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProtocolServiceResponse{}
	_body, _err := client.CreateProtocolServiceWithOptions(request, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRecycleBinDeleteJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRecycleBinDeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRecycleBinRestoreJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRecycleBinRestoreJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSnapshot"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessGroup"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessRule"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyId)) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteDataFlowWithOptions(request *DeleteDataFlowRequest, runtime *util.RuntimeOptions) (_result *DeleteDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataFlow(request *DeleteDataFlowRequest) (_result *DeleteDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataFlowResponse{}
	_body, _err := client.DeleteDataFlowWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteFilesetWithOptions(request *DeleteFilesetRequest, runtime *util.RuntimeOptions) (_result *DeleteFilesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FsetId)) {
		query["FsetId"] = request.FsetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFileset"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFileset(request *DeleteFilesetRequest) (_result *DeleteFilesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFilesetResponse{}
	_body, _err := client.DeleteFilesetWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLDAPConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.LifecyclePolicyName)) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLifecyclePolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteProtocolMountTargetWithOptions(request *DeleteProtocolMountTargetRequest, runtime *util.RuntimeOptions) (_result *DeleteProtocolMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ExportId)) {
		query["ExportId"] = request.ExportId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProtocolMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProtocolMountTarget(request *DeleteProtocolMountTargetRequest) (_result *DeleteProtocolMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProtocolMountTargetResponse{}
	_body, _err := client.DeleteProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProtocolServiceWithOptions(request *DeleteProtocolServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteProtocolServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProtocolService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProtocolService(request *DeleteProtocolServiceRequest) (_result *DeleteProtocolServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProtocolServiceResponse{}
	_body, _err := client.DeleteProtocolServiceWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshot"),
		Version:     tea.String("2017-06-26"),
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

func (client *Client) DescribeAccessGroupsWithOptions(request *DescribeAccessGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccessGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UseUTCDateTime)) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessGroups"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      tea.String("DescribeAccessRules"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyId)) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      tea.String("DescribeAutoSnapshotPolicies"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoSnapshotPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyIds)) {
		query["AutoSnapshotPolicyIds"] = request.AutoSnapshotPolicyIds
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemIds)) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      tea.String("DescribeAutoSnapshotTasks"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoSnapshotTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIP)) {
		query["ClientIP"] = request.ClientIP
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBlackListClients"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBlackListClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDataFlowTasksWithOptions(request *DescribeDataFlowTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDataFlowTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataFlowTasks"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataFlowTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataFlowTasks(request *DescribeDataFlowTasksRequest) (_result *DescribeDataFlowTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataFlowTasksResponse{}
	_body, _err := client.DescribeDataFlowTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataFlowsWithOptions(request *DescribeDataFlowsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataFlows"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataFlows(request *DescribeDataFlowsRequest) (_result *DescribeDataFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataFlowsResponse{}
	_body, _err := client.DescribeDataFlowsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDirQuotas"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDirQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeFileSystemStatisticsWithOptions(request *DescribeFileSystemStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeFileSystemStatisticsResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFileSystemStatistics"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileSystemStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeFileSystemsWithOptions(request *DescribeFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DescribeFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFileSystems"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeFilesetsWithOptions(request *DescribeFilesetsRequest, runtime *util.RuntimeOptions) (_result *DescribeFilesetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFilesets"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFilesetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFilesets(request *DescribeFilesetsRequest) (_result *DescribeFilesetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFilesetsResponse{}
	_body, _err := client.DescribeFilesetsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLDAPConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLifecyclePolicies"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLifecyclePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogAnalysis"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeMountTargetsWithOptions(request *DescribeMountTargetsRequest, runtime *util.RuntimeOptions) (_result *DescribeMountTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DualStackMountTargetDomain)) {
		query["DualStackMountTargetDomain"] = request.DualStackMountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
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
		Action:      tea.String("DescribeMountTargets"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeMountedClientsWithOptions(request *DescribeMountedClientsRequest, runtime *util.RuntimeOptions) (_result *DescribeMountedClientsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIP)) {
		query["ClientIP"] = request.ClientIP
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMountedClients"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMountedClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeProtocolMountTargetWithOptions(request *DescribeProtocolMountTargetRequest, runtime *util.RuntimeOptions) (_result *DescribeProtocolMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProtocolMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProtocolMountTarget(request *DescribeProtocolMountTargetRequest) (_result *DescribeProtocolMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtocolMountTargetResponse{}
	_body, _err := client.DescribeProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProtocolServiceWithOptions(request *DescribeProtocolServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeProtocolServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceIds)) {
		query["ProtocolServiceIds"] = request.ProtocolServiceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProtocolService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProtocolService(request *DescribeProtocolServiceRequest) (_result *DescribeProtocolServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtocolServiceResponse{}
	_body, _err := client.DescribeProtocolServiceWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-06-26"),
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

func (client *Client) DescribeSmbAclWithOptions(request *DescribeSmbAclRequest, runtime *util.RuntimeOptions) (_result *DescribeSmbAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSmbAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmbAcl(request *DescribeSmbAclRequest) (_result *DescribeSmbAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmbAclResponse{}
	_body, _err := client.DescribeSmbAclWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotType)) {
		query["SnapshotType"] = request.SnapshotType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSnapshots"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UseUTCDateTime)) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStoragePackages"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStoragePackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTags"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAndCleanRecycleBin"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAndCleanRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DisableSmbAclWithOptions(request *DisableSmbAclRequest, runtime *util.RuntimeOptions) (_result *DisableSmbAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSmbAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSmbAcl(request *DisableSmbAclRequest) (_result *DisableSmbAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableSmbAclResponse{}
	_body, _err := client.DisableSmbAclWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedDays)) {
		query["ReservedDays"] = request.ReservedDays
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableRecycleBin"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) EnableSmbAclWithOptions(request *EnableSmbAclRequest, runtime *util.RuntimeOptions) (_result *EnableSmbAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Keytab)) {
		query["Keytab"] = request.Keytab
	}

	if !tea.BoolValue(util.IsUnset(request.KeytabMd5)) {
		query["KeytabMd5"] = request.KeytabMd5
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSmbAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSmbAcl(request *EnableSmbAclRequest) (_result *EnableSmbAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSmbAclResponse{}
	_body, _err := client.EnableSmbAclWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDirectoryOrFileProperties"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDirectoryOrFilePropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecycleBinAttribute"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecycleBinAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryOnly)) {
		query["DirectoryOnly"] = request.DirectoryOnly
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDirectoriesAndFiles"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDirectoriesAndFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLifecycleRetrieveJobs"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLifecycleRetrieveJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecentlyRecycledDirectories"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecentlyRecycledDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecycleBinJobs"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecycleBinJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecycledDirectoriesAndFiles"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecycledDirectoriesAndFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccessGroup"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6SourceCidrIp)) {
		query["Ipv6SourceCidrIp"] = request.Ipv6SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RWAccessType)) {
		query["RWAccessType"] = request.RWAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrIp)) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessType)) {
		query["UserAccessType"] = request.UserAccessType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccessRule"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyId)) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoSnapshotPolicyName)) {
		query["AutoSnapshotPolicyName"] = request.AutoSnapshotPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatWeekdays)) {
		query["RepeatWeekdays"] = request.RepeatWeekdays
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.TimePoints)) {
		query["TimePoints"] = request.TimePoints
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAutoSnapshotPolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyDataFlowWithOptions(request *ModifyDataFlowRequest, runtime *util.RuntimeOptions) (_result *ModifyDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Throughput)) {
		query["Throughput"] = request.Throughput
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDataFlow(request *ModifyDataFlowRequest) (_result *ModifyDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataFlowResponse{}
	_body, _err := client.ModifyDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDataFlowAutoRefreshWithOptions(request *ModifyDataFlowAutoRefreshRequest, runtime *util.RuntimeOptions) (_result *ModifyDataFlowAutoRefreshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRefreshInterval)) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRefreshPolicy)) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataFlowAutoRefresh"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDataFlowAutoRefresh(request *ModifyDataFlowAutoRefreshRequest) (_result *ModifyDataFlowAutoRefreshResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataFlowAutoRefreshResponse{}
	_body, _err := client.ModifyDataFlowAutoRefreshWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyFilesetWithOptions(request *ModifyFilesetRequest, runtime *util.RuntimeOptions) (_result *ModifyFilesetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FsetId)) {
		query["FsetId"] = request.FsetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFileset"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFilesetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFileset(request *ModifyFilesetRequest) (_result *ModifyFilesetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFilesetResponse{}
	_body, _err := client.ModifyFilesetWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindDN)) {
		query["BindDN"] = request.BindDN
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchBase)) {
		query["SearchBase"] = request.SearchBase
	}

	if !tea.BoolValue(util.IsUnset(request.URI)) {
		query["URI"] = request.URI
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLDAPConfig"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLDAPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.LifecyclePolicyName)) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleRuleName)) {
		query["LifecycleRuleName"] = request.LifecycleRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLifecyclePolicy"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLifecyclePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.DualStackMountTargetDomain)) {
		query["DualStackMountTargetDomain"] = request.DualStackMountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountTargetDomain)) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyProtocolMountTargetWithOptions(request *ModifyProtocolMountTargetRequest, runtime *util.RuntimeOptions) (_result *ModifyProtocolMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ExportId)) {
		query["ExportId"] = request.ExportId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtocolMountTarget"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtocolMountTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtocolMountTarget(request *ModifyProtocolMountTargetRequest) (_result *ModifyProtocolMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtocolMountTargetResponse{}
	_body, _err := client.ModifyProtocolMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProtocolServiceWithOptions(request *ModifyProtocolServiceRequest, runtime *util.RuntimeOptions) (_result *ModifyProtocolServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolServiceId)) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtocolService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtocolServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtocolService(request *ModifyProtocolServiceRequest) (_result *ModifyProtocolServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtocolServiceResponse{}
	_body, _err := client.ModifyProtocolServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySmbAclWithOptions(request *ModifySmbAclRequest, runtime *util.RuntimeOptions) (_result *ModifySmbAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableAnonymousAccess)) {
		query["EnableAnonymousAccess"] = request.EnableAnonymousAccess
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptData)) {
		query["EncryptData"] = request.EncryptData
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.HomeDirPath)) {
		query["HomeDirPath"] = request.HomeDirPath
	}

	if !tea.BoolValue(util.IsUnset(request.Keytab)) {
		query["Keytab"] = request.Keytab
	}

	if !tea.BoolValue(util.IsUnset(request.KeytabMd5)) {
		query["KeytabMd5"] = request.KeytabMd5
	}

	if !tea.BoolValue(util.IsUnset(request.RejectUnencryptedAccess)) {
		query["RejectUnencryptedAccess"] = request.RejectUnencryptedAccess
	}

	if !tea.BoolValue(util.IsUnset(request.SuperAdminSid)) {
		query["SuperAdminSid"] = request.SuperAdminSid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySmbAcl"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySmbAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySmbAcl(request *ModifySmbAclRequest) (_result *ModifySmbAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySmbAclResponse{}
	_body, _err := client.ModifySmbAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenNASServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenNASServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenNASService"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenNASServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIP)) {
		query["ClientIP"] = request.ClientIP
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveClientFromBlackList"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveClientFromBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTags"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryLifecycleRetrieveJob"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileCountLimit)) {
		query["FileCountLimit"] = request.FileCountLimit
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaType)) {
		query["QuotaType"] = request.QuotaType
	}

	if !tea.BoolValue(util.IsUnset(request.SizeLimit)) {
		query["SizeLimit"] = request.SizeLimit
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDirQuota"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDirQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) StartDataFlowWithOptions(request *StartDataFlowRequest, runtime *util.RuntimeOptions) (_result *StartDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDataFlow(request *StartDataFlowRequest) (_result *StartDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDataFlowResponse{}
	_body, _err := client.StartDataFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDataFlowWithOptions(request *StopDataFlowRequest, runtime *util.RuntimeOptions) (_result *StopDataFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataFlowId)) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDataFlow"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDataFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDataFlow(request *StopDataFlowRequest) (_result *StopDataFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDataFlowResponse{}
	_body, _err := client.StopDataFlowWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRecycleBinAttribute"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRecycleBinAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		query["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeFileSystem"),
		Version:     tea.String("2017-06-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
