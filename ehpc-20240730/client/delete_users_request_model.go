// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteUsersRequest
	GetClusterId() *string
	SetUser(v []*DeleteUsersRequestUser) *DeleteUsersRequest
	GetUser() []*DeleteUsersRequestUser
}

type DeleteUsersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The users that you want to delete.
	//
	// This parameter is required.
	User []*DeleteUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DeleteUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteUsersRequest) GetUser() []*DeleteUsersRequestUser {
	return s.User
}

func (s *DeleteUsersRequest) SetClusterId(v string) *DeleteUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersRequest) SetUser(v []*DeleteUsersRequestUser) *DeleteUsersRequest {
	s.User = v
	return s
}

func (s *DeleteUsersRequest) Validate() error {
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

type DeleteUsersRequestUser struct {
	// The name of user N that you want to delete.
	//
	// Valid values of N: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteUsersRequestUser) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsersRequestUser) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequestUser) GetUserName() *string {
	return s.UserName
}

func (s *DeleteUsersRequestUser) SetUserName(v string) *DeleteUsersRequestUser {
	s.UserName = &v
	return s
}

func (s *DeleteUsersRequestUser) Validate() error {
	return dara.Validate(s)
}
