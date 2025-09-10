// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusDashboardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusDashboards(v []*ListPrometheusDashboardsResponseBodyPrometheusDashboards) *ListPrometheusDashboardsResponseBody
	GetPrometheusDashboards() []*ListPrometheusDashboardsResponseBodyPrometheusDashboards
	SetRequestId(v string) *ListPrometheusDashboardsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPrometheusDashboardsResponseBody
	GetTotalCount() *int32
}

type ListPrometheusDashboardsResponseBody struct {
	PrometheusDashboards []*ListPrometheusDashboardsResponseBodyPrometheusDashboards `json:"prometheusDashboards,omitempty" xml:"prometheusDashboards,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 66
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPrometheusDashboardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusDashboardsResponseBody) GetPrometheusDashboards() []*ListPrometheusDashboardsResponseBodyPrometheusDashboards {
	return s.PrometheusDashboards
}

func (s *ListPrometheusDashboardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusDashboardsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPrometheusDashboardsResponseBody) SetPrometheusDashboards(v []*ListPrometheusDashboardsResponseBodyPrometheusDashboards) *ListPrometheusDashboardsResponseBody {
	s.PrometheusDashboards = v
	return s
}

func (s *ListPrometheusDashboardsResponseBody) SetRequestId(v string) *ListPrometheusDashboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusDashboardsResponseBody) SetTotalCount(v int32) *ListPrometheusDashboardsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPrometheusDashboardsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusDashboardsResponseBodyPrometheusDashboards struct {
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// pod
	Name *string   `json:"name,omitempty" xml:"name,omitempty"`
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// ceshi
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1987395500251724
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// example:
	//
	// https://gnew.console.aliyun.com/d/xxx-17460385-807-7-6/cs-cost-application
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListPrometheusDashboardsResponseBodyPrometheusDashboards) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusDashboardsResponseBodyPrometheusDashboards) GoString() string {
	return s.String()
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) GetId() *string {
	return s.Id
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) GetName() *string {
	return s.Name
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) GetTags() []*string {
	return s.Tags
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) GetTitle() *string {
	return s.Title
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) GetUid() *string {
	return s.Uid
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) GetUrl() *string {
	return s.Url
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) SetId(v string) *ListPrometheusDashboardsResponseBodyPrometheusDashboards {
	s.Id = &v
	return s
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) SetName(v string) *ListPrometheusDashboardsResponseBodyPrometheusDashboards {
	s.Name = &v
	return s
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) SetTags(v []*string) *ListPrometheusDashboardsResponseBodyPrometheusDashboards {
	s.Tags = v
	return s
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) SetTitle(v string) *ListPrometheusDashboardsResponseBodyPrometheusDashboards {
	s.Title = &v
	return s
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) SetUid(v string) *ListPrometheusDashboardsResponseBodyPrometheusDashboards {
	s.Uid = &v
	return s
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) SetUrl(v string) *ListPrometheusDashboardsResponseBodyPrometheusDashboards {
	s.Url = &v
	return s
}

func (s *ListPrometheusDashboardsResponseBodyPrometheusDashboards) Validate() error {
	return dara.Validate(s)
}
