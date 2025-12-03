// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBranchInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetBranchInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetBranchInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetBranchInfoResponseBody
	GetRequestId() *string
	SetResult(v *GetBranchInfoResponseBodyResult) *GetBranchInfoResponseBody
	GetResult() *GetBranchInfoResponseBodyResult
	SetSuccess(v bool) *GetBranchInfoResponseBody
	GetSuccess() *bool
}

type GetBranchInfoResponseBody struct {
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
	// 6177543A-8D54-5736-A93B-E0195A1512CB
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetBranchInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetBranchInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBranchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetBranchInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBranchInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBranchInfoResponseBody) GetResult() *GetBranchInfoResponseBodyResult {
	return s.Result
}

func (s *GetBranchInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBranchInfoResponseBody) SetErrorCode(v string) *GetBranchInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetErrorMessage(v string) *GetBranchInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetRequestId(v string) *GetBranchInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetResult(v *GetBranchInfoResponseBodyResult) *GetBranchInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetBranchInfoResponseBody) SetSuccess(v bool) *GetBranchInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetBranchInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBranchInfoResponseBodyResult struct {
	Commit *GetBranchInfoResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
	// example:
	//
	// master
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Protected *string `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s GetBranchInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetBranchInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResult) GetCommit() *GetBranchInfoResponseBodyResultCommit {
	return s.Commit
}

func (s *GetBranchInfoResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetBranchInfoResponseBodyResult) GetProtected() *string {
	return s.Protected
}

func (s *GetBranchInfoResponseBodyResult) SetCommit(v *GetBranchInfoResponseBodyResultCommit) *GetBranchInfoResponseBodyResult {
	s.Commit = v
	return s
}

func (s *GetBranchInfoResponseBodyResult) SetName(v string) *GetBranchInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetBranchInfoResponseBodyResult) SetProtected(v string) *GetBranchInfoResponseBodyResult {
	s.Protected = &v
	return s
}

func (s *GetBranchInfoResponseBodyResult) Validate() error {
	if s.Commit != nil {
		if err := s.Commit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBranchInfoResponseBodyResultCommit struct {
	Author *GetBranchInfoResponseBodyResultCommitAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	// example:
	//
	// username@example.com
	AuthorEmail *string `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	AuthorName  *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	// example:
	//
	// 2022-03-18 08:00:00
	AuthoredDate *string `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	// example:
	//
	// 0
	CommentsCount *int64 `json:"commentsCount,omitempty" xml:"commentsCount,omitempty"`
	// example:
	//
	// 2022-03-18 09:00:00
	CommittedDate *string                                         `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	Committer     *GetBranchInfoResponseBodyResultCommitCommitter `json:"committer,omitempty" xml:"committer,omitempty" type:"Struct"`
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
	// e0297d8fb0393c833a8531e7cc8832739e3cba6d
	Id        *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// e0297d8f
	ShortId   *string                                         `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Signature *GetBranchInfoResponseBodyResultCommitSignature `json:"signature,omitempty" xml:"signature,omitempty" type:"Struct"`
	Title     *string                                         `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommit) String() string {
	return dara.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommit) GetAuthor() *GetBranchInfoResponseBodyResultCommitAuthor {
	return s.Author
}

func (s *GetBranchInfoResponseBodyResultCommit) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *GetBranchInfoResponseBodyResultCommit) GetAuthorName() *string {
	return s.AuthorName
}

func (s *GetBranchInfoResponseBodyResultCommit) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *GetBranchInfoResponseBodyResultCommit) GetCommentsCount() *int64 {
	return s.CommentsCount
}

func (s *GetBranchInfoResponseBodyResultCommit) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *GetBranchInfoResponseBodyResultCommit) GetCommitter() *GetBranchInfoResponseBodyResultCommitCommitter {
	return s.Committer
}

func (s *GetBranchInfoResponseBodyResultCommit) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *GetBranchInfoResponseBodyResultCommit) GetCommitterName() *string {
	return s.CommitterName
}

func (s *GetBranchInfoResponseBodyResultCommit) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetBranchInfoResponseBodyResultCommit) GetId() *string {
	return s.Id
}

func (s *GetBranchInfoResponseBodyResultCommit) GetMessage() *string {
	return s.Message
}

func (s *GetBranchInfoResponseBodyResultCommit) GetParentIds() []*string {
	return s.ParentIds
}

func (s *GetBranchInfoResponseBodyResultCommit) GetShortId() *string {
	return s.ShortId
}

func (s *GetBranchInfoResponseBodyResultCommit) GetSignature() *GetBranchInfoResponseBodyResultCommitSignature {
	return s.Signature
}

func (s *GetBranchInfoResponseBodyResultCommit) GetTitle() *string {
	return s.Title
}

func (s *GetBranchInfoResponseBodyResultCommit) SetAuthor(v *GetBranchInfoResponseBodyResultCommitAuthor) *GetBranchInfoResponseBodyResultCommit {
	s.Author = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetAuthorEmail(v string) *GetBranchInfoResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetAuthorName(v string) *GetBranchInfoResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetAuthoredDate(v string) *GetBranchInfoResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommentsCount(v int64) *GetBranchInfoResponseBodyResultCommit {
	s.CommentsCount = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommittedDate(v string) *GetBranchInfoResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommitter(v *GetBranchInfoResponseBodyResultCommitCommitter) *GetBranchInfoResponseBodyResultCommit {
	s.Committer = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommitterEmail(v string) *GetBranchInfoResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCommitterName(v string) *GetBranchInfoResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetCreatedAt(v string) *GetBranchInfoResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetId(v string) *GetBranchInfoResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetMessage(v string) *GetBranchInfoResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetParentIds(v []*string) *GetBranchInfoResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetShortId(v string) *GetBranchInfoResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetSignature(v *GetBranchInfoResponseBodyResultCommitSignature) *GetBranchInfoResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) SetTitle(v string) *GetBranchInfoResponseBodyResultCommit {
	s.Title = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommit) Validate() error {
	if s.Author != nil {
		if err := s.Author.Validate(); err != nil {
			return err
		}
	}
	if s.Committer != nil {
		if err := s.Committer.Validate(); err != nil {
			return err
		}
	}
	if s.Signature != nil {
		if err := s.Signature.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBranchInfoResponseBodyResultCommitAuthor struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 28056
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// testtest
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommitAuthor) String() string {
	return dara.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommitAuthor) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) GetEmail() *string {
	return s.Email
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) GetId() *int64 {
	return s.Id
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) GetName() *string {
	return s.Name
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) GetState() *string {
	return s.State
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) GetUsername() *string {
	return s.Username
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetAvatarUrl(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetEmail(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.Email = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetId(v int64) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.Id = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetName(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.Name = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetState(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.State = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) SetUsername(v string) *GetBranchInfoResponseBodyResultCommitAuthor {
	s.Username = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitAuthor) Validate() error {
	return dara.Validate(s)
}

type GetBranchInfoResponseBodyResultCommitCommitter struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 5035
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup-commit
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// commitcommit
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommitCommitter) String() string {
	return dara.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommitCommitter) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) GetEmail() *string {
	return s.Email
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) GetId() *int64 {
	return s.Id
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) GetName() *string {
	return s.Name
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) GetState() *string {
	return s.State
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) GetUsername() *string {
	return s.Username
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetAvatarUrl(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.AvatarUrl = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetEmail(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.Email = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetId(v int64) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.Id = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetName(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.Name = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetState(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.State = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) SetUsername(v string) *GetBranchInfoResponseBodyResultCommitCommitter {
	s.Username = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitCommitter) Validate() error {
	return dara.Validate(s)
}

type GetBranchInfoResponseBodyResultCommitSignature struct {
	// example:
	//
	// ""
	GpgKeyId *string `json:"gpgKeyId,omitempty" xml:"gpgKeyId,omitempty"`
	// example:
	//
	// verified
	VerificationStatus *string `json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommitSignature) String() string {
	return dara.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommitSignature) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommitSignature) GetGpgKeyId() *string {
	return s.GpgKeyId
}

func (s *GetBranchInfoResponseBodyResultCommitSignature) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *GetBranchInfoResponseBodyResultCommitSignature) SetGpgKeyId(v string) *GetBranchInfoResponseBodyResultCommitSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitSignature) SetVerificationStatus(v string) *GetBranchInfoResponseBodyResultCommitSignature {
	s.VerificationStatus = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitSignature) Validate() error {
	return dara.Validate(s)
}
