// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeEngineJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetComputeEngineJobDetailRequest
	GetInstanceId() *string
	SetJobId(v string) *GetComputeEngineJobDetailRequest
	GetJobId() *string
	SetOwnerAccount(v string) *GetComputeEngineJobDetailRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetComputeEngineJobDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetComputeEngineJobDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetComputeEngineJobDetailRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetComputeEngineJobDetailRequest
	GetSecurityToken() *string
}

type GetComputeEngineJobDetailRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetComputeEngineJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEngineJobDetailRequest) GoString() string {
	return s.String()
}

func (s *GetComputeEngineJobDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetComputeEngineJobDetailRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetComputeEngineJobDetailRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetComputeEngineJobDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetComputeEngineJobDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetComputeEngineJobDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetComputeEngineJobDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetComputeEngineJobDetailRequest) SetInstanceId(v string) *GetComputeEngineJobDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetComputeEngineJobDetailRequest) SetJobId(v string) *GetComputeEngineJobDetailRequest {
	s.JobId = &v
	return s
}

func (s *GetComputeEngineJobDetailRequest) SetOwnerAccount(v string) *GetComputeEngineJobDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetComputeEngineJobDetailRequest) SetOwnerId(v int64) *GetComputeEngineJobDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetComputeEngineJobDetailRequest) SetResourceOwnerAccount(v string) *GetComputeEngineJobDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetComputeEngineJobDetailRequest) SetResourceOwnerId(v int64) *GetComputeEngineJobDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetComputeEngineJobDetailRequest) SetSecurityToken(v string) *GetComputeEngineJobDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetComputeEngineJobDetailRequest) Validate() error {
	return dara.Validate(s)
}
