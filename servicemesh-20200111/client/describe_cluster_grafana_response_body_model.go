// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterGrafanaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDashboards(v []*DescribeClusterGrafanaResponseBodyDashboards) *DescribeClusterGrafanaResponseBody
	GetDashboards() []*DescribeClusterGrafanaResponseBodyDashboards
	SetRequestId(v string) *DescribeClusterGrafanaResponseBody
	GetRequestId() *string
}

type DescribeClusterGrafanaResponseBody struct {
	// The information about Grafana dashboards.
	Dashboards []*DescribeClusterGrafanaResponseBodyDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterGrafanaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterGrafanaResponseBody) GetDashboards() []*DescribeClusterGrafanaResponseBodyDashboards {
	return s.Dashboards
}

func (s *DescribeClusterGrafanaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterGrafanaResponseBody) SetDashboards(v []*DescribeClusterGrafanaResponseBodyDashboards) *DescribeClusterGrafanaResponseBody {
	s.Dashboards = v
	return s
}

func (s *DescribeClusterGrafanaResponseBody) SetRequestId(v string) *DescribeClusterGrafanaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterGrafanaResponseBody) Validate() error {
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

type DescribeClusterGrafanaResponseBodyDashboards struct {
	// The title of the Grafana dashboard.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The endpoint of the Grafana dashboard.
	//
	// example:
	//
	// test.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeClusterGrafanaResponseBodyDashboards) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterGrafanaResponseBodyDashboards) GoString() string {
	return s.String()
}

func (s *DescribeClusterGrafanaResponseBodyDashboards) GetTitle() *string {
	return s.Title
}

func (s *DescribeClusterGrafanaResponseBodyDashboards) GetUrl() *string {
	return s.Url
}

func (s *DescribeClusterGrafanaResponseBodyDashboards) SetTitle(v string) *DescribeClusterGrafanaResponseBodyDashboards {
	s.Title = &v
	return s
}

func (s *DescribeClusterGrafanaResponseBodyDashboards) SetUrl(v string) *DescribeClusterGrafanaResponseBodyDashboards {
	s.Url = &v
	return s
}

func (s *DescribeClusterGrafanaResponseBodyDashboards) Validate() error {
	return dara.Validate(s)
}
