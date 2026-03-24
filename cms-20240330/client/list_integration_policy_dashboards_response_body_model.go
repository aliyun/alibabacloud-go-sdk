// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyDashboardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDashboards(v []*ListIntegrationPolicyDashboardsResponseBodyDashboards) *ListIntegrationPolicyDashboardsResponseBody
	GetDashboards() []*ListIntegrationPolicyDashboardsResponseBodyDashboards
	SetRequestId(v string) *ListIntegrationPolicyDashboardsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListIntegrationPolicyDashboardsResponseBody
	GetTotal() *int32
}

type ListIntegrationPolicyDashboardsResponseBody struct {
	// The list of dashboards.
	Dashboards []*ListIntegrationPolicyDashboardsResponseBodyDashboards `json:"dashboards,omitempty" xml:"dashboards,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of components.
	//
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListIntegrationPolicyDashboardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyDashboardsResponseBody) GetDashboards() []*ListIntegrationPolicyDashboardsResponseBodyDashboards {
	return s.Dashboards
}

func (s *ListIntegrationPolicyDashboardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationPolicyDashboardsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListIntegrationPolicyDashboardsResponseBody) SetDashboards(v []*ListIntegrationPolicyDashboardsResponseBodyDashboards) *ListIntegrationPolicyDashboardsResponseBody {
	s.Dashboards = v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBody) SetRequestId(v string) *ListIntegrationPolicyDashboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBody) SetTotal(v int32) *ListIntegrationPolicyDashboardsResponseBody {
	s.Total = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBody) Validate() error {
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

type ListIntegrationPolicyDashboardsResponseBodyDashboards struct {
	// The dashboard engine. Valid values:
	//
	// grafana: shared Grafana.
	//
	// cms: the self-developed dashboard engine of CloudMonitor.
	//
	// example:
	//
	// grafana
	Engine *string `json:"engine,omitempty" xml:"engine,omitempty"`
	// The UID of the dashboard folder.
	//
	// example:
	//
	// Env-AAA
	FolderUid *string `json:"folderUid,omitempty" xml:"folderUid,omitempty"`
	// The dashboard name.
	//
	// example:
	//
	// k8s-pod
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hongkong
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The list of tags.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The title of the UI module. This is different from the name parameter.
	//
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The ID of the current Alibaba Cloud account. This parameter is read-only.
	//
	// example:
	//
	// 1258199346721590
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// The webhook URL for PagerDuty integration. V1 and V2 are supported.
	//
	// example:
	//
	// https://agi.alicdn.com/user/0/0_0_5255362686.png?x-oss-process=image/quality,q_75/format,jpg&file=1734574878007.jpg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListIntegrationPolicyDashboardsResponseBodyDashboards) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyDashboardsResponseBodyDashboards) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) GetEngine() *string {
	return s.Engine
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) GetFolderUid() *string {
	return s.FolderUid
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) GetRegion() *string {
	return s.Region
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) GetTags() []*string {
	return s.Tags
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) GetTitle() *string {
	return s.Title
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) GetUid() *string {
	return s.Uid
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) GetUrl() *string {
	return s.Url
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) SetEngine(v string) *ListIntegrationPolicyDashboardsResponseBodyDashboards {
	s.Engine = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) SetFolderUid(v string) *ListIntegrationPolicyDashboardsResponseBodyDashboards {
	s.FolderUid = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) SetName(v string) *ListIntegrationPolicyDashboardsResponseBodyDashboards {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) SetRegion(v string) *ListIntegrationPolicyDashboardsResponseBodyDashboards {
	s.Region = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) SetTags(v []*string) *ListIntegrationPolicyDashboardsResponseBodyDashboards {
	s.Tags = v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) SetTitle(v string) *ListIntegrationPolicyDashboardsResponseBodyDashboards {
	s.Title = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) SetUid(v string) *ListIntegrationPolicyDashboardsResponseBodyDashboards {
	s.Uid = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) SetUrl(v string) *ListIntegrationPolicyDashboardsResponseBodyDashboards {
	s.Url = &v
	return s
}

func (s *ListIntegrationPolicyDashboardsResponseBodyDashboards) Validate() error {
	return dara.Validate(s)
}
