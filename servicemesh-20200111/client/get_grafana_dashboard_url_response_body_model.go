// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrafanaDashboardUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDashboards(v []*GetGrafanaDashboardUrlResponseBodyDashboards) *GetGrafanaDashboardUrlResponseBody
	GetDashboards() []*GetGrafanaDashboardUrlResponseBodyDashboards
	SetRequestId(v string) *GetGrafanaDashboardUrlResponseBody
	GetRequestId() *string
}

type GetGrafanaDashboardUrlResponseBody struct {
	// The information about the dashboard.
	Dashboards []*GetGrafanaDashboardUrlResponseBodyDashboards `json:"Dashboards,omitempty" xml:"Dashboards,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 76DBB8A0-5AA6-5A56-9A8A-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGrafanaDashboardUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGrafanaDashboardUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetGrafanaDashboardUrlResponseBody) GetDashboards() []*GetGrafanaDashboardUrlResponseBodyDashboards {
	return s.Dashboards
}

func (s *GetGrafanaDashboardUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGrafanaDashboardUrlResponseBody) SetDashboards(v []*GetGrafanaDashboardUrlResponseBodyDashboards) *GetGrafanaDashboardUrlResponseBody {
	s.Dashboards = v
	return s
}

func (s *GetGrafanaDashboardUrlResponseBody) SetRequestId(v string) *GetGrafanaDashboardUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGrafanaDashboardUrlResponseBody) Validate() error {
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

type GetGrafanaDashboardUrlResponseBodyDashboards struct {
	// The name of the dashboard.
	//
	// example:
	//
	// Cloud ASM Istio Http Gateway
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the dashboard.
	//
	// example:
	//
	// https://g.console.aliyun.com/d/181863583797****-14651340-200-2/alibaba-cloud-mesh-service?orgId=32****&refresh=60s
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetGrafanaDashboardUrlResponseBodyDashboards) String() string {
	return dara.Prettify(s)
}

func (s GetGrafanaDashboardUrlResponseBodyDashboards) GoString() string {
	return s.String()
}

func (s *GetGrafanaDashboardUrlResponseBodyDashboards) GetTitle() *string {
	return s.Title
}

func (s *GetGrafanaDashboardUrlResponseBodyDashboards) GetUrl() *string {
	return s.Url
}

func (s *GetGrafanaDashboardUrlResponseBodyDashboards) SetTitle(v string) *GetGrafanaDashboardUrlResponseBodyDashboards {
	s.Title = &v
	return s
}

func (s *GetGrafanaDashboardUrlResponseBodyDashboards) SetUrl(v string) *GetGrafanaDashboardUrlResponseBodyDashboards {
	s.Url = &v
	return s
}

func (s *GetGrafanaDashboardUrlResponseBodyDashboards) Validate() error {
	return dara.Validate(s)
}
