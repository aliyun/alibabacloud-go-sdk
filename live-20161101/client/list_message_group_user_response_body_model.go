// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMessageGroupUserResponseBody
	GetRequestId() *string
	SetResult(v *ListMessageGroupUserResponseBodyResult) *ListMessageGroupUserResponseBody
	GetResult() *ListMessageGroupUserResponseBodyResult
}

type ListMessageGroupUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *ListMessageGroupUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListMessageGroupUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageGroupUserResponseBody) GetResult() *ListMessageGroupUserResponseBodyResult {
	return s.Result
}

func (s *ListMessageGroupUserResponseBody) SetRequestId(v string) *ListMessageGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageGroupUserResponseBody) SetResult(v *ListMessageGroupUserResponseBodyResult) *ListMessageGroupUserResponseBody {
	s.Result = v
	return s
}

func (s *ListMessageGroupUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMessageGroupUserResponseBodyResult struct {
	// Indicates whether the current page is followed by another page. Valid values:
	//
	// 	- true: The current page is followed by another page.
	//
	// 	- false: The current page is not followed by another page.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The total number of users in the message group.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Details about the users.
	UserList []*ListMessageGroupUserResponseBodyResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListMessageGroupUserResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserResponseBodyResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMessageGroupUserResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *ListMessageGroupUserResponseBodyResult) GetUserList() []*ListMessageGroupUserResponseBodyResultUserList {
	return s.UserList
}

func (s *ListMessageGroupUserResponseBodyResult) SetHasMore(v bool) *ListMessageGroupUserResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListMessageGroupUserResponseBodyResult) SetTotal(v int32) *ListMessageGroupUserResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListMessageGroupUserResponseBodyResult) SetUserList(v []*ListMessageGroupUserResponseBodyResultUserList) *ListMessageGroupUserResponseBodyResult {
	s.UserList = v
	return s
}

func (s *ListMessageGroupUserResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListMessageGroupUserResponseBodyResultUserList struct {
	// The time when the user joined the message group. The value is a UTC timestamp.
	//
	// example:
	//
	// 12**45
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// de1**a0,hu**9
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMessageGroupUserResponseBodyResultUserList) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserResponseBodyResultUserList) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserResponseBodyResultUserList) GetJoinTime() *int64 {
	return s.JoinTime
}

func (s *ListMessageGroupUserResponseBodyResultUserList) GetUserId() *string {
	return s.UserId
}

func (s *ListMessageGroupUserResponseBodyResultUserList) SetJoinTime(v int64) *ListMessageGroupUserResponseBodyResultUserList {
	s.JoinTime = &v
	return s
}

func (s *ListMessageGroupUserResponseBodyResultUserList) SetUserId(v string) *ListMessageGroupUserResponseBodyResultUserList {
	s.UserId = &v
	return s
}

func (s *ListMessageGroupUserResponseBodyResultUserList) Validate() error {
	return dara.Validate(s)
}
