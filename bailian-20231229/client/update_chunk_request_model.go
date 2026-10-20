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
	// The ID of the text chunk to modify. You can obtain this value by calling the **ListChunks*	- operation. The value is in the Node.Metadata._id field of the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-5ip55o1zrzxx_09fe52x_xxxxx_033b551e10024029992e79767b151fxx_10024xx_0
	ChunkId *string `json:"ChunkId,omitempty" xml:"ChunkId,omitempty"`
	// The file ID, which is the `FileId` returned by the **AddFile*	- operation. You can also obtain it from the <props="china">[application data](https://bailian.console.aliyun.com/?tab=app#/data-center)-Files<props="intl">[application data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center)-Files tab in the Alibaba Cloud Model Studio console by clicking the ID icon next to the corresponding file.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc_c134aa2073204a5d936d870bf960f56axxxxxxxx
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Specifies whether the text chunk participates in knowledge base retrieval. Valid values:
	//
	// - true: Participates.
	//
	// - false: Does not participate.
	//
	// Default value: true.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDisplayedChunkContent *bool `json:"IsDisplayedChunkContent,omitempty" xml:"IsDisplayedChunkContent,omitempty"`
	// The knowledge base ID, which is the `Data.Id` returned when you call the **CreateIndex*	- operation, or retrieved from the <props="china">[knowledge base](https://bailian.console.aliyun.com/?tab=app#/knowledge-base)<props="intl">[knowledge base](https://modelstudio.console.alibabacloud.com/?tab=app#/knowledge-base) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The new content of the text chunk. The content length must be between 10 and 6,000 characters, and must not exceed the maximum segment length set when the knowledge base was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// The type of certainty obtained in philosophy is not scientific certainty (that is, certainty that is the same for every rational mind), but a kind of personal conviction that can only be obtained in the overall nature of humanity. Every form of philosophy differs from science in that no philosophy has received universal acceptance..
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The new title of the text chunk. The length limit is 0–50 characters. An empty string is allowed. If you pass an empty string, the existing title is cleared. If you do not pass this parameter, the original title is retained.
	//
	// example:
	//
	// What is philosophy
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
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
