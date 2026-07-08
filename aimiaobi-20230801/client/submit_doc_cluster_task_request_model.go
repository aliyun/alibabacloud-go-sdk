// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocClusterTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SubmitDocClusterTaskRequest
	GetAgentKey() *string
	SetDocuments(v []*SubmitDocClusterTaskRequestDocuments) *SubmitDocClusterTaskRequest
	GetDocuments() []*SubmitDocClusterTaskRequestDocuments
	SetSummaryLength(v int32) *SubmitDocClusterTaskRequest
	GetSummaryLength() *int32
	SetTitleLength(v int32) *SubmitDocClusterTaskRequest
	GetTitleLength() *int32
	SetTopicCount(v int32) *SubmitDocClusterTaskRequest
	GetTopicCount() *int32
}

type SubmitDocClusterTaskRequest struct {
	// The unique ID of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// A list of documents.
	//
	// This parameter is required.
	Documents []*SubmitDocClusterTaskRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The character limit for the generated cluster summary.
	//
	// example:
	//
	// 49
	SummaryLength *int32 `json:"SummaryLength,omitempty" xml:"SummaryLength,omitempty"`
	// The character limit for the generated cluster title.
	//
	// example:
	//
	// 69
	TitleLength *int32 `json:"TitleLength,omitempty" xml:"TitleLength,omitempty"`
	// The maximum number of clusters.
	//
	// example:
	//
	// 15
	TopicCount *int32 `json:"TopicCount,omitempty" xml:"TopicCount,omitempty"`
}

func (s SubmitDocClusterTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SubmitDocClusterTaskRequest) GetDocuments() []*SubmitDocClusterTaskRequestDocuments {
	return s.Documents
}

func (s *SubmitDocClusterTaskRequest) GetSummaryLength() *int32 {
	return s.SummaryLength
}

func (s *SubmitDocClusterTaskRequest) GetTitleLength() *int32 {
	return s.TitleLength
}

func (s *SubmitDocClusterTaskRequest) GetTopicCount() *int32 {
	return s.TopicCount
}

func (s *SubmitDocClusterTaskRequest) SetAgentKey(v string) *SubmitDocClusterTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitDocClusterTaskRequest) SetDocuments(v []*SubmitDocClusterTaskRequestDocuments) *SubmitDocClusterTaskRequest {
	s.Documents = v
	return s
}

func (s *SubmitDocClusterTaskRequest) SetSummaryLength(v int32) *SubmitDocClusterTaskRequest {
	s.SummaryLength = &v
	return s
}

func (s *SubmitDocClusterTaskRequest) SetTitleLength(v int32) *SubmitDocClusterTaskRequest {
	s.TitleLength = &v
	return s
}

func (s *SubmitDocClusterTaskRequest) SetTopicCount(v int32) *SubmitDocClusterTaskRequest {
	s.TopicCount = &v
	return s
}

func (s *SubmitDocClusterTaskRequest) Validate() error {
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitDocClusterTaskRequestDocuments struct {
	// The content of the document.
	//
	// This parameter is required.
	//
	// example:
	//
	// 文档内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The document ID. It is used to identify the document in the returned cluster results. If you do not provide an ID for any document, the array index is used instead. An error is reported if you provide IDs for only some of the documents.
	//
	// example:
	//
	// 文档ID。用于在返回聚类文章时标识文章。如果文章列表中都不传则使用数组索引作为ID。如果部分传则会报错
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// The title of the document.
	//
	// example:
	//
	// 文档标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SubmitDocClusterTaskRequestDocuments) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocClusterTaskRequestDocuments) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskRequestDocuments) GetContent() *string {
	return s.Content
}

func (s *SubmitDocClusterTaskRequestDocuments) GetDocId() *string {
	return s.DocId
}

func (s *SubmitDocClusterTaskRequestDocuments) GetTitle() *string {
	return s.Title
}

func (s *SubmitDocClusterTaskRequestDocuments) SetContent(v string) *SubmitDocClusterTaskRequestDocuments {
	s.Content = &v
	return s
}

func (s *SubmitDocClusterTaskRequestDocuments) SetDocId(v string) *SubmitDocClusterTaskRequestDocuments {
	s.DocId = &v
	return s
}

func (s *SubmitDocClusterTaskRequestDocuments) SetTitle(v string) *SubmitDocClusterTaskRequestDocuments {
	s.Title = &v
	return s
}

func (s *SubmitDocClusterTaskRequestDocuments) Validate() error {
	return dara.Validate(s)
}
