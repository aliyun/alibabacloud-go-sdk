// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpareIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *GetSpareIpRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *GetSpareIpRequest
	GetClientToken() *string
	SetDryRun(v bool) *GetSpareIpRequest
	GetDryRun() *bool
	SetRegionId(v string) *GetSpareIpRequest
	GetRegionId() *string
	SetSpareIp(v string) *GetSpareIpRequest
	GetSpareIp() *string
}

type GetSpareIpRequest struct {
	// The instance ID of the Alibaba Cloud Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of a request.
	//
	// Generate a unique value from your client to ensure that different requests have unique ClientToken values. ClientToken supports only ASCII characters.
	//
	// example:
	//
	// 1F4B6A4A-C89E-489E-BAF1-52777EE148EF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without actually creating the resource. The system checks the required parameters, request syntax, and business limitations. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check passes, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the Alibaba Cloud Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The CNAME spare IP address. When an acceleration area is abnormal, traffic is switched to this IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.100.XX.XX
	SpareIp *string `json:"SpareIp,omitempty" xml:"SpareIp,omitempty"`
}

func (s GetSpareIpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSpareIpRequest) GoString() string {
	return s.String()
}

func (s *GetSpareIpRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetSpareIpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetSpareIpRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetSpareIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSpareIpRequest) GetSpareIp() *string {
	return s.SpareIp
}

func (s *GetSpareIpRequest) SetAcceleratorId(v string) *GetSpareIpRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetSpareIpRequest) SetClientToken(v string) *GetSpareIpRequest {
	s.ClientToken = &v
	return s
}

func (s *GetSpareIpRequest) SetDryRun(v bool) *GetSpareIpRequest {
	s.DryRun = &v
	return s
}

func (s *GetSpareIpRequest) SetRegionId(v string) *GetSpareIpRequest {
	s.RegionId = &v
	return s
}

func (s *GetSpareIpRequest) SetSpareIp(v string) *GetSpareIpRequest {
	s.SpareIp = &v
	return s
}

func (s *GetSpareIpRequest) Validate() error {
	return dara.Validate(s)
}
