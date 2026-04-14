// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusVirtualInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *ListPrometheusVirtualInstancesRequest
	GetNamespace() *string
	SetTenantId(v string) *ListPrometheusVirtualInstancesRequest
	GetTenantId() *string
}

type ListPrometheusVirtualInstancesRequest struct {
	// Optional cloud product
	//
	// example:
	//
	// ack-csi-fuse
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	TenantId  *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListPrometheusVirtualInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusVirtualInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusVirtualInstancesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListPrometheusVirtualInstancesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPrometheusVirtualInstancesRequest) SetNamespace(v string) *ListPrometheusVirtualInstancesRequest {
	s.Namespace = &v
	return s
}

func (s *ListPrometheusVirtualInstancesRequest) SetTenantId(v string) *ListPrometheusVirtualInstancesRequest {
	s.TenantId = &v
	return s
}

func (s *ListPrometheusVirtualInstancesRequest) Validate() error {
	return dara.Validate(s)
}
