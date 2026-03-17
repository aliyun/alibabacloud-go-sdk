// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudConnectNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateCloudConnectNetworkRequest
	GetCidrBlock() *string
	SetDescription(v string) *CreateCloudConnectNetworkRequest
	GetDescription() *string
	SetName(v string) *CreateCloudConnectNetworkRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateCloudConnectNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCloudConnectNetworkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateCloudConnectNetworkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateCloudConnectNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCloudConnectNetworkRequest
	GetResourceOwnerId() *int64
	SetSnatCidrBlock(v string) *CreateCloudConnectNetworkRequest
	GetSnatCidrBlock() *string
}

type CreateCloudConnectNetworkRequest struct {
	// The private CIDR block.
	//
	// example:
	//
	// 172.XX.XX.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the CCN instance.
	//
	// The description must be 2 to 128 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
	//
	// example:
	//
	// ccndesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the CCN instance.
	//
	// The name must be 2 to 100 characters in length and can contain letters, digits, periods (.), underscores (_),and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// ccnname
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CCN instance is deployed.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The private CIDR block used for Source Network Address Translation (SNAT).
	//
	// example:
	//
	// 172.XX.XX.0/25
	SnatCidrBlock *string `json:"SnatCidrBlock,omitempty" xml:"SnatCidrBlock,omitempty"`
}

func (s CreateCloudConnectNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudConnectNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudConnectNetworkRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateCloudConnectNetworkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCloudConnectNetworkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCloudConnectNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCloudConnectNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCloudConnectNetworkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudConnectNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCloudConnectNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCloudConnectNetworkRequest) GetSnatCidrBlock() *string {
	return s.SnatCidrBlock
}

func (s *CreateCloudConnectNetworkRequest) SetCidrBlock(v string) *CreateCloudConnectNetworkRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) SetDescription(v string) *CreateCloudConnectNetworkRequest {
	s.Description = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) SetName(v string) *CreateCloudConnectNetworkRequest {
	s.Name = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) SetOwnerAccount(v string) *CreateCloudConnectNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) SetOwnerId(v int64) *CreateCloudConnectNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) SetRegionId(v string) *CreateCloudConnectNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) SetResourceOwnerAccount(v string) *CreateCloudConnectNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) SetResourceOwnerId(v int64) *CreateCloudConnectNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) SetSnatCidrBlock(v string) *CreateCloudConnectNetworkRequest {
	s.SnatCidrBlock = &v
	return s
}

func (s *CreateCloudConnectNetworkRequest) Validate() error {
	return dara.Validate(s)
}
