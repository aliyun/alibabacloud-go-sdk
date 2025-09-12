// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsNamespacedQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLdpsNamespacedQuotaRequest
	GetInstanceId() *string
	SetNamespace(v string) *GetLdpsNamespacedQuotaRequest
	GetNamespace() *string
	SetOwnerAccount(v string) *GetLdpsNamespacedQuotaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLdpsNamespacedQuotaRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetLdpsNamespacedQuotaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetLdpsNamespacedQuotaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLdpsNamespacedQuotaRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLdpsNamespacedQuotaRequest
	GetSecurityToken() *string
}

type GetLdpsNamespacedQuotaRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Namespace            *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLdpsNamespacedQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsNamespacedQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetLdpsNamespacedQuotaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLdpsNamespacedQuotaRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLdpsNamespacedQuotaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLdpsNamespacedQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLdpsNamespacedQuotaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLdpsNamespacedQuotaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLdpsNamespacedQuotaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLdpsNamespacedQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLdpsNamespacedQuotaRequest) SetInstanceId(v string) *GetLdpsNamespacedQuotaRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetNamespace(v string) *GetLdpsNamespacedQuotaRequest {
	s.Namespace = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetOwnerAccount(v string) *GetLdpsNamespacedQuotaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetOwnerId(v int64) *GetLdpsNamespacedQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetRegionId(v string) *GetLdpsNamespacedQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetResourceOwnerAccount(v string) *GetLdpsNamespacedQuotaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetResourceOwnerId(v int64) *GetLdpsNamespacedQuotaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetSecurityToken(v string) *GetLdpsNamespacedQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) Validate() error {
	return dara.Validate(s)
}
