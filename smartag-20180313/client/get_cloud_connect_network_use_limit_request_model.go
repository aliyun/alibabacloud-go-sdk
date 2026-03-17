// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudConnectNetworkUseLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetCloudConnectNetworkUseLimitRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetCloudConnectNetworkUseLimitRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetCloudConnectNetworkUseLimitRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetCloudConnectNetworkUseLimitRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCloudConnectNetworkUseLimitRequest
	GetResourceOwnerId() *int64
}

type GetCloudConnectNetworkUseLimitRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CCN instances are deployed.
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

func (s GetCloudConnectNetworkUseLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudConnectNetworkUseLimitRequest) GoString() string {
	return s.String()
}

func (s *GetCloudConnectNetworkUseLimitRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetCloudConnectNetworkUseLimitRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCloudConnectNetworkUseLimitRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCloudConnectNetworkUseLimitRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCloudConnectNetworkUseLimitRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCloudConnectNetworkUseLimitRequest) SetOwnerAccount(v string) *GetCloudConnectNetworkUseLimitRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitRequest) SetOwnerId(v int64) *GetCloudConnectNetworkUseLimitRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitRequest) SetRegionId(v string) *GetCloudConnectNetworkUseLimitRequest {
	s.RegionId = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitRequest) SetResourceOwnerAccount(v string) *GetCloudConnectNetworkUseLimitRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitRequest) SetResourceOwnerId(v int64) *GetCloudConnectNetworkUseLimitRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitRequest) Validate() error {
	return dara.Validate(s)
}
