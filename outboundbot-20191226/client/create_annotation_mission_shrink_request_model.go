// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnotationMissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateAnnotationMissionShrinkRequest
	GetAgentId() *string
	SetAgentKey(v string) *CreateAnnotationMissionShrinkRequest
	GetAgentKey() *string
	SetAnnotationMissionDataSourceType(v int32) *CreateAnnotationMissionShrinkRequest
	GetAnnotationMissionDataSourceType() *int32
	SetAnnotationMissionDebugDataSourceListShrink(v string) *CreateAnnotationMissionShrinkRequest
	GetAnnotationMissionDebugDataSourceListShrink() *string
	SetAnnotationMissionDebugDataSourceListJsonString(v string) *CreateAnnotationMissionShrinkRequest
	GetAnnotationMissionDebugDataSourceListJsonString() *string
	SetAnnotationMissionName(v string) *CreateAnnotationMissionShrinkRequest
	GetAnnotationMissionName() *string
	SetChatbotId(v string) *CreateAnnotationMissionShrinkRequest
	GetChatbotId() *string
	SetConversationTimeEndFilter(v int64) *CreateAnnotationMissionShrinkRequest
	GetConversationTimeEndFilter() *int64
	SetConversationTimeStartFilter(v int64) *CreateAnnotationMissionShrinkRequest
	GetConversationTimeStartFilter() *int64
	SetExcludeOtherSession(v bool) *CreateAnnotationMissionShrinkRequest
	GetExcludeOtherSession() *bool
	SetFinished(v bool) *CreateAnnotationMissionShrinkRequest
	GetFinished() *bool
	SetInstanceId(v string) *CreateAnnotationMissionShrinkRequest
	GetInstanceId() *string
	SetSamplingCount(v int32) *CreateAnnotationMissionShrinkRequest
	GetSamplingCount() *int32
	SetSamplingRate(v int32) *CreateAnnotationMissionShrinkRequest
	GetSamplingRate() *int32
	SetSamplingType(v int32) *CreateAnnotationMissionShrinkRequest
	GetSamplingType() *int32
	SetScriptId(v string) *CreateAnnotationMissionShrinkRequest
	GetScriptId() *string
	SetSessionEndReasonFilterList(v []*int32) *CreateAnnotationMissionShrinkRequest
	GetSessionEndReasonFilterList() []*int32
	SetSessionEndReasonFilterListJsonString(v string) *CreateAnnotationMissionShrinkRequest
	GetSessionEndReasonFilterListJsonString() *string
}

type CreateAnnotationMissionShrinkRequest struct {
	// Agent ID
	//
	// > You can obtain it from the DescribeInstance API.
	//
	// example:
	//
	// 1168702
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Workspace key
	//
	// > You can obtain this from the DescribeInstance API.
	//
	// example:
	//
	// 9137ab9c27044921860030adf8590ec4_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Data Type of the annotation Job.
	//
	// - 1: Outbound call
	//
	// example:
	//
	// 1
	AnnotationMissionDataSourceType *int32 `json:"AnnotationMissionDataSourceType,omitempty" xml:"AnnotationMissionDataSourceType,omitempty"`
	// List of annotation data sources
	//
	// > This parameter serves the same purpose as AnnotationMissionDebugDataSourceListJsonString. You can use either one.
	AnnotationMissionDebugDataSourceListShrink *string `json:"AnnotationMissionDebugDataSourceList,omitempty" xml:"AnnotationMissionDebugDataSourceList,omitempty"`
	// Debug data in JSON format
	//
	// > This parameter serves the same purpose as AnnotationMissionDebugDataSourceList. You can use either one. The format is [1], [2], or [1,2]. You can include multiple filter conditions in the array. The corresponding enumeration values are as follows:
	//
	// - 1: Outbound job
	//
	// - 2: Testing job
	//
	// example:
	//
	// [1]
	AnnotationMissionDebugDataSourceListJsonString *string `json:"AnnotationMissionDebugDataSourceListJsonString,omitempty" xml:"AnnotationMissionDebugDataSourceListJsonString,omitempty"`
	// Task Name of the annotation Job
	//
	// example:
	//
	// 询问卖车-标注任务-20230506-112934
	AnnotationMissionName *string `json:"AnnotationMissionName,omitempty" xml:"AnnotationMissionName,omitempty"`
	// Chatbot ID
	//
	// > You can obtain it from the DescribeScript API.
	//
	// example:
	//
	// chatbot-cn-fqEnFZBYnb
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// End time for filtering conversations by time
	//
	// example:
	//
	// 1682956800000
	ConversationTimeEndFilter *int64 `json:"ConversationTimeEndFilter,omitempty" xml:"ConversationTimeEndFilter,omitempty"`
	// Start time for filtering conversations by time
	//
	// example:
	//
	// 1683302400000
	ConversationTimeStartFilter *int64 `json:"ConversationTimeStartFilter,omitempty" xml:"ConversationTimeStartFilter,omitempty"`
	// Specifies whether to exclude call records already annotated by other Jobs. The default value is false if this parameter is not specified.
	//
	// example:
	//
	// true
	ExcludeOtherSession *bool `json:"ExcludeOtherSession,omitempty" xml:"ExcludeOtherSession,omitempty"`
	// Whether the business ended normally.
	//
	// > This parameter takes effect only when SessionEndReasonFilterList is not provided.
	//
	// > - true: Normal conversation records
	//
	// > - false: Abnormal conversation records
	//
	// example:
	//
	// True
	Finished *bool `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// Instance ID
	//
	// example:
	//
	// bf3b51a5-e88a-4636-98b0-1a34725a085b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Custom sampling quantity
	//
	// > This parameter takes effect only when SamplingType is 3 and is Required in that case; otherwise, creation will fail.
	//
	// example:
	//
	// 1
	SamplingCount *int32 `json:"SamplingCount,omitempty" xml:"SamplingCount,omitempty"`
	// Sampling rate
	//
	// > This parameter takes effect only when SamplingType is 2 and is Required in that case; otherwise, creation will fail.
	//
	// example:
	//
	// 1
	SamplingRate *int32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// Sampling type.
	//
	// - 1: Full volume
	//
	// - 2: Percentage
	//
	// - 3: Custom quantity
	//
	// example:
	//
	// 1
	SamplingType *int32 `json:"SamplingType,omitempty" xml:"SamplingType,omitempty"`
	// Outbound scenario ID
	//
	// example:
	//
	// 6236f21e-2e04-4dad-a47b-ae77e6a48325
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Filter condition for call result status.
	//
	// > This parameter is equivalent to SessionEndReasonFilterListJsonString. You can specify either one of them.
	SessionEndReasonFilterList []*int32 `json:"SessionEndReasonFilterList,omitempty" xml:"SessionEndReasonFilterList,omitempty" type:"Repeated"`
	// Filter condition for call result status.
	//
	// > This parameter serves the same purpose as SessionEndReasonFilterList. You can specify either one of them.
	//
	// The format is [1] or [1,2], etc. The array can contain multiple filter conditions. The corresponding enumeration values are as follows.
	//
	// **Enumeration values for call record filtering**
	//
	// - 1: Normal completion
	//
	// - 2: Robot hangs up after rejection
	//
	// - 3: Hangs up due to silence timeout
	//
	// - 4: User hangs up after rejection
	//
	// - 5: User hangs up without reason
	//
	// - 6: Hit intent and transferred to agent
	//
	// - 7: Rejected and transferred to agent
	//
	// - 8: No interaction from user side
	//
	// - 9: System Exception interruption
	//
	// - 10: Hit intent and transferred to IVR
	//
	// - 11: Rejected and transferred to IVR
	//
	// example:
	//
	// [1]
	SessionEndReasonFilterListJsonString *string `json:"SessionEndReasonFilterListJsonString,omitempty" xml:"SessionEndReasonFilterListJsonString,omitempty"`
}

func (s CreateAnnotationMissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnotationMissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAnnotationMissionShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAnnotationMissionShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateAnnotationMissionShrinkRequest) GetAnnotationMissionDataSourceType() *int32 {
	return s.AnnotationMissionDataSourceType
}

func (s *CreateAnnotationMissionShrinkRequest) GetAnnotationMissionDebugDataSourceListShrink() *string {
	return s.AnnotationMissionDebugDataSourceListShrink
}

func (s *CreateAnnotationMissionShrinkRequest) GetAnnotationMissionDebugDataSourceListJsonString() *string {
	return s.AnnotationMissionDebugDataSourceListJsonString
}

func (s *CreateAnnotationMissionShrinkRequest) GetAnnotationMissionName() *string {
	return s.AnnotationMissionName
}

func (s *CreateAnnotationMissionShrinkRequest) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *CreateAnnotationMissionShrinkRequest) GetConversationTimeEndFilter() *int64 {
	return s.ConversationTimeEndFilter
}

func (s *CreateAnnotationMissionShrinkRequest) GetConversationTimeStartFilter() *int64 {
	return s.ConversationTimeStartFilter
}

func (s *CreateAnnotationMissionShrinkRequest) GetExcludeOtherSession() *bool {
	return s.ExcludeOtherSession
}

func (s *CreateAnnotationMissionShrinkRequest) GetFinished() *bool {
	return s.Finished
}

func (s *CreateAnnotationMissionShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAnnotationMissionShrinkRequest) GetSamplingCount() *int32 {
	return s.SamplingCount
}

func (s *CreateAnnotationMissionShrinkRequest) GetSamplingRate() *int32 {
	return s.SamplingRate
}

func (s *CreateAnnotationMissionShrinkRequest) GetSamplingType() *int32 {
	return s.SamplingType
}

func (s *CreateAnnotationMissionShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateAnnotationMissionShrinkRequest) GetSessionEndReasonFilterList() []*int32 {
	return s.SessionEndReasonFilterList
}

func (s *CreateAnnotationMissionShrinkRequest) GetSessionEndReasonFilterListJsonString() *string {
	return s.SessionEndReasonFilterListJsonString
}

func (s *CreateAnnotationMissionShrinkRequest) SetAgentId(v string) *CreateAnnotationMissionShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetAgentKey(v string) *CreateAnnotationMissionShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetAnnotationMissionDataSourceType(v int32) *CreateAnnotationMissionShrinkRequest {
	s.AnnotationMissionDataSourceType = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetAnnotationMissionDebugDataSourceListShrink(v string) *CreateAnnotationMissionShrinkRequest {
	s.AnnotationMissionDebugDataSourceListShrink = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetAnnotationMissionDebugDataSourceListJsonString(v string) *CreateAnnotationMissionShrinkRequest {
	s.AnnotationMissionDebugDataSourceListJsonString = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetAnnotationMissionName(v string) *CreateAnnotationMissionShrinkRequest {
	s.AnnotationMissionName = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetChatbotId(v string) *CreateAnnotationMissionShrinkRequest {
	s.ChatbotId = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetConversationTimeEndFilter(v int64) *CreateAnnotationMissionShrinkRequest {
	s.ConversationTimeEndFilter = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetConversationTimeStartFilter(v int64) *CreateAnnotationMissionShrinkRequest {
	s.ConversationTimeStartFilter = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetExcludeOtherSession(v bool) *CreateAnnotationMissionShrinkRequest {
	s.ExcludeOtherSession = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetFinished(v bool) *CreateAnnotationMissionShrinkRequest {
	s.Finished = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetInstanceId(v string) *CreateAnnotationMissionShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetSamplingCount(v int32) *CreateAnnotationMissionShrinkRequest {
	s.SamplingCount = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetSamplingRate(v int32) *CreateAnnotationMissionShrinkRequest {
	s.SamplingRate = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetSamplingType(v int32) *CreateAnnotationMissionShrinkRequest {
	s.SamplingType = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetScriptId(v string) *CreateAnnotationMissionShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetSessionEndReasonFilterList(v []*int32) *CreateAnnotationMissionShrinkRequest {
	s.SessionEndReasonFilterList = v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) SetSessionEndReasonFilterListJsonString(v string) *CreateAnnotationMissionShrinkRequest {
	s.SessionEndReasonFilterListJsonString = &v
	return s
}

func (s *CreateAnnotationMissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
