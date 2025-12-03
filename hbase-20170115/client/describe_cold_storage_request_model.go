// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeColdStorageRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DescribeColdStorageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeColdStorageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeColdStorageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeColdStorageRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeColdStorageRequest
	GetZoneId() *string
}

type DescribeColdStorageRequest struct {
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeColdStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeColdStorageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeColdStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColdStorageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeColdStorageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeColdStorageRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeColdStorageRequest) SetClusterId(v string) *DescribeColdStorageRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeColdStorageRequest) SetOwnerId(v int64) *DescribeColdStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeColdStorageRequest) SetRegionId(v string) *DescribeColdStorageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColdStorageRequest) SetResourceOwnerAccount(v string) *DescribeColdStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeColdStorageRequest) SetResourceOwnerId(v int64) *DescribeColdStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeColdStorageRequest) SetZoneId(v string) *DescribeColdStorageRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeColdStorageRequest) Validate() error {
	return dara.Validate(s)
}
