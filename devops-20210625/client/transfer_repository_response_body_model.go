// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *TransferRepositoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *TransferRepositoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *TransferRepositoryResponseBody
	GetRequestId() *string
	SetResult(v *TransferRepositoryResponseBodyResult) *TransferRepositoryResponseBody
	GetResult() *TransferRepositoryResponseBodyResult
	SetSuccess(v bool) *TransferRepositoryResponseBody
	GetSuccess() *bool
}

type TransferRepositoryResponseBody struct {
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *TransferRepositoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TransferRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *TransferRepositoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TransferRepositoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *TransferRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferRepositoryResponseBody) GetResult() *TransferRepositoryResponseBodyResult {
	return s.Result
}

func (s *TransferRepositoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferRepositoryResponseBody) SetErrorCode(v string) *TransferRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransferRepositoryResponseBody) SetErrorMessage(v string) *TransferRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TransferRepositoryResponseBody) SetRequestId(v string) *TransferRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferRepositoryResponseBody) SetResult(v *TransferRepositoryResponseBodyResult) *TransferRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *TransferRepositoryResponseBody) SetSuccess(v bool) *TransferRepositoryResponseBody {
	s.Success = &v
	return s
}

func (s *TransferRepositoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TransferRepositoryResponseBodyResult struct {
	// example:
	//
	// 30
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// false
	Archived *bool `json:"archived,omitempty" xml:"archived,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 19238
	CreatorId *int64 `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// false
	DemoProject *bool   `json:"demoProject,omitempty" xml:"demoProject,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// example:
	//
	// 19285
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	LastActivityAt *string `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	// example:
	//
	// test-repo
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// 100003
	NamespaceId *int64 `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// example:
	//
	// test-repo
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org/test-create-codeup
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// 0
	StarCount *int64 `json:"starCount,omitempty" xml:"starCount,omitempty"`
	// example:
	//
	// true
	Starred *bool `json:"starred,omitempty" xml:"starred,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
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

func (s TransferRepositoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s TransferRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TransferRepositoryResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *TransferRepositoryResponseBodyResult) GetArchived() *bool {
	return s.Archived
}

func (s *TransferRepositoryResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *TransferRepositoryResponseBodyResult) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *TransferRepositoryResponseBodyResult) GetDemoProject() *bool {
	return s.DemoProject
}

func (s *TransferRepositoryResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *TransferRepositoryResponseBodyResult) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *TransferRepositoryResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *TransferRepositoryResponseBodyResult) GetLastActivityAt() *string {
	return s.LastActivityAt
}

func (s *TransferRepositoryResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *TransferRepositoryResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *TransferRepositoryResponseBodyResult) GetNamespaceId() *int64 {
	return s.NamespaceId
}

func (s *TransferRepositoryResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *TransferRepositoryResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *TransferRepositoryResponseBodyResult) GetStarCount() *int64 {
	return s.StarCount
}

func (s *TransferRepositoryResponseBodyResult) GetStarred() *bool {
	return s.Starred
}

func (s *TransferRepositoryResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *TransferRepositoryResponseBodyResult) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *TransferRepositoryResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *TransferRepositoryResponseBodyResult) SetAccessLevel(v int32) *TransferRepositoryResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetArchived(v bool) *TransferRepositoryResponseBodyResult {
	s.Archived = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetCreatedAt(v string) *TransferRepositoryResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetCreatorId(v int64) *TransferRepositoryResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetDemoProject(v bool) *TransferRepositoryResponseBodyResult {
	s.DemoProject = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetDescription(v string) *TransferRepositoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetEncrypted(v bool) *TransferRepositoryResponseBodyResult {
	s.Encrypted = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetId(v int64) *TransferRepositoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetLastActivityAt(v string) *TransferRepositoryResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetName(v string) *TransferRepositoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetNameWithNamespace(v string) *TransferRepositoryResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetNamespaceId(v int64) *TransferRepositoryResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetPath(v string) *TransferRepositoryResponseBodyResult {
	s.Path = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetPathWithNamespace(v string) *TransferRepositoryResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetStarCount(v int64) *TransferRepositoryResponseBodyResult {
	s.StarCount = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetStarred(v bool) *TransferRepositoryResponseBodyResult {
	s.Starred = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetUpdatedAt(v string) *TransferRepositoryResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetVisibilityLevel(v int32) *TransferRepositoryResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) SetWebUrl(v string) *TransferRepositoryResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *TransferRepositoryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
