// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreatePipelineRequest
	GetContent() *string
	SetName(v string) *CreatePipelineRequest
	GetName() *string
}

type CreatePipelineRequest struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequest) GetContent() *string {
	return s.Content
}

func (s *CreatePipelineRequest) GetName() *string {
	return s.Name
}

func (s *CreatePipelineRequest) SetContent(v string) *CreatePipelineRequest {
	s.Content = &v
	return s
}

func (s *CreatePipelineRequest) SetName(v string) *CreatePipelineRequest {
	s.Name = &v
	return s
}

func (s *CreatePipelineRequest) Validate() error {
	return dara.Validate(s)
}
