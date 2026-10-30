// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteRouteMapRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteRouteMapRequest
	GetDryRun() *bool
	SetRouteMapId(v string) *DeleteRouteMapRequest
	GetRouteMapId() *string
}

type DeleteRouteMapRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client and make sure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: The system prechecks the request without deleting the route map. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the precheck, the system returns an error. If the request passes the precheck, the system returns the DryRunOperation error code.
	//
	// - **false*	- (default): The system sends a normal request and deletes the route map after the request passes the precheck.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the route map.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-w4yf7toozfol3q****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DeleteRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRouteMapRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRouteMapRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *DeleteRouteMapRequest) SetClientToken(v string) *DeleteRouteMapRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRouteMapRequest) SetDryRun(v bool) *DeleteRouteMapRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRouteMapRequest) SetRouteMapId(v string) *DeleteRouteMapRequest {
	s.RouteMapId = &v
	return s
}

func (s *DeleteRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
