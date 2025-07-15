// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualPhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVirtualPhysicalConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateVirtualPhysicalConnectionResponseBody
	GetSuccess() *string
}

type UpdateVirtualPhysicalConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7035627E-1C1D-5BC7-A830-F897A35912D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the VLAN ID of the hosted connection is changed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateVirtualPhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualPhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVirtualPhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVirtualPhysicalConnectionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateVirtualPhysicalConnectionResponseBody) SetRequestId(v string) *UpdateVirtualPhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionResponseBody) SetSuccess(v string) *UpdateVirtualPhysicalConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
