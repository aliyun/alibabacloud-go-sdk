// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Pipeline) *CreatePipelineRequest
	GetBody() *Pipeline
}

type CreatePipelineRequest struct {
	Body *Pipeline `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequest) GetBody() *Pipeline {
	return s.Body
}

func (s *CreatePipelineRequest) SetBody(v *Pipeline) *CreatePipelineRequest {
	s.Body = v
	return s
}

func (s *CreatePipelineRequest) Validate() error {
	return dara.Validate(s)
}
