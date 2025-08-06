// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetStorageListRequest
	GetAppId() *string
	SetDivision(v string) *GetStorageListRequest
	GetDivision() *string
	SetOwnerAccount(v string) *GetStorageListRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetStorageListRequest
	GetOwnerId() *string
	SetPageNumber(v int32) *GetStorageListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetStorageListRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *GetStorageListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *GetStorageListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetStorageListRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *GetStorageListRequest
	GetResourceRealOwnerId() *int64
	SetStorageRegion(v string) *GetStorageListRequest
	GetStorageRegion() *string
	SetStorageStatus(v string) *GetStorageListRequest
	GetStorageStatus() *string
	SetStorageType(v string) *GetStorageListRequest
	GetStorageType() *string
}

type GetStorageListRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Division             *string `json:"Division,omitempty" xml:"Division,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	StorageRegion        *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
	StorageStatus        *string `json:"StorageStatus,omitempty" xml:"StorageStatus,omitempty"`
	StorageType          *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s GetStorageListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageListRequest) GoString() string {
	return s.String()
}

func (s *GetStorageListRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetStorageListRequest) GetDivision() *string {
	return s.Division
}

func (s *GetStorageListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetStorageListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetStorageListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetStorageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetStorageListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetStorageListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetStorageListRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetStorageListRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetStorageListRequest) GetStorageRegion() *string {
	return s.StorageRegion
}

func (s *GetStorageListRequest) GetStorageStatus() *string {
	return s.StorageStatus
}

func (s *GetStorageListRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *GetStorageListRequest) SetAppId(v string) *GetStorageListRequest {
	s.AppId = &v
	return s
}

func (s *GetStorageListRequest) SetDivision(v string) *GetStorageListRequest {
	s.Division = &v
	return s
}

func (s *GetStorageListRequest) SetOwnerAccount(v string) *GetStorageListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetStorageListRequest) SetOwnerId(v string) *GetStorageListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetStorageListRequest) SetPageNumber(v int32) *GetStorageListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStorageListRequest) SetPageSize(v int32) *GetStorageListRequest {
	s.PageSize = &v
	return s
}

func (s *GetStorageListRequest) SetResourceGroupId(v string) *GetStorageListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetStorageListRequest) SetResourceOwnerAccount(v string) *GetStorageListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetStorageListRequest) SetResourceOwnerId(v string) *GetStorageListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetStorageListRequest) SetResourceRealOwnerId(v int64) *GetStorageListRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetStorageListRequest) SetStorageRegion(v string) *GetStorageListRequest {
	s.StorageRegion = &v
	return s
}

func (s *GetStorageListRequest) SetStorageStatus(v string) *GetStorageListRequest {
	s.StorageStatus = &v
	return s
}

func (s *GetStorageListRequest) SetStorageType(v string) *GetStorageListRequest {
	s.StorageType = &v
	return s
}

func (s *GetStorageListRequest) Validate() error {
	return dara.Validate(s)
}
