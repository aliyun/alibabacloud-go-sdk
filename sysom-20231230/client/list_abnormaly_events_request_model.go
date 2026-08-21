// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbnormalyEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListAbnormalyEventsRequest
	GetXDebugId() *string
	SetCluster(v string) *ListAbnormalyEventsRequest
	GetCluster() *string
	SetCurrent(v int32) *ListAbnormalyEventsRequest
	GetCurrent() *int32
	SetEnd(v float32) *ListAbnormalyEventsRequest
	GetEnd() *float32
	SetEvent(v string) *ListAbnormalyEventsRequest
	GetEvent() *string
	SetInstance(v string) *ListAbnormalyEventsRequest
	GetInstance() *string
	SetLevel(v string) *ListAbnormalyEventsRequest
	GetLevel() *string
	SetNamespace(v string) *ListAbnormalyEventsRequest
	GetNamespace() *string
	SetPageSize(v int32) *ListAbnormalyEventsRequest
	GetPageSize() *int32
	SetPod(v string) *ListAbnormalyEventsRequest
	GetPod() *string
	SetShowPod(v int32) *ListAbnormalyEventsRequest
	GetShowPod() *int32
	SetStart(v float32) *ListAbnormalyEventsRequest
	GetStart() *float32
	SetXSysomInvokeSource(v string) *ListAbnormalyEventsRequest
	GetXSysomInvokeSource() *string
}

type ListAbnormalyEventsRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 2d33d4be-bf54-4070-82ca-c1dc2d8b1562
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The current page number. This parameter is present during paginated queries.
	//
	// example:
	//
	// 1
	Current *int32 `json:"current,omitempty" xml:"current,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// The name of the anomaly event.
	//
	// example:
	//
	// 节点根文件系统使用检测
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The level of the anomaly event.
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
	// The number of entries per page. Default value: 5. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The pod name.
	//
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
	// Specifies whether to display pod anomaly events.
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

func (s ListAbnormalyEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAbnormalyEventsRequest) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListAbnormalyEventsRequest) GetCluster() *string {
	return s.Cluster
}

func (s *ListAbnormalyEventsRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAbnormalyEventsRequest) GetEnd() *float32 {
	return s.End
}

func (s *ListAbnormalyEventsRequest) GetEvent() *string {
	return s.Event
}

func (s *ListAbnormalyEventsRequest) GetInstance() *string {
	return s.Instance
}

func (s *ListAbnormalyEventsRequest) GetLevel() *string {
	return s.Level
}

func (s *ListAbnormalyEventsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAbnormalyEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAbnormalyEventsRequest) GetPod() *string {
	return s.Pod
}

func (s *ListAbnormalyEventsRequest) GetShowPod() *int32 {
	return s.ShowPod
}

func (s *ListAbnormalyEventsRequest) GetStart() *float32 {
	return s.Start
}

func (s *ListAbnormalyEventsRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListAbnormalyEventsRequest) SetXDebugId(v string) *ListAbnormalyEventsRequest {
	s.XDebugId = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetCluster(v string) *ListAbnormalyEventsRequest {
	s.Cluster = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetCurrent(v int32) *ListAbnormalyEventsRequest {
	s.Current = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetEnd(v float32) *ListAbnormalyEventsRequest {
	s.End = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetEvent(v string) *ListAbnormalyEventsRequest {
	s.Event = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetInstance(v string) *ListAbnormalyEventsRequest {
	s.Instance = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetLevel(v string) *ListAbnormalyEventsRequest {
	s.Level = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetNamespace(v string) *ListAbnormalyEventsRequest {
	s.Namespace = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetPageSize(v int32) *ListAbnormalyEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetPod(v string) *ListAbnormalyEventsRequest {
	s.Pod = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetShowPod(v int32) *ListAbnormalyEventsRequest {
	s.ShowPod = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetStart(v float32) *ListAbnormalyEventsRequest {
	s.Start = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetXSysomInvokeSource(v string) *ListAbnormalyEventsRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListAbnormalyEventsRequest) Validate() error {
	return dara.Validate(s)
}
