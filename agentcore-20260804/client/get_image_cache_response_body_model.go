// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageCacheResponseBody
	GetCode() *string
	SetData(v *GetImageCacheResponseBodyData) *GetImageCacheResponseBody
	GetData() *GetImageCacheResponseBodyData
	SetHttpStatusCode(v int32) *GetImageCacheResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetImageCacheResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageCacheResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageCacheResponseBody
	GetSuccess() *bool
}

type GetImageCacheResponseBody struct {
	// The business response code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The query result of the image cache details.
	//
	// This parameter is required.
	Data *GetImageCacheResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful. This does not indicate that the cache creation is complete.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetImageCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageCacheResponseBody) GetData() *GetImageCacheResponseBodyData {
	return s.Data
}

func (s *GetImageCacheResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetImageCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageCacheResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageCacheResponseBody) SetCode(v string) *GetImageCacheResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageCacheResponseBody) SetData(v *GetImageCacheResponseBodyData) *GetImageCacheResponseBody {
	s.Data = v
	return s
}

func (s *GetImageCacheResponseBody) SetHttpStatusCode(v int32) *GetImageCacheResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageCacheResponseBody) SetMessage(v string) *GetImageCacheResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageCacheResponseBody) SetRequestId(v string) *GetImageCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageCacheResponseBody) SetSuccess(v bool) *GetImageCacheResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageCacheResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageCacheResponseBodyData struct {
	// The source ACR Enterprise instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The number of unique agents in the current workspace that have useImageCache set to true and are associated with this logical resource. The value is greater than or equal to 0. This uses the same collection as the imageCacheId filter in ListAllAgents. Disabled agents are not counted, regardless of the number of occupied instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AgentCount *int64 `json:"agentCount,omitempty" xml:"agentCount,omitempty"`
	// The actual cache size in bytes. The known value is greater than or equal to 0. If unknown, the value is null. The compressed size of the source image cannot be used as a substitute.
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
	// Indicates whether deletion is allowed. The value is true if there are no configuration references with the switch enabled, no valid instances occupying the resource, and the resource is not being deleted in the current workspace. If the reference status is unknown, the query fails instead of allowing deletion by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// The creation time of the logical resource in the current workspace, in RFC3339 format with time zone. Content updates for the latest tag do not refresh this time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-10T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The reason why deletion is blocked. The value is null if canDelete is true. Otherwise, the value is returned based on the priority of DELETING, AGENT_ENABLED, and INSTANCES_RELEASING. Even if the agent count is 0, old instances may still occupy the resource.
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
	// The cache failure reason. This must be a non-null object when the status is FAILED, and null for other statuses. It only describes the current content and the current creation round.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	Error *GetImageCacheResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
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
	// - PENDING: Creating.
	//
	// - READY: Available.
	//
	// - FAILED: Abnormal.
	//
	// - DELETING: Deleting in the current workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The source image tag with case preserved. Regular tags cannot be registered with different digests, except for the exact lowercase latest tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The workspace ID, which must match the requested workspace. Resources with the same content from other workspaces are not returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetImageCacheResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponseBodyData) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *GetImageCacheResponseBodyData) GetAgentCount() *int64 {
	return s.AgentCount
}

func (s *GetImageCacheResponseBodyData) GetCacheSizeBytes() *int64 {
	return s.CacheSizeBytes
}

func (s *GetImageCacheResponseBodyData) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *GetImageCacheResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetImageCacheResponseBodyData) GetDeleteBlockReason() *string {
	return s.DeleteBlockReason
}

func (s *GetImageCacheResponseBodyData) GetDigest() *string {
	return s.Digest
}

func (s *GetImageCacheResponseBodyData) GetError() *GetImageCacheResponseBodyDataError {
	return s.Error
}

func (s *GetImageCacheResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageCacheResponseBodyData) GetRegistry() *string {
	return s.Registry
}

func (s *GetImageCacheResponseBodyData) GetRepository() *string {
	return s.Repository
}

func (s *GetImageCacheResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetImageCacheResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *GetImageCacheResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetImageCacheResponseBodyData) SetAcrInstanceId(v string) *GetImageCacheResponseBodyData {
	s.AcrInstanceId = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetAgentCount(v int64) *GetImageCacheResponseBodyData {
	s.AgentCount = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetCacheSizeBytes(v int64) *GetImageCacheResponseBodyData {
	s.CacheSizeBytes = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetCanDelete(v bool) *GetImageCacheResponseBodyData {
	s.CanDelete = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetCreatedAt(v string) *GetImageCacheResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetDeleteBlockReason(v string) *GetImageCacheResponseBodyData {
	s.DeleteBlockReason = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetDigest(v string) *GetImageCacheResponseBodyData {
	s.Digest = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetError(v *GetImageCacheResponseBodyDataError) *GetImageCacheResponseBodyData {
	s.Error = v
	return s
}

func (s *GetImageCacheResponseBodyData) SetImageId(v string) *GetImageCacheResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetRegistry(v string) *GetImageCacheResponseBodyData {
	s.Registry = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetRepository(v string) *GetImageCacheResponseBodyData {
	s.Repository = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetStatus(v string) *GetImageCacheResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetTag(v string) *GetImageCacheResponseBodyData {
	s.Tag = &v
	return s
}

func (s *GetImageCacheResponseBodyData) SetWorkspaceId(v string) *GetImageCacheResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetImageCacheResponseBodyData) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageCacheResponseBodyDataError struct {
	// The secure and non-empty error code for the current creation failure.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACSCacheBuildFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The secure and non-empty display message for the creation failure, excluding credentials, internal addresses, or stack traces.
	//
	// This parameter is required.
	//
	// example:
	//
	// Image cache creation failed. Please try again
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetImageCacheResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *GetImageCacheResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *GetImageCacheResponseBodyDataError) SetCode(v string) *GetImageCacheResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *GetImageCacheResponseBodyDataError) SetMessage(v string) *GetImageCacheResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *GetImageCacheResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}
