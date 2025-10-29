// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreatePipelinesRequest
	GetClientToken() *string
	SetBody(v []*CreatePipelinesRequestBody) *CreatePipelinesRequest
	GetBody() []*CreatePipelinesRequestBody
	SetTrigger(v bool) *CreatePipelinesRequest
	GetTrigger() *bool
}

type CreatePipelinesRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Body        []*CreatePipelinesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Trigger *bool `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s CreatePipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelinesRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelinesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePipelinesRequest) GetBody() []*CreatePipelinesRequestBody {
	return s.Body
}

func (s *CreatePipelinesRequest) GetTrigger() *bool {
	return s.Trigger
}

func (s *CreatePipelinesRequest) SetClientToken(v string) *CreatePipelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePipelinesRequest) SetBody(v []*CreatePipelinesRequestBody) *CreatePipelinesRequest {
	s.Body = v
	return s
}

func (s *CreatePipelinesRequest) SetTrigger(v bool) *CreatePipelinesRequest {
	s.Trigger = &v
	return s
}

func (s *CreatePipelinesRequest) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePipelinesRequestBody struct {
	// example:
	//
	// 50
	BatchDelay *int32 `json:"batchDelay,omitempty" xml:"batchDelay,omitempty"`
	// example:
	//
	// 125
	BatchSize *int32 `json:"batchSize,omitempty" xml:"batchSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// input { } filter { } output { }
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// this is a test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pipeline-test
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// example:
	//
	// 1024
	QueueCheckPointWrites *int32 `json:"queueCheckPointWrites,omitempty" xml:"queueCheckPointWrites,omitempty"`
	// example:
	//
	// 1024
	QueueMaxBytes *int32 `json:"queueMaxBytes,omitempty" xml:"queueMaxBytes,omitempty"`
	// example:
	//
	// MEMORY
	QueueType *string `json:"queueType,omitempty" xml:"queueType,omitempty"`
	// example:
	//
	// 2
	Workers *int32 `json:"workers,omitempty" xml:"workers,omitempty"`
}

func (s CreatePipelinesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelinesRequestBody) GoString() string {
	return s.String()
}

func (s *CreatePipelinesRequestBody) GetBatchDelay() *int32 {
	return s.BatchDelay
}

func (s *CreatePipelinesRequestBody) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *CreatePipelinesRequestBody) GetConfig() *string {
	return s.Config
}

func (s *CreatePipelinesRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreatePipelinesRequestBody) GetPipelineId() *string {
	return s.PipelineId
}

func (s *CreatePipelinesRequestBody) GetQueueCheckPointWrites() *int32 {
	return s.QueueCheckPointWrites
}

func (s *CreatePipelinesRequestBody) GetQueueMaxBytes() *int32 {
	return s.QueueMaxBytes
}

func (s *CreatePipelinesRequestBody) GetQueueType() *string {
	return s.QueueType
}

func (s *CreatePipelinesRequestBody) GetWorkers() *int32 {
	return s.Workers
}

func (s *CreatePipelinesRequestBody) SetBatchDelay(v int32) *CreatePipelinesRequestBody {
	s.BatchDelay = &v
	return s
}

func (s *CreatePipelinesRequestBody) SetBatchSize(v int32) *CreatePipelinesRequestBody {
	s.BatchSize = &v
	return s
}

func (s *CreatePipelinesRequestBody) SetConfig(v string) *CreatePipelinesRequestBody {
	s.Config = &v
	return s
}

func (s *CreatePipelinesRequestBody) SetDescription(v string) *CreatePipelinesRequestBody {
	s.Description = &v
	return s
}

func (s *CreatePipelinesRequestBody) SetPipelineId(v string) *CreatePipelinesRequestBody {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelinesRequestBody) SetQueueCheckPointWrites(v int32) *CreatePipelinesRequestBody {
	s.QueueCheckPointWrites = &v
	return s
}

func (s *CreatePipelinesRequestBody) SetQueueMaxBytes(v int32) *CreatePipelinesRequestBody {
	s.QueueMaxBytes = &v
	return s
}

func (s *CreatePipelinesRequestBody) SetQueueType(v string) *CreatePipelinesRequestBody {
	s.QueueType = &v
	return s
}

func (s *CreatePipelinesRequestBody) SetWorkers(v int32) *CreatePipelinesRequestBody {
	s.Workers = &v
	return s
}

func (s *CreatePipelinesRequestBody) Validate() error {
	return dara.Validate(s)
}
