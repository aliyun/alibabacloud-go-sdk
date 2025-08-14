// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineId(v string) *DeletePipelineRequest
	GetPipelineId() *string
}

type DeletePipelineRequest struct {
	// The ID of the MPS queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s DeletePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineRequest) GoString() string {
	return s.String()
}

func (s *DeletePipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DeletePipelineRequest) SetPipelineId(v string) *DeletePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *DeletePipelineRequest) Validate() error {
	return dara.Validate(s)
}
