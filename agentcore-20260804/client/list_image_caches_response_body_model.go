// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListImageCachesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListImageCachesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListImageCachesResponseBodyItems) *ListImageCachesResponseBody
	GetItems() []*ListImageCachesResponseBodyItems
	SetMaxResults(v int32) *ListImageCachesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListImageCachesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListImageCachesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListImageCachesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImageCachesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListImageCachesResponseBody
	GetTotalCount() *int64
}

type ListImageCachesResponseBody struct {
	// The business response code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The records on the current page. An empty array is returned if no matching items are found.
	//
	// This parameter is required.
	Items []*ListImageCachesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The actual page size used for this request, which is not the number of items on the current page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The opaque cursor for the next page. An empty string indicates the end of the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// opaque-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// The total number of records under the current filter conditions. Do not use this value to determine whether the traversal is complete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListImageCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageCachesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListImageCachesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListImageCachesResponseBody) GetItems() []*ListImageCachesResponseBodyItems {
	return s.Items
}

func (s *ListImageCachesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImageCachesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListImageCachesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImageCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageCachesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImageCachesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListImageCachesResponseBody) SetCode(v string) *ListImageCachesResponseBody {
	s.Code = &v
	return s
}

func (s *ListImageCachesResponseBody) SetHttpStatusCode(v int32) *ListImageCachesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListImageCachesResponseBody) SetItems(v []*ListImageCachesResponseBodyItems) *ListImageCachesResponseBody {
	s.Items = v
	return s
}

func (s *ListImageCachesResponseBody) SetMaxResults(v int32) *ListImageCachesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListImageCachesResponseBody) SetMessage(v string) *ListImageCachesResponseBody {
	s.Message = &v
	return s
}

func (s *ListImageCachesResponseBody) SetNextToken(v string) *ListImageCachesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImageCachesResponseBody) SetRequestId(v string) *ListImageCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageCachesResponseBody) SetSuccess(v bool) *ListImageCachesResponseBody {
	s.Success = &v
	return s
}

func (s *ListImageCachesResponseBody) SetTotalCount(v int64) *ListImageCachesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImageCachesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImageCachesResponseBodyItems struct {
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
	// The resource creation time. The time is in RFC 3339 format. The Z at the end of the example indicates the UTC time zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-10T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The reason why deletion is blocked.
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
	Error *ListImageCachesResponseBodyItemsError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
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

func (s ListImageCachesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListImageCachesResponseBodyItems) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *ListImageCachesResponseBodyItems) GetAgentCount() *int64 {
	return s.AgentCount
}

func (s *ListImageCachesResponseBodyItems) GetCacheSizeBytes() *int64 {
	return s.CacheSizeBytes
}

func (s *ListImageCachesResponseBodyItems) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ListImageCachesResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListImageCachesResponseBodyItems) GetDeleteBlockReason() *string {
	return s.DeleteBlockReason
}

func (s *ListImageCachesResponseBodyItems) GetDigest() *string {
	return s.Digest
}

func (s *ListImageCachesResponseBodyItems) GetError() *ListImageCachesResponseBodyItemsError {
	return s.Error
}

func (s *ListImageCachesResponseBodyItems) GetImageId() *string {
	return s.ImageId
}

func (s *ListImageCachesResponseBodyItems) GetRegistry() *string {
	return s.Registry
}

func (s *ListImageCachesResponseBodyItems) GetRepository() *string {
	return s.Repository
}

func (s *ListImageCachesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListImageCachesResponseBodyItems) GetTag() *string {
	return s.Tag
}

func (s *ListImageCachesResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListImageCachesResponseBodyItems) SetAcrInstanceId(v string) *ListImageCachesResponseBodyItems {
	s.AcrInstanceId = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetAgentCount(v int64) *ListImageCachesResponseBodyItems {
	s.AgentCount = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetCacheSizeBytes(v int64) *ListImageCachesResponseBodyItems {
	s.CacheSizeBytes = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetCanDelete(v bool) *ListImageCachesResponseBodyItems {
	s.CanDelete = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetCreatedAt(v string) *ListImageCachesResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetDeleteBlockReason(v string) *ListImageCachesResponseBodyItems {
	s.DeleteBlockReason = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetDigest(v string) *ListImageCachesResponseBodyItems {
	s.Digest = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetError(v *ListImageCachesResponseBodyItemsError) *ListImageCachesResponseBodyItems {
	s.Error = v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetImageId(v string) *ListImageCachesResponseBodyItems {
	s.ImageId = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetRegistry(v string) *ListImageCachesResponseBodyItems {
	s.Registry = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetRepository(v string) *ListImageCachesResponseBodyItems {
	s.Repository = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetStatus(v string) *ListImageCachesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetTag(v string) *ListImageCachesResponseBodyItems {
	s.Tag = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) SetWorkspaceId(v string) *ListImageCachesResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListImageCachesResponseBodyItems) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListImageCachesResponseBodyItemsError struct {
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
	// Image cache creation failed. Try again
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListImageCachesResponseBodyItemsError) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesResponseBodyItemsError) GoString() string {
	return s.String()
}

func (s *ListImageCachesResponseBodyItemsError) GetCode() *string {
	return s.Code
}

func (s *ListImageCachesResponseBodyItemsError) GetMessage() *string {
	return s.Message
}

func (s *ListImageCachesResponseBodyItemsError) SetCode(v string) *ListImageCachesResponseBodyItemsError {
	s.Code = &v
	return s
}

func (s *ListImageCachesResponseBodyItemsError) SetMessage(v string) *ListImageCachesResponseBodyItemsError {
	s.Message = &v
	return s
}

func (s *ListImageCachesResponseBodyItemsError) Validate() error {
	return dara.Validate(s)
}
