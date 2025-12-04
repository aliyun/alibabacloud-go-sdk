// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CancelRestartInstanceRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CancelRestartInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelRestartInstanceRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *CancelRestartInstanceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *CancelRestartInstanceRequest
	GetPageSize() *int32
	SetRegionId(v string) *CancelRestartInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelRestartInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelRestartInstanceRequest
	GetResourceOwnerId() *int64
	SetRestartTime(v string) *CancelRestartInstanceRequest
	GetRestartTime() *string
}

type CancelRestartInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1t9lbb7a4z7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: 30. Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
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
	RestartTime          *string `json:"RestartTime,omitempty" xml:"RestartTime,omitempty"`
}

func (s CancelRestartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *CancelRestartInstanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CancelRestartInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelRestartInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelRestartInstanceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CancelRestartInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CancelRestartInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelRestartInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelRestartInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelRestartInstanceRequest) GetRestartTime() *string {
	return s.RestartTime
}

func (s *CancelRestartInstanceRequest) SetDBClusterId(v string) *CancelRestartInstanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelRestartInstanceRequest) SetOwnerAccount(v string) *CancelRestartInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelRestartInstanceRequest) SetOwnerId(v int64) *CancelRestartInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelRestartInstanceRequest) SetPageNumber(v int32) *CancelRestartInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *CancelRestartInstanceRequest) SetPageSize(v int32) *CancelRestartInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *CancelRestartInstanceRequest) SetRegionId(v string) *CancelRestartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CancelRestartInstanceRequest) SetResourceOwnerAccount(v string) *CancelRestartInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelRestartInstanceRequest) SetResourceOwnerId(v int64) *CancelRestartInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelRestartInstanceRequest) SetRestartTime(v string) *CancelRestartInstanceRequest {
	s.RestartTime = &v
	return s
}

func (s *CancelRestartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
