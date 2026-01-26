// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetPrometheusMonitoringResponseBody
	GetCode() *int32
	SetData(v *GetPrometheusMonitoringResponseBodyData) *GetPrometheusMonitoringResponseBody
	GetData() *GetPrometheusMonitoringResponseBodyData
	SetMessage(v string) *GetPrometheusMonitoringResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPrometheusMonitoringResponseBody
	GetRequestId() *string
}

type GetPrometheusMonitoringResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *GetPrometheusMonitoringResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPrometheusMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusMonitoringResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPrometheusMonitoringResponseBody) GetData() *GetPrometheusMonitoringResponseBodyData {
	return s.Data
}

func (s *GetPrometheusMonitoringResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPrometheusMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrometheusMonitoringResponseBody) SetCode(v int32) *GetPrometheusMonitoringResponseBody {
	s.Code = &v
	return s
}

func (s *GetPrometheusMonitoringResponseBody) SetData(v *GetPrometheusMonitoringResponseBodyData) *GetPrometheusMonitoringResponseBody {
	s.Data = v
	return s
}

func (s *GetPrometheusMonitoringResponseBody) SetMessage(v string) *GetPrometheusMonitoringResponseBody {
	s.Message = &v
	return s
}

func (s *GetPrometheusMonitoringResponseBody) SetRequestId(v string) *GetPrometheusMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusMonitoringResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPrometheusMonitoringResponseBodyData struct {
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
	// customJob1
	MonitoringName *string `json:"MonitoringName,omitempty" xml:"MonitoringName,omitempty"`
	// The status of the monitoring configuration. Valid values: run and stop.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the monitoring configuration. Valid values: serviceMonitor, podMonitor, customJob, and probe.
	//
	// example:
	//
	// serviceMonitor
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPrometheusMonitoringResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusMonitoringResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPrometheusMonitoringResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetPrometheusMonitoringResponseBodyData) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *GetPrometheusMonitoringResponseBodyData) GetMonitoringName() *string {
	return s.MonitoringName
}

func (s *GetPrometheusMonitoringResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetPrometheusMonitoringResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetPrometheusMonitoringResponseBodyData) SetClusterId(v string) *GetPrometheusMonitoringResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetPrometheusMonitoringResponseBodyData) SetConfigYaml(v string) *GetPrometheusMonitoringResponseBodyData {
	s.ConfigYaml = &v
	return s
}

func (s *GetPrometheusMonitoringResponseBodyData) SetMonitoringName(v string) *GetPrometheusMonitoringResponseBodyData {
	s.MonitoringName = &v
	return s
}

func (s *GetPrometheusMonitoringResponseBodyData) SetStatus(v string) *GetPrometheusMonitoringResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPrometheusMonitoringResponseBodyData) SetType(v string) *GetPrometheusMonitoringResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetPrometheusMonitoringResponseBodyData) Validate() error {
	return dara.Validate(s)
}
