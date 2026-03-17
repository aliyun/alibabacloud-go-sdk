// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindVbrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *UnbindVbrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnbindVbrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnbindVbrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnbindVbrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindVbrRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *UnbindVbrRequest
	GetSmartAGId() *string
	SetSmartAGUid(v int64) *UnbindVbrRequest
	GetSmartAGUid() *int64
	SetVbrId(v string) *UnbindVbrRequest
	GetVbrId() *string
	SetVbrRegionId(v string) *UnbindVbrRequest
	GetVbrRegionId() *string
}

type UnbindVbrRequest struct {
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
	// sag-pek29fu47tmpj0****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 253460731706911258
	SmartAGUid *int64 `json:"SmartAGUid,omitempty" xml:"SmartAGUid,omitempty"`
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-wz90rs6ef1m2fq0yn****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The ID of the region where the VBR is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionId *string `json:"VbrRegionId,omitempty" xml:"VbrRegionId,omitempty"`
}

func (s UnbindVbrRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindVbrRequest) GoString() string {
	return s.String()
}

func (s *UnbindVbrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnbindVbrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindVbrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnbindVbrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindVbrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindVbrRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UnbindVbrRequest) GetSmartAGUid() *int64 {
	return s.SmartAGUid
}

func (s *UnbindVbrRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *UnbindVbrRequest) GetVbrRegionId() *string {
	return s.VbrRegionId
}

func (s *UnbindVbrRequest) SetOwnerAccount(v string) *UnbindVbrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindVbrRequest) SetOwnerId(v int64) *UnbindVbrRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindVbrRequest) SetRegionId(v string) *UnbindVbrRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindVbrRequest) SetResourceOwnerAccount(v string) *UnbindVbrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindVbrRequest) SetResourceOwnerId(v int64) *UnbindVbrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindVbrRequest) SetSmartAGId(v string) *UnbindVbrRequest {
	s.SmartAGId = &v
	return s
}

func (s *UnbindVbrRequest) SetSmartAGUid(v int64) *UnbindVbrRequest {
	s.SmartAGUid = &v
	return s
}

func (s *UnbindVbrRequest) SetVbrId(v string) *UnbindVbrRequest {
	s.VbrId = &v
	return s
}

func (s *UnbindVbrRequest) SetVbrRegionId(v string) *UnbindVbrRequest {
	s.VbrRegionId = &v
	return s
}

func (s *UnbindVbrRequest) Validate() error {
	return dara.Validate(s)
}
