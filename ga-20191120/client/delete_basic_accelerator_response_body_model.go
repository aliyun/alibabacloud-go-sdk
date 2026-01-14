// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteBasicAcceleratorResponseBody
	GetAcceleratorId() *string
	SetRequestId(v string) *DeleteBasicAcceleratorResponseBody
	GetRequestId() *string
}

type DeleteBasicAcceleratorResponseBody struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicAcceleratorResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteBasicAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBasicAcceleratorResponseBody) SetAcceleratorId(v string) *DeleteBasicAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteBasicAcceleratorResponseBody) SetRequestId(v string) *DeleteBasicAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBasicAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
