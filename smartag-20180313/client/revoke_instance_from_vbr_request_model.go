// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromVbrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *RevokeInstanceFromVbrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeInstanceFromVbrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RevokeInstanceFromVbrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RevokeInstanceFromVbrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeInstanceFromVbrRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *RevokeInstanceFromVbrRequest
	GetSmartAGId() *string
	SetVbrInstanceId(v string) *RevokeInstanceFromVbrRequest
	GetVbrInstanceId() *string
}

type RevokeInstanceFromVbrRequest struct {
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
}

func (s RevokeInstanceFromVbrRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromVbrRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromVbrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeInstanceFromVbrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeInstanceFromVbrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeInstanceFromVbrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeInstanceFromVbrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeInstanceFromVbrRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *RevokeInstanceFromVbrRequest) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *RevokeInstanceFromVbrRequest) SetOwnerAccount(v string) *RevokeInstanceFromVbrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetOwnerId(v int64) *RevokeInstanceFromVbrRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetRegionId(v string) *RevokeInstanceFromVbrRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetResourceOwnerAccount(v string) *RevokeInstanceFromVbrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetResourceOwnerId(v int64) *RevokeInstanceFromVbrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetSmartAGId(v string) *RevokeInstanceFromVbrRequest {
	s.SmartAGId = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetVbrInstanceId(v string) *RevokeInstanceFromVbrRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) Validate() error {
	return dara.Validate(s)
}
