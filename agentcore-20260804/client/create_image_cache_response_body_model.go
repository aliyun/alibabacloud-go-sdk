// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateImageCacheResponseBody
	GetCode() *string
	SetData(v *CreateImageCacheResponseBodyData) *CreateImageCacheResponseBody
	GetData() *CreateImageCacheResponseBodyData
	SetHttpStatusCode(v int32) *CreateImageCacheResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateImageCacheResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateImageCacheResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateImageCacheResponseBody
	GetSuccess() *bool
}

type CreateImageCacheResponseBody struct {
	// The business response code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result of creating the image cache.
	//
	// This parameter is required.
	Data *CreateImageCacheResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateImageCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateImageCacheResponseBody) GetData() *CreateImageCacheResponseBodyData {
	return s.Data
}

func (s *CreateImageCacheResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateImageCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateImageCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageCacheResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageCacheResponseBody) SetCode(v string) *CreateImageCacheResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetData(v *CreateImageCacheResponseBodyData) *CreateImageCacheResponseBody {
	s.Data = v
	return s
}

func (s *CreateImageCacheResponseBody) SetHttpStatusCode(v int32) *CreateImageCacheResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetMessage(v string) *CreateImageCacheResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetRequestId(v string) *CreateImageCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetSuccess(v bool) *CreateImageCacheResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageCacheResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageCacheResponseBodyData struct {
	// The ID of the source ACR Enterprise instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The deduplicated number of agents that have useImageCache set to true and are associated with this logical resource in the current workspace. The value is greater than or equal to 0. This set is the same as the set filtered by imageCacheId in ListAllAgents. Agents with the feature disabled are not counted, and this value is independent of the number of occupied instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AgentCount *int64 `json:"agentCount,omitempty" xml:"agentCount,omitempty"`
	// The actual cache size in bytes. The known value is greater than or equal to 0. If the value is unknown, it is null. The compressed size of the source image cannot be used as a substitute.
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
	// Indicates whether deletion is allowed. The value is true when there are no configuration references with the switch enabled, no valid instances occupying the resource, and the resource is not being deleted in the current workspace. If the reference status is unknown, the query should fail, and deletion cannot be allowed by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// The creation time of the logical resource in the current workspace, in RFC3339 format with time zone. The time is not refreshed when the content of the latest tag is updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-10T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The reason why deletion is blocked. The value is null when canDelete is true. Otherwise, the value is returned based on the priority of DELETING, AGENT_ENABLED, and INSTANCES_RELEASING. Even if the number of agents is 0, old instances may still occupy the resource.
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
	// The current content digest in the format of sha256: followed by a 64-character lowercase hexadecimal string. It only describes the currently registered content and does not prove that old instances have been switched.
	//
	// This parameter is required.
	//
	// example:
	//
	// sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	Digest *string `json:"digest,omitempty" xml:"digest,omitempty"`
	// The reason for the cache failure. This is a non-null object when the status is FAILED, and null for other statuses. It only describes the current content and the current creation round.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	Error *CreateImageCacheResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The logical resource ID manageable by the current workspace, which is not a shared physical cache ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-example
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// The domain name of the source image, excluding the protocol, path, and credentials.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-registry.cn-hangzhou.cr.aliyuncs.com
	Registry *string `json:"registry,omitempty" xml:"registry,omitempty"`
	// The path of the source image repository in the namespace/repository format, which already includes the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// team/app
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// The status of the image cache. Valid values: PENDING: creating, READY: available, FAILED: abnormal, DELETING: deleting in the current workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the source image. Case is preserved. Regular tags cannot be registered with different digests, with the exact lowercase latest tag being the only exception.
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

func (s CreateImageCacheResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponseBodyData) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateImageCacheResponseBodyData) GetAgentCount() *int64 {
	return s.AgentCount
}

func (s *CreateImageCacheResponseBodyData) GetCacheSizeBytes() *int64 {
	return s.CacheSizeBytes
}

func (s *CreateImageCacheResponseBodyData) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *CreateImageCacheResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateImageCacheResponseBodyData) GetDeleteBlockReason() *string {
	return s.DeleteBlockReason
}

func (s *CreateImageCacheResponseBodyData) GetDigest() *string {
	return s.Digest
}

func (s *CreateImageCacheResponseBodyData) GetError() *CreateImageCacheResponseBodyDataError {
	return s.Error
}

func (s *CreateImageCacheResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *CreateImageCacheResponseBodyData) GetRegistry() *string {
	return s.Registry
}

func (s *CreateImageCacheResponseBodyData) GetRepository() *string {
	return s.Repository
}

func (s *CreateImageCacheResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateImageCacheResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *CreateImageCacheResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateImageCacheResponseBodyData) SetAcrInstanceId(v string) *CreateImageCacheResponseBodyData {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetAgentCount(v int64) *CreateImageCacheResponseBodyData {
	s.AgentCount = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetCacheSizeBytes(v int64) *CreateImageCacheResponseBodyData {
	s.CacheSizeBytes = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetCanDelete(v bool) *CreateImageCacheResponseBodyData {
	s.CanDelete = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetCreatedAt(v string) *CreateImageCacheResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetDeleteBlockReason(v string) *CreateImageCacheResponseBodyData {
	s.DeleteBlockReason = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetDigest(v string) *CreateImageCacheResponseBodyData {
	s.Digest = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetError(v *CreateImageCacheResponseBodyDataError) *CreateImageCacheResponseBodyData {
	s.Error = v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetImageId(v string) *CreateImageCacheResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetRegistry(v string) *CreateImageCacheResponseBodyData {
	s.Registry = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetRepository(v string) *CreateImageCacheResponseBodyData {
	s.Repository = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetStatus(v string) *CreateImageCacheResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetTag(v string) *CreateImageCacheResponseBodyData {
	s.Tag = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) SetWorkspaceId(v string) *CreateImageCacheResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateImageCacheResponseBodyData) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageCacheResponseBodyDataError struct {
	// The secure and non-null error code for the current creation failure.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACSCacheBuildFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The secure and non-null display reason for the creation failure, excluding credentials, internal addresses, or stack traces.
	//
	// This parameter is required.
	//
	// example:
	//
	// Failed to create the image cache. Please try again
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateImageCacheResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *CreateImageCacheResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *CreateImageCacheResponseBodyDataError) SetCode(v string) *CreateImageCacheResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *CreateImageCacheResponseBodyDataError) SetMessage(v string) *CreateImageCacheResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *CreateImageCacheResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}
