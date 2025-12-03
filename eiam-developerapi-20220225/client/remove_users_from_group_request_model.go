// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserIds(v []*string) *RemoveUsersFromGroupRequest
	GetUserIds() []*string
}

type RemoveUsersFromGroupRequest struct {
	// The account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [user_d6sbsuumeta4h66ec3il7yxxxx}
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RemoveUsersFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *RemoveUsersFromGroupRequest) SetUserIds(v []*string) *RemoveUsersFromGroupRequest {
	s.UserIds = v
	return s
}

func (s *RemoveUsersFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
