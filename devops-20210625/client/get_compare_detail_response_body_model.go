// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompareDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetCompareDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetCompareDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetCompareDetailResponseBody
	GetRequestId() *string
	SetResult(v *GetCompareDetailResponseBodyResult) *GetCompareDetailResponseBody
	GetResult() *GetCompareDetailResponseBodyResult
	SetSuccess(v bool) *GetCompareDetailResponseBody
	GetSuccess() *bool
}

type GetCompareDetailResponseBody struct {
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetCompareDetailResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCompareDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCompareDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompareDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCompareDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCompareDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCompareDetailResponseBody) GetResult() *GetCompareDetailResponseBodyResult {
	return s.Result
}

func (s *GetCompareDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCompareDetailResponseBody) SetErrorCode(v string) *GetCompareDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCompareDetailResponseBody) SetErrorMessage(v string) *GetCompareDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCompareDetailResponseBody) SetRequestId(v string) *GetCompareDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCompareDetailResponseBody) SetResult(v *GetCompareDetailResponseBodyResult) *GetCompareDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetCompareDetailResponseBody) SetSuccess(v bool) *GetCompareDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetCompareDetailResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCompareDetailResponseBodyResult struct {
	Commits  []*GetCompareDetailResponseBodyResultCommits `json:"commits,omitempty" xml:"commits,omitempty" type:"Repeated"`
	Diffs    []*GetCompareDetailResponseBodyResultDiffs   `json:"diffs,omitempty" xml:"diffs,omitempty" type:"Repeated"`
	Messages []*string                                    `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s GetCompareDetailResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCompareDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCompareDetailResponseBodyResult) GetCommits() []*GetCompareDetailResponseBodyResultCommits {
	return s.Commits
}

func (s *GetCompareDetailResponseBodyResult) GetDiffs() []*GetCompareDetailResponseBodyResultDiffs {
	return s.Diffs
}

func (s *GetCompareDetailResponseBodyResult) GetMessages() []*string {
	return s.Messages
}

func (s *GetCompareDetailResponseBodyResult) SetCommits(v []*GetCompareDetailResponseBodyResultCommits) *GetCompareDetailResponseBodyResult {
	s.Commits = v
	return s
}

func (s *GetCompareDetailResponseBodyResult) SetDiffs(v []*GetCompareDetailResponseBodyResultDiffs) *GetCompareDetailResponseBodyResult {
	s.Diffs = v
	return s
}

func (s *GetCompareDetailResponseBodyResult) SetMessages(v []*string) *GetCompareDetailResponseBodyResult {
	s.Messages = v
	return s
}

func (s *GetCompareDetailResponseBodyResult) Validate() error {
	if s.Commits != nil {
		for _, item := range s.Commits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Diffs != nil {
		for _, item := range s.Diffs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCompareDetailResponseBodyResultCommits struct {
	Author *GetCompareDetailResponseBodyResultCommitsAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	// example:
	//
	// username@example.com
	AuthorEmail *string `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	AuthorName  *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	// example:
	//
	// 2023-01-03T15:41:26+08:00
	AuthoredDate *string `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	// example:
	//
	// 1
	CommentsCount *int64 `json:"commentsCount,omitempty" xml:"commentsCount,omitempty"`
	// example:
	//
	// 2023-01-03T15:41:26+08:00
	CommittedDate *string                                             `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	Committer     *GetCompareDetailResponseBodyResultCommitsCommitter `json:"committer,omitempty" xml:"committer,omitempty" type:"Struct"`
	// example:
	//
	// username@example.com
	CommitterEmail *string `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	CommitterName  *string `json:"committerName,omitempty" xml:"committerName,omitempty"`
	// example:
	//
	// 2023-01-03T15:41:26+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// b8f6f28520b1936aafe2e638373e19ccafa42b02
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ""
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// b8f6f285
	ShortId *string `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetCompareDetailResponseBodyResultCommits) String() string {
	return dara.Prettify(s)
}

func (s GetCompareDetailResponseBodyResultCommits) GoString() string {
	return s.String()
}

func (s *GetCompareDetailResponseBodyResultCommits) GetAuthor() *GetCompareDetailResponseBodyResultCommitsAuthor {
	return s.Author
}

func (s *GetCompareDetailResponseBodyResultCommits) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *GetCompareDetailResponseBodyResultCommits) GetAuthorName() *string {
	return s.AuthorName
}

func (s *GetCompareDetailResponseBodyResultCommits) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *GetCompareDetailResponseBodyResultCommits) GetCommentsCount() *int64 {
	return s.CommentsCount
}

func (s *GetCompareDetailResponseBodyResultCommits) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *GetCompareDetailResponseBodyResultCommits) GetCommitter() *GetCompareDetailResponseBodyResultCommitsCommitter {
	return s.Committer
}

func (s *GetCompareDetailResponseBodyResultCommits) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *GetCompareDetailResponseBodyResultCommits) GetCommitterName() *string {
	return s.CommitterName
}

func (s *GetCompareDetailResponseBodyResultCommits) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCompareDetailResponseBodyResultCommits) GetId() *string {
	return s.Id
}

func (s *GetCompareDetailResponseBodyResultCommits) GetMessage() *string {
	return s.Message
}

func (s *GetCompareDetailResponseBodyResultCommits) GetParentIds() []*string {
	return s.ParentIds
}

func (s *GetCompareDetailResponseBodyResultCommits) GetShortId() *string {
	return s.ShortId
}

func (s *GetCompareDetailResponseBodyResultCommits) GetTitle() *string {
	return s.Title
}

func (s *GetCompareDetailResponseBodyResultCommits) SetAuthor(v *GetCompareDetailResponseBodyResultCommitsAuthor) *GetCompareDetailResponseBodyResultCommits {
	s.Author = v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetAuthorEmail(v string) *GetCompareDetailResponseBodyResultCommits {
	s.AuthorEmail = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetAuthorName(v string) *GetCompareDetailResponseBodyResultCommits {
	s.AuthorName = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetAuthoredDate(v string) *GetCompareDetailResponseBodyResultCommits {
	s.AuthoredDate = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetCommentsCount(v int64) *GetCompareDetailResponseBodyResultCommits {
	s.CommentsCount = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetCommittedDate(v string) *GetCompareDetailResponseBodyResultCommits {
	s.CommittedDate = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetCommitter(v *GetCompareDetailResponseBodyResultCommitsCommitter) *GetCompareDetailResponseBodyResultCommits {
	s.Committer = v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetCommitterEmail(v string) *GetCompareDetailResponseBodyResultCommits {
	s.CommitterEmail = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetCommitterName(v string) *GetCompareDetailResponseBodyResultCommits {
	s.CommitterName = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetCreatedAt(v string) *GetCompareDetailResponseBodyResultCommits {
	s.CreatedAt = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetId(v string) *GetCompareDetailResponseBodyResultCommits {
	s.Id = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetMessage(v string) *GetCompareDetailResponseBodyResultCommits {
	s.Message = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetParentIds(v []*string) *GetCompareDetailResponseBodyResultCommits {
	s.ParentIds = v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetShortId(v string) *GetCompareDetailResponseBodyResultCommits {
	s.ShortId = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) SetTitle(v string) *GetCompareDetailResponseBodyResultCommits {
	s.Title = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommits) Validate() error {
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

type GetCompareDetailResponseBodyResultCommitsAuthor struct {
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
	// 21396
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// Codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetCompareDetailResponseBodyResultCommitsAuthor) String() string {
	return dara.Prettify(s)
}

func (s GetCompareDetailResponseBodyResultCommitsAuthor) GoString() string {
	return s.String()
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) GetEmail() *string {
	return s.Email
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) GetId() *int64 {
	return s.Id
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) GetName() *string {
	return s.Name
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) GetState() *string {
	return s.State
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) GetUsername() *string {
	return s.Username
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) SetAvatarUrl(v string) *GetCompareDetailResponseBodyResultCommitsAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) SetEmail(v string) *GetCompareDetailResponseBodyResultCommitsAuthor {
	s.Email = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) SetId(v int64) *GetCompareDetailResponseBodyResultCommitsAuthor {
	s.Id = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) SetName(v string) *GetCompareDetailResponseBodyResultCommitsAuthor {
	s.Name = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) SetState(v string) *GetCompareDetailResponseBodyResultCommitsAuthor {
	s.State = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) SetUsername(v string) *GetCompareDetailResponseBodyResultCommitsAuthor {
	s.Username = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsAuthor) Validate() error {
	return dara.Validate(s)
}

type GetCompareDetailResponseBodyResultCommitsCommitter struct {
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
	// 43910
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// Codeup-commiter
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetCompareDetailResponseBodyResultCommitsCommitter) String() string {
	return dara.Prettify(s)
}

func (s GetCompareDetailResponseBodyResultCommitsCommitter) GoString() string {
	return s.String()
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) GetEmail() *string {
	return s.Email
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) GetId() *int64 {
	return s.Id
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) GetName() *string {
	return s.Name
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) GetState() *string {
	return s.State
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) GetUsername() *string {
	return s.Username
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) SetAvatarUrl(v string) *GetCompareDetailResponseBodyResultCommitsCommitter {
	s.AvatarUrl = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) SetEmail(v string) *GetCompareDetailResponseBodyResultCommitsCommitter {
	s.Email = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) SetId(v int64) *GetCompareDetailResponseBodyResultCommitsCommitter {
	s.Id = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) SetName(v string) *GetCompareDetailResponseBodyResultCommitsCommitter {
	s.Name = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) SetState(v string) *GetCompareDetailResponseBodyResultCommitsCommitter {
	s.State = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) SetUsername(v string) *GetCompareDetailResponseBodyResultCommitsCommitter {
	s.Username = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultCommitsCommitter) Validate() error {
	return dara.Validate(s)
}

type GetCompareDetailResponseBodyResultDiffs struct {
	// example:
	//
	// 100644
	AMode *string `json:"aMode,omitempty" xml:"aMode,omitempty"`
	// example:
	//
	// 100644
	BMode *string `json:"bMode,omitempty" xml:"bMode,omitempty"`
	// example:
	//
	// false
	DeletedFile *bool   `json:"deletedFile,omitempty" xml:"deletedFile,omitempty"`
	Diff        *string `json:"diff,omitempty" xml:"diff,omitempty"`
	// example:
	//
	// false
	IsBinary *bool `json:"isBinary,omitempty" xml:"isBinary,omitempty"`
	// example:
	//
	// false
	IsNewLfs *bool `json:"isNewLfs,omitempty" xml:"isNewLfs,omitempty"`
	// example:
	//
	// false
	IsOldLfs *bool `json:"isOldLfs,omitempty" xml:"isOldLfs,omitempty"`
	// example:
	//
	// false
	NewFile *bool `json:"newFile,omitempty" xml:"newFile,omitempty"`
	// example:
	//
	// cb75846da2df3d3d7f290c3569236fcf3dd17224
	NewId *string `json:"newId,omitempty" xml:"newId,omitempty"`
	// example:
	//
	// new_test.txt
	NewPath *string `json:"newPath,omitempty" xml:"newPath,omitempty"`
	// example:
	//
	// 6c268061a546378276559c713d0ad377d4dsjfh
	OldId *string `json:"oldId,omitempty" xml:"oldId,omitempty"`
	// example:
	//
	// test.txt
	OldPath *string `json:"oldPath,omitempty" xml:"oldPath,omitempty"`
	// example:
	//
	// false
	RenamedFile *bool `json:"renamedFile,omitempty" xml:"renamedFile,omitempty"`
}

func (s GetCompareDetailResponseBodyResultDiffs) String() string {
	return dara.Prettify(s)
}

func (s GetCompareDetailResponseBodyResultDiffs) GoString() string {
	return s.String()
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetAMode() *string {
	return s.AMode
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetBMode() *string {
	return s.BMode
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetDeletedFile() *bool {
	return s.DeletedFile
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetDiff() *string {
	return s.Diff
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetIsBinary() *bool {
	return s.IsBinary
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetIsNewLfs() *bool {
	return s.IsNewLfs
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetIsOldLfs() *bool {
	return s.IsOldLfs
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetNewFile() *bool {
	return s.NewFile
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetNewId() *string {
	return s.NewId
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetNewPath() *string {
	return s.NewPath
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetOldId() *string {
	return s.OldId
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetOldPath() *string {
	return s.OldPath
}

func (s *GetCompareDetailResponseBodyResultDiffs) GetRenamedFile() *bool {
	return s.RenamedFile
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetAMode(v string) *GetCompareDetailResponseBodyResultDiffs {
	s.AMode = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetBMode(v string) *GetCompareDetailResponseBodyResultDiffs {
	s.BMode = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetDeletedFile(v bool) *GetCompareDetailResponseBodyResultDiffs {
	s.DeletedFile = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetDiff(v string) *GetCompareDetailResponseBodyResultDiffs {
	s.Diff = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetIsBinary(v bool) *GetCompareDetailResponseBodyResultDiffs {
	s.IsBinary = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetIsNewLfs(v bool) *GetCompareDetailResponseBodyResultDiffs {
	s.IsNewLfs = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetIsOldLfs(v bool) *GetCompareDetailResponseBodyResultDiffs {
	s.IsOldLfs = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetNewFile(v bool) *GetCompareDetailResponseBodyResultDiffs {
	s.NewFile = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetNewId(v string) *GetCompareDetailResponseBodyResultDiffs {
	s.NewId = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetNewPath(v string) *GetCompareDetailResponseBodyResultDiffs {
	s.NewPath = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetOldId(v string) *GetCompareDetailResponseBodyResultDiffs {
	s.OldId = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetOldPath(v string) *GetCompareDetailResponseBodyResultDiffs {
	s.OldPath = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) SetRenamedFile(v bool) *GetCompareDetailResponseBodyResultDiffs {
	s.RenamedFile = &v
	return s
}

func (s *GetCompareDetailResponseBodyResultDiffs) Validate() error {
	return dara.Validate(s)
}
