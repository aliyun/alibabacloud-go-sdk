// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySrvNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionType(v string) *ModifySrvNetworkAddressRequest
	GetConnectionType() *string
	SetDBInstanceId(v string) *ModifySrvNetworkAddressRequest
	GetDBInstanceId() *string
	SetNewConnectionString(v string) *ModifySrvNetworkAddressRequest
	GetNewConnectionString() *string
	SetOwnerAccount(v string) *ModifySrvNetworkAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySrvNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySrvNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySrvNetworkAddressRequest
	GetResourceOwnerId() *int64
}

type ModifySrvNetworkAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1fd530f271****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aliyuntest111
	NewConnectionString  *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifySrvNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySrvNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifySrvNetworkAddressRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *ModifySrvNetworkAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySrvNetworkAddressRequest) GetNewConnectionString() *string {
	return s.NewConnectionString
}

func (s *ModifySrvNetworkAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySrvNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySrvNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySrvNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySrvNetworkAddressRequest) SetConnectionType(v string) *ModifySrvNetworkAddressRequest {
	s.ConnectionType = &v
	return s
}

func (s *ModifySrvNetworkAddressRequest) SetDBInstanceId(v string) *ModifySrvNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySrvNetworkAddressRequest) SetNewConnectionString(v string) *ModifySrvNetworkAddressRequest {
	s.NewConnectionString = &v
	return s
}

func (s *ModifySrvNetworkAddressRequest) SetOwnerAccount(v string) *ModifySrvNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySrvNetworkAddressRequest) SetOwnerId(v int64) *ModifySrvNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySrvNetworkAddressRequest) SetResourceOwnerAccount(v string) *ModifySrvNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySrvNetworkAddressRequest) SetResourceOwnerId(v int64) *ModifySrvNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySrvNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
