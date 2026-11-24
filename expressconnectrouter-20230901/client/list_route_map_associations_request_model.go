// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *ListRouteMapAssociationsRequest
	GetAssociationId() *string
	SetClientToken(v string) *ListRouteMapAssociationsRequest
	GetClientToken() *string
	SetDryRun(v bool) *ListRouteMapAssociationsRequest
	GetDryRun() *bool
	SetEcrId(v string) *ListRouteMapAssociationsRequest
	GetEcrId() *string
	SetMaxResults(v int32) *ListRouteMapAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRouteMapAssociationsRequest
	GetNextToken() *string
	SetRouteMapId(v string) *ListRouteMapAssociationsRequest
	GetRouteMapId() *string
	SetVersion(v string) *ListRouteMapAssociationsRequest
	GetVersion() *string
}

type ListRouteMapAssociationsRequest struct {
	// The ID of the route map association.
	//
	// example:
	//
	// ecr-routemap-assoc-qdoeneegffq4yv****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can generate the value from your client and make sure that the value is unique across different requests. The client token supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether this is a dry run on the request. Valid values:
	//
	// - **true**: performs a dry run. The system checks the request but does not modify EPS resources.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query. Valid values:
	//
	// - If this is the first query or no next query is to be sent, this parameter is not required.
	//
	// - If a next query is to be sent, set the value to the **NextToken*	- value returned by the last API call.
	//
	// example:
	//
	// e0a2dbeb69a8beeeb8194e92b702df3fd3e7bfe6ce7bfc16e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the route map in the Cloud Enterprise Network (CEN).
	//
	// example:
	//
	// rm-6m58b17417tg5a****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListRouteMapAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListRouteMapAssociationsRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *ListRouteMapAssociationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListRouteMapAssociationsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListRouteMapAssociationsRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ListRouteMapAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRouteMapAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRouteMapAssociationsRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *ListRouteMapAssociationsRequest) GetVersion() *string {
	return s.Version
}

func (s *ListRouteMapAssociationsRequest) SetAssociationId(v string) *ListRouteMapAssociationsRequest {
	s.AssociationId = &v
	return s
}

func (s *ListRouteMapAssociationsRequest) SetClientToken(v string) *ListRouteMapAssociationsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListRouteMapAssociationsRequest) SetDryRun(v bool) *ListRouteMapAssociationsRequest {
	s.DryRun = &v
	return s
}

func (s *ListRouteMapAssociationsRequest) SetEcrId(v string) *ListRouteMapAssociationsRequest {
	s.EcrId = &v
	return s
}

func (s *ListRouteMapAssociationsRequest) SetMaxResults(v int32) *ListRouteMapAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRouteMapAssociationsRequest) SetNextToken(v string) *ListRouteMapAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRouteMapAssociationsRequest) SetRouteMapId(v string) *ListRouteMapAssociationsRequest {
	s.RouteMapId = &v
	return s
}

func (s *ListRouteMapAssociationsRequest) SetVersion(v string) *ListRouteMapAssociationsRequest {
	s.Version = &v
	return s
}

func (s *ListRouteMapAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
