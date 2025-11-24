// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterAccessLogDashboardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetK8sClusterId(v string) *DescribeGuestClusterAccessLogDashboardsRequest
	GetK8sClusterId() *string
}

type DescribeGuestClusterAccessLogDashboardsRequest struct {
	// The ID of the cluster on the data plane.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
}

func (s DescribeGuestClusterAccessLogDashboardsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterAccessLogDashboardsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterAccessLogDashboardsRequest) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeGuestClusterAccessLogDashboardsRequest) SetK8sClusterId(v string) *DescribeGuestClusterAccessLogDashboardsRequest {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsRequest) Validate() error {
	return dara.Validate(s)
}
