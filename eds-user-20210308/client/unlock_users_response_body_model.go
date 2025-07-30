// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnlockUsersResponseBody
	GetRequestId() *string
	SetUnlockUsersResult(v *UnlockUsersResponseBodyUnlockUsersResult) *UnlockUsersResponseBody
	GetUnlockUsersResult() *UnlockUsersResponseBodyUnlockUsersResult
}

type UnlockUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 023F4EC4-3602-4A3E-A514-4970847D59DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of unlocking the convenience user.
	UnlockUsersResult *UnlockUsersResponseBodyUnlockUsersResult `json:"UnlockUsersResult,omitempty" xml:"UnlockUsersResult,omitempty" type:"Struct"`
}

func (s UnlockUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockUsersResponseBody) GetUnlockUsersResult() *UnlockUsersResponseBodyUnlockUsersResult {
	return s.UnlockUsersResult
}

func (s *UnlockUsersResponseBody) SetRequestId(v string) *UnlockUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockUsersResponseBody) SetUnlockUsersResult(v *UnlockUsersResponseBodyUnlockUsersResult) *UnlockUsersResponseBody {
	s.UnlockUsersResult = v
	return s
}

func (s *UnlockUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type UnlockUsersResponseBodyUnlockUsersResult struct {
	// The convenience users that failed to be unlocked.
	FailedUsers []*UnlockUsersResponseBodyUnlockUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	// The convenience users that were unlocked.
	UnlockedUsers []*string `json:"UnlockedUsers,omitempty" xml:"UnlockedUsers,omitempty" type:"Repeated"`
}

func (s UnlockUsersResponseBodyUnlockUsersResult) String() string {
	return dara.Prettify(s)
}

func (s UnlockUsersResponseBodyUnlockUsersResult) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) GetFailedUsers() []*UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	return s.FailedUsers
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) GetUnlockedUsers() []*string {
	return s.UnlockedUsers
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) SetFailedUsers(v []*UnlockUsersResponseBodyUnlockUsersResultFailedUsers) *UnlockUsersResponseBodyUnlockUsersResult {
	s.FailedUsers = v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) SetUnlockedUsers(v []*string) *UnlockUsersResponseBodyUnlockUsersResult {
	s.UnlockedUsers = v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) Validate() error {
	return dara.Validate(s)
}

type UnlockUsersResponseBodyUnlockUsersResultFailedUsers struct {
	// The ID of the convenience user that failed to be unlocked.
	//
	// example:
	//
	// test123
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidUsername
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// test123 is an invalid username.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s UnlockUsersResponseBodyUnlockUsersResultFailedUsers) String() string {
	return dara.Prettify(s)
}

func (s UnlockUsersResponseBodyUnlockUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetEndUserId(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetErrorCode(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetErrorMessage(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) Validate() error {
	return dara.Validate(s)
}
