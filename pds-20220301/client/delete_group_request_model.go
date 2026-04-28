// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteGroupRequest
	GetGroupId() *string
}

type DeleteGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// g123
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteGroupRequest) SetGroupId(v string) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupRequest) Validate() error {
	return dara.Validate(s)
}
