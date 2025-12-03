// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CloseBackupRequest
	GetClusterId() *string
	SetOwnerId(v int64) *CloseBackupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CloseBackupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CloseBackupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloseBackupRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *CloseBackupRequest
	GetZoneId() *string
}

type CloseBackupRequest struct {
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CloseBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseBackupRequest) GoString() string {
	return s.String()
}

func (s *CloseBackupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CloseBackupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloseBackupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloseBackupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloseBackupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloseBackupRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CloseBackupRequest) SetClusterId(v string) *CloseBackupRequest {
	s.ClusterId = &v
	return s
}

func (s *CloseBackupRequest) SetOwnerId(v int64) *CloseBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseBackupRequest) SetRegionId(v string) *CloseBackupRequest {
	s.RegionId = &v
	return s
}

func (s *CloseBackupRequest) SetResourceOwnerAccount(v string) *CloseBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloseBackupRequest) SetResourceOwnerId(v int64) *CloseBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloseBackupRequest) SetZoneId(v string) *CloseBackupRequest {
	s.ZoneId = &v
	return s
}

func (s *CloseBackupRequest) Validate() error {
	return dara.Validate(s)
}
