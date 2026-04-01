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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable automatic storage expansion. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disable**
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The threshold in percentage based on which an automatic storage expansion is triggered. If the available storage reaches the threshold, ApsaraDB RDS increases the storage capacity of the instance. Valid values:
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
	// >  If you set the StorageAutoScale parameter to **Enable**, you must specify this parameter.
	//
	// example:
	//
	// 50
	StorageThreshold *int32 `json:"StorageThreshold,omitempty" xml:"StorageThreshold,omitempty"`
	// The maximum storage capacity that is allowed for an automatic storage expansion. The value of this parameter must be greater than or equal to the current storage capacity of the RDS instance.
	//
	// 	- If the RDS instance uses ESSDs, the maximum value of this parameter can be set to 32000 GB.
	//
	// 	- If the RDS instance uses standard SSDs, the maximum value of this parameter can be set to 6000 GB.
	//
	// >  If you set the **StorageAutoScale*	- parameter to **Enable**, you must specify this parameter.
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
