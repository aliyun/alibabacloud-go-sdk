// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvPodMonitorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvPodMonitorsResponseBody
	GetCode() *int32
	SetData(v []*ListEnvPodMonitorsResponseBodyData) *ListEnvPodMonitorsResponseBody
	GetData() []*ListEnvPodMonitorsResponseBodyData
	SetMessage(v string) *ListEnvPodMonitorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvPodMonitorsResponseBody
	GetRequestId() *string
}

type ListEnvPodMonitorsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	Data []*ListEnvPodMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 4C518054-852F-4023-ABC1-4AF95FF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnvPodMonitorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvPodMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvPodMonitorsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvPodMonitorsResponseBody) GetData() []*ListEnvPodMonitorsResponseBodyData {
	return s.Data
}

func (s *ListEnvPodMonitorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvPodMonitorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvPodMonitorsResponseBody) SetCode(v int32) *ListEnvPodMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBody) SetData(v []*ListEnvPodMonitorsResponseBodyData) *ListEnvPodMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvPodMonitorsResponseBody) SetMessage(v string) *ListEnvPodMonitorsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBody) SetRequestId(v string) *ListEnvPodMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEnvPodMonitorsResponseBodyData struct {
	// The name of the add-on to which the PodMonitor belongs.
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
	// 1.0.5
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// The YAML configuration string.
	//
	// example:
	//
	// Refer to supplementary instructions.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The time when the PodMonitor was created. The value of this parameter is a timestamp.
	//
	// example:
	//
	// 2011-01-02T11:34:22Z
	CreationTimestamp *string `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	// The endpoints of the PodMonitor.
	Endpoints []*ListEnvPodMonitorsResponseBodyDataEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the PodMonitor.
	//
	// example:
	//
	// pm1
	PodMonitorName *string `json:"PodMonitorName,omitempty" xml:"PodMonitorName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the PodMonitor.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEnvPodMonitorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvPodMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvPodMonitorsResponseBodyData) GetAddonName() *string {
	return s.AddonName
}

func (s *ListEnvPodMonitorsResponseBodyData) GetAddonReleaseName() *string {
	return s.AddonReleaseName
}

func (s *ListEnvPodMonitorsResponseBodyData) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *ListEnvPodMonitorsResponseBodyData) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *ListEnvPodMonitorsResponseBodyData) GetCreationTimestamp() *string {
	return s.CreationTimestamp
}

func (s *ListEnvPodMonitorsResponseBodyData) GetEndpoints() []*ListEnvPodMonitorsResponseBodyDataEndpoints {
	return s.Endpoints
}

func (s *ListEnvPodMonitorsResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvPodMonitorsResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListEnvPodMonitorsResponseBodyData) GetPodMonitorName() *string {
	return s.PodMonitorName
}

func (s *ListEnvPodMonitorsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvPodMonitorsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListEnvPodMonitorsResponseBodyData) SetAddonName(v string) *ListEnvPodMonitorsResponseBodyData {
	s.AddonName = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetAddonReleaseName(v string) *ListEnvPodMonitorsResponseBodyData {
	s.AddonReleaseName = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetAddonVersion(v string) *ListEnvPodMonitorsResponseBodyData {
	s.AddonVersion = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetConfigYaml(v string) *ListEnvPodMonitorsResponseBodyData {
	s.ConfigYaml = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetCreationTimestamp(v string) *ListEnvPodMonitorsResponseBodyData {
	s.CreationTimestamp = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetEndpoints(v []*ListEnvPodMonitorsResponseBodyDataEndpoints) *ListEnvPodMonitorsResponseBodyData {
	s.Endpoints = v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetEnvironmentId(v string) *ListEnvPodMonitorsResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetNamespace(v string) *ListEnvPodMonitorsResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetPodMonitorName(v string) *ListEnvPodMonitorsResponseBodyData {
	s.PodMonitorName = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetRegionId(v string) *ListEnvPodMonitorsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) SetStatus(v string) *ListEnvPodMonitorsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListEnvPodMonitorsResponseBodyDataEndpoints struct {
	// The collection interval.
	//
	// example:
	//
	// 30s
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of pods that match the PodMonitor endpoint.
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
	// 9182
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 3306
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s ListEnvPodMonitorsResponseBodyDataEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListEnvPodMonitorsResponseBodyDataEndpoints) GoString() string {
	return s.String()
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) GetInterval() *string {
	return s.Interval
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) GetMatchedTargetCount() *int32 {
	return s.MatchedTargetCount
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) GetPath() *string {
	return s.Path
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) GetPort() *string {
	return s.Port
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) SetInterval(v string) *ListEnvPodMonitorsResponseBodyDataEndpoints {
	s.Interval = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) SetMatchedTargetCount(v int32) *ListEnvPodMonitorsResponseBodyDataEndpoints {
	s.MatchedTargetCount = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) SetPath(v string) *ListEnvPodMonitorsResponseBodyDataEndpoints {
	s.Path = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) SetPort(v string) *ListEnvPodMonitorsResponseBodyDataEndpoints {
	s.Port = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) SetTargetPort(v int32) *ListEnvPodMonitorsResponseBodyDataEndpoints {
	s.TargetPort = &v
	return s
}

func (s *ListEnvPodMonitorsResponseBodyDataEndpoints) Validate() error {
	return dara.Validate(s)
}
