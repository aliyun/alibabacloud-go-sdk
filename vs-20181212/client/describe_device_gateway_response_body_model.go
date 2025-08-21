// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHost(v string) *DescribeDeviceGatewayResponseBody
	GetHost() *string
	SetPort(v int64) *DescribeDeviceGatewayResponseBody
	GetPort() *int64
	SetProtocol(v string) *DescribeDeviceGatewayResponseBody
	GetProtocol() *string
	SetRequestId(v string) *DescribeDeviceGatewayResponseBody
	GetRequestId() *string
	SetToken(v string) *DescribeDeviceGatewayResponseBody
	GetToken() *string
}

type DescribeDeviceGatewayResponseBody struct {
	// example:
	//
	// 192.168.0.1
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// 8080
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// gb28181
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// f5578fbc-694c-461d-a2a2-eb090775cef0
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeDeviceGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceGatewayResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeDeviceGatewayResponseBody) GetPort() *int64 {
	return s.Port
}

func (s *DescribeDeviceGatewayResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDeviceGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceGatewayResponseBody) GetToken() *string {
	return s.Token
}

func (s *DescribeDeviceGatewayResponseBody) SetHost(v string) *DescribeDeviceGatewayResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) SetPort(v int64) *DescribeDeviceGatewayResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) SetProtocol(v string) *DescribeDeviceGatewayResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) SetRequestId(v string) *DescribeDeviceGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) SetToken(v string) *DescribeDeviceGatewayResponseBody {
	s.Token = &v
	return s
}

func (s *DescribeDeviceGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
