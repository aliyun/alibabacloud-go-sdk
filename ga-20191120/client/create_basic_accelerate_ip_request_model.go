// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAccelerateIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateBasicAccelerateIpRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateBasicAccelerateIpRequest
	GetClientToken() *string
	SetIpSetId(v string) *CreateBasicAccelerateIpRequest
	GetIpSetId() *string
	SetRegionId(v string) *CreateBasicAccelerateIpRequest
	GetRegionId() *string
}

type CreateBasicAccelerateIpRequest struct {
	// The ID of the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
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
	// The ID of the acceleration region.
	//
	// You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/2253380.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicAccelerateIpRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicAccelerateIpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBasicAccelerateIpRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *CreateBasicAccelerateIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBasicAccelerateIpRequest) SetAcceleratorId(v string) *CreateBasicAccelerateIpRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpRequest) SetClientToken(v string) *CreateBasicAccelerateIpRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicAccelerateIpRequest) SetIpSetId(v string) *CreateBasicAccelerateIpRequest {
	s.IpSetId = &v
	return s
}

func (s *CreateBasicAccelerateIpRequest) SetRegionId(v string) *CreateBasicAccelerateIpRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBasicAccelerateIpRequest) Validate() error {
	return dara.Validate(s)
}
