// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserIds(v []*string) *AddUsersToGroupRequest
	GetUserIds() []*string
}

type AddUsersToGroupRequest struct {
	// The account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [user_d6sbsuumeta4h66ec3il7yxxxx}
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddUsersToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *AddUsersToGroupRequest) SetUserIds(v []*string) *AddUsersToGroupRequest {
	s.UserIds = v
	return s
}

func (s *AddUsersToGroupRequest) Validate() error {
	return dara.Validate(s)
}
