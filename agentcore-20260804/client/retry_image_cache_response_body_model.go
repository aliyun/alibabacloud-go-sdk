// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryImageCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryImageCacheResponseBody
	GetCode() *string
	SetData(v *RetryImageCacheResponseBodyData) *RetryImageCacheResponseBody
	GetData() *RetryImageCacheResponseBodyData
	SetHttpStatusCode(v int32) *RetryImageCacheResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RetryImageCacheResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryImageCacheResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryImageCacheResponseBody
	GetSuccess() *bool
}

type RetryImageCacheResponseBody struct {
	// The business response code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The retry image cache result.
	//
	// This parameter is required.
	Data *RetryImageCacheResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used for troubleshooting.
	//
	// This parameter is required.
	//
	// example:
	//
	// req-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. This does not mean that the cache production is complete.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RetryImageCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *RetryImageCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryImageCacheResponseBody) GetData() *RetryImageCacheResponseBodyData {
	return s.Data
}

func (s *RetryImageCacheResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RetryImageCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryImageCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryImageCacheResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryImageCacheResponseBody) SetCode(v string) *RetryImageCacheResponseBody {
	s.Code = &v
	return s
}

func (s *RetryImageCacheResponseBody) SetData(v *RetryImageCacheResponseBodyData) *RetryImageCacheResponseBody {
	s.Data = v
	return s
}

func (s *RetryImageCacheResponseBody) SetHttpStatusCode(v int32) *RetryImageCacheResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RetryImageCacheResponseBody) SetMessage(v string) *RetryImageCacheResponseBody {
	s.Message = &v
	return s
}

func (s *RetryImageCacheResponseBody) SetRequestId(v string) *RetryImageCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryImageCacheResponseBody) SetSuccess(v bool) *RetryImageCacheResponseBody {
	s.Success = &v
	return s
}

func (s *RetryImageCacheResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetryImageCacheResponseBodyData struct {
	// The source ACR Enterprise instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The deduplicated number of agents in the current workspace that have useImageCache set to true and are associated with this logical resource. The value is greater than or equal to 0. This uses the same set as the imageCacheId filter in ListAllAgents. Agents with the feature disabled are not counted, regardless of the number of occupied instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AgentCount *int64 `json:"agentCount,omitempty" xml:"agentCount,omitempty"`
	// The actual cache size in bytes. Known values are greater than or equal to 0. If unknown, the value is null. The compressed size of the source image cannot be used as a substitute.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1073741824
	CacheSizeBytes *int64 `json:"cacheSizeBytes,omitempty" xml:"cacheSizeBytes,omitempty"`
	// Indicates whether deletion is allowed. The value is true when the current workspace has no configuration references with the switch enabled, no valid instances are occupying the resource, and the resource is not being deleted. If the reference status is unknown, the query fails instead of allowing deletion by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// The creation time of the logical resource in the current workspace, in RFC3339 format with time zone. Content updates for the "latest" tag do not refresh this time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-10T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The reason why deletion is blocked. The value is null when canDelete is true. Otherwise, the value is returned based on the priority of DELETING, AGENT_ENABLED, and INSTANCES_RELEASING. Old instances may still occupy the resource even if the agent count is 0.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// INSTANCES_RELEASING
	DeleteBlockReason *string `json:"deleteBlockReason,omitempty" xml:"deleteBlockReason,omitempty"`
	// The current content digest in the sha256: format followed by a 64-character lowercase hexadecimal string. It only describes the currently registered content and does not prove that old instances have been switched.
	//
	// This parameter is required.
	//
	// example:
	//
	// sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	Digest *string `json:"digest,omitempty" xml:"digest,omitempty"`
	// The reason for cache failure. This is a non-empty object when the status is FAILED, and null for other statuses. It only describes the current content and the current production round.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	Error *RetryImageCacheResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The logical resource ID manageable by the current workspace, not the shared physical cache ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-example
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// The source image domain name, excluding the protocol, path, and credentials.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-registry.cn-hangzhou.cr.aliyuncs.com
	Registry *string `json:"registry,omitempty" xml:"registry,omitempty"`
	// The source image repository path in the namespace/repository format, which includes the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// team/app
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// The image cache status. Valid values:
	//
	// - PENDING: The cache is being created.
	//
	// - READY: The cache is available.
	//
	// - FAILED: An exception occurred.
	//
	// - DELETING: The cache is being deleted in the current workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The source image tag with case preserved. Regular tags cannot be registered with different digests, except for the exact lowercase "latest" tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The workspace ID, which is consistent with the requested workspace. Resources with the same content from other workspaces are not returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s RetryImageCacheResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetryImageCacheResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetryImageCacheResponseBodyData) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *RetryImageCacheResponseBodyData) GetAgentCount() *int64 {
	return s.AgentCount
}

func (s *RetryImageCacheResponseBodyData) GetCacheSizeBytes() *int64 {
	return s.CacheSizeBytes
}

func (s *RetryImageCacheResponseBodyData) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *RetryImageCacheResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *RetryImageCacheResponseBodyData) GetDeleteBlockReason() *string {
	return s.DeleteBlockReason
}

func (s *RetryImageCacheResponseBodyData) GetDigest() *string {
	return s.Digest
}

func (s *RetryImageCacheResponseBodyData) GetError() *RetryImageCacheResponseBodyDataError {
	return s.Error
}

func (s *RetryImageCacheResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *RetryImageCacheResponseBodyData) GetRegistry() *string {
	return s.Registry
}

func (s *RetryImageCacheResponseBodyData) GetRepository() *string {
	return s.Repository
}

func (s *RetryImageCacheResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *RetryImageCacheResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *RetryImageCacheResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RetryImageCacheResponseBodyData) SetAcrInstanceId(v string) *RetryImageCacheResponseBodyData {
	s.AcrInstanceId = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetAgentCount(v int64) *RetryImageCacheResponseBodyData {
	s.AgentCount = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetCacheSizeBytes(v int64) *RetryImageCacheResponseBodyData {
	s.CacheSizeBytes = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetCanDelete(v bool) *RetryImageCacheResponseBodyData {
	s.CanDelete = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetCreatedAt(v string) *RetryImageCacheResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetDeleteBlockReason(v string) *RetryImageCacheResponseBodyData {
	s.DeleteBlockReason = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetDigest(v string) *RetryImageCacheResponseBodyData {
	s.Digest = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetError(v *RetryImageCacheResponseBodyDataError) *RetryImageCacheResponseBodyData {
	s.Error = v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetImageId(v string) *RetryImageCacheResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetRegistry(v string) *RetryImageCacheResponseBodyData {
	s.Registry = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetRepository(v string) *RetryImageCacheResponseBodyData {
	s.Repository = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetStatus(v string) *RetryImageCacheResponseBodyData {
	s.Status = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetTag(v string) *RetryImageCacheResponseBodyData {
	s.Tag = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) SetWorkspaceId(v string) *RetryImageCacheResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *RetryImageCacheResponseBodyData) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetryImageCacheResponseBodyDataError struct {
	// The secure and non-empty error code for the current production failure.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACSCacheBuildFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The secure and non-empty display reason for the production failure, excluding credentials, internal addresses, or stack traces.
	//
	// This parameter is required.
	//
	// example:
	//
	// Image cache production failed. Please retry
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s RetryImageCacheResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s RetryImageCacheResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *RetryImageCacheResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *RetryImageCacheResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *RetryImageCacheResponseBodyDataError) SetCode(v string) *RetryImageCacheResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *RetryImageCacheResponseBodyDataError) SetMessage(v string) *RetryImageCacheResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *RetryImageCacheResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}
