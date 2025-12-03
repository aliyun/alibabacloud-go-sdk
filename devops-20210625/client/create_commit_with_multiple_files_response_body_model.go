// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommitWithMultipleFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateCommitWithMultipleFilesResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateCommitWithMultipleFilesResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateCommitWithMultipleFilesResponseBody
	GetRequestId() *string
	SetResult(v *CreateCommitWithMultipleFilesResponseBodyResult) *CreateCommitWithMultipleFilesResponseBody
	GetResult() *CreateCommitWithMultipleFilesResponseBodyResult
	SetSuccess(v bool) *CreateCommitWithMultipleFilesResponseBody
	GetSuccess() *bool
}

type CreateCommitWithMultipleFilesResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// A7586FEB-E48D-5579-983F-74981FBFF627
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateCommitWithMultipleFilesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCommitWithMultipleFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitWithMultipleFilesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommitWithMultipleFilesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateCommitWithMultipleFilesResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateCommitWithMultipleFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCommitWithMultipleFilesResponseBody) GetResult() *CreateCommitWithMultipleFilesResponseBodyResult {
	return s.Result
}

func (s *CreateCommitWithMultipleFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCommitWithMultipleFilesResponseBody) SetErrorCode(v string) *CreateCommitWithMultipleFilesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBody) SetErrorMsg(v string) *CreateCommitWithMultipleFilesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBody) SetRequestId(v string) *CreateCommitWithMultipleFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBody) SetResult(v *CreateCommitWithMultipleFilesResponseBodyResult) *CreateCommitWithMultipleFilesResponseBody {
	s.Result = v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBody) SetSuccess(v bool) *CreateCommitWithMultipleFilesResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCommitWithMultipleFilesResponseBodyResult struct {
	Author *CreateCommitWithMultipleFilesResponseBodyResultAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
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
	// 2022-03-18 16:00:00
	CommittedDate *string                                                   `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	Committer     *CreateCommitWithMultipleFilesResponseBodyResultCommitter `json:"committer,omitempty" xml:"committer,omitempty" type:"Struct"`
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
	ShortId *string `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCommitWithMultipleFilesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitWithMultipleFilesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetAuthor() *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	return s.Author
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetAuthorName() *string {
	return s.AuthorName
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetCommitter() *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	return s.Committer
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetCommitterName() *string {
	return s.CommitterName
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetParentIds() []*string {
	return s.ParentIds
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetShortId() *string {
	return s.ShortId
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetAuthor(v *CreateCommitWithMultipleFilesResponseBodyResultAuthor) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.Author = v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetAuthorEmail(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.AuthorEmail = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetAuthorName(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.AuthorName = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetAuthoredDate(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.AuthoredDate = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetCommittedDate(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.CommittedDate = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetCommitter(v *CreateCommitWithMultipleFilesResponseBodyResultCommitter) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.Committer = v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetCommitterEmail(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.CommitterEmail = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetCommitterName(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.CommitterName = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetCreatedAt(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetId(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetMessage(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.Message = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetParentIds(v []*string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.ParentIds = v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetShortId(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.ShortId = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) SetTitle(v string) *CreateCommitWithMultipleFilesResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResult) Validate() error {
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
	return nil
}

type CreateCommitWithMultipleFilesResponseBodyResultAuthor struct {
	// example:
	//
	// 284692704493684695
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
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
	// test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// example:
	//
	// ""
	WebsiteUrl *string `json:"websiteUrl,omitempty" xml:"websiteUrl,omitempty"`
}

func (s CreateCommitWithMultipleFilesResponseBodyResultAuthor) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitWithMultipleFilesResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) GetEmail() *string {
	return s.Email
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) GetId() *int64 {
	return s.Id
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) GetName() *string {
	return s.Name
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) GetState() *string {
	return s.State
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) GetUsername() *string {
	return s.Username
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) GetWebsiteUrl() *string {
	return s.WebsiteUrl
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) SetAliyunPk(v string) *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	s.AliyunPk = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) SetAvatarUrl(v string) *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) SetEmail(v string) *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) SetId(v int64) *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) SetName(v string) *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) SetState(v string) *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	s.State = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) SetUsername(v string) *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	s.Username = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) SetWebsiteUrl(v string) *CreateCommitWithMultipleFilesResponseBodyResultAuthor {
	s.WebsiteUrl = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultAuthor) Validate() error {
	return dara.Validate(s)
}

type CreateCommitWithMultipleFilesResponseBodyResultCommitter struct {
	// example:
	//
	// 235671547828975455
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
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
	// committer-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// example:
	//
	// ""
	WebsiteUrl *string `json:"websiteUrl,omitempty" xml:"websiteUrl,omitempty"`
}

func (s CreateCommitWithMultipleFilesResponseBodyResultCommitter) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitWithMultipleFilesResponseBodyResultCommitter) GoString() string {
	return s.String()
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) GetEmail() *string {
	return s.Email
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) GetId() *int64 {
	return s.Id
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) GetName() *string {
	return s.Name
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) GetState() *string {
	return s.State
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) GetUsername() *string {
	return s.Username
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) GetWebsiteUrl() *string {
	return s.WebsiteUrl
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) SetAliyunPk(v string) *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	s.AliyunPk = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) SetAvatarUrl(v string) *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	s.AvatarUrl = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) SetEmail(v string) *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	s.Email = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) SetId(v int64) *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	s.Id = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) SetName(v string) *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	s.Name = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) SetState(v string) *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	s.State = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) SetUsername(v string) *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	s.Username = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) SetWebsiteUrl(v string) *CreateCommitWithMultipleFilesResponseBodyResultCommitter {
	s.WebsiteUrl = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponseBodyResultCommitter) Validate() error {
	return dara.Validate(s)
}
