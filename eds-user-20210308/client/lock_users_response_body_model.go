// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLockUsersResult(v *LockUsersResponseBodyLockUsersResult) *LockUsersResponseBody
	GetLockUsersResult() *LockUsersResponseBodyLockUsersResult
	SetRequestId(v string) *LockUsersResponseBody
	GetRequestId() *string
}

type LockUsersResponseBody struct {
	// The result of the LockUsers operation.
	LockUsersResult *LockUsersResponseBodyLockUsersResult `json:"LockUsersResult,omitempty" xml:"LockUsersResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBody) GetLockUsersResult() *LockUsersResponseBodyLockUsersResult {
	return s.LockUsersResult
}

func (s *LockUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockUsersResponseBody) SetLockUsersResult(v *LockUsersResponseBodyLockUsersResult) *LockUsersResponseBody {
	s.LockUsersResult = v
	return s
}

func (s *LockUsersResponseBody) SetRequestId(v string) *LockUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockUsersResponseBody) Validate() error {
	if s.LockUsersResult != nil {
		if err := s.LockUsersResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LockUsersResponseBodyLockUsersResult struct {
	// A list of convenience accounts that failed to lock.
	FailedUsers []*LockUsersResponseBodyLockUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	// A list of successfully locked convenience accounts.
	LockedUsers []*string `json:"LockedUsers,omitempty" xml:"LockedUsers,omitempty" type:"Repeated"`
}

func (s LockUsersResponseBodyLockUsersResult) String() string {
	return dara.Prettify(s)
}

func (s LockUsersResponseBodyLockUsersResult) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyLockUsersResult) GetFailedUsers() []*LockUsersResponseBodyLockUsersResultFailedUsers {
	return s.FailedUsers
}

func (s *LockUsersResponseBodyLockUsersResult) GetLockedUsers() []*string {
	return s.LockedUsers
}

func (s *LockUsersResponseBodyLockUsersResult) SetFailedUsers(v []*LockUsersResponseBodyLockUsersResultFailedUsers) *LockUsersResponseBodyLockUsersResult {
	s.FailedUsers = v
	return s
}

func (s *LockUsersResponseBodyLockUsersResult) SetLockedUsers(v []*string) *LockUsersResponseBodyLockUsersResult {
	s.LockedUsers = v
	return s
}

func (s *LockUsersResponseBodyLockUsersResult) Validate() error {
	if s.FailedUsers != nil {
		for _, item := range s.FailedUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LockUsersResponseBodyLockUsersResultFailedUsers struct {
	// The username of the convenience account that failed to lock.
	//
	// example:
	//
	// alice
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

func (s LockUsersResponseBodyLockUsersResultFailedUsers) String() string {
	return dara.Prettify(s)
}

func (s LockUsersResponseBodyLockUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetEndUserId(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetErrorCode(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetErrorMessage(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) Validate() error {
	return dara.Validate(s)
}
