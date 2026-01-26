// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetPrometheusIntegrationResponseBody
	GetCode() *int32
	SetData(v *GetPrometheusIntegrationResponseBodyData) *GetPrometheusIntegrationResponseBody
	GetData() *GetPrometheusIntegrationResponseBodyData
	SetMessage(v string) *GetPrometheusIntegrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPrometheusIntegrationResponseBody
	GetRequestId() *string
}

type GetPrometheusIntegrationResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. If another status code is returned, the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *GetPrometheusIntegrationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9BEF2832-9D95-5E3E-9B10-74887CA17B94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPrometheusIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusIntegrationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPrometheusIntegrationResponseBody) GetData() *GetPrometheusIntegrationResponseBodyData {
	return s.Data
}

func (s *GetPrometheusIntegrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPrometheusIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrometheusIntegrationResponseBody) SetCode(v int32) *GetPrometheusIntegrationResponseBody {
	s.Code = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBody) SetData(v *GetPrometheusIntegrationResponseBodyData) *GetPrometheusIntegrationResponseBody {
	s.Data = v
	return s
}

func (s *GetPrometheusIntegrationResponseBody) SetMessage(v string) *GetPrometheusIntegrationResponseBody {
	s.Message = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBody) SetRequestId(v string) *GetPrometheusIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPrometheusIntegrationResponseBodyData struct {
	// Indicates whether the exporter can be deleted.
	//
	// example:
	//
	// true
	CanDelete *bool `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	// Indicates whether the exporter can be modified.
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
	// The container name.
	//
	// example:
	//
	// kafka-exporter-1
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
	// tidb-exporter
	ExporterType *string `json:"ExporterType,omitempty" xml:"ExporterType,omitempty"`
	// The ID of the exporter.
	//
	// example:
	//
	// 2893
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the exporter.
	//
	// example:
	//
	// lpd-skyeye
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The integration type. Valid values: kafka, mysql, redis, snmp, emr, nubela, and tidb.
	//
	// example:
	//
	// tidb
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
	// True
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The parameters of the exporter. Format: JSON string.
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
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
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
	// 127.0.0.1:3422
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The version information.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetPrometheusIntegrationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusIntegrationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPrometheusIntegrationResponseBodyData) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *GetPrometheusIntegrationResponseBodyData) GetCanEditor() *bool {
	return s.CanEditor
}

func (s *GetPrometheusIntegrationResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetPrometheusIntegrationResponseBodyData) GetContainerName() *string {
	return s.ContainerName
}

func (s *GetPrometheusIntegrationResponseBodyData) GetDescribe() *string {
	return s.Describe
}

func (s *GetPrometheusIntegrationResponseBodyData) GetExporterType() *string {
	return s.ExporterType
}

func (s *GetPrometheusIntegrationResponseBodyData) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetPrometheusIntegrationResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetPrometheusIntegrationResponseBodyData) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *GetPrometheusIntegrationResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *GetPrometheusIntegrationResponseBodyData) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *GetPrometheusIntegrationResponseBodyData) GetParam() *string {
	return s.Param
}

func (s *GetPrometheusIntegrationResponseBodyData) GetShowDescribe() *bool {
	return s.ShowDescribe
}

func (s *GetPrometheusIntegrationResponseBodyData) GetShowLog() *string {
	return s.ShowLog
}

func (s *GetPrometheusIntegrationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetPrometheusIntegrationResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *GetPrometheusIntegrationResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetPrometheusIntegrationResponseBodyData) SetCanDelete(v bool) *GetPrometheusIntegrationResponseBodyData {
	s.CanDelete = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetCanEditor(v bool) *GetPrometheusIntegrationResponseBodyData {
	s.CanEditor = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetClusterId(v string) *GetPrometheusIntegrationResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetContainerName(v string) *GetPrometheusIntegrationResponseBodyData {
	s.ContainerName = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetDescribe(v string) *GetPrometheusIntegrationResponseBodyData {
	s.Describe = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetExporterType(v string) *GetPrometheusIntegrationResponseBodyData {
	s.ExporterType = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetInstanceId(v int64) *GetPrometheusIntegrationResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetInstanceName(v string) *GetPrometheusIntegrationResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetIntegrationType(v string) *GetPrometheusIntegrationResponseBodyData {
	s.IntegrationType = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetNamespace(v string) *GetPrometheusIntegrationResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetNeedUpgrade(v bool) *GetPrometheusIntegrationResponseBodyData {
	s.NeedUpgrade = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetParam(v string) *GetPrometheusIntegrationResponseBodyData {
	s.Param = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetShowDescribe(v bool) *GetPrometheusIntegrationResponseBodyData {
	s.ShowDescribe = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetShowLog(v string) *GetPrometheusIntegrationResponseBodyData {
	s.ShowLog = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetStatus(v string) *GetPrometheusIntegrationResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetTarget(v string) *GetPrometheusIntegrationResponseBodyData {
	s.Target = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) SetVersion(v string) *GetPrometheusIntegrationResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetPrometheusIntegrationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
