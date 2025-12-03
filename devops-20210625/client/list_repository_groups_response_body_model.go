// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListRepositoryGroupsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRepositoryGroupsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoryGroupsResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoryGroupsResponseBodyResult) *ListRepositoryGroupsResponseBody
	GetResult() []*ListRepositoryGroupsResponseBodyResult
	SetSuccess(v bool) *ListRepositoryGroupsResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListRepositoryGroupsResponseBody
	GetTotal() *int64
}

type ListRepositoryGroupsResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRepositoryGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoryGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRepositoryGroupsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoryGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryGroupsResponseBody) GetResult() []*ListRepositoryGroupsResponseBodyResult {
	return s.Result
}

func (s *ListRepositoryGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoryGroupsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListRepositoryGroupsResponseBody) SetErrorCode(v string) *ListRepositoryGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryGroupsResponseBody) SetErrorMessage(v string) *ListRepositoryGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryGroupsResponseBody) SetRequestId(v string) *ListRepositoryGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryGroupsResponseBody) SetResult(v []*ListRepositoryGroupsResponseBodyResult) *ListRepositoryGroupsResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryGroupsResponseBody) SetSuccess(v bool) *ListRepositoryGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryGroupsResponseBody) SetTotal(v int64) *ListRepositoryGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryGroupsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRepositoryGroupsResponseBodyResult struct {
	// example:
	//
	// 40
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt   *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 10
	GroupCount *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	// example:
	//
	// 19285
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-codeup
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// 19230
	OwnerId *int64 `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// 26842
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// test-codeup
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 60de7a6852743a5162b5f957/test-codeup
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// 10
	ProjectCount *int64 `json:"projectCount,omitempty" xml:"projectCount,omitempty"`
	// example:
	//
	// Group
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// example:
	//
	// ""
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s ListRepositoryGroupsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryGroupsResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListRepositoryGroupsResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListRepositoryGroupsResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRepositoryGroupsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListRepositoryGroupsResponseBodyResult) GetGroupCount() *int64 {
	return s.GroupCount
}

func (s *ListRepositoryGroupsResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListRepositoryGroupsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListRepositoryGroupsResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *ListRepositoryGroupsResponseBodyResult) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListRepositoryGroupsResponseBodyResult) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListRepositoryGroupsResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *ListRepositoryGroupsResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *ListRepositoryGroupsResponseBodyResult) GetProjectCount() *int64 {
	return s.ProjectCount
}

func (s *ListRepositoryGroupsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListRepositoryGroupsResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListRepositoryGroupsResponseBodyResult) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *ListRepositoryGroupsResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *ListRepositoryGroupsResponseBodyResult) SetAccessLevel(v int32) *ListRepositoryGroupsResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetAvatarUrl(v string) *ListRepositoryGroupsResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetCreatedAt(v string) *ListRepositoryGroupsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetDescription(v string) *ListRepositoryGroupsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetGroupCount(v int64) *ListRepositoryGroupsResponseBodyResult {
	s.GroupCount = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetId(v int64) *ListRepositoryGroupsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetName(v string) *ListRepositoryGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetNameWithNamespace(v string) *ListRepositoryGroupsResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetOwnerId(v int64) *ListRepositoryGroupsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetParentId(v int64) *ListRepositoryGroupsResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetPath(v string) *ListRepositoryGroupsResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetPathWithNamespace(v string) *ListRepositoryGroupsResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetProjectCount(v int64) *ListRepositoryGroupsResponseBodyResult {
	s.ProjectCount = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetType(v string) *ListRepositoryGroupsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetUpdatedAt(v string) *ListRepositoryGroupsResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetVisibilityLevel(v int32) *ListRepositoryGroupsResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) SetWebUrl(v string) *ListRepositoryGroupsResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *ListRepositoryGroupsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
