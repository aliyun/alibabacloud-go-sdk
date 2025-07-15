// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletionProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeletionProtectionRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeletionProtectionRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DeletionProtectionRequest
	GetOwnerId() *int64
	SetProtectionEnable(v bool) *DeletionProtectionRequest
	GetProtectionEnable() *bool
	SetRegionId(v string) *DeletionProtectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeletionProtectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletionProtectionRequest
	GetResourceOwnerId() *int64
	SetType(v string) *DeletionProtectionRequest
	GetType() *string
}

type DeletionProtectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance for which you want to enable deletion protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-uf6id7gvguruxe41v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable deletion protection. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ProtectionEnable *bool `json:"ProtectionEnable,omitempty" xml:"ProtectionEnable,omitempty"`
	// The ID of the region where the instance for which you want to enable deletion protection is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of instance. Valid values:
	//
	// 	- **EIP**: Elastic IP Address (EIP)
	//
	// 	- **CBWP**: EIP bandwidth plan
	//
	// 	- **NATGW**: NAT gateway
	//
	// This parameter is required.
	//
	// example:
	//
	// EIP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeletionProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *DeletionProtectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeletionProtectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeletionProtectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletionProtectionRequest) GetProtectionEnable() *bool {
	return s.ProtectionEnable
}

func (s *DeletionProtectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletionProtectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletionProtectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletionProtectionRequest) GetType() *string {
	return s.Type
}

func (s *DeletionProtectionRequest) SetClientToken(v string) *DeletionProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeletionProtectionRequest) SetInstanceId(v string) *DeletionProtectionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeletionProtectionRequest) SetOwnerId(v int64) *DeletionProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletionProtectionRequest) SetProtectionEnable(v bool) *DeletionProtectionRequest {
	s.ProtectionEnable = &v
	return s
}

func (s *DeletionProtectionRequest) SetRegionId(v string) *DeletionProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *DeletionProtectionRequest) SetResourceOwnerAccount(v string) *DeletionProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletionProtectionRequest) SetResourceOwnerId(v int64) *DeletionProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletionProtectionRequest) SetType(v string) *DeletionProtectionRequest {
	s.Type = &v
	return s
}

func (s *DeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
