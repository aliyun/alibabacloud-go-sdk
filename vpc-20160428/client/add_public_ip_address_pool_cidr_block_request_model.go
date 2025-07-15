// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPublicIpAddressPoolCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *AddPublicIpAddressPoolCidrBlockRequest
	GetCidrBlock() *string
	SetCidrMask(v int32) *AddPublicIpAddressPoolCidrBlockRequest
	GetCidrMask() *int32
	SetClientToken(v string) *AddPublicIpAddressPoolCidrBlockRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddPublicIpAddressPoolCidrBlockRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *AddPublicIpAddressPoolCidrBlockRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddPublicIpAddressPoolCidrBlockRequest
	GetOwnerId() *int64
	SetPublicIpAddressPoolId(v string) *AddPublicIpAddressPoolCidrBlockRequest
	GetPublicIpAddressPoolId() *string
	SetRegionId(v string) *AddPublicIpAddressPoolCidrBlockRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddPublicIpAddressPoolCidrBlockRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddPublicIpAddressPoolCidrBlockRequest
	GetResourceOwnerId() *int64
}

type AddPublicIpAddressPoolCidrBlockRequest struct {
	// The CIDR block.
	//
	// >  You can specify only one of **CidrBlock*	- and **CidrMask**.
	//
	// example:
	//
	// 47.0.XX.XX/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The subnet mask of the CIDR block. After you enter the subnet mask, the system automatically allocates IP addresses.
	//
	// Valid values: **24*	- to **28**.
	//
	// >  You can specify only one of **CidrBlock*	- and **CidrMask**.
	//
	// example:
	//
	// 24
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the IP address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the IP address pool to which you want to add the CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddPublicIpAddressPoolCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPublicIpAddressPoolCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetCidrMask() *int32 {
	return s.CidrMask
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetCidrBlock(v string) *AddPublicIpAddressPoolCidrBlockRequest {
	s.CidrBlock = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetCidrMask(v int32) *AddPublicIpAddressPoolCidrBlockRequest {
	s.CidrMask = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetClientToken(v string) *AddPublicIpAddressPoolCidrBlockRequest {
	s.ClientToken = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetDryRun(v bool) *AddPublicIpAddressPoolCidrBlockRequest {
	s.DryRun = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetOwnerAccount(v string) *AddPublicIpAddressPoolCidrBlockRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetOwnerId(v int64) *AddPublicIpAddressPoolCidrBlockRequest {
	s.OwnerId = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetPublicIpAddressPoolId(v string) *AddPublicIpAddressPoolCidrBlockRequest {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetRegionId(v string) *AddPublicIpAddressPoolCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetResourceOwnerAccount(v string) *AddPublicIpAddressPoolCidrBlockRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) SetResourceOwnerId(v int64) *AddPublicIpAddressPoolCidrBlockRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
