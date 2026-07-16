// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChunkShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineId(v string) *AddChunkShrinkRequest
	GetPipelineId() *string
	SetDataId(v string) *AddChunkShrinkRequest
	GetDataId() *string
	SetFieldShrink(v string) *AddChunkShrinkRequest
	GetFieldShrink() *string
}

type AddChunkShrinkRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The file ID.
	//
	// example:
	//
	// doc_xxx
	DataId *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The chunk content to insert, passed as key-value pairs. For document search knowledge bases, use the following fixed key list:
	//
	// - content (**String**): **Required**. The body content of the chunk.
	//
	// - title (**String**): **Optional**. The title of the chunk.
	//
	// - image_urls (**Array**): **Optional**. Image URLs contained in the chunk. A maximum of 10 images are supported.
	//
	// For data query and image Q&A knowledge bases, the keys are not fixed and are determined by the data source spreadsheet of the knowledge base. The key is the Excel column header, and the value is the corresponding column value.
	//
	// example:
	//
	// {
	//
	//   "content": "The Bailian platform supports parsing multiple document formats including PDF, Word, and PPT.",
	//
	//   "title": "Document Parsing and Chunking",
	//
	//   "image_urls": [
	//
	// "https://example.com/images/chunk-flow.png",
	//
	//   "https://example.com/images/parsing-result.png"
	//
	//   ]
	//
	// }
	FieldShrink *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s AddChunkShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddChunkShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddChunkShrinkRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *AddChunkShrinkRequest) GetDataId() *string {
	return s.DataId
}

func (s *AddChunkShrinkRequest) GetFieldShrink() *string {
	return s.FieldShrink
}

func (s *AddChunkShrinkRequest) SetPipelineId(v string) *AddChunkShrinkRequest {
	s.PipelineId = &v
	return s
}

func (s *AddChunkShrinkRequest) SetDataId(v string) *AddChunkShrinkRequest {
	s.DataId = &v
	return s
}

func (s *AddChunkShrinkRequest) SetFieldShrink(v string) *AddChunkShrinkRequest {
	s.FieldShrink = &v
	return s
}

func (s *AddChunkShrinkRequest) Validate() error {
	return dara.Validate(s)
}
