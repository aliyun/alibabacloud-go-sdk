// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePipelineResponseBody
	GetRequestId() *string
	SetResult(v *DescribePipelineResponseBodyResult) *DescribePipelineResponseBody
	GetResult() *DescribePipelineResponseBodyResult
}

type DescribePipelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 829F38F6-E2D6-4109-90A6-888160BD1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned pipeline information. For more information, see [logstash.yml](https://www.elastic.co/guide/en/logstash/6.7/logstash-settings-file.html).
	Result *DescribePipelineResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePipelineResponseBody) GetResult() *DescribePipelineResponseBodyResult {
	return s.Result
}

func (s *DescribePipelineResponseBody) SetRequestId(v string) *DescribePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePipelineResponseBody) SetResult(v *DescribePipelineResponseBodyResult) *DescribePipelineResponseBody {
	s.Result = v
	return s
}

func (s *DescribePipelineResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePipelineResponseBodyResult struct {
	// The batch delay of the pipeline. Unit: milliseconds.
	//
	// example:
	//
	// 50
	BatchDelay *int32 `json:"batchDelay,omitempty" xml:"batchDelay,omitempty"`
	// The batch size of the pipeline.
	//
	// example:
	//
	// 125
	BatchSize *int32 `json:"batchSize,omitempty" xml:"batchSize,omitempty"`
	// The specific configuration of the pipeline.
	//
	// example:
	//
	// input {  }  filter {  }  output {  }
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The pipeline description.
	//
	// example:
	//
	// this is a test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The time when the pipeline was created.
	//
	// example:
	//
	// 2020-06-20T07:26:47.000+0000
	GmtCreatedTime *string `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	// The time when the pipeline was last updated.
	//
	// example:
	//
	// 2020-06-20T07:26:47.000+0000
	GmtUpdateTime *string `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	// The pipeline ID.
	//
	// example:
	//
	// pipeline_test
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// The pipeline status. Valid values:
	//
	// - NOT_DEPLOYED: not deployed.
	//
	// - RUNNING: running.
	//
	// - DELETED: deleted. This status is not displayed in the console.
	//
	// example:
	//
	// RUNNING
	PipelineStatus *string `json:"pipelineStatus,omitempty" xml:"pipelineStatus,omitempty"`
	// The number of queue checkpoint writes.
	//
	// example:
	//
	// 1024
	QueueCheckPointWrites *int32 `json:"queueCheckPointWrites,omitempty" xml:"queueCheckPointWrites,omitempty"`
	// The total capacity of the queue, in bytes. Unit: MB.
	//
	// example:
	//
	// 1024
	QueueMaxBytes *int32 `json:"queueMaxBytes,omitempty" xml:"queueMaxBytes,omitempty"`
	// The queue type. Valid values:
	//
	// - MEMORY: a traditional memory-based queue.
	//
	// - PERSISTED: a disk-based ACKed queue (persistent queue).
	//
	// example:
	//
	// MEMORY
	QueueType *string `json:"queueType,omitempty" xml:"queueType,omitempty"`
	// The number of pipeline worker threads.
	//
	// example:
	//
	// 2
	Workers *int32 `json:"workers,omitempty" xml:"workers,omitempty"`
}

func (s DescribePipelineResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBodyResult) GetBatchDelay() *int32 {
	return s.BatchDelay
}

func (s *DescribePipelineResponseBodyResult) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *DescribePipelineResponseBodyResult) GetConfig() *string {
	return s.Config
}

func (s *DescribePipelineResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribePipelineResponseBodyResult) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *DescribePipelineResponseBodyResult) GetGmtUpdateTime() *string {
	return s.GmtUpdateTime
}

func (s *DescribePipelineResponseBodyResult) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DescribePipelineResponseBodyResult) GetPipelineStatus() *string {
	return s.PipelineStatus
}

func (s *DescribePipelineResponseBodyResult) GetQueueCheckPointWrites() *int32 {
	return s.QueueCheckPointWrites
}

func (s *DescribePipelineResponseBodyResult) GetQueueMaxBytes() *int32 {
	return s.QueueMaxBytes
}

func (s *DescribePipelineResponseBodyResult) GetQueueType() *string {
	return s.QueueType
}

func (s *DescribePipelineResponseBodyResult) GetWorkers() *int32 {
	return s.Workers
}

func (s *DescribePipelineResponseBodyResult) SetBatchDelay(v int32) *DescribePipelineResponseBodyResult {
	s.BatchDelay = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetBatchSize(v int32) *DescribePipelineResponseBodyResult {
	s.BatchSize = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetConfig(v string) *DescribePipelineResponseBodyResult {
	s.Config = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetDescription(v string) *DescribePipelineResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetGmtCreatedTime(v string) *DescribePipelineResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetGmtUpdateTime(v string) *DescribePipelineResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetPipelineId(v string) *DescribePipelineResponseBodyResult {
	s.PipelineId = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetPipelineStatus(v string) *DescribePipelineResponseBodyResult {
	s.PipelineStatus = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetQueueCheckPointWrites(v int32) *DescribePipelineResponseBodyResult {
	s.QueueCheckPointWrites = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetQueueMaxBytes(v int32) *DescribePipelineResponseBodyResult {
	s.QueueMaxBytes = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetQueueType(v string) *DescribePipelineResponseBodyResult {
	s.QueueType = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetWorkers(v int32) *DescribePipelineResponseBodyResult {
	s.Workers = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
