// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUsers(v []*string) *RemoveUsersRequest
	GetUsers() []*string
}

type RemoveUsersRequest struct {
	// The usernames of the convenience users that you want to remove.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s RemoveUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequest) GetUsers() []*string {
	return s.Users
}

func (s *RemoveUsersRequest) SetUsers(v []*string) *RemoveUsersRequest {
	s.Users = v
	return s
}

func (s *RemoveUsersRequest) Validate() error {
	return dara.Validate(s)
}
