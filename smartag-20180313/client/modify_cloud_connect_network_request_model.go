// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudConnectNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *ModifyCloudConnectNetworkRequest
	GetCcnId() *string
	SetCidrBlock(v string) *ModifyCloudConnectNetworkRequest
	GetCidrBlock() *string
	SetDescription(v string) *ModifyCloudConnectNetworkRequest
	GetDescription() *string
	SetInterworkingStatus(v string) *ModifyCloudConnectNetworkRequest
	GetInterworkingStatus() *string
	SetName(v string) *ModifyCloudConnectNetworkRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyCloudConnectNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCloudConnectNetworkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCloudConnectNetworkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCloudConnectNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCloudConnectNetworkRequest
	GetResourceOwnerId() *int64
}

type ModifyCloudConnectNetworkRequest struct {
	// The ID of the CCN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-l9340rlu5ens*****
	CcnId *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	// The private CIDR block.
	//
	// example:
	//
	// 10.10.10.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the CCN instance.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to allow the SAG instances associated with the same CCN instance to communicate with each other.
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// example:
	//
	// enable
	InterworkingStatus *string `json:"InterworkingStatus,omitempty" xml:"InterworkingStatus,omitempty"`
	// The name of the CCN instance.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Name
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CCN instance is deployed.
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

func (s ModifyCloudConnectNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudConnectNetworkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudConnectNetworkRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *ModifyCloudConnectNetworkRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *ModifyCloudConnectNetworkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCloudConnectNetworkRequest) GetInterworkingStatus() *string {
	return s.InterworkingStatus
}

func (s *ModifyCloudConnectNetworkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCloudConnectNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCloudConnectNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCloudConnectNetworkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudConnectNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCloudConnectNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCloudConnectNetworkRequest) SetCcnId(v string) *ModifyCloudConnectNetworkRequest {
	s.CcnId = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetCidrBlock(v string) *ModifyCloudConnectNetworkRequest {
	s.CidrBlock = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetDescription(v string) *ModifyCloudConnectNetworkRequest {
	s.Description = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetInterworkingStatus(v string) *ModifyCloudConnectNetworkRequest {
	s.InterworkingStatus = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetName(v string) *ModifyCloudConnectNetworkRequest {
	s.Name = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetOwnerAccount(v string) *ModifyCloudConnectNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetOwnerId(v int64) *ModifyCloudConnectNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetRegionId(v string) *ModifyCloudConnectNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetResourceOwnerAccount(v string) *ModifyCloudConnectNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) SetResourceOwnerId(v int64) *ModifyCloudConnectNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCloudConnectNetworkRequest) Validate() error {
	return dara.Validate(s)
}
