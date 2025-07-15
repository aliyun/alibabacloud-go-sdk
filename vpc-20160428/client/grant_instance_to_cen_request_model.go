// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *GrantInstanceToCenRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *GrantInstanceToCenRequest
	GetCenOwnerId() *int64
	SetClientToken(v string) *GrantInstanceToCenRequest
	GetClientToken() *string
	SetInstanceId(v string) *GrantInstanceToCenRequest
	GetInstanceId() *string
	SetInstanceType(v string) *GrantInstanceToCenRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *GrantInstanceToCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantInstanceToCenRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GrantInstanceToCenRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GrantInstanceToCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GrantInstanceToCenRequest
	GetResourceOwnerId() *int64
}

type GrantInstanceToCenRequest struct {
	// The ID of the CEN instance to which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The user ID (UID) of the Apsara Stack tenant account to which the CEN instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6o8d1dj8sjwxi6o****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VBR**
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantInstanceToCenRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToCenRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *GrantInstanceToCenRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *GrantInstanceToCenRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GrantInstanceToCenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GrantInstanceToCenRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GrantInstanceToCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantInstanceToCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantInstanceToCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantInstanceToCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GrantInstanceToCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantInstanceToCenRequest) SetCenId(v string) *GrantInstanceToCenRequest {
	s.CenId = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetCenOwnerId(v int64) *GrantInstanceToCenRequest {
	s.CenOwnerId = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetClientToken(v string) *GrantInstanceToCenRequest {
	s.ClientToken = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetInstanceId(v string) *GrantInstanceToCenRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetInstanceType(v string) *GrantInstanceToCenRequest {
	s.InstanceType = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetOwnerAccount(v string) *GrantInstanceToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetOwnerId(v int64) *GrantInstanceToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetRegionId(v string) *GrantInstanceToCenRequest {
	s.RegionId = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetResourceOwnerAccount(v string) *GrantInstanceToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantInstanceToCenRequest) SetResourceOwnerId(v int64) *GrantInstanceToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantInstanceToCenRequest) Validate() error {
	return dara.Validate(s)
}
