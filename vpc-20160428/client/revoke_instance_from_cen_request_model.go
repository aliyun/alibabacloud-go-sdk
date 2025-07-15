// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *RevokeInstanceFromCenRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *RevokeInstanceFromCenRequest
	GetCenOwnerId() *int64
	SetClientToken(v string) *RevokeInstanceFromCenRequest
	GetClientToken() *string
	SetInstanceId(v string) *RevokeInstanceFromCenRequest
	GetInstanceId() *string
	SetInstanceType(v string) *RevokeInstanceFromCenRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *RevokeInstanceFromCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeInstanceFromCenRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RevokeInstanceFromCenRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RevokeInstanceFromCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeInstanceFromCenRequest
	GetResourceOwnerId() *int64
}

type RevokeInstanceFromCenRequest struct {
	// The ID of the CEN instance to which the network instance is attached.
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
	// 	- **CCN**
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

func (s RevokeInstanceFromCenRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromCenRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *RevokeInstanceFromCenRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *RevokeInstanceFromCenRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RevokeInstanceFromCenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeInstanceFromCenRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RevokeInstanceFromCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeInstanceFromCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeInstanceFromCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeInstanceFromCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeInstanceFromCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeInstanceFromCenRequest) SetCenId(v string) *RevokeInstanceFromCenRequest {
	s.CenId = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetCenOwnerId(v int64) *RevokeInstanceFromCenRequest {
	s.CenOwnerId = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetClientToken(v string) *RevokeInstanceFromCenRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetInstanceId(v string) *RevokeInstanceFromCenRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetInstanceType(v string) *RevokeInstanceFromCenRequest {
	s.InstanceType = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetOwnerAccount(v string) *RevokeInstanceFromCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetOwnerId(v int64) *RevokeInstanceFromCenRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetRegionId(v string) *RevokeInstanceFromCenRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetResourceOwnerAccount(v string) *RevokeInstanceFromCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) SetResourceOwnerId(v int64) *RevokeInstanceFromCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeInstanceFromCenRequest) Validate() error {
	return dara.Validate(s)
}
