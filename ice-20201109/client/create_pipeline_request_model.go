// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreatePipelineRequest
	GetName() *string
	SetPriority(v int32) *CreatePipelineRequest
	GetPriority() *int32
	SetSpeed(v string) *CreatePipelineRequest
	GetSpeed() *string
}

type CreatePipelineRequest struct {
	// The name of the MPS queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-pipeline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority. Default value: 6. Valid values: 1 to 10. A greater value specifies a higher priority.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the MPS queue. Valid values:
	//
	// 1.  Standard: standard MPS queue.
	//
	// 2.  Boost: MPS queue with transcoding speed boosted.
	//
	// 3.  NarrowBandHDV2: MPS queue that supports Narrowband HD 2.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// Standard
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
}

func (s CreatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequest) GetName() *string {
	return s.Name
}

func (s *CreatePipelineRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreatePipelineRequest) GetSpeed() *string {
	return s.Speed
}

func (s *CreatePipelineRequest) SetName(v string) *CreatePipelineRequest {
	s.Name = &v
	return s
}

func (s *CreatePipelineRequest) SetPriority(v int32) *CreatePipelineRequest {
	s.Priority = &v
	return s
}

func (s *CreatePipelineRequest) SetSpeed(v string) *CreatePipelineRequest {
	s.Speed = &v
	return s
}

func (s *CreatePipelineRequest) Validate() error {
	return dara.Validate(s)
}
