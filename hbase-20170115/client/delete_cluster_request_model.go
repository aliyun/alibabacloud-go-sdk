// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteClusterRequest
	GetClientToken() *string
	SetClusterId(v string) *DeleteClusterRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DeleteClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteClusterRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DeleteClusterRequest
	GetZoneId() *string
}

type DeleteClusterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteClusterRequest) SetClientToken(v string) *DeleteClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetOwnerId(v int64) *DeleteClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteClusterRequest) SetRegionId(v string) *DeleteClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteClusterRequest) SetResourceOwnerAccount(v string) *DeleteClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteClusterRequest) SetResourceOwnerId(v int64) *DeleteClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteClusterRequest) SetZoneId(v string) *DeleteClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteClusterRequest) Validate() error {
	return dara.Validate(s)
}
