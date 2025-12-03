// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySparkRelateHBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *QuerySparkRelateHBaseRequest
	GetClusterId() *string
	SetOwnerId(v int64) *QuerySparkRelateHBaseRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QuerySparkRelateHBaseRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *QuerySparkRelateHBaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySparkRelateHBaseRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *QuerySparkRelateHBaseRequest
	GetZoneId() *string
}

type QuerySparkRelateHBaseRequest struct {
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QuerySparkRelateHBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySparkRelateHBaseRequest) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QuerySparkRelateHBaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySparkRelateHBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QuerySparkRelateHBaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySparkRelateHBaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySparkRelateHBaseRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *QuerySparkRelateHBaseRequest) SetClusterId(v string) *QuerySparkRelateHBaseRequest {
	s.ClusterId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetOwnerId(v int64) *QuerySparkRelateHBaseRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetRegionId(v string) *QuerySparkRelateHBaseRequest {
	s.RegionId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetResourceOwnerAccount(v string) *QuerySparkRelateHBaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetResourceOwnerId(v int64) *QuerySparkRelateHBaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) SetZoneId(v string) *QuerySparkRelateHBaseRequest {
	s.ZoneId = &v
	return s
}

func (s *QuerySparkRelateHBaseRequest) Validate() error {
	return dara.Validate(s)
}
