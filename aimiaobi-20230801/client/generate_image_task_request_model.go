// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GenerateImageTaskRequest
	GetAgentKey() *string
	SetArticleTaskId(v string) *GenerateImageTaskRequest
	GetArticleTaskId() *string
	SetParagraphList(v []*GenerateImageTaskRequestParagraphList) *GenerateImageTaskRequest
	GetParagraphList() []*GenerateImageTaskRequestParagraphList
	SetSize(v string) *GenerateImageTaskRequest
	GetSize() *string
	SetStyle(v string) *GenerateImageTaskRequest
	GetStyle() *string
}

type GenerateImageTaskRequest struct {
	// The unique identifier of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The task ID of the article. If you do not have one, you can assign a universally unique identifier (UUID).
	//
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	ArticleTaskId *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	// The content of the paragraphs.
	//
	// This parameter is required.
	ParagraphList []*GenerateImageTaskRequestParagraphList `json:"ParagraphList,omitempty" xml:"ParagraphList,omitempty" type:"Repeated"`
	// The size of the image to generate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024*1024
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The style.
	//
	// This parameter is required.
	//
	// example:
	//
	// \\"<auto>\\"
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
}

func (s GenerateImageTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageTaskRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GenerateImageTaskRequest) GetArticleTaskId() *string {
	return s.ArticleTaskId
}

func (s *GenerateImageTaskRequest) GetParagraphList() []*GenerateImageTaskRequestParagraphList {
	return s.ParagraphList
}

func (s *GenerateImageTaskRequest) GetSize() *string {
	return s.Size
}

func (s *GenerateImageTaskRequest) GetStyle() *string {
	return s.Style
}

func (s *GenerateImageTaskRequest) SetAgentKey(v string) *GenerateImageTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateImageTaskRequest) SetArticleTaskId(v string) *GenerateImageTaskRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *GenerateImageTaskRequest) SetParagraphList(v []*GenerateImageTaskRequestParagraphList) *GenerateImageTaskRequest {
	s.ParagraphList = v
	return s
}

func (s *GenerateImageTaskRequest) SetSize(v string) *GenerateImageTaskRequest {
	s.Size = &v
	return s
}

func (s *GenerateImageTaskRequest) SetStyle(v string) *GenerateImageTaskRequest {
	s.Style = &v
	return s
}

func (s *GenerateImageTaskRequest) Validate() error {
	if s.ParagraphList != nil {
		for _, item := range s.ParagraphList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateImageTaskRequestParagraphList struct {
	// The content of the paragraph.
	//
	// This parameter is required.
	//
	// example:
	//
	// 一直忧伤的猫
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The paragraph ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique ID of the task.
	//
	// > By default, you do not need to specify this parameter. The system automatically generates a task ID. If you specify the same TaskId for subsequent tasks, these tasks are considered part of the same conversation group.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The current status of the task.
	//
	// - PENDING: The task is in the queue.
	//
	// - RUNNING: The task is in progress.
	//
	// - SUSPENDED: The task is suspended.
	//
	// - SUCCEEDED: The task was successful.
	//
	// - FAILED: The task failed.
	//
	// - UNKNOWN: The task does not exist or its status is unknown.
	//
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GenerateImageTaskRequestParagraphList) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageTaskRequestParagraphList) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskRequestParagraphList) GetContent() *string {
	return s.Content
}

func (s *GenerateImageTaskRequestParagraphList) GetId() *int64 {
	return s.Id
}

func (s *GenerateImageTaskRequestParagraphList) GetTaskId() *string {
	return s.TaskId
}

func (s *GenerateImageTaskRequestParagraphList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GenerateImageTaskRequestParagraphList) SetContent(v string) *GenerateImageTaskRequestParagraphList {
	s.Content = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetId(v int64) *GenerateImageTaskRequestParagraphList {
	s.Id = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetTaskId(v string) *GenerateImageTaskRequestParagraphList {
	s.TaskId = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetTaskStatus(v string) *GenerateImageTaskRequestParagraphList {
	s.TaskStatus = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) Validate() error {
	return dara.Validate(s)
}
