// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineId(v string) *GetPipelineRequest
	GetPipelineId() *string
}

type GetPipelineRequest struct {
	// The ID of the MPS queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s GetPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetPipelineRequest) SetPipelineId(v string) *GetPipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineRequest) Validate() error {
	return dara.Validate(s)
}
