// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterServiceConfigHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClusterServiceConfigHistoryRequest
	GetClusterId() *string
	SetOwnerId(v int64) *ListClusterServiceConfigHistoryRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListClusterServiceConfigHistoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListClusterServiceConfigHistoryRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListClusterServiceConfigHistoryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListClusterServiceConfigHistoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListClusterServiceConfigHistoryRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *ListClusterServiceConfigHistoryRequest
	GetZoneId() *string
}

type ListClusterServiceConfigHistoryRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClusterServiceConfigHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterServiceConfigHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListClusterServiceConfigHistoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClusterServiceConfigHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterServiceConfigHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClusterServiceConfigHistoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListClusterServiceConfigHistoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListClusterServiceConfigHistoryRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListClusterServiceConfigHistoryRequest) SetClusterId(v string) *ListClusterServiceConfigHistoryRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetOwnerId(v int64) *ListClusterServiceConfigHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetPageNumber(v int32) *ListClusterServiceConfigHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetPageSize(v int32) *ListClusterServiceConfigHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetRegionId(v string) *ListClusterServiceConfigHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetResourceOwnerAccount(v string) *ListClusterServiceConfigHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetResourceOwnerId(v int64) *ListClusterServiceConfigHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetZoneId(v string) *ListClusterServiceConfigHistoryRequest {
	s.ZoneId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) Validate() error {
	return dara.Validate(s)
}
