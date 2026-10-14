// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterKubeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterKubeConfigRequest
	GetClusterId() *string
}

type DescribeClusterKubeConfigRequest struct {
	// The cluster ID.
	//
	// This parameter is essentially required. If it is not provided, the service returns BadRequest. The target cluster must also have completed control plane deployment and certificate generation. You cannot obtain the kubeconfig for a cluster in the initial state.
	//
	// example:
	//
	// c8f0377146d104687ac562eef9403****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterKubeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterKubeConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterKubeConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterKubeConfigRequest) SetClusterId(v string) *DescribeClusterKubeConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterKubeConfigRequest) Validate() error {
	return dara.Validate(s)
}
