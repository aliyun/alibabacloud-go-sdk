// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAbnormalEventsCountRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type GetAbnormalEventsCountRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	Level    *string `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// test-pod
	Pod     *string `json:"pod,omitempty" xml:"pod,omitempty"`
	ShowPod *int32  `json:"showPod,omitempty" xml:"showPod,omitempty"`
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetAbnormalEventsCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAbnormalEventsCountRequest) GoString() string {
	return s.String()
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

func (s *GetAbnormalEventsCountRequest) Validate() error {
	return dara.Validate(s)
}
