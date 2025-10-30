// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGroupTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *ListSwimmingLaneGroupTagsRequest
	GetGroupId() *int64
	SetNamespaceId(v string) *ListSwimmingLaneGroupTagsRequest
	GetNamespaceId() *string
}

type ListSwimmingLaneGroupTagsRequest struct {
	// The ID of the lane group.
	//
	// example:
	//
	// 2074
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of a namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListSwimmingLaneGroupTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupTagsRequest) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupTagsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListSwimmingLaneGroupTagsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListSwimmingLaneGroupTagsRequest) SetGroupId(v int64) *ListSwimmingLaneGroupTagsRequest {
	s.GroupId = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsRequest) SetNamespaceId(v string) *ListSwimmingLaneGroupTagsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsRequest) Validate() error {
	return dara.Validate(s)
}
