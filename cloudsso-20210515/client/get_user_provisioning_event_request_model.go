// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetUserProvisioningEventRequest
	GetDirectoryId() *string
	SetEventId(v string) *GetUserProvisioningEventRequest
	GetEventId() *string
}

type GetUserProvisioningEventRequest struct {
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
}

func (s GetUserProvisioningEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningEventRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningEventRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *GetUserProvisioningEventRequest) SetDirectoryId(v string) *GetUserProvisioningEventRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningEventRequest) SetEventId(v string) *GetUserProvisioningEventRequest {
	s.EventId = &v
	return s
}

func (s *GetUserProvisioningEventRequest) Validate() error {
	return dara.Validate(s)
}
