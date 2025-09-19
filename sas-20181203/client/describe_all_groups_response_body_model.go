// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeAllGroupsResponseBody
	GetCount() *int32
	SetGroups(v []*DescribeAllGroupsResponseBodyGroups) *DescribeAllGroupsResponseBody
	GetGroups() []*DescribeAllGroupsResponseBodyGroups
	SetRequestId(v string) *DescribeAllGroupsResponseBody
	GetRequestId() *string
}

type DescribeAllGroupsResponseBody struct {
	// The total number of server groups.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The grouping information about the servers.
	Groups []*DescribeAllGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAllGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllGroupsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAllGroupsResponseBody) GetGroups() []*DescribeAllGroupsResponseBodyGroups {
	return s.Groups
}

func (s *DescribeAllGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllGroupsResponseBody) SetCount(v int32) *DescribeAllGroupsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAllGroupsResponseBody) SetGroups(v []*DescribeAllGroupsResponseBodyGroups) *DescribeAllGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeAllGroupsResponseBody) SetRequestId(v string) *DescribeAllGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAllGroupsResponseBodyGroups struct {
	// The type of the server group. Valid values:
	//
	// 	- **0**: default group
	//
	// 	- **1**: other groups
	//
	// example:
	//
	// 1
	GroupFlag *int32 `json:"GroupFlag,omitempty" xml:"GroupFlag,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// 8834224
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the server group.
	//
	// example:
	//
	// abc
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeAllGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeAllGroupsResponseBodyGroups) GetGroupFlag() *int32 {
	return s.GroupFlag
}

func (s *DescribeAllGroupsResponseBodyGroups) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DescribeAllGroupsResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeAllGroupsResponseBodyGroups) SetGroupFlag(v int32) *DescribeAllGroupsResponseBodyGroups {
	s.GroupFlag = &v
	return s
}

func (s *DescribeAllGroupsResponseBodyGroups) SetGroupId(v int32) *DescribeAllGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeAllGroupsResponseBodyGroups) SetGroupName(v string) *DescribeAllGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeAllGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}
