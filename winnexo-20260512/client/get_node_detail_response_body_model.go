// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentWorkProcess(v *GetNodeDetailResponseBodyAgentWorkProcess) *GetNodeDetailResponseBody
	GetAgentWorkProcess() *GetNodeDetailResponseBodyAgentWorkProcess
	SetArtifacts(v []*GetNodeDetailResponseBodyArtifacts) *GetNodeDetailResponseBody
	GetArtifacts() []*GetNodeDetailResponseBodyArtifacts
	SetAttemptSeq(v int64) *GetNodeDetailResponseBody
	GetAttemptSeq() *int64
	SetCode(v string) *GetNodeDetailResponseBody
	GetCode() *string
	SetConversationId(v string) *GetNodeDetailResponseBody
	GetConversationId() *string
	SetConversationWritable(v bool) *GetNodeDetailResponseBody
	GetConversationWritable() *bool
	SetHandler(v string) *GetNodeDetailResponseBody
	GetHandler() *string
	SetMessage(v string) *GetNodeDetailResponseBody
	GetMessage() *string
	SetNodeId(v string) *GetNodeDetailResponseBody
	GetNodeId() *string
	SetNodeName(v string) *GetNodeDetailResponseBody
	GetNodeName() *string
	SetNodeType(v string) *GetNodeDetailResponseBody
	GetNodeType() *string
	SetRequestId(v string) *GetNodeDetailResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *GetNodeDetailResponseBody
	GetResult() map[string]interface{}
	SetResumePending(v bool) *GetNodeDetailResponseBody
	GetResumePending() *bool
	SetRunLog(v []*GetNodeDetailResponseBodyRunLog) *GetNodeDetailResponseBody
	GetRunLog() []*GetNodeDetailResponseBodyRunLog
	SetStateVersion(v int64) *GetNodeDetailResponseBody
	GetStateVersion() *int64
	SetStatus(v string) *GetNodeDetailResponseBody
	GetStatus() *string
	SetTodoIds(v []*string) *GetNodeDetailResponseBody
	GetTodoIds() []*string
	SetTurnSeq(v int64) *GetNodeDetailResponseBody
	GetTurnSeq() *int64
	SetWaitType(v string) *GetNodeDetailResponseBody
	GetWaitType() *string
}

type GetNodeDetailResponseBody struct {
	// The actual work process grouped by thoughts, skills, tools, replies, and artifacts.
	//
	// example:
	//
	// {"thoughts":[],"skillCalls":[],"toolCalls":[],"replies":[],"artifacts":[]}
	AgentWorkProcess *GetNodeDetailResponseBodyAgentWorkProcess `json:"agentWorkProcess,omitempty" xml:"agentWorkProcess,omitempty" type:"Struct"`
	// The artifacts: saved node work records and artifacts.
	//
	// example:
	//
	// []
	Artifacts []*GetNodeDetailResponseBodyArtifacts `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Repeated"`
	// The sequence number of the node execution attempt.
	//
	// example:
	//
	// 1
	AttemptSeq *int64 `json:"attemptSeq,omitempty" xml:"attemptSeq,omitempty"`
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The node conversation ID. This parameter is empty if no conversation is created.
	//
	// example:
	//
	// conversation_example
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// Indicates whether the current node conversation can continue to receive input.
	//
	// example:
	//
	// true
	ConversationWritable *bool `json:"conversationWritable,omitempty" xml:"conversationWritable,omitempty"`
	// The display name of the node owner. This parameter is empty if the owner is unknown.
	//
	// example:
	//
	// Digital worker
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The node key.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The display name of the node, which shares the same source as the node list.
	//
	// example:
	//
	// Query signed amount
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// The node type. The actual enum value is returned.
	//
	// example:
	//
	// AGENT
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The node result summary, excluding inline content.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// Indicates whether there are pending follow-up turns to be dispatched.
	//
	// example:
	//
	// false
	ResumePending *bool `json:"resumePending,omitempty" xml:"resumePending,omitempty"`
	// The run log: saved node work records and artifacts.
	//
	// example:
	//
	// []
	RunLog []*GetNodeDetailResponseBodyRunLog `json:"runLog,omitempty" xml:"runLog,omitempty" type:"Repeated"`
	// stateVersion
	//
	// example:
	//
	// 2
	StateVersion *int64 `json:"stateVersion,omitempty" xml:"stateVersion,omitempty"`
	// The node running status. The actual enum value is returned.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of to-do IDs associated with the current node.
	//
	// example:
	//
	// []
	TodoIds []*string `json:"todoIds,omitempty" xml:"todoIds,omitempty" type:"Repeated"`
	// The sequence number of the conversation turn in the current attempt.
	//
	// example:
	//
	// 2
	TurnSeq *int64 `json:"turnSeq,omitempty" xml:"turnSeq,omitempty"`
	// The reason why the node is waiting. This parameter is empty if the node is not waiting.
	//
	// example:
	//
	// HUMAN_INPUT
	WaitType *string `json:"waitType,omitempty" xml:"waitType,omitempty"`
}

func (s GetNodeDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBody) GetAgentWorkProcess() *GetNodeDetailResponseBodyAgentWorkProcess {
	return s.AgentWorkProcess
}

func (s *GetNodeDetailResponseBody) GetArtifacts() []*GetNodeDetailResponseBodyArtifacts {
	return s.Artifacts
}

func (s *GetNodeDetailResponseBody) GetAttemptSeq() *int64 {
	return s.AttemptSeq
}

func (s *GetNodeDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNodeDetailResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetNodeDetailResponseBody) GetConversationWritable() *bool {
	return s.ConversationWritable
}

func (s *GetNodeDetailResponseBody) GetHandler() *string {
	return s.Handler
}

func (s *GetNodeDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNodeDetailResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeDetailResponseBody) GetNodeName() *string {
	return s.NodeName
}

func (s *GetNodeDetailResponseBody) GetNodeType() *string {
	return s.NodeType
}

func (s *GetNodeDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeDetailResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *GetNodeDetailResponseBody) GetResumePending() *bool {
	return s.ResumePending
}

func (s *GetNodeDetailResponseBody) GetRunLog() []*GetNodeDetailResponseBodyRunLog {
	return s.RunLog
}

func (s *GetNodeDetailResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *GetNodeDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetNodeDetailResponseBody) GetTodoIds() []*string {
	return s.TodoIds
}

func (s *GetNodeDetailResponseBody) GetTurnSeq() *int64 {
	return s.TurnSeq
}

func (s *GetNodeDetailResponseBody) GetWaitType() *string {
	return s.WaitType
}

func (s *GetNodeDetailResponseBody) SetAgentWorkProcess(v *GetNodeDetailResponseBodyAgentWorkProcess) *GetNodeDetailResponseBody {
	s.AgentWorkProcess = v
	return s
}

func (s *GetNodeDetailResponseBody) SetArtifacts(v []*GetNodeDetailResponseBodyArtifacts) *GetNodeDetailResponseBody {
	s.Artifacts = v
	return s
}

func (s *GetNodeDetailResponseBody) SetAttemptSeq(v int64) *GetNodeDetailResponseBody {
	s.AttemptSeq = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetCode(v string) *GetNodeDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetConversationId(v string) *GetNodeDetailResponseBody {
	s.ConversationId = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetConversationWritable(v bool) *GetNodeDetailResponseBody {
	s.ConversationWritable = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetHandler(v string) *GetNodeDetailResponseBody {
	s.Handler = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetMessage(v string) *GetNodeDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetNodeId(v string) *GetNodeDetailResponseBody {
	s.NodeId = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetNodeName(v string) *GetNodeDetailResponseBody {
	s.NodeName = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetNodeType(v string) *GetNodeDetailResponseBody {
	s.NodeType = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetRequestId(v string) *GetNodeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetResult(v map[string]interface{}) *GetNodeDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetNodeDetailResponseBody) SetResumePending(v bool) *GetNodeDetailResponseBody {
	s.ResumePending = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetRunLog(v []*GetNodeDetailResponseBodyRunLog) *GetNodeDetailResponseBody {
	s.RunLog = v
	return s
}

func (s *GetNodeDetailResponseBody) SetStateVersion(v int64) *GetNodeDetailResponseBody {
	s.StateVersion = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetStatus(v string) *GetNodeDetailResponseBody {
	s.Status = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetTodoIds(v []*string) *GetNodeDetailResponseBody {
	s.TodoIds = v
	return s
}

func (s *GetNodeDetailResponseBody) SetTurnSeq(v int64) *GetNodeDetailResponseBody {
	s.TurnSeq = &v
	return s
}

func (s *GetNodeDetailResponseBody) SetWaitType(v string) *GetNodeDetailResponseBody {
	s.WaitType = &v
	return s
}

func (s *GetNodeDetailResponseBody) Validate() error {
	if s.AgentWorkProcess != nil {
		if err := s.AgentWorkProcess.Validate(); err != nil {
			return err
		}
	}
	if s.Artifacts != nil {
		for _, item := range s.Artifacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RunLog != nil {
		for _, item := range s.RunLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNodeDetailResponseBodyAgentWorkProcess struct {
	// The artifacts: saved node work records and artifacts.
	//
	// example:
	//
	// []
	Artifacts []*GetNodeDetailResponseBodyAgentWorkProcessArtifacts `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Repeated"`
	// The replies: saved node work records and artifacts.
	//
	// example:
	//
	// []
	Replies []*GetNodeDetailResponseBodyAgentWorkProcessReplies `json:"replies,omitempty" xml:"replies,omitempty" type:"Repeated"`
	// The skill calls: saved node work records and artifacts.
	//
	// example:
	//
	// []
	SkillCalls []*GetNodeDetailResponseBodyAgentWorkProcessSkillCalls `json:"skillCalls,omitempty" xml:"skillCalls,omitempty" type:"Repeated"`
	// The thoughts: saved node work records and artifacts.
	//
	// example:
	//
	// []
	Thoughts []*GetNodeDetailResponseBodyAgentWorkProcessThoughts `json:"thoughts,omitempty" xml:"thoughts,omitempty" type:"Repeated"`
	// The tool calls: saved node work records and artifacts.
	//
	// example:
	//
	// []
	ToolCalls []*GetNodeDetailResponseBodyAgentWorkProcessToolCalls `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s GetNodeDetailResponseBodyAgentWorkProcess) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBodyAgentWorkProcess) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) GetArtifacts() []*GetNodeDetailResponseBodyAgentWorkProcessArtifacts {
	return s.Artifacts
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) GetReplies() []*GetNodeDetailResponseBodyAgentWorkProcessReplies {
	return s.Replies
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) GetSkillCalls() []*GetNodeDetailResponseBodyAgentWorkProcessSkillCalls {
	return s.SkillCalls
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) GetThoughts() []*GetNodeDetailResponseBodyAgentWorkProcessThoughts {
	return s.Thoughts
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) GetToolCalls() []*GetNodeDetailResponseBodyAgentWorkProcessToolCalls {
	return s.ToolCalls
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) SetArtifacts(v []*GetNodeDetailResponseBodyAgentWorkProcessArtifacts) *GetNodeDetailResponseBodyAgentWorkProcess {
	s.Artifacts = v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) SetReplies(v []*GetNodeDetailResponseBodyAgentWorkProcessReplies) *GetNodeDetailResponseBodyAgentWorkProcess {
	s.Replies = v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) SetSkillCalls(v []*GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) *GetNodeDetailResponseBodyAgentWorkProcess {
	s.SkillCalls = v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) SetThoughts(v []*GetNodeDetailResponseBodyAgentWorkProcessThoughts) *GetNodeDetailResponseBodyAgentWorkProcess {
	s.Thoughts = v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) SetToolCalls(v []*GetNodeDetailResponseBodyAgentWorkProcessToolCalls) *GetNodeDetailResponseBodyAgentWorkProcess {
	s.ToolCalls = v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcess) Validate() error {
	if s.Artifacts != nil {
		for _, item := range s.Artifacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Replies != nil {
		for _, item := range s.Replies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SkillCalls != nil {
		for _, item := range s.SkillCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Thoughts != nil {
		for _, item := range s.Thoughts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ToolCalls != nil {
		for _, item := range s.ToolCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNodeDetailResponseBodyAgentWorkProcessArtifacts struct {
	// The registered artifact ID. This parameter may be empty for artifacts derived only from the result summary.
	//
	// example:
	//
	// artifact_example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// The unique key of the artifact in the node result.
	//
	// example:
	//
	// amount_report
	ArtifactKey *string `json:"artifactKey,omitempty" xml:"artifactKey,omitempty"`
	// The stable reference of the artifact, including kind, id, and url. The URL of a registered artifact is a re-authentication entry point and can be empty.
	//
	// example:
	//
	// {"kind":"INLINE","id":"amount_report","url":"inline://amount_report"}
	Ref map[string]*string `json:"ref,omitempty" xml:"ref,omitempty"`
	// The artifact summary, excluding inline content.
	//
	// example:
	//
	// Signed amount statistics completed
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The display title of the artifact.
	//
	// example:
	//
	// Signed amount statistics report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The artifact type.
	//
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetNodeDetailResponseBodyAgentWorkProcessArtifacts) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBodyAgentWorkProcessArtifacts) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) GetArtifactKey() *string {
	return s.ArtifactKey
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) GetRef() map[string]*string {
	return s.Ref
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) GetSummary() *string {
	return s.Summary
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) GetTitle() *string {
	return s.Title
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) GetType() *string {
	return s.Type
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) SetArtifactId(v string) *GetNodeDetailResponseBodyAgentWorkProcessArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) SetArtifactKey(v string) *GetNodeDetailResponseBodyAgentWorkProcessArtifacts {
	s.ArtifactKey = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) SetRef(v map[string]*string) *GetNodeDetailResponseBodyAgentWorkProcessArtifacts {
	s.Ref = v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) SetSummary(v string) *GetNodeDetailResponseBodyAgentWorkProcessArtifacts {
	s.Summary = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) SetTitle(v string) *GetNodeDetailResponseBodyAgentWorkProcessArtifacts {
	s.Title = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) SetType(v string) *GetNodeDetailResponseBodyAgentWorkProcessArtifacts {
	s.Type = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessArtifacts) Validate() error {
	return dara.Validate(s)
}

type GetNodeDetailResponseBodyAgentWorkProcessReplies struct {
	// The record type. Valid values: THINK, REPLY, TOOL_CALL, SKILL_CALL, or ARTIFACT.
	//
	// example:
	//
	// REPLY
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The saved original record content. Tool invocation and artifact metadata are JSON strings.
	//
	// example:
	//
	// Data query for this turn completed
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// The time taken for this work record, in milliseconds. The value is 0 for artifact records.
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The key of the parent node, not the conversation ID.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The time when the event is recorded, in ISO-8601 format. This is an empty string if the source record has no time.
	//
	// example:
	//
	// 2026-10-14T08:00:00Z
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetNodeDetailResponseBodyAgentWorkProcessReplies) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBodyAgentWorkProcessReplies) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) GetAction() *string {
	return s.Action
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) GetDetail() *string {
	return s.Detail
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) GetDuration() *int64 {
	return s.Duration
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) SetAction(v string) *GetNodeDetailResponseBodyAgentWorkProcessReplies {
	s.Action = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) SetDetail(v string) *GetNodeDetailResponseBodyAgentWorkProcessReplies {
	s.Detail = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) SetDuration(v int64) *GetNodeDetailResponseBodyAgentWorkProcessReplies {
	s.Duration = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) SetNodeId(v string) *GetNodeDetailResponseBodyAgentWorkProcessReplies {
	s.NodeId = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) SetTimestamp(v string) *GetNodeDetailResponseBodyAgentWorkProcessReplies {
	s.Timestamp = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessReplies) Validate() error {
	return dara.Validate(s)
}

type GetNodeDetailResponseBodyAgentWorkProcessSkillCalls struct {
	// The record type. Valid values: THINK, REPLY, TOOL_CALL, SKILL_CALL, or ARTIFACT.
	//
	// example:
	//
	// SKILL_CALL
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The saved original record content. Tool invocation and artifact metadata are JSON strings.
	//
	// example:
	//
	// Data query for this turn completed
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// The time taken for this work record, in milliseconds. The value is 0 for artifact records.
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The key of the parent node, not the conversation ID.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The time when the event is recorded, in ISO-8601 format. This is an empty string if the source record has no time.
	//
	// example:
	//
	// 2026-10-14T08:00:00Z
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) GetAction() *string {
	return s.Action
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) GetDetail() *string {
	return s.Detail
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) GetDuration() *int64 {
	return s.Duration
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) SetAction(v string) *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls {
	s.Action = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) SetDetail(v string) *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls {
	s.Detail = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) SetDuration(v int64) *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls {
	s.Duration = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) SetNodeId(v string) *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls {
	s.NodeId = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) SetTimestamp(v string) *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls {
	s.Timestamp = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessSkillCalls) Validate() error {
	return dara.Validate(s)
}

type GetNodeDetailResponseBodyAgentWorkProcessThoughts struct {
	// The record type. Valid values: THINK, REPLY, TOOL_CALL, SKILL_CALL, or ARTIFACT.
	//
	// example:
	//
	// THINK
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The saved original record content. Tool invocation and artifact metadata are JSON strings.
	//
	// example:
	//
	// Data query for this turn completed
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// The time taken for this work record, in milliseconds. The value is 0 for artifact records.
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The key of the parent node, not the conversation ID.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The time when the event is recorded, in ISO-8601 format. This is an empty string if the source record has no time.
	//
	// example:
	//
	// 2026-10-14T08:00:00Z
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetNodeDetailResponseBodyAgentWorkProcessThoughts) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBodyAgentWorkProcessThoughts) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) GetAction() *string {
	return s.Action
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) GetDetail() *string {
	return s.Detail
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) GetDuration() *int64 {
	return s.Duration
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) SetAction(v string) *GetNodeDetailResponseBodyAgentWorkProcessThoughts {
	s.Action = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) SetDetail(v string) *GetNodeDetailResponseBodyAgentWorkProcessThoughts {
	s.Detail = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) SetDuration(v int64) *GetNodeDetailResponseBodyAgentWorkProcessThoughts {
	s.Duration = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) SetNodeId(v string) *GetNodeDetailResponseBodyAgentWorkProcessThoughts {
	s.NodeId = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) SetTimestamp(v string) *GetNodeDetailResponseBodyAgentWorkProcessThoughts {
	s.Timestamp = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessThoughts) Validate() error {
	return dara.Validate(s)
}

type GetNodeDetailResponseBodyAgentWorkProcessToolCalls struct {
	// The record type. Valid values: THINK, REPLY, TOOL_CALL, SKILL_CALL, or ARTIFACT.
	//
	// example:
	//
	// TOOL_CALL
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The saved original record content. Tool invocation and artifact metadata are JSON strings.
	//
	// example:
	//
	// Data query for this turn completed
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// The time taken for this work record, in milliseconds. The value is 0 for artifact records.
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The key of the parent node, not the conversation ID.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The time when the event is recorded, in ISO-8601 format. This is an empty string if the source record has no time.
	//
	// example:
	//
	// 2026-10-14T08:00:00Z
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetNodeDetailResponseBodyAgentWorkProcessToolCalls) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBodyAgentWorkProcessToolCalls) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) GetAction() *string {
	return s.Action
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) GetDetail() *string {
	return s.Detail
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) GetDuration() *int64 {
	return s.Duration
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) SetAction(v string) *GetNodeDetailResponseBodyAgentWorkProcessToolCalls {
	s.Action = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) SetDetail(v string) *GetNodeDetailResponseBodyAgentWorkProcessToolCalls {
	s.Detail = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) SetDuration(v int64) *GetNodeDetailResponseBodyAgentWorkProcessToolCalls {
	s.Duration = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) SetNodeId(v string) *GetNodeDetailResponseBodyAgentWorkProcessToolCalls {
	s.NodeId = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) SetTimestamp(v string) *GetNodeDetailResponseBodyAgentWorkProcessToolCalls {
	s.Timestamp = &v
	return s
}

func (s *GetNodeDetailResponseBodyAgentWorkProcessToolCalls) Validate() error {
	return dara.Validate(s)
}

type GetNodeDetailResponseBodyArtifacts struct {
	// The registered artifact ID. This parameter may be empty for artifacts derived only from the result summary.
	//
	// example:
	//
	// artifact_example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// The unique key of the artifact in the node result.
	//
	// example:
	//
	// amount_report
	ArtifactKey *string `json:"artifactKey,omitempty" xml:"artifactKey,omitempty"`
	// The stable reference of the artifact, including kind, id, and url. The URL of a registered artifact is a re-authentication entry point and can be empty.
	//
	// example:
	//
	// {"kind":"INLINE","id":"amount_report","url":"inline://amount_report"}
	Ref map[string]*string `json:"ref,omitempty" xml:"ref,omitempty"`
	// The artifact summary, excluding inline content.
	//
	// example:
	//
	// Signed amount statistics completed
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The display title of the artifact.
	//
	// example:
	//
	// Signed amount statistics report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The artifact type.
	//
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetNodeDetailResponseBodyArtifacts) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBodyArtifacts) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *GetNodeDetailResponseBodyArtifacts) GetArtifactKey() *string {
	return s.ArtifactKey
}

func (s *GetNodeDetailResponseBodyArtifacts) GetRef() map[string]*string {
	return s.Ref
}

func (s *GetNodeDetailResponseBodyArtifacts) GetSummary() *string {
	return s.Summary
}

func (s *GetNodeDetailResponseBodyArtifacts) GetTitle() *string {
	return s.Title
}

func (s *GetNodeDetailResponseBodyArtifacts) GetType() *string {
	return s.Type
}

func (s *GetNodeDetailResponseBodyArtifacts) SetArtifactId(v string) *GetNodeDetailResponseBodyArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *GetNodeDetailResponseBodyArtifacts) SetArtifactKey(v string) *GetNodeDetailResponseBodyArtifacts {
	s.ArtifactKey = &v
	return s
}

func (s *GetNodeDetailResponseBodyArtifacts) SetRef(v map[string]*string) *GetNodeDetailResponseBodyArtifacts {
	s.Ref = v
	return s
}

func (s *GetNodeDetailResponseBodyArtifacts) SetSummary(v string) *GetNodeDetailResponseBodyArtifacts {
	s.Summary = &v
	return s
}

func (s *GetNodeDetailResponseBodyArtifacts) SetTitle(v string) *GetNodeDetailResponseBodyArtifacts {
	s.Title = &v
	return s
}

func (s *GetNodeDetailResponseBodyArtifacts) SetType(v string) *GetNodeDetailResponseBodyArtifacts {
	s.Type = &v
	return s
}

func (s *GetNodeDetailResponseBodyArtifacts) Validate() error {
	return dara.Validate(s)
}

type GetNodeDetailResponseBodyRunLog struct {
	// The record type. Valid values: THINK, REPLY, TOOL_CALL, SKILL_CALL, or ARTIFACT.
	//
	// example:
	//
	// REPLY
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The saved original record content. Tool invocation and artifact metadata are JSON strings.
	//
	// example:
	//
	// Data query for this turn completed
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// The time taken for this work record, in milliseconds. The value is 0 for artifact records.
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The key of the parent node, not the conversation ID.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The time when the event is recorded, in ISO-8601 format. This is an empty string if the source record has no time.
	//
	// example:
	//
	// 2026-10-14T08:00:00Z
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetNodeDetailResponseBodyRunLog) String() string {
	return dara.Prettify(s)
}

func (s GetNodeDetailResponseBodyRunLog) GoString() string {
	return s.String()
}

func (s *GetNodeDetailResponseBodyRunLog) GetAction() *string {
	return s.Action
}

func (s *GetNodeDetailResponseBodyRunLog) GetDetail() *string {
	return s.Detail
}

func (s *GetNodeDetailResponseBodyRunLog) GetDuration() *int64 {
	return s.Duration
}

func (s *GetNodeDetailResponseBodyRunLog) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNodeDetailResponseBodyRunLog) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetNodeDetailResponseBodyRunLog) SetAction(v string) *GetNodeDetailResponseBodyRunLog {
	s.Action = &v
	return s
}

func (s *GetNodeDetailResponseBodyRunLog) SetDetail(v string) *GetNodeDetailResponseBodyRunLog {
	s.Detail = &v
	return s
}

func (s *GetNodeDetailResponseBodyRunLog) SetDuration(v int64) *GetNodeDetailResponseBodyRunLog {
	s.Duration = &v
	return s
}

func (s *GetNodeDetailResponseBodyRunLog) SetNodeId(v string) *GetNodeDetailResponseBodyRunLog {
	s.NodeId = &v
	return s
}

func (s *GetNodeDetailResponseBodyRunLog) SetTimestamp(v string) *GetNodeDetailResponseBodyRunLog {
	s.Timestamp = &v
	return s
}

func (s *GetNodeDetailResponseBodyRunLog) Validate() error {
	return dara.Validate(s)
}
