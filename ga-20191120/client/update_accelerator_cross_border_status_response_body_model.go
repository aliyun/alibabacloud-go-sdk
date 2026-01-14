// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorCrossBorderStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAcceleratorCrossBorderStatusResponseBody
	GetRequestId() *string
}

type UpdateAcceleratorCrossBorderStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAcceleratorCrossBorderStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorCrossBorderStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorCrossBorderStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAcceleratorCrossBorderStatusResponseBody) SetRequestId(v string) *UpdateAcceleratorCrossBorderStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
