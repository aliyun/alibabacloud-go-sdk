// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChunkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineId(v string) *AddChunkRequest
	GetPipelineId() *string
	SetDataId(v string) *AddChunkRequest
	GetDataId() *string
	SetField(v map[string]interface{}) *AddChunkRequest
	GetField() map[string]interface{}
}

type AddChunkRequest struct {
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
	Field map[string]interface{} `json:"field,omitempty" xml:"field,omitempty"`
}

func (s AddChunkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddChunkRequest) GoString() string {
	return s.String()
}

func (s *AddChunkRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *AddChunkRequest) GetDataId() *string {
	return s.DataId
}

func (s *AddChunkRequest) GetField() map[string]interface{} {
	return s.Field
}

func (s *AddChunkRequest) SetPipelineId(v string) *AddChunkRequest {
	s.PipelineId = &v
	return s
}

func (s *AddChunkRequest) SetDataId(v string) *AddChunkRequest {
	s.DataId = &v
	return s
}

func (s *AddChunkRequest) SetField(v map[string]interface{}) *AddChunkRequest {
	s.Field = v
	return s
}

func (s *AddChunkRequest) Validate() error {
	return dara.Validate(s)
}
