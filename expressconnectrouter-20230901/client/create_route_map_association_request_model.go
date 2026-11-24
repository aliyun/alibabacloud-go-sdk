// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRouteMapAssociationRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateRouteMapAssociationRequest
	GetDryRun() *bool
	SetEcrId(v string) *CreateRouteMapAssociationRequest
	GetEcrId() *string
	SetRegionIdList(v []*string) *CreateRouteMapAssociationRequest
	GetRegionIdList() []*string
	SetRouteMapId(v string) *CreateRouteMapAssociationRequest
	GetRouteMapId() *string
	SetVersion(v string) *CreateRouteMapAssociationRequest
	GetVersion() *string
}

type CreateRouteMapAssociationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client and ensure that the value is unique across different requests. The client token supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether it is a dry run only. Valid values:
	//
	// - **true**: sends a dry run request without modifying EPS resources.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The Express Connect Router (ECR) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The list of region IDs of the Express Connect Router (ECR) where the routing policy is distributed and takes effect.
	//
	// > If this parameter is left empty, the routing policy is not distributed to any region.
	RegionIdList []*string `json:"RegionIdList,omitempty" xml:"RegionIdList,omitempty" type:"Repeated"`
	// The ID of the route map in the Cloud Enterprise Network (CEN).
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-6m58b17417tg5a****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateRouteMapAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapAssociationRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteMapAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouteMapAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRouteMapAssociationRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *CreateRouteMapAssociationRequest) GetRegionIdList() []*string {
	return s.RegionIdList
}

func (s *CreateRouteMapAssociationRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *CreateRouteMapAssociationRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateRouteMapAssociationRequest) SetClientToken(v string) *CreateRouteMapAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteMapAssociationRequest) SetDryRun(v bool) *CreateRouteMapAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRouteMapAssociationRequest) SetEcrId(v string) *CreateRouteMapAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *CreateRouteMapAssociationRequest) SetRegionIdList(v []*string) *CreateRouteMapAssociationRequest {
	s.RegionIdList = v
	return s
}

func (s *CreateRouteMapAssociationRequest) SetRouteMapId(v string) *CreateRouteMapAssociationRequest {
	s.RouteMapId = &v
	return s
}

func (s *CreateRouteMapAssociationRequest) SetVersion(v string) *CreateRouteMapAssociationRequest {
	s.Version = &v
	return s
}

func (s *CreateRouteMapAssociationRequest) Validate() error {
	return dara.Validate(s)
}
