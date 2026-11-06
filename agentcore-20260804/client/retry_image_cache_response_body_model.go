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
	// The result of retrying image cache creation.
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
	// Indicates whether the request is successful. A successful request does not mean that the cache creation is complete.
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
	// The ID of the source ACR instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The number of agents that use the cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AgentCount *int64 `json:"agentCount,omitempty" xml:"agentCount,omitempty"`
	// The size of the cache in bytes.
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
	// Indicates whether the cache can be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// The time when the resource was created. The time follows the RFC 3339 format. The Z at the end of the example indicates the UTC time zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-10T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The reason why the cache cannot be deleted.
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
	// The current content digest.
	//
	// This parameter is required.
	//
	// example:
	//
	// sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	Digest *string `json:"digest,omitempty" xml:"digest,omitempty"`
	// The reason for the cache failure.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	Error *RetryImageCacheResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The image cache ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-example
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// The domain name of the source image.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-registry.cn-hangzhou.cr.aliyuncs.com
	Registry *string `json:"registry,omitempty" xml:"registry,omitempty"`
	// The path of the source image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// team/app
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// The cache status.
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the source image.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The workspace ID.
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
	// The error code for the creation failure.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACSCacheBuildFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the creation failure.
	//
	// This parameter is required.
	//
	// example:
	//
	// The image cache creation failed. Please retry
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
