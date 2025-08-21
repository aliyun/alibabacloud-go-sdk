// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterKubeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterKubeConfigResponseBody
	GetClusterId() *string
	SetKubeconfig(v string) *DescribeClusterKubeConfigResponseBody
	GetKubeconfig() *string
	SetRequestId(v string) *DescribeClusterKubeConfigResponseBody
	GetRequestId() *string
}

type DescribeClusterKubeConfigResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c8f0377146d104687ac562eef9403****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster certificate.
	//
	// example:
	//
	// apiVersion: v1
	//
	// clusters:
	//
	// - cluster:
	//
	//     certificate-authority-data:***
	//
	//     server: https://****:6443
	//
	//   name: kubernetes
	//
	// contexts:
	//
	// - context:
	//
	//     cluster: kubernetes
	//
	//     user: "2580306074811*****"
	//
	//   name: 258*******
	//
	// kind: Config
	//
	// users:
	//
	// - name: "2580306074811*****"
	//
	//   user:
	//
	//     client-certificate-data:***
	//
	//     client-key-data: ***
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterKubeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterKubeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterKubeConfigResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterKubeConfigResponseBody) GetKubeconfig() *string {
	return s.Kubeconfig
}

func (s *DescribeClusterKubeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterKubeConfigResponseBody) SetClusterId(v string) *DescribeClusterKubeConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterKubeConfigResponseBody) SetKubeconfig(v string) *DescribeClusterKubeConfigResponseBody {
	s.Kubeconfig = &v
	return s
}

func (s *DescribeClusterKubeConfigResponseBody) SetRequestId(v string) *DescribeClusterKubeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterKubeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
