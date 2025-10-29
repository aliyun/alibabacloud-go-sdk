// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElastictaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetElastictaskResponseBody
	GetRequestId() *string
	SetResult(v *GetElastictaskResponseBodyResult) *GetElastictaskResponseBody
	GetResult() *GetElastictaskResponseBodyResult
}

type GetElastictaskResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetElastictaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetElastictaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetElastictaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetElastictaskResponseBody) GetResult() *GetElastictaskResponseBodyResult {
	return s.Result
}

func (s *GetElastictaskResponseBody) SetRequestId(v string) *GetElastictaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetElastictaskResponseBody) SetResult(v *GetElastictaskResponseBodyResult) *GetElastictaskResponseBody {
	s.Result = v
	return s
}

func (s *GetElastictaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetElastictaskResponseBodyResult struct {
	ElasticExpansionTask *GetElastictaskResponseBodyResultElasticExpansionTask `json:"elasticExpansionTask,omitempty" xml:"elasticExpansionTask,omitempty" type:"Struct"`
	ElasticShrinkTask    *GetElastictaskResponseBodyResultElasticShrinkTask    `json:"elasticShrinkTask,omitempty" xml:"elasticShrinkTask,omitempty" type:"Struct"`
}

func (s GetElastictaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetElastictaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponseBodyResult) GetElasticExpansionTask() *GetElastictaskResponseBodyResultElasticExpansionTask {
	return s.ElasticExpansionTask
}

func (s *GetElastictaskResponseBodyResult) GetElasticShrinkTask() *GetElastictaskResponseBodyResultElasticShrinkTask {
	return s.ElasticShrinkTask
}

func (s *GetElastictaskResponseBodyResult) SetElasticExpansionTask(v *GetElastictaskResponseBodyResultElasticExpansionTask) *GetElastictaskResponseBodyResult {
	s.ElasticExpansionTask = v
	return s
}

func (s *GetElastictaskResponseBodyResult) SetElasticShrinkTask(v *GetElastictaskResponseBodyResultElasticShrinkTask) *GetElastictaskResponseBodyResult {
	s.ElasticShrinkTask = v
	return s
}

func (s *GetElastictaskResponseBodyResult) Validate() error {
	if s.ElasticExpansionTask != nil {
		if err := s.ElasticExpansionTask.Validate(); err != nil {
			return err
		}
	}
	if s.ElasticShrinkTask != nil {
		if err := s.ElasticShrinkTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetElastictaskResponseBodyResultElasticExpansionTask struct {
	// example:
	//
	// 0 0 0 ? 	- MON
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// example:
	//
	// 2
	ElasticNodeCount *int32 `json:"elasticNodeCount,omitempty" xml:"elasticNodeCount,omitempty"`
	// example:
	//
	// 2
	ReplicaCount  *int32    `json:"replicaCount,omitempty" xml:"replicaCount,omitempty"`
	TargetIndices []*string `json:"targetIndices,omitempty" xml:"targetIndices,omitempty" type:"Repeated"`
	// example:
	//
	// crontab
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetElastictaskResponseBodyResultElasticExpansionTask) String() string {
	return dara.Prettify(s)
}

func (s GetElastictaskResponseBodyResultElasticExpansionTask) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) GetElasticNodeCount() *int32 {
	return s.ElasticNodeCount
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) GetTargetIndices() []*string {
	return s.TargetIndices
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetCronExpression(v string) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.CronExpression = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetElasticNodeCount(v int32) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.ElasticNodeCount = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetReplicaCount(v int32) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.ReplicaCount = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetTargetIndices(v []*string) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.TargetIndices = v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetTriggerType(v string) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.TriggerType = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) Validate() error {
	return dara.Validate(s)
}

type GetElastictaskResponseBodyResultElasticShrinkTask struct {
	// example:
	//
	// 4 4 4 ? 	- WED
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// example:
	//
	// 2
	ElasticNodeCount *int32 `json:"elasticNodeCount,omitempty" xml:"elasticNodeCount,omitempty"`
	// example:
	//
	// 2
	ReplicaCount  *int32    `json:"replicaCount,omitempty" xml:"replicaCount,omitempty"`
	TargetIndices []*string `json:"targetIndices,omitempty" xml:"targetIndices,omitempty" type:"Repeated"`
	// example:
	//
	// crontab
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetElastictaskResponseBodyResultElasticShrinkTask) String() string {
	return dara.Prettify(s)
}

func (s GetElastictaskResponseBodyResultElasticShrinkTask) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) GetElasticNodeCount() *int32 {
	return s.ElasticNodeCount
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) GetTargetIndices() []*string {
	return s.TargetIndices
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetCronExpression(v string) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.CronExpression = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetElasticNodeCount(v int32) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.ElasticNodeCount = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetReplicaCount(v int32) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.ReplicaCount = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetTargetIndices(v []*string) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.TargetIndices = v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetTriggerType(v string) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.TriggerType = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) Validate() error {
	return dara.Validate(s)
}
