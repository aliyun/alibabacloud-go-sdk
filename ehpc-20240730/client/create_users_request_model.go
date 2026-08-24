// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateUsersRequest
	GetClusterId() *string
	SetUser(v []*CreateUsersRequestUser) *CreateUsersRequest
	GetUser() []*CreateUsersRequestUser
}

type CreateUsersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of users.
	User []*CreateUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s CreateUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersRequest) GoString() string {
	return s.String()
}

func (s *CreateUsersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateUsersRequest) GetUser() []*CreateUsersRequestUser {
	return s.User
}

func (s *CreateUsersRequest) SetClusterId(v string) *CreateUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateUsersRequest) SetUser(v []*CreateUsersRequestUser) *CreateUsersRequest {
	s.User = v
	return s
}

func (s *CreateUsersRequest) Validate() error {
	if s.User != nil {
		for _, item := range s.User {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateUsersRequestUser struct {
	// The public key of the Nth user to add.
	//
	// Valid values of N: 1 to 20.
	//
	// This parameter is mutually exclusive with the Password parameter. This parameter takes effect when the cluster authentication method is set to key (not recommended).
	//
	// example:
	//
	// Abc****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The user group of the Nth user to add. Valid values:
	//
	// - users: ordinary permission group. This group is suitable for regular users who only need to commit and debug jobs.
	//
	// - wheel: sudo permission group. This group is suitable for administrators who need to perform cluster management. In addition to committing and debugging jobs, users in this group can execute sudo commands to install software, restart nodes, and perform other operations.
	//
	// Valid values of N: 1 to 20.
	//
	// example:
	//
	// users
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The password of the Nth user to add. The password must be 8 to 30 characters in length and contain at least three of the following four character types:
	//
	// - Uppercase letters
	//
	// - Lowercase letters
	//
	// - Digits
	//
	// - Special characters: ()~!@#$%^&*-_+=|{}[]:;\\"/<>,.?/
	//
	// Valid values of N: 1 to 20.
	//
	// This parameter is mutually exclusive with the AuthKey parameter. This parameter takes effect when the cluster authentication method is set to password (recommended).
	//
	// example:
	//
	// 1@a2****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username of the Nth user to add. The username must be 1 to 30 characters in length, start with a letter, and can contain digits and special characters (.).
	//
	// Valid values of N: 1 to 20.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUsersRequestUser) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersRequestUser) GoString() string {
	return s.String()
}

func (s *CreateUsersRequestUser) GetAuthKey() *string {
	return s.AuthKey
}

func (s *CreateUsersRequestUser) GetGroup() *string {
	return s.Group
}

func (s *CreateUsersRequestUser) GetPassword() *string {
	return s.Password
}

func (s *CreateUsersRequestUser) GetUserName() *string {
	return s.UserName
}

func (s *CreateUsersRequestUser) SetAuthKey(v string) *CreateUsersRequestUser {
	s.AuthKey = &v
	return s
}

func (s *CreateUsersRequestUser) SetGroup(v string) *CreateUsersRequestUser {
	s.Group = &v
	return s
}

func (s *CreateUsersRequestUser) SetPassword(v string) *CreateUsersRequestUser {
	s.Password = &v
	return s
}

func (s *CreateUsersRequestUser) SetUserName(v string) *CreateUsersRequestUser {
	s.UserName = &v
	return s
}

func (s *CreateUsersRequestUser) Validate() error {
	return dara.Validate(s)
}
