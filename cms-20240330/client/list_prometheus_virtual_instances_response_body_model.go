// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusVirtualInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListPrometheusVirtualInstancesResponseBodyInstances) *ListPrometheusVirtualInstancesResponseBody
	GetInstances() []*ListPrometheusVirtualInstancesResponseBodyInstances
	SetMaxResults(v string) *ListPrometheusVirtualInstancesResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListPrometheusVirtualInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPrometheusVirtualInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListPrometheusVirtualInstancesResponseBody
	GetTotalCount() *string
}

type ListPrometheusVirtualInstancesResponseBody struct {
	// Instance information.
	Instances  []*ListPrometheusVirtualInstancesResponseBodyInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	MaxResults *string                                                `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *string `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPrometheusVirtualInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusVirtualInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusVirtualInstancesResponseBody) GetInstances() []*ListPrometheusVirtualInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListPrometheusVirtualInstancesResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListPrometheusVirtualInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrometheusVirtualInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusVirtualInstancesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListPrometheusVirtualInstancesResponseBody) SetInstances(v []*ListPrometheusVirtualInstancesResponseBodyInstances) *ListPrometheusVirtualInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBody) SetMaxResults(v string) *ListPrometheusVirtualInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBody) SetNextToken(v string) *ListPrometheusVirtualInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBody) SetRequestId(v string) *ListPrometheusVirtualInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBody) SetTotalCount(v string) *ListPrometheusVirtualInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusVirtualInstancesResponseBodyInstances struct {
	// Creation time
	//
	// example:
	//
	// 1750315319946
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// HTTP API URL.
	//
	// example:
	//
	// http://xxxxxxxx
	HttpApiUrl *string `json:"httpApiUrl,omitempty" xml:"httpApiUrl,omitempty"`
	// Applicable data source type: PROMETHEUS_DS
	//
	// Prometheus instance ID
	//
	// example:
	//
	// rw-b8cfbbe94194ac37fe83f3d2d16a
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Applicable query type: CMS_BASIC_QUERY.
	//
	// Namespace of the metric
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai-cloudspe
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// User ID.
	//
	// example:
	//
	// 17073812345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListPrometheusVirtualInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusVirtualInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) GetHttpApiUrl() *string {
	return s.HttpApiUrl
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) GetNamespace() *string {
	return s.Namespace
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) GetUserId() *string {
	return s.UserId
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) SetCreatedAt(v string) *ListPrometheusVirtualInstancesResponseBodyInstances {
	s.CreatedAt = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) SetHttpApiUrl(v string) *ListPrometheusVirtualInstancesResponseBodyInstances {
	s.HttpApiUrl = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) SetInstanceId(v string) *ListPrometheusVirtualInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) SetNamespace(v string) *ListPrometheusVirtualInstancesResponseBodyInstances {
	s.Namespace = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) SetRegionId(v string) *ListPrometheusVirtualInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) SetUserId(v string) *ListPrometheusVirtualInstancesResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
