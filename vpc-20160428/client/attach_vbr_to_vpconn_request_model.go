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
	// - **true**: performs a dry run without associating the VBR instance with shared Express Connect circuits. The system checks whether the required parameters are specified, the request format is valid, and the instance status is correct. If the check fails, the corresponding error is returned. If the check passes, the request ID is returned.
	//
	// - **false*	- (default): sends a normal request. After the check passes, the VBR instance is directly associated with shared Express Connect circuits.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the shared Express Connect circuits.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client token must be unique among different requests. The maximum length is 64 ASCII characters.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The VBR instance ID.
	//
	// >The ID of the VBR instance to be migrated. The VBR must currently be directly attached to an Express Connect circuit owned by the caller, and must be the same VBR specified in CreateVpconnFromVbr.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp133sn3nwjvu7twc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The ID of the shared Express Connect circuits (VirtualPhysicalConnection) instance.
	//
	// >The shared Express Connect circuits instance ID returned by CreateVpconnFromVbr. The instance must have been confirmed and accepted by the tenant (Confirmed) and be in the Enabled state.
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
