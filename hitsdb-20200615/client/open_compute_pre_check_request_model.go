// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenComputePreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuLimit(v string) *OpenComputePreCheckRequest
	GetCpuLimit() *string
	SetInstanceId(v string) *OpenComputePreCheckRequest
	GetInstanceId() *string
	SetMemoryLimit(v string) *OpenComputePreCheckRequest
	GetMemoryLimit() *string
	SetOwnerAccount(v string) *OpenComputePreCheckRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenComputePreCheckRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *OpenComputePreCheckRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenComputePreCheckRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *OpenComputePreCheckRequest
	GetSecurityToken() *string
}

type OpenComputePreCheckRequest struct {
	CpuLimit *string `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemoryLimit          *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OpenComputePreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenComputePreCheckRequest) GoString() string {
	return s.String()
}

func (s *OpenComputePreCheckRequest) GetCpuLimit() *string {
	return s.CpuLimit
}

func (s *OpenComputePreCheckRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OpenComputePreCheckRequest) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *OpenComputePreCheckRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenComputePreCheckRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenComputePreCheckRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenComputePreCheckRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenComputePreCheckRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *OpenComputePreCheckRequest) SetCpuLimit(v string) *OpenComputePreCheckRequest {
	s.CpuLimit = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetInstanceId(v string) *OpenComputePreCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetMemoryLimit(v string) *OpenComputePreCheckRequest {
	s.MemoryLimit = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetOwnerAccount(v string) *OpenComputePreCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetOwnerId(v int64) *OpenComputePreCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetResourceOwnerAccount(v string) *OpenComputePreCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetResourceOwnerId(v int64) *OpenComputePreCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetSecurityToken(v string) *OpenComputePreCheckRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenComputePreCheckRequest) Validate() error {
	return dara.Validate(s)
}
