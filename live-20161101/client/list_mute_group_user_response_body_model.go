// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMuteGroupUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMuteGroupUserResponseBody
	GetRequestId() *string
	SetResult(v *ListMuteGroupUserResponseBodyResult) *ListMuteGroupUserResponseBody
	GetResult() *ListMuteGroupUserResponseBodyResult
}

type ListMuteGroupUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *ListMuteGroupUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListMuteGroupUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMuteGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListMuteGroupUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMuteGroupUserResponseBody) GetResult() *ListMuteGroupUserResponseBodyResult {
	return s.Result
}

func (s *ListMuteGroupUserResponseBody) SetRequestId(v string) *ListMuteGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMuteGroupUserResponseBody) SetResult(v *ListMuteGroupUserResponseBodyResult) *ListMuteGroupUserResponseBody {
	s.Result = v
	return s
}

func (s *ListMuteGroupUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMuteGroupUserResponseBodyResult struct {
	// Indicates whether the current page is followed by another page. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The total number of muted members.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The list of muted users.
	UserList []*ListMuteGroupUserResponseBodyResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListMuteGroupUserResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMuteGroupUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMuteGroupUserResponseBodyResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMuteGroupUserResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *ListMuteGroupUserResponseBodyResult) GetUserList() []*ListMuteGroupUserResponseBodyResultUserList {
	return s.UserList
}

func (s *ListMuteGroupUserResponseBodyResult) SetHasMore(v bool) *ListMuteGroupUserResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListMuteGroupUserResponseBodyResult) SetTotal(v int32) *ListMuteGroupUserResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListMuteGroupUserResponseBodyResult) SetUserList(v []*ListMuteGroupUserResponseBodyResultUserList) *ListMuteGroupUserResponseBodyResult {
	s.UserList = v
	return s
}

func (s *ListMuteGroupUserResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListMuteGroupUserResponseBodyResultUserList struct {
	// The ID of the muted user.
	//
	// example:
	//
	// 1sd***,yu***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMuteGroupUserResponseBodyResultUserList) String() string {
	return dara.Prettify(s)
}

func (s ListMuteGroupUserResponseBodyResultUserList) GoString() string {
	return s.String()
}

func (s *ListMuteGroupUserResponseBodyResultUserList) GetUserId() *string {
	return s.UserId
}

func (s *ListMuteGroupUserResponseBodyResultUserList) SetUserId(v string) *ListMuteGroupUserResponseBodyResultUserList {
	s.UserId = &v
	return s
}

func (s *ListMuteGroupUserResponseBodyResultUserList) Validate() error {
	return dara.Validate(s)
}
