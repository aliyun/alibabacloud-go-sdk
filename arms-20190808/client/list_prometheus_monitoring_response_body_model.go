// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListPrometheusMonitoringResponseBody
	GetCode() *int32
	SetData(v []*ListPrometheusMonitoringResponseBodyData) *ListPrometheusMonitoringResponseBody
	GetData() []*ListPrometheusMonitoringResponseBodyData
	SetMessage(v string) *ListPrometheusMonitoringResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrometheusMonitoringResponseBody
	GetRequestId() *string
}

type ListPrometheusMonitoringResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data []*ListPrometheusMonitoringResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D80ADAAC-8C32-5479-BD14-C28CF832****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusMonitoringResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPrometheusMonitoringResponseBody) GetData() []*ListPrometheusMonitoringResponseBodyData {
	return s.Data
}

func (s *ListPrometheusMonitoringResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrometheusMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusMonitoringResponseBody) SetCode(v int32) *ListPrometheusMonitoringResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrometheusMonitoringResponseBody) SetData(v []*ListPrometheusMonitoringResponseBodyData) *ListPrometheusMonitoringResponseBody {
	s.Data = v
	return s
}

func (s *ListPrometheusMonitoringResponseBody) SetMessage(v string) *ListPrometheusMonitoringResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrometheusMonitoringResponseBody) SetRequestId(v string) *ListPrometheusMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusMonitoringResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusMonitoringResponseBodyData struct {
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// c589a1b8db05c4561aefbb898ca8fb1cf
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The monitoring configuration. The value is a YAML string.
	//
	// example:
	//
	// apiVersion: monitoring.coreos.com/v1
	//
	// kind: ServiceMonitor
	//
	// metadata:
	//
	//   name: tomcat-demo
	//
	//   namespace: default
	//
	// spec:
	//
	//   endpoints:
	//
	//     - interval: 30s
	//
	//       path: /metrics
	//
	//       port: tomcat-monitor
	//
	//   namespaceSelector:
	//
	//     any: true
	//
	//   selector:
	//
	//     matchLabels:
	//
	//       app: tomcat
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The name of the monitoring configuration.
	//
	// example:
	//
	// podMonitor1
	MonitoringName *string `json:"MonitoringName,omitempty" xml:"MonitoringName,omitempty"`
	// The status of the monitoring configuration.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the monitoring configuration.
	//
	// example:
	//
	// podMonitor
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusMonitoringResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusMonitoringResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrometheusMonitoringResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusMonitoringResponseBodyData) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *ListPrometheusMonitoringResponseBodyData) GetMonitoringName() *string {
	return s.MonitoringName
}

func (s *ListPrometheusMonitoringResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListPrometheusMonitoringResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListPrometheusMonitoringResponseBodyData) SetClusterId(v string) *ListPrometheusMonitoringResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusMonitoringResponseBodyData) SetConfigYaml(v string) *ListPrometheusMonitoringResponseBodyData {
	s.ConfigYaml = &v
	return s
}

func (s *ListPrometheusMonitoringResponseBodyData) SetMonitoringName(v string) *ListPrometheusMonitoringResponseBodyData {
	s.MonitoringName = &v
	return s
}

func (s *ListPrometheusMonitoringResponseBodyData) SetStatus(v string) *ListPrometheusMonitoringResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListPrometheusMonitoringResponseBodyData) SetType(v string) *ListPrometheusMonitoringResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListPrometheusMonitoringResponseBodyData) Validate() error {
	return dara.Validate(s)
}
