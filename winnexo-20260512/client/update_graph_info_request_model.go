// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGraphInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessProfile(v string) *UpdateGraphInfoRequest
	GetBusinessProfile() *string
	SetDisplayName(v string) *UpdateGraphInfoRequest
	GetDisplayName() *string
	SetGraphName(v string) *UpdateGraphInfoRequest
	GetGraphName() *string
	SetTenantId(v string) *UpdateGraphInfoRequest
	GetTenantId() *string
}

type UpdateGraphInfoRequest struct {
	// The business description of the knowledge graph. If not configured, the value is an empty string.
	//
	// example:
	//
	// Customer domain knowledge graph
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// The display name of the knowledge graph.
	//
	// example:
	//
	// CRM Graph
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the knowledge graph.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateGraphInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGraphInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateGraphInfoRequest) GetBusinessProfile() *string {
	return s.BusinessProfile
}

func (s *UpdateGraphInfoRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateGraphInfoRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *UpdateGraphInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGraphInfoRequest) SetBusinessProfile(v string) *UpdateGraphInfoRequest {
	s.BusinessProfile = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetDisplayName(v string) *UpdateGraphInfoRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetGraphName(v string) *UpdateGraphInfoRequest {
	s.GraphName = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetTenantId(v string) *UpdateGraphInfoRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGraphInfoRequest) Validate() error {
	return dara.Validate(s)
}
