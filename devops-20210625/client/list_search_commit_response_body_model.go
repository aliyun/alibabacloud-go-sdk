// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchCommitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSearchCommitResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSearchCommitResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSearchCommitResponseBody
	GetRequestId() *string
	SetResult(v []*ListSearchCommitResponseBodyResult) *ListSearchCommitResponseBody
	GetResult() []*ListSearchCommitResponseBodyResult
	SetSuccess(v bool) *ListSearchCommitResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListSearchCommitResponseBody
	GetTotal() *int64
}

type ListSearchCommitResponseBody struct {
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
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListSearchCommitResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSearchCommitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchCommitResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchCommitResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSearchCommitResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSearchCommitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchCommitResponseBody) GetResult() []*ListSearchCommitResponseBodyResult {
	return s.Result
}

func (s *ListSearchCommitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSearchCommitResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListSearchCommitResponseBody) SetErrorCode(v string) *ListSearchCommitResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSearchCommitResponseBody) SetErrorMessage(v string) *ListSearchCommitResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSearchCommitResponseBody) SetRequestId(v string) *ListSearchCommitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchCommitResponseBody) SetResult(v []*ListSearchCommitResponseBodyResult) *ListSearchCommitResponseBody {
	s.Result = v
	return s
}

func (s *ListSearchCommitResponseBody) SetSuccess(v bool) *ListSearchCommitResponseBody {
	s.Success = &v
	return s
}

func (s *ListSearchCommitResponseBody) SetTotal(v int64) *ListSearchCommitResponseBody {
	s.Total = &v
	return s
}

func (s *ListSearchCommitResponseBody) Validate() error {
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

type ListSearchCommitResponseBodyResult struct {
	// example:
	//
	// 60d54f3daccf2bbd6659f3ad/gitlabhq/master/config/environments/test.rb
	DocId            *string                                             `json:"docId,omitempty" xml:"docId,omitempty"`
	HighlightTextMap *ListSearchCommitResponseBodyResultHighlightTextMap `json:"highlightTextMap,omitempty" xml:"highlightTextMap,omitempty" type:"Struct"`
	Source           *ListSearchCommitResponseBodyResultSource           `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s ListSearchCommitResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSearchCommitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSearchCommitResponseBodyResult) GetDocId() *string {
	return s.DocId
}

func (s *ListSearchCommitResponseBodyResult) GetHighlightTextMap() *ListSearchCommitResponseBodyResultHighlightTextMap {
	return s.HighlightTextMap
}

func (s *ListSearchCommitResponseBodyResult) GetSource() *ListSearchCommitResponseBodyResultSource {
	return s.Source
}

func (s *ListSearchCommitResponseBodyResult) SetDocId(v string) *ListSearchCommitResponseBodyResult {
	s.DocId = &v
	return s
}

func (s *ListSearchCommitResponseBodyResult) SetHighlightTextMap(v *ListSearchCommitResponseBodyResultHighlightTextMap) *ListSearchCommitResponseBodyResult {
	s.HighlightTextMap = v
	return s
}

func (s *ListSearchCommitResponseBodyResult) SetSource(v *ListSearchCommitResponseBodyResultSource) *ListSearchCommitResponseBodyResult {
	s.Source = v
	return s
}

func (s *ListSearchCommitResponseBodyResult) Validate() error {
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

type ListSearchCommitResponseBodyResultHighlightTextMap struct {
	// example:
	//
	// df1b701cb0f3f7ca92320d49d31995821f2d045c
	CommitId      *string `json:"commitId,omitempty" xml:"commitId,omitempty"`
	CommitMessage *string `json:"commitMessage,omitempty" xml:"commitMessage,omitempty"`
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListSearchCommitResponseBodyResultHighlightTextMap) String() string {
	return dara.Prettify(s)
}

func (s ListSearchCommitResponseBodyResultHighlightTextMap) GoString() string {
	return s.String()
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) GetCommitId() *string {
	return s.CommitId
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) GetCommitMessage() *string {
	return s.CommitMessage
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) GetTitle() *string {
	return s.Title
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) SetCommitId(v string) *ListSearchCommitResponseBodyResultHighlightTextMap {
	s.CommitId = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) SetCommitMessage(v string) *ListSearchCommitResponseBodyResultHighlightTextMap {
	s.CommitMessage = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) SetOrganizationId(v string) *ListSearchCommitResponseBodyResultHighlightTextMap {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) SetTitle(v string) *ListSearchCommitResponseBodyResultHighlightTextMap {
	s.Title = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultHighlightTextMap) Validate() error {
	return dara.Validate(s)
}

type ListSearchCommitResponseBodyResultSource struct {
	Author *ListSearchCommitResponseBodyResultSourceAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	// example:
	//
	// 2022-11-11 11:11:11
	AuthorTime *string `json:"authorTime,omitempty" xml:"authorTime,omitempty"`
	// example:
	//
	// a748f5ecb17a93900d4808944bfcc96dc158ee2d
	CommitId      *string `json:"commitId,omitempty" xml:"commitId,omitempty"`
	CommitMessage *string `json:"commitMessage,omitempty" xml:"commitMessage,omitempty"`
	// example:
	//
	// 61133b011bd96aa110f1b500
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// codeup/test-repo
	RepoPath *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListSearchCommitResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s ListSearchCommitResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *ListSearchCommitResponseBodyResultSource) GetAuthor() *ListSearchCommitResponseBodyResultSourceAuthor {
	return s.Author
}

func (s *ListSearchCommitResponseBodyResultSource) GetAuthorTime() *string {
	return s.AuthorTime
}

func (s *ListSearchCommitResponseBodyResultSource) GetCommitId() *string {
	return s.CommitId
}

func (s *ListSearchCommitResponseBodyResultSource) GetCommitMessage() *string {
	return s.CommitMessage
}

func (s *ListSearchCommitResponseBodyResultSource) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchCommitResponseBodyResultSource) GetRepoPath() *string {
	return s.RepoPath
}

func (s *ListSearchCommitResponseBodyResultSource) GetTitle() *string {
	return s.Title
}

func (s *ListSearchCommitResponseBodyResultSource) SetAuthor(v *ListSearchCommitResponseBodyResultSourceAuthor) *ListSearchCommitResponseBodyResultSource {
	s.Author = v
	return s
}

func (s *ListSearchCommitResponseBodyResultSource) SetAuthorTime(v string) *ListSearchCommitResponseBodyResultSource {
	s.AuthorTime = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultSource) SetCommitId(v string) *ListSearchCommitResponseBodyResultSource {
	s.CommitId = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultSource) SetCommitMessage(v string) *ListSearchCommitResponseBodyResultSource {
	s.CommitMessage = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultSource) SetOrganizationId(v string) *ListSearchCommitResponseBodyResultSource {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultSource) SetRepoPath(v string) *ListSearchCommitResponseBodyResultSource {
	s.RepoPath = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultSource) SetTitle(v string) *ListSearchCommitResponseBodyResultSource {
	s.Title = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultSource) Validate() error {
	if s.Author != nil {
		if err := s.Author.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSearchCommitResponseBodyResultSourceAuthor struct {
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListSearchCommitResponseBodyResultSourceAuthor) String() string {
	return dara.Prettify(s)
}

func (s ListSearchCommitResponseBodyResultSourceAuthor) GoString() string {
	return s.String()
}

func (s *ListSearchCommitResponseBodyResultSourceAuthor) GetEmail() *string {
	return s.Email
}

func (s *ListSearchCommitResponseBodyResultSourceAuthor) GetName() *string {
	return s.Name
}

func (s *ListSearchCommitResponseBodyResultSourceAuthor) SetEmail(v string) *ListSearchCommitResponseBodyResultSourceAuthor {
	s.Email = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultSourceAuthor) SetName(v string) *ListSearchCommitResponseBodyResultSourceAuthor {
	s.Name = &v
	return s
}

func (s *ListSearchCommitResponseBodyResultSourceAuthor) Validate() error {
	return dara.Validate(s)
}
