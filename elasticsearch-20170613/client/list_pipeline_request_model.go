// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListPipelineRequest
	GetPage() *int32
	SetPipelineId(v string) *ListPipelineRequest
	GetPipelineId() *string
	SetSize(v int32) *ListPipelineRequest
	GetSize() *int32
}

type ListPipelineRequest struct {
	// The header of the response.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// pipeline_test
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 15
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListPipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListPipelineRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListPipelineRequest) SetPage(v int32) *ListPipelineRequest {
	s.Page = &v
	return s
}

func (s *ListPipelineRequest) SetPipelineId(v string) *ListPipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineRequest) SetSize(v int32) *ListPipelineRequest {
	s.Size = &v
	return s
}

func (s *ListPipelineRequest) Validate() error {
	return dara.Validate(s)
}
