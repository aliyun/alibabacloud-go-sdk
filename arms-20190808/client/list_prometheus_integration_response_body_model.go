// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListPrometheusIntegrationResponseBody
	GetCode() *int32
	SetData(v []*ListPrometheusIntegrationResponseBodyData) *ListPrometheusIntegrationResponseBody
	GetData() []*ListPrometheusIntegrationResponseBodyData
	SetMessage(v string) *ListPrometheusIntegrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrometheusIntegrationResponseBody
	GetRequestId() *string
}

type ListPrometheusIntegrationResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried exporters.
	Data []*ListPrometheusIntegrationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1F1D8840-5330-5804-A8DB-C3C5C5CED6BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusIntegrationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPrometheusIntegrationResponseBody) GetData() []*ListPrometheusIntegrationResponseBodyData {
	return s.Data
}

func (s *ListPrometheusIntegrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrometheusIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusIntegrationResponseBody) SetCode(v int32) *ListPrometheusIntegrationResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBody) SetData(v []*ListPrometheusIntegrationResponseBodyData) *ListPrometheusIntegrationResponseBody {
	s.Data = v
	return s
}

func (s *ListPrometheusIntegrationResponseBody) SetMessage(v string) *ListPrometheusIntegrationResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBody) SetRequestId(v string) *ListPrometheusIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusIntegrationResponseBodyData struct {
	// Indicates whether the exporter can be deleted.
	//
	// example:
	//
	// true
	CanDelete *bool `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	// Indicates whether the exporter can be edited.
	//
	// example:
	//
	// true
	CanEditor *bool `json:"CanEditor,omitempty" xml:"CanEditor,omitempty"`
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// c589a1b8db05c4561aefbb898ca8fb1cf
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// container-1
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The description of the exporter.
	//
	// example:
	//
	// "{}"
	Describe *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// The type of the exporter.
	//
	// example:
	//
	// kafka-exporter
	ExporterType *string `json:"ExporterType,omitempty" xml:"ExporterType,omitempty"`
	// The ID of the exporter.
	//
	// example:
	//
	// 29374
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the exporter.
	//
	// example:
	//
	// inet
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The integration type. Valid values: kafka, mysql, redis, snmp, emr, nubela, and tidb.
	//
	// example:
	//
	// Kafka, mysql, redis, snmp, emr, nubela, and tidb
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Indicates whether an upgrade is required.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The configurations of the exporter. The value is a JSON string.
	//
	// example:
	//
	// {
	//
	//       "port": "5554",
	//
	//       "name": "kafka-test12",
	//
	//       "kafka_instance": "kafka-test",
	//
	//       "__label_value": "kafka-test",
	//
	//       "scrape_interval": 33,
	//
	//       "metrics_path": "/metrics",
	//
	//       "__label_key": "kafka-test"
	//
	// }
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The pod name of the exporter.
	//
	// example:
	//
	// kafka-exporter-1
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// Indicates whether the description is displayed.
	//
	// example:
	//
	// true
	ShowDescribe *bool `json:"ShowDescribe,omitempty" xml:"ShowDescribe,omitempty"`
	// Indicates whether the exporter logs are displayed.
	//
	// example:
	//
	// true
	ShowLog *bool `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	// The status of the exporter.
	//
	// example:
	//
	// installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The monitored IP address.
	//
	// example:
	//
	// 121.40.62.240:3342
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPrometheusIntegrationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusIntegrationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrometheusIntegrationResponseBodyData) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ListPrometheusIntegrationResponseBodyData) GetCanEditor() *bool {
	return s.CanEditor
}

func (s *ListPrometheusIntegrationResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusIntegrationResponseBodyData) GetContainerName() *string {
	return s.ContainerName
}

func (s *ListPrometheusIntegrationResponseBodyData) GetDescribe() *string {
	return s.Describe
}

func (s *ListPrometheusIntegrationResponseBodyData) GetExporterType() *string {
	return s.ExporterType
}

func (s *ListPrometheusIntegrationResponseBodyData) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListPrometheusIntegrationResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListPrometheusIntegrationResponseBodyData) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *ListPrometheusIntegrationResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListPrometheusIntegrationResponseBodyData) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *ListPrometheusIntegrationResponseBodyData) GetParam() *string {
	return s.Param
}

func (s *ListPrometheusIntegrationResponseBodyData) GetPodName() *string {
	return s.PodName
}

func (s *ListPrometheusIntegrationResponseBodyData) GetShowDescribe() *bool {
	return s.ShowDescribe
}

func (s *ListPrometheusIntegrationResponseBodyData) GetShowLog() *bool {
	return s.ShowLog
}

func (s *ListPrometheusIntegrationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListPrometheusIntegrationResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *ListPrometheusIntegrationResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListPrometheusIntegrationResponseBodyData) SetCanDelete(v bool) *ListPrometheusIntegrationResponseBodyData {
	s.CanDelete = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetCanEditor(v bool) *ListPrometheusIntegrationResponseBodyData {
	s.CanEditor = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetClusterId(v string) *ListPrometheusIntegrationResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetContainerName(v string) *ListPrometheusIntegrationResponseBodyData {
	s.ContainerName = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetDescribe(v string) *ListPrometheusIntegrationResponseBodyData {
	s.Describe = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetExporterType(v string) *ListPrometheusIntegrationResponseBodyData {
	s.ExporterType = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetInstanceId(v int64) *ListPrometheusIntegrationResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetInstanceName(v string) *ListPrometheusIntegrationResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetIntegrationType(v string) *ListPrometheusIntegrationResponseBodyData {
	s.IntegrationType = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetNamespace(v string) *ListPrometheusIntegrationResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetNeedUpgrade(v bool) *ListPrometheusIntegrationResponseBodyData {
	s.NeedUpgrade = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetParam(v string) *ListPrometheusIntegrationResponseBodyData {
	s.Param = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetPodName(v string) *ListPrometheusIntegrationResponseBodyData {
	s.PodName = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetShowDescribe(v bool) *ListPrometheusIntegrationResponseBodyData {
	s.ShowDescribe = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetShowLog(v bool) *ListPrometheusIntegrationResponseBodyData {
	s.ShowLog = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetStatus(v string) *ListPrometheusIntegrationResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetTarget(v string) *ListPrometheusIntegrationResponseBodyData {
	s.Target = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) SetVersion(v string) *ListPrometheusIntegrationResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListPrometheusIntegrationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
