// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMachineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineGroupID(v string) *DeleteMachineGroupResponseBody
	GetMachineGroupID() *string
	SetRequestId(v string) *DeleteMachineGroupResponseBody
	GetRequestId() *string
}

type DeleteMachineGroupResponseBody struct {
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMachineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMachineGroupResponseBody) GetMachineGroupID() *string {
	return s.MachineGroupID
}

func (s *DeleteMachineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMachineGroupResponseBody) SetMachineGroupID(v string) *DeleteMachineGroupResponseBody {
	s.MachineGroupID = &v
	return s
}

func (s *DeleteMachineGroupResponseBody) SetRequestId(v string) *DeleteMachineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMachineGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
