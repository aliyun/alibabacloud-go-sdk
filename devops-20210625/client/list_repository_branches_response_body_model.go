// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryBranchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListRepositoryBranchesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRepositoryBranchesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoryBranchesResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoryBranchesResponseBodyResult) *ListRepositoryBranchesResponseBody
	GetResult() []*ListRepositoryBranchesResponseBodyResult
	SetSuccess(v bool) *ListRepositoryBranchesResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListRepositoryBranchesResponseBody
	GetTotal() *int64
}

type ListRepositoryBranchesResponseBody struct {
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
	// 1F4F342D-493A-5B2C-B133-BA78B30FF834
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRepositoryBranchesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoryBranchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRepositoryBranchesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoryBranchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryBranchesResponseBody) GetResult() []*ListRepositoryBranchesResponseBodyResult {
	return s.Result
}

func (s *ListRepositoryBranchesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoryBranchesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListRepositoryBranchesResponseBody) SetErrorCode(v string) *ListRepositoryBranchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetErrorMessage(v string) *ListRepositoryBranchesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetRequestId(v string) *ListRepositoryBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetResult(v []*ListRepositoryBranchesResponseBodyResult) *ListRepositoryBranchesResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetSuccess(v bool) *ListRepositoryBranchesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetTotal(v int64) *ListRepositoryBranchesResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) Validate() error {
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

type ListRepositoryBranchesResponseBodyResult struct {
	Commit *ListRepositoryBranchesResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
	// example:
	//
	// testBranch
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Protected *string `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s ListRepositoryBranchesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryBranchesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBodyResult) GetCommit() *ListRepositoryBranchesResponseBodyResultCommit {
	return s.Commit
}

func (s *ListRepositoryBranchesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListRepositoryBranchesResponseBodyResult) GetProtected() *string {
	return s.Protected
}

func (s *ListRepositoryBranchesResponseBodyResult) SetCommit(v *ListRepositoryBranchesResponseBodyResultCommit) *ListRepositoryBranchesResponseBodyResult {
	s.Commit = v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) SetName(v string) *ListRepositoryBranchesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) SetProtected(v string) *ListRepositoryBranchesResponseBodyResult {
	s.Protected = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) Validate() error {
	if s.Commit != nil {
		if err := s.Commit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRepositoryBranchesResponseBodyResultCommit struct {
	// example:
	//
	// username@example.com
	AuthorEmail *string `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	// example:
	//
	// test-codeup
	AuthorName *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	// example:
	//
	// 2022-03-18 10:00:00
	AuthoredDate *string `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	// example:
	//
	// 2022-03-18 11:00:00
	CommittedDate *string `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	// example:
	//
	// username@example.com
	CommitterEmail *string `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	// example:
	//
	// commit-codeup
	CommitterName *string `json:"committerName,omitempty" xml:"committerName,omitempty"`
	// example:
	//
	// 2022-03-18 10:00:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// de02b625ba8488f92eb204bcb3773a40c1b4ddac
	Id        *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// e0297d8f
	ShortId *string `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListRepositoryBranchesResponseBodyResultCommit) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryBranchesResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetAuthorName() *string {
	return s.AuthorName
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetCommitterName() *string {
	return s.CommitterName
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetId() *string {
	return s.Id
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetMessage() *string {
	return s.Message
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetParentIds() []*string {
	return s.ParentIds
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetShortId() *string {
	return s.ShortId
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) GetTitle() *string {
	return s.Title
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetAuthorEmail(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetAuthorName(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetAuthoredDate(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetCommittedDate(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetCommitterEmail(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetCommitterName(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetCreatedAt(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetId(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetMessage(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetParentIds(v []*string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetShortId(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) SetTitle(v string) *ListRepositoryBranchesResponseBodyResultCommit {
	s.Title = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommit) Validate() error {
	return dara.Validate(s)
}
