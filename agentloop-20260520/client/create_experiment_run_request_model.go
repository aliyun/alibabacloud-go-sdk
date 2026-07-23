// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompletedAt(v int64) *CreateExperimentRunRequest
	GetCompletedAt() *int64
	SetCompletedTasks(v int32) *CreateExperimentRunRequest
	GetCompletedTasks() *int32
	SetExecutedAt(v int64) *CreateExperimentRunRequest
	GetExecutedAt() *int64
	SetExperimentPlanId(v string) *CreateExperimentRunRequest
	GetExperimentPlanId() *string
	SetFailedTasks(v int32) *CreateExperimentRunRequest
	GetFailedTasks() *int32
	SetOfflineExperiments(v []*OfflineExperimentConfig) *CreateExperimentRunRequest
	GetOfflineExperiments() []*OfflineExperimentConfig
	SetRecordName(v string) *CreateExperimentRunRequest
	GetRecordName() *string
	SetStatus(v string) *CreateExperimentRunRequest
	GetStatus() *string
	SetTotalTasks(v int32) *CreateExperimentRunRequest
	GetTotalTasks() *int32
	SetClientToken(v string) *CreateExperimentRunRequest
	GetClientToken() *string
}

type CreateExperimentRunRequest struct {
	// The completion time, in millisecond-level UNIX timestamp.
	//
	// example:
	//
	// 1784721811392
	CompletedAt *int64 `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// The number of completed tasks. If not specified, the default value is 0.
	//
	// example:
	//
	// 0
	CompletedTasks *int32 `json:"completedTasks,omitempty" xml:"completedTasks,omitempty"`
	// The execution time, in millisecond-level UNIX timestamp.
	//
	// example:
	//
	// 1784721775379
	ExecutedAt *int64 `json:"executedAt,omitempty" xml:"executedAt,omitempty"`
	// The experiment plan ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exp-plan-0242d983f5d340fd8479cf2c19eb279e
	ExperimentPlanId *string `json:"experimentPlanId,omitempty" xml:"experimentPlanId,omitempty"`
	// The number of failed tasks. If not specified, the default value is 0.
	//
	// example:
	//
	// 0
	FailedTasks *int32 `json:"failedTasks,omitempty" xml:"failedTasks,omitempty"`
	// The list of offline experiment configurations. Required when the plan type is offline. The number of items ranges from 1 to 5.
	//
	// example:
	//
	// [{"label": "experimentA", "name": "experimentA"}]
	OfflineExperiments []*OfflineExperimentConfig `json:"offlineExperiments,omitempty" xml:"offlineExperiments,omitempty" type:"Repeated"`
	// The experiment record name. If not specified, the default value is the plan name plus a timestamp.
	//
	// example:
	//
	// arms_agent_experiment 2026/07/22 20:02:55
	RecordName *string `json:"recordName,omitempty" xml:"recordName,omitempty"`
	// The initial status. If not specified, the default value is `pending`.
	//
	// example:
	//
	// pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total number of tasks. For online experiments, if not specified, the value is calculated based on the number of generated tasks.
	//
	// example:
	//
	// 40
	TotalTasks *int32 `json:"totalTasks,omitempty" xml:"totalTasks,omitempty"`
	// Optional.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateExperimentRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentRunRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRunRequest) GetCompletedAt() *int64 {
	return s.CompletedAt
}

func (s *CreateExperimentRunRequest) GetCompletedTasks() *int32 {
	return s.CompletedTasks
}

func (s *CreateExperimentRunRequest) GetExecutedAt() *int64 {
	return s.ExecutedAt
}

func (s *CreateExperimentRunRequest) GetExperimentPlanId() *string {
	return s.ExperimentPlanId
}

func (s *CreateExperimentRunRequest) GetFailedTasks() *int32 {
	return s.FailedTasks
}

func (s *CreateExperimentRunRequest) GetOfflineExperiments() []*OfflineExperimentConfig {
	return s.OfflineExperiments
}

func (s *CreateExperimentRunRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateExperimentRunRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateExperimentRunRequest) GetTotalTasks() *int32 {
	return s.TotalTasks
}

func (s *CreateExperimentRunRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExperimentRunRequest) SetCompletedAt(v int64) *CreateExperimentRunRequest {
	s.CompletedAt = &v
	return s
}

func (s *CreateExperimentRunRequest) SetCompletedTasks(v int32) *CreateExperimentRunRequest {
	s.CompletedTasks = &v
	return s
}

func (s *CreateExperimentRunRequest) SetExecutedAt(v int64) *CreateExperimentRunRequest {
	s.ExecutedAt = &v
	return s
}

func (s *CreateExperimentRunRequest) SetExperimentPlanId(v string) *CreateExperimentRunRequest {
	s.ExperimentPlanId = &v
	return s
}

func (s *CreateExperimentRunRequest) SetFailedTasks(v int32) *CreateExperimentRunRequest {
	s.FailedTasks = &v
	return s
}

func (s *CreateExperimentRunRequest) SetOfflineExperiments(v []*OfflineExperimentConfig) *CreateExperimentRunRequest {
	s.OfflineExperiments = v
	return s
}

func (s *CreateExperimentRunRequest) SetRecordName(v string) *CreateExperimentRunRequest {
	s.RecordName = &v
	return s
}

func (s *CreateExperimentRunRequest) SetStatus(v string) *CreateExperimentRunRequest {
	s.Status = &v
	return s
}

func (s *CreateExperimentRunRequest) SetTotalTasks(v int32) *CreateExperimentRunRequest {
	s.TotalTasks = &v
	return s
}

func (s *CreateExperimentRunRequest) SetClientToken(v string) *CreateExperimentRunRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExperimentRunRequest) Validate() error {
	if s.OfflineExperiments != nil {
		for _, item := range s.OfflineExperiments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
