// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDBInstanceSrvNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *AllocateDBInstanceSrvNetworkAddressRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *AllocateDBInstanceSrvNetworkAddressRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *AllocateDBInstanceSrvNetworkAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateDBInstanceSrvNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AllocateDBInstanceSrvNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateDBInstanceSrvNetworkAddressRequest
	GetResourceOwnerId() *int64
	SetSrvConnectionType(v string) *AllocateDBInstanceSrvNetworkAddressRequest
	GetSrvConnectionType() *string
}

type AllocateDBInstanceSrvNetworkAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dds-2ze5eb9514e31364
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// d-bp1b7bb3bbe****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// vpc
	SrvConnectionType *string `json:"SrvConnectionType,omitempty" xml:"SrvConnectionType,omitempty"`
}

func (s AllocateDBInstanceSrvNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateDBInstanceSrvNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) GetSrvConnectionType() *string {
	return s.SrvConnectionType
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) SetDBInstanceId(v string) *AllocateDBInstanceSrvNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) SetNodeId(v string) *AllocateDBInstanceSrvNetworkAddressRequest {
	s.NodeId = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) SetOwnerAccount(v string) *AllocateDBInstanceSrvNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) SetOwnerId(v int64) *AllocateDBInstanceSrvNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) SetResourceOwnerAccount(v string) *AllocateDBInstanceSrvNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) SetResourceOwnerId(v int64) *AllocateDBInstanceSrvNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) SetSrvConnectionType(v string) *AllocateDBInstanceSrvNetworkAddressRequest {
	s.SrvConnectionType = &v
	return s
}

func (s *AllocateDBInstanceSrvNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
