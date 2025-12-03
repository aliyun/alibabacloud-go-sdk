// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemCommentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v *CreateWorkitemCommentResponseBodyComment) *CreateWorkitemCommentResponseBody
	GetComment() *CreateWorkitemCommentResponseBodyComment
	SetErrorCode(v string) *CreateWorkitemCommentResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateWorkitemCommentResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateWorkitemCommentResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateWorkitemCommentResponseBody
	GetSuccess() *string
}

type CreateWorkitemCommentResponseBody struct {
	Comment *CreateWorkitemCommentResponseBodyComment `json:"Comment,omitempty" xml:"Comment,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkitemCommentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemCommentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentResponseBody) GetComment() *CreateWorkitemCommentResponseBodyComment {
	return s.Comment
}

func (s *CreateWorkitemCommentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateWorkitemCommentResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateWorkitemCommentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkitemCommentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateWorkitemCommentResponseBody) SetComment(v *CreateWorkitemCommentResponseBodyComment) *CreateWorkitemCommentResponseBody {
	s.Comment = v
	return s
}

func (s *CreateWorkitemCommentResponseBody) SetErrorCode(v string) *CreateWorkitemCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemCommentResponseBody) SetErrorMsg(v string) *CreateWorkitemCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWorkitemCommentResponseBody) SetRequestId(v string) *CreateWorkitemCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemCommentResponseBody) SetSuccess(v string) *CreateWorkitemCommentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkitemCommentResponseBody) Validate() error {
	if s.Comment != nil {
		if err := s.Comment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkitemCommentResponseBodyComment struct {
	// example:
	//
	// 26842
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1667202662000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// RICHTEXT/MARKDOWN
	FormatType *string `json:"formatType,omitempty" xml:"formatType,omitempty"`
	// example:
	//
	// true/false
	IsTop *bool `json:"isTop,omitempty" xml:"isTop,omitempty"`
	// example:
	//
	// 1637995553000
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// 26842
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// 79d523806bdxxxxx95612a5154
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	// example:
	//
	// workitem
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// example:
	//
	// 1667202662000
	TopTime *int64                                        `json:"topTime,omitempty" xml:"topTime,omitempty"`
	User    *CreateWorkitemCommentResponseBodyCommentUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
}

func (s CreateWorkitemCommentResponseBodyComment) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemCommentResponseBodyComment) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentResponseBodyComment) GetId() *int64 {
	return s.Id
}

func (s *CreateWorkitemCommentResponseBodyComment) GetContent() *string {
	return s.Content
}

func (s *CreateWorkitemCommentResponseBodyComment) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateWorkitemCommentResponseBodyComment) GetFormatType() *string {
	return s.FormatType
}

func (s *CreateWorkitemCommentResponseBodyComment) GetIsTop() *bool {
	return s.IsTop
}

func (s *CreateWorkitemCommentResponseBodyComment) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *CreateWorkitemCommentResponseBodyComment) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateWorkitemCommentResponseBodyComment) GetTargetIdentifier() *string {
	return s.TargetIdentifier
}

func (s *CreateWorkitemCommentResponseBodyComment) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateWorkitemCommentResponseBodyComment) GetTopTime() *int64 {
	return s.TopTime
}

func (s *CreateWorkitemCommentResponseBodyComment) GetUser() *CreateWorkitemCommentResponseBodyCommentUser {
	return s.User
}

func (s *CreateWorkitemCommentResponseBodyComment) SetId(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.Id = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetContent(v string) *CreateWorkitemCommentResponseBodyComment {
	s.Content = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetCreateTime(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.CreateTime = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetFormatType(v string) *CreateWorkitemCommentResponseBodyComment {
	s.FormatType = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetIsTop(v bool) *CreateWorkitemCommentResponseBodyComment {
	s.IsTop = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetModifiedTime(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.ModifiedTime = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetParentId(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.ParentId = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetTargetIdentifier(v string) *CreateWorkitemCommentResponseBodyComment {
	s.TargetIdentifier = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetTargetType(v string) *CreateWorkitemCommentResponseBodyComment {
	s.TargetType = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetTopTime(v int64) *CreateWorkitemCommentResponseBodyComment {
	s.TopTime = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) SetUser(v *CreateWorkitemCommentResponseBodyCommentUser) *CreateWorkitemCommentResponseBodyComment {
	s.User = v
	return s
}

func (s *CreateWorkitemCommentResponseBodyComment) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkitemCommentResponseBodyCommentUser struct {
	// example:
	//
	// 13164xxxxx41068
	Account *string `json:"account,omitempty" xml:"account,omitempty"`
	// example:
	//
	// https://xxxxx/thumbnail/112m7150e256dbba0d6456bafbb65c81f90d/w/200/h/200
	Avatar      *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1e9903d8bxxxxxxe9df9286ef5
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	NickName   *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	RealName   *string `json:"realName,omitempty" xml:"realName,omitempty"`
	// example:
	//
	// workitem
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s CreateWorkitemCommentResponseBodyCommentUser) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemCommentResponseBodyCommentUser) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) GetAccount() *string {
	return s.Account
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) GetAvatar() *string {
	return s.Avatar
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) GetNickName() *string {
	return s.NickName
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) GetRealName() *string {
	return s.RealName
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetAccount(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.Account = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetAvatar(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.Avatar = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetDisplayName(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.DisplayName = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetIdentifier(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetNickName(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.NickName = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetRealName(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.RealName = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) SetTargetType(v string) *CreateWorkitemCommentResponseBodyCommentUser {
	s.TargetType = &v
	return s
}

func (s *CreateWorkitemCommentResponseBodyCommentUser) Validate() error {
	return dara.Validate(s)
}
