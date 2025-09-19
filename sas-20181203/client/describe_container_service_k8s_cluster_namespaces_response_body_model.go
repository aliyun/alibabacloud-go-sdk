// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClusterNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetK8sClusterNamespaces(v []*DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces) *DescribeContainerServiceK8sClusterNamespacesResponseBody
	GetK8sClusterNamespaces() []*DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces
	SetRequestId(v string) *DescribeContainerServiceK8sClusterNamespacesResponseBody
	GetRequestId() *string
}

type DescribeContainerServiceK8sClusterNamespacesResponseBody struct {
	// The namespaces.
	K8sClusterNamespaces []*DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces `json:"K8sClusterNamespaces,omitempty" xml:"K8sClusterNamespaces,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0C8487EF-50C2-54BB-8634-10F8C35D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerServiceK8sClusterNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClusterNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponseBody) GetK8sClusterNamespaces() []*DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces {
	return s.K8sClusterNamespaces
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponseBody) SetK8sClusterNamespaces(v []*DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces) *DescribeContainerServiceK8sClusterNamespacesResponseBody {
	s.K8sClusterNamespaces = v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponseBody) SetRequestId(v string) *DescribeContainerServiceK8sClusterNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces struct {
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces) SetNamespace(v string) *DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponseBodyK8sClusterNamespaces) Validate() error {
	return dara.Validate(s)
}
