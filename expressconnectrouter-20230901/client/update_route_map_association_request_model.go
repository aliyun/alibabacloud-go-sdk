// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *UpdateRouteMapAssociationRequest
	GetAssociationId() *string
	SetClientToken(v string) *UpdateRouteMapAssociationRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateRouteMapAssociationRequest
	GetDryRun() *bool
	SetEcrId(v string) *UpdateRouteMapAssociationRequest
	GetEcrId() *string
	SetRegionIdList(v []*string) *UpdateRouteMapAssociationRequest
	GetRegionIdList() []*string
	SetVersion(v string) *UpdateRouteMapAssociationRequest
	GetVersion() *string
}

type UpdateRouteMapAssociationRequest struct {
	// The ID of the routing policy association.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-routemap-assoc-qdoeneegffq4yv****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value on your client to ensure that it is unique across different requests. The client token only supports ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - **true**: Performs a dry run and does not modify service resources.
	//
	// - **false*	- (default): Sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The list of region IDs where the routing policy is delivered and takes effect for the Express Connect Router (ECR).
	//
	// > If this parameter is left empty, the routing policy is not delivered to any region.
	RegionIdList []*string `json:"RegionIdList,omitempty" xml:"RegionIdList,omitempty" type:"Repeated"`
	Version      *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateRouteMapAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapAssociationRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapAssociationRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *UpdateRouteMapAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRouteMapAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateRouteMapAssociationRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *UpdateRouteMapAssociationRequest) GetRegionIdList() []*string {
	return s.RegionIdList
}

func (s *UpdateRouteMapAssociationRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateRouteMapAssociationRequest) SetAssociationId(v string) *UpdateRouteMapAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *UpdateRouteMapAssociationRequest) SetClientToken(v string) *UpdateRouteMapAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRouteMapAssociationRequest) SetDryRun(v bool) *UpdateRouteMapAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRouteMapAssociationRequest) SetEcrId(v string) *UpdateRouteMapAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *UpdateRouteMapAssociationRequest) SetRegionIdList(v []*string) *UpdateRouteMapAssociationRequest {
	s.RegionIdList = v
	return s
}

func (s *UpdateRouteMapAssociationRequest) SetVersion(v string) *UpdateRouteMapAssociationRequest {
	s.Version = &v
	return s
}

func (s *UpdateRouteMapAssociationRequest) Validate() error {
	return dara.Validate(s)
}
