// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupMachineGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineGroups(v []*MachineGroup) *ListResourceGroupMachineGroupsResponseBody
	GetMachineGroups() []*MachineGroup
	SetRequestId(v string) *ListResourceGroupMachineGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListResourceGroupMachineGroupsResponseBody
	GetTotalCount() *string
}

type ListResourceGroupMachineGroupsResponseBody struct {
	MachineGroups []*MachineGroup `json:"MachineGroups,omitempty" xml:"MachineGroups,omitempty" type:"Repeated"`
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupMachineGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMachineGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMachineGroupsResponseBody) GetMachineGroups() []*MachineGroup {
	return s.MachineGroups
}

func (s *ListResourceGroupMachineGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupMachineGroupsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListResourceGroupMachineGroupsResponseBody) SetMachineGroups(v []*MachineGroup) *ListResourceGroupMachineGroupsResponseBody {
	s.MachineGroups = v
	return s
}

func (s *ListResourceGroupMachineGroupsResponseBody) SetRequestId(v string) *ListResourceGroupMachineGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupMachineGroupsResponseBody) SetTotalCount(v string) *ListResourceGroupMachineGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceGroupMachineGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
