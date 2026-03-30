// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserInRecycleBinResponseBody
	GetRequestId() *string
	SetUser(v *GetUserInRecycleBinResponseBodyUser) *GetUserInRecycleBinResponseBody
	GetUser() *GetUserInRecycleBinResponseBodyUser
}

type GetUserInRecycleBinResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *GetUserInRecycleBinResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserInRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserInRecycleBinResponseBody) GetUser() *GetUserInRecycleBinResponseBodyUser {
	return s.User
}

func (s *GetUserInRecycleBinResponseBody) SetRequestId(v string) *GetUserInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInRecycleBinResponseBody) SetUser(v *GetUserInRecycleBinResponseBodyUser) *GetUserInRecycleBinResponseBody {
	s.User = v
	return s
}

func (s *GetUserInRecycleBinResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserInRecycleBinResponseBodyUser struct {
	// The time when the RAM user was created.
	//
	// example:
	//
	// 2020-10-11T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the RAM user will be permanently deleted from the recycle bin.
	//
	// example:
	//
	// 2020-11-12T09:12:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the RAM user was deleted and moved to the recycle bin.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	RecycleDate *string `json:"RecycleDate,omitempty" xml:"RecycleDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserInRecycleBinResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetUserInRecycleBinResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserInRecycleBinResponseBodyUser) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetUserInRecycleBinResponseBodyUser) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *GetUserInRecycleBinResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserInRecycleBinResponseBodyUser) GetRecycleDate() *string {
	return s.RecycleDate
}

func (s *GetUserInRecycleBinResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *GetUserInRecycleBinResponseBodyUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetUserInRecycleBinResponseBodyUser) SetCreateDate(v string) *GetUserInRecycleBinResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetDeleteDate(v string) *GetUserInRecycleBinResponseBodyUser {
	s.DeleteDate = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetDisplayName(v string) *GetUserInRecycleBinResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetRecycleDate(v string) *GetUserInRecycleBinResponseBodyUser {
	s.RecycleDate = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetUserId(v string) *GetUserInRecycleBinResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetUserPrincipalName(v string) *GetUserInRecycleBinResponseBodyUser {
	s.UserPrincipalName = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) Validate() error {
	return dara.Validate(s)
}
