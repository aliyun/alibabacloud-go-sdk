// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAppointmentRestartInstanceNodeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CancelAppointmentRestartInstanceNodeListRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CancelAppointmentRestartInstanceNodeListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelAppointmentRestartInstanceNodeListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *CancelAppointmentRestartInstanceNodeListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *CancelAppointmentRestartInstanceNodeListRequest
	GetPageSize() *int32
	SetRegionId(v string) *CancelAppointmentRestartInstanceNodeListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelAppointmentRestartInstanceNodeListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelAppointmentRestartInstanceNodeListRequest
	GetResourceOwnerId() *int64
}

type CancelAppointmentRestartInstanceNodeListRequest struct {
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

func (s CancelAppointmentRestartInstanceNodeListRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAppointmentRestartInstanceNodeListRequest) GoString() string {
	return s.String()
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) SetDBClusterId(v string) *CancelAppointmentRestartInstanceNodeListRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) SetOwnerAccount(v string) *CancelAppointmentRestartInstanceNodeListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) SetOwnerId(v int64) *CancelAppointmentRestartInstanceNodeListRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) SetPageNumber(v int32) *CancelAppointmentRestartInstanceNodeListRequest {
	s.PageNumber = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) SetPageSize(v int32) *CancelAppointmentRestartInstanceNodeListRequest {
	s.PageSize = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) SetRegionId(v string) *CancelAppointmentRestartInstanceNodeListRequest {
	s.RegionId = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) SetResourceOwnerAccount(v string) *CancelAppointmentRestartInstanceNodeListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) SetResourceOwnerId(v int64) *CancelAppointmentRestartInstanceNodeListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListRequest) Validate() error {
	return dara.Validate(s)
}
