// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteGlobalResourceRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DeleteGlobalResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteGlobalResourceRequest
	GetRegionId() *string
	SetResourceName(v string) *DeleteGlobalResourceRequest
	GetResourceName() *string
	SetResourceOwnerAccount(v string) *DeleteGlobalResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteGlobalResourceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DeleteGlobalResourceRequest
	GetResourceType() *string
	SetZoneId(v string) *DeleteGlobalResourceRequest
	GetZoneId() *string
}

type DeleteGlobalResourceRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteGlobalResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteGlobalResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteGlobalResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGlobalResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DeleteGlobalResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteGlobalResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteGlobalResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteGlobalResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteGlobalResourceRequest) SetClusterId(v string) *DeleteGlobalResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetOwnerId(v int64) *DeleteGlobalResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetRegionId(v string) *DeleteGlobalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceName(v string) *DeleteGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceOwnerAccount(v string) *DeleteGlobalResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceOwnerId(v int64) *DeleteGlobalResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceType(v string) *DeleteGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetZoneId(v string) *DeleteGlobalResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) Validate() error {
	return dara.Validate(s)
}
