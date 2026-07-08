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
	// Communication configuration parameters.
	ChatConfig *RunSearchSimilarArticlesRequestChatConfig `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty" type:"Struct"`
	// Document type.
	//
	// example:
	//
	// html
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Article title.
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Article URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Unique identifier of the Alibaba Cloud Model Studio workspace. To get this ID, see [Get the workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
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
	// Search configuration parameters.
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
	// Category UUID
	CategoryUuids []*string `json:"CategoryUuids,omitempty" xml:"CategoryUuids,omitempty" type:"Repeated"`
	// Creation Time cutoff, in UNIX timestamp format.
	//
	// example:
	//
	// 111
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// Start Creation Time.
	//
	// example:
	//
	// 111
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// Document ID
	DocIds []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	// Document types: text, image, video, audio, pdf, word, ppt, etc.
	DocTypes []*string `json:"DocTypes,omitempty" xml:"DocTypes,omitempty" type:"Repeated"`
	// Document UUID
	DocUuids []*string `json:"DocUuids,omitempty" xml:"DocUuids,omitempty" type:"Repeated"`
	// End Time
	//
	// example:
	//
	// 111
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Extension Field 1
	//
	// example:
	//
	// xx
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	// Extension Field 2
	//
	// example:
	//
	// xx
	Extend2 *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	// Extension Field 3
	//
	// example:
	//
	// xx
	Extend3 *string `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	// Search sources.
	SearchSources []*RunSearchSimilarArticlesRequestChatConfigSearchParamSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	// Start Time
	//
	// example:
	//
	// 1725983999999
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Tag Name
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// Search source type:
	//
	// - SystemSearch: Built-in system search.
	//
	// - CustomSemanticSearch: Custom semantic index search.
	//
	// - ThirdSearch: Third-party API search.
	//
	// example:
	//
	// SystemSearch
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Unique identifier of the search source.
	//
	// example:
	//
	// QuarkCommonNews
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// Search source name (optional).
	//
	// example:
	//
	// 互联网搜索
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
