// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportHotTopicPlanningProposalsShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAgentKey(v string) *ExportHotTopicPlanningProposalsShrinkRequest
  GetAgentKey() *string 
  SetCustomViewPointIdsShrink(v string) *ExportHotTopicPlanningProposalsShrinkRequest
  GetCustomViewPointIdsShrink() *string 
  SetExportType(v string) *ExportHotTopicPlanningProposalsShrinkRequest
  GetExportType() *string 
  SetTitlesShrink(v string) *ExportHotTopicPlanningProposalsShrinkRequest
  GetTitlesShrink() *string 
  SetTopic(v string) *ExportHotTopicPlanningProposalsShrinkRequest
  GetTopic() *string 
  SetTopicSource(v string) *ExportHotTopicPlanningProposalsShrinkRequest
  GetTopicSource() *string 
  SetViewPointType(v string) *ExportHotTopicPlanningProposalsShrinkRequest
  GetViewPointType() *string 
}

type ExportHotTopicPlanningProposalsShrinkRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // xxxxx_p_efm
  AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
  // example:
  // 
  // 025c6cee437741368098b790c90166f8
  CustomViewPointIdsShrink *string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty"`
  // example:
  // 
  // 导出文档类型，word:导出为word,xmind:导处为xmind
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  TitlesShrink *string `json:"Titles,omitempty" xml:"Titles,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 热榜主题
  Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 热榜源
  TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
  // example:
  // 
  // 选题策划类型：CustomViewPoints:自定义视角，HotViewPoints:热门视角、TimedViewPoints:时效性视角、WebReviewPoints:网友视角、FreshViewPoints:新颖视角
  ViewPointType *string `json:"ViewPointType,omitempty" xml:"ViewPointType,omitempty"`
}

func (s ExportHotTopicPlanningProposalsShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportHotTopicPlanningProposalsShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) GetAgentKey() *string  {
  return s.AgentKey
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) GetCustomViewPointIdsShrink() *string  {
  return s.CustomViewPointIdsShrink
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) GetTitlesShrink() *string  {
  return s.TitlesShrink
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) GetTopic() *string  {
  return s.Topic
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) GetTopicSource() *string  {
  return s.TopicSource
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) GetViewPointType() *string  {
  return s.ViewPointType
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetAgentKey(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
  s.AgentKey = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetCustomViewPointIdsShrink(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
  s.CustomViewPointIdsShrink = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetExportType(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
  s.ExportType = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetTitlesShrink(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
  s.TitlesShrink = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetTopic(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
  s.Topic = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetTopicSource(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
  s.TopicSource = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetViewPointType(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
  s.ViewPointType = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) Validate() error {
  return dara.Validate(s)
}

