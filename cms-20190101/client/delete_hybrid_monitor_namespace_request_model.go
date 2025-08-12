// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DeleteHybridMonitorNamespaceRequest
	GetNamespace() *string
	SetRegionId(v string) *DeleteHybridMonitorNamespaceRequest
	GetRegionId() *string
}

type DeleteHybridMonitorNamespaceRequest struct {
	// The name of the namespace.
	//
	// For information about how to obtain the name of a namespace, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHybridMonitorNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteHybridMonitorNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHybridMonitorNamespaceRequest) SetNamespace(v string) *DeleteHybridMonitorNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteHybridMonitorNamespaceRequest) SetRegionId(v string) *DeleteHybridMonitorNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHybridMonitorNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
