// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupStructResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupFather(v int32) *DescribeGroupStructResponseBody
	GetGroupFather() *int32
	SetGroupFlag(v int32) *DescribeGroupStructResponseBody
	GetGroupFlag() *int32
	SetGroupId(v int64) *DescribeGroupStructResponseBody
	GetGroupId() *int64
	SetGroupIndex(v int32) *DescribeGroupStructResponseBody
	GetGroupIndex() *int32
	SetGroupLevel(v int32) *DescribeGroupStructResponseBody
	GetGroupLevel() *int32
	SetGroupName(v string) *DescribeGroupStructResponseBody
	GetGroupName() *string
	SetGroups(v []*string) *DescribeGroupStructResponseBody
	GetGroups() []*string
	SetMachineNum(v int32) *DescribeGroupStructResponseBody
	GetMachineNum() *int32
	SetRequestId(v string) *DescribeGroupStructResponseBody
	GetRequestId() *string
}

type DescribeGroupStructResponseBody struct {
	// The parent node of the group.
	//
	// example:
	//
	// 958****
	GroupFather *int32 `json:"GroupFather,omitempty" xml:"GroupFather,omitempty"`
	// The type of the server group. Valid values:
	//
	// 	- **0**: the default group
	//
	// 	- **1**: other groups
	//
	// example:
	//
	// 0
	GroupFlag *int32 `json:"GroupFlag,omitempty" xml:"GroupFlag,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// 958****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The sequence number.
	//
	// example:
	//
	// 1
	GroupIndex *int32 `json:"GroupIndex,omitempty" xml:"GroupIndex,omitempty"`
	// The level of the application group.
	//
	// example:
	//
	// 2
	GroupLevel *int32 `json:"GroupLevel,omitempty" xml:"GroupLevel,omitempty"`
	// The name of the server group.
	//
	// example:
	//
	// TestGroupName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// An array that consists of child groups.
	Groups []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The number of servers in the group.
	//
	// example:
	//
	// 30
	MachineNum *int32 `json:"MachineNum,omitempty" xml:"MachineNum,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9FBC6E47-7508-58C9-9E76-528E118C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupStructResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupStructResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupStructResponseBody) GetGroupFather() *int32 {
	return s.GroupFather
}

func (s *DescribeGroupStructResponseBody) GetGroupFlag() *int32 {
	return s.GroupFlag
}

func (s *DescribeGroupStructResponseBody) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeGroupStructResponseBody) GetGroupIndex() *int32 {
	return s.GroupIndex
}

func (s *DescribeGroupStructResponseBody) GetGroupLevel() *int32 {
	return s.GroupLevel
}

func (s *DescribeGroupStructResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGroupStructResponseBody) GetGroups() []*string {
	return s.Groups
}

func (s *DescribeGroupStructResponseBody) GetMachineNum() *int32 {
	return s.MachineNum
}

func (s *DescribeGroupStructResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupStructResponseBody) SetGroupFather(v int32) *DescribeGroupStructResponseBody {
	s.GroupFather = &v
	return s
}

func (s *DescribeGroupStructResponseBody) SetGroupFlag(v int32) *DescribeGroupStructResponseBody {
	s.GroupFlag = &v
	return s
}

func (s *DescribeGroupStructResponseBody) SetGroupId(v int64) *DescribeGroupStructResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupStructResponseBody) SetGroupIndex(v int32) *DescribeGroupStructResponseBody {
	s.GroupIndex = &v
	return s
}

func (s *DescribeGroupStructResponseBody) SetGroupLevel(v int32) *DescribeGroupStructResponseBody {
	s.GroupLevel = &v
	return s
}

func (s *DescribeGroupStructResponseBody) SetGroupName(v string) *DescribeGroupStructResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupStructResponseBody) SetGroups(v []*string) *DescribeGroupStructResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeGroupStructResponseBody) SetMachineNum(v int32) *DescribeGroupStructResponseBody {
	s.MachineNum = &v
	return s
}

func (s *DescribeGroupStructResponseBody) SetRequestId(v string) *DescribeGroupStructResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupStructResponseBody) Validate() error {
	return dara.Validate(s)
}
