// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobStatusRunning interface {
	dara.Model
	String() string
	GoString() string
	SetObservedFlinkJobRestarts(v int64) *JobStatusRunning
	GetObservedFlinkJobRestarts() *int64
	SetObservedFlinkJobStatus(v string) *JobStatusRunning
	GetObservedFlinkJobStatus() *string
}

type JobStatusRunning struct {
	// example:
	//
	// 4
	ObservedFlinkJobRestarts *int64 `json:"observedFlinkJobRestarts,omitempty" xml:"observedFlinkJobRestarts,omitempty"`
	// example:
	//
	// RUNNING
	ObservedFlinkJobStatus *string `json:"observedFlinkJobStatus,omitempty" xml:"observedFlinkJobStatus,omitempty"`
}

func (s JobStatusRunning) String() string {
	return dara.Prettify(s)
}

func (s JobStatusRunning) GoString() string {
	return s.String()
}

func (s *JobStatusRunning) GetObservedFlinkJobRestarts() *int64 {
	return s.ObservedFlinkJobRestarts
}

func (s *JobStatusRunning) GetObservedFlinkJobStatus() *string {
	return s.ObservedFlinkJobStatus
}

func (s *JobStatusRunning) SetObservedFlinkJobRestarts(v int64) *JobStatusRunning {
	s.ObservedFlinkJobRestarts = &v
	return s
}

func (s *JobStatusRunning) SetObservedFlinkJobStatus(v string) *JobStatusRunning {
	s.ObservedFlinkJobStatus = &v
	return s
}

func (s *JobStatusRunning) Validate() error {
	return dara.Validate(s)
}
