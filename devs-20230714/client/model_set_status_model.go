// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSetStatus interface {
	dara.Model
	String() string
	GoString() string
	SetObservedGeneration(v int64) *ModelSetStatus
	GetObservedGeneration() *int64
	SetObservedTime(v string) *ModelSetStatus
	GetObservedTime() *string
	SetPhase(v string) *ModelSetStatus
	GetPhase() *string
}

type ModelSetStatus struct {
	ObservedGeneration *int64  `json:"observedGeneration,omitempty" xml:"observedGeneration,omitempty"`
	ObservedTime       *string `json:"observedTime,omitempty" xml:"observedTime,omitempty"`
	Phase              *string `json:"phase,omitempty" xml:"phase,omitempty"`
}

func (s ModelSetStatus) String() string {
	return dara.Prettify(s)
}

func (s ModelSetStatus) GoString() string {
	return s.String()
}

func (s *ModelSetStatus) GetObservedGeneration() *int64 {
	return s.ObservedGeneration
}

func (s *ModelSetStatus) GetObservedTime() *string {
	return s.ObservedTime
}

func (s *ModelSetStatus) GetPhase() *string {
	return s.Phase
}

func (s *ModelSetStatus) SetObservedGeneration(v int64) *ModelSetStatus {
	s.ObservedGeneration = &v
	return s
}

func (s *ModelSetStatus) SetObservedTime(v string) *ModelSetStatus {
	s.ObservedTime = &v
	return s
}

func (s *ModelSetStatus) SetPhase(v string) *ModelSetStatus {
	s.Phase = &v
	return s
}

func (s *ModelSetStatus) Validate() error {
	return dara.Validate(s)
}
