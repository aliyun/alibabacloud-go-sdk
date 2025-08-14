// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeline(v *GetPipelineResponseBodyPipeline) *GetPipelineResponseBody
	GetPipeline() *GetPipelineResponseBodyPipeline
	SetRequestId(v string) *GetPipelineResponseBody
	GetRequestId() *string
}

type GetPipelineResponseBody struct {
	// The information about the MPS queue.
	Pipeline *GetPipelineResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBody) GetPipeline() *GetPipelineResponseBodyPipeline {
	return s.Pipeline
}

func (s *GetPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineResponseBody) SetPipeline(v *GetPipelineResponseBodyPipeline) *GetPipelineResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetPipelineResponseBody) SetRequestId(v string) *GetPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPipelineResponseBodyPipeline struct {
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
	// The priority of the MPS queue. Valid values: 1 to 10.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the MPS queue.
	//
	// Valid values:
	//
	// 	- Standard: standard MPS queue.
	//
	// 	- Boost: MPS queue with transcoding speed boosted.
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

func (s GetPipelineResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBodyPipeline) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPipelineResponseBodyPipeline) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetPipelineResponseBodyPipeline) GetName() *string {
	return s.Name
}

func (s *GetPipelineResponseBodyPipeline) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetPipelineResponseBodyPipeline) GetPriority() *int32 {
	return s.Priority
}

func (s *GetPipelineResponseBodyPipeline) GetSpeed() *string {
	return s.Speed
}

func (s *GetPipelineResponseBodyPipeline) GetStatus() *string {
	return s.Status
}

func (s *GetPipelineResponseBodyPipeline) SetCreateTime(v string) *GetPipelineResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetModifiedTime(v string) *GetPipelineResponseBodyPipeline {
	s.ModifiedTime = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetName(v string) *GetPipelineResponseBodyPipeline {
	s.Name = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetPipelineId(v string) *GetPipelineResponseBodyPipeline {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetPriority(v int32) *GetPipelineResponseBodyPipeline {
	s.Priority = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetSpeed(v string) *GetPipelineResponseBodyPipeline {
	s.Speed = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) SetStatus(v string) *GetPipelineResponseBodyPipeline {
	s.Status = &v
	return s
}

func (s *GetPipelineResponseBodyPipeline) Validate() error {
	return dara.Validate(s)
}
