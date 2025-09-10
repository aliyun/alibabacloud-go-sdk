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
}

type ListPrometheusVirtualInstancesRequest struct {
	// example:
	//
	// ack-csi-fuse
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
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

func (s *ListPrometheusVirtualInstancesRequest) SetNamespace(v string) *ListPrometheusVirtualInstancesRequest {
	s.Namespace = &v
	return s
}

func (s *ListPrometheusVirtualInstancesRequest) Validate() error {
	return dara.Validate(s)
}
