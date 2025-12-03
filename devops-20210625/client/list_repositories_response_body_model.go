// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int32) *ListRepositoriesResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *ListRepositoriesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoriesResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoriesResponseBodyResult) *ListRepositoriesResponseBody
	GetResult() []*ListRepositoriesResponseBodyResult
	SetSuccess(v bool) *ListRepositoriesResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListRepositoriesResponseBody
	GetTotal() *int64
}

type ListRepositoriesResponseBody struct {
	ErrorCode *int32 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 4D6AF7CC-B43B-5454-86AB-023D25E44868
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRepositoriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 149
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *ListRepositoriesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoriesResponseBody) GetResult() []*ListRepositoriesResponseBodyResult {
	return s.Result
}

func (s *ListRepositoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoriesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListRepositoriesResponseBody) SetErrorCode(v int32) *ListRepositoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetErrorMessage(v string) *ListRepositoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetRequestId(v string) *ListRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetResult(v []*ListRepositoriesResponseBodyResult) *ListRepositoriesResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoriesResponseBody) SetSuccess(v bool) *ListRepositoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetTotal(v int64) *ListRepositoriesResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoriesResponseBody) Validate() error {
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

type ListRepositoriesResponseBodyResult struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 40
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// false
	Archive   *bool   `json:"archive,omitempty" xml:"archive,omitempty"`
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	CreatedAt    *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	ImportStatus *string `json:"importStatus,omitempty" xml:"importStatus,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	LastActivityAt *string `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	// example:
	//
	// codeupTest
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// 745
	NamespaceId *int64 `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// example:
	//
	// test-codeup
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org/test-codeup
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// false
	Star *bool `json:"star,omitempty" xml:"star,omitempty"`
	// example:
	//
	// 0
	StarCount *int64 `json:"starCount,omitempty" xml:"starCount,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *string `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// example:
	//
	// ""
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s ListRepositoriesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListRepositoriesResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListRepositoriesResponseBodyResult) GetArchive() *bool {
	return s.Archive
}

func (s *ListRepositoriesResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListRepositoriesResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRepositoriesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListRepositoriesResponseBodyResult) GetImportStatus() *string {
	return s.ImportStatus
}

func (s *ListRepositoriesResponseBodyResult) GetLastActivityAt() *string {
	return s.LastActivityAt
}

func (s *ListRepositoriesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListRepositoriesResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *ListRepositoriesResponseBodyResult) GetNamespaceId() *int64 {
	return s.NamespaceId
}

func (s *ListRepositoriesResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *ListRepositoriesResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *ListRepositoriesResponseBodyResult) GetStar() *bool {
	return s.Star
}

func (s *ListRepositoriesResponseBodyResult) GetStarCount() *int64 {
	return s.StarCount
}

func (s *ListRepositoriesResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListRepositoriesResponseBodyResult) GetVisibilityLevel() *string {
	return s.VisibilityLevel
}

func (s *ListRepositoriesResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *ListRepositoriesResponseBodyResult) SetId(v int64) *ListRepositoriesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetAccessLevel(v int32) *ListRepositoriesResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetArchive(v bool) *ListRepositoriesResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetAvatarUrl(v string) *ListRepositoriesResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetCreatedAt(v string) *ListRepositoriesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetDescription(v string) *ListRepositoriesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetImportStatus(v string) *ListRepositoriesResponseBodyResult {
	s.ImportStatus = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetLastActivityAt(v string) *ListRepositoriesResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetName(v string) *ListRepositoriesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetNameWithNamespace(v string) *ListRepositoriesResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetNamespaceId(v int64) *ListRepositoriesResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetPath(v string) *ListRepositoriesResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetPathWithNamespace(v string) *ListRepositoriesResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetStar(v bool) *ListRepositoriesResponseBodyResult {
	s.Star = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetStarCount(v int64) *ListRepositoriesResponseBodyResult {
	s.StarCount = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetUpdatedAt(v string) *ListRepositoriesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetVisibilityLevel(v string) *ListRepositoriesResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetWebUrl(v string) *ListRepositoriesResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
