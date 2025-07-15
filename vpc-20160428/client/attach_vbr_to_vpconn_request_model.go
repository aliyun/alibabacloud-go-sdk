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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including required parameters, request syntax, and instance status. If the request fails the dry run, an error message is returned. If the request passes the dry run, the request ID is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the hosted connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp133sn3nwjvu7twc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The ID of the hosted connection.
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
