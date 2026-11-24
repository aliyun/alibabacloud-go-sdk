// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *DeleteRouteMapAssociationRequest
	GetAssociationId() *string
	SetClientToken(v string) *DeleteRouteMapAssociationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteRouteMapAssociationRequest
	GetDryRun() *bool
	SetEcrId(v string) *DeleteRouteMapAssociationRequest
	GetEcrId() *string
	SetVersion(v string) *DeleteRouteMapAssociationRequest
	GetVersion() *string
}

type DeleteRouteMapAssociationRequest struct {
	// The ID of the route policy association.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-routemap-assoc-qdoeneegffq4yv****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can generate the value from your client and ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether it is only a dry run for this request. Valid values:
	//
	// - **true**: sends a check request without modifying the EPS resource.
	//
	// - **false*	- (default): sends a normal request.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DeleteRouteMapAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapAssociationRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapAssociationRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DeleteRouteMapAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRouteMapAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRouteMapAssociationRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DeleteRouteMapAssociationRequest) GetVersion() *string {
	return s.Version
}

func (s *DeleteRouteMapAssociationRequest) SetAssociationId(v string) *DeleteRouteMapAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *DeleteRouteMapAssociationRequest) SetClientToken(v string) *DeleteRouteMapAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRouteMapAssociationRequest) SetDryRun(v bool) *DeleteRouteMapAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRouteMapAssociationRequest) SetEcrId(v string) *DeleteRouteMapAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *DeleteRouteMapAssociationRequest) SetVersion(v string) *DeleteRouteMapAssociationRequest {
	s.Version = &v
	return s
}

func (s *DeleteRouteMapAssociationRequest) Validate() error {
	return dara.Validate(s)
}
