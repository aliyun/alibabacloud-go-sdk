// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllSucceed(v bool) *CreateUsersResponseBody
	GetAllSucceed() *bool
	SetCreateResult(v *CreateUsersResponseBodyCreateResult) *CreateUsersResponseBody
	GetCreateResult() *CreateUsersResponseBodyCreateResult
	SetRequestId(v string) *CreateUsersResponseBody
	GetRequestId() *string
}

type CreateUsersResponseBody struct {
	AllSucceed *bool `json:"AllSucceed,omitempty" xml:"AllSucceed,omitempty"`
	// The result of user creation.
	CreateResult *CreateUsersResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBody) GetAllSucceed() *bool {
	return s.AllSucceed
}

func (s *CreateUsersResponseBody) GetCreateResult() *CreateUsersResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUsersResponseBody) SetAllSucceed(v bool) *CreateUsersResponseBody {
	s.AllSucceed = &v
	return s
}

func (s *CreateUsersResponseBody) SetCreateResult(v *CreateUsersResponseBodyCreateResult) *CreateUsersResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateUsersResponseBody) SetRequestId(v string) *CreateUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUsersResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUsersResponseBodyCreateResult struct {
	// Details of the created convenience users.
	CreatedUsers []*CreateUsersResponseBodyCreateResultCreatedUsers `json:"CreatedUsers,omitempty" xml:"CreatedUsers,omitempty" type:"Repeated"`
	// Details of the convenience users that failed to be created.
	FailedUsers []*CreateUsersResponseBodyCreateResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
}

func (s CreateUsersResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResult) GetCreatedUsers() []*CreateUsersResponseBodyCreateResultCreatedUsers {
	return s.CreatedUsers
}

func (s *CreateUsersResponseBodyCreateResult) GetFailedUsers() []*CreateUsersResponseBodyCreateResultFailedUsers {
	return s.FailedUsers
}

func (s *CreateUsersResponseBodyCreateResult) SetCreatedUsers(v []*CreateUsersResponseBodyCreateResultCreatedUsers) *CreateUsersResponseBodyCreateResult {
	s.CreatedUsers = v
	return s
}

func (s *CreateUsersResponseBodyCreateResult) SetFailedUsers(v []*CreateUsersResponseBodyCreateResultFailedUsers) *CreateUsersResponseBodyCreateResult {
	s.FailedUsers = v
	return s
}

func (s *CreateUsersResponseBodyCreateResult) Validate() error {
	if s.CreatedUsers != nil {
		for _, item := range s.CreatedUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateUsersResponseBodyCreateResultCreatedUsers struct {
	// The email address of the end user.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the end user.
	//
	// example:
	//
	// test1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The mobile number of the end user.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The display name of the end user.
	//
	// example:
	//
	// Bean
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks of the end user.
	//
	// example:
	//
	// remark1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateUsersResponseBodyCreateResultCreatedUsers) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResultCreatedUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) GetEmail() *string {
	return s.Email
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) GetPhone() *string {
	return s.Phone
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) GetRealNickName() *string {
	return s.RealNickName
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) GetRemark() *string {
	return s.Remark
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetEmail(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetEndUserId(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetPhone(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Phone = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetRealNickName(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.RealNickName = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetRemark(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Remark = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) Validate() error {
	return dara.Validate(s)
}

type CreateUsersResponseBodyCreateResultFailedUsers struct {
	// The email address of the end user.
	//
	// example:
	//
	// username2@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the end user.
	//
	// example:
	//
	// test2
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// ExistedEndUserId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The username test is used by another user.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The mobile number of the end user.
	//
	// example:
	//
	// 1390000****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s CreateUsersResponseBodyCreateResultFailedUsers) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResultFailedUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) GetEmail() *string {
	return s.Email
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) GetPhone() *string {
	return s.Phone
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetEmail(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetEndUserId(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetErrorCode(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetErrorMessage(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetPhone(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.Phone = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) Validate() error {
	return dara.Validate(s)
}
