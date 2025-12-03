// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iXpackRelateDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *XpackRelateDBRequest
	GetClusterId() *string
	SetDbClusterIds(v string) *XpackRelateDBRequest
	GetDbClusterIds() *string
	SetOwnerId(v int64) *XpackRelateDBRequest
	GetOwnerId() *int64
	SetRegionId(v string) *XpackRelateDBRequest
	GetRegionId() *string
	SetRelateDbType(v string) *XpackRelateDBRequest
	GetRelateDbType() *string
	SetResourceOwnerAccount(v string) *XpackRelateDBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *XpackRelateDBRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *XpackRelateDBRequest
	GetZoneId() *string
}

type XpackRelateDBRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	DbClusterIds *string `json:"DbClusterIds,omitempty" xml:"DbClusterIds,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RelateDbType         *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s XpackRelateDBRequest) String() string {
	return dara.Prettify(s)
}

func (s XpackRelateDBRequest) GoString() string {
	return s.String()
}

func (s *XpackRelateDBRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *XpackRelateDBRequest) GetDbClusterIds() *string {
	return s.DbClusterIds
}

func (s *XpackRelateDBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *XpackRelateDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *XpackRelateDBRequest) GetRelateDbType() *string {
	return s.RelateDbType
}

func (s *XpackRelateDBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *XpackRelateDBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *XpackRelateDBRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *XpackRelateDBRequest) SetClusterId(v string) *XpackRelateDBRequest {
	s.ClusterId = &v
	return s
}

func (s *XpackRelateDBRequest) SetDbClusterIds(v string) *XpackRelateDBRequest {
	s.DbClusterIds = &v
	return s
}

func (s *XpackRelateDBRequest) SetOwnerId(v int64) *XpackRelateDBRequest {
	s.OwnerId = &v
	return s
}

func (s *XpackRelateDBRequest) SetRegionId(v string) *XpackRelateDBRequest {
	s.RegionId = &v
	return s
}

func (s *XpackRelateDBRequest) SetRelateDbType(v string) *XpackRelateDBRequest {
	s.RelateDbType = &v
	return s
}

func (s *XpackRelateDBRequest) SetResourceOwnerAccount(v string) *XpackRelateDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *XpackRelateDBRequest) SetResourceOwnerId(v int64) *XpackRelateDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *XpackRelateDBRequest) SetZoneId(v string) *XpackRelateDBRequest {
	s.ZoneId = &v
	return s
}

func (s *XpackRelateDBRequest) Validate() error {
	return dara.Validate(s)
}
