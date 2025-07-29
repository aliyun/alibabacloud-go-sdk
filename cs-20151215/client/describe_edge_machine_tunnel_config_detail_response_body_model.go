// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachineTunnelConfigDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody
	GetDeviceName() *string
	SetModel(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody
	GetModel() *string
	SetProductKey(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody
	GetProductKey() *string
	SetRequestId(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody
	GetRequestId() *string
	SetSn(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody
	GetSn() *string
	SetToken(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody
	GetToken() *string
	SetTunnelEndpoint(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody
	GetTunnelEndpoint() *string
}

type DescribeEdgeMachineTunnelConfigDetailResponseBody struct {
	// The device name.
	//
	// example:
	//
	// TEST0621N0FF****
	DeviceName *string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// The model of the cloud-native box.
	//
	// example:
	//
	// ACK-A-S001
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Product Key
	//
	// example:
	//
	// a11rXul****
	ProductKey *string `json:"product_key,omitempty" xml:"product_key,omitempty"`
	// Request ID
	//
	// example:
	//
	// bfd12953-31cb-42f1-8a36-7b80ec345094
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The serial number of the cloud-native box.
	//
	// example:
	//
	// Q2CB5XZAFBFG****
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
	// Token
	//
	// example:
	//
	// abcd****
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The tunnel endpoint.
	//
	// example:
	//
	// wss://frontend-iotx-r-debug.aliyun-inc.test
	TunnelEndpoint *string `json:"tunnel_endpoint,omitempty" xml:"tunnel_endpoint,omitempty"`
}

func (s DescribeEdgeMachineTunnelConfigDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachineTunnelConfigDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) GetModel() *string {
	return s.Model
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) GetProductKey() *string {
	return s.ProductKey
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) GetSn() *string {
	return s.Sn
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) GetToken() *string {
	return s.Token
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) GetTunnelEndpoint() *string {
	return s.TunnelEndpoint
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) SetDeviceName(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody {
	s.DeviceName = &v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) SetModel(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody {
	s.Model = &v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) SetProductKey(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody {
	s.ProductKey = &v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) SetRequestId(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) SetSn(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody {
	s.Sn = &v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) SetToken(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody {
	s.Token = &v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) SetTunnelEndpoint(v string) *DescribeEdgeMachineTunnelConfigDetailResponseBody {
	s.TunnelEndpoint = &v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
