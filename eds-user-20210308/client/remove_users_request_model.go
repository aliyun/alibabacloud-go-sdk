// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *RemoveUsersRequest
	GetBusinessChannel() *string
	SetUnbindDesktops(v bool) *RemoveUsersRequest
	GetUnbindDesktops() *bool
	SetUsers(v []*string) *RemoveUsersRequest
	GetUsers() []*string
}

type RemoveUsersRequest struct {
	// The business channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// Specifies whether to automatically revoke cloud desktop authorization before deleting convenience users. Valid values:
	//
	// - true: Asynchronously revokes cloud desktop authorization and deletes convenience users. The API returns the task acceptance result.
	//
	// - false (default): Synchronously deletes convenience users without automatically revoking cloud desktop authorization. If associated resources prevent deletion, a deletion failure is returned.
	//
	// When this parameter is set to true, the caller must also have the ecd:ModifyUserEntitlement permission. Task acceptance does not mean that the convenience users have been deleted.
	UnbindDesktops *bool `json:"UnbindDesktops,omitempty" xml:"UnbindDesktops,omitempty"`
	// The list of usernames of the convenience users to be deleted.
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

func (s *RemoveUsersRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *RemoveUsersRequest) GetUnbindDesktops() *bool {
	return s.UnbindDesktops
}

func (s *RemoveUsersRequest) GetUsers() []*string {
	return s.Users
}

func (s *RemoveUsersRequest) SetBusinessChannel(v string) *RemoveUsersRequest {
	s.BusinessChannel = &v
	return s
}

func (s *RemoveUsersRequest) SetUnbindDesktops(v bool) *RemoveUsersRequest {
	s.UnbindDesktops = &v
	return s
}

func (s *RemoveUsersRequest) SetUsers(v []*string) *RemoveUsersRequest {
	s.Users = v
	return s
}

func (s *RemoveUsersRequest) Validate() error {
	return dara.Validate(s)
}
