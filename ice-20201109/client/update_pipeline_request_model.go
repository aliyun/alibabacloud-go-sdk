// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdatePipelineRequest
	GetName() *string
	SetPipelineId(v string) *UpdatePipelineRequest
	GetPipelineId() *string
	SetPriority(v int32) *UpdatePipelineRequest
	GetPriority() *int32
	SetStatus(v string) *UpdatePipelineRequest
	GetStatus() *string
}

type UpdatePipelineRequest struct {
	// The name of the MPS queue.
	//
	// example:
	//
	// test-pipeline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the MPS queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the MPS queue. Valid values: 1 to 10.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
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
	// Paused
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *UpdatePipelineRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdatePipelineRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdatePipelineRequest) SetName(v string) *UpdatePipelineRequest {
	s.Name = &v
	return s
}

func (s *UpdatePipelineRequest) SetPipelineId(v string) *UpdatePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *UpdatePipelineRequest) SetPriority(v int32) *UpdatePipelineRequest {
	s.Priority = &v
	return s
}

func (s *UpdatePipelineRequest) SetStatus(v string) *UpdatePipelineRequest {
	s.Status = &v
	return s
}

func (s *UpdatePipelineRequest) Validate() error {
	return dara.Validate(s)
}
