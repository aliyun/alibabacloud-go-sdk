// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeSlotsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeSlotsRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeSlotsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSlotsRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeSlotsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSlotsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSlotsRequest
	GetResourceOwnerId() *int64
}

type DescribeSlotsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can leave this parameter empty.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSlotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlotsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSlotsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSlotsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSlotsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSlotsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlotsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSlotsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSlotsRequest) SetClientToken(v string) *DescribeSlotsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSlotsRequest) SetDBInstanceId(v string) *DescribeSlotsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlotsRequest) SetOwnerAccount(v string) *DescribeSlotsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlotsRequest) SetOwnerId(v int64) *DescribeSlotsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlotsRequest) SetResourceGroupId(v string) *DescribeSlotsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlotsRequest) SetResourceOwnerAccount(v string) *DescribeSlotsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlotsRequest) SetResourceOwnerId(v int64) *DescribeSlotsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlotsRequest) Validate() error {
	return dara.Validate(s)
}
