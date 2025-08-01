// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbnormalyEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *ListAbnormalyEventsRequest
	GetCluster() *string
	SetCurrent(v int32) *ListAbnormalyEventsRequest
	GetCurrent() *int32
	SetEnd(v float32) *ListAbnormalyEventsRequest
	GetEnd() *float32
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
}

type ListAbnormalyEventsRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// potential
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
	// example:
	//
	// 1
	ShowPod *int32 `json:"showPod,omitempty" xml:"showPod,omitempty"`
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListAbnormalyEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAbnormalyEventsRequest) GoString() string {
	return s.String()
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

func (s *ListAbnormalyEventsRequest) Validate() error {
	return dara.Validate(s)
}
