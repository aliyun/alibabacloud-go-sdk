// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicIpSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteBasicIpSetRequest
	GetClientToken() *string
	SetIpSetId(v string) *DeleteBasicIpSetRequest
	GetIpSetId() *string
	SetRegionId(v string) *DeleteBasicIpSetRequest
	GetRegionId() *string
}

type DeleteBasicIpSetRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBasicIpSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicIpSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicIpSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBasicIpSetRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *DeleteBasicIpSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBasicIpSetRequest) SetClientToken(v string) *DeleteBasicIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicIpSetRequest) SetIpSetId(v string) *DeleteBasicIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *DeleteBasicIpSetRequest) SetRegionId(v string) *DeleteBasicIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBasicIpSetRequest) Validate() error {
	return dara.Validate(s)
}
