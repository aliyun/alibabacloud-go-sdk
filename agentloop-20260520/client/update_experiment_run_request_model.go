// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompletedAt(v int64) *UpdateExperimentRunRequest
	GetCompletedAt() *int64
	SetCompletedTasks(v int32) *UpdateExperimentRunRequest
	GetCompletedTasks() *int32
	SetExecutedAt(v int64) *UpdateExperimentRunRequest
	GetExecutedAt() *int64
	SetFailedTasks(v int32) *UpdateExperimentRunRequest
	GetFailedTasks() *int32
	SetRecordName(v string) *UpdateExperimentRunRequest
	GetRecordName() *string
	SetStatus(v string) *UpdateExperimentRunRequest
	GetStatus() *string
	SetTotalTasks(v int32) *UpdateExperimentRunRequest
	GetTotalTasks() *int32
	SetClientToken(v string) *UpdateExperimentRunRequest
	GetClientToken() *string
}

type UpdateExperimentRunRequest struct {
	// The experiment completion time. A millisecond-level UNIX timestamp.
	//
	// example:
	//
	// 1784719989371
	CompletedAt *int64 `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// The number of completed tasks.
	//
	// example:
	//
	// 10
	CompletedTasks *int32 `json:"completedTasks,omitempty" xml:"completedTasks,omitempty"`
	// The experiment execution time. A millisecond-level UNIX timestamp.
	//
	// example:
	//
	// 1784719439255
	ExecutedAt *int64 `json:"executedAt,omitempty" xml:"executedAt,omitempty"`
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedTasks *int32 `json:"failedTasks,omitempty" xml:"failedTasks,omitempty"`
	// The experiment record name.
	//
	// example:
	//
	// rca_benchmark_eval_experiment 2026/07/22 19:23:59
	RecordName *string `json:"recordName,omitempty" xml:"recordName,omitempty"`
	// The experiment record status. Set to cancelled to cancel execution.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 20
	TotalTasks *int32 `json:"totalTasks,omitempty" xml:"totalTasks,omitempty"`
	// Optional.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateExperimentRunRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentRunRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRunRequest) GetCompletedAt() *int64 {
	return s.CompletedAt
}

func (s *UpdateExperimentRunRequest) GetCompletedTasks() *int32 {
	return s.CompletedTasks
}

func (s *UpdateExperimentRunRequest) GetExecutedAt() *int64 {
	return s.ExecutedAt
}

func (s *UpdateExperimentRunRequest) GetFailedTasks() *int32 {
	return s.FailedTasks
}

func (s *UpdateExperimentRunRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *UpdateExperimentRunRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateExperimentRunRequest) GetTotalTasks() *int32 {
	return s.TotalTasks
}

func (s *UpdateExperimentRunRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateExperimentRunRequest) SetCompletedAt(v int64) *UpdateExperimentRunRequest {
	s.CompletedAt = &v
	return s
}

func (s *UpdateExperimentRunRequest) SetCompletedTasks(v int32) *UpdateExperimentRunRequest {
	s.CompletedTasks = &v
	return s
}

func (s *UpdateExperimentRunRequest) SetExecutedAt(v int64) *UpdateExperimentRunRequest {
	s.ExecutedAt = &v
	return s
}

func (s *UpdateExperimentRunRequest) SetFailedTasks(v int32) *UpdateExperimentRunRequest {
	s.FailedTasks = &v
	return s
}

func (s *UpdateExperimentRunRequest) SetRecordName(v string) *UpdateExperimentRunRequest {
	s.RecordName = &v
	return s
}

func (s *UpdateExperimentRunRequest) SetStatus(v string) *UpdateExperimentRunRequest {
	s.Status = &v
	return s
}

func (s *UpdateExperimentRunRequest) SetTotalTasks(v int32) *UpdateExperimentRunRequest {
	s.TotalTasks = &v
	return s
}

func (s *UpdateExperimentRunRequest) SetClientToken(v string) *UpdateExperimentRunRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateExperimentRunRequest) Validate() error {
	return dara.Validate(s)
}
