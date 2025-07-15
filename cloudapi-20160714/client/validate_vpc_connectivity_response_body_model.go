// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateVpcConnectivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnected(v bool) *ValidateVpcConnectivityResponseBody
	GetConnected() *bool
	SetIpType(v string) *ValidateVpcConnectivityResponseBody
	GetIpType() *string
	SetRequestId(v string) *ValidateVpcConnectivityResponseBody
	GetRequestId() *string
}

type ValidateVpcConnectivityResponseBody struct {
	// Indicates whether the API Gateway instance is connected to the port. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Connected *bool `json:"Connected,omitempty" xml:"Connected,omitempty"`
	// Indicates whether the instance in the authorization is an ECS instance or an SLB instance when the instance ID in the authorization is an IP address. Valid values:
	//
	// 	- **ECS**
	//
	// 	- **SLB**
	//
	// 	- **INVALID**: The instance type corresponding to the IP address is invalid.
	//
	// example:
	//
	// ECS
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A591B5B-0EC2-5463-B8B8-1984AE3AEBF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateVpcConnectivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateVpcConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateVpcConnectivityResponseBody) GetConnected() *bool {
	return s.Connected
}

func (s *ValidateVpcConnectivityResponseBody) GetIpType() *string {
	return s.IpType
}

func (s *ValidateVpcConnectivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateVpcConnectivityResponseBody) SetConnected(v bool) *ValidateVpcConnectivityResponseBody {
	s.Connected = &v
	return s
}

func (s *ValidateVpcConnectivityResponseBody) SetIpType(v string) *ValidateVpcConnectivityResponseBody {
	s.IpType = &v
	return s
}

func (s *ValidateVpcConnectivityResponseBody) SetRequestId(v string) *ValidateVpcConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateVpcConnectivityResponseBody) Validate() error {
	return dara.Validate(s)
}
