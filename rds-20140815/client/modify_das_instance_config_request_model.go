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
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and can be up to 64 ASCII characters in length. The value cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Instance ID of the instance. You can call the DescribeDBInstances operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable automatic storage expansion. Valid values:
	//
	// 	- **Enable**: Enables automatic storage expansion.
	//
	// 	- **Disable**: Disables automatic storage expansion.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The threshold percentage that triggers automatic storage expansion. When the percentage of remaining storage space reaches the specified threshold, automatic storage expansion is triggered. Valid values:
	//
	// 	- **10**
	//
	// 	- **20**
	//
	// 	- **30**
	//
	// 	- **40**
	//
	// 	- **50**
	//
	// >This parameter is required when **StorageAutoScale*	- is set to **Enable**.
	//
	// example:
	//
	// 50
	StorageThreshold *int32 `json:"StorageThreshold,omitempty" xml:"StorageThreshold,omitempty"`
	// The upper limit for automatic storage expansion. The value must be greater than or equal to the current total storage capacity of the instance.
	//
	// - Upper limit for ESSDs: 32,000 GB
	//
	// - Upper limit for standard SSDs: 6,000 GB
	//
	// >This parameter is required when **StorageAutoScale*	- is set to **Enable**.
	//
	// example:
	//
	// 1000
	StorageUpperBound *int32 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
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
