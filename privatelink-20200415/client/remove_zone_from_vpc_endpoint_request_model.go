// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveZoneFromVpcEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveZoneFromVpcEndpointRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveZoneFromVpcEndpointRequest
	GetDryRun() *bool
	SetEndpointId(v string) *RemoveZoneFromVpcEndpointRequest
	GetEndpointId() *string
	SetRegionId(v string) *RemoveZoneFromVpcEndpointRequest
	GetRegionId() *string
	SetZoneId(v string) *RemoveZoneFromVpcEndpointRequest
	GetZoneId() *string
}

type RemoveZoneFromVpcEndpointRequest struct {
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
	// The ID of the endpoint for which you want to delete the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint for which you want to delete the zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the zone that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RemoveZoneFromVpcEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveZoneFromVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *RemoveZoneFromVpcEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveZoneFromVpcEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveZoneFromVpcEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *RemoveZoneFromVpcEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveZoneFromVpcEndpointRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *RemoveZoneFromVpcEndpointRequest) SetClientToken(v string) *RemoveZoneFromVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) SetDryRun(v bool) *RemoveZoneFromVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) SetEndpointId(v string) *RemoveZoneFromVpcEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) SetRegionId(v string) *RemoveZoneFromVpcEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) SetZoneId(v string) *RemoveZoneFromVpcEndpointRequest {
	s.ZoneId = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointRequest) Validate() error {
	return dara.Validate(s)
}
