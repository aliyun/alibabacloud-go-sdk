// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlinkApiProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlinkApiPath(v string) *FlinkApiProxyRequest
	GetFlinkApiPath() *string
	SetNamespace(v string) *FlinkApiProxyRequest
	GetNamespace() *string
	SetResourceId(v string) *FlinkApiProxyRequest
	GetResourceId() *string
	SetResourceType(v string) *FlinkApiProxyRequest
	GetResourceType() *string
}

type FlinkApiProxyRequest struct {
	// The path of the Flink UI.
	//
	// This parameter is required.
	//
	// example:
	//
	// /jobs/4df35f8e54554b23bf7dcd38a151****
	FlinkApiPath *string `json:"flinkApiPath,omitempty" xml:"flinkApiPath,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a27a3aa-c5b9-4dc1-8c86-be57d2d6****
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- jobs
	//
	// 	- sessionclusters
	//
	// This parameter is required.
	//
	// example:
	//
	// jobs
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s FlinkApiProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s FlinkApiProxyRequest) GoString() string {
	return s.String()
}

func (s *FlinkApiProxyRequest) GetFlinkApiPath() *string {
	return s.FlinkApiPath
}

func (s *FlinkApiProxyRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *FlinkApiProxyRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *FlinkApiProxyRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *FlinkApiProxyRequest) SetFlinkApiPath(v string) *FlinkApiProxyRequest {
	s.FlinkApiPath = &v
	return s
}

func (s *FlinkApiProxyRequest) SetNamespace(v string) *FlinkApiProxyRequest {
	s.Namespace = &v
	return s
}

func (s *FlinkApiProxyRequest) SetResourceId(v string) *FlinkApiProxyRequest {
	s.ResourceId = &v
	return s
}

func (s *FlinkApiProxyRequest) SetResourceType(v string) *FlinkApiProxyRequest {
	s.ResourceType = &v
	return s
}

func (s *FlinkApiProxyRequest) Validate() error {
	return dara.Validate(s)
}
