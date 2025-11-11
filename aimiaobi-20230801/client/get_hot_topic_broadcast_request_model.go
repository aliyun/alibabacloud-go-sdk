// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotTopicBroadcastRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalcTotalToken(v bool) *GetHotTopicBroadcastRequest
	GetCalcTotalToken() *bool
	SetCategory(v string) *GetHotTopicBroadcastRequest
	GetCategory() *string
	SetCurrent(v int32) *GetHotTopicBroadcastRequest
	GetCurrent() *int32
	SetHotTopicVersion(v string) *GetHotTopicBroadcastRequest
	GetHotTopicVersion() *string
	SetLocationQuery(v string) *GetHotTopicBroadcastRequest
	GetLocationQuery() *string
	SetLocations(v []*string) *GetHotTopicBroadcastRequest
	GetLocations() []*string
	SetQuery(v string) *GetHotTopicBroadcastRequest
	GetQuery() *string
	SetSize(v int32) *GetHotTopicBroadcastRequest
	GetSize() *int32
	SetStepForCustomSummaryStyleConfig(v *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) *GetHotTopicBroadcastRequest
	GetStepForCustomSummaryStyleConfig() *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig
	SetStepForNewsBroadcastContentConfig(v *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) *GetHotTopicBroadcastRequest
	GetStepForNewsBroadcastContentConfig() *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig
	SetTopics(v []*string) *GetHotTopicBroadcastRequest
	GetTopics() []*string
	SetWorkspaceId(v string) *GetHotTopicBroadcastRequest
	GetWorkspaceId() *string
}

type GetHotTopicBroadcastRequest struct {
	// example:
	//
	// false
	CalcTotalToken *bool `json:"CalcTotalToken,omitempty" xml:"CalcTotalToken,omitempty"`
	// example:
	//
	// 分类筛选
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 2024-10-11_13
	HotTopicVersion *string   `json:"HotTopicVersion,omitempty" xml:"HotTopicVersion,omitempty"`
	LocationQuery   *string   `json:"LocationQuery,omitempty" xml:"LocationQuery,omitempty"`
	Locations       []*string `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Repeated"`
	Query           *string   `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 5
	Size                              *int32                                                        `json:"Size,omitempty" xml:"Size,omitempty"`
	StepForCustomSummaryStyleConfig   *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig   `json:"StepForCustomSummaryStyleConfig,omitempty" xml:"StepForCustomSummaryStyleConfig,omitempty" type:"Struct"`
	StepForNewsBroadcastContentConfig *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig `json:"StepForNewsBroadcastContentConfig,omitempty" xml:"StepForNewsBroadcastContentConfig,omitempty" type:"Struct"`
	// example:
	//
	// ["主题1","主题2"]
	Topics []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetHotTopicBroadcastRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastRequest) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastRequest) GetCalcTotalToken() *bool {
	return s.CalcTotalToken
}

func (s *GetHotTopicBroadcastRequest) GetCategory() *string {
	return s.Category
}

func (s *GetHotTopicBroadcastRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *GetHotTopicBroadcastRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *GetHotTopicBroadcastRequest) GetLocationQuery() *string {
	return s.LocationQuery
}

func (s *GetHotTopicBroadcastRequest) GetLocations() []*string {
	return s.Locations
}

func (s *GetHotTopicBroadcastRequest) GetQuery() *string {
	return s.Query
}

func (s *GetHotTopicBroadcastRequest) GetSize() *int32 {
	return s.Size
}

func (s *GetHotTopicBroadcastRequest) GetStepForCustomSummaryStyleConfig() *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig {
	return s.StepForCustomSummaryStyleConfig
}

func (s *GetHotTopicBroadcastRequest) GetStepForNewsBroadcastContentConfig() *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig {
	return s.StepForNewsBroadcastContentConfig
}

func (s *GetHotTopicBroadcastRequest) GetTopics() []*string {
	return s.Topics
}

func (s *GetHotTopicBroadcastRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetHotTopicBroadcastRequest) SetCalcTotalToken(v bool) *GetHotTopicBroadcastRequest {
	s.CalcTotalToken = &v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetCategory(v string) *GetHotTopicBroadcastRequest {
	s.Category = &v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetCurrent(v int32) *GetHotTopicBroadcastRequest {
	s.Current = &v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetHotTopicVersion(v string) *GetHotTopicBroadcastRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetLocationQuery(v string) *GetHotTopicBroadcastRequest {
	s.LocationQuery = &v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetLocations(v []*string) *GetHotTopicBroadcastRequest {
	s.Locations = v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetQuery(v string) *GetHotTopicBroadcastRequest {
	s.Query = &v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetSize(v int32) *GetHotTopicBroadcastRequest {
	s.Size = &v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetStepForCustomSummaryStyleConfig(v *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) *GetHotTopicBroadcastRequest {
	s.StepForCustomSummaryStyleConfig = v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetStepForNewsBroadcastContentConfig(v *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) *GetHotTopicBroadcastRequest {
	s.StepForNewsBroadcastContentConfig = v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetTopics(v []*string) *GetHotTopicBroadcastRequest {
	s.Topics = v
	return s
}

func (s *GetHotTopicBroadcastRequest) SetWorkspaceId(v string) *GetHotTopicBroadcastRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetHotTopicBroadcastRequest) Validate() error {
	if s.StepForCustomSummaryStyleConfig != nil {
		if err := s.StepForCustomSummaryStyleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.StepForNewsBroadcastContentConfig != nil {
		if err := s.StepForNewsBroadcastContentConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig struct {
	// example:
	//
	// 90
	SummaryImageCount *int32 `json:"SummaryImageCount,omitempty" xml:"SummaryImageCount,omitempty"`
	// example:
	//
	// 摘要模型
	SummaryModel *string `json:"SummaryModel,omitempty" xml:"SummaryModel,omitempty"`
	// example:
	//
	// 摘要-自定义Prompt
	SummaryPrompt *string `json:"SummaryPrompt,omitempty" xml:"SummaryPrompt,omitempty"`
}

func (s GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) GetSummaryImageCount() *int32 {
	return s.SummaryImageCount
}

func (s *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) GetSummaryModel() *string {
	return s.SummaryModel
}

func (s *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) GetSummaryPrompt() *string {
	return s.SummaryPrompt
}

func (s *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) SetSummaryImageCount(v int32) *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig {
	s.SummaryImageCount = &v
	return s
}

func (s *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) SetSummaryModel(v string) *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig {
	s.SummaryModel = &v
	return s
}

func (s *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) SetSummaryPrompt(v string) *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig {
	s.SummaryPrompt = &v
	return s
}

func (s *GetHotTopicBroadcastRequestStepForCustomSummaryStyleConfig) Validate() error {
	return dara.Validate(s)
}

type GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig struct {
	// example:
	//
	// ["科技","经济","时政","娱乐"]
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// Deprecated
	CustomHotValueWeights []*GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights `json:"CustomHotValueWeights,omitempty" xml:"CustomHotValueWeights,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TopicCount *int32 `json:"TopicCount,omitempty" xml:"TopicCount,omitempty"`
}

func (s GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) GetCategories() []*string {
	return s.Categories
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) GetCustomHotValueWeights() []*GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights {
	return s.CustomHotValueWeights
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) GetTopicCount() *int32 {
	return s.TopicCount
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) SetCategories(v []*string) *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig {
	s.Categories = v
	return s
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) SetCustomHotValueWeights(v []*GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights) *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig {
	s.CustomHotValueWeights = v
	return s
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) SetTopicCount(v int32) *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig {
	s.TopicCount = &v
	return s
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfig) Validate() error {
	if s.CustomHotValueWeights != nil {
		for _, item := range s.CustomHotValueWeights {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights struct {
	// example:
	//
	// views
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights) GetDimension() *string {
	return s.Dimension
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights) SetDimension(v string) *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights {
	s.Dimension = &v
	return s
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights) SetWeight(v int32) *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights {
	s.Weight = &v
	return s
}

func (s *GetHotTopicBroadcastRequestStepForNewsBroadcastContentConfigCustomHotValueWeights) Validate() error {
	return dara.Validate(s)
}
