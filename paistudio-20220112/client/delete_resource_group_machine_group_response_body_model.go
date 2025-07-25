// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupMachineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineGroupID(v string) *DeleteResourceGroupMachineGroupResponseBody
	GetMachineGroupID() *string
	SetRequestId(v string) *DeleteResourceGroupMachineGroupResponseBody
	GetRequestId() *string
}

type DeleteResourceGroupMachineGroupResponseBody struct {
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceGroupMachineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupMachineGroupResponseBody) GetMachineGroupID() *string {
	return s.MachineGroupID
}

func (s *DeleteResourceGroupMachineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceGroupMachineGroupResponseBody) SetMachineGroupID(v string) *DeleteResourceGroupMachineGroupResponseBody {
	s.MachineGroupID = &v
	return s
}

func (s *DeleteResourceGroupMachineGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupMachineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupMachineGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
