// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemCommentListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommentList(v []*GetWorkitemCommentListResponseBodyCommentList) *GetWorkitemCommentListResponseBody
	GetCommentList() []*GetWorkitemCommentListResponseBodyCommentList
	SetErrorCode(v string) *GetWorkitemCommentListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetWorkitemCommentListResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetWorkitemCommentListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetWorkitemCommentListResponseBody
	GetSuccess() *string
}

type GetWorkitemCommentListResponseBody struct {
	CommentList []*GetWorkitemCommentListResponseBodyCommentList `json:"commentList,omitempty" xml:"commentList,omitempty" type:"Repeated"`
	// example:
	//
	// Openapi.RequestError
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

func (s GetWorkitemCommentListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemCommentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemCommentListResponseBody) GetCommentList() []*GetWorkitemCommentListResponseBodyCommentList {
	return s.CommentList
}

func (s *GetWorkitemCommentListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkitemCommentListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetWorkitemCommentListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkitemCommentListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetWorkitemCommentListResponseBody) SetCommentList(v []*GetWorkitemCommentListResponseBodyCommentList) *GetWorkitemCommentListResponseBody {
	s.CommentList = v
	return s
}

func (s *GetWorkitemCommentListResponseBody) SetErrorCode(v string) *GetWorkitemCommentListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkitemCommentListResponseBody) SetErrorMsg(v string) *GetWorkitemCommentListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkitemCommentListResponseBody) SetRequestId(v string) *GetWorkitemCommentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemCommentListResponseBody) SetSuccess(v string) *GetWorkitemCommentListResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkitemCommentListResponseBody) Validate() error {
	if s.CommentList != nil {
		for _, item := range s.CommentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkitemCommentListResponseBodyCommentList struct {
	// example:
	//
	// 066961
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1667205617061
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// RICHTEXT/MARKDOWN
	FormatType *string `json:"formatType,omitempty" xml:"formatType,omitempty"`
	// example:
	//
	// 6573
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
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
	// 1013131897677566
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// 9144ef6b72d8exxxxx9e61a4d0
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	// example:
	//
	// workitem
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// example:
	//
	// 1667205617089
	TopTime *int64                                             `json:"topTime,omitempty" xml:"topTime,omitempty"`
	User    *GetWorkitemCommentListResponseBodyCommentListUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
}

func (s GetWorkitemCommentListResponseBodyCommentList) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemCommentListResponseBodyCommentList) GoString() string {
	return s.String()
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetContent() *string {
	return s.Content
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetFormatType() *string {
	return s.FormatType
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetId() *int64 {
	return s.Id
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetIsTop() *bool {
	return s.IsTop
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetTargetIdentifier() *string {
	return s.TargetIdentifier
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetTargetType() *string {
	return s.TargetType
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetTopTime() *int64 {
	return s.TopTime
}

func (s *GetWorkitemCommentListResponseBodyCommentList) GetUser() *GetWorkitemCommentListResponseBodyCommentListUser {
	return s.User
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetContent(v string) *GetWorkitemCommentListResponseBodyCommentList {
	s.Content = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetCreateTime(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.CreateTime = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetFormatType(v string) *GetWorkitemCommentListResponseBodyCommentList {
	s.FormatType = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetId(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.Id = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetIsTop(v bool) *GetWorkitemCommentListResponseBodyCommentList {
	s.IsTop = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetModifiedTime(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.ModifiedTime = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetParentId(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.ParentId = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetTargetIdentifier(v string) *GetWorkitemCommentListResponseBodyCommentList {
	s.TargetIdentifier = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetTargetType(v string) *GetWorkitemCommentListResponseBodyCommentList {
	s.TargetType = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetTopTime(v int64) *GetWorkitemCommentListResponseBodyCommentList {
	s.TopTime = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) SetUser(v *GetWorkitemCommentListResponseBodyCommentListUser) *GetWorkitemCommentListResponseBodyCommentList {
	s.User = v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentList) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkitemCommentListResponseBodyCommentListUser struct {
	Account    *string `json:"account,omitempty" xml:"account,omitempty"`
	Avatar     *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	NickName   *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	RealName   *string `json:"realName,omitempty" xml:"realName,omitempty"`
}

func (s GetWorkitemCommentListResponseBodyCommentListUser) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemCommentListResponseBodyCommentListUser) GoString() string {
	return s.String()
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) GetAccount() *string {
	return s.Account
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) GetAvatar() *string {
	return s.Avatar
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) GetNickName() *string {
	return s.NickName
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) GetRealName() *string {
	return s.RealName
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) SetAccount(v string) *GetWorkitemCommentListResponseBodyCommentListUser {
	s.Account = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) SetAvatar(v string) *GetWorkitemCommentListResponseBodyCommentListUser {
	s.Avatar = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) SetIdentifier(v string) *GetWorkitemCommentListResponseBodyCommentListUser {
	s.Identifier = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) SetNickName(v string) *GetWorkitemCommentListResponseBodyCommentListUser {
	s.NickName = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) SetRealName(v string) *GetWorkitemCommentListResponseBodyCommentListUser {
	s.RealName = &v
	return s
}

func (s *GetWorkitemCommentListResponseBodyCommentListUser) Validate() error {
	return dara.Validate(s)
}
