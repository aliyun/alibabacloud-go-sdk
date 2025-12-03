// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryCommitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListRepositoryCommitsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRepositoryCommitsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoryCommitsResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoryCommitsResponseBodyResult) *ListRepositoryCommitsResponseBody
	GetResult() []*ListRepositoryCommitsResponseBodyResult
	SetSuccess(v bool) *ListRepositoryCommitsResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListRepositoryCommitsResponseBody
	GetTotal() *int64
}

type ListRepositoryCommitsResponseBody struct {
	// example:
	//
	// OpenApi.error
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F1138237-CF7F-56BF-95D4-9AA937CCE8E5
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRepositoryCommitsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 145
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoryCommitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRepositoryCommitsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoryCommitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryCommitsResponseBody) GetResult() []*ListRepositoryCommitsResponseBodyResult {
	return s.Result
}

func (s *ListRepositoryCommitsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoryCommitsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListRepositoryCommitsResponseBody) SetErrorCode(v string) *ListRepositoryCommitsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetErrorMessage(v string) *ListRepositoryCommitsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetRequestId(v string) *ListRepositoryCommitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetResult(v []*ListRepositoryCommitsResponseBodyResult) *ListRepositoryCommitsResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetSuccess(v bool) *ListRepositoryCommitsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetTotal(v int64) *ListRepositoryCommitsResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) Validate() error {
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

type ListRepositoryCommitsResponseBodyResult struct {
	Author *ListRepositoryCommitsResponseBodyResultAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
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
	// 2022-03-18 16:00:00
	CommittedDate *string                                           `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	Committer     *ListRepositoryCommitsResponseBodyResultCommitter `json:"committer,omitempty" xml:"committer,omitempty" type:"Struct"`
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
	// de02b625ba8488f92eb204bcb3773a40c1b4ddac
	Id        *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// de02b625
	ShortId   *string                                           `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Signature *ListRepositoryCommitsResponseBodyResultSignature `json:"signature,omitempty" xml:"signature,omitempty" type:"Struct"`
	Title     *string                                           `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListRepositoryCommitsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponseBodyResult) GetAuthor() *ListRepositoryCommitsResponseBodyResultAuthor {
	return s.Author
}

func (s *ListRepositoryCommitsResponseBodyResult) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *ListRepositoryCommitsResponseBodyResult) GetAuthorName() *string {
	return s.AuthorName
}

func (s *ListRepositoryCommitsResponseBodyResult) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *ListRepositoryCommitsResponseBodyResult) GetCommentsCount() *int64 {
	return s.CommentsCount
}

func (s *ListRepositoryCommitsResponseBodyResult) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *ListRepositoryCommitsResponseBodyResult) GetCommitter() *ListRepositoryCommitsResponseBodyResultCommitter {
	return s.Committer
}

func (s *ListRepositoryCommitsResponseBodyResult) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *ListRepositoryCommitsResponseBodyResult) GetCommitterName() *string {
	return s.CommitterName
}

func (s *ListRepositoryCommitsResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRepositoryCommitsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListRepositoryCommitsResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *ListRepositoryCommitsResponseBodyResult) GetParentIds() []*string {
	return s.ParentIds
}

func (s *ListRepositoryCommitsResponseBodyResult) GetShortId() *string {
	return s.ShortId
}

func (s *ListRepositoryCommitsResponseBodyResult) GetSignature() *ListRepositoryCommitsResponseBodyResultSignature {
	return s.Signature
}

func (s *ListRepositoryCommitsResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *ListRepositoryCommitsResponseBodyResult) SetAuthor(v *ListRepositoryCommitsResponseBodyResultAuthor) *ListRepositoryCommitsResponseBodyResult {
	s.Author = v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetAuthorEmail(v string) *ListRepositoryCommitsResponseBodyResult {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetAuthorName(v string) *ListRepositoryCommitsResponseBodyResult {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetAuthoredDate(v string) *ListRepositoryCommitsResponseBodyResult {
	s.AuthoredDate = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCommentsCount(v int64) *ListRepositoryCommitsResponseBodyResult {
	s.CommentsCount = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCommittedDate(v string) *ListRepositoryCommitsResponseBodyResult {
	s.CommittedDate = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCommitter(v *ListRepositoryCommitsResponseBodyResultCommitter) *ListRepositoryCommitsResponseBodyResult {
	s.Committer = v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCommitterEmail(v string) *ListRepositoryCommitsResponseBodyResult {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCommitterName(v string) *ListRepositoryCommitsResponseBodyResult {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCreatedAt(v string) *ListRepositoryCommitsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetId(v string) *ListRepositoryCommitsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetMessage(v string) *ListRepositoryCommitsResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetParentIds(v []*string) *ListRepositoryCommitsResponseBodyResult {
	s.ParentIds = v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetShortId(v string) *ListRepositoryCommitsResponseBodyResult {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetSignature(v *ListRepositoryCommitsResponseBodyResultSignature) *ListRepositoryCommitsResponseBodyResult {
	s.Signature = v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetTitle(v string) *ListRepositoryCommitsResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) Validate() error {
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

type ListRepositoryCommitsResponseBodyResultAuthor struct {
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 7914
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
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

func (s ListRepositoryCommitsResponseBodyResultAuthor) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitsResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) GetEmail() *string {
	return s.Email
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) GetId() *int64 {
	return s.Id
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) GetName() *string {
	return s.Name
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) GetState() *string {
	return s.State
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) GetUsername() *string {
	return s.Username
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) GetWebsiteUrl() *string {
	return s.WebsiteUrl
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) SetAvatarUrl(v string) *ListRepositoryCommitsResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) SetEmail(v string) *ListRepositoryCommitsResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) SetId(v int64) *ListRepositoryCommitsResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) SetName(v string) *ListRepositoryCommitsResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) SetState(v string) *ListRepositoryCommitsResponseBodyResultAuthor {
	s.State = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) SetUsername(v string) *ListRepositoryCommitsResponseBodyResultAuthor {
	s.Username = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) SetWebsiteUrl(v string) *ListRepositoryCommitsResponseBodyResultAuthor {
	s.WebsiteUrl = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultAuthor) Validate() error {
	return dara.Validate(s)
}

type ListRepositoryCommitsResponseBodyResultCommitter struct {
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
	// 41031
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

func (s ListRepositoryCommitsResponseBodyResultCommitter) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitsResponseBodyResultCommitter) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) GetEmail() *string {
	return s.Email
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) GetId() *int64 {
	return s.Id
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) GetName() *string {
	return s.Name
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) GetState() *string {
	return s.State
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) GetUsername() *string {
	return s.Username
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) GetWebsiteUrl() *string {
	return s.WebsiteUrl
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) SetAvatarUrl(v string) *ListRepositoryCommitsResponseBodyResultCommitter {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) SetEmail(v string) *ListRepositoryCommitsResponseBodyResultCommitter {
	s.Email = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) SetId(v int64) *ListRepositoryCommitsResponseBodyResultCommitter {
	s.Id = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) SetName(v string) *ListRepositoryCommitsResponseBodyResultCommitter {
	s.Name = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) SetState(v string) *ListRepositoryCommitsResponseBodyResultCommitter {
	s.State = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) SetUsername(v string) *ListRepositoryCommitsResponseBodyResultCommitter {
	s.Username = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) SetWebsiteUrl(v string) *ListRepositoryCommitsResponseBodyResultCommitter {
	s.WebsiteUrl = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultCommitter) Validate() error {
	return dara.Validate(s)
}

type ListRepositoryCommitsResponseBodyResultSignature struct {
	// example:
	//
	// ”“
	GpgKeyId *string `json:"gpgKeyId,omitempty" xml:"gpgKeyId,omitempty"`
	// example:
	//
	// verified
	VerificationStatus *string `json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
}

func (s ListRepositoryCommitsResponseBodyResultSignature) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitsResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponseBodyResultSignature) GetGpgKeyId() *string {
	return s.GpgKeyId
}

func (s *ListRepositoryCommitsResponseBodyResultSignature) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *ListRepositoryCommitsResponseBodyResultSignature) SetGpgKeyId(v string) *ListRepositoryCommitsResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultSignature) SetVerificationStatus(v string) *ListRepositoryCommitsResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultSignature) Validate() error {
	return dara.Validate(s)
}
