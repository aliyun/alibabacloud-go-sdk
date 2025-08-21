// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIp(v string) *DescribeDeviceGatewayRequest
	GetClientIp() *string
	SetExpire(v int64) *DescribeDeviceGatewayRequest
	GetExpire() *int64
	SetId(v string) *DescribeDeviceGatewayRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeDeviceGatewayRequest
	GetOwnerId() *int64
}

type DescribeDeviceGatewayRequest struct {
	// example:
	//
	// 192.168.0.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 3600
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24611****70597051-cn-beijing
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDeviceGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceGatewayRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceGatewayRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeDeviceGatewayRequest) GetExpire() *int64 {
	return s.Expire
}

func (s *DescribeDeviceGatewayRequest) GetId() *string {
	return s.Id
}

func (s *DescribeDeviceGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDeviceGatewayRequest) SetClientIp(v string) *DescribeDeviceGatewayRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeDeviceGatewayRequest) SetExpire(v int64) *DescribeDeviceGatewayRequest {
	s.Expire = &v
	return s
}

func (s *DescribeDeviceGatewayRequest) SetId(v string) *DescribeDeviceGatewayRequest {
	s.Id = &v
	return s
}

func (s *DescribeDeviceGatewayRequest) SetOwnerId(v int64) *DescribeDeviceGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeviceGatewayRequest) Validate() error {
	return dara.Validate(s)
}
