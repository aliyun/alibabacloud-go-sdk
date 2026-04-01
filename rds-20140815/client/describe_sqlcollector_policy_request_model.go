// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLCollectorPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeSQLCollectorPolicyRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeSQLCollectorPolicyRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeSQLCollectorPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLCollectorPolicyRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeSQLCollectorPolicyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSQLCollectorPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLCollectorPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribeSQLCollectorPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmyxxxxxxxxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQLCollectorPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSQLCollectorPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLCollectorPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLCollectorPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLCollectorPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSQLCollectorPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLCollectorPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLCollectorPolicyRequest) SetClientToken(v string) *DescribeSQLCollectorPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSQLCollectorPolicyRequest) SetDBInstanceId(v string) *DescribeSQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLCollectorPolicyRequest) SetOwnerAccount(v string) *DescribeSQLCollectorPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLCollectorPolicyRequest) SetOwnerId(v int64) *DescribeSQLCollectorPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLCollectorPolicyRequest) SetResourceGroupId(v string) *DescribeSQLCollectorPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSQLCollectorPolicyRequest) SetResourceOwnerAccount(v string) *DescribeSQLCollectorPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLCollectorPolicyRequest) SetResourceOwnerId(v int64) *DescribeSQLCollectorPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLCollectorPolicyRequest) Validate() error {
	return dara.Validate(s)
}
