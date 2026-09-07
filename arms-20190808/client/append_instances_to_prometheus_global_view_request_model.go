// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendInstancesToPrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v string) *AppendInstancesToPrometheusGlobalViewRequest
	GetClusters() *string
	SetGlobalViewClusterId(v string) *AppendInstancesToPrometheusGlobalViewRequest
	GetGlobalViewClusterId() *string
	SetGroupName(v string) *AppendInstancesToPrometheusGlobalViewRequest
	GetGroupName() *string
	SetRegionId(v string) *AppendInstancesToPrometheusGlobalViewRequest
	GetRegionId() *string
}

type AppendInstancesToPrometheusGlobalViewRequest struct {
	// The list of global aggregation instances. The value is a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [         {             "sourceName": "数据源名称- ArmsPrometheus No.1",             "sourceType":"AlibabaPrometheus",             "userId":"UserID",             "clusterId":"ClusterId",         },         {             "sourceName": "数据源名称 - MetrcStore No.2",             "sourceType":"MetricStore",             "dataSource":"MetricStore的 remote read 地址",             "extras":{                 "username":"BasicAuthUsername",                 "password":"BasicAuthPassword"             }         },         {             "sourceName": "Custom ",             "sourceType":"CustomPrometheus",             "dataSource":"自建Prometheus数据源 remoteread地址",             "extras":{                 "username":"BasicAuthUsername",                 "password":"BasicAuthPassword"             }         },         {           	"sourceName": "Other one ",             "sourceType":"Others",             "dataSource":"其他数据源如Tencent remoteread地址",             "headers":{                 "AnyHeaderToFill":"需要填充的Headers"             },             "extras":{                 "username":"BasicAuthUsername",                 "password":"BasicAuthPassword"             }         }   // ....... more addre ]
	Clusters *string `json:"Clusters,omitempty" xml:"Clusters,omitempty"`
	// The ID of the global aggregation instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// global-v2-cn-1670100631025794-6gjc0qgb
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	// The name of the global aggregation instance.
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
}

func (s AppendInstancesToPrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s AppendInstancesToPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) GetClusters() *string {
	return s.Clusters
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) GetGlobalViewClusterId() *string {
	return s.GlobalViewClusterId
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) SetClusters(v string) *AppendInstancesToPrometheusGlobalViewRequest {
	s.Clusters = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *AppendInstancesToPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) SetGroupName(v string) *AppendInstancesToPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) SetRegionId(v string) *AppendInstancesToPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) Validate() error {
	return dara.Validate(s)
}
