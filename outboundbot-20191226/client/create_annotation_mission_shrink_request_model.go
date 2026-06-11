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
	// The agent ID.
	//
	// > You can obtain this ID from the \\`DescribeInstance\\` operation.
	//
	// example:
	//
	// 1168702
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The workspace key.
	//
	// > You can obtain this key from the \\`DescribeInstance\\` operation.
	//
	// example:
	//
	// 9137ab9c27044921860030adf8590ec4_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The data type of the annotation task.
	//
	// - 1: Outbound call
	//
	// example:
	//
	// 1
	AnnotationMissionDataSourceType *int32 `json:"AnnotationMissionDataSourceType,omitempty" xml:"AnnotationMissionDataSourceType,omitempty"`
	// The list of annotation data sources.
	//
	// > This parameter has the same function as \\`AnnotationMissionDebugDataSourceListJsonString\\`. You can specify either of them.
	AnnotationMissionDebugDataSourceListShrink *string `json:"AnnotationMissionDebugDataSourceList,omitempty" xml:"AnnotationMissionDebugDataSourceList,omitempty"`
	// The JSON string for the test data.
	//
	// > This parameter has the same function as \\`AnnotationMissionDebugDataSourceList\\`. You can specify either of them. The format is \\`[1]\\`, \\`[2]\\`, or \\`[1,2]\\`. You can specify multiple filter conditions in the array. The enumeration values are as follows:
	//
	// - 1: Outbound call task
	//
	// - 2: Test task
	//
	// example:
	//
	// [1]
	AnnotationMissionDebugDataSourceListJsonString *string `json:"AnnotationMissionDebugDataSourceListJsonString,omitempty" xml:"AnnotationMissionDebugDataSourceListJsonString,omitempty"`
	// The name of the annotation task.
	//
	// example:
	//
	// 询问卖车-标注任务-20230506-112934
	AnnotationMissionName *string `json:"AnnotationMissionName,omitempty" xml:"AnnotationMissionName,omitempty"`
	// The bot ID.
	//
	// > You can obtain this ID from the \\`DescribeScript\\` operation.
	//
	// example:
	//
	// chatbot-cn-fqEnFZBYnb
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// The end time for filtering calls.
	//
	// example:
	//
	// 1682956800000
	ConversationTimeEndFilter *int64 `json:"ConversationTimeEndFilter,omitempty" xml:"ConversationTimeEndFilter,omitempty"`
	// The start time for filtering calls.
	//
	// example:
	//
	// 1683302400000
	ConversationTimeStartFilter *int64 `json:"ConversationTimeStartFilter,omitempty" xml:"ConversationTimeStartFilter,omitempty"`
	// Specifies whether to exclude call records that have been annotated in other tasks. If you do not specify this parameter, the default value is \\`false\\`.
	//
	// example:
	//
	// true
	ExcludeOtherSession *bool `json:"ExcludeOtherSession,omitempty" xml:"ExcludeOtherSession,omitempty"`
	// Indicates whether the business process ended normally.
	//
	// > This parameter takes effect only when \\`SessionEndReasonFilterList\\` is not specified.
	//
	// >
	//
	// > - \\`true\\`: The call record is normal.
	//
	// >
	//
	// > - \\`false\\`: The call did not end normally.
	//
	// example:
	//
	// True
	Finished *bool `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// bf3b51a5-e88a-4636-98b0-1a34725a085b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The custom sampling amount.
	//
	// > This parameter is required and takes effect only when \\`SamplingType\\` is set to 3. Otherwise, the task fails to be created.
	//
	// example:
	//
	// 1
	SamplingCount *int32 `json:"SamplingCount,omitempty" xml:"SamplingCount,omitempty"`
	// The sampling percentage.
	//
	// > This parameter is required and takes effect only when \\`SamplingType\\` is set to 2. Otherwise, the task fails to be created.
	//
	// example:
	//
	// 1
	SamplingRate *int32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// The sampling type.
	//
	// - 1: Full data
	//
	// - 2: Percentage
	//
	// - 3: Custom amount
	//
	// example:
	//
	// 1
	SamplingType *int32 `json:"SamplingType,omitempty" xml:"SamplingType,omitempty"`
	// The outbound scenario ID.
	//
	// example:
	//
	// 6236f21e-2e04-4dad-a47b-ae77e6a48325
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The filter condition for call completion statuses.
	//
	// > This parameter has the same function as \\`SessionEndReasonFilterListJsonString\\`. You can specify either of them.
	SessionEndReasonFilterList []*int32 `json:"SessionEndReasonFilterList,omitempty" xml:"SessionEndReasonFilterList,omitempty" type:"Repeated"`
	// The filter condition for call completion statuses.
	//
	// > This parameter has the same function as \\`SessionEndReasonFilterList\\`. You can specify either of them. The format is \\`[1]\\` or \\`[1,2]\\`. You can specify multiple filter conditions in the array. The enumeration values for the filter conditions are as follows.
	//
	// **Enumeration values for filtering call records**
	//
	// - 1: The call ended normally.
	//
	// - 2: The bot hung up after a recognition failure.
	//
	// - 3: The call was hung up due to a silence timeout.
	//
	// - 4: The user hung up after a recognition failure.
	//
	// - 5: The user hung up for no reason.
	//
	// - 6: The call was transferred to a manual agent after an intent was hit.
	//
	// - 7: The call was transferred to a manual agent after a recognition failure.
	//
	// - 8: No interaction from the user.
	//
	// - 9: The call was interrupted by a system exception.
	//
	// - 10: The call was transferred to an IVR after an intent was hit.
	//
	// - 11: The call was transferred to an IVR after a recognition failure.
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
