// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroups(v []*ListHostGroupsResponseBodyHostGroups) *ListHostGroupsResponseBody
	GetHostGroups() []*ListHostGroupsResponseBodyHostGroups
	SetRequestId(v string) *ListHostGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostGroupsResponseBody
	GetTotalCount() *int32
}

type ListHostGroupsResponseBody struct {
	// The asset groups returned.
	HostGroups []*ListHostGroupsResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of asset groups returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBody) GetHostGroups() []*ListHostGroupsResponseBodyHostGroups {
	return s.HostGroups
}

func (s *ListHostGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostGroupsResponseBody) SetHostGroups(v []*ListHostGroupsResponseBodyHostGroups) *ListHostGroupsResponseBody {
	s.HostGroups = v
	return s
}

func (s *ListHostGroupsResponseBody) SetRequestId(v string) *ListHostGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetTotalCount(v int32) *ListHostGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostGroupsResponseBody) Validate() error {
	if s.HostGroups != nil {
		for _, item := range s.HostGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHostGroupsResponseBodyHostGroups struct {
	// The remarks of the asset group.
	//
	// example:
	//
	// Description
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The asset group ID.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The name of the asset group.
	//
	// example:
	//
	// Host group 1
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The number of hosts in the asset group.
	//
	// example:
	//
	// 1
	MemberCount *int32 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
}

func (s ListHostGroupsResponseBodyHostGroups) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsResponseBodyHostGroups) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBodyHostGroups) GetComment() *string {
	return s.Comment
}

func (s *ListHostGroupsResponseBodyHostGroups) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ListHostGroupsResponseBodyHostGroups) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *ListHostGroupsResponseBodyHostGroups) GetMemberCount() *int32 {
	return s.MemberCount
}

func (s *ListHostGroupsResponseBodyHostGroups) SetComment(v string) *ListHostGroupsResponseBodyHostGroups {
	s.Comment = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetHostGroupId(v string) *ListHostGroupsResponseBodyHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetHostGroupName(v string) *ListHostGroupsResponseBodyHostGroups {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetMemberCount(v int32) *ListHostGroupsResponseBodyHostGroups {
	s.MemberCount = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) Validate() error {
	return dara.Validate(s)
}
