// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcEndpointZoneConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableVpcEndpointZoneConnectionRequest
	GetClientToken() *string
	SetDryRun(v bool) *DisableVpcEndpointZoneConnectionRequest
	GetDryRun() *bool
	SetEndpointId(v string) *DisableVpcEndpointZoneConnectionRequest
	GetEndpointId() *string
	SetRegionId(v string) *DisableVpcEndpointZoneConnectionRequest
	GetRegionId() *string
	SetReplacedResource(v bool) *DisableVpcEndpointZoneConnectionRequest
	GetReplacedResource() *bool
	SetServiceId(v string) *DisableVpcEndpointZoneConnectionRequest
	GetServiceId() *string
	SetZoneId(v string) *DisableVpcEndpointZoneConnectionRequest
	GetZoneId() *string
}

type DisableVpcEndpointZoneConnectionRequest struct {
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
	// The ID of the region where the connection request from the endpoint is rejected.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to close connections in the endpoint zone after migration. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// >  Set the value to true if you want to close connections in the endpoint zone after migration.
	//
	// example:
	//
	// false
	ReplacedResource *bool `json:"ReplacedResource,omitempty" xml:"ReplacedResource,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the zone that is associated with the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DisableVpcEndpointZoneConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcEndpointZoneConnectionRequest) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointZoneConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableVpcEndpointZoneConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableVpcEndpointZoneConnectionRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DisableVpcEndpointZoneConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableVpcEndpointZoneConnectionRequest) GetReplacedResource() *bool {
	return s.ReplacedResource
}

func (s *DisableVpcEndpointZoneConnectionRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DisableVpcEndpointZoneConnectionRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetClientToken(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetDryRun(v bool) *DisableVpcEndpointZoneConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetEndpointId(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.EndpointId = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetRegionId(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetReplacedResource(v bool) *DisableVpcEndpointZoneConnectionRequest {
	s.ReplacedResource = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetServiceId(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.ServiceId = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) SetZoneId(v string) *DisableVpcEndpointZoneConnectionRequest {
	s.ZoneId = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionRequest) Validate() error {
	return dara.Validate(s)
}
