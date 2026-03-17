// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSagInstanceToVbrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GrantSagInstanceToVbrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantSagInstanceToVbrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GrantSagInstanceToVbrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GrantSagInstanceToVbrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GrantSagInstanceToVbrRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *GrantSagInstanceToVbrRequest
	GetSmartAGId() *string
	SetVbrInstanceId(v string) *GrantSagInstanceToVbrRequest
	GetVbrInstanceId() *string
	SetVbrRegionId(v string) *GrantSagInstanceToVbrRequest
	GetVbrRegionId() *string
	SetVbrUid(v int64) *GrantSagInstanceToVbrRequest
	GetVbrUid() *int64
}

type GrantSagInstanceToVbrRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-0nnteglltw6z4b****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp13gtbhdp0pfqg6s****
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	// The ID of the region where the VBR is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionId *string `json:"VbrRegionId,omitempty" xml:"VbrRegionId,omitempty"`
	// The user ID (UID) of the account to which the VBR belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231212121121212112
	VbrUid *int64 `json:"VbrUid,omitempty" xml:"VbrUid,omitempty"`
}

func (s GrantSagInstanceToVbrRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantSagInstanceToVbrRequest) GoString() string {
	return s.String()
}

func (s *GrantSagInstanceToVbrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantSagInstanceToVbrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantSagInstanceToVbrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantSagInstanceToVbrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GrantSagInstanceToVbrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantSagInstanceToVbrRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *GrantSagInstanceToVbrRequest) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *GrantSagInstanceToVbrRequest) GetVbrRegionId() *string {
	return s.VbrRegionId
}

func (s *GrantSagInstanceToVbrRequest) GetVbrUid() *int64 {
	return s.VbrUid
}

func (s *GrantSagInstanceToVbrRequest) SetOwnerAccount(v string) *GrantSagInstanceToVbrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) SetOwnerId(v int64) *GrantSagInstanceToVbrRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) SetRegionId(v string) *GrantSagInstanceToVbrRequest {
	s.RegionId = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) SetResourceOwnerAccount(v string) *GrantSagInstanceToVbrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) SetResourceOwnerId(v int64) *GrantSagInstanceToVbrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) SetSmartAGId(v string) *GrantSagInstanceToVbrRequest {
	s.SmartAGId = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) SetVbrInstanceId(v string) *GrantSagInstanceToVbrRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) SetVbrRegionId(v string) *GrantSagInstanceToVbrRequest {
	s.VbrRegionId = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) SetVbrUid(v int64) *GrantSagInstanceToVbrRequest {
	s.VbrUid = &v
	return s
}

func (s *GrantSagInstanceToVbrRequest) Validate() error {
	return dara.Validate(s)
}
