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
	// The image cache details.
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
	// The request ID, used for troubleshooting.
	//
	// This parameter is required.
	//
	// example:
	//
	// req-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. This does not mean that the cache creation is complete.
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
	// The source ACR instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The number of agents using the cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AgentCount *int64 `json:"agentCount,omitempty" xml:"agentCount,omitempty"`
	// The cache size in bytes.
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
	// Indicates whether deletion is allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// The resource creation time. The format is RFC 3339. The Z at the end of the example indicates the UTC time zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-10T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The reason for blocking deletion.
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
	// The cache failure reason.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	Error *GetImageCacheResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The image cache ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-example
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// The source image domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-registry.cn-hangzhou.cr.aliyuncs.com
	Registry *string `json:"registry,omitempty" xml:"registry,omitempty"`
	// The source image repository path.
	//
	// This parameter is required.
	//
	// example:
	//
	// team/app
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// The image cache status.
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The source image tag.
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
	// The error code for creation failure.
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
	// Image cache creation failed. Try again
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
