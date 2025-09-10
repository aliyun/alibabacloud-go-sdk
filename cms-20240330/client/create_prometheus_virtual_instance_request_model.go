// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusVirtualInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *CreatePrometheusVirtualInstanceRequest
	GetNamespace() *string
}

type CreatePrometheusVirtualInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cms_prometheus
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s CreatePrometheusVirtualInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusVirtualInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusVirtualInstanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreatePrometheusVirtualInstanceRequest) SetNamespace(v string) *CreatePrometheusVirtualInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *CreatePrometheusVirtualInstanceRequest) Validate() error {
	return dara.Validate(s)
}
