// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudConnectNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *DeleteCloudConnectNetworkRequest
	GetCcnId() *string
	SetOwnerAccount(v string) *DeleteCloudConnectNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCloudConnectNetworkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCloudConnectNetworkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteCloudConnectNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCloudConnectNetworkRequest
	GetResourceOwnerId() *int64
}

type DeleteCloudConnectNetworkRequest struct {
	// The ID of the CCN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccn-bxuau4ezctts2*****
	CcnId        *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CCN instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCloudConnectNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudConnectNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudConnectNetworkRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *DeleteCloudConnectNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCloudConnectNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCloudConnectNetworkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCloudConnectNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCloudConnectNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCloudConnectNetworkRequest) SetCcnId(v string) *DeleteCloudConnectNetworkRequest {
	s.CcnId = &v
	return s
}

func (s *DeleteCloudConnectNetworkRequest) SetOwnerAccount(v string) *DeleteCloudConnectNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCloudConnectNetworkRequest) SetOwnerId(v int64) *DeleteCloudConnectNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCloudConnectNetworkRequest) SetRegionId(v string) *DeleteCloudConnectNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCloudConnectNetworkRequest) SetResourceOwnerAccount(v string) *DeleteCloudConnectNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCloudConnectNetworkRequest) SetResourceOwnerId(v int64) *DeleteCloudConnectNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCloudConnectNetworkRequest) Validate() error {
	return dara.Validate(s)
}
