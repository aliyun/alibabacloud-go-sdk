// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorCrossBorderModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAcceleratorCrossBorderModeResponseBody
	GetRequestId() *string
}

type UpdateAcceleratorCrossBorderModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAcceleratorCrossBorderModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorCrossBorderModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorCrossBorderModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAcceleratorCrossBorderModeResponseBody) SetRequestId(v string) *UpdateAcceleratorCrossBorderModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAcceleratorCrossBorderModeResponseBody) Validate() error {
	return dara.Validate(s)
}
