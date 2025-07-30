// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeTairKVCacheCustomInstanceDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetAutoPay() *bool
	SetDiskId(v string) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetDiskId() *string
	SetDiskSize(v string) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetDiskSize() *string
	SetInstanceId(v string) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ResizeTairKVCacheCustomInstanceDiskRequest
	GetSecurityToken() *string
}

type ResizeTairKVCacheCustomInstanceDiskRequest struct {
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d-5v1aggi3ffoxufb57**
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// example:
	//
	// 5000
	DiskSize *string `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tc-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResizeTairKVCacheCustomInstanceDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeTairKVCacheCustomInstanceDiskRequest) GoString() string {
	return s.String()
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetDiskSize() *string {
	return s.DiskSize
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetAutoPay(v bool) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.AutoPay = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetDiskId(v string) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetDiskSize(v string) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.DiskSize = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetInstanceId(v string) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetOwnerAccount(v string) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetOwnerId(v int64) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetResourceOwnerAccount(v string) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetResourceOwnerId(v int64) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) SetSecurityToken(v string) *ResizeTairKVCacheCustomInstanceDiskRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResizeTairKVCacheCustomInstanceDiskRequest) Validate() error {
	return dara.Validate(s)
}
