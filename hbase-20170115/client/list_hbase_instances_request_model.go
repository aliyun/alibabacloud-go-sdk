// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHbaseInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ListHbaseInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListHbaseInstancesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListHbaseInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListHbaseInstancesRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ListHbaseInstancesRequest
	GetSecurityToken() *string
	SetVpcId(v string) *ListHbaseInstancesRequest
	GetVpcId() *string
	SetZoneId(v string) *ListHbaseInstancesRequest
	GetZoneId() *string
}

type ListHbaseInstancesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListHbaseInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHbaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListHbaseInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListHbaseInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListHbaseInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListHbaseInstancesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListHbaseInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListHbaseInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListHbaseInstancesRequest) SetOwnerAccount(v string) *ListHbaseInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetOwnerId(v int64) *ListHbaseInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetResourceOwnerAccount(v string) *ListHbaseInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetResourceOwnerId(v int64) *ListHbaseInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetSecurityToken(v string) *ListHbaseInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetVpcId(v string) *ListHbaseInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *ListHbaseInstancesRequest) SetZoneId(v string) *ListHbaseInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *ListHbaseInstancesRequest) Validate() error {
	return dara.Validate(s)
}
