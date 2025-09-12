// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenComputeEngineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuLimit(v string) *OpenComputeEngineRequest
	GetCpuLimit() *string
	SetInstanceId(v string) *OpenComputeEngineRequest
	GetInstanceId() *string
	SetMemoryLimit(v string) *OpenComputeEngineRequest
	GetMemoryLimit() *string
	SetOwnerAccount(v string) *OpenComputeEngineRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenComputeEngineRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *OpenComputeEngineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenComputeEngineRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *OpenComputeEngineRequest
	GetSecurityToken() *string
}

type OpenComputeEngineRequest struct {
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

func (s OpenComputeEngineRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenComputeEngineRequest) GoString() string {
	return s.String()
}

func (s *OpenComputeEngineRequest) GetCpuLimit() *string {
	return s.CpuLimit
}

func (s *OpenComputeEngineRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OpenComputeEngineRequest) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *OpenComputeEngineRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenComputeEngineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenComputeEngineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenComputeEngineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenComputeEngineRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *OpenComputeEngineRequest) SetCpuLimit(v string) *OpenComputeEngineRequest {
	s.CpuLimit = &v
	return s
}

func (s *OpenComputeEngineRequest) SetInstanceId(v string) *OpenComputeEngineRequest {
	s.InstanceId = &v
	return s
}

func (s *OpenComputeEngineRequest) SetMemoryLimit(v string) *OpenComputeEngineRequest {
	s.MemoryLimit = &v
	return s
}

func (s *OpenComputeEngineRequest) SetOwnerAccount(v string) *OpenComputeEngineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenComputeEngineRequest) SetOwnerId(v int64) *OpenComputeEngineRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenComputeEngineRequest) SetResourceOwnerAccount(v string) *OpenComputeEngineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenComputeEngineRequest) SetResourceOwnerId(v int64) *OpenComputeEngineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenComputeEngineRequest) SetSecurityToken(v string) *OpenComputeEngineRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenComputeEngineRequest) Validate() error {
	return dara.Validate(s)
}
