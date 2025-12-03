// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateGroupResponseBody
	GetGroupId() *string
}

type CreateGroupResponseBody struct {
	// The group ID.
	//
	// example:
	//
	// group_wovwffm62xifdziem7an7xxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupResponseBody) SetGroupId(v string) *CreateGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
