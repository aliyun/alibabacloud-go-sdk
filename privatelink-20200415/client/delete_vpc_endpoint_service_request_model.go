// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcEndpointServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpcEndpointServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteVpcEndpointServiceRequest
	GetDryRun() *bool
	SetRegionId(v string) *DeleteVpcEndpointServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *DeleteVpcEndpointServiceRequest
	GetServiceId() *string
}

type DeleteVpcEndpointServiceRequest struct {
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
	// The region ID of the endpoint service. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the endpoint service that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteVpcEndpointServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpcEndpointServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVpcEndpointServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpcEndpointServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteVpcEndpointServiceRequest) SetClientToken(v string) *DeleteVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcEndpointServiceRequest) SetDryRun(v bool) *DeleteVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVpcEndpointServiceRequest) SetRegionId(v string) *DeleteVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpcEndpointServiceRequest) SetServiceId(v string) *DeleteVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteVpcEndpointServiceRequest) Validate() error {
	return dara.Validate(s)
}
