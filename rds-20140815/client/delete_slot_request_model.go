// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSlotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSlotRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DeleteSlotRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DeleteSlotRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSlotRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DeleteSlotRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteSlotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSlotRequest
	GetResourceOwnerId() *int64
	SetSlotName(v string) *DeleteSlotRequest
	GetSlotName() *string
	SetSlotStatus(v string) *DeleteSlotRequest
	GetSlotStatus() *string
}

type DeleteSlotRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SlotName *string `json:"SlotName,omitempty" xml:"SlotName,omitempty"`
	// This parameter is required.
	SlotStatus *string `json:"SlotStatus,omitempty" xml:"SlotStatus,omitempty"`
}

func (s DeleteSlotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSlotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSlotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSlotRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteSlotRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSlotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSlotRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteSlotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSlotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSlotRequest) GetSlotName() *string {
	return s.SlotName
}

func (s *DeleteSlotRequest) GetSlotStatus() *string {
	return s.SlotStatus
}

func (s *DeleteSlotRequest) SetClientToken(v string) *DeleteSlotRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSlotRequest) SetDBInstanceId(v string) *DeleteSlotRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteSlotRequest) SetOwnerAccount(v string) *DeleteSlotRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSlotRequest) SetOwnerId(v int64) *DeleteSlotRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSlotRequest) SetResourceGroupId(v string) *DeleteSlotRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteSlotRequest) SetResourceOwnerAccount(v string) *DeleteSlotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSlotRequest) SetResourceOwnerId(v int64) *DeleteSlotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSlotRequest) SetSlotName(v string) *DeleteSlotRequest {
	s.SlotName = &v
	return s
}

func (s *DeleteSlotRequest) SetSlotStatus(v string) *DeleteSlotRequest {
	s.SlotStatus = &v
	return s
}

func (s *DeleteSlotRequest) Validate() error {
	return dara.Validate(s)
}
