// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*ListGroupsForApplicationResponseBodyGroups) *ListGroupsForApplicationResponseBody
	GetGroups() []*ListGroupsForApplicationResponseBodyGroups
	SetRequestId(v string) *ListGroupsForApplicationResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupsForApplicationResponseBody
	GetTotalCount() *int64
}

type ListGroupsForApplicationResponseBody struct {
	// The group IDs.
	Groups []*ListGroupsForApplicationResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsForApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponseBody) GetGroups() []*ListGroupsForApplicationResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsForApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsForApplicationResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsForApplicationResponseBody) SetGroups(v []*ListGroupsForApplicationResponseBodyGroups) *ListGroupsForApplicationResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForApplicationResponseBody) SetRequestId(v string) *ListGroupsForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForApplicationResponseBody) SetTotalCount(v int64) *ListGroupsForApplicationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsForApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGroupsForApplicationResponseBodyGroups struct {
	// The group ID.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListGroupsForApplicationResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForApplicationResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsForApplicationResponseBodyGroups) SetGroupId(v string) *ListGroupsForApplicationResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForApplicationResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}
