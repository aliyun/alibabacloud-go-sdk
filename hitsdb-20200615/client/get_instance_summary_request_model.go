// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetInstanceSummaryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetInstanceSummaryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetInstanceSummaryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetInstanceSummaryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetInstanceSummaryRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetInstanceSummaryRequest
	GetSecurityToken() *string
}

type GetInstanceSummaryRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetInstanceSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceSummaryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetInstanceSummaryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetInstanceSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceSummaryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetInstanceSummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetInstanceSummaryRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetInstanceSummaryRequest) SetOwnerAccount(v string) *GetInstanceSummaryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetInstanceSummaryRequest) SetOwnerId(v int64) *GetInstanceSummaryRequest {
	s.OwnerId = &v
	return s
}

func (s *GetInstanceSummaryRequest) SetRegionId(v string) *GetInstanceSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceSummaryRequest) SetResourceOwnerAccount(v string) *GetInstanceSummaryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetInstanceSummaryRequest) SetResourceOwnerId(v int64) *GetInstanceSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetInstanceSummaryRequest) SetSecurityToken(v string) *GetInstanceSummaryRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetInstanceSummaryRequest) Validate() error {
	return dara.Validate(s)
}
