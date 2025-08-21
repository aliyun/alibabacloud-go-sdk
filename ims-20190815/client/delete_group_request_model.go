// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteGroupRequest
	GetGroupName() *string
}

type DeleteGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteGroupRequest) SetGroupName(v string) *DeleteGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteGroupRequest) Validate() error {
	return dara.Validate(s)
}
