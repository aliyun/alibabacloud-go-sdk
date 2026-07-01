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
	// The CIDR block to be allocated.
	//
	// > You must specify CidrBlock or CidrMask.
	//
	// example:
	//
	// 172.68.0.0/26
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The mask length of the CIDR block to be allocated.
	//
	// > You must specify CidrBlock or CidrMask.
	//
	// example:
	//
	// 26
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a unique value for this parameter from your client. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the RequestId as the ClientToken. The RequestId of each request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The region of the IPAM pool.
	//
	// > If the IPAM pool has a region property, this parameter specifies the region where the pool is active. If the IPAM pool does not have a region property, this parameter specifies the managed region of IPAM.
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
