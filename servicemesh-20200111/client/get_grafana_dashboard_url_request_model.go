// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrafanaDashboardUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetK8sClusterId(v string) *GetGrafanaDashboardUrlRequest
	GetK8sClusterId() *string
	SetServiceMeshId(v string) *GetGrafanaDashboardUrlRequest
	GetServiceMeshId() *string
	SetTitle(v string) *GetGrafanaDashboardUrlRequest
	GetTitle() *string
}

type GetGrafanaDashboardUrlRequest struct {
	// The ID of the Container Service for Kubernetes (ACK) or ACK Serverless cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c94ca2d27f7aa47ab84ed73e6f084****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The name of the dashboard.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cloud ASM Istio Http Gateway
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetGrafanaDashboardUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGrafanaDashboardUrlRequest) GoString() string {
	return s.String()
}

func (s *GetGrafanaDashboardUrlRequest) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *GetGrafanaDashboardUrlRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetGrafanaDashboardUrlRequest) GetTitle() *string {
	return s.Title
}

func (s *GetGrafanaDashboardUrlRequest) SetK8sClusterId(v string) *GetGrafanaDashboardUrlRequest {
	s.K8sClusterId = &v
	return s
}

func (s *GetGrafanaDashboardUrlRequest) SetServiceMeshId(v string) *GetGrafanaDashboardUrlRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetGrafanaDashboardUrlRequest) SetTitle(v string) *GetGrafanaDashboardUrlRequest {
	s.Title = &v
	return s
}

func (s *GetGrafanaDashboardUrlRequest) Validate() error {
	return dara.Validate(s)
}
