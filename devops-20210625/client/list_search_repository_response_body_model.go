// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSearchRepositoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSearchRepositoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSearchRepositoryResponseBody
	GetRequestId() *string
	SetResult(v []*ListSearchRepositoryResponseBodyResult) *ListSearchRepositoryResponseBody
	GetResult() []*ListSearchRepositoryResponseBodyResult
	SetSuccess(v bool) *ListSearchRepositoryResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListSearchRepositoryResponseBody
	GetTotal() *int64
}

type ListSearchRepositoryResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListSearchRepositoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 30
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSearchRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchRepositoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSearchRepositoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSearchRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchRepositoryResponseBody) GetResult() []*ListSearchRepositoryResponseBodyResult {
	return s.Result
}

func (s *ListSearchRepositoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSearchRepositoryResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListSearchRepositoryResponseBody) SetErrorCode(v string) *ListSearchRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSearchRepositoryResponseBody) SetErrorMessage(v string) *ListSearchRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSearchRepositoryResponseBody) SetRequestId(v string) *ListSearchRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchRepositoryResponseBody) SetResult(v []*ListSearchRepositoryResponseBodyResult) *ListSearchRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *ListSearchRepositoryResponseBody) SetSuccess(v bool) *ListSearchRepositoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListSearchRepositoryResponseBody) SetTotal(v int64) *ListSearchRepositoryResponseBody {
	s.Total = &v
	return s
}

func (s *ListSearchRepositoryResponseBody) Validate() error {
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

type ListSearchRepositoryResponseBodyResult struct {
	// example:
	//
	// 60d54f3daccf2bbd6659f3ad/gitlabhq/master/config/environments/test.rb
	DocId            *string                                                 `json:"docId,omitempty" xml:"docId,omitempty"`
	HighlightTextMap *ListSearchRepositoryResponseBodyResultHighlightTextMap `json:"highlightTextMap,omitempty" xml:"highlightTextMap,omitempty" type:"Struct"`
	Source           *ListSearchRepositoryResponseBodyResultSource           `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s ListSearchRepositoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSearchRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSearchRepositoryResponseBodyResult) GetDocId() *string {
	return s.DocId
}

func (s *ListSearchRepositoryResponseBodyResult) GetHighlightTextMap() *ListSearchRepositoryResponseBodyResultHighlightTextMap {
	return s.HighlightTextMap
}

func (s *ListSearchRepositoryResponseBodyResult) GetSource() *ListSearchRepositoryResponseBodyResultSource {
	return s.Source
}

func (s *ListSearchRepositoryResponseBodyResult) SetDocId(v string) *ListSearchRepositoryResponseBodyResult {
	s.DocId = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResult) SetHighlightTextMap(v *ListSearchRepositoryResponseBodyResultHighlightTextMap) *ListSearchRepositoryResponseBodyResult {
	s.HighlightTextMap = v
	return s
}

func (s *ListSearchRepositoryResponseBodyResult) SetSource(v *ListSearchRepositoryResponseBodyResultSource) *ListSearchRepositoryResponseBodyResult {
	s.Source = v
	return s
}

func (s *ListSearchRepositoryResponseBodyResult) Validate() error {
	if s.HighlightTextMap != nil {
		if err := s.HighlightTextMap.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSearchRepositoryResponseBodyResultHighlightTextMap struct {
	// example:
	//
	// xxx
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// xx
	ReadMe *string `json:"readMe,omitempty" xml:"readMe,omitempty"`
	// example:
	//
	// codeup/test-repo
	RepoNameWithNamespace *string `json:"repoNameWithNamespace,omitempty" xml:"repoNameWithNamespace,omitempty"`
	// example:
	//
	// codeup/test-repo
	RepoPath *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
}

func (s ListSearchRepositoryResponseBodyResultHighlightTextMap) String() string {
	return dara.Prettify(s)
}

func (s ListSearchRepositoryResponseBodyResultHighlightTextMap) GoString() string {
	return s.String()
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) GetDescription() *string {
	return s.Description
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) GetReadMe() *string {
	return s.ReadMe
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) GetRepoNameWithNamespace() *string {
	return s.RepoNameWithNamespace
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) GetRepoPath() *string {
	return s.RepoPath
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) SetCreatorUserId(v string) *ListSearchRepositoryResponseBodyResultHighlightTextMap {
	s.CreatorUserId = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) SetDescription(v string) *ListSearchRepositoryResponseBodyResultHighlightTextMap {
	s.Description = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) SetOrganizationId(v string) *ListSearchRepositoryResponseBodyResultHighlightTextMap {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) SetReadMe(v string) *ListSearchRepositoryResponseBodyResultHighlightTextMap {
	s.ReadMe = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) SetRepoNameWithNamespace(v string) *ListSearchRepositoryResponseBodyResultHighlightTextMap {
	s.RepoNameWithNamespace = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) SetRepoPath(v string) *ListSearchRepositoryResponseBodyResultHighlightTextMap {
	s.RepoPath = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultHighlightTextMap) Validate() error {
	return dara.Validate(s)
}

type ListSearchRepositoryResponseBodyResultSource struct {
	// example:
	//
	// 2022-10-10 10:10:10
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2022-11-11 11:11:11
	LastActivityTime *string `json:"lastActivityTime,omitempty" xml:"lastActivityTime,omitempty"`
	// example:
	//
	// 62a94a8611fc0f0c9e2a7bc1
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// xxx
	ReadMe *string `json:"readMe,omitempty" xml:"readMe,omitempty"`
	// example:
	//
	// test-repo
	RepoName *string `json:"repoName,omitempty" xml:"repoName,omitempty"`
	// example:
	//
	// codeup/test-repo
	RepoPath *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s ListSearchRepositoryResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s ListSearchRepositoryResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *ListSearchRepositoryResponseBodyResultSource) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSearchRepositoryResponseBodyResultSource) GetDescription() *string {
	return s.Description
}

func (s *ListSearchRepositoryResponseBodyResultSource) GetLastActivityTime() *string {
	return s.LastActivityTime
}

func (s *ListSearchRepositoryResponseBodyResultSource) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchRepositoryResponseBodyResultSource) GetReadMe() *string {
	return s.ReadMe
}

func (s *ListSearchRepositoryResponseBodyResultSource) GetRepoName() *string {
	return s.RepoName
}

func (s *ListSearchRepositoryResponseBodyResultSource) GetRepoPath() *string {
	return s.RepoPath
}

func (s *ListSearchRepositoryResponseBodyResultSource) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *ListSearchRepositoryResponseBodyResultSource) SetCreateTime(v string) *ListSearchRepositoryResponseBodyResultSource {
	s.CreateTime = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultSource) SetDescription(v string) *ListSearchRepositoryResponseBodyResultSource {
	s.Description = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultSource) SetLastActivityTime(v string) *ListSearchRepositoryResponseBodyResultSource {
	s.LastActivityTime = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultSource) SetOrganizationId(v string) *ListSearchRepositoryResponseBodyResultSource {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultSource) SetReadMe(v string) *ListSearchRepositoryResponseBodyResultSource {
	s.ReadMe = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultSource) SetRepoName(v string) *ListSearchRepositoryResponseBodyResultSource {
	s.RepoName = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultSource) SetRepoPath(v string) *ListSearchRepositoryResponseBodyResultSource {
	s.RepoPath = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultSource) SetVisibilityLevel(v int32) *ListSearchRepositoryResponseBodyResultSource {
	s.VisibilityLevel = &v
	return s
}

func (s *ListSearchRepositoryResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}
