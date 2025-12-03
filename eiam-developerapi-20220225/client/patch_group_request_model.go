// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *PatchGroupRequest
	GetGroupName() *string
}

type PatchGroupRequest struct {
	// The group name.
	//
	// example:
	//
	// name001
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s PatchGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchGroupRequest) GoString() string {
	return s.String()
}

func (s *PatchGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *PatchGroupRequest) SetGroupName(v string) *PatchGroupRequest {
	s.GroupName = &v
	return s
}

func (s *PatchGroupRequest) Validate() error {
	return dara.Validate(s)
}
