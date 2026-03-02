// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v string) *DeleteMonitorGroupMemberRequest
	GetContactIds() *string
}

type DeleteMonitorGroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// []
	ContactIds *string `json:"contactIds,omitempty" xml:"contactIds,omitempty"`
}

func (s DeleteMonitorGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupMemberRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *DeleteMonitorGroupMemberRequest) SetContactIds(v string) *DeleteMonitorGroupMemberRequest {
	s.ContactIds = &v
	return s
}

func (s *DeleteMonitorGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
