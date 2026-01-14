// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSpareIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteSpareIpsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DeleteSpareIpsRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteSpareIpsRequest
	GetDryRun() *bool
	SetRegionId(v string) *DeleteSpareIpsRequest
	GetRegionId() *string
	SetSpareIps(v []*string) *DeleteSpareIpsRequest
	GetSpareIps() []*string
}

type DeleteSpareIpsRequest struct {
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
	// The secondary IP addresses to be deleted for the CNAME. If an acceleration area of the GA instance becomes unavailable, GA redirects the access traffic to the secondary IP addresses.
	//
	// Separate the IP addresses with commas (,). You can specify up to two secondary IP addresses.
	//
	// This parameter is required.
	SpareIps []*string `json:"SpareIps,omitempty" xml:"SpareIps,omitempty" type:"Repeated"`
}

func (s DeleteSpareIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSpareIpsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpareIpsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteSpareIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSpareIpsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteSpareIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSpareIpsRequest) GetSpareIps() []*string {
	return s.SpareIps
}

func (s *DeleteSpareIpsRequest) SetAcceleratorId(v string) *DeleteSpareIpsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteSpareIpsRequest) SetClientToken(v string) *DeleteSpareIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSpareIpsRequest) SetDryRun(v bool) *DeleteSpareIpsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteSpareIpsRequest) SetRegionId(v string) *DeleteSpareIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSpareIpsRequest) SetSpareIps(v []*string) *DeleteSpareIpsRequest {
	s.SpareIps = v
	return s
}

func (s *DeleteSpareIpsRequest) Validate() error {
	return dara.Validate(s)
}
