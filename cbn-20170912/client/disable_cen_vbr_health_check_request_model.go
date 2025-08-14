// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCenVbrHealthCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DisableCenVbrHealthCheckRequest
	GetCenId() *string
	SetOwnerAccount(v string) *DisableCenVbrHealthCheckRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableCenVbrHealthCheckRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DisableCenVbrHealthCheckRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableCenVbrHealthCheckRequest
	GetResourceOwnerId() *int64
	SetVbrInstanceId(v string) *DisableCenVbrHealthCheckRequest
	GetVbrInstanceId() *string
	SetVbrInstanceOwnerId(v int64) *DisableCenVbrHealthCheckRequest
	GetVbrInstanceOwnerId() *int64
	SetVbrInstanceRegionId(v string) *DisableCenVbrHealthCheckRequest
	GetVbrInstanceRegionId() *string
}

type DisableCenVbrHealthCheckRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-sjfoejfghhjgghjghkg****
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-wz95o9aylj181n5****
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VBR belongs.
	//
	// > This parameter is required if the VBR and the CEN instance belong to different Alibaba Cloud accounts.
	//
	// example:
	//
	// 1250123456123456
	VbrInstanceOwnerId *int64 `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
	// The ID of the region where the VBR is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VbrInstanceRegionId *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
}

func (s DisableCenVbrHealthCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCenVbrHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *DisableCenVbrHealthCheckRequest) GetCenId() *string {
	return s.CenId
}

func (s *DisableCenVbrHealthCheckRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableCenVbrHealthCheckRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableCenVbrHealthCheckRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableCenVbrHealthCheckRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableCenVbrHealthCheckRequest) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *DisableCenVbrHealthCheckRequest) GetVbrInstanceOwnerId() *int64 {
	return s.VbrInstanceOwnerId
}

func (s *DisableCenVbrHealthCheckRequest) GetVbrInstanceRegionId() *string {
	return s.VbrInstanceRegionId
}

func (s *DisableCenVbrHealthCheckRequest) SetCenId(v string) *DisableCenVbrHealthCheckRequest {
	s.CenId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetOwnerAccount(v string) *DisableCenVbrHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetOwnerId(v int64) *DisableCenVbrHealthCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetResourceOwnerAccount(v string) *DisableCenVbrHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetResourceOwnerId(v int64) *DisableCenVbrHealthCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceId(v string) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceOwnerId(v int64) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceOwnerId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *DisableCenVbrHealthCheckRequest {
	s.VbrInstanceRegionId = &v
	return s
}

func (s *DisableCenVbrHealthCheckRequest) Validate() error {
	return dara.Validate(s)
}
