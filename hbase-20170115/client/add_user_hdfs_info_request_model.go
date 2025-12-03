// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserHdfsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddUserHdfsInfoRequest
	GetClusterId() *string
	SetExtInfo(v string) *AddUserHdfsInfoRequest
	GetExtInfo() *string
	SetOwnerId(v int64) *AddUserHdfsInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddUserHdfsInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddUserHdfsInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddUserHdfsInfoRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *AddUserHdfsInfoRequest
	GetZoneId() *string
}

type AddUserHdfsInfoRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddUserHdfsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserHdfsInfoRequest) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddUserHdfsInfoRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *AddUserHdfsInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddUserHdfsInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddUserHdfsInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddUserHdfsInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddUserHdfsInfoRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddUserHdfsInfoRequest) SetClusterId(v string) *AddUserHdfsInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetExtInfo(v string) *AddUserHdfsInfoRequest {
	s.ExtInfo = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetOwnerId(v int64) *AddUserHdfsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetRegionId(v string) *AddUserHdfsInfoRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetResourceOwnerAccount(v string) *AddUserHdfsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetResourceOwnerId(v int64) *AddUserHdfsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetZoneId(v string) *AddUserHdfsInfoRequest {
	s.ZoneId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) Validate() error {
	return dara.Validate(s)
}
