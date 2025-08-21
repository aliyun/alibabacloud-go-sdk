// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElastictaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyElastictaskResponseBody
	GetRequestId() *string
	SetResult(v *ModifyElastictaskResponseBodyResult) *ModifyElastictaskResponseBody
	GetResult() *ModifyElastictaskResponseBodyResult
}

type ModifyElastictaskResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ModifyElastictaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyElastictaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElastictaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElastictaskResponseBody) GetResult() *ModifyElastictaskResponseBodyResult {
	return s.Result
}

func (s *ModifyElastictaskResponseBody) SetRequestId(v string) *ModifyElastictaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElastictaskResponseBody) SetResult(v *ModifyElastictaskResponseBodyResult) *ModifyElastictaskResponseBody {
	s.Result = v
	return s
}

func (s *ModifyElastictaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyElastictaskResponseBodyResult struct {
	ElasticExpansionTask *ModifyElastictaskResponseBodyResultElasticExpansionTask `json:"elasticExpansionTask,omitempty" xml:"elasticExpansionTask,omitempty" type:"Struct"`
	ElasticShrinkTask    *ModifyElastictaskResponseBodyResultElasticShrinkTask    `json:"elasticShrinkTask,omitempty" xml:"elasticShrinkTask,omitempty" type:"Struct"`
}

func (s ModifyElastictaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifyElastictaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponseBodyResult) GetElasticExpansionTask() *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	return s.ElasticExpansionTask
}

func (s *ModifyElastictaskResponseBodyResult) GetElasticShrinkTask() *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	return s.ElasticShrinkTask
}

func (s *ModifyElastictaskResponseBodyResult) SetElasticExpansionTask(v *ModifyElastictaskResponseBodyResultElasticExpansionTask) *ModifyElastictaskResponseBodyResult {
	s.ElasticExpansionTask = v
	return s
}

func (s *ModifyElastictaskResponseBodyResult) SetElasticShrinkTask(v *ModifyElastictaskResponseBodyResultElasticShrinkTask) *ModifyElastictaskResponseBodyResult {
	s.ElasticShrinkTask = v
	return s
}

func (s *ModifyElastictaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ModifyElastictaskResponseBodyResultElasticExpansionTask struct {
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

func (s ModifyElastictaskResponseBodyResultElasticExpansionTask) String() string {
	return dara.Prettify(s)
}

func (s ModifyElastictaskResponseBodyResultElasticExpansionTask) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) GetElasticNodeCount() *int32 {
	return s.ElasticNodeCount
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) GetTargetIndices() []*string {
	return s.TargetIndices
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetCronExpression(v string) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.CronExpression = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetElasticNodeCount(v int32) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.ElasticNodeCount = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetReplicaCount(v int32) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.ReplicaCount = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetTargetIndices(v []*string) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.TargetIndices = v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetTriggerType(v string) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.TriggerType = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) Validate() error {
	return dara.Validate(s)
}

type ModifyElastictaskResponseBodyResultElasticShrinkTask struct {
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

func (s ModifyElastictaskResponseBodyResultElasticShrinkTask) String() string {
	return dara.Prettify(s)
}

func (s ModifyElastictaskResponseBodyResultElasticShrinkTask) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) GetElasticNodeCount() *int32 {
	return s.ElasticNodeCount
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) GetTargetIndices() []*string {
	return s.TargetIndices
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetCronExpression(v string) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.CronExpression = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetElasticNodeCount(v int32) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.ElasticNodeCount = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetReplicaCount(v int32) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.ReplicaCount = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetTargetIndices(v []*string) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.TargetIndices = v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetTriggerType(v string) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.TriggerType = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) Validate() error {
	return dara.Validate(s)
}
