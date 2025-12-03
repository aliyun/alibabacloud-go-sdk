// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateCommentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateCommentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateCommentResponseBody
	GetRequestId() *string
	SetResult(v *CreateCommentResponseBodyResult) *CreateCommentResponseBody
	GetResult() *CreateCommentResponseBodyResult
	SetSuccess(v bool) *CreateCommentResponseBody
	GetSuccess() *bool
}

type CreateCommentResponseBody struct {
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateCommentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCommentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCommentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateCommentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateCommentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCommentResponseBody) GetResult() *CreateCommentResponseBodyResult {
	return s.Result
}

func (s *CreateCommentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCommentResponseBody) SetErrorCode(v string) *CreateCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCommentResponseBody) SetErrorMessage(v string) *CreateCommentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateCommentResponseBody) SetRequestId(v string) *CreateCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommentResponseBody) SetResult(v *CreateCommentResponseBodyResult) *CreateCommentResponseBody {
	s.Result = v
	return s
}

func (s *CreateCommentResponseBody) SetSuccess(v bool) *CreateCommentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCommentResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCommentResponseBodyResult struct {
	Author *CreateCommentResponseBodyResultAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
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
	// example:
	//
	// xxxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// false
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// example:
	//
	// src/main/update.txt
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
	ParentCommentBizId *string                                         `json:"parentCommentBizId,omitempty" xml:"parentCommentBizId,omitempty"`
	RelatedPatchSet    *CreateCommentResponseBodyResultRelatedPatchSet `json:"relatedPatchSet,omitempty" xml:"relatedPatchSet,omitempty" type:"Struct"`
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

func (s CreateCommentResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateCommentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCommentResponseBodyResult) GetAuthor() *CreateCommentResponseBodyResultAuthor {
	return s.Author
}

func (s *CreateCommentResponseBodyResult) GetCommentBizId() *string {
	return s.CommentBizId
}

func (s *CreateCommentResponseBodyResult) GetCommentTime() *string {
	return s.CommentTime
}

func (s *CreateCommentResponseBodyResult) GetCommentType() *string {
	return s.CommentType
}

func (s *CreateCommentResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *CreateCommentResponseBodyResult) GetDeleted() *bool {
	return s.Deleted
}

func (s *CreateCommentResponseBodyResult) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateCommentResponseBodyResult) GetLastEditTime() *string {
	return s.LastEditTime
}

func (s *CreateCommentResponseBodyResult) GetLineNumber() *string {
	return s.LineNumber
}

func (s *CreateCommentResponseBodyResult) GetParentCommentBizId() *string {
	return s.ParentCommentBizId
}

func (s *CreateCommentResponseBodyResult) GetRelatedPatchSet() *CreateCommentResponseBodyResultRelatedPatchSet {
	return s.RelatedPatchSet
}

func (s *CreateCommentResponseBodyResult) GetResolved() *bool {
	return s.Resolved
}

func (s *CreateCommentResponseBodyResult) GetRootCommentBizId() *string {
	return s.RootCommentBizId
}

func (s *CreateCommentResponseBodyResult) GetState() *string {
	return s.State
}

func (s *CreateCommentResponseBodyResult) SetAuthor(v *CreateCommentResponseBodyResultAuthor) *CreateCommentResponseBodyResult {
	s.Author = v
	return s
}

func (s *CreateCommentResponseBodyResult) SetCommentBizId(v string) *CreateCommentResponseBodyResult {
	s.CommentBizId = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetCommentTime(v string) *CreateCommentResponseBodyResult {
	s.CommentTime = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetCommentType(v string) *CreateCommentResponseBodyResult {
	s.CommentType = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetContent(v string) *CreateCommentResponseBodyResult {
	s.Content = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetDeleted(v bool) *CreateCommentResponseBodyResult {
	s.Deleted = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetFilePath(v string) *CreateCommentResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetLastEditTime(v string) *CreateCommentResponseBodyResult {
	s.LastEditTime = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetLineNumber(v string) *CreateCommentResponseBodyResult {
	s.LineNumber = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetParentCommentBizId(v string) *CreateCommentResponseBodyResult {
	s.ParentCommentBizId = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetRelatedPatchSet(v *CreateCommentResponseBodyResultRelatedPatchSet) *CreateCommentResponseBodyResult {
	s.RelatedPatchSet = v
	return s
}

func (s *CreateCommentResponseBodyResult) SetResolved(v bool) *CreateCommentResponseBodyResult {
	s.Resolved = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetRootCommentBizId(v string) *CreateCommentResponseBodyResult {
	s.RootCommentBizId = &v
	return s
}

func (s *CreateCommentResponseBodyResult) SetState(v string) *CreateCommentResponseBodyResult {
	s.State = &v
	return s
}

func (s *CreateCommentResponseBodyResult) Validate() error {
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

type CreateCommentResponseBodyResultAuthor struct {
	// example:
	//
	// 284692704493684695
	AliyunPk  *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
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
}

func (s CreateCommentResponseBodyResultAuthor) String() string {
	return dara.Prettify(s)
}

func (s CreateCommentResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *CreateCommentResponseBodyResultAuthor) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *CreateCommentResponseBodyResultAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateCommentResponseBodyResultAuthor) GetEmail() *string {
	return s.Email
}

func (s *CreateCommentResponseBodyResultAuthor) GetId() *int64 {
	return s.Id
}

func (s *CreateCommentResponseBodyResultAuthor) GetName() *string {
	return s.Name
}

func (s *CreateCommentResponseBodyResultAuthor) GetState() *string {
	return s.State
}

func (s *CreateCommentResponseBodyResultAuthor) GetUsername() *string {
	return s.Username
}

func (s *CreateCommentResponseBodyResultAuthor) SetAliyunPk(v string) *CreateCommentResponseBodyResultAuthor {
	s.AliyunPk = &v
	return s
}

func (s *CreateCommentResponseBodyResultAuthor) SetAvatarUrl(v string) *CreateCommentResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *CreateCommentResponseBodyResultAuthor) SetEmail(v string) *CreateCommentResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *CreateCommentResponseBodyResultAuthor) SetId(v int64) *CreateCommentResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *CreateCommentResponseBodyResultAuthor) SetName(v string) *CreateCommentResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *CreateCommentResponseBodyResultAuthor) SetState(v string) *CreateCommentResponseBodyResultAuthor {
	s.State = &v
	return s
}

func (s *CreateCommentResponseBodyResultAuthor) SetUsername(v string) *CreateCommentResponseBodyResultAuthor {
	s.Username = &v
	return s
}

func (s *CreateCommentResponseBodyResultAuthor) Validate() error {
	return dara.Validate(s)
}

type CreateCommentResponseBodyResultRelatedPatchSet struct {
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

func (s CreateCommentResponseBodyResultRelatedPatchSet) String() string {
	return dara.Prettify(s)
}

func (s CreateCommentResponseBodyResultRelatedPatchSet) GoString() string {
	return s.String()
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) GetCommitId() *string {
	return s.CommitId
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) GetPatchSetBizId() *string {
	return s.PatchSetBizId
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) GetPatchSetName() *string {
	return s.PatchSetName
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) GetPatchSetNo() *string {
	return s.PatchSetNo
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) GetRelatedMergeItemType() *string {
	return s.RelatedMergeItemType
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) GetShortId() *string {
	return s.ShortId
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) SetCommitId(v string) *CreateCommentResponseBodyResultRelatedPatchSet {
	s.CommitId = &v
	return s
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) SetCreatedAt(v string) *CreateCommentResponseBodyResultRelatedPatchSet {
	s.CreatedAt = &v
	return s
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) SetPatchSetBizId(v string) *CreateCommentResponseBodyResultRelatedPatchSet {
	s.PatchSetBizId = &v
	return s
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) SetPatchSetName(v string) *CreateCommentResponseBodyResultRelatedPatchSet {
	s.PatchSetName = &v
	return s
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) SetPatchSetNo(v string) *CreateCommentResponseBodyResultRelatedPatchSet {
	s.PatchSetNo = &v
	return s
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) SetRelatedMergeItemType(v string) *CreateCommentResponseBodyResultRelatedPatchSet {
	s.RelatedMergeItemType = &v
	return s
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) SetShortId(v string) *CreateCommentResponseBodyResultRelatedPatchSet {
	s.ShortId = &v
	return s
}

func (s *CreateCommentResponseBodyResultRelatedPatchSet) Validate() error {
	return dara.Validate(s)
}
