// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicAccelerateIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpId(v string) *DeleteBasicAccelerateIpResponseBody
	GetAccelerateIpId() *string
	SetRequestId(v string) *DeleteBasicAccelerateIpResponseBody
	GetRequestId() *string
}

type DeleteBasicAccelerateIpResponseBody struct {
	// The ID of the accelerated IP address that is deleted.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicAccelerateIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicAccelerateIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpResponseBody) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *DeleteBasicAccelerateIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBasicAccelerateIpResponseBody) SetAccelerateIpId(v string) *DeleteBasicAccelerateIpResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *DeleteBasicAccelerateIpResponseBody) SetRequestId(v string) *DeleteBasicAccelerateIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBasicAccelerateIpResponseBody) Validate() error {
	return dara.Validate(s)
}
