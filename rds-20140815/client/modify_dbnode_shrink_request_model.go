// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDBNodeShrinkRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyDBNodeShrinkRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyDBNodeShrinkRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v string) *ModifyDBNodeShrinkRequest
	GetDBInstanceStorage() *string
	SetDBInstanceStorageType(v string) *ModifyDBNodeShrinkRequest
	GetDBInstanceStorageType() *string
	SetDBNodeShrink(v string) *ModifyDBNodeShrinkRequest
	GetDBNodeShrink() *string
	SetDryRun(v bool) *ModifyDBNodeShrinkRequest
	GetDryRun() *bool
	SetEffectiveTime(v string) *ModifyDBNodeShrinkRequest
	GetEffectiveTime() *string
	SetOwnerAccount(v string) *ModifyDBNodeShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeShrinkRequest
	GetOwnerId() *int64
	SetProduceAsync(v bool) *ModifyDBNodeShrinkRequest
	GetProduceAsync() *bool
	SetResourceOwnerAccount(v string) *ModifyDBNodeShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyDBNodeShrinkRequest struct {
	AutoPay     *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStorage     *string `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	DBNodeShrink          *string `json:"DBNode,omitempty" xml:"DBNode,omitempty"`
	DryRun                *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EffectiveTime         *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProduceAsync          *bool   `json:"ProduceAsync,omitempty" xml:"ProduceAsync,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBNodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDBNodeShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBNodeShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBNodeShrinkRequest) GetDBInstanceStorage() *string {
	return s.DBInstanceStorage
}

func (s *ModifyDBNodeShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBNodeShrinkRequest) GetDBNodeShrink() *string {
	return s.DBNodeShrink
}

func (s *ModifyDBNodeShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDBNodeShrinkRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBNodeShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeShrinkRequest) GetProduceAsync() *bool {
	return s.ProduceAsync
}

func (s *ModifyDBNodeShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeShrinkRequest) SetAutoPay(v bool) *ModifyDBNodeShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetClientToken(v string) *ModifyDBNodeShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDBInstanceId(v string) *ModifyDBNodeShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDBInstanceStorage(v string) *ModifyDBNodeShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDBInstanceStorageType(v string) *ModifyDBNodeShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDBNodeShrink(v string) *ModifyDBNodeShrinkRequest {
	s.DBNodeShrink = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDryRun(v bool) *ModifyDBNodeShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetEffectiveTime(v string) *ModifyDBNodeShrinkRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetOwnerAccount(v string) *ModifyDBNodeShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetOwnerId(v int64) *ModifyDBNodeShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetProduceAsync(v bool) *ModifyDBNodeShrinkRequest {
	s.ProduceAsync = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBNodeShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
