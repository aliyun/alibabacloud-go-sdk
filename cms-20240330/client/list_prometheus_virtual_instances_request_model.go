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
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Optional cloud product
	//
	// example:
	//
	// ack-csi-fuse
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TenantId  *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
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
