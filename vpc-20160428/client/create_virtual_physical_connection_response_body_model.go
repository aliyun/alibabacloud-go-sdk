// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualPhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVirtualPhysicalConnectionResponseBody
	GetRequestId() *string
	SetVirtualPhysicalConnection(v string) *CreateVirtualPhysicalConnectionResponseBody
	GetVirtualPhysicalConnection() *string
}

type CreateVirtualPhysicalConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD14EA74-E9C3-59A9-942A-DFEC7E12818D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the hosted connection.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	VirtualPhysicalConnection *string `json:"VirtualPhysicalConnection,omitempty" xml:"VirtualPhysicalConnection,omitempty"`
}

func (s CreateVirtualPhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualPhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualPhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirtualPhysicalConnectionResponseBody) GetVirtualPhysicalConnection() *string {
	return s.VirtualPhysicalConnection
}

func (s *CreateVirtualPhysicalConnectionResponseBody) SetRequestId(v string) *CreateVirtualPhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionResponseBody) SetVirtualPhysicalConnection(v string) *CreateVirtualPhysicalConnectionResponseBody {
	s.VirtualPhysicalConnection = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
