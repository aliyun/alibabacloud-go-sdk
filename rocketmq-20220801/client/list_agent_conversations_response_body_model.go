// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentConversationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAgentConversationsResponseBodyData) *ListAgentConversationsResponseBody
	GetData() []*ListAgentConversationsResponseBodyData
	SetMaxResults(v int32) *ListAgentConversationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentConversationsResponseBody
	GetNextToken() *string
}

type ListAgentConversationsResponseBody struct {
	// example:
	//
	// [{"id":"0198c8a91f4d7f01a3c21e97","title":"消息堆积排查"}]
	Data []*ListAgentConversationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJpZCI6IjAxOThjOGE5In0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAgentConversationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentConversationsResponseBody) GetData() []*ListAgentConversationsResponseBodyData {
	return s.Data
}

func (s *ListAgentConversationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentConversationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentConversationsResponseBody) SetData(v []*ListAgentConversationsResponseBodyData) *ListAgentConversationsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentConversationsResponseBody) SetMaxResults(v int32) *ListAgentConversationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgentConversationsResponseBody) SetNextToken(v string) *ListAgentConversationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentConversationsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentConversationsResponseBodyData struct {
	// example:
	//
	// {"role":"owner","admin":false,"canRead":true,"canWrite":true,"canManage":true}
	Access interface{} `json:"access,omitempty" xml:"access,omitempty"`
	// example:
	//
	// {"runtime":"claude","connectionId":"connection-example","modelId":"model-example","connectionName":"示例连接","modelDisplayName":"示例模型"}
	Agent interface{} `json:"agent,omitempty" xml:"agent,omitempty"`
	// example:
	//
	// {"sessionId":"0198c8a91f4d7f01a3c21e97","purposes":[]}
	AnalysisCoverage interface{} `json:"analysisCoverage,omitempty" xml:"analysisCoverage,omitempty"`
	// example:
	//
	// 0198c8a91f4d7f01a3c21e97
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// example:
	//
	// 2026-08-19T12:00:00.000Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 1000000000000001
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// waiting
	CurrentRunStatus *string `json:"currentRunStatus,omitempty" xml:"currentRunStatus,omitempty"`
	// example:
	//
	// null
	EditMonitorId *string `json:"editMonitorId,omitempty" xml:"editMonitorId,omitempty"`
	// example:
	//
	// 0198c8a91f4d7f01a3c21e97
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// {"userId":"1000000000000001","name":null,"nickName":null,"email":null,"avatar":null}
	Owner interface{} `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// {"id":"project-example","name":"RocketMQ 诊断","color":"#1677ff","initial":"R","agentProfile":"customer"}
	Project interface{} `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// rocketmq-customer
	ScenarioId *string `json:"scenarioId,omitempty" xml:"scenarioId,omitempty"`
	// example:
	//
	// private
	ShareScope *string `json:"shareScope,omitempty" xml:"shareScope,omitempty"`
	// example:
	//
	// null
	SourceSessionId *string `json:"sourceSessionId,omitempty" xml:"sourceSessionId,omitempty"`
	// example:
	//
	// null
	Tagging interface{} `json:"tagging,omitempty" xml:"tagging,omitempty"`
	// example:
	//
	// 消息堆积排查
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 3
	ToolCallCount *int64 `json:"toolCallCount,omitempty" xml:"toolCallCount,omitempty"`
	// example:
	//
	// 2026-08-19T12:05:00.000Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ["rocketmq-agent"]
	UsedSkills []*string `json:"usedSkills,omitempty" xml:"usedSkills,omitempty" type:"Repeated"`
}

func (s ListAgentConversationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentConversationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentConversationsResponseBodyData) GetAccess() interface{} {
	return s.Access
}

func (s *ListAgentConversationsResponseBodyData) GetAgent() interface{} {
	return s.Agent
}

func (s *ListAgentConversationsResponseBodyData) GetAnalysisCoverage() interface{} {
	return s.AnalysisCoverage
}

func (s *ListAgentConversationsResponseBodyData) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAgentConversationsResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListAgentConversationsResponseBodyData) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListAgentConversationsResponseBodyData) GetCurrentRunStatus() *string {
	return s.CurrentRunStatus
}

func (s *ListAgentConversationsResponseBodyData) GetEditMonitorId() *string {
	return s.EditMonitorId
}

func (s *ListAgentConversationsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListAgentConversationsResponseBodyData) GetOwner() interface{} {
	return s.Owner
}

func (s *ListAgentConversationsResponseBodyData) GetProject() interface{} {
	return s.Project
}

func (s *ListAgentConversationsResponseBodyData) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ListAgentConversationsResponseBodyData) GetShareScope() *string {
	return s.ShareScope
}

func (s *ListAgentConversationsResponseBodyData) GetSourceSessionId() *string {
	return s.SourceSessionId
}

func (s *ListAgentConversationsResponseBodyData) GetTagging() interface{} {
	return s.Tagging
}

func (s *ListAgentConversationsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListAgentConversationsResponseBodyData) GetToolCallCount() *int64 {
	return s.ToolCallCount
}

func (s *ListAgentConversationsResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListAgentConversationsResponseBodyData) GetUsedSkills() []*string {
	return s.UsedSkills
}

func (s *ListAgentConversationsResponseBodyData) SetAccess(v interface{}) *ListAgentConversationsResponseBodyData {
	s.Access = v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetAgent(v interface{}) *ListAgentConversationsResponseBodyData {
	s.Agent = v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetAnalysisCoverage(v interface{}) *ListAgentConversationsResponseBodyData {
	s.AnalysisCoverage = v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetConversationId(v string) *ListAgentConversationsResponseBodyData {
	s.ConversationId = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetCreatedAt(v string) *ListAgentConversationsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetCreatedBy(v string) *ListAgentConversationsResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetCurrentRunStatus(v string) *ListAgentConversationsResponseBodyData {
	s.CurrentRunStatus = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetEditMonitorId(v string) *ListAgentConversationsResponseBodyData {
	s.EditMonitorId = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetId(v string) *ListAgentConversationsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetOwner(v interface{}) *ListAgentConversationsResponseBodyData {
	s.Owner = v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetProject(v interface{}) *ListAgentConversationsResponseBodyData {
	s.Project = v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetScenarioId(v string) *ListAgentConversationsResponseBodyData {
	s.ScenarioId = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetShareScope(v string) *ListAgentConversationsResponseBodyData {
	s.ShareScope = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetSourceSessionId(v string) *ListAgentConversationsResponseBodyData {
	s.SourceSessionId = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetTagging(v interface{}) *ListAgentConversationsResponseBodyData {
	s.Tagging = v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetTitle(v string) *ListAgentConversationsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetToolCallCount(v int64) *ListAgentConversationsResponseBodyData {
	s.ToolCallCount = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetUpdatedAt(v string) *ListAgentConversationsResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListAgentConversationsResponseBodyData) SetUsedSkills(v []*string) *ListAgentConversationsResponseBodyData {
	s.UsedSkills = v
	return s
}

func (s *ListAgentConversationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
