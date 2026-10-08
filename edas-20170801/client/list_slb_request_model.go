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
	// The address type. Valid values:
	//
	// - Internet: public address.
	//
	// - Intranet: private network address.
	//
	// example:
	//
	// internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The SLB type. Valid values:
	//
	// - clb: classic load balancing.
	//
	// - alb: application load balancing.
	//
	// example:
	//
	// clb
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
	// The VPC ID.
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
