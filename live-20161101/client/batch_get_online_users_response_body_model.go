// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetOnlineUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchGetOnlineUsersResponseBody
	GetRequestId() *string
	SetResult(v *BatchGetOnlineUsersResponseBodyResult) *BatchGetOnlineUsersResponseBody
	GetResult() *BatchGetOnlineUsersResponseBodyResult
}

type BatchGetOnlineUsersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Result *BatchGetOnlineUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s BatchGetOnlineUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetOnlineUsersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetOnlineUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetOnlineUsersResponseBody) GetResult() *BatchGetOnlineUsersResponseBodyResult {
	return s.Result
}

func (s *BatchGetOnlineUsersResponseBody) SetRequestId(v string) *BatchGetOnlineUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetOnlineUsersResponseBody) SetResult(v *BatchGetOnlineUsersResponseBodyResult) *BatchGetOnlineUsersResponseBody {
	s.Result = v
	return s
}

func (s *BatchGetOnlineUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchGetOnlineUsersResponseBodyResult struct {
	// The information about users.
	OnlineUsers []*BatchGetOnlineUsersResponseBodyResultOnlineUsers `json:"OnlineUsers,omitempty" xml:"OnlineUsers,omitempty" type:"Repeated"`
}

func (s BatchGetOnlineUsersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetOnlineUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchGetOnlineUsersResponseBodyResult) GetOnlineUsers() []*BatchGetOnlineUsersResponseBodyResultOnlineUsers {
	return s.OnlineUsers
}

func (s *BatchGetOnlineUsersResponseBodyResult) SetOnlineUsers(v []*BatchGetOnlineUsersResponseBodyResultOnlineUsers) *BatchGetOnlineUsersResponseBodyResult {
	s.OnlineUsers = v
	return s
}

func (s *BatchGetOnlineUsersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type BatchGetOnlineUsersResponseBodyResultOnlineUsers struct {
	// The time when the user joined the group. The value is a UTC timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 12**45
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// Indicates whether the user is online. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// de1**a0
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchGetOnlineUsersResponseBodyResultOnlineUsers) String() string {
	return dara.Prettify(s)
}

func (s BatchGetOnlineUsersResponseBodyResultOnlineUsers) GoString() string {
	return s.String()
}

func (s *BatchGetOnlineUsersResponseBodyResultOnlineUsers) GetJoinTime() *int64 {
	return s.JoinTime
}

func (s *BatchGetOnlineUsersResponseBodyResultOnlineUsers) GetOnline() *bool {
	return s.Online
}

func (s *BatchGetOnlineUsersResponseBodyResultOnlineUsers) GetUserId() *string {
	return s.UserId
}

func (s *BatchGetOnlineUsersResponseBodyResultOnlineUsers) SetJoinTime(v int64) *BatchGetOnlineUsersResponseBodyResultOnlineUsers {
	s.JoinTime = &v
	return s
}

func (s *BatchGetOnlineUsersResponseBodyResultOnlineUsers) SetOnline(v bool) *BatchGetOnlineUsersResponseBodyResultOnlineUsers {
	s.Online = &v
	return s
}

func (s *BatchGetOnlineUsersResponseBodyResultOnlineUsers) SetUserId(v string) *BatchGetOnlineUsersResponseBodyResultOnlineUsers {
	s.UserId = &v
	return s
}

func (s *BatchGetOnlineUsersResponseBodyResultOnlineUsers) Validate() error {
	return dara.Validate(s)
}
