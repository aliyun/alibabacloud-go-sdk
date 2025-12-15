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
	// The users that you want to add.
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
	// The public key of the user.
	//
	// You can add up to 20 users in a call.
	//
	// Specify one of the Password and AuthKey parameters. The AuthKey parameter takes effect only when the cluster authentication method is set to Key. Key authentication is not recommended.
	//
	// example:
	//
	// Abc****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The permission group to which the user belongs. Valid values:
	//
	// users: ordinary permissions, which are suitable for ordinary users that need only to submit and debug jobs. wheel: sudo permissions, which are suitable for administrators who need to manage clusters. In addition to submitting and debugging jobs, you can also run sudo commands to install software and restart nodes. You can add up to 20 users in a call.
	//
	// example:
	//
	// users
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The password of the user. The password must be 6 to 30 characters in length and must contain three of the following character types:
	//
	// 	- Uppercase letters
	//
	// 	- Lowercase letters
	//
	// 	- Digits
	//
	// 	- Special characters ()~!@#$%^&\\*-_+=|{}[]:;\\"/<>,.?/
	//
	// You can add up to 20 users in a call.
	//
	// Specify one of the Password and AuthKey parameters. The Password parameter takes effect only when the cluster authentication method is set to Password. Password authentication is recommended.
	//
	// example:
	//
	// 1@a2****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username. The username must be 1 to 30 characters in length. It must start with a letter and can contain digits, letters, and periods (.).
	//
	// You can add up to 20 users in a call.
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
