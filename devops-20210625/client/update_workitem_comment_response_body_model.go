// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkitemCommentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v *UpdateWorkitemCommentResponseBodyComment) *UpdateWorkitemCommentResponseBody
	GetComment() *UpdateWorkitemCommentResponseBodyComment
	SetErrorCode(v string) *UpdateWorkitemCommentResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateWorkitemCommentResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *UpdateWorkitemCommentResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateWorkitemCommentResponseBody
	GetSuccess() *string
}

type UpdateWorkitemCommentResponseBody struct {
	Comment *UpdateWorkitemCommentResponseBodyComment `json:"comment,omitempty" xml:"comment,omitempty" type:"Struct"`
	// example:
	//
	// InvalidParam.NotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateWorkitemCommentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemCommentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentResponseBody) GetComment() *UpdateWorkitemCommentResponseBodyComment {
	return s.Comment
}

func (s *UpdateWorkitemCommentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateWorkitemCommentResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateWorkitemCommentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkitemCommentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateWorkitemCommentResponseBody) SetComment(v *UpdateWorkitemCommentResponseBodyComment) *UpdateWorkitemCommentResponseBody {
	s.Comment = v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) SetErrorCode(v string) *UpdateWorkitemCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) SetErrorMsg(v string) *UpdateWorkitemCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) SetRequestId(v string) *UpdateWorkitemCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) SetSuccess(v string) *UpdateWorkitemCommentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBody) Validate() error {
	if s.Comment != nil {
		if err := s.Comment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkitemCommentResponseBodyComment struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1667205617061
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// MARKDOWN/RICHTEXT
	FormatType *string `json:"formatType,omitempty" xml:"formatType,omitempty"`
	// example:
	//
	// 1964584
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true/false
	IsTop *bool `json:"isTop,omitempty" xml:"isTop,omitempty"`
	// example:
	//
	// 1646323200000
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// 2684432
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// deafe5f33xxxxx6a259d8dafd
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	// example:
	//
	// workitem
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// example:
	//
	// 1667205617089
	TopTime *int64                                        `json:"topTime,omitempty" xml:"topTime,omitempty"`
	User    *UpdateWorkitemCommentResponseBodyCommentUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
}

func (s UpdateWorkitemCommentResponseBodyComment) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemCommentResponseBodyComment) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetContent() *string {
	return s.Content
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetFormatType() *string {
	return s.FormatType
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetId() *string {
	return s.Id
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetIsTop() *bool {
	return s.IsTop
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetTargetIdentifier() *string {
	return s.TargetIdentifier
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetTopTime() *int64 {
	return s.TopTime
}

func (s *UpdateWorkitemCommentResponseBodyComment) GetUser() *UpdateWorkitemCommentResponseBodyCommentUser {
	return s.User
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetContent(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.Content = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetCreateTime(v int64) *UpdateWorkitemCommentResponseBodyComment {
	s.CreateTime = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetFormatType(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.FormatType = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetId(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.Id = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetIsTop(v bool) *UpdateWorkitemCommentResponseBodyComment {
	s.IsTop = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetModifiedTime(v int64) *UpdateWorkitemCommentResponseBodyComment {
	s.ModifiedTime = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetParentId(v int64) *UpdateWorkitemCommentResponseBodyComment {
	s.ParentId = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetTargetIdentifier(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.TargetIdentifier = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetTargetType(v string) *UpdateWorkitemCommentResponseBodyComment {
	s.TargetType = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetTopTime(v int64) *UpdateWorkitemCommentResponseBodyComment {
	s.TopTime = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) SetUser(v *UpdateWorkitemCommentResponseBodyCommentUser) *UpdateWorkitemCommentResponseBodyComment {
	s.User = v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyComment) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkitemCommentResponseBodyCommentUser struct {
	// example:
	//
	// 1316458xxxxx41068
	Account *string `json:"account,omitempty" xml:"account,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112m7150e256dbba0d6456bafbb65c81f90d/w/200/h/200
	Avatar      *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 9144ef6b72d8exxxxx9e61a4d0
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	NickName   *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	RealName   *string `json:"realName,omitempty" xml:"realName,omitempty"`
}

func (s UpdateWorkitemCommentResponseBodyCommentUser) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemCommentResponseBodyCommentUser) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) GetAccount() *string {
	return s.Account
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) GetNickName() *string {
	return s.NickName
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) GetRealName() *string {
	return s.RealName
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetAccount(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.Account = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetAvatar(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.Avatar = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetDisplayName(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.DisplayName = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetIdentifier(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.Identifier = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetNickName(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.NickName = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) SetRealName(v string) *UpdateWorkitemCommentResponseBodyCommentUser {
	s.RealName = &v
	return s
}

func (s *UpdateWorkitemCommentResponseBodyCommentUser) Validate() error {
	return dara.Validate(s)
}
