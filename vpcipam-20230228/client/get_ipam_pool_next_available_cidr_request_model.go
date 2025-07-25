// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpamPoolNextAvailableCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *GetIpamPoolNextAvailableCidrRequest
	GetCidrBlock() *string
	SetCidrMask(v int32) *GetIpamPoolNextAvailableCidrRequest
	GetCidrMask() *int32
	SetClientToken(v string) *GetIpamPoolNextAvailableCidrRequest
	GetClientToken() *string
	SetIpamPoolId(v string) *GetIpamPoolNextAvailableCidrRequest
	GetIpamPoolId() *string
	SetRegionId(v string) *GetIpamPoolNextAvailableCidrRequest
	GetRegionId() *string
}

type GetIpamPoolNextAvailableCidrRequest struct {
	// CIDR to be allocated.
	//
	// >  You must enter at least one of the CidrBlock and CidrMask fields.
	//
	// example:
	//
	// 172.68.0.0/26
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The length of the CIDR mask to be allocated.
	//
	// >  You must enter at least one of the CidrBlock and CidrMask fields.
	//
	// example:
	//
	// 26
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The region of the IPAM pool.
	//
	// >  If the IPAM pool has the region attribute, this parameter is set to the effective region of the IPAM pool. If not, this is set to the hosted region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIpamPoolNextAvailableCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIpamPoolNextAvailableCidrRequest) GoString() string {
	return s.String()
}

func (s *GetIpamPoolNextAvailableCidrRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *GetIpamPoolNextAvailableCidrRequest) GetCidrMask() *int32 {
	return s.CidrMask
}

func (s *GetIpamPoolNextAvailableCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetIpamPoolNextAvailableCidrRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *GetIpamPoolNextAvailableCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetCidrBlock(v string) *GetIpamPoolNextAvailableCidrRequest {
	s.CidrBlock = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetCidrMask(v int32) *GetIpamPoolNextAvailableCidrRequest {
	s.CidrMask = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetClientToken(v string) *GetIpamPoolNextAvailableCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetIpamPoolId(v string) *GetIpamPoolNextAvailableCidrRequest {
	s.IpamPoolId = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) SetRegionId(v string) *GetIpamPoolNextAvailableCidrRequest {
	s.RegionId = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrRequest) Validate() error {
	return dara.Validate(s)
}
