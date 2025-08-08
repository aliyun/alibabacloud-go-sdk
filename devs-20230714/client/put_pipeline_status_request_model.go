// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutPipelineStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Pipeline) *PutPipelineStatusRequest
	GetBody() *Pipeline
	SetForce(v bool) *PutPipelineStatusRequest
	GetForce() *bool
}

type PutPipelineStatusRequest struct {
	Body *Pipeline `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutPipelineStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s PutPipelineStatusRequest) GoString() string {
	return s.String()
}

func (s *PutPipelineStatusRequest) GetBody() *Pipeline {
	return s.Body
}

func (s *PutPipelineStatusRequest) GetForce() *bool {
	return s.Force
}

func (s *PutPipelineStatusRequest) SetBody(v *Pipeline) *PutPipelineStatusRequest {
	s.Body = v
	return s
}

func (s *PutPipelineStatusRequest) SetForce(v bool) *PutPipelineStatusRequest {
	s.Force = &v
	return s
}

func (s *PutPipelineStatusRequest) Validate() error {
	return dara.Validate(s)
}
