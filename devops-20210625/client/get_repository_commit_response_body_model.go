// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryCommitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetRepositoryCommitResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetRepositoryCommitResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetRepositoryCommitResponseBody
	GetRequestId() *string
	SetResult(v *GetRepositoryCommitResponseBodyResult) *GetRepositoryCommitResponseBody
	GetResult() *GetRepositoryCommitResponseBodyResult
	SetSuccess(v bool) *GetRepositoryCommitResponseBody
	GetSuccess() *bool
}

type GetRepositoryCommitResponseBody struct {
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
	// A7586FEB-E48D-5579-983F-74981FBFF627
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetRepositoryCommitResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRepositoryCommitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryCommitResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRepositoryCommitResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRepositoryCommitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepositoryCommitResponseBody) GetResult() *GetRepositoryCommitResponseBodyResult {
	return s.Result
}

func (s *GetRepositoryCommitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRepositoryCommitResponseBody) SetErrorCode(v string) *GetRepositoryCommitResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryCommitResponseBody) SetErrorMessage(v string) *GetRepositoryCommitResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryCommitResponseBody) SetRequestId(v string) *GetRepositoryCommitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryCommitResponseBody) SetResult(v *GetRepositoryCommitResponseBodyResult) *GetRepositoryCommitResponseBody {
	s.Result = v
	return s
}

func (s *GetRepositoryCommitResponseBody) SetSuccess(v bool) *GetRepositoryCommitResponseBody {
	s.Success = &v
	return s
}

func (s *GetRepositoryCommitResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRepositoryCommitResponseBodyResult struct {
	Author *GetRepositoryCommitResponseBodyResultAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
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
	// 2022-03-18 15:00:00
	AuthoredDate *string `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	// example:
	//
	// 1
	CommentsCount *int64 `json:"commentsCount,omitempty" xml:"commentsCount,omitempty"`
	// example:
	//
	// 2022-03-18 15:00:02
	CommittedDate *string                                         `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	Committer     *GetRepositoryCommitResponseBodyResultCommitter `json:"committer,omitempty" xml:"committer,omitempty" type:"Struct"`
	// example:
	//
	// username@example.com
	CommitterEmail *string `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	// example:
	//
	// committer-codeup
	CommitterName *string `json:"committerName,omitempty" xml:"committerName,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// ff4fb5ac6d1f44f452654336d2dba468ae6c8d04
	Id        *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// ff4fb5ac
	ShortId   *string                                         `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Signature *GetRepositoryCommitResponseBodyResultSignature `json:"signature,omitempty" xml:"signature,omitempty" type:"Struct"`
	Title     *string                                         `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetRepositoryCommitResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryCommitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponseBodyResult) GetAuthor() *GetRepositoryCommitResponseBodyResultAuthor {
	return s.Author
}

func (s *GetRepositoryCommitResponseBodyResult) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *GetRepositoryCommitResponseBodyResult) GetAuthorName() *string {
	return s.AuthorName
}

func (s *GetRepositoryCommitResponseBodyResult) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *GetRepositoryCommitResponseBodyResult) GetCommentsCount() *int64 {
	return s.CommentsCount
}

func (s *GetRepositoryCommitResponseBodyResult) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *GetRepositoryCommitResponseBodyResult) GetCommitter() *GetRepositoryCommitResponseBodyResultCommitter {
	return s.Committer
}

func (s *GetRepositoryCommitResponseBodyResult) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *GetRepositoryCommitResponseBodyResult) GetCommitterName() *string {
	return s.CommitterName
}

func (s *GetRepositoryCommitResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetRepositoryCommitResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetRepositoryCommitResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *GetRepositoryCommitResponseBodyResult) GetParentIds() []*string {
	return s.ParentIds
}

func (s *GetRepositoryCommitResponseBodyResult) GetShortId() *string {
	return s.ShortId
}

func (s *GetRepositoryCommitResponseBodyResult) GetSignature() *GetRepositoryCommitResponseBodyResultSignature {
	return s.Signature
}

func (s *GetRepositoryCommitResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetRepositoryCommitResponseBodyResult) SetAuthor(v *GetRepositoryCommitResponseBodyResultAuthor) *GetRepositoryCommitResponseBodyResult {
	s.Author = v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetAuthorEmail(v string) *GetRepositoryCommitResponseBodyResult {
	s.AuthorEmail = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetAuthorName(v string) *GetRepositoryCommitResponseBodyResult {
	s.AuthorName = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetAuthoredDate(v string) *GetRepositoryCommitResponseBodyResult {
	s.AuthoredDate = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCommentsCount(v int64) *GetRepositoryCommitResponseBodyResult {
	s.CommentsCount = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCommittedDate(v string) *GetRepositoryCommitResponseBodyResult {
	s.CommittedDate = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCommitter(v *GetRepositoryCommitResponseBodyResultCommitter) *GetRepositoryCommitResponseBodyResult {
	s.Committer = v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCommitterEmail(v string) *GetRepositoryCommitResponseBodyResult {
	s.CommitterEmail = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCommitterName(v string) *GetRepositoryCommitResponseBodyResult {
	s.CommitterName = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCreatedAt(v string) *GetRepositoryCommitResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetId(v string) *GetRepositoryCommitResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetMessage(v string) *GetRepositoryCommitResponseBodyResult {
	s.Message = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetParentIds(v []*string) *GetRepositoryCommitResponseBodyResult {
	s.ParentIds = v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetShortId(v string) *GetRepositoryCommitResponseBodyResult {
	s.ShortId = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetSignature(v *GetRepositoryCommitResponseBodyResultSignature) *GetRepositoryCommitResponseBodyResult {
	s.Signature = v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetTitle(v string) *GetRepositoryCommitResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) Validate() error {
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

type GetRepositoryCommitResponseBodyResultAuthor struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c4ef67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 19927
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// test-nickname
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// example:
	//
	// ""
	WebsiteUrl *string `json:"websiteUrl,omitempty" xml:"websiteUrl,omitempty"`
}

func (s GetRepositoryCommitResponseBodyResultAuthor) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryCommitResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) GetEmail() *string {
	return s.Email
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) GetId() *int64 {
	return s.Id
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) GetName() *string {
	return s.Name
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) GetState() *string {
	return s.State
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) GetUsername() *string {
	return s.Username
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) GetWebsiteUrl() *string {
	return s.WebsiteUrl
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) SetAvatarUrl(v string) *GetRepositoryCommitResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) SetEmail(v string) *GetRepositoryCommitResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) SetId(v int64) *GetRepositoryCommitResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) SetName(v string) *GetRepositoryCommitResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) SetState(v string) *GetRepositoryCommitResponseBodyResultAuthor {
	s.State = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) SetUsername(v string) *GetRepositoryCommitResponseBodyResultAuthor {
	s.Username = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) SetWebsiteUrl(v string) *GetRepositoryCommitResponseBodyResultAuthor {
	s.WebsiteUrl = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultAuthor) Validate() error {
	return dara.Validate(s)
}

type GetRepositoryCommitResponseBodyResultCommitter struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c4ef67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 24661
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// committer-codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// nickname
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// example:
	//
	// ""
	WebsiteUrl *string `json:"websiteUrl,omitempty" xml:"websiteUrl,omitempty"`
}

func (s GetRepositoryCommitResponseBodyResultCommitter) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryCommitResponseBodyResultCommitter) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) GetEmail() *string {
	return s.Email
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) GetId() *int64 {
	return s.Id
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) GetName() *string {
	return s.Name
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) GetState() *string {
	return s.State
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) GetUsername() *string {
	return s.Username
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) GetWebsiteUrl() *string {
	return s.WebsiteUrl
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) SetAvatarUrl(v string) *GetRepositoryCommitResponseBodyResultCommitter {
	s.AvatarUrl = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) SetEmail(v string) *GetRepositoryCommitResponseBodyResultCommitter {
	s.Email = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) SetId(v int64) *GetRepositoryCommitResponseBodyResultCommitter {
	s.Id = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) SetName(v string) *GetRepositoryCommitResponseBodyResultCommitter {
	s.Name = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) SetState(v string) *GetRepositoryCommitResponseBodyResultCommitter {
	s.State = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) SetUsername(v string) *GetRepositoryCommitResponseBodyResultCommitter {
	s.Username = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) SetWebsiteUrl(v string) *GetRepositoryCommitResponseBodyResultCommitter {
	s.WebsiteUrl = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultCommitter) Validate() error {
	return dara.Validate(s)
}

type GetRepositoryCommitResponseBodyResultSignature struct {
	// example:
	//
	// 34d2c47c7ce46a5c4639c5ffe208
	GpgKeyId *string `json:"gpgKeyId,omitempty" xml:"gpgKeyId,omitempty"`
	// example:
	//
	// verified
	VerificationStatus *string `json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
}

func (s GetRepositoryCommitResponseBodyResultSignature) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryCommitResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponseBodyResultSignature) GetGpgKeyId() *string {
	return s.GpgKeyId
}

func (s *GetRepositoryCommitResponseBodyResultSignature) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *GetRepositoryCommitResponseBodyResultSignature) SetGpgKeyId(v string) *GetRepositoryCommitResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultSignature) SetVerificationStatus(v string) *GetRepositoryCommitResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultSignature) Validate() error {
	return dara.Validate(s)
}
