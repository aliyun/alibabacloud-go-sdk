// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterAccessLogDashboardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDashboards(v []*DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) *DescribeGuestClusterAccessLogDashboardsResponseBody
	GetDashboards() []*DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards
	SetK8sClusterId(v string) *DescribeGuestClusterAccessLogDashboardsResponseBody
	GetK8sClusterId() *string
	SetRequestId(v string) *DescribeGuestClusterAccessLogDashboardsResponseBody
	GetRequestId() *string
}

type DescribeGuestClusterAccessLogDashboardsResponseBody struct {
	// The access log dashboards of the cluster on the data plane.
	Dashboards []*DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The ID of the cluster on the data plane.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGuestClusterAccessLogDashboardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterAccessLogDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) GetDashboards() []*DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards {
	return s.Dashboards
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) SetDashboards(v []*DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) *DescribeGuestClusterAccessLogDashboardsResponseBody {
	s.Dashboards = v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) SetK8sClusterId(v string) *DescribeGuestClusterAccessLogDashboardsResponseBody {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) SetRequestId(v string) *DescribeGuestClusterAccessLogDashboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBody) Validate() error {
	if s.Dashboards != nil {
		for _, item := range s.Dashboards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards struct {
	// The title of the dashboard.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of a dashboard.
	//
	// example:
	//
	// test.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) GetTitle() *string {
	return s.Title
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) GetUrl() *string {
	return s.Url
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) SetTitle(v string) *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards {
	s.Title = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) SetUrl(v string) *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards {
	s.Url = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponseBodyDashboards) Validate() error {
	return dara.Validate(s)
}
