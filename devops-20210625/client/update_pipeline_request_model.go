// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdatePipelineRequest
	GetContent() *string
	SetName(v string) *UpdatePipelineRequest
	GetName() *string
	SetPipelineId(v string) *UpdatePipelineRequest
	GetPipelineId() *string
}

type UpdatePipelineRequest struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111xxx
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
}

func (s UpdatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequest) GetContent() *string {
	return s.Content
}

func (s *UpdatePipelineRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *UpdatePipelineRequest) SetContent(v string) *UpdatePipelineRequest {
	s.Content = &v
	return s
}

func (s *UpdatePipelineRequest) SetName(v string) *UpdatePipelineRequest {
	s.Name = &v
	return s
}

func (s *UpdatePipelineRequest) SetPipelineId(v string) *UpdatePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *UpdatePipelineRequest) Validate() error {
	return dara.Validate(s)
}
