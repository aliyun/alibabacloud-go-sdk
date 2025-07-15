// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalAccelerationInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalAccelerationInstanceId(v string) *CreateGlobalAccelerationInstanceResponseBody
	GetGlobalAccelerationInstanceId() *string
	SetIpAddress(v string) *CreateGlobalAccelerationInstanceResponseBody
	GetIpAddress() *string
	SetRequestId(v string) *CreateGlobalAccelerationInstanceResponseBody
	GetRequestId() *string
}

type CreateGlobalAccelerationInstanceResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1fi6sq7npnicmjj****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	// The public IP address of the GA instance.
	//
	// If **BandwidthType*	- is set to **Sharing**, this parameter is not returned.
	//
	// example:
	//
	// 12.xx.xx.78
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalAccelerationInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalAccelerationInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalAccelerationInstanceResponseBody) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *CreateGlobalAccelerationInstanceResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *CreateGlobalAccelerationInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGlobalAccelerationInstanceResponseBody) SetGlobalAccelerationInstanceId(v string) *CreateGlobalAccelerationInstanceResponseBody {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceResponseBody) SetIpAddress(v string) *CreateGlobalAccelerationInstanceResponseBody {
	s.IpAddress = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceResponseBody) SetRequestId(v string) *CreateGlobalAccelerationInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
