// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunQuickWritingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticles(v []*RunQuickWritingRequestArticles) *RunQuickWritingRequest
	GetArticles() []*RunQuickWritingRequestArticles
	SetPrompt(v string) *RunQuickWritingRequest
	GetPrompt() *string
	SetSearchSources(v []*RunQuickWritingRequestSearchSources) *RunQuickWritingRequest
	GetSearchSources() []*RunQuickWritingRequestSearchSources
	SetTaskId(v string) *RunQuickWritingRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunQuickWritingRequest
	GetWorkspaceId() *string
}

type RunQuickWritingRequest struct {
	Articles []*RunQuickWritingRequestArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// This parameter is required.
	Prompt        *string                                `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	SearchSources []*RunQuickWritingRequestSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunQuickWritingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingRequest) GoString() string {
	return s.String()
}

func (s *RunQuickWritingRequest) GetArticles() []*RunQuickWritingRequestArticles {
	return s.Articles
}

func (s *RunQuickWritingRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunQuickWritingRequest) GetSearchSources() []*RunQuickWritingRequestSearchSources {
	return s.SearchSources
}

func (s *RunQuickWritingRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunQuickWritingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunQuickWritingRequest) SetArticles(v []*RunQuickWritingRequestArticles) *RunQuickWritingRequest {
	s.Articles = v
	return s
}

func (s *RunQuickWritingRequest) SetPrompt(v string) *RunQuickWritingRequest {
	s.Prompt = &v
	return s
}

func (s *RunQuickWritingRequest) SetSearchSources(v []*RunQuickWritingRequestSearchSources) *RunQuickWritingRequest {
	s.SearchSources = v
	return s
}

func (s *RunQuickWritingRequest) SetTaskId(v string) *RunQuickWritingRequest {
	s.TaskId = &v
	return s
}

func (s *RunQuickWritingRequest) SetWorkspaceId(v string) *RunQuickWritingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunQuickWritingRequest) Validate() error {
	if s.Articles != nil {
		for _, item := range s.Articles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchSources != nil {
		for _, item := range s.SearchSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunQuickWritingRequestArticles struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunQuickWritingRequestArticles) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingRequestArticles) GoString() string {
	return s.String()
}

func (s *RunQuickWritingRequestArticles) GetContent() *string {
	return s.Content
}

func (s *RunQuickWritingRequestArticles) GetTitle() *string {
	return s.Title
}

func (s *RunQuickWritingRequestArticles) GetUrl() *string {
	return s.Url
}

func (s *RunQuickWritingRequestArticles) SetContent(v string) *RunQuickWritingRequestArticles {
	s.Content = &v
	return s
}

func (s *RunQuickWritingRequestArticles) SetTitle(v string) *RunQuickWritingRequestArticles {
	s.Title = &v
	return s
}

func (s *RunQuickWritingRequestArticles) SetUrl(v string) *RunQuickWritingRequestArticles {
	s.Url = &v
	return s
}

func (s *RunQuickWritingRequestArticles) Validate() error {
	return dara.Validate(s)
}

type RunQuickWritingRequestSearchSources struct {
	// example:
	//
	// SystemSearch
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// QuarkCommonNews
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s RunQuickWritingRequestSearchSources) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingRequestSearchSources) GoString() string {
	return s.String()
}

func (s *RunQuickWritingRequestSearchSources) GetCode() *string {
	return s.Code
}

func (s *RunQuickWritingRequestSearchSources) GetDatasetName() *string {
	return s.DatasetName
}

func (s *RunQuickWritingRequestSearchSources) SetCode(v string) *RunQuickWritingRequestSearchSources {
	s.Code = &v
	return s
}

func (s *RunQuickWritingRequestSearchSources) SetDatasetName(v string) *RunQuickWritingRequestSearchSources {
	s.DatasetName = &v
	return s
}

func (s *RunQuickWritingRequestSearchSources) Validate() error {
	return dara.Validate(s)
}
