// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestCommentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListMergeRequestCommentsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMergeRequestCommentsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListMergeRequestCommentsResponseBody
	GetRequestId() *string
	SetResult(v []*ListMergeRequestCommentsResponseBodyResult) *ListMergeRequestCommentsResponseBody
	GetResult() []*ListMergeRequestCommentsResponseBodyResult
	SetSuccess(v bool) *ListMergeRequestCommentsResponseBody
	GetSuccess() *bool
}

type ListMergeRequestCommentsResponseBody struct {
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
	// 56C33A95-C04F-59F0-B3CD-E2A2EB9FADBB
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListMergeRequestCommentsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMergeRequestCommentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMergeRequestCommentsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMergeRequestCommentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMergeRequestCommentsResponseBody) GetResult() []*ListMergeRequestCommentsResponseBodyResult {
	return s.Result
}

func (s *ListMergeRequestCommentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMergeRequestCommentsResponseBody) SetErrorCode(v string) *ListMergeRequestCommentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetErrorMessage(v string) *ListMergeRequestCommentsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetRequestId(v string) *ListMergeRequestCommentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetResult(v []*ListMergeRequestCommentsResponseBodyResult) *ListMergeRequestCommentsResponseBody {
	s.Result = v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetSuccess(v bool) *ListMergeRequestCommentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) Validate() error {
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

type ListMergeRequestCommentsResponseBodyResult struct {
	Author        *ListMergeRequestCommentsResponseBodyResultAuthor          `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	ChildComments []*ListMergeRequestCommentsResponseBodyResultChildComments `json:"childComments,omitempty" xml:"childComments,omitempty" type:"Repeated"`
	// example:
	//
	// 5c399e3685e542a28db16d93e9f82abb
	CommentBizId *string `json:"commentBizId,omitempty" xml:"commentBizId,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CommentTime *string `json:"commentTime,omitempty" xml:"commentTime,omitempty"`
	// example:
	//
	// GLOBAL_COMMENT
	CommentType *string `json:"commentType,omitempty" xml:"commentType,omitempty"`
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// src/main/test.java
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	LastEditTime *string `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	// example:
	//
	// 3
	LineNumber *string `json:"lineNumber,omitempty" xml:"lineNumber,omitempty"`
	// example:
	//
	// 4c5dcec6a8dc41e69c369737dadc1841
	ParentCommentBizId *string                                                    `json:"parentCommentBizId,omitempty" xml:"parentCommentBizId,omitempty"`
	RelatedPatchSet    *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet `json:"relatedPatchSet,omitempty" xml:"relatedPatchSet,omitempty" type:"Struct"`
	// example:
	//
	// false
	Resolved *bool `json:"resolved,omitempty" xml:"resolved,omitempty"`
	// example:
	//
	// dcf2b23cebfc418f98dbd35e423d9fd3
	RootCommentBizId *string `json:"rootCommentBizId,omitempty" xml:"rootCommentBizId,omitempty"`
	// example:
	//
	// OPENED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetAuthor() *ListMergeRequestCommentsResponseBodyResultAuthor {
	return s.Author
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetChildComments() []*ListMergeRequestCommentsResponseBodyResultChildComments {
	return s.ChildComments
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetCommentBizId() *string {
	return s.CommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetCommentTime() *string {
	return s.CommentTime
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetCommentType() *string {
	return s.CommentType
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetDeleted() *bool {
	return s.Deleted
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetFilePath() *string {
	return s.FilePath
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetLastEditTime() *string {
	return s.LastEditTime
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetLineNumber() *string {
	return s.LineNumber
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetParentCommentBizId() *string {
	return s.ParentCommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetRelatedPatchSet() *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet {
	return s.RelatedPatchSet
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetResolved() *bool {
	return s.Resolved
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetRootCommentBizId() *string {
	return s.RootCommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetAuthor(v *ListMergeRequestCommentsResponseBodyResultAuthor) *ListMergeRequestCommentsResponseBodyResult {
	s.Author = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetChildComments(v []*ListMergeRequestCommentsResponseBodyResultChildComments) *ListMergeRequestCommentsResponseBodyResult {
	s.ChildComments = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.CommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetCommentTime(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.CommentTime = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetCommentType(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.CommentType = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetContent(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetDeleted(v bool) *ListMergeRequestCommentsResponseBodyResult {
	s.Deleted = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetFilePath(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetLastEditTime(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.LastEditTime = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetLineNumber(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.LineNumber = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetParentCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.ParentCommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetRelatedPatchSet(v *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) *ListMergeRequestCommentsResponseBodyResult {
	s.RelatedPatchSet = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetResolved(v bool) *ListMergeRequestCommentsResponseBodyResult {
	s.Resolved = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetRootCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.RootCommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetState(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) Validate() error {
	if s.Author != nil {
		if err := s.Author.Validate(); err != nil {
			return err
		}
	}
	if s.ChildComments != nil {
		for _, item := range s.ChildComments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedPatchSet != nil {
		if err := s.RelatedPatchSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMergeRequestCommentsResponseBodyResultAuthor struct {
	// example:
	//
	// 284692704493684695
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultAuthor) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) GetEmail() *string {
	return s.Email
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) GetId() *int64 {
	return s.Id
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) GetState() *string {
	return s.State
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) GetUsername() *string {
	return s.Username
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetAliyunPk(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.AliyunPk = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetAvatarUrl(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetEmail(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetId(v int64) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetName(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetState(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.State = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetUsername(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.Username = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) Validate() error {
	return dara.Validate(s)
}

type ListMergeRequestCommentsResponseBodyResultChildComments struct {
	Author *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	// example:
	//
	// 63f0e293357f48f9846ddc4dbbebd0e3
	CommentBizId *string `json:"commentBizId,omitempty" xml:"commentBizId,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CommentTime *string `json:"commentTime,omitempty" xml:"commentTime,omitempty"`
	// example:
	//
	// GLOBAL_COMMENT
	CommentType *string `json:"commentType,omitempty" xml:"commentType,omitempty"`
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// src/main/test.java
	FilePath           *string                                                                      `json:"filePath,omitempty" xml:"filePath,omitempty"`
	FinalChildComments []*ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments `json:"finalChildComments,omitempty" xml:"finalChildComments,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-03-18 14:24:54
	LastEditTime *string `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	// example:
	//
	// 3
	LineNumber *string `json:"lineNumber,omitempty" xml:"lineNumber,omitempty"`
	// example:
	//
	// 5c399e3685e542a28db16d93e9f82abb
	ParentCommentBizId *string                                                                 `json:"parentCommentBizId,omitempty" xml:"parentCommentBizId,omitempty"`
	RelatedPatchSet    *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet `json:"relatedPatchSet,omitempty" xml:"relatedPatchSet,omitempty" type:"Struct"`
	// example:
	//
	// false
	Resolved *bool `json:"resolved,omitempty" xml:"resolved,omitempty"`
	// example:
	//
	// dcf2b23cebfc418f98dbd35e423d9fd3
	RootCommentBizId *string `json:"rootCommentBizId,omitempty" xml:"rootCommentBizId,omitempty"`
	// example:
	//
	// OPENED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultChildComments) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultChildComments) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetAuthor() *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor {
	return s.Author
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetCommentBizId() *string {
	return s.CommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetCommentTime() *string {
	return s.CommentTime
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetCommentType() *string {
	return s.CommentType
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetContent() *string {
	return s.Content
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetDeleted() *bool {
	return s.Deleted
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetFilePath() *string {
	return s.FilePath
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetFinalChildComments() []*ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	return s.FinalChildComments
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetLastEditTime() *string {
	return s.LastEditTime
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetLineNumber() *string {
	return s.LineNumber
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetParentCommentBizId() *string {
	return s.ParentCommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetRelatedPatchSet() *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet {
	return s.RelatedPatchSet
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetResolved() *bool {
	return s.Resolved
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetRootCommentBizId() *string {
	return s.RootCommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) GetState() *string {
	return s.State
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetAuthor(v *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.Author = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.CommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetCommentTime(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.CommentTime = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetCommentType(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.CommentType = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetContent(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.Content = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetDeleted(v bool) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.Deleted = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetFilePath(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.FilePath = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetFinalChildComments(v []*ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.FinalChildComments = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetLastEditTime(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.LastEditTime = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetLineNumber(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.LineNumber = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetParentCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.ParentCommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetRelatedPatchSet(v *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.RelatedPatchSet = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetResolved(v bool) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.Resolved = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetRootCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.RootCommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) SetState(v string) *ListMergeRequestCommentsResponseBodyResultChildComments {
	s.State = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildComments) Validate() error {
	if s.Author != nil {
		if err := s.Author.Validate(); err != nil {
			return err
		}
	}
	if s.FinalChildComments != nil {
		for _, item := range s.FinalChildComments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedPatchSet != nil {
		if err := s.RelatedPatchSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor struct {
	// example:
	//
	// 204485087002425236
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) GetEmail() *string {
	return s.Email
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) GetId() *int64 {
	return s.Id
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) GetState() *string {
	return s.State
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) GetUsername() *string {
	return s.Username
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) SetAliyunPk(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor {
	s.AliyunPk = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) SetAvatarUrl(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) SetEmail(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor {
	s.Email = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) SetId(v int64) *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor {
	s.Id = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) SetName(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor {
	s.Name = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) SetState(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor {
	s.State = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) SetUsername(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor {
	s.Username = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsAuthor) Validate() error {
	return dara.Validate(s)
}

type ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments struct {
	Author *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	// example:
	//
	// 5c399e3685e542a28db16d93e9f82abb
	CommentBizId *string `json:"commentBizId,omitempty" xml:"commentBizId,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CommentTime *string `json:"commentTime,omitempty" xml:"commentTime,omitempty"`
	// example:
	//
	// GLOBAL_COMMENT
	CommentType *string `json:"commentType,omitempty" xml:"commentType,omitempty"`
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// src/main/test.java
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	LastEditTime *string `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	// example:
	//
	// 3
	LineNumber *string `json:"lineNumber,omitempty" xml:"lineNumber,omitempty"`
	// example:
	//
	// 4c5dcec6a8dc41e69c369737dadc1841
	ParentCommentBizId *string                                                                                   `json:"parentCommentBizId,omitempty" xml:"parentCommentBizId,omitempty"`
	RelatedPatchSet    *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet `json:"relatedPatchSet,omitempty" xml:"relatedPatchSet,omitempty" type:"Struct"`
	// example:
	//
	// false
	Resolved *bool `json:"resolved,omitempty" xml:"resolved,omitempty"`
	// example:
	//
	// dcf2b23cebfc418f98dbd35e423d9fd3
	RootCommentBizId *string `json:"rootCommentBizId,omitempty" xml:"rootCommentBizId,omitempty"`
	// example:
	//
	// OPENED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetAuthor() *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor {
	return s.Author
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetCommentBizId() *string {
	return s.CommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetCommentTime() *string {
	return s.CommentTime
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetCommentType() *string {
	return s.CommentType
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetContent() *string {
	return s.Content
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetDeleted() *bool {
	return s.Deleted
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetFilePath() *string {
	return s.FilePath
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetLastEditTime() *string {
	return s.LastEditTime
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetLineNumber() *string {
	return s.LineNumber
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetParentCommentBizId() *string {
	return s.ParentCommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetRelatedPatchSet() *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet {
	return s.RelatedPatchSet
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetResolved() *bool {
	return s.Resolved
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetRootCommentBizId() *string {
	return s.RootCommentBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) GetState() *string {
	return s.State
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetAuthor(v *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.Author = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.CommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetCommentTime(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.CommentTime = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetCommentType(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.CommentType = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetContent(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.Content = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetDeleted(v bool) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.Deleted = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetFilePath(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.FilePath = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetLastEditTime(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.LastEditTime = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetLineNumber(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.LineNumber = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetParentCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.ParentCommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetRelatedPatchSet(v *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.RelatedPatchSet = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetResolved(v bool) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.Resolved = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetRootCommentBizId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.RootCommentBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) SetState(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments {
	s.State = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildComments) Validate() error {
	if s.Author != nil {
		if err := s.Author.Validate(); err != nil {
			return err
		}
	}
	if s.RelatedPatchSet != nil {
		if err := s.RelatedPatchSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor struct {
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
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) GetEmail() *string {
	return s.Email
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) GetId() *int64 {
	return s.Id
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) GetState() *string {
	return s.State
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) GetUsername() *string {
	return s.Username
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) SetAliyunPk(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor {
	s.AliyunPk = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) SetAvatarUrl(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) SetEmail(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor {
	s.Email = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) SetId(v int64) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor {
	s.Id = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) SetName(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor {
	s.Name = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) SetState(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor {
	s.State = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) SetUsername(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor {
	s.Username = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsAuthor) Validate() error {
	return dara.Validate(s)
}

type ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet struct {
	// example:
	//
	// 1a072f5367c21f9de3464b8c0ee8546e47764d2d
	CommitId *string `json:"commitId,omitempty" xml:"commitId,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 513fcfd81a9142d2bb0db4f72c0aa15b
	PatchSetBizId *string `json:"patchSetBizId,omitempty" xml:"patchSetBizId,omitempty"`
	PatchSetName  *string `json:"patchSetName,omitempty" xml:"patchSetName,omitempty"`
	// example:
	//
	// 1
	PatchSetNo *string `json:"patchSetNo,omitempty" xml:"patchSetNo,omitempty"`
	// example:
	//
	// MERGE_SOURCE
	RelatedMergeItemType *string `json:"relatedMergeItemType,omitempty" xml:"relatedMergeItemType,omitempty"`
	// example:
	//
	// 1a072f53
	ShortId *string `json:"shortId,omitempty" xml:"shortId,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) GetCommitId() *string {
	return s.CommitId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) GetPatchSetBizId() *string {
	return s.PatchSetBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) GetPatchSetName() *string {
	return s.PatchSetName
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) GetPatchSetNo() *string {
	return s.PatchSetNo
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) GetRelatedMergeItemType() *string {
	return s.RelatedMergeItemType
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) GetShortId() *string {
	return s.ShortId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) SetCommitId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet {
	s.CommitId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) SetCreatedAt(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet {
	s.CreatedAt = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) SetPatchSetBizId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet {
	s.PatchSetBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) SetPatchSetName(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet {
	s.PatchSetName = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) SetPatchSetNo(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet {
	s.PatchSetNo = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) SetRelatedMergeItemType(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet {
	s.RelatedMergeItemType = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) SetShortId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet {
	s.ShortId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsFinalChildCommentsRelatedPatchSet) Validate() error {
	return dara.Validate(s)
}

type ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet struct {
	// example:
	//
	// 1a072f5367c21f9de3464b8c0ee8546e47764d2d
	CommitId *string `json:"commitId,omitempty" xml:"commitId,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 513fcfd81a9142d2bb0db4f72c0aa15b
	PatchSetBizId *string `json:"patchSetBizId,omitempty" xml:"patchSetBizId,omitempty"`
	PatchSetName  *string `json:"patchSetName,omitempty" xml:"patchSetName,omitempty"`
	// example:
	//
	// 1
	PatchSetNo *string `json:"patchSetNo,omitempty" xml:"patchSetNo,omitempty"`
	// example:
	//
	// MERGE_SOURCE
	RelatedMergeItemType *string `json:"relatedMergeItemType,omitempty" xml:"relatedMergeItemType,omitempty"`
	// example:
	//
	// 1a072f53
	ShortId *string `json:"shortId,omitempty" xml:"shortId,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) GetCommitId() *string {
	return s.CommitId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) GetPatchSetBizId() *string {
	return s.PatchSetBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) GetPatchSetName() *string {
	return s.PatchSetName
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) GetPatchSetNo() *string {
	return s.PatchSetNo
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) GetRelatedMergeItemType() *string {
	return s.RelatedMergeItemType
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) GetShortId() *string {
	return s.ShortId
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) SetCommitId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet {
	s.CommitId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) SetCreatedAt(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet {
	s.CreatedAt = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) SetPatchSetBizId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet {
	s.PatchSetBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) SetPatchSetName(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet {
	s.PatchSetName = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) SetPatchSetNo(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet {
	s.PatchSetNo = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) SetRelatedMergeItemType(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet {
	s.RelatedMergeItemType = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) SetShortId(v string) *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet {
	s.ShortId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultChildCommentsRelatedPatchSet) Validate() error {
	return dara.Validate(s)
}

type ListMergeRequestCommentsResponseBodyResultRelatedPatchSet struct {
	// example:
	//
	// 1a072f5367c21f9de3464b8c0ee8546e47764d2d
	CommitId *string `json:"commitId,omitempty" xml:"commitId,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 513fcfd81a9142d2bb0db4f72c0aa15b
	PatchSetBizId *string `json:"patchSetBizId,omitempty" xml:"patchSetBizId,omitempty"`
	PatchSetName  *string `json:"patchSetName,omitempty" xml:"patchSetName,omitempty"`
	// example:
	//
	// 1
	PatchSetNo *string `json:"patchSetNo,omitempty" xml:"patchSetNo,omitempty"`
	// example:
	//
	// MERGE_SOURCE
	RelatedMergeItemType *string `json:"relatedMergeItemType,omitempty" xml:"relatedMergeItemType,omitempty"`
	// example:
	//
	// 1a072f53
	ShortId *string `json:"shortId,omitempty" xml:"shortId,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) GetCommitId() *string {
	return s.CommitId
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) GetPatchSetBizId() *string {
	return s.PatchSetBizId
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) GetPatchSetName() *string {
	return s.PatchSetName
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) GetPatchSetNo() *string {
	return s.PatchSetNo
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) GetRelatedMergeItemType() *string {
	return s.RelatedMergeItemType
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) GetShortId() *string {
	return s.ShortId
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) SetCommitId(v string) *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet {
	s.CommitId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) SetCreatedAt(v string) *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet {
	s.CreatedAt = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) SetPatchSetBizId(v string) *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet {
	s.PatchSetBizId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) SetPatchSetName(v string) *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet {
	s.PatchSetName = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) SetPatchSetNo(v string) *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet {
	s.PatchSetNo = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) SetRelatedMergeItemType(v string) *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet {
	s.RelatedMergeItemType = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) SetShortId(v string) *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet {
	s.ShortId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultRelatedPatchSet) Validate() error {
	return dara.Validate(s)
}
