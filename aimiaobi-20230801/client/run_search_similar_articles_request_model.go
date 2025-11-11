// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchSimilarArticlesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatConfig(v *RunSearchSimilarArticlesRequestChatConfig) *RunSearchSimilarArticlesRequest
	GetChatConfig() *RunSearchSimilarArticlesRequestChatConfig
	SetDocType(v string) *RunSearchSimilarArticlesRequest
	GetDocType() *string
	SetTitle(v string) *RunSearchSimilarArticlesRequest
	GetTitle() *string
	SetUrl(v string) *RunSearchSimilarArticlesRequest
	GetUrl() *string
	SetWorkspaceId(v string) *RunSearchSimilarArticlesRequest
	GetWorkspaceId() *string
}

type RunSearchSimilarArticlesRequest struct {
	ChatConfig *RunSearchSimilarArticlesRequestChatConfig `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty" type:"Struct"`
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

func (s RunSearchSimilarArticlesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesRequest) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesRequest) GetChatConfig() *RunSearchSimilarArticlesRequestChatConfig {
	return s.ChatConfig
}

func (s *RunSearchSimilarArticlesRequest) GetDocType() *string {
	return s.DocType
}

func (s *RunSearchSimilarArticlesRequest) GetTitle() *string {
	return s.Title
}

func (s *RunSearchSimilarArticlesRequest) GetUrl() *string {
	return s.Url
}

func (s *RunSearchSimilarArticlesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunSearchSimilarArticlesRequest) SetChatConfig(v *RunSearchSimilarArticlesRequestChatConfig) *RunSearchSimilarArticlesRequest {
	s.ChatConfig = v
	return s
}

func (s *RunSearchSimilarArticlesRequest) SetDocType(v string) *RunSearchSimilarArticlesRequest {
	s.DocType = &v
	return s
}

func (s *RunSearchSimilarArticlesRequest) SetTitle(v string) *RunSearchSimilarArticlesRequest {
	s.Title = &v
	return s
}

func (s *RunSearchSimilarArticlesRequest) SetUrl(v string) *RunSearchSimilarArticlesRequest {
	s.Url = &v
	return s
}

func (s *RunSearchSimilarArticlesRequest) SetWorkspaceId(v string) *RunSearchSimilarArticlesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunSearchSimilarArticlesRequest) Validate() error {
	if s.ChatConfig != nil {
		if err := s.ChatConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchSimilarArticlesRequestChatConfig struct {
	SearchParam *RunSearchSimilarArticlesRequestChatConfigSearchParam `json:"SearchParam,omitempty" xml:"SearchParam,omitempty" type:"Struct"`
}

func (s RunSearchSimilarArticlesRequestChatConfig) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesRequestChatConfig) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesRequestChatConfig) GetSearchParam() *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	return s.SearchParam
}

func (s *RunSearchSimilarArticlesRequestChatConfig) SetSearchParam(v *RunSearchSimilarArticlesRequestChatConfigSearchParam) *RunSearchSimilarArticlesRequestChatConfig {
	s.SearchParam = v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfig) Validate() error {
	if s.SearchParam != nil {
		if err := s.SearchParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSearchSimilarArticlesRequestChatConfigSearchParam struct {
	SearchSources []*RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
}

func (s RunSearchSimilarArticlesRequestChatConfigSearchParam) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesRequestChatConfigSearchParam) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetSearchSources() []*RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources {
	return s.SearchSources
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetSearchSources(v []*RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.SearchSources = v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) Validate() error {
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

type RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources struct {
	// example:
	//
	// SystemSearch
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) GetCode() *string {
	return s.Code
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) GetDatasetName() *string {
	return s.DatasetName
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) GetName() *string {
	return s.Name
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) SetCode(v string) *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources {
	s.Code = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) SetDatasetName(v string) *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources {
	s.DatasetName = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) SetName(v string) *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources {
	s.Name = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) Validate() error {
	return dara.Validate(s)
}
