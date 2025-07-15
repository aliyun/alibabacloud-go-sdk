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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and instance status. If the request fails the dry run, an error message is returned. If the request passes the dry run, the system returns the ID of the request.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The payer for the shared Express Connect circuit. Valid values:
	//
	// 	- **PayByPhysicalConnectionOwner**: the owner of the shared Express Connect circuit
	//
	// 	- **PayByVirtualPhysicalConnectionOwner**: the owner of the hosted connection
	//
	// example:
	//
	// PayByVirtualPhysicalConnectionOwner
	OrderMode *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
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
	// The ID of the associated VBR.
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
