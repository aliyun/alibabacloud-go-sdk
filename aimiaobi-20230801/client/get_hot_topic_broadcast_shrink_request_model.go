// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotTopicBroadcastShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalcTotalToken(v bool) *GetHotTopicBroadcastShrinkRequest
	GetCalcTotalToken() *bool
	SetCategory(v string) *GetHotTopicBroadcastShrinkRequest
	GetCategory() *string
	SetCurrent(v int32) *GetHotTopicBroadcastShrinkRequest
	GetCurrent() *int32
	SetHotTopicVersion(v string) *GetHotTopicBroadcastShrinkRequest
	GetHotTopicVersion() *string
	SetLocationQuery(v string) *GetHotTopicBroadcastShrinkRequest
	GetLocationQuery() *string
	SetLocationsShrink(v string) *GetHotTopicBroadcastShrinkRequest
	GetLocationsShrink() *string
	SetQuery(v string) *GetHotTopicBroadcastShrinkRequest
	GetQuery() *string
	SetSize(v int32) *GetHotTopicBroadcastShrinkRequest
	GetSize() *int32
	SetStepForCustomSummaryStyleConfigShrink(v string) *GetHotTopicBroadcastShrinkRequest
	GetStepForCustomSummaryStyleConfigShrink() *string
	SetStepForNewsBroadcastContentConfigShrink(v string) *GetHotTopicBroadcastShrinkRequest
	GetStepForNewsBroadcastContentConfigShrink() *string
	SetTopicsShrink(v string) *GetHotTopicBroadcastShrinkRequest
	GetTopicsShrink() *string
	SetWorkspaceId(v string) *GetHotTopicBroadcastShrinkRequest
	GetWorkspaceId() *string
}

type GetHotTopicBroadcastShrinkRequest struct {
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
	HotTopicVersion *string `json:"HotTopicVersion,omitempty" xml:"HotTopicVersion,omitempty"`
	LocationQuery   *string `json:"LocationQuery,omitempty" xml:"LocationQuery,omitempty"`
	LocationsShrink *string `json:"Locations,omitempty" xml:"Locations,omitempty"`
	Query           *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 5
	Size                                    *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	StepForCustomSummaryStyleConfigShrink   *string `json:"StepForCustomSummaryStyleConfig,omitempty" xml:"StepForCustomSummaryStyleConfig,omitempty"`
	StepForNewsBroadcastContentConfigShrink *string `json:"StepForNewsBroadcastContentConfig,omitempty" xml:"StepForNewsBroadcastContentConfig,omitempty"`
	// example:
	//
	// ["主题1","主题2"]
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetHotTopicBroadcastShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastShrinkRequest) GetCalcTotalToken() *bool {
	return s.CalcTotalToken
}

func (s *GetHotTopicBroadcastShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *GetHotTopicBroadcastShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *GetHotTopicBroadcastShrinkRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *GetHotTopicBroadcastShrinkRequest) GetLocationQuery() *string {
	return s.LocationQuery
}

func (s *GetHotTopicBroadcastShrinkRequest) GetLocationsShrink() *string {
	return s.LocationsShrink
}

func (s *GetHotTopicBroadcastShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *GetHotTopicBroadcastShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *GetHotTopicBroadcastShrinkRequest) GetStepForCustomSummaryStyleConfigShrink() *string {
	return s.StepForCustomSummaryStyleConfigShrink
}

func (s *GetHotTopicBroadcastShrinkRequest) GetStepForNewsBroadcastContentConfigShrink() *string {
	return s.StepForNewsBroadcastContentConfigShrink
}

func (s *GetHotTopicBroadcastShrinkRequest) GetTopicsShrink() *string {
	return s.TopicsShrink
}

func (s *GetHotTopicBroadcastShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetHotTopicBroadcastShrinkRequest) SetCalcTotalToken(v bool) *GetHotTopicBroadcastShrinkRequest {
	s.CalcTotalToken = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetCategory(v string) *GetHotTopicBroadcastShrinkRequest {
	s.Category = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetCurrent(v int32) *GetHotTopicBroadcastShrinkRequest {
	s.Current = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetHotTopicVersion(v string) *GetHotTopicBroadcastShrinkRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetLocationQuery(v string) *GetHotTopicBroadcastShrinkRequest {
	s.LocationQuery = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetLocationsShrink(v string) *GetHotTopicBroadcastShrinkRequest {
	s.LocationsShrink = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetQuery(v string) *GetHotTopicBroadcastShrinkRequest {
	s.Query = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetSize(v int32) *GetHotTopicBroadcastShrinkRequest {
	s.Size = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetStepForCustomSummaryStyleConfigShrink(v string) *GetHotTopicBroadcastShrinkRequest {
	s.StepForCustomSummaryStyleConfigShrink = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetStepForNewsBroadcastContentConfigShrink(v string) *GetHotTopicBroadcastShrinkRequest {
	s.StepForNewsBroadcastContentConfigShrink = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetTopicsShrink(v string) *GetHotTopicBroadcastShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) SetWorkspaceId(v string) *GetHotTopicBroadcastShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetHotTopicBroadcastShrinkRequest) Validate() error {
	return dara.Validate(s)
}
