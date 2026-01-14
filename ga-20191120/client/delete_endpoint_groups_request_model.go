// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEndpointGroupsRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteEndpointGroupsRequest
	GetDryRun() *bool
	SetEndpointGroupIds(v []*string) *DeleteEndpointGroupsRequest
	GetEndpointGroupIds() []*string
	SetRegionId(v string) *DeleteEndpointGroupsRequest
	GetRegionId() *string
}

type DeleteEndpointGroupsRequest struct {
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
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IDs of endpoint groups.
	//
	// You can specify up to 10 endpoint group IDs in each request.
	//
	// This parameter is required.
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEndpointGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEndpointGroupsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteEndpointGroupsRequest) GetEndpointGroupIds() []*string {
	return s.EndpointGroupIds
}

func (s *DeleteEndpointGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEndpointGroupsRequest) SetClientToken(v string) *DeleteEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEndpointGroupsRequest) SetDryRun(v bool) *DeleteEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteEndpointGroupsRequest) SetEndpointGroupIds(v []*string) *DeleteEndpointGroupsRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *DeleteEndpointGroupsRequest) SetRegionId(v string) *DeleteEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEndpointGroupsRequest) Validate() error {
	return dara.Validate(s)
}
