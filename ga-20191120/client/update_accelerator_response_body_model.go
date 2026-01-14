// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAcceleratorResponseBody
	GetRequestId() *string
}

type UpdateAcceleratorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAcceleratorResponseBody) SetRequestId(v string) *UpdateAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
