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
	// The GA instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 1F4B6A4A-C89E-489E-BAF1-52777EE148EF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The secondary IP address that is associated with the CNAME. If the acceleration area becomes unavailable, GA redirects traffic to the secondary IP address.
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
