// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSwimmingLaneGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *ListAllSwimmingLaneGroupsRequest
	GetNamespaceId() *string
}

type ListAllSwimmingLaneGroupsRequest struct {
	// The ID of a namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListAllSwimmingLaneGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLaneGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLaneGroupsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAllSwimmingLaneGroupsRequest) SetNamespaceId(v string) *ListAllSwimmingLaneGroupsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsRequest) Validate() error {
	return dara.Validate(s)
}
