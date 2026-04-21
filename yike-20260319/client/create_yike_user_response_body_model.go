// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateYikeUserResponseBody
	GetRequestId() *string
	SetUserInfo(v *CreateYikeUserResponseBodyUserInfo) *CreateYikeUserResponseBody
	GetUserInfo() *CreateYikeUserResponseBodyUserInfo
}

type CreateYikeUserResponseBody struct {
	// RequestId
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserInfo  *CreateYikeUserResponseBodyUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreateYikeUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYikeUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateYikeUserResponseBody) GetUserInfo() *CreateYikeUserResponseBodyUserInfo {
	return s.UserInfo
}

func (s *CreateYikeUserResponseBody) SetRequestId(v string) *CreateYikeUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateYikeUserResponseBody) SetUserInfo(v *CreateYikeUserResponseBodyUserInfo) *CreateYikeUserResponseBody {
	s.UserInfo = v
	return s
}

func (s *CreateYikeUserResponseBody) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateYikeUserResponseBodyUserInfo struct {
	// example:
	//
	// nick
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// spaceId
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateYikeUserResponseBodyUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeUserResponseBodyUserInfo) GoString() string {
	return s.String()
}

func (s *CreateYikeUserResponseBodyUserInfo) GetNickname() *string {
	return s.Nickname
}

func (s *CreateYikeUserResponseBodyUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *CreateYikeUserResponseBodyUserInfo) GetUserName() *string {
	return s.UserName
}

func (s *CreateYikeUserResponseBodyUserInfo) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateYikeUserResponseBodyUserInfo) SetNickname(v string) *CreateYikeUserResponseBodyUserInfo {
	s.Nickname = &v
	return s
}

func (s *CreateYikeUserResponseBodyUserInfo) SetUserId(v string) *CreateYikeUserResponseBodyUserInfo {
	s.UserId = &v
	return s
}

func (s *CreateYikeUserResponseBodyUserInfo) SetUserName(v string) *CreateYikeUserResponseBodyUserInfo {
	s.UserName = &v
	return s
}

func (s *CreateYikeUserResponseBodyUserInfo) SetWorkspaceId(v string) *CreateYikeUserResponseBodyUserInfo {
	s.WorkspaceId = &v
	return s
}

func (s *CreateYikeUserResponseBodyUserInfo) Validate() error {
	return dara.Validate(s)
}
