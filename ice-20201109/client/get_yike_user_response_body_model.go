// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetYikeUserResponseBody
	GetRequestId() *string
	SetUserInfo(v *GetYikeUserResponseBodyUserInfo) *GetYikeUserResponseBody
	GetUserInfo() *GetYikeUserResponseBodyUserInfo
}

type GetYikeUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user information.
	UserInfo *GetYikeUserResponseBodyUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetYikeUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeUserResponseBody) GetUserInfo() *GetYikeUserResponseBodyUserInfo {
	return s.UserInfo
}

func (s *GetYikeUserResponseBody) SetRequestId(v string) *GetYikeUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeUserResponseBody) SetUserInfo(v *GetYikeUserResponseBodyUserInfo) *GetYikeUserResponseBody {
	s.UserInfo = v
	return s
}

func (s *GetYikeUserResponseBody) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetYikeUserResponseBodyUserInfo struct {
	// The user\\"s nickname.
	//
	// example:
	//
	// nick
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// The user name.
	//
	// example:
	//
	// name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// spaceId
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// id
	YikeUserId *string `json:"YikeUserId,omitempty" xml:"YikeUserId,omitempty"`
}

func (s GetYikeUserResponseBodyUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetYikeUserResponseBodyUserInfo) GoString() string {
	return s.String()
}

func (s *GetYikeUserResponseBodyUserInfo) GetNickname() *string {
	return s.Nickname
}

func (s *GetYikeUserResponseBodyUserInfo) GetUserName() *string {
	return s.UserName
}

func (s *GetYikeUserResponseBodyUserInfo) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetYikeUserResponseBodyUserInfo) GetYikeUserId() *string {
	return s.YikeUserId
}

func (s *GetYikeUserResponseBodyUserInfo) SetNickname(v string) *GetYikeUserResponseBodyUserInfo {
	s.Nickname = &v
	return s
}

func (s *GetYikeUserResponseBodyUserInfo) SetUserName(v string) *GetYikeUserResponseBodyUserInfo {
	s.UserName = &v
	return s
}

func (s *GetYikeUserResponseBodyUserInfo) SetWorkspaceId(v string) *GetYikeUserResponseBodyUserInfo {
	s.WorkspaceId = &v
	return s
}

func (s *GetYikeUserResponseBodyUserInfo) SetYikeUserId(v string) *GetYikeUserResponseBodyUserInfo {
	s.YikeUserId = &v
	return s
}

func (s *GetYikeUserResponseBodyUserInfo) Validate() error {
	return dara.Validate(s)
}
