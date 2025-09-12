// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteCustomResourceRequest
	GetInstanceId() *string
	SetName(v string) *DeleteCustomResourceRequest
	GetName() *string
	SetOwnerAccount(v string) *DeleteCustomResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCustomResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCustomResourceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteCustomResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCustomResourceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteCustomResourceRequest
	GetSecurityToken() *string
}

type DeleteCustomResourceRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteCustomResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCustomResourceRequest) GetName() *string {
	return s.Name
}

func (s *DeleteCustomResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCustomResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCustomResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCustomResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCustomResourceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteCustomResourceRequest) SetInstanceId(v string) *DeleteCustomResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetName(v string) *DeleteCustomResourceRequest {
	s.Name = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetOwnerAccount(v string) *DeleteCustomResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetOwnerId(v int64) *DeleteCustomResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetRegionId(v string) *DeleteCustomResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetResourceOwnerAccount(v string) *DeleteCustomResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetResourceOwnerId(v int64) *DeleteCustomResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetSecurityToken(v string) *DeleteCustomResourceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteCustomResourceRequest) Validate() error {
	return dara.Validate(s)
}
