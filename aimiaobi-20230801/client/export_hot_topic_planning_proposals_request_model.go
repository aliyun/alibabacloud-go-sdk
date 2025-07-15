// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportHotTopicPlanningProposalsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAgentKey(v string) *ExportHotTopicPlanningProposalsRequest
  GetAgentKey() *string 
  SetCustomViewPointIds(v []*string) *ExportHotTopicPlanningProposalsRequest
  GetCustomViewPointIds() []*string 
  SetExportType(v string) *ExportHotTopicPlanningProposalsRequest
  GetExportType() *string 
  SetTitles(v []*string) *ExportHotTopicPlanningProposalsRequest
  GetTitles() []*string 
  SetTopic(v string) *ExportHotTopicPlanningProposalsRequest
  GetTopic() *string 
  SetTopicSource(v string) *ExportHotTopicPlanningProposalsRequest
  GetTopicSource() *string 
  SetViewPointType(v string) *ExportHotTopicPlanningProposalsRequest
  GetViewPointType() *string 
}

type ExportHotTopicPlanningProposalsRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // xxxxx_p_efm
  AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
  // example:
  // 
  // 025c6cee437741368098b790c90166f8
  CustomViewPointIds []*string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty" type:"Repeated"`
  // example:
  // 
  // 导出文档类型，word:导出为word,xmind:导处为xmind
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  Titles []*string `json:"Titles,omitempty" xml:"Titles,omitempty" type:"Repeated"`
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

func (s ExportHotTopicPlanningProposalsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportHotTopicPlanningProposalsRequest) GoString() string {
  return s.String()
}

func (s *ExportHotTopicPlanningProposalsRequest) GetAgentKey() *string  {
  return s.AgentKey
}

func (s *ExportHotTopicPlanningProposalsRequest) GetCustomViewPointIds() []*string  {
  return s.CustomViewPointIds
}

func (s *ExportHotTopicPlanningProposalsRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportHotTopicPlanningProposalsRequest) GetTitles() []*string  {
  return s.Titles
}

func (s *ExportHotTopicPlanningProposalsRequest) GetTopic() *string  {
  return s.Topic
}

func (s *ExportHotTopicPlanningProposalsRequest) GetTopicSource() *string  {
  return s.TopicSource
}

func (s *ExportHotTopicPlanningProposalsRequest) GetViewPointType() *string  {
  return s.ViewPointType
}

func (s *ExportHotTopicPlanningProposalsRequest) SetAgentKey(v string) *ExportHotTopicPlanningProposalsRequest {
  s.AgentKey = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetCustomViewPointIds(v []*string) *ExportHotTopicPlanningProposalsRequest {
  s.CustomViewPointIds = v
  return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetExportType(v string) *ExportHotTopicPlanningProposalsRequest {
  s.ExportType = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetTitles(v []*string) *ExportHotTopicPlanningProposalsRequest {
  s.Titles = v
  return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetTopic(v string) *ExportHotTopicPlanningProposalsRequest {
  s.Topic = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetTopicSource(v string) *ExportHotTopicPlanningProposalsRequest {
  s.TopicSource = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetViewPointType(v string) *ExportHotTopicPlanningProposalsRequest {
  s.ViewPointType = &v
  return s
}

func (s *ExportHotTopicPlanningProposalsRequest) Validate() error {
  return dara.Validate(s)
}

