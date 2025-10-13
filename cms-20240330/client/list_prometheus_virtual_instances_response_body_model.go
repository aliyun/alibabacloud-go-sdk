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
	SetRequestId(v string) *ListPrometheusVirtualInstancesResponseBody
	GetRequestId() *string
}

type ListPrometheusVirtualInstancesResponseBody struct {
	Instances []*ListPrometheusVirtualInstancesResponseBodyInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *ListPrometheusVirtualInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusVirtualInstancesResponseBody) SetInstances(v []*ListPrometheusVirtualInstancesResponseBodyInstances) *ListPrometheusVirtualInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListPrometheusVirtualInstancesResponseBody) SetRequestId(v string) *ListPrometheusVirtualInstancesResponseBody {
	s.RequestId = &v
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
	// example:
	//
	// 1750315319946
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// http://xxxxxxxx
	HttpApiUrl *string `json:"httpApiUrl,omitempty" xml:"httpApiUrl,omitempty"`
	// example:
	//
	// rw-b8cfbbe94194ac37fe83f3d2d16a
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// arms-prom
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// cn-shanghai-cloudspe
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
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
