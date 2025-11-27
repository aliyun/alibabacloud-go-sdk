// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemoveUsersResult(v *RemoveUsersResponseBodyRemoveUsersResult) *RemoveUsersResponseBody
	GetRemoveUsersResult() *RemoveUsersResponseBodyRemoveUsersResult
	SetRequestId(v string) *RemoveUsersResponseBody
	GetRequestId() *string
}

type RemoveUsersResponseBody struct {
	// The result of removing the convenience user.
	RemoveUsersResult *RemoveUsersResponseBodyRemoveUsersResult `json:"RemoveUsersResult,omitempty" xml:"RemoveUsersResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBody) GetRemoveUsersResult() *RemoveUsersResponseBodyRemoveUsersResult {
	return s.RemoveUsersResult
}

func (s *RemoveUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUsersResponseBody) SetRemoveUsersResult(v *RemoveUsersResponseBodyRemoveUsersResult) *RemoveUsersResponseBody {
	s.RemoveUsersResult = v
	return s
}

func (s *RemoveUsersResponseBody) SetRequestId(v string) *RemoveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUsersResponseBody) Validate() error {
	if s.RemoveUsersResult != nil {
		if err := s.RemoveUsersResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveUsersResponseBodyRemoveUsersResult struct {
	// The convenience users that failed to be removed.
	FailedUsers []*RemoveUsersResponseBodyRemoveUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	// The convenience users that were removed.
	RemovedUsers []*string `json:"RemovedUsers,omitempty" xml:"RemovedUsers,omitempty" type:"Repeated"`
}

func (s RemoveUsersResponseBodyRemoveUsersResult) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersResponseBodyRemoveUsersResult) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) GetFailedUsers() []*RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	return s.FailedUsers
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) GetRemovedUsers() []*string {
	return s.RemovedUsers
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) SetFailedUsers(v []*RemoveUsersResponseBodyRemoveUsersResultFailedUsers) *RemoveUsersResponseBodyRemoveUsersResult {
	s.FailedUsers = v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) SetRemovedUsers(v []*string) *RemoveUsersResponseBodyRemoveUsersResult {
	s.RemovedUsers = v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) Validate() error {
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

type RemoveUsersResponseBodyRemoveUsersResultFailedUsers struct {
	// The ID of the convenience user that failed to be removed.
	//
	// example:
	//
	// test2
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
	// test2 is an invalid username.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s RemoveUsersResponseBodyRemoveUsersResultFailedUsers) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersResponseBodyRemoveUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetEndUserId(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetErrorCode(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetErrorMessage(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) Validate() error {
	return dara.Validate(s)
}
