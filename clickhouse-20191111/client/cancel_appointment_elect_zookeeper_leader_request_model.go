// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAppointmentElectZookeeperLeaderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CancelAppointmentElectZookeeperLeaderRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CancelAppointmentElectZookeeperLeaderRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelAppointmentElectZookeeperLeaderRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *CancelAppointmentElectZookeeperLeaderRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *CancelAppointmentElectZookeeperLeaderRequest
	GetPageSize() *int32
	SetRegionId(v string) *CancelAppointmentElectZookeeperLeaderRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelAppointmentElectZookeeperLeaderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelAppointmentElectZookeeperLeaderRequest
	GetResourceOwnerId() *int64
}

type CancelAppointmentElectZookeeperLeaderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelAppointmentElectZookeeperLeaderRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAppointmentElectZookeeperLeaderRequest) GoString() string {
	return s.String()
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) SetDBClusterId(v string) *CancelAppointmentElectZookeeperLeaderRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) SetOwnerAccount(v string) *CancelAppointmentElectZookeeperLeaderRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) SetOwnerId(v int64) *CancelAppointmentElectZookeeperLeaderRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) SetPageNumber(v int32) *CancelAppointmentElectZookeeperLeaderRequest {
	s.PageNumber = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) SetPageSize(v int32) *CancelAppointmentElectZookeeperLeaderRequest {
	s.PageSize = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) SetRegionId(v string) *CancelAppointmentElectZookeeperLeaderRequest {
	s.RegionId = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) SetResourceOwnerAccount(v string) *CancelAppointmentElectZookeeperLeaderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) SetResourceOwnerId(v int64) *CancelAppointmentElectZookeeperLeaderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelAppointmentElectZookeeperLeaderRequest) Validate() error {
	return dara.Validate(s)
}
