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
	// Request ID.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results.
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessageGroupUserResponseBodyResult struct {
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
	// Total number of message group users.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// User list.
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

type ListMessageGroupUserResponseBodyResultUserList struct {
	// UTC timestamp when the user joined the message group.
	//
	// example:
	//
	// 12**45
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// User ID.
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
