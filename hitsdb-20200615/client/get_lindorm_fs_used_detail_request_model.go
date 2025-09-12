// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormFsUsedDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormFsUsedDetailRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormFsUsedDetailRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormFsUsedDetailRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetLindormFsUsedDetailRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetLindormFsUsedDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormFsUsedDetailRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormFsUsedDetailRequest
	GetSecurityToken() *string
}

type GetLindormFsUsedDetailRequest struct {
	// The ID of the instance. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-xxxx
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormFsUsedDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormFsUsedDetailRequest) GoString() string {
	return s.String()
}

func (s *GetLindormFsUsedDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormFsUsedDetailRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormFsUsedDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormFsUsedDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormFsUsedDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormFsUsedDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormFsUsedDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormFsUsedDetailRequest) SetInstanceId(v string) *GetLindormFsUsedDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetOwnerAccount(v string) *GetLindormFsUsedDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetOwnerId(v int64) *GetLindormFsUsedDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetRegionId(v string) *GetLindormFsUsedDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetResourceOwnerAccount(v string) *GetLindormFsUsedDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetResourceOwnerId(v int64) *GetLindormFsUsedDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetSecurityToken(v string) *GetLindormFsUsedDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) Validate() error {
	return dara.Validate(s)
}
