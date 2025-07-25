// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationDefaultCidrMask(v int32) *UpdateIpamPoolRequest
	GetAllocationDefaultCidrMask() *int32
	SetAllocationMaxCidrMask(v int32) *UpdateIpamPoolRequest
	GetAllocationMaxCidrMask() *int32
	SetAllocationMinCidrMask(v int32) *UpdateIpamPoolRequest
	GetAllocationMinCidrMask() *int32
	SetAutoImport(v bool) *UpdateIpamPoolRequest
	GetAutoImport() *bool
	SetClearAllocationDefaultCidrMask(v bool) *UpdateIpamPoolRequest
	GetClearAllocationDefaultCidrMask() *bool
	SetClientToken(v string) *UpdateIpamPoolRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateIpamPoolRequest
	GetDryRun() *bool
	SetIpamPoolDescription(v string) *UpdateIpamPoolRequest
	GetIpamPoolDescription() *string
	SetIpamPoolId(v string) *UpdateIpamPoolRequest
	GetIpamPoolId() *string
	SetIpamPoolName(v string) *UpdateIpamPoolRequest
	GetIpamPoolName() *string
	SetOwnerAccount(v string) *UpdateIpamPoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateIpamPoolRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateIpamPoolRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateIpamPoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateIpamPoolRequest
	GetResourceOwnerId() *int64
}

type UpdateIpamPoolRequest struct {
	// The new default network mask for the IPAM pool.
	//
	// The mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// The new maximum network mask for the IPAM pool.
	//
	// The mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// The new minimum network mask for the IPAM pool.
	//
	// The mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// Whether the pool has the auto-import feature enabled.
	//
	// example:
	//
	// true
	AutoImport *bool `json:"AutoImport,omitempty" xml:"AutoImport,omitempty"`
	// Specifies whether to delete the default network mask for the IPAM pool. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ClearAllocationDefaultCidrMask *bool `json:"ClearAllocationDefaultCidrMask,omitempty" xml:"ClearAllocationDefaultCidrMask,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new description of the IPAM pool.
	//
	// It must be 2 to 268 characters in length. It must start with a letter but cannot start with a `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The new name of the IPAM pool.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpamPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolRequest) GetAllocationDefaultCidrMask() *int32 {
	return s.AllocationDefaultCidrMask
}

func (s *UpdateIpamPoolRequest) GetAllocationMaxCidrMask() *int32 {
	return s.AllocationMaxCidrMask
}

func (s *UpdateIpamPoolRequest) GetAllocationMinCidrMask() *int32 {
	return s.AllocationMinCidrMask
}

func (s *UpdateIpamPoolRequest) GetAutoImport() *bool {
	return s.AutoImport
}

func (s *UpdateIpamPoolRequest) GetClearAllocationDefaultCidrMask() *bool {
	return s.ClearAllocationDefaultCidrMask
}

func (s *UpdateIpamPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIpamPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateIpamPoolRequest) GetIpamPoolDescription() *string {
	return s.IpamPoolDescription
}

func (s *UpdateIpamPoolRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *UpdateIpamPoolRequest) GetIpamPoolName() *string {
	return s.IpamPoolName
}

func (s *UpdateIpamPoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateIpamPoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateIpamPoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpamPoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateIpamPoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateIpamPoolRequest) SetAllocationDefaultCidrMask(v int32) *UpdateIpamPoolRequest {
	s.AllocationDefaultCidrMask = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetAllocationMaxCidrMask(v int32) *UpdateIpamPoolRequest {
	s.AllocationMaxCidrMask = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetAllocationMinCidrMask(v int32) *UpdateIpamPoolRequest {
	s.AllocationMinCidrMask = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetAutoImport(v bool) *UpdateIpamPoolRequest {
	s.AutoImport = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetClearAllocationDefaultCidrMask(v bool) *UpdateIpamPoolRequest {
	s.ClearAllocationDefaultCidrMask = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetClientToken(v string) *UpdateIpamPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetDryRun(v bool) *UpdateIpamPoolRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetIpamPoolDescription(v string) *UpdateIpamPoolRequest {
	s.IpamPoolDescription = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetIpamPoolId(v string) *UpdateIpamPoolRequest {
	s.IpamPoolId = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetIpamPoolName(v string) *UpdateIpamPoolRequest {
	s.IpamPoolName = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetOwnerAccount(v string) *UpdateIpamPoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetOwnerId(v int64) *UpdateIpamPoolRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetRegionId(v string) *UpdateIpamPoolRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetResourceOwnerAccount(v string) *UpdateIpamPoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpamPoolRequest) SetResourceOwnerId(v int64) *UpdateIpamPoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateIpamPoolRequest) Validate() error {
	return dara.Validate(s)
}
