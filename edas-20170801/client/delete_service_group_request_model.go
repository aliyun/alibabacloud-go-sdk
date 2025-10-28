// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteServiceGroupRequest
	GetGroupId() *string
}

type DeleteServiceGroupRequest struct {
	// The ID of the service group that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 789d9cda-74b1-****-****-05e21a0a7661
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteServiceGroupRequest) SetGroupId(v string) *DeleteServiceGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteServiceGroupRequest) Validate() error {
	return dara.Validate(s)
}
