// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDasInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDasInstanceConfigRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyDasInstanceConfigRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyDasInstanceConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDasInstanceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDasInstanceConfigRequest
	GetResourceOwnerId() *int64
	SetStorageAutoScale(v string) *ModifyDasInstanceConfigRequest
	GetStorageAutoScale() *string
	SetStorageThreshold(v int32) *ModifyDasInstanceConfigRequest
	GetStorageThreshold() *int32
	SetStorageUpperBound(v int32) *ModifyDasInstanceConfigRequest
	GetStorageUpperBound() *int32
}

type ModifyDasInstanceConfigRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StorageAutoScale  *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	StorageThreshold  *int32  `json:"StorageThreshold,omitempty" xml:"StorageThreshold,omitempty"`
	StorageUpperBound *int32  `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
}

func (s ModifyDasInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDasInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDasInstanceConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDasInstanceConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDasInstanceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDasInstanceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDasInstanceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDasInstanceConfigRequest) GetStorageAutoScale() *string {
	return s.StorageAutoScale
}

func (s *ModifyDasInstanceConfigRequest) GetStorageThreshold() *int32 {
	return s.StorageThreshold
}

func (s *ModifyDasInstanceConfigRequest) GetStorageUpperBound() *int32 {
	return s.StorageUpperBound
}

func (s *ModifyDasInstanceConfigRequest) SetClientToken(v string) *ModifyDasInstanceConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDasInstanceConfigRequest) SetDBInstanceId(v string) *ModifyDasInstanceConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDasInstanceConfigRequest) SetOwnerId(v int64) *ModifyDasInstanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDasInstanceConfigRequest) SetResourceOwnerAccount(v string) *ModifyDasInstanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDasInstanceConfigRequest) SetResourceOwnerId(v int64) *ModifyDasInstanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDasInstanceConfigRequest) SetStorageAutoScale(v string) *ModifyDasInstanceConfigRequest {
	s.StorageAutoScale = &v
	return s
}

func (s *ModifyDasInstanceConfigRequest) SetStorageThreshold(v int32) *ModifyDasInstanceConfigRequest {
	s.StorageThreshold = &v
	return s
}

func (s *ModifyDasInstanceConfigRequest) SetStorageUpperBound(v int32) *ModifyDasInstanceConfigRequest {
	s.StorageUpperBound = &v
	return s
}

func (s *ModifyDasInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
