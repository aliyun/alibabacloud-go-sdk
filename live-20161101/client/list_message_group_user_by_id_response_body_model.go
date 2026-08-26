// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupUserByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMessageGroupUserByIdResponseBody
	GetRequestId() *string
	SetResult(v *ListMessageGroupUserByIdResponseBodyResult) *ListMessageGroupUserByIdResponseBody
	GetResult() *ListMessageGroupUserByIdResponseBodyResult
}

type ListMessageGroupUserByIdResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	Result *ListMessageGroupUserByIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListMessageGroupUserByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserByIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageGroupUserByIdResponseBody) GetResult() *ListMessageGroupUserByIdResponseBodyResult {
	return s.Result
}

func (s *ListMessageGroupUserByIdResponseBody) SetRequestId(v string) *ListMessageGroupUserByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageGroupUserByIdResponseBody) SetResult(v *ListMessageGroupUserByIdResponseBodyResult) *ListMessageGroupUserByIdResponseBody {
	s.Result = v
	return s
}

func (s *ListMessageGroupUserByIdResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessageGroupUserByIdResponseBodyResult struct {
	// Indicates whether there is a next page. Valid values:
	//
	// - true: There is a next page.
	//
	// - false: There is no next page.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// Total number of users queried.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// User list information.
	UserList []*ListMessageGroupUserByIdResponseBodyResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListMessageGroupUserByIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserByIdResponseBodyResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMessageGroupUserByIdResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *ListMessageGroupUserByIdResponseBodyResult) GetUserList() []*ListMessageGroupUserByIdResponseBodyResultUserList {
	return s.UserList
}

func (s *ListMessageGroupUserByIdResponseBodyResult) SetHasMore(v bool) *ListMessageGroupUserByIdResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResult) SetTotal(v int32) *ListMessageGroupUserByIdResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResult) SetUserList(v []*ListMessageGroupUserByIdResponseBodyResultUserList) *ListMessageGroupUserByIdResponseBodyResult {
	s.UserList = v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResult) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessageGroupUserByIdResponseBodyResultUserList struct {
	// Indicates whether the user is muted. Valid values:
	//
	// - true: Muted.
	//
	// - false: Not muted.
	//
	// example:
	//
	// true
	IsMute *bool `json:"IsMute,omitempty" xml:"IsMute,omitempty"`
	// Mute type. Valid values:
	//
	// - group: All members in the message group are muted.
	//
	// - user: Individual user is muted.
	MuteBy []*string `json:"MuteBy,omitempty" xml:"MuteBy,omitempty" type:"Repeated"`
	// Profile picture URL.
	//
	// example:
	//
	// "http://www.aliyundoc.com/xxyy.png"
	UserAvatar *string `json:"UserAvatar,omitempty" xml:"UserAvatar,omitempty"`
	// Custom user information content.
	//
	// example:
	//
	// 12e
	UserExtension *string `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	// User ID.
	//
	// example:
	//
	// ad***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// User nickname.
	//
	// example:
	//
	// xxyy
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s ListMessageGroupUserByIdResponseBodyResultUserList) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserByIdResponseBodyResultUserList) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) GetIsMute() *bool {
	return s.IsMute
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) GetMuteBy() []*string {
	return s.MuteBy
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) GetUserAvatar() *string {
	return s.UserAvatar
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) GetUserExtension() *string {
	return s.UserExtension
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) GetUserId() *string {
	return s.UserId
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) GetUserNick() *string {
	return s.UserNick
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) SetIsMute(v bool) *ListMessageGroupUserByIdResponseBodyResultUserList {
	s.IsMute = &v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) SetMuteBy(v []*string) *ListMessageGroupUserByIdResponseBodyResultUserList {
	s.MuteBy = v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) SetUserAvatar(v string) *ListMessageGroupUserByIdResponseBodyResultUserList {
	s.UserAvatar = &v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) SetUserExtension(v string) *ListMessageGroupUserByIdResponseBodyResultUserList {
	s.UserExtension = &v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) SetUserId(v string) *ListMessageGroupUserByIdResponseBodyResultUserList {
	s.UserId = &v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) SetUserNick(v string) *ListMessageGroupUserByIdResponseBodyResultUserList {
	s.UserNick = &v
	return s
}

func (s *ListMessageGroupUserByIdResponseBodyResultUserList) Validate() error {
	return dara.Validate(s)
}
