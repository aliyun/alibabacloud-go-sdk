// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpAddress(v string) *GetBasicAccelerateIpResponseBody
	GetAccelerateIpAddress() *string
	SetAccelerateIpId(v string) *GetBasicAccelerateIpResponseBody
	GetAccelerateIpId() *string
	SetAcceleratorId(v string) *GetBasicAccelerateIpResponseBody
	GetAcceleratorId() *string
	SetIpSetId(v string) *GetBasicAccelerateIpResponseBody
	GetIpSetId() *string
	SetRequestId(v string) *GetBasicAccelerateIpResponseBody
	GetRequestId() *string
	SetState(v string) *GetBasicAccelerateIpResponseBody
	GetState() *string
}

type GetBasicAccelerateIpResponseBody struct {
	// The accelerated IP address of the basic GA instance.
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
	// The ID of the basic GA instance to which the queried accelerated IP address belongs.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the acceleration region of the basic GA instance.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the accelerated IP address. Valid values:
	//
	// 	- **active**: The accelerated IP address is available.
	//
	// 	- **binding**: The accelerated IP address is being associated.
	//
	// 	- **bound**: The accelerated IP address is associated.
	//
	// 	- **unbinding**: The accelerated IP address is being disassociated.
	//
	// 	- **deleting**: The GA instance is being deleted.
	//
	// >  If the accelerated IP address is being created, this parameter is not returned.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicAccelerateIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpResponseBody) GetAccelerateIpAddress() *string {
	return s.AccelerateIpAddress
}

func (s *GetBasicAccelerateIpResponseBody) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *GetBasicAccelerateIpResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetBasicAccelerateIpResponseBody) GetIpSetId() *string {
	return s.IpSetId
}

func (s *GetBasicAccelerateIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicAccelerateIpResponseBody) GetState() *string {
	return s.State
}

func (s *GetBasicAccelerateIpResponseBody) SetAccelerateIpAddress(v string) *GetBasicAccelerateIpResponseBody {
	s.AccelerateIpAddress = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetAccelerateIpId(v string) *GetBasicAccelerateIpResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetAcceleratorId(v string) *GetBasicAccelerateIpResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetIpSetId(v string) *GetBasicAccelerateIpResponseBody {
	s.IpSetId = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetRequestId(v string) *GetBasicAccelerateIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetState(v string) *GetBasicAccelerateIpResponseBody {
	s.State = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) Validate() error {
	return dara.Validate(s)
}
