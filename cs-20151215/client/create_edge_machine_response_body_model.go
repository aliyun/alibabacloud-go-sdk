// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeMachineId(v string) *CreateEdgeMachineResponseBody
	GetEdgeMachineId() *string
	SetRequestId(v string) *CreateEdgeMachineResponseBody
	GetRequestId() *string
}

type CreateEdgeMachineResponseBody struct {
	// The ID of the cloud-native box.
	//
	// example:
	//
	// cc0725ddf688744979cd98445f67e****
	EdgeMachineId *string `json:"edge_machine_id,omitempty" xml:"edge_machine_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// "request_id": "6e7b377a-c5ed-4388-8026-689e1b34****",
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s CreateEdgeMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeMachineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEdgeMachineResponseBody) GetEdgeMachineId() *string {
	return s.EdgeMachineId
}

func (s *CreateEdgeMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEdgeMachineResponseBody) SetEdgeMachineId(v string) *CreateEdgeMachineResponseBody {
	s.EdgeMachineId = &v
	return s
}

func (s *CreateEdgeMachineResponseBody) SetRequestId(v string) *CreateEdgeMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEdgeMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
