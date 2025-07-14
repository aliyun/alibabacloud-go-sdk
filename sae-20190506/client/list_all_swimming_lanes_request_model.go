// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSwimmingLanesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *ListAllSwimmingLanesRequest
	GetGroupId() *int64
	SetNamespaceId(v string) *ListAllSwimmingLanesRequest
	GetNamespaceId() *string
}

type ListAllSwimmingLanesRequest struct {
	// example:
	//
	// 2074
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListAllSwimmingLanesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesRequest) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListAllSwimmingLanesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAllSwimmingLanesRequest) SetGroupId(v int64) *ListAllSwimmingLanesRequest {
	s.GroupId = &v
	return s
}

func (s *ListAllSwimmingLanesRequest) SetNamespaceId(v string) *ListAllSwimmingLanesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAllSwimmingLanesRequest) Validate() error {
	return dara.Validate(s)
}
