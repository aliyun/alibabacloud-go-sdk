// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *ListSlbRequest
	GetAddressType() *string
	SetSlbType(v string) *ListSlbRequest
	GetSlbType() *string
	SetVpcId(v string) *ListSlbRequest
	GetVpcId() *string
}

type ListSlbRequest struct {
	// The type of the IP addresses. Valid values:
	//
	// 	- Internet: Users can connect to the SLB instance over the Internet.
	//
	// 	- Intranet: Users can connect to the SLB instance over the internal network.
	//
	// example:
	//
	// internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The type of the SLB instance. Valid values:
	//
	// 	- clb: Classic Load Balancer (CLB)
	//
	// 	- alb: Application Load Balancer (ALB)
	//
	// example:
	//
	// clb
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1f90rfybszjogyw****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSlbRequest) GoString() string {
	return s.String()
}

func (s *ListSlbRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *ListSlbRequest) GetSlbType() *string {
	return s.SlbType
}

func (s *ListSlbRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListSlbRequest) SetAddressType(v string) *ListSlbRequest {
	s.AddressType = &v
	return s
}

func (s *ListSlbRequest) SetSlbType(v string) *ListSlbRequest {
	s.SlbType = &v
	return s
}

func (s *ListSlbRequest) SetVpcId(v string) *ListSlbRequest {
	s.VpcId = &v
	return s
}

func (s *ListSlbRequest) Validate() error {
	return dara.Validate(s)
}
