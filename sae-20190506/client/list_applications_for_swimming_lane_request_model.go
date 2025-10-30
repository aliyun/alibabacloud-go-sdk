// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *ListApplicationsForSwimmingLaneRequest
	GetGroupId() *int64
	SetNamespaceId(v string) *ListApplicationsForSwimmingLaneRequest
	GetNamespaceId() *string
	SetTag(v string) *ListApplicationsForSwimmingLaneRequest
	GetTag() *string
}

type ListApplicationsForSwimmingLaneRequest struct {
	// The ID of the application group. You can call the [DescribeApplicationGroups](https://help.aliyun.com/document_detail/126249.html) operation to obtain the ID.
	//
	// example:
	//
	// b2a8a925-477a-eswa-b823-d5e22500****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of a namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The canary tag
	//
	// example:
	//
	// {"alicloud.service.tag":"gray"}
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListApplicationsForSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForSwimmingLaneRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListApplicationsForSwimmingLaneRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListApplicationsForSwimmingLaneRequest) GetTag() *string {
	return s.Tag
}

func (s *ListApplicationsForSwimmingLaneRequest) SetGroupId(v int64) *ListApplicationsForSwimmingLaneRequest {
	s.GroupId = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneRequest) SetNamespaceId(v string) *ListApplicationsForSwimmingLaneRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneRequest) SetTag(v string) *ListApplicationsForSwimmingLaneRequest {
	s.Tag = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}
