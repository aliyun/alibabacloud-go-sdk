// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchSimilarArticlesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatConfigShrink(v string) *RunSearchSimilarArticlesShrinkRequest
	GetChatConfigShrink() *string
	SetDocType(v string) *RunSearchSimilarArticlesShrinkRequest
	GetDocType() *string
	SetTitle(v string) *RunSearchSimilarArticlesShrinkRequest
	GetTitle() *string
	SetUrl(v string) *RunSearchSimilarArticlesShrinkRequest
	GetUrl() *string
	SetWorkspaceId(v string) *RunSearchSimilarArticlesShrinkRequest
	GetWorkspaceId() *string
}

type RunSearchSimilarArticlesShrinkRequest struct {
	ChatConfigShrink *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// example:
	//
	// html
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunSearchSimilarArticlesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesShrinkRequest) GetChatConfigShrink() *string {
	return s.ChatConfigShrink
}

func (s *RunSearchSimilarArticlesShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *RunSearchSimilarArticlesShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *RunSearchSimilarArticlesShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *RunSearchSimilarArticlesShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunSearchSimilarArticlesShrinkRequest) SetChatConfigShrink(v string) *RunSearchSimilarArticlesShrinkRequest {
	s.ChatConfigShrink = &v
	return s
}

func (s *RunSearchSimilarArticlesShrinkRequest) SetDocType(v string) *RunSearchSimilarArticlesShrinkRequest {
	s.DocType = &v
	return s
}

func (s *RunSearchSimilarArticlesShrinkRequest) SetTitle(v string) *RunSearchSimilarArticlesShrinkRequest {
	s.Title = &v
	return s
}

func (s *RunSearchSimilarArticlesShrinkRequest) SetUrl(v string) *RunSearchSimilarArticlesShrinkRequest {
	s.Url = &v
	return s
}

func (s *RunSearchSimilarArticlesShrinkRequest) SetWorkspaceId(v string) *RunSearchSimilarArticlesShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunSearchSimilarArticlesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
