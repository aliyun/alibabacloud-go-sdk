// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcEndpointConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableVpcEndpointConnectionRequest
	GetClientToken() *string
	SetDryRun(v bool) *DisableVpcEndpointConnectionRequest
	GetDryRun() *bool
	SetEndpointId(v string) *DisableVpcEndpointConnectionRequest
	GetEndpointId() *string
	SetRegionId(v string) *DisableVpcEndpointConnectionRequest
	GetRegionId() *string
	SetServiceId(v string) *DisableVpcEndpointConnectionRequest
	GetServiceId() *string
}

type DisableVpcEndpointConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the connection request from the endpoint is rejected. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DisableVpcEndpointConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcEndpointConnectionRequest) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableVpcEndpointConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableVpcEndpointConnectionRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DisableVpcEndpointConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableVpcEndpointConnectionRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DisableVpcEndpointConnectionRequest) SetClientToken(v string) *DisableVpcEndpointConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) SetDryRun(v bool) *DisableVpcEndpointConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) SetEndpointId(v string) *DisableVpcEndpointConnectionRequest {
	s.EndpointId = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) SetRegionId(v string) *DisableVpcEndpointConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) SetServiceId(v string) *DisableVpcEndpointConnectionRequest {
	s.ServiceId = &v
	return s
}

func (s *DisableVpcEndpointConnectionRequest) Validate() error {
	return dara.Validate(s)
}
