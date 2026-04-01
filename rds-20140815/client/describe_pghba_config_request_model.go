// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePGHbaConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribePGHbaConfigRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribePGHbaConfigRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribePGHbaConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePGHbaConfigRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribePGHbaConfigRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePGHbaConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePGHbaConfigRequest
	GetResourceOwnerId() *int64
}

type DescribePGHbaConfigRequest struct {
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// 1
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1lymyn1v3i****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePGHbaConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePGHbaConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePGHbaConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePGHbaConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribePGHbaConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePGHbaConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePGHbaConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePGHbaConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePGHbaConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePGHbaConfigRequest) SetClientToken(v string) *DescribePGHbaConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePGHbaConfigRequest) SetDBInstanceId(v string) *DescribePGHbaConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePGHbaConfigRequest) SetOwnerAccount(v string) *DescribePGHbaConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePGHbaConfigRequest) SetOwnerId(v int64) *DescribePGHbaConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePGHbaConfigRequest) SetResourceGroupId(v string) *DescribePGHbaConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePGHbaConfigRequest) SetResourceOwnerAccount(v string) *DescribePGHbaConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePGHbaConfigRequest) SetResourceOwnerId(v int64) *DescribePGHbaConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePGHbaConfigRequest) Validate() error {
	return dara.Validate(s)
}
