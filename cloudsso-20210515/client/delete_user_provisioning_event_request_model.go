// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserProvisioningEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *DeleteUserProvisioningEventRequest
	GetDirectoryId() *string
	SetEventId(v string) *DeleteUserProvisioningEventRequest
	GetEventId() *string
	SetUserProvisioningId(v string) *DeleteUserProvisioningEventRequest
	GetUserProvisioningId() *string
}

type DeleteUserProvisioningEventRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the RAM user provisioning event.
	//
	// You can call the [ListUserProvisioningEvents](https://help.aliyun.com/document_detail/2636305.html) operation to query the value of `EventId`.
	//
	// example:
	//
	// upe-wjKyNDmZvyZOiRcJ****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the RAM user provisioning.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s DeleteUserProvisioningEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserProvisioningEventRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningEventRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteUserProvisioningEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *DeleteUserProvisioningEventRequest) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *DeleteUserProvisioningEventRequest) SetDirectoryId(v string) *DeleteUserProvisioningEventRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteUserProvisioningEventRequest) SetEventId(v string) *DeleteUserProvisioningEventRequest {
	s.EventId = &v
	return s
}

func (s *DeleteUserProvisioningEventRequest) SetUserProvisioningId(v string) *DeleteUserProvisioningEventRequest {
	s.UserProvisioningId = &v
	return s
}

func (s *DeleteUserProvisioningEventRequest) Validate() error {
	return dara.Validate(s)
}
