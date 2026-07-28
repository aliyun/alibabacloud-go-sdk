// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpconnFromVbrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *CreateVpconnFromVbrRequest
	GetDryRun() *bool
	SetOrderMode(v string) *CreateVpconnFromVbrRequest
	GetOrderMode() *string
	SetRegionId(v string) *CreateVpconnFromVbrRequest
	GetRegionId() *string
	SetToken(v string) *CreateVpconnFromVbrRequest
	GetToken() *string
	SetVbrId(v string) *CreateVpconnFromVbrRequest
	GetVbrId() *string
}

type CreateVpconnFromVbrRequest struct {
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: sends a check request without transforming the shared Express Connect circuits mode. The system checks the required parameters, request format, and instance status. If the check fails, the corresponding error is returned. If the check succeeds, the request ID is returned.
	//
	// - **false*	- (default): sends a Normal request and transforms the shared Express Connect circuits mode after the check succeeds.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The payer of the shared Express Connect circuits. Valid values:
	//
	// - **PayByPhysicalConnectionOwner**: The owner of the Express Connect circuit associated with the shared Express Connect circuits pays the fee.
	//
	// - **PayByVirtualPhysicalConnectionOwner**: The owner of the shared Express Connect circuits pays the fee.
	//
	// example:
	//
	// PayByVirtualPhysicalConnectionOwner
	OrderMode *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	// The region ID of the shared Express Connect circuits.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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
	// The instance ID of the cross-account VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp136flp1mf8mlq6r****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s CreateVpconnFromVbrRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpconnFromVbrRequest) GoString() string {
	return s.String()
}

func (s *CreateVpconnFromVbrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpconnFromVbrRequest) GetOrderMode() *string {
	return s.OrderMode
}

func (s *CreateVpconnFromVbrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpconnFromVbrRequest) GetToken() *string {
	return s.Token
}

func (s *CreateVpconnFromVbrRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *CreateVpconnFromVbrRequest) SetDryRun(v bool) *CreateVpconnFromVbrRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpconnFromVbrRequest) SetOrderMode(v string) *CreateVpconnFromVbrRequest {
	s.OrderMode = &v
	return s
}

func (s *CreateVpconnFromVbrRequest) SetRegionId(v string) *CreateVpconnFromVbrRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpconnFromVbrRequest) SetToken(v string) *CreateVpconnFromVbrRequest {
	s.Token = &v
	return s
}

func (s *CreateVpconnFromVbrRequest) SetVbrId(v string) *CreateVpconnFromVbrRequest {
	s.VbrId = &v
	return s
}

func (s *CreateVpconnFromVbrRequest) Validate() error {
	return dara.Validate(s)
}
