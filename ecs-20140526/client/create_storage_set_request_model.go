// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateStorageSetRequest
	GetClientToken() *string
	SetDescription(v string) *CreateStorageSetRequest
	GetDescription() *string
	SetMaxPartitionNumber(v int32) *CreateStorageSetRequest
	GetMaxPartitionNumber() *int32
	SetOwnerAccount(v string) *CreateStorageSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateStorageSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateStorageSetRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateStorageSetRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateStorageSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateStorageSetRequest
	GetResourceOwnerId() *int64
	SetStorageSetName(v string) *CreateStorageSetRequest
	GetStorageSetName() *string
	SetTag(v []*CreateStorageSetRequestTag) *CreateStorageSetRequest
	GetTag() []*CreateStorageSetRequestTag
	SetZoneId(v string) *CreateStorageSetRequest
	GetZoneId() *string
}

type CreateStorageSetRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxPartitionNumber *int32  `json:"MaxPartitionNumber,omitempty" xml:"MaxPartitionNumber,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StorageSetName       *string                       `json:"StorageSetName,omitempty" xml:"StorageSetName,omitempty"`
	Tag                  []*CreateStorageSetRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateStorageSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageSetRequest) GoString() string {
	return s.String()
}

func (s *CreateStorageSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStorageSetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStorageSetRequest) GetMaxPartitionNumber() *int32 {
	return s.MaxPartitionNumber
}

func (s *CreateStorageSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateStorageSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateStorageSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStorageSetRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateStorageSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateStorageSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateStorageSetRequest) GetStorageSetName() *string {
	return s.StorageSetName
}

func (s *CreateStorageSetRequest) GetTag() []*CreateStorageSetRequestTag {
	return s.Tag
}

func (s *CreateStorageSetRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateStorageSetRequest) SetClientToken(v string) *CreateStorageSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStorageSetRequest) SetDescription(v string) *CreateStorageSetRequest {
	s.Description = &v
	return s
}

func (s *CreateStorageSetRequest) SetMaxPartitionNumber(v int32) *CreateStorageSetRequest {
	s.MaxPartitionNumber = &v
	return s
}

func (s *CreateStorageSetRequest) SetOwnerAccount(v string) *CreateStorageSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateStorageSetRequest) SetOwnerId(v int64) *CreateStorageSetRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateStorageSetRequest) SetRegionId(v string) *CreateStorageSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStorageSetRequest) SetResourceGroupId(v string) *CreateStorageSetRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStorageSetRequest) SetResourceOwnerAccount(v string) *CreateStorageSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateStorageSetRequest) SetResourceOwnerId(v int64) *CreateStorageSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateStorageSetRequest) SetStorageSetName(v string) *CreateStorageSetRequest {
	s.StorageSetName = &v
	return s
}

func (s *CreateStorageSetRequest) SetTag(v []*CreateStorageSetRequestTag) *CreateStorageSetRequest {
	s.Tag = v
	return s
}

func (s *CreateStorageSetRequest) SetZoneId(v string) *CreateStorageSetRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateStorageSetRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateStorageSetRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStorageSetRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageSetRequestTag) GoString() string {
	return s.String()
}

func (s *CreateStorageSetRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateStorageSetRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateStorageSetRequestTag) SetKey(v string) *CreateStorageSetRequestTag {
	s.Key = &v
	return s
}

func (s *CreateStorageSetRequestTag) SetValue(v string) *CreateStorageSetRequestTag {
	s.Value = &v
	return s
}

func (s *CreateStorageSetRequestTag) Validate() error {
	return dara.Validate(s)
}
