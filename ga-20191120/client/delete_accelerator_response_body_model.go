// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteAcceleratorResponseBody
	GetAcceleratorId() *string
	SetRequestId(v string) *DeleteAcceleratorResponseBody
	GetRequestId() *string
}

type DeleteAcceleratorResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAcceleratorResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAcceleratorResponseBody) SetAcceleratorId(v string) *DeleteAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteAcceleratorResponseBody) SetRequestId(v string) *DeleteAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
