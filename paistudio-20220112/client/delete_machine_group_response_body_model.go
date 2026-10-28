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
	// The ID of the machine group.
	//
	// example:
	//
	// mg1
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
