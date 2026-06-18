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
	// The file ID. This is the `FileId` returned by the **AddFile*	- operation. You can also obtain it from the <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) - Files<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) - Files tab in the Model Studio console by clicking the ID icon next to the corresponding file.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc_c134aa2073204a5d936d870bf960f56axxxxxxxx
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Specifies whether this text chunk participates in knowledge base retrieval. Valid values:
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
	// The knowledge base ID. This is the `Data.Id` returned by the **CreateIndex*	- operation, or you can obtain it from the <props="china">[Knowledge Base](https://bailian.console.aliyun.com/?tab=app#/knowledge-base)<props="intl">[Knowledge Base](https://modelstudio.console.alibabacloud.com/?tab=app#/knowledge-base) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The new content of the text chunk. The content length must be between 10 and 6000 characters and cannot exceed the maximum segment length specified when the knowledge base was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 在哲学中所获得的确定性类型不是科学的确定性(即对每个人的理智来说都一样的确定性)，而是一种要在人类的整体本质中才能获得的亲证。哲学的每一形态都不同于科学，因为所有的哲学都没有得到一致的认可...
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The new title of the text chunk. The length is limited to 0 to 50 characters. An empty string is allowed. If you pass an empty string, the existing title is cleared. If you do not pass this parameter, the original title is retained.
	//
	// example:
	//
	// 什么是哲学
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
