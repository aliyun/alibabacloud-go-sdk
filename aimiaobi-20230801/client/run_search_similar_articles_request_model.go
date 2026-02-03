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
	CategoryUuids   []*string                                                            `json:"CategoryUuids,omitempty" xml:"CategoryUuids,omitempty" type:"Repeated"`
	CreateTimeEnd   *int64                                                               `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart *int64                                                               `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	DocIds          []*string                                                            `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	DocTypes        []*string                                                            `json:"DocTypes,omitempty" xml:"DocTypes,omitempty" type:"Repeated"`
	DocUuids        []*string                                                            `json:"DocUuids,omitempty" xml:"DocUuids,omitempty" type:"Repeated"`
	EndTime         *int64                                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Extend1         *string                                                              `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	Extend2         *string                                                              `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	Extend3         *string                                                              `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	SearchSources   []*RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	StartTime       *int64                                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tags            []*string                                                            `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RunSearchSimilarArticlesRequestChatConfigSearchParam) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesRequestChatConfigSearchParam) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetCategoryUuids() []*string {
	return s.CategoryUuids
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetDocIds() []*string {
	return s.DocIds
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetDocTypes() []*string {
	return s.DocTypes
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetDocUuids() []*string {
	return s.DocUuids
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetExtend1() *string {
	return s.Extend1
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetExtend2() *string {
	return s.Extend2
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetExtend3() *string {
	return s.Extend3
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetSearchSources() []*RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources {
	return s.SearchSources
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetStartTime() *int64 {
	return s.StartTime
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) GetTags() []*string {
	return s.Tags
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetCategoryUuids(v []*string) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.CategoryUuids = v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetCreateTimeEnd(v int64) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.CreateTimeEnd = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetCreateTimeStart(v int64) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.CreateTimeStart = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetDocIds(v []*string) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.DocIds = v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetDocTypes(v []*string) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.DocTypes = v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetDocUuids(v []*string) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.DocUuids = v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetEndTime(v int64) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.EndTime = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetExtend1(v string) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.Extend1 = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetExtend2(v string) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.Extend2 = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetExtend3(v string) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.Extend3 = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetSearchSources(v []*RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.SearchSources = v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetStartTime(v int64) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.StartTime = &v
	return s
}

func (s *RunSearchSimilarArticlesRequestChatConfigSearchParam) SetTags(v []*string) *RunSearchSimilarArticlesRequestChatConfigSearchParam {
	s.Tags = v
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
