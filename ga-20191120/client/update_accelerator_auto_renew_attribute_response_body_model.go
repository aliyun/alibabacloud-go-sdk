// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateAcceleratorAutoRenewAttributeResponseBody
	GetAcceleratorId() *string
	SetRequestId(v string) *UpdateAcceleratorAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type UpdateAcceleratorAutoRenewAttributeResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F591955F-5CB5-4CCE-A75D-17CF2085CE22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAcceleratorAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorAutoRenewAttributeResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateAcceleratorAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAcceleratorAutoRenewAttributeResponseBody) SetAcceleratorId(v string) *UpdateAcceleratorAutoRenewAttributeResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeResponseBody) SetRequestId(v string) *UpdateAcceleratorAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
