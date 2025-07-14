// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DeleteSwimmingLaneGroupRequest
	GetGroupId() *int64
	SetNamespaceId(v string) *DeleteSwimmingLaneGroupRequest
	GetNamespaceId() *string
}

type DeleteSwimmingLaneGroupRequest struct {
	// example:
	//
	// 2074
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// cn-beijing:demo
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteSwimmingLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteSwimmingLaneGroupRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteSwimmingLaneGroupRequest) SetGroupId(v int64) *DeleteSwimmingLaneGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteSwimmingLaneGroupRequest) SetNamespaceId(v string) *DeleteSwimmingLaneGroupRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteSwimmingLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
