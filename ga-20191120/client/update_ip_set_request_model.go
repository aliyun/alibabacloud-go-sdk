// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *UpdateIpSetRequest
	GetBandwidth() *int32
	SetClientToken(v string) *UpdateIpSetRequest
	GetClientToken() *string
	SetIpSetId(v string) *UpdateIpSetRequest
	GetIpSetId() *string
	SetRegionId(v string) *UpdateIpSetRequest
	GetRegionId() *string
}

type UpdateIpSetRequest struct {
	// The new bandwidth that you want to allocate to the acceleration region. Unit: Mbit/s.
	//
	// You must allocate at least 2 Mbit/s of bandwidth to each acceleration region.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 7D2F7E4E-B958-439C-9821-56D6213A31EC
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the acceleration region that you want to modify.
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

func (s UpdateIpSetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpSetRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateIpSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIpSetRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *UpdateIpSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpSetRequest) SetBandwidth(v int32) *UpdateIpSetRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateIpSetRequest) SetClientToken(v string) *UpdateIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpSetRequest) SetIpSetId(v string) *UpdateIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *UpdateIpSetRequest) SetRegionId(v string) *UpdateIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpSetRequest) Validate() error {
	return dara.Validate(s)
}
