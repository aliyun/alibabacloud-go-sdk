// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusVirtualInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPrometheusVirtualInstancesRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListPrometheusVirtualInstancesRequest
	GetNamespace() *string
	SetNextToken(v string) *ListPrometheusVirtualInstancesRequest
	GetNextToken() *string
	SetTenantId(v string) *ListPrometheusVirtualInstancesRequest
	GetTenantId() *string
}

type ListPrometheusVirtualInstancesRequest struct {
	// The maximum number of instances to return per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The optional cloud services.
	//
	// example:
	//
	// ack-csi-fuse
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The next token for the paging query.
	//
	// example:
	//
	// a1b2c3d4e5f60718293a4b5c6d7e****8f9a0b1c2d3e4f508f9a0b1c2d3e4f508f9a0b1c2d3e4f50
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890123456
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListPrometheusVirtualInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusVirtualInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusVirtualInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrometheusVirtualInstancesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListPrometheusVirtualInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrometheusVirtualInstancesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPrometheusVirtualInstancesRequest) SetMaxResults(v int32) *ListPrometheusVirtualInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPrometheusVirtualInstancesRequest) SetNamespace(v string) *ListPrometheusVirtualInstancesRequest {
	s.Namespace = &v
	return s
}

func (s *ListPrometheusVirtualInstancesRequest) SetNextToken(v string) *ListPrometheusVirtualInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPrometheusVirtualInstancesRequest) SetTenantId(v string) *ListPrometheusVirtualInstancesRequest {
	s.TenantId = &v
	return s
}

func (s *ListPrometheusVirtualInstancesRequest) Validate() error {
	return dara.Validate(s)
}
