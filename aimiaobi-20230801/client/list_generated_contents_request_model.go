// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGeneratedContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListGeneratedContentsRequest
	GetAgentKey() *string
	SetContentDomain(v string) *ListGeneratedContentsRequest
	GetContentDomain() *string
	SetCurrent(v int32) *ListGeneratedContentsRequest
	GetCurrent() *int32
	SetDataType(v string) *ListGeneratedContentsRequest
	GetDataType() *string
	SetEndTime(v string) *ListGeneratedContentsRequest
	GetEndTime() *string
	SetQuery(v string) *ListGeneratedContentsRequest
	GetQuery() *string
	SetSize(v int32) *ListGeneratedContentsRequest
	GetSize() *int32
	SetStartTime(v string) *ListGeneratedContentsRequest
	GetStartTime() *string
	SetTaskId(v string) *ListGeneratedContentsRequest
	GetTaskId() *string
	SetTitle(v string) *ListGeneratedContentsRequest
	GetTitle() *string
}

type ListGeneratedContentsRequest struct {
	// Workspace ID: [AgentKey](https://help.aliyun.com/document_detail/2587494.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Content domain (content category)
	//
	// - media: Media writing
	//
	// - government: Government document writing
	//
	// - office: Office writing
	//
	// - market: Marketing writing
	//
	// - custom: Custom writing
	//
	// - commentGenerate: Opinion generation
	//
	// example:
	//
	// media
	ContentDomain *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// Data type filter
	//
	// - plainText: Plain text
	//
	// - richText: Rich text
	//
	// - html: HTML
	//
	// - pdf: PDF
	//
	// - word: Word
	//
	// - excel: Excel
	//
	// - csv: CSV
	//
	// - image: Image
	//
	// - video: Video
	//
	// - audio: Audio
	//
	// example:
	//
	// plainText
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// End time
	//
	// example:
	//
	// 2024-01-04 11:46:07
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Search keyword: Supports fuzzy search on titles and content
	//
	// example:
	//
	// 检索Query
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Items per page. Default is 10.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Start time
	//
	// example:
	//
	// 2024-01-04 11:46:07
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Task ID
	//
	// > You do not need to specify TaskId. The system generates it automatically. If you use the same TaskId for multiple tasks, those tasks belong to the same conversation.
	//
	// example:
	//
	// task-03d46184ee7d8749
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Title text
	//
	// example:
	//
	// 杭州亚运会
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListGeneratedContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGeneratedContentsRequest) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListGeneratedContentsRequest) GetContentDomain() *string {
	return s.ContentDomain
}

func (s *ListGeneratedContentsRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListGeneratedContentsRequest) GetDataType() *string {
	return s.DataType
}

func (s *ListGeneratedContentsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListGeneratedContentsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListGeneratedContentsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListGeneratedContentsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListGeneratedContentsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListGeneratedContentsRequest) GetTitle() *string {
	return s.Title
}

func (s *ListGeneratedContentsRequest) SetAgentKey(v string) *ListGeneratedContentsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetContentDomain(v string) *ListGeneratedContentsRequest {
	s.ContentDomain = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetCurrent(v int32) *ListGeneratedContentsRequest {
	s.Current = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetDataType(v string) *ListGeneratedContentsRequest {
	s.DataType = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetEndTime(v string) *ListGeneratedContentsRequest {
	s.EndTime = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetQuery(v string) *ListGeneratedContentsRequest {
	s.Query = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetSize(v int32) *ListGeneratedContentsRequest {
	s.Size = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetStartTime(v string) *ListGeneratedContentsRequest {
	s.StartTime = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetTaskId(v string) *ListGeneratedContentsRequest {
	s.TaskId = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetTitle(v string) *ListGeneratedContentsRequest {
	s.Title = &v
	return s
}

func (s *ListGeneratedContentsRequest) Validate() error {
	return dara.Validate(s)
}
