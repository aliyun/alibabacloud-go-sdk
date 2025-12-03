// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateBranchResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateBranchResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateBranchResponseBody
	GetRequestId() *string
	SetResult(v *CreateBranchResponseBodyResult) *CreateBranchResponseBody
	GetResult() *CreateBranchResponseBodyResult
	SetSuccess(v bool) *CreateBranchResponseBody
	GetSuccess() *bool
}

type CreateBranchResponseBody struct {
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
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateBranchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBranchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateBranchResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBranchResponseBody) GetResult() *CreateBranchResponseBodyResult {
	return s.Result
}

func (s *CreateBranchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBranchResponseBody) SetErrorCode(v string) *CreateBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateBranchResponseBody) SetErrorMessage(v string) *CreateBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateBranchResponseBody) SetRequestId(v string) *CreateBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBranchResponseBody) SetResult(v *CreateBranchResponseBodyResult) *CreateBranchResponseBody {
	s.Result = v
	return s
}

func (s *CreateBranchResponseBody) SetSuccess(v bool) *CreateBranchResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBranchResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBranchResponseBodyResult struct {
	Commit *CreateBranchResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
	// example:
	//
	// createBranch
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Protected *bool `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s CreateBranchResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBodyResult) GetCommit() *CreateBranchResponseBodyResultCommit {
	return s.Commit
}

func (s *CreateBranchResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateBranchResponseBodyResult) GetProtected() *bool {
	return s.Protected
}

func (s *CreateBranchResponseBodyResult) SetCommit(v *CreateBranchResponseBodyResultCommit) *CreateBranchResponseBodyResult {
	s.Commit = v
	return s
}

func (s *CreateBranchResponseBodyResult) SetName(v string) *CreateBranchResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateBranchResponseBodyResult) SetProtected(v bool) *CreateBranchResponseBodyResult {
	s.Protected = &v
	return s
}

func (s *CreateBranchResponseBodyResult) Validate() error {
	if s.Commit != nil {
		if err := s.Commit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBranchResponseBodyResultCommit struct {
	// example:
	//
	// username@example.com
	AuthorEmail *string `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	AuthorName  *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	// example:
	//
	// 2022-03-18 09:00:00
	AuthoredDate *string `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	// example:
	//
	// 2022-03-18 10:00:00
	CommittedDate *string `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	// example:
	//
	// username@example.com
	CommitterEmail *string `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	CommitterName  *string `json:"committerName,omitempty" xml:"committerName,omitempty"`
	// example:
	//
	// 2022-03-18 10:00:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// ff4fb5ac6d1f44f452654336d2dba468ae6c8d04
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// create branch
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// ff4fb5ac
	ShortId *string `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateBranchResponseBodyResultCommit) String() string {
	return dara.Prettify(s)
}

func (s CreateBranchResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBodyResultCommit) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *CreateBranchResponseBodyResultCommit) GetAuthorName() *string {
	return s.AuthorName
}

func (s *CreateBranchResponseBodyResultCommit) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *CreateBranchResponseBodyResultCommit) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *CreateBranchResponseBodyResultCommit) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *CreateBranchResponseBodyResultCommit) GetCommitterName() *string {
	return s.CommitterName
}

func (s *CreateBranchResponseBodyResultCommit) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateBranchResponseBodyResultCommit) GetId() *string {
	return s.Id
}

func (s *CreateBranchResponseBodyResultCommit) GetMessage() *string {
	return s.Message
}

func (s *CreateBranchResponseBodyResultCommit) GetParentIds() []*string {
	return s.ParentIds
}

func (s *CreateBranchResponseBodyResultCommit) GetShortId() *string {
	return s.ShortId
}

func (s *CreateBranchResponseBodyResultCommit) GetTitle() *string {
	return s.Title
}

func (s *CreateBranchResponseBodyResultCommit) SetAuthorEmail(v string) *CreateBranchResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetAuthorName(v string) *CreateBranchResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetAuthoredDate(v string) *CreateBranchResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetCommittedDate(v string) *CreateBranchResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetCommitterEmail(v string) *CreateBranchResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetCommitterName(v string) *CreateBranchResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetCreatedAt(v string) *CreateBranchResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetId(v string) *CreateBranchResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetMessage(v string) *CreateBranchResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetParentIds(v []*string) *CreateBranchResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetShortId(v string) *CreateBranchResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) SetTitle(v string) *CreateBranchResponseBodyResultCommit {
	s.Title = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommit) Validate() error {
	return dara.Validate(s)
}
