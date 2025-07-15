// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipGatewayInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEipGatewayInfoResponseBody
	GetCode() *string
	SetEipInfos(v *DescribeEipGatewayInfoResponseBodyEipInfos) *DescribeEipGatewayInfoResponseBody
	GetEipInfos() *DescribeEipGatewayInfoResponseBodyEipInfos
	SetMessage(v string) *DescribeEipGatewayInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEipGatewayInfoResponseBody
	GetRequestId() *string
}

type DescribeEipGatewayInfoResponseBody struct {
	// The status code of the operation.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information about the EIP.
	EipInfos *DescribeEipGatewayInfoResponseBodyEipInfos `json:"EipInfos,omitempty" xml:"EipInfos,omitempty" type:"Struct"`
	// The result of the operation.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C0FD0EED-F90D-4479-803D-DD62335357E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEipGatewayInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipGatewayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipGatewayInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEipGatewayInfoResponseBody) GetEipInfos() *DescribeEipGatewayInfoResponseBodyEipInfos {
	return s.EipInfos
}

func (s *DescribeEipGatewayInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEipGatewayInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEipGatewayInfoResponseBody) SetCode(v string) *DescribeEipGatewayInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEipGatewayInfoResponseBody) SetEipInfos(v *DescribeEipGatewayInfoResponseBodyEipInfos) *DescribeEipGatewayInfoResponseBody {
	s.EipInfos = v
	return s
}

func (s *DescribeEipGatewayInfoResponseBody) SetMessage(v string) *DescribeEipGatewayInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEipGatewayInfoResponseBody) SetRequestId(v string) *DescribeEipGatewayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEipGatewayInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEipGatewayInfoResponseBodyEipInfos struct {
	EipInfo []*DescribeEipGatewayInfoResponseBodyEipInfosEipInfo `json:"EipInfo,omitempty" xml:"EipInfo,omitempty" type:"Repeated"`
}

func (s DescribeEipGatewayInfoResponseBodyEipInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipGatewayInfoResponseBodyEipInfos) GoString() string {
	return s.String()
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfos) GetEipInfo() []*DescribeEipGatewayInfoResponseBodyEipInfosEipInfo {
	return s.EipInfo
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfos) SetEipInfo(v []*DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) *DescribeEipGatewayInfoResponseBodyEipInfos {
	s.EipInfo = v
	return s
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeEipGatewayInfoResponseBodyEipInfosEipInfo struct {
	// The IP address of the EIP.
	//
	// example:
	//
	// 47.XX.XX.236
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP address of the gateway that is associated with the EIP.
	//
	// example:
	//
	// 47.XX.XX.1
	IpGw *string `json:"IpGw,omitempty" xml:"IpGw,omitempty"`
	// The subnet mask of the EIP.
	//
	// example:
	//
	// 255.255.255.0
	IpMask *string `json:"IpMask,omitempty" xml:"IpMask,omitempty"`
}

func (s DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) GoString() string {
	return s.String()
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) GetIp() *string {
	return s.Ip
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) GetIpGw() *string {
	return s.IpGw
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) GetIpMask() *string {
	return s.IpMask
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) SetIp(v string) *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo {
	s.Ip = &v
	return s
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) SetIpGw(v string) *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo {
	s.IpGw = &v
	return s
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) SetIpMask(v string) *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo {
	s.IpMask = &v
	return s
}

func (s *DescribeEipGatewayInfoResponseBodyEipInfosEipInfo) Validate() error {
	return dara.Validate(s)
}
