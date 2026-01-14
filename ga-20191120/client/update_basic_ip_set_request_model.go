// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicIpSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *UpdateBasicIpSetRequest
	GetBandwidth() *int32
	SetClientToken(v string) *UpdateBasicIpSetRequest
	GetClientToken() *string
	SetIpSetId(v string) *UpdateBasicIpSetRequest
	GetIpSetId() *string
	SetRegionId(v string) *UpdateBasicIpSetRequest
	GetRegionId() *string
}

type UpdateBasicIpSetRequest struct {
	// The bandwidth of the acceleration region. Unit: Mbit/s.
	//
	// The minimum bandwidth is 2 Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the acceleration region of the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The ID of the region where the basic GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBasicIpSetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicIpSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicIpSetRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateBasicIpSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateBasicIpSetRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *UpdateBasicIpSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBasicIpSetRequest) SetBandwidth(v int32) *UpdateBasicIpSetRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateBasicIpSetRequest) SetClientToken(v string) *UpdateBasicIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBasicIpSetRequest) SetIpSetId(v string) *UpdateBasicIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *UpdateBasicIpSetRequest) SetRegionId(v string) *UpdateBasicIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBasicIpSetRequest) Validate() error {
	return dara.Validate(s)
}
