// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAccelerateIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpAddress(v string) *CreateBasicAccelerateIpResponseBody
	GetAccelerateIpAddress() *string
	SetAccelerateIpId(v string) *CreateBasicAccelerateIpResponseBody
	GetAccelerateIpId() *string
	SetAcceleratorId(v string) *CreateBasicAccelerateIpResponseBody
	GetAcceleratorId() *string
	SetIpSetId(v string) *CreateBasicAccelerateIpResponseBody
	GetIpSetId() *string
	SetRequestId(v string) *CreateBasicAccelerateIpResponseBody
	GetRequestId() *string
	SetState(v string) *CreateBasicAccelerateIpResponseBody
	GetState() *string
}

type CreateBasicAccelerateIpResponseBody struct {
	// The accelerated IP address of the basic GA instance.
	//
	// >  This parameter is unavailable.
	//
	// example:
	//
	// 116.132.XX.XX
	AccelerateIpAddress *string `json:"AccelerateIpAddress,omitempty" xml:"AccelerateIpAddress,omitempty"`
	// The ID of the accelerated IP address.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the acceleration region.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the accelerated IP address. Valid values:
	//
	// **init:*	- The accelerated IP address is being initialized.
	//
	// example:
	//
	// null
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateBasicAccelerateIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpResponseBody) GetAccelerateIpAddress() *string {
	return s.AccelerateIpAddress
}

func (s *CreateBasicAccelerateIpResponseBody) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *CreateBasicAccelerateIpResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicAccelerateIpResponseBody) GetIpSetId() *string {
	return s.IpSetId
}

func (s *CreateBasicAccelerateIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicAccelerateIpResponseBody) GetState() *string {
	return s.State
}

func (s *CreateBasicAccelerateIpResponseBody) SetAccelerateIpAddress(v string) *CreateBasicAccelerateIpResponseBody {
	s.AccelerateIpAddress = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetAccelerateIpId(v string) *CreateBasicAccelerateIpResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetAcceleratorId(v string) *CreateBasicAccelerateIpResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetIpSetId(v string) *CreateBasicAccelerateIpResponseBody {
	s.IpSetId = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetRequestId(v string) *CreateBasicAccelerateIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetState(v string) *CreateBasicAccelerateIpResponseBody {
	s.State = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) Validate() error {
	return dara.Validate(s)
}
