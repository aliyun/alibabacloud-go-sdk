// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAbnormalEventsCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetAbnormalEventsCountRequest
	GetXDebugId() *string
	SetCluster(v string) *GetAbnormalEventsCountRequest
	GetCluster() *string
	SetEnd(v float32) *GetAbnormalEventsCountRequest
	GetEnd() *float32
	SetInstance(v string) *GetAbnormalEventsCountRequest
	GetInstance() *string
	SetLevel(v string) *GetAbnormalEventsCountRequest
	GetLevel() *string
	SetNamespace(v string) *GetAbnormalEventsCountRequest
	GetNamespace() *string
	SetPod(v string) *GetAbnormalEventsCountRequest
	GetPod() *string
	SetShowPod(v int32) *GetAbnormalEventsCountRequest
	GetShowPod() *int32
	SetStart(v float32) *GetAbnormalEventsCountRequest
	GetStart() *float32
	SetXSysomInvokeSource(v string) *GetAbnormalEventsCountRequest
	GetXSysomInvokeSource() *string
}

type GetAbnormalEventsCountRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 2ijff4be-bf24-4070-89ca-c47c879b0g32
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The severity level of abnormal events.
	//
	// example:
	//
	// potential
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The namespace of the pod.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The pod name.
	//
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
	// Specifies whether to display abnormal events of the pod.
	//
	// example:
	//
	// 1
	ShowPod *int32 `json:"showPod,omitempty" xml:"showPod,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1725797727754
	Start              *float32 `json:"start,omitempty" xml:"start,omitempty"`
	XSysomInvokeSource *string  `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetAbnormalEventsCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAbnormalEventsCountRequest) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetAbnormalEventsCountRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetAbnormalEventsCountRequest) GetEnd() *float32 {
	return s.End
}

func (s *GetAbnormalEventsCountRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetAbnormalEventsCountRequest) GetLevel() *string {
	return s.Level
}

func (s *GetAbnormalEventsCountRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetAbnormalEventsCountRequest) GetPod() *string {
	return s.Pod
}

func (s *GetAbnormalEventsCountRequest) GetShowPod() *int32 {
	return s.ShowPod
}

func (s *GetAbnormalEventsCountRequest) GetStart() *float32 {
	return s.Start
}

func (s *GetAbnormalEventsCountRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetAbnormalEventsCountRequest) SetXDebugId(v string) *GetAbnormalEventsCountRequest {
	s.XDebugId = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetCluster(v string) *GetAbnormalEventsCountRequest {
	s.Cluster = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetEnd(v float32) *GetAbnormalEventsCountRequest {
	s.End = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetInstance(v string) *GetAbnormalEventsCountRequest {
	s.Instance = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetLevel(v string) *GetAbnormalEventsCountRequest {
	s.Level = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetNamespace(v string) *GetAbnormalEventsCountRequest {
	s.Namespace = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetPod(v string) *GetAbnormalEventsCountRequest {
	s.Pod = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetShowPod(v int32) *GetAbnormalEventsCountRequest {
	s.ShowPod = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetStart(v float32) *GetAbnormalEventsCountRequest {
	s.Start = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetXSysomInvokeSource(v string) *GetAbnormalEventsCountRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) Validate() error {
	return dara.Validate(s)
}
