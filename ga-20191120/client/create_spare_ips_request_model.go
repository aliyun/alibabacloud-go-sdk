// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSpareIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateSpareIpsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateSpareIpsRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateSpareIpsRequest
	GetDryRun() *bool
	SetRegionId(v string) *CreateSpareIpsRequest
	GetRegionId() *string
	SetSpareIps(v []*string) *CreateSpareIpsRequest
	GetSpareIps() []*string
}

type CreateSpareIpsRequest struct {
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
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 1F4B6A4A-C89E-489E-BAF1-52777EE148EF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true:*	- performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (defalut): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
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
	// The secondary IP addresses to be created for the CNAME. If an acceleration area of the GA instance becomes unavailable, access traffic is redirected to the secondary IP addresses.
	//
	// You can specify up to two secondary IP addresses.
	//
	// This parameter is required.
	SpareIps []*string `json:"SpareIps,omitempty" xml:"SpareIps,omitempty" type:"Repeated"`
}

func (s CreateSpareIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSpareIpsRequest) GoString() string {
	return s.String()
}

func (s *CreateSpareIpsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateSpareIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSpareIpsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateSpareIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSpareIpsRequest) GetSpareIps() []*string {
	return s.SpareIps
}

func (s *CreateSpareIpsRequest) SetAcceleratorId(v string) *CreateSpareIpsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateSpareIpsRequest) SetClientToken(v string) *CreateSpareIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSpareIpsRequest) SetDryRun(v bool) *CreateSpareIpsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSpareIpsRequest) SetRegionId(v string) *CreateSpareIpsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSpareIpsRequest) SetSpareIps(v []*string) *CreateSpareIpsRequest {
	s.SpareIps = v
	return s
}

func (s *CreateSpareIpsRequest) Validate() error {
	return dara.Validate(s)
}
