// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHBaseHaDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *QueryHBaseHaDBRequest
	GetClusterId() *string
	SetOwnerId(v int64) *QueryHBaseHaDBRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryHBaseHaDBRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *QueryHBaseHaDBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryHBaseHaDBRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *QueryHBaseHaDBRequest
	GetZoneId() *string
}

type QueryHBaseHaDBRequest struct {
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryHBaseHaDBRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHBaseHaDBRequest) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryHBaseHaDBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryHBaseHaDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryHBaseHaDBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryHBaseHaDBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryHBaseHaDBRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *QueryHBaseHaDBRequest) SetClusterId(v string) *QueryHBaseHaDBRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetOwnerId(v int64) *QueryHBaseHaDBRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetRegionId(v string) *QueryHBaseHaDBRequest {
	s.RegionId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetResourceOwnerAccount(v string) *QueryHBaseHaDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetResourceOwnerId(v int64) *QueryHBaseHaDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) SetZoneId(v string) *QueryHBaseHaDBRequest {
	s.ZoneId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) Validate() error {
	return dara.Validate(s)
}
