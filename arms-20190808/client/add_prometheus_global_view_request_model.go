// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v string) *AddPrometheusGlobalViewRequest
	GetClusters() *string
	SetGroupName(v string) *AddPrometheusGlobalViewRequest
	GetGroupName() *string
	SetRegionId(v string) *AddPrometheusGlobalViewRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AddPrometheusGlobalViewRequest
	GetResourceGroupId() *string
	SetTag(v []*AddPrometheusGlobalViewRequestTag) *AddPrometheusGlobalViewRequest
	GetTag() []*AddPrometheusGlobalViewRequestTag
}

type AddPrometheusGlobalViewRequest struct {
	// The queried global aggregation instances. The value is a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [         {             "sourceName": "Data source name- ArmsPrometheus No.1",             "sourceType":"AlibabaPrometheus",             "userId":"UserID",             "clusterId":"ClusterId",         },         {             "sourceName": "Data source name - MetrcStore No.2",             "sourceType":"MetricStore",             "dataSource":"MetricStore remote read address",             "extras":{                 "username":"BasicAuthUsername",                 "password":"BasicAuthPassword"             }         },         {             "sourceName": "Custom ",             "sourceType":"CustomPrometheus",             "dataSource":"Build your own Prometheus data source remoteread address",             "extras":{                 "username":"BasicAuthUsername",                 "password":"BasicAuthPassword"             }         },         {           	"sourceName": "Other one ",             "sourceType":"Others",             "dataSource":"Other data sources such as Tencent remoteread address",             "headers":{                 "AnyHeaderToFill":"Headers to be populated"             },             "extras":{                 "username":"BasicAuthUsername",                 "password":"BasicAuthPassword"             }         }   // ....... more addre ]
	Clusters *string `json:"Clusters,omitempty" xml:"Clusters,omitempty"`
	// The name of the aggregation instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// zyGlobalView
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2eq4pecazwfy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tag []*AddPrometheusGlobalViewRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AddPrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewRequest) GetClusters() *string {
	return s.Clusters
}

func (s *AddPrometheusGlobalViewRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddPrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPrometheusGlobalViewRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddPrometheusGlobalViewRequest) GetTag() []*AddPrometheusGlobalViewRequestTag {
	return s.Tag
}

func (s *AddPrometheusGlobalViewRequest) SetClusters(v string) *AddPrometheusGlobalViewRequest {
	s.Clusters = &v
	return s
}

func (s *AddPrometheusGlobalViewRequest) SetGroupName(v string) *AddPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *AddPrometheusGlobalViewRequest) SetRegionId(v string) *AddPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *AddPrometheusGlobalViewRequest) SetResourceGroupId(v string) *AddPrometheusGlobalViewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddPrometheusGlobalViewRequest) SetTag(v []*AddPrometheusGlobalViewRequestTag) *AddPrometheusGlobalViewRequest {
	s.Tag = v
	return s
}

func (s *AddPrometheusGlobalViewRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddPrometheusGlobalViewRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// http.status_code
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddPrometheusGlobalViewRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewRequestTag) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddPrometheusGlobalViewRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddPrometheusGlobalViewRequestTag) SetKey(v string) *AddPrometheusGlobalViewRequestTag {
	s.Key = &v
	return s
}

func (s *AddPrometheusGlobalViewRequestTag) SetValue(v string) *AddPrometheusGlobalViewRequestTag {
	s.Value = &v
	return s
}

func (s *AddPrometheusGlobalViewRequestTag) Validate() error {
	return dara.Validate(s)
}
