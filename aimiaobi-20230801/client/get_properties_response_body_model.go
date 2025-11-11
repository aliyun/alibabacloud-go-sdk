// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPropertiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPropertiesResponseBody
	GetCode() *string
	SetData(v *GetPropertiesResponseBodyData) *GetPropertiesResponseBody
	GetData() *GetPropertiesResponseBodyData
	SetHttpStatusCode(v int32) *GetPropertiesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPropertiesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPropertiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPropertiesResponseBody
	GetSuccess() *bool
}

type GetPropertiesResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPropertiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPropertiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPropertiesResponseBody) GetData() *GetPropertiesResponseBodyData {
	return s.Data
}

func (s *GetPropertiesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPropertiesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPropertiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPropertiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPropertiesResponseBody) SetCode(v string) *GetPropertiesResponseBody {
	s.Code = &v
	return s
}

func (s *GetPropertiesResponseBody) SetData(v *GetPropertiesResponseBodyData) *GetPropertiesResponseBody {
	s.Data = v
	return s
}

func (s *GetPropertiesResponseBody) SetHttpStatusCode(v int32) *GetPropertiesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPropertiesResponseBody) SetMessage(v string) *GetPropertiesResponseBody {
	s.Message = &v
	return s
}

func (s *GetPropertiesResponseBody) SetRequestId(v string) *GetPropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPropertiesResponseBody) SetSuccess(v bool) *GetPropertiesResponseBody {
	s.Success = &v
	return s
}

func (s *GetPropertiesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPropertiesResponseBodyData struct {
	ChatConfig              map[string]interface{}                                `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	ConsoleConfig           *GetPropertiesResponseBodyDataConsoleConfig           `json:"ConsoleConfig,omitempty" xml:"ConsoleConfig,omitempty" type:"Struct"`
	GeneralConfigMap        map[string]interface{}                                `json:"GeneralConfigMap,omitempty" xml:"GeneralConfigMap,omitempty"`
	IntelligentSearchConfig *GetPropertiesResponseBodyDataIntelligentSearchConfig `json:"IntelligentSearchConfig,omitempty" xml:"IntelligentSearchConfig,omitempty" type:"Struct"`
	MiaosouConfig           *GetPropertiesResponseBodyDataMiaosouConfig           `json:"MiaosouConfig,omitempty" xml:"MiaosouConfig,omitempty" type:"Struct"`
	SearchSourceList        []*GetPropertiesResponseBodyDataSearchSourceList      `json:"SearchSourceList,omitempty" xml:"SearchSourceList,omitempty" type:"Repeated"`
	SearchSources           []*GetPropertiesResponseBodyDataSearchSources         `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SlrAuthorized            *bool                                                    `json:"SlrAuthorized,omitempty" xml:"SlrAuthorized,omitempty"`
	UserInfo                 *GetPropertiesResponseBodyDataUserInfo                   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	WanxiangImageSizeConfig  []*GetPropertiesResponseBodyDataWanxiangImageSizeConfig  `json:"WanxiangImageSizeConfig,omitempty" xml:"WanxiangImageSizeConfig,omitempty" type:"Repeated"`
	WanxiangImageStyleConfig []*GetPropertiesResponseBodyDataWanxiangImageStyleConfig `json:"WanxiangImageStyleConfig,omitempty" xml:"WanxiangImageStyleConfig,omitempty" type:"Repeated"`
}

func (s GetPropertiesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyData) GetChatConfig() map[string]interface{} {
	return s.ChatConfig
}

func (s *GetPropertiesResponseBodyData) GetConsoleConfig() *GetPropertiesResponseBodyDataConsoleConfig {
	return s.ConsoleConfig
}

func (s *GetPropertiesResponseBodyData) GetGeneralConfigMap() map[string]interface{} {
	return s.GeneralConfigMap
}

func (s *GetPropertiesResponseBodyData) GetIntelligentSearchConfig() *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	return s.IntelligentSearchConfig
}

func (s *GetPropertiesResponseBodyData) GetMiaosouConfig() *GetPropertiesResponseBodyDataMiaosouConfig {
	return s.MiaosouConfig
}

func (s *GetPropertiesResponseBodyData) GetSearchSourceList() []*GetPropertiesResponseBodyDataSearchSourceList {
	return s.SearchSourceList
}

func (s *GetPropertiesResponseBodyData) GetSearchSources() []*GetPropertiesResponseBodyDataSearchSources {
	return s.SearchSources
}

func (s *GetPropertiesResponseBodyData) GetSlrAuthorized() *bool {
	return s.SlrAuthorized
}

func (s *GetPropertiesResponseBodyData) GetUserInfo() *GetPropertiesResponseBodyDataUserInfo {
	return s.UserInfo
}

func (s *GetPropertiesResponseBodyData) GetWanxiangImageSizeConfig() []*GetPropertiesResponseBodyDataWanxiangImageSizeConfig {
	return s.WanxiangImageSizeConfig
}

func (s *GetPropertiesResponseBodyData) GetWanxiangImageStyleConfig() []*GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	return s.WanxiangImageStyleConfig
}

func (s *GetPropertiesResponseBodyData) SetChatConfig(v map[string]interface{}) *GetPropertiesResponseBodyData {
	s.ChatConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetConsoleConfig(v *GetPropertiesResponseBodyDataConsoleConfig) *GetPropertiesResponseBodyData {
	s.ConsoleConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetGeneralConfigMap(v map[string]interface{}) *GetPropertiesResponseBodyData {
	s.GeneralConfigMap = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetIntelligentSearchConfig(v *GetPropertiesResponseBodyDataIntelligentSearchConfig) *GetPropertiesResponseBodyData {
	s.IntelligentSearchConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetMiaosouConfig(v *GetPropertiesResponseBodyDataMiaosouConfig) *GetPropertiesResponseBodyData {
	s.MiaosouConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetSearchSourceList(v []*GetPropertiesResponseBodyDataSearchSourceList) *GetPropertiesResponseBodyData {
	s.SearchSourceList = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetSearchSources(v []*GetPropertiesResponseBodyDataSearchSources) *GetPropertiesResponseBodyData {
	s.SearchSources = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetSlrAuthorized(v bool) *GetPropertiesResponseBodyData {
	s.SlrAuthorized = &v
	return s
}

func (s *GetPropertiesResponseBodyData) SetUserInfo(v *GetPropertiesResponseBodyDataUserInfo) *GetPropertiesResponseBodyData {
	s.UserInfo = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetWanxiangImageSizeConfig(v []*GetPropertiesResponseBodyDataWanxiangImageSizeConfig) *GetPropertiesResponseBodyData {
	s.WanxiangImageSizeConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetWanxiangImageStyleConfig(v []*GetPropertiesResponseBodyDataWanxiangImageStyleConfig) *GetPropertiesResponseBodyData {
	s.WanxiangImageStyleConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) Validate() error {
	if s.ConsoleConfig != nil {
		if err := s.ConsoleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IntelligentSearchConfig != nil {
		if err := s.IntelligentSearchConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MiaosouConfig != nil {
		if err := s.MiaosouConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SearchSourceList != nil {
		for _, item := range s.SearchSourceList {
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
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	if s.WanxiangImageSizeConfig != nil {
		for _, item := range s.WanxiangImageSizeConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WanxiangImageStyleConfig != nil {
		for _, item := range s.WanxiangImageStyleConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPropertiesResponseBodyDataConsoleConfig struct {
	// example:
	//
	// xx
	TipContent *string `json:"TipContent,omitempty" xml:"TipContent,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetPropertiesResponseBodyDataConsoleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataConsoleConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) GetTipContent() *string {
	return s.TipContent
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) GetTitle() *string {
	return s.Title
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) SetTipContent(v string) *GetPropertiesResponseBodyDataConsoleConfig {
	s.TipContent = &v
	return s
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) SetTitle(v string) *GetPropertiesResponseBodyDataConsoleConfig {
	s.Title = &v
	return s
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataIntelligentSearchConfig struct {
	CopilotPreciseSearchSources []*GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources `json:"CopilotPreciseSearchSources,omitempty" xml:"CopilotPreciseSearchSources,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	ProductDescription *string                                                              `json:"ProductDescription,omitempty" xml:"ProductDescription,omitempty"`
	SearchSamples      []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples `json:"SearchSamples,omitempty" xml:"SearchSamples,omitempty" type:"Repeated"`
	SearchSources      []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) GetCopilotPreciseSearchSources() []*GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources {
	return s.CopilotPreciseSearchSources
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) GetProductDescription() *string {
	return s.ProductDescription
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) GetSearchSamples() []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	return s.SearchSamples
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) GetSearchSources() []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	return s.SearchSources
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetCopilotPreciseSearchSources(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.CopilotPreciseSearchSources = v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetProductDescription(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.ProductDescription = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetSearchSamples(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.SearchSamples = v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetSearchSources(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.SearchSources = v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) Validate() error {
	if s.CopilotPreciseSearchSources != nil {
		for _, item := range s.CopilotPreciseSearchSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchSamples != nil {
		for _, item := range s.SearchSamples {
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

type GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources struct {
	// example:
	//
	// x
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// x
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// x
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) GetCode() *string {
	return s.Code
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) GetName() *string {
	return s.Name
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) SetCode(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources {
	s.Code = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) SetDatasetName(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources {
	s.DatasetName = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) SetName(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigCopilotPreciseSearchSources) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples struct {
	Articles []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) GetArticles() []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	return s.Articles
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) GetPrompt() *string {
	return s.Prompt
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) GetText() *string {
	return s.Text
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetArticles(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Articles = v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetPrompt(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Prompt = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetText(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Text = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) Validate() error {
	if s.Articles != nil {
		for _, item := range s.Articles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles struct {
	// example:
	//
	// true
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
	// example:
	//
	// false
	Stared *bool `json:"Stared,omitempty" xml:"Stared,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxx.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) GetSelect() *bool {
	return s.Select
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) GetStared() *bool {
	return s.Stared
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) GetTitle() *string {
	return s.Title
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) GetUrl() *string {
	return s.Url
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetSelect(v bool) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Select = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetStared(v bool) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Stared = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetTitle(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Title = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetUrl(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Url = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources struct {
	// example:
	//
	// xx
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xx
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) GetCode() *string {
	return s.Code
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) GetName() *string {
	return s.Name
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetCode(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.Code = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetDatasetName(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.DatasetName = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetName(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataMiaosouConfig struct {
	// example:
	//
	// 1
	MaxDocSize *int64                                                  `json:"MaxDocSize,omitempty" xml:"MaxDocSize,omitempty"`
	ModelInfos []*GetPropertiesResponseBodyDataMiaosouConfigModelInfos `json:"ModelInfos,omitempty" xml:"ModelInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	UseDocSize *int64 `json:"UseDocSize,omitempty" xml:"UseDocSize,omitempty"`
}

func (s GetPropertiesResponseBodyDataMiaosouConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataMiaosouConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataMiaosouConfig) GetMaxDocSize() *int64 {
	return s.MaxDocSize
}

func (s *GetPropertiesResponseBodyDataMiaosouConfig) GetModelInfos() []*GetPropertiesResponseBodyDataMiaosouConfigModelInfos {
	return s.ModelInfos
}

func (s *GetPropertiesResponseBodyDataMiaosouConfig) GetUseDocSize() *int64 {
	return s.UseDocSize
}

func (s *GetPropertiesResponseBodyDataMiaosouConfig) SetMaxDocSize(v int64) *GetPropertiesResponseBodyDataMiaosouConfig {
	s.MaxDocSize = &v
	return s
}

func (s *GetPropertiesResponseBodyDataMiaosouConfig) SetModelInfos(v []*GetPropertiesResponseBodyDataMiaosouConfigModelInfos) *GetPropertiesResponseBodyDataMiaosouConfig {
	s.ModelInfos = v
	return s
}

func (s *GetPropertiesResponseBodyDataMiaosouConfig) SetUseDocSize(v int64) *GetPropertiesResponseBodyDataMiaosouConfig {
	s.UseDocSize = &v
	return s
}

func (s *GetPropertiesResponseBodyDataMiaosouConfig) Validate() error {
	if s.ModelInfos != nil {
		for _, item := range s.ModelInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPropertiesResponseBodyDataMiaosouConfigModelInfos struct {
	// example:
	//
	// quanmiao-max
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// 全妙-Max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s GetPropertiesResponseBodyDataMiaosouConfigModelInfos) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataMiaosouConfigModelInfos) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataMiaosouConfigModelInfos) GetModelId() *string {
	return s.ModelId
}

func (s *GetPropertiesResponseBodyDataMiaosouConfigModelInfos) GetModelName() *string {
	return s.ModelName
}

func (s *GetPropertiesResponseBodyDataMiaosouConfigModelInfos) SetModelId(v string) *GetPropertiesResponseBodyDataMiaosouConfigModelInfos {
	s.ModelId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataMiaosouConfigModelInfos) SetModelName(v string) *GetPropertiesResponseBodyDataMiaosouConfigModelInfos {
	s.ModelName = &v
	return s
}

func (s *GetPropertiesResponseBodyDataMiaosouConfigModelInfos) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataSearchSourceList struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPropertiesResponseBodyDataSearchSourceList) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataSearchSourceList) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataSearchSourceList) GetCode() *string {
	return s.Code
}

func (s *GetPropertiesResponseBodyDataSearchSourceList) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetPropertiesResponseBodyDataSearchSourceList) GetName() *string {
	return s.Name
}

func (s *GetPropertiesResponseBodyDataSearchSourceList) SetCode(v string) *GetPropertiesResponseBodyDataSearchSourceList {
	s.Code = &v
	return s
}

func (s *GetPropertiesResponseBodyDataSearchSourceList) SetDatasetName(v string) *GetPropertiesResponseBodyDataSearchSourceList {
	s.DatasetName = &v
	return s
}

func (s *GetPropertiesResponseBodyDataSearchSourceList) SetName(v string) *GetPropertiesResponseBodyDataSearchSourceList {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataSearchSourceList) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataSearchSources struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// SystemSearch
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataSearchSources) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataSearchSources) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataSearchSources) GetLabel() *string {
	return s.Label
}

func (s *GetPropertiesResponseBodyDataSearchSources) GetValue() *string {
	return s.Value
}

func (s *GetPropertiesResponseBodyDataSearchSources) SetLabel(v string) *GetPropertiesResponseBodyDataSearchSources {
	s.Label = &v
	return s
}

func (s *GetPropertiesResponseBodyDataSearchSources) SetValue(v string) *GetPropertiesResponseBodyDataSearchSources {
	s.Value = &v
	return s
}

func (s *GetPropertiesResponseBodyDataSearchSources) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataUserInfo struct {
	// example:
	//
	// 1
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 1
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetPropertiesResponseBodyDataUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataUserInfo) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataUserInfo) GetAgentId() *string {
	return s.AgentId
}

func (s *GetPropertiesResponseBodyDataUserInfo) GetTenantId() *string {
	return s.TenantId
}

func (s *GetPropertiesResponseBodyDataUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *GetPropertiesResponseBodyDataUserInfo) GetUsername() *string {
	return s.Username
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetAgentId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.AgentId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetTenantId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.TenantId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetUserId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.UserId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetUsername(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.Username = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataWanxiangImageSizeConfig struct {
	// example:
	//
	// 1:1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1024*1024
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataWanxiangImageSizeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataWanxiangImageSizeConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) GetName() *string {
	return s.Name
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) GetValue() *string {
	return s.Value
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) SetName(v string) *GetPropertiesResponseBodyDataWanxiangImageSizeConfig {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) SetValue(v string) *GetPropertiesResponseBodyDataWanxiangImageSizeConfig {
	s.Value = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) Validate() error {
	return dara.Validate(s)
}

type GetPropertiesResponseBodyDataWanxiangImageStyleConfig struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01RzKicz1W0YWzYkWcK_!!6000000002726-2-tps-132-104.png
	Pic *string `json:"Pic,omitempty" xml:"Pic,omitempty"`
	// example:
	//
	// <auto>
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataWanxiangImageStyleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponseBodyDataWanxiangImageStyleConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) GetName() *string {
	return s.Name
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) GetPic() *string {
	return s.Pic
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) GetValue() *string {
	return s.Value
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetName(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetPic(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Pic = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetValue(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Value = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) Validate() error {
	return dara.Validate(s)
}
