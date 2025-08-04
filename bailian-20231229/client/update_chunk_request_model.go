// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChunkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkId(v string) *UpdateChunkRequest
	GetChunkId() *string
	SetDataId(v string) *UpdateChunkRequest
	GetDataId() *string
	SetIsDisplayedChunkContent(v bool) *UpdateChunkRequest
	GetIsDisplayedChunkContent() *bool
	SetPipelineId(v string) *UpdateChunkRequest
	GetPipelineId() *string
	SetContent(v string) *UpdateChunkRequest
	GetContent() *string
	SetTitle(v string) *UpdateChunkRequest
	GetTitle() *string
}

type UpdateChunkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// llm-5ip55o1zrzxx_09fe52x_table_033b551e10024029992e79767b151fxx_10024xx_0
	ChunkId *string `json:"ChunkId,omitempty" xml:"ChunkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// doc_c134aa2073204a5d936d870bf960f56axxxxxxxx
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDisplayedChunkContent *bool `json:"IsDisplayedChunkContent,omitempty" xml:"IsDisplayedChunkContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateChunkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChunkRequest) GoString() string {
	return s.String()
}

func (s *UpdateChunkRequest) GetChunkId() *string {
	return s.ChunkId
}

func (s *UpdateChunkRequest) GetDataId() *string {
	return s.DataId
}

func (s *UpdateChunkRequest) GetIsDisplayedChunkContent() *bool {
	return s.IsDisplayedChunkContent
}

func (s *UpdateChunkRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *UpdateChunkRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateChunkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateChunkRequest) SetChunkId(v string) *UpdateChunkRequest {
	s.ChunkId = &v
	return s
}

func (s *UpdateChunkRequest) SetDataId(v string) *UpdateChunkRequest {
	s.DataId = &v
	return s
}

func (s *UpdateChunkRequest) SetIsDisplayedChunkContent(v bool) *UpdateChunkRequest {
	s.IsDisplayedChunkContent = &v
	return s
}

func (s *UpdateChunkRequest) SetPipelineId(v string) *UpdateChunkRequest {
	s.PipelineId = &v
	return s
}

func (s *UpdateChunkRequest) SetContent(v string) *UpdateChunkRequest {
	s.Content = &v
	return s
}

func (s *UpdateChunkRequest) SetTitle(v string) *UpdateChunkRequest {
	s.Title = &v
	return s
}

func (s *UpdateChunkRequest) Validate() error {
	return dara.Validate(s)
}
