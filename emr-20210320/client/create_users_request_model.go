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
	SetRegionId(v string) *CreateUsersRequest
	GetRegionId() *string
	SetUsers(v []*CreateUsersRequestUsers) *CreateUsersRequest
	GetUsers() []*CreateUsersRequestUsers
}

type CreateUsersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of users. The number of users must be from 0 to 10.
	//
	// This parameter is required.
	Users []*CreateUsersRequestUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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

func (s *CreateUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateUsersRequest) GetUsers() []*CreateUsersRequestUsers {
	return s.Users
}

func (s *CreateUsersRequest) SetClusterId(v string) *CreateUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateUsersRequest) SetRegionId(v string) *CreateUsersRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUsersRequest) SetUsers(v []*CreateUsersRequestUsers) *CreateUsersRequest {
	s.Users = v
	return s
}

func (s *CreateUsersRequest) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateUsersRequestUsers struct {
	// The password of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// *Ab123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// xi
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUsersRequestUsers) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersRequestUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersRequestUsers) GetPassword() *string {
	return s.Password
}

func (s *CreateUsersRequestUsers) GetUserName() *string {
	return s.UserName
}

func (s *CreateUsersRequestUsers) SetPassword(v string) *CreateUsersRequestUsers {
	s.Password = &v
	return s
}

func (s *CreateUsersRequestUsers) SetUserName(v string) *CreateUsersRequestUsers {
	s.UserName = &v
	return s
}

func (s *CreateUsersRequestUsers) Validate() error {
	return dara.Validate(s)
}
