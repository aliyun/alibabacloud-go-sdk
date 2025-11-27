// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeBurstConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeComputeBurstConfigRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeComputeBurstConfigRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeComputeBurstConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeComputeBurstConfigRequest
	GetOwnerId() *int64
	SetRegionId(v []byte) *DescribeComputeBurstConfigRequest
	GetRegionId() []byte
	SetResourceGroupId(v string) *DescribeComputeBurstConfigRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeComputeBurstConfigRequest
	GetResourceOwnerAccount() *string
}

type DescribeComputeBurstConfigRequest struct {
	// The client token that is used to ensure the idempotence of requests and prevent repeated requests from being submitted. You can use the client to generate the value, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-2ze63v2p3o3k****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId []byte `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeComputeBurstConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeBurstConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeComputeBurstConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeComputeBurstConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeComputeBurstConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeComputeBurstConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeComputeBurstConfigRequest) GetRegionId() []byte {
	return s.RegionId
}

func (s *DescribeComputeBurstConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeComputeBurstConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeComputeBurstConfigRequest) SetClientToken(v string) *DescribeComputeBurstConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeComputeBurstConfigRequest) SetDBInstanceId(v string) *DescribeComputeBurstConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeComputeBurstConfigRequest) SetOwnerAccount(v string) *DescribeComputeBurstConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeComputeBurstConfigRequest) SetOwnerId(v int64) *DescribeComputeBurstConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeComputeBurstConfigRequest) SetRegionId(v []byte) *DescribeComputeBurstConfigRequest {
	s.RegionId = v
	return s
}

func (s *DescribeComputeBurstConfigRequest) SetResourceGroupId(v string) *DescribeComputeBurstConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeComputeBurstConfigRequest) SetResourceOwnerAccount(v string) *DescribeComputeBurstConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeComputeBurstConfigRequest) Validate() error {
	return dara.Validate(s)
}
