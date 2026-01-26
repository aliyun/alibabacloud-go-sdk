// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvServiceMonitorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvServiceMonitorsResponseBody
	GetCode() *int32
	SetData(v []*ListEnvServiceMonitorsResponseBodyData) *ListEnvServiceMonitorsResponseBody
	GetData() []*ListEnvServiceMonitorsResponseBodyData
	SetMessage(v string) *ListEnvServiceMonitorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvServiceMonitorsResponseBody
	GetRequestId() *string
}

type ListEnvServiceMonitorsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data []*ListEnvServiceMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 1A474FF8-7861-4D00-81B5-5BC3DA4E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnvServiceMonitorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvServiceMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvServiceMonitorsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvServiceMonitorsResponseBody) GetData() []*ListEnvServiceMonitorsResponseBodyData {
	return s.Data
}

func (s *ListEnvServiceMonitorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvServiceMonitorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvServiceMonitorsResponseBody) SetCode(v int32) *ListEnvServiceMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBody) SetData(v []*ListEnvServiceMonitorsResponseBodyData) *ListEnvServiceMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvServiceMonitorsResponseBody) SetMessage(v string) *ListEnvServiceMonitorsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBody) SetRequestId(v string) *ListEnvServiceMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBody) Validate() error {
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

type ListEnvServiceMonitorsResponseBodyData struct {
	// The name of the add-on to which the ServiceMonitor belongs.
	//
	// example:
	//
	// mysql
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The instance name of the add-on.
	//
	// example:
	//
	// mysql1
	AddonReleaseName *string `json:"AddonReleaseName,omitempty" xml:"AddonReleaseName,omitempty"`
	// The version of the add-on.
	//
	// example:
	//
	// 1.1.0
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// The YAML configuration string.
	//
	// example:
	//
	// Refer to supplementary instructions.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The time when the ServiceMonitor was created. The value of this parameter is a timestamp.
	//
	// example:
	//
	// 2011-10-11T22:32:11Z
	CreationTimestamp *string `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	// The endpoints of the ServiceMonitor.
	Endpoints []*ListEnvServiceMonitorsResponseBodyDataEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The number of matched services.
	//
	// example:
	//
	// 1
	MatchedServiceCount *int32 `json:"MatchedServiceCount,omitempty" xml:"MatchedServiceCount,omitempty"`
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the ServiceMonitor.
	//
	// example:
	//
	// sm1
	ServiceMonitorName *string `json:"ServiceMonitorName,omitempty" xml:"ServiceMonitorName,omitempty"`
	// The status of the ServiceMonitor.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEnvServiceMonitorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvServiceMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetAddonName() *string {
	return s.AddonName
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetCreationTimestamp() *string {
	return s.CreationTimestamp
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetEndpoints() []*ListEnvServiceMonitorsResponseBodyDataEndpoints {
	return s.Endpoints
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetMatchedServiceCount() *int32 {
	return s.MatchedServiceCount
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetServiceMonitorName() *string {
	return s.ServiceMonitorName
}

func (s *ListEnvServiceMonitorsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetAddonName(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.AddonName = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetAddonReleaseName(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.AddonReleaseName = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetAddonVersion(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.AddonVersion = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetConfigYaml(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.ConfigYaml = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetCreationTimestamp(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.CreationTimestamp = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetEndpoints(v []*ListEnvServiceMonitorsResponseBodyDataEndpoints) *ListEnvServiceMonitorsResponseBodyData {
	s.Endpoints = v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetEnvironmentId(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetMatchedServiceCount(v int32) *ListEnvServiceMonitorsResponseBodyData {
	s.MatchedServiceCount = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetNamespace(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetRegionId(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetServiceMonitorName(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.ServiceMonitorName = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) SetStatus(v string) *ListEnvServiceMonitorsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyData) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvServiceMonitorsResponseBodyDataEndpoints struct {
	// The collection interval.
	//
	// example:
	//
	// 30s
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of pods that match the ServiceMonitor endpoint.
	//
	// example:
	//
	// 1
	MatchedTargetCount *int32 `json:"MatchedTargetCount,omitempty" xml:"MatchedTargetCount,omitempty"`
	// The collection path.
	//
	// example:
	//
	// /metrics
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The external port.
	//
	// example:
	//
	// 9101
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 443
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s ListEnvServiceMonitorsResponseBodyDataEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListEnvServiceMonitorsResponseBodyDataEndpoints) GoString() string {
	return s.String()
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) GetInterval() *string {
	return s.Interval
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) GetMatchedTargetCount() *int32 {
	return s.MatchedTargetCount
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) GetPath() *string {
	return s.Path
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) GetPort() *string {
	return s.Port
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) SetInterval(v string) *ListEnvServiceMonitorsResponseBodyDataEndpoints {
	s.Interval = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) SetMatchedTargetCount(v int32) *ListEnvServiceMonitorsResponseBodyDataEndpoints {
	s.MatchedTargetCount = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) SetPath(v string) *ListEnvServiceMonitorsResponseBodyDataEndpoints {
	s.Path = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) SetPort(v string) *ListEnvServiceMonitorsResponseBodyDataEndpoints {
	s.Port = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) SetTargetPort(v int32) *ListEnvServiceMonitorsResponseBodyDataEndpoints {
	s.TargetPort = &v
	return s
}

func (s *ListEnvServiceMonitorsResponseBodyDataEndpoints) Validate() error {
	return dara.Validate(s)
}
