// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeline(v *CreatePipelineResponseBodyPipeline) *CreatePipelineResponseBody
	GetPipeline() *CreatePipelineResponseBodyPipeline
	SetRequestId(v string) *CreatePipelineResponseBody
	GetRequestId() *string
}

type CreatePipelineResponseBody struct {
	// The information about the MPS queue.
	Pipeline *CreatePipelineResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponseBody) GetPipeline() *CreatePipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *CreatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelineResponseBody) SetPipeline(v *CreatePipelineResponseBodyPipeline) *CreatePipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *CreatePipelineResponseBody) SetRequestId(v string) *CreatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineResponseBody) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineResponseBodyPipeline struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2022-07-12T16:17:54Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the template was last modified.
	//
	// example:
	//
	// 2022-07-12T16:17:54Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the MPS queue.
	//
	// example:
	//
	// test-pipeline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the MPS queue.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the MPS queue.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the MPS queue.
	//
	// Valid values:
	//
	// 	- Boost: MPS queue with transcoding speed boosted.
	//
	// 	- Standard: standard MPS queue.
	//
	// 	- NarrowBandHDV2: MPS queue that supports Narrowband HD 2.0.
	//
	// example:
	//
	// Standard
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The state of the MPS queue.
	//
	// Valid values:
	//
	// 	- Active
	//
	// 	- Paused
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreatePipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponseBodyPipeline) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreatePipelineResponseBodyPipeline) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *CreatePipelineResponseBodyPipeline) GetName() *string {
	return s.Name
}

func (s *CreatePipelineResponseBodyPipeline) GetPipelineId() *string {
	return s.PipelineId
}

func (s *CreatePipelineResponseBodyPipeline) GetPriority() *int32 {
	return s.Priority
}

func (s *CreatePipelineResponseBodyPipeline) GetSpeed() *string {
	return s.Speed
}

func (s *CreatePipelineResponseBodyPipeline) GetStatus() *string {
	return s.Status
}

func (s *CreatePipelineResponseBodyPipeline) SetCreateTime(v string) *CreatePipelineResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *CreatePipelineResponseBodyPipeline) SetModifiedTime(v string) *CreatePipelineResponseBodyPipeline {
	s.ModifiedTime = &v
	return s
}

func (s *CreatePipelineResponseBodyPipeline) SetName(v string) *CreatePipelineResponseBodyPipeline {
	s.Name = &v
	return s
}

func (s *CreatePipelineResponseBodyPipeline) SetPipelineId(v string) *CreatePipelineResponseBodyPipeline {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelineResponseBodyPipeline) SetPriority(v int32) *CreatePipelineResponseBodyPipeline {
	s.Priority = &v
	return s
}

func (s *CreatePipelineResponseBodyPipeline) SetSpeed(v string) *CreatePipelineResponseBodyPipeline {
	s.Speed = &v
	return s
}

func (s *CreatePipelineResponseBodyPipeline) SetStatus(v string) *CreatePipelineResponseBodyPipeline {
	s.Status = &v
	return s
}

func (s *CreatePipelineResponseBodyPipeline) Validate() error {
	return dara.Validate(s)
}
