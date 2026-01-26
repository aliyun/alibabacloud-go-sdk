// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDashboardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDashboardVos(v []*ListDashboardsResponseBodyDashboardVos) *ListDashboardsResponseBody
	GetDashboardVos() []*ListDashboardsResponseBodyDashboardVos
	SetEnvironmentId(v string) *ListDashboardsResponseBody
	GetEnvironmentId() *string
	SetGrafanaServiceOpened(v string) *ListDashboardsResponseBody
	GetGrafanaServiceOpened() *string
	SetPrometheusServiceOpened(v string) *ListDashboardsResponseBody
	GetPrometheusServiceOpened() *string
	SetRequestId(v string) *ListDashboardsResponseBody
	GetRequestId() *string
}

type ListDashboardsResponseBody struct {
	// The information about the Grafana dashboard.
	DashboardVos []*ListDashboardsResponseBodyDashboardVos `json:"DashboardVos,omitempty" xml:"DashboardVos,omitempty" type:"Repeated"`
	// The ID of the environment instance.
	//
	// example:
	//
	// env-ebd54733482581fc8c4237******
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// Indicates whether Managed Service for Grafana is activated.
	//
	// example:
	//
	// true
	GrafanaServiceOpened *string `json:"GrafanaServiceOpened,omitempty" xml:"GrafanaServiceOpened,omitempty"`
	// Whether or not to turn on Prometheus service.
	//
	// example:
	//
	// true
	PrometheusServiceOpened *string `json:"PrometheusServiceOpened,omitempty" xml:"PrometheusServiceOpened,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A0CEDF1-06FE-44AC-8E21-21A5BE65****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDashboardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBody) GetDashboardVos() []*ListDashboardsResponseBodyDashboardVos {
	return s.DashboardVos
}

func (s *ListDashboardsResponseBody) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListDashboardsResponseBody) GetGrafanaServiceOpened() *string {
	return s.GrafanaServiceOpened
}

func (s *ListDashboardsResponseBody) GetPrometheusServiceOpened() *string {
	return s.PrometheusServiceOpened
}

func (s *ListDashboardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDashboardsResponseBody) SetDashboardVos(v []*ListDashboardsResponseBodyDashboardVos) *ListDashboardsResponseBody {
	s.DashboardVos = v
	return s
}

func (s *ListDashboardsResponseBody) SetEnvironmentId(v string) *ListDashboardsResponseBody {
	s.EnvironmentId = &v
	return s
}

func (s *ListDashboardsResponseBody) SetGrafanaServiceOpened(v string) *ListDashboardsResponseBody {
	s.GrafanaServiceOpened = &v
	return s
}

func (s *ListDashboardsResponseBody) SetPrometheusServiceOpened(v string) *ListDashboardsResponseBody {
	s.PrometheusServiceOpened = &v
	return s
}

func (s *ListDashboardsResponseBody) SetRequestId(v string) *ListDashboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardsResponseBody) Validate() error {
	if s.DashboardVos != nil {
		for _, item := range s.DashboardVos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDashboardsResponseBodyDashboardVos struct {
	// The type of the Grafana dashboard. This parameter has the same effect as the Exporter parameter whereas provides clearer implication.
	//
	// example:
	//
	// Node
	DashboardType *string `json:"DashboardType,omitempty" xml:"DashboardType,omitempty"`
	// The type of the exporter access source. Valid values:
	//
	// 	- Prometheus
	//
	// 	- Node
	//
	// 	- GPU
	//
	// 	- Redis
	//
	// 	- MySQL
	//
	// 	- Kafka
	//
	// 	- NGINX V2
	//
	// 	- Nginx
	//
	// 	- ZooKeeper
	//
	// 	- MongoDB
	//
	// 	- RabbitMQ
	//
	// 	- PostgreSQL
	//
	// 	- Kubernetes
	//
	// 	- Client Library
	//
	// 	- Elasticsearch
	//
	// 	- RocketMQ
	//
	// example:
	//
	// Nginx
	Exporter *string `json:"Exporter,omitempty" xml:"Exporter,omitempty"`
	// The URL of the Grafana dashboard.
	//
	// example:
	//
	// http://g.console.aliyun.com/d/1131971649496228-*****-59/ApiServer?orgId=3**&refresh=60s
	HttpUrl *string `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	// The URL of the Grafana dashboard.
	//
	// example:
	//
	// http://g.console.aliyun.com/d/1131971649496228-*****-59/ApiServer?orgId=3**&refresh=60s
	HttpsUrl *string `json:"HttpsUrl,omitempty" xml:"HttpsUrl,omitempty"`
	// The information about the Grafana dashboard.
	I18nChild *ListDashboardsResponseBodyDashboardVosI18nChild `json:"I18nChild,omitempty" xml:"I18nChild,omitempty" type:"Struct"`
	// The ID of the Grafana dashboard. The value is unique only when you install the Grafana dashboard.
	//
	// example:
	//
	// 1100**
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the exporter is provided by Application Real-Time Monitoring Service (ARMS).
	//
	// 	- `true:` The exporter is provided by ARMS.
	//
	// 	- `false:`: The exporter is not provided by ARMS.
	//
	// example:
	//
	// false
	IsArmsExporter *bool `json:"IsArmsExporter,omitempty" xml:"IsArmsExporter,omitempty"`
	// The category of the Grafana dashboard. Valid values: BASIC, THIRD, LIMIT, and CUSTOM.
	//
	// example:
	//
	// BASIC
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The language of the Grafana dashboard.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the Grafana dashboard. This parameter is different from the **Title*	- parameter as this parameter cannot be changed.
	//
	// example:
	//
	// k8s-node-overview
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the Grafana dashboard has a new version that is available for upgrade.
	//
	// example:
	//
	// false
	NeedUpdate *bool `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	// The tags of the Grafana dashboard.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the Grafana dashboard was created. The value is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1590136924
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The title of the Grafana dashboard.
	//
	// example:
	//
	// ApiServer
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The type of the Grafana dashboard. Valid values:
	//
	// 	- `dash-db`: a dashboard
	//
	// 	- `dash-folder`: a folder that can include a dashboard
	//
	// example:
	//
	// dash-db
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unique identifier of the Grafana dashboard.
	//
	// example:
	//
	// 1131971649496228-*****-59
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The complete URL of the Grafana dashboard.
	//
	// example:
	//
	// http://g.console.aliyun.com/d/1131971649496228-*****-59/ApiServer?orgId=3**&refresh=60s
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The version of the Grafana dashboard. The combination of version and name uniquely identifies a dashboard.
	//
	// example:
	//
	// v2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDashboardsResponseBodyDashboardVos) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsResponseBodyDashboardVos) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBodyDashboardVos) GetDashboardType() *string {
	return s.DashboardType
}

func (s *ListDashboardsResponseBodyDashboardVos) GetExporter() *string {
	return s.Exporter
}

func (s *ListDashboardsResponseBodyDashboardVos) GetHttpUrl() *string {
	return s.HttpUrl
}

func (s *ListDashboardsResponseBodyDashboardVos) GetHttpsUrl() *string {
	return s.HttpsUrl
}

func (s *ListDashboardsResponseBodyDashboardVos) GetI18nChild() *ListDashboardsResponseBodyDashboardVosI18nChild {
	return s.I18nChild
}

func (s *ListDashboardsResponseBodyDashboardVos) GetId() *string {
	return s.Id
}

func (s *ListDashboardsResponseBodyDashboardVos) GetIsArmsExporter() *bool {
	return s.IsArmsExporter
}

func (s *ListDashboardsResponseBodyDashboardVos) GetKind() *string {
	return s.Kind
}

func (s *ListDashboardsResponseBodyDashboardVos) GetLanguage() *string {
	return s.Language
}

func (s *ListDashboardsResponseBodyDashboardVos) GetName() *string {
	return s.Name
}

func (s *ListDashboardsResponseBodyDashboardVos) GetNeedUpdate() *bool {
	return s.NeedUpdate
}

func (s *ListDashboardsResponseBodyDashboardVos) GetTags() []*string {
	return s.Tags
}

func (s *ListDashboardsResponseBodyDashboardVos) GetTime() *string {
	return s.Time
}

func (s *ListDashboardsResponseBodyDashboardVos) GetTitle() *string {
	return s.Title
}

func (s *ListDashboardsResponseBodyDashboardVos) GetType() *string {
	return s.Type
}

func (s *ListDashboardsResponseBodyDashboardVos) GetUid() *string {
	return s.Uid
}

func (s *ListDashboardsResponseBodyDashboardVos) GetUrl() *string {
	return s.Url
}

func (s *ListDashboardsResponseBodyDashboardVos) GetVersion() *string {
	return s.Version
}

func (s *ListDashboardsResponseBodyDashboardVos) SetDashboardType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.DashboardType = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetExporter(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Exporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetHttpUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.HttpUrl = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetHttpsUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.HttpsUrl = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetI18nChild(v *ListDashboardsResponseBodyDashboardVosI18nChild) *ListDashboardsResponseBodyDashboardVos {
	s.I18nChild = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetId(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Id = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetIsArmsExporter(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.IsArmsExporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetKind(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Kind = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetLanguage(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Language = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetName(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Name = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetNeedUpdate(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.NeedUpdate = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTags(v []*string) *ListDashboardsResponseBodyDashboardVos {
	s.Tags = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTime(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Time = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTitle(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Title = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Type = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUid(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Uid = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Url = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetVersion(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Version = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) Validate() error {
	if s.I18nChild != nil {
		if err := s.I18nChild.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDashboardsResponseBodyDashboardVosI18nChild struct {
	// The type of the Grafana dashboard. This parameter has the same effect as the Exporter parameter whereas provides clearer implication.
	//
	// example:
	//
	// Node
	DashboardType *string `json:"DashboardType,omitempty" xml:"DashboardType,omitempty"`
	// The type of the exporter access source. Valid values:
	//
	// 	- Prometheus
	//
	// 	- Node
	//
	// 	- GPU
	//
	// 	- Redis
	//
	// 	- MySQL
	//
	// 	- Kafka
	//
	// 	- NGINX V2
	//
	// 	- Nginx
	//
	// 	- ZooKeeper
	//
	// 	- MongoDB
	//
	// 	- RabbitMQ
	//
	// 	- PostgreSQL
	//
	// 	- Kubernetes
	//
	// 	- Client Library
	//
	// 	- Elasticsearch
	//
	// 	- RocketMQ
	//
	// example:
	//
	// Nginx
	Exporter *string `json:"Exporter,omitempty" xml:"Exporter,omitempty"`
	// The URL of the Grafana dashboard.
	//
	// example:
	//
	// http://g.console.aliyun.com/d/1131971649496228-*****-59/ApiServer?orgId=3**&refresh=60s
	HttpUrl *string `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	// The URL of the Grafana dashboard.
	//
	// example:
	//
	// http://g.console.aliyun.com/d/1131971649496228-*****-59/ApiServer?orgId=3**&refresh=60s
	HttpsUrl *string `json:"HttpsUrl,omitempty" xml:"HttpsUrl,omitempty"`
	// The ID of the Grafana dashboard. The value is unique only when you install the Grafana dashboard.
	//
	// example:
	//
	// 1100**
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the exporter is provided by ARMS.
	//
	// 	- `true:` The exporter is provided by ARMS.
	//
	// 	- `false:`: The exporter is not provided by ARMS.
	//
	// example:
	//
	// false
	IsArmsExporter *bool `json:"IsArmsExporter,omitempty" xml:"IsArmsExporter,omitempty"`
	// The category of the Grafana dashboard. Valid values: BASIC, THIRD, LIMIT, and CUSTOM.
	//
	// example:
	//
	// BASIC
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The language of the Grafana dashboard.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the Grafana dashboard. This parameter is different from the **Title*	- parameter as this parameter cannot be changed.
	//
	// example:
	//
	// k8s-node-overview
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the Grafana dashboard has a new version that is available for upgrade.
	//
	// example:
	//
	// false
	NeedUpdate *bool `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	// The tags of the Grafana dashboard.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the Grafana dashboard was created. The value is a timestamp.
	//
	// example:
	//
	// 1590136924
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The title of the Grafana dashboard.
	//
	// example:
	//
	// ApiServer
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The type of the Grafana dashboard. Valid values:
	//
	// 	- `dash-db`: a dashboard
	//
	// 	- `dash-folder`: a folder that can include a dashboard
	//
	// example:
	//
	// dash-db
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unique identifier of the Grafana dashboard.
	//
	// example:
	//
	// 1131971649496228-*****-59
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The complete URL of the Grafana dashboard.
	//
	// example:
	//
	// http://g.console.aliyun.com/d/1131971649496228-*****-59/ApiServer?orgId=3**&refresh=60s
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The version of the Grafana dashboard. The combination of version and name uniquely identifies a dashboard.
	//
	// example:
	//
	// v2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDashboardsResponseBodyDashboardVosI18nChild) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsResponseBodyDashboardVosI18nChild) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetDashboardType() *string {
	return s.DashboardType
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetExporter() *string {
	return s.Exporter
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetHttpUrl() *string {
	return s.HttpUrl
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetHttpsUrl() *string {
	return s.HttpsUrl
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetId() *string {
	return s.Id
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetIsArmsExporter() *bool {
	return s.IsArmsExporter
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetKind() *string {
	return s.Kind
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetLanguage() *string {
	return s.Language
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetName() *string {
	return s.Name
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetNeedUpdate() *bool {
	return s.NeedUpdate
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetTags() []*string {
	return s.Tags
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetTime() *string {
	return s.Time
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetTitle() *string {
	return s.Title
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetType() *string {
	return s.Type
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetUid() *string {
	return s.Uid
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetUrl() *string {
	return s.Url
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) GetVersion() *string {
	return s.Version
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetDashboardType(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.DashboardType = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetExporter(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Exporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetHttpUrl(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.HttpUrl = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetHttpsUrl(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.HttpsUrl = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetId(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Id = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetIsArmsExporter(v bool) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.IsArmsExporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetKind(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Kind = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetLanguage(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Language = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetName(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Name = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetNeedUpdate(v bool) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.NeedUpdate = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetTags(v []*string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Tags = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetTime(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Time = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetTitle(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Title = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetType(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Type = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetUid(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Uid = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetUrl(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Url = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetVersion(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Version = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) Validate() error {
	return dara.Validate(s)
}
