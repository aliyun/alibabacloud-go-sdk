// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkRelateHBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *SparkRelateHBaseRequest
	GetClusterId() *string
	SetHBaseClusterIds(v string) *SparkRelateHBaseRequest
	GetHBaseClusterIds() *string
	SetOwnerId(v int64) *SparkRelateHBaseRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SparkRelateHBaseRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SparkRelateHBaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SparkRelateHBaseRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *SparkRelateHBaseRequest
	GetZoneId() *string
}

type SparkRelateHBaseRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	HBaseClusterIds      *string `json:"HBaseClusterIds,omitempty" xml:"HBaseClusterIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SparkRelateHBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s SparkRelateHBaseRequest) GoString() string {
	return s.String()
}

func (s *SparkRelateHBaseRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *SparkRelateHBaseRequest) GetHBaseClusterIds() *string {
	return s.HBaseClusterIds
}

func (s *SparkRelateHBaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SparkRelateHBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SparkRelateHBaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SparkRelateHBaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SparkRelateHBaseRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *SparkRelateHBaseRequest) SetClusterId(v string) *SparkRelateHBaseRequest {
	s.ClusterId = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetHBaseClusterIds(v string) *SparkRelateHBaseRequest {
	s.HBaseClusterIds = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetOwnerId(v int64) *SparkRelateHBaseRequest {
	s.OwnerId = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetRegionId(v string) *SparkRelateHBaseRequest {
	s.RegionId = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetResourceOwnerAccount(v string) *SparkRelateHBaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetResourceOwnerId(v int64) *SparkRelateHBaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SparkRelateHBaseRequest) SetZoneId(v string) *SparkRelateHBaseRequest {
	s.ZoneId = &v
	return s
}

func (s *SparkRelateHBaseRequest) Validate() error {
	return dara.Validate(s)
}
