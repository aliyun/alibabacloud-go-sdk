// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserHdfsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteUserHdfsInfoRequest
	GetClusterId() *string
	SetNameService(v string) *DeleteUserHdfsInfoRequest
	GetNameService() *string
	SetOwnerId(v int64) *DeleteUserHdfsInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteUserHdfsInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteUserHdfsInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteUserHdfsInfoRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DeleteUserHdfsInfoRequest
	GetZoneId() *string
}

type DeleteUserHdfsInfoRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	NameService *string `json:"NameService,omitempty" xml:"NameService,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteUserHdfsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserHdfsInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteUserHdfsInfoRequest) GetNameService() *string {
	return s.NameService
}

func (s *DeleteUserHdfsInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteUserHdfsInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUserHdfsInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteUserHdfsInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteUserHdfsInfoRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteUserHdfsInfoRequest) SetClusterId(v string) *DeleteUserHdfsInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetNameService(v string) *DeleteUserHdfsInfoRequest {
	s.NameService = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetOwnerId(v int64) *DeleteUserHdfsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetRegionId(v string) *DeleteUserHdfsInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetResourceOwnerAccount(v string) *DeleteUserHdfsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetResourceOwnerId(v int64) *DeleteUserHdfsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetZoneId(v string) *DeleteUserHdfsInfoRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) Validate() error {
	return dara.Validate(s)
}
