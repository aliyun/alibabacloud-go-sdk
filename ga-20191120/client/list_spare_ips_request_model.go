// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpareIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListSpareIpsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *ListSpareIpsRequest
	GetClientToken() *string
	SetDryRun(v bool) *ListSpareIpsRequest
	GetDryRun() *bool
	SetRegionId(v string) *ListSpareIpsRequest
	GetRegionId() *string
}

type ListSpareIpsRequest struct {
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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
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
}

func (s ListSpareIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSpareIpsRequest) GoString() string {
	return s.String()
}

func (s *ListSpareIpsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListSpareIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListSpareIpsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListSpareIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSpareIpsRequest) SetAcceleratorId(v string) *ListSpareIpsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListSpareIpsRequest) SetClientToken(v string) *ListSpareIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSpareIpsRequest) SetDryRun(v bool) *ListSpareIpsRequest {
	s.DryRun = &v
	return s
}

func (s *ListSpareIpsRequest) SetRegionId(v string) *ListSpareIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSpareIpsRequest) Validate() error {
	return dara.Validate(s)
}
