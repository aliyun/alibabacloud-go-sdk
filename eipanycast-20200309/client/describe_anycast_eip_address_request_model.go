// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastId(v string) *DescribeAnycastEipAddressRequest
	GetAnycastId() *string
	SetBindInstanceId(v string) *DescribeAnycastEipAddressRequest
	GetBindInstanceId() *string
	SetIp(v string) *DescribeAnycastEipAddressRequest
	GetIp() *string
}

type DescribeAnycastEipAddressRequest struct {
	// The ID of the Anycast EIP.
	//
	// > You must specify **Ip*	- or **AnycastId**.
	//
	// example:
	//
	// aeip-bp1ix34fralt4ykf3****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The ID of the endpoint with which the Anycast EIP is associated.
	//
	// example:
	//
	// lb-2zebb08phyczzawe****
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The IP address of the Anycast EIP.
	//
	// > You must specify **Ip*	- or **AnycastId**.
	//
	// example:
	//
	// 139.95.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeAnycastEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressRequest) GetAnycastId() *string {
	return s.AnycastId
}

func (s *DescribeAnycastEipAddressRequest) GetBindInstanceId() *string {
	return s.BindInstanceId
}

func (s *DescribeAnycastEipAddressRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeAnycastEipAddressRequest) SetAnycastId(v string) *DescribeAnycastEipAddressRequest {
	s.AnycastId = &v
	return s
}

func (s *DescribeAnycastEipAddressRequest) SetBindInstanceId(v string) *DescribeAnycastEipAddressRequest {
	s.BindInstanceId = &v
	return s
}

func (s *DescribeAnycastEipAddressRequest) SetIp(v string) *DescribeAnycastEipAddressRequest {
	s.Ip = &v
	return s
}

func (s *DescribeAnycastEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
