// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetUserPasswordResponseBody
	GetRequestId() *string
	SetResetUsersResult(v *ResetUserPasswordResponseBodyResetUsersResult) *ResetUserPasswordResponseBody
	GetResetUsersResult() *ResetUserPasswordResponseBodyResetUsersResult
}

type ResetUserPasswordResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 32D05B39-E6EE-4D7A-9FD0-762A26859D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of resetting the password of the convenience user.
	ResetUsersResult *ResetUserPasswordResponseBodyResetUsersResult `json:"ResetUsersResult,omitempty" xml:"ResetUsersResult,omitempty" type:"Struct"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetUserPasswordResponseBody) GetResetUsersResult() *ResetUserPasswordResponseBodyResetUsersResult {
	return s.ResetUsersResult
}

func (s *ResetUserPasswordResponseBody) SetRequestId(v string) *ResetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetResetUsersResult(v *ResetUserPasswordResponseBodyResetUsersResult) *ResetUserPasswordResponseBody {
	s.ResetUsersResult = v
	return s
}

func (s *ResetUserPasswordResponseBody) Validate() error {
	if s.ResetUsersResult != nil {
		if err := s.ResetUsersResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetUserPasswordResponseBodyResetUsersResult struct {
	// The information about the convenience users whose passwords failed to be reset.
	FailedUsers []*ResetUserPasswordResponseBodyResetUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	// The convenience users to which the system sent a password reset email.
	ResetUsers []*string `json:"ResetUsers,omitempty" xml:"ResetUsers,omitempty" type:"Repeated"`
}

func (s ResetUserPasswordResponseBodyResetUsersResult) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponseBodyResetUsersResult) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) GetFailedUsers() []*ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	return s.FailedUsers
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) GetResetUsers() []*string {
	return s.ResetUsers
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) SetFailedUsers(v []*ResetUserPasswordResponseBodyResetUsersResultFailedUsers) *ResetUserPasswordResponseBodyResetUsersResult {
	s.FailedUsers = v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) SetResetUsers(v []*string) *ResetUserPasswordResponseBodyResetUsersResult {
	s.ResetUsers = v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) Validate() error {
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

type ResetUserPasswordResponseBodyResetUsersResultFailedUsers struct {
	// The ID of the convenience user whose password failed to be reset.
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

func (s ResetUserPasswordResponseBodyResetUsersResultFailedUsers) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponseBodyResetUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetEndUserId(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetErrorCode(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetErrorMessage(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) Validate() error {
	return dara.Validate(s)
}
