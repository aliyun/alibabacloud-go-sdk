// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVbrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *BindVbrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BindVbrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *BindVbrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *BindVbrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindVbrRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *BindVbrRequest
	GetSmartAGId() *string
	SetSmartAGUid(v int64) *BindVbrRequest
	GetSmartAGUid() *int64
	SetVbrId(v string) *BindVbrRequest
	GetVbrId() *string
	SetVbrRegionId(v string) *BindVbrRequest
	GetVbrRegionId() *string
}

type BindVbrRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-eoqbp1fmrsgbrn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 1250123456123456
	SmartAGUid *int64 `json:"SmartAGUid,omitempty" xml:"SmartAGUid,omitempty"`
	// The ID of the VBR that you want to associate with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-gc7ccdvtp3l4fec0j****
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

func (s BindVbrRequest) String() string {
	return dara.Prettify(s)
}

func (s BindVbrRequest) GoString() string {
	return s.String()
}

func (s *BindVbrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BindVbrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindVbrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindVbrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindVbrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindVbrRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *BindVbrRequest) GetSmartAGUid() *int64 {
	return s.SmartAGUid
}

func (s *BindVbrRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *BindVbrRequest) GetVbrRegionId() *string {
	return s.VbrRegionId
}

func (s *BindVbrRequest) SetOwnerAccount(v string) *BindVbrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindVbrRequest) SetOwnerId(v int64) *BindVbrRequest {
	s.OwnerId = &v
	return s
}

func (s *BindVbrRequest) SetRegionId(v string) *BindVbrRequest {
	s.RegionId = &v
	return s
}

func (s *BindVbrRequest) SetResourceOwnerAccount(v string) *BindVbrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindVbrRequest) SetResourceOwnerId(v int64) *BindVbrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindVbrRequest) SetSmartAGId(v string) *BindVbrRequest {
	s.SmartAGId = &v
	return s
}

func (s *BindVbrRequest) SetSmartAGUid(v int64) *BindVbrRequest {
	s.SmartAGUid = &v
	return s
}

func (s *BindVbrRequest) SetVbrId(v string) *BindVbrRequest {
	s.VbrId = &v
	return s
}

func (s *BindVbrRequest) SetVbrRegionId(v string) *BindVbrRequest {
	s.VbrRegionId = &v
	return s
}

func (s *BindVbrRequest) Validate() error {
	return dara.Validate(s)
}
