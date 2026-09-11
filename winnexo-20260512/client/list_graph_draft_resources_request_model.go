// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphDraftResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraphName(v string) *ListGraphDraftResourcesRequest
	GetGraphName() *string
	SetTenantId(v string) *ListGraphDraftResourcesRequest
	GetTenantId() *string
}

type ListGraphDraftResourcesRequest struct {
	// The graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGraphDraftResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGraphDraftResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListGraphDraftResourcesRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *ListGraphDraftResourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGraphDraftResourcesRequest) SetGraphName(v string) *ListGraphDraftResourcesRequest {
	s.GraphName = &v
	return s
}

func (s *ListGraphDraftResourcesRequest) SetTenantId(v string) *ListGraphDraftResourcesRequest {
	s.TenantId = &v
	return s
}

func (s *ListGraphDraftResourcesRequest) Validate() error {
	return dara.Validate(s)
}
