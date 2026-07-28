// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVbrToVpconnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *AttachVbrToVpconnRequest
	GetDryRun() *bool
	SetRegionId(v string) *AttachVbrToVpconnRequest
	GetRegionId() *string
	SetToken(v string) *AttachVbrToVpconnRequest
	GetToken() *string
	SetVbrId(v string) *AttachVbrToVpconnRequest
	GetVbrId() *string
	SetVpconnId(v string) *AttachVbrToVpconnRequest
	GetVpconnId() *string
}

type AttachVbrToVpconnRequest struct {
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without associating the VBR instance with the shared Express Connect circuits. The system checks the required parameters, request format, and instance status. If the check fails, the corresponding error is returned. If the check succeeds, the request ID is returned.
	//
	// - **false*	- (default): sends a normal request. After the check succeeds, the VBR instance is directly associated with the shared Express Connect circuits.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the shared Express Connect circuits.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client token must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The instance ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp133sn3nwjvu7twc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The instance ID of the shared Express Connect circuits.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1mrgfbtmc9brre7****
	VpconnId *string `json:"VpconnId,omitempty" xml:"VpconnId,omitempty"`
}

func (s AttachVbrToVpconnRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachVbrToVpconnRequest) GoString() string {
	return s.String()
}

func (s *AttachVbrToVpconnRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AttachVbrToVpconnRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachVbrToVpconnRequest) GetToken() *string {
	return s.Token
}

func (s *AttachVbrToVpconnRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *AttachVbrToVpconnRequest) GetVpconnId() *string {
	return s.VpconnId
}

func (s *AttachVbrToVpconnRequest) SetDryRun(v bool) *AttachVbrToVpconnRequest {
	s.DryRun = &v
	return s
}

func (s *AttachVbrToVpconnRequest) SetRegionId(v string) *AttachVbrToVpconnRequest {
	s.RegionId = &v
	return s
}

func (s *AttachVbrToVpconnRequest) SetToken(v string) *AttachVbrToVpconnRequest {
	s.Token = &v
	return s
}

func (s *AttachVbrToVpconnRequest) SetVbrId(v string) *AttachVbrToVpconnRequest {
	s.VbrId = &v
	return s
}

func (s *AttachVbrToVpconnRequest) SetVpconnId(v string) *AttachVbrToVpconnRequest {
	s.VpconnId = &v
	return s
}

func (s *AttachVbrToVpconnRequest) Validate() error {
	return dara.Validate(s)
}
