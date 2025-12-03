// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupRepositoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListGroupRepositoriesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListGroupRepositoriesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListGroupRepositoriesResponseBody
	GetRequestId() *string
	SetResult(v []*ListGroupRepositoriesResponseBodyResult) *ListGroupRepositoriesResponseBody
	GetResult() []*ListGroupRepositoriesResponseBodyResult
	SetSuccess(v bool) *ListGroupRepositoriesResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListGroupRepositoriesResponseBody
	GetTotal() *int64
}

type ListGroupRepositoriesResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 9293CBDA-B5BC-5AD6-A8F4-C7873AC7A3DF
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListGroupRepositoriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListGroupRepositoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListGroupRepositoriesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListGroupRepositoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupRepositoriesResponseBody) GetResult() []*ListGroupRepositoriesResponseBodyResult {
	return s.Result
}

func (s *ListGroupRepositoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGroupRepositoriesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListGroupRepositoriesResponseBody) SetErrorCode(v string) *ListGroupRepositoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetErrorMessage(v string) *ListGroupRepositoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetRequestId(v string) *ListGroupRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetResult(v []*ListGroupRepositoriesResponseBodyResult) *ListGroupRepositoriesResponseBody {
	s.Result = v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetSuccess(v bool) *ListGroupRepositoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetTotal(v int64) *ListGroupRepositoriesResponseBody {
	s.Total = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) Validate() error {
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

type ListGroupRepositoriesResponseBodyResult struct {
	// example:
	//
	// false
	Archived *bool `json:"archived,omitempty" xml:"archived,omitempty"`
	// example:
	//
	// 10
	CommitCount *int64 `json:"commitCount,omitempty" xml:"commitCount,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 19238
	CreatorId   *int64  `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://xxx/test/test
	HttpUrl *string `json:"httpUrl,omitempty" xml:"httpUrl,omitempty"`
	// example:
	//
	// 89616
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// http://xxx/x
	ImportUrl *string `json:"importUrl,omitempty" xml:"importUrl,omitempty"`
	// example:
	//
	// true
	IsStared *bool `json:"isStared,omitempty" xml:"isStared,omitempty"`
	// example:
	//
	// true
	IssuesEnabled *bool `json:"issuesEnabled,omitempty" xml:"issuesEnabled,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	LastActivityAt *string `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	// example:
	//
	// true
	MergeRequestsEnabled *bool `json:"mergeRequestsEnabled,omitempty" xml:"mergeRequestsEnabled,omitempty"`
	// example:
	//
	// test-group-repo
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// 100003
	NamespaceId *bool `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// example:
	//
	// test-group-repo
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org/test-group-repo
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// true
	PrivateFlag *bool `json:"privateFlag,omitempty" xml:"privateFlag,omitempty"`
	// example:
	//
	// false
	SnippetsEnabled *bool `json:"snippetsEnabled,omitempty" xml:"snippetsEnabled,omitempty"`
	// example:
	//
	// git@xxx:xxx/test/test.git
	SshUrl *string `json:"sshUrl,omitempty" xml:"sshUrl,omitempty"`
	// example:
	//
	// 0
	StarCount *int32 `json:"starCount,omitempty" xml:"starCount,omitempty"`
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
	// example:
	//
	// true
	WikiEnabled *bool `json:"wikiEnabled,omitempty" xml:"wikiEnabled,omitempty"`
}

func (s ListGroupRepositoriesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListGroupRepositoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesResponseBodyResult) GetArchived() *bool {
	return s.Archived
}

func (s *ListGroupRepositoriesResponseBodyResult) GetCommitCount() *int64 {
	return s.CommitCount
}

func (s *ListGroupRepositoriesResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListGroupRepositoriesResponseBodyResult) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ListGroupRepositoriesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListGroupRepositoriesResponseBodyResult) GetHttpUrl() *string {
	return s.HttpUrl
}

func (s *ListGroupRepositoriesResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListGroupRepositoriesResponseBodyResult) GetImportUrl() *string {
	return s.ImportUrl
}

func (s *ListGroupRepositoriesResponseBodyResult) GetIsStared() *bool {
	return s.IsStared
}

func (s *ListGroupRepositoriesResponseBodyResult) GetIssuesEnabled() *bool {
	return s.IssuesEnabled
}

func (s *ListGroupRepositoriesResponseBodyResult) GetLastActivityAt() *string {
	return s.LastActivityAt
}

func (s *ListGroupRepositoriesResponseBodyResult) GetMergeRequestsEnabled() *bool {
	return s.MergeRequestsEnabled
}

func (s *ListGroupRepositoriesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListGroupRepositoriesResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *ListGroupRepositoriesResponseBodyResult) GetNamespaceId() *bool {
	return s.NamespaceId
}

func (s *ListGroupRepositoriesResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *ListGroupRepositoriesResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *ListGroupRepositoriesResponseBodyResult) GetPrivateFlag() *bool {
	return s.PrivateFlag
}

func (s *ListGroupRepositoriesResponseBodyResult) GetSnippetsEnabled() *bool {
	return s.SnippetsEnabled
}

func (s *ListGroupRepositoriesResponseBodyResult) GetSshUrl() *string {
	return s.SshUrl
}

func (s *ListGroupRepositoriesResponseBodyResult) GetStarCount() *int32 {
	return s.StarCount
}

func (s *ListGroupRepositoriesResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListGroupRepositoriesResponseBodyResult) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *ListGroupRepositoriesResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *ListGroupRepositoriesResponseBodyResult) GetWikiEnabled() *bool {
	return s.WikiEnabled
}

func (s *ListGroupRepositoriesResponseBodyResult) SetArchived(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.Archived = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetCommitCount(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.CommitCount = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetCreatedAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetCreatorId(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetDescription(v string) *ListGroupRepositoriesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetHttpUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.HttpUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetId(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetImportUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.ImportUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetIsStared(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.IsStared = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetIssuesEnabled(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.IssuesEnabled = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetLastActivityAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetMergeRequestsEnabled(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.MergeRequestsEnabled = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetName(v string) *ListGroupRepositoriesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetNameWithNamespace(v string) *ListGroupRepositoriesResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetNamespaceId(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetPath(v string) *ListGroupRepositoriesResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetPathWithNamespace(v string) *ListGroupRepositoriesResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetPrivateFlag(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.PrivateFlag = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetSnippetsEnabled(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.SnippetsEnabled = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetSshUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.SshUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetStarCount(v int32) *ListGroupRepositoriesResponseBodyResult {
	s.StarCount = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetUpdatedAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetVisibilityLevel(v int32) *ListGroupRepositoriesResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetWebUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetWikiEnabled(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.WikiEnabled = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
