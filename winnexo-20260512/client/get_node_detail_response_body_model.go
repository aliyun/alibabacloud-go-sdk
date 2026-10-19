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
	// 按思考、技能、工具、回复和产物分组的实际工作过程
	//
	// example:
	//
	// {"thoughts":[],"skillCalls":[],"toolCalls":[],"replies":[],"artifacts":[]}
	AgentWorkProcess *GetNodeDetailResponseBodyAgentWorkProcess `json:"agentWorkProcess,omitempty" xml:"agentWorkProcess,omitempty" type:"Struct"`
	// 已保存的节点产物摘要与稳定引用，不包含inline正文；无产物时为空列表
	//
	// example:
	//
	// []
	Artifacts []*GetNodeDetailResponseBodyArtifacts `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Repeated"`
	// 节点执行尝试序号
	//
	// example:
	//
	// 1
	AttemptSeq *int64 `json:"attemptSeq,omitempty" xml:"attemptSeq,omitempty"`
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 节点会话ID，未创建时为空
	//
	// example:
	//
	// conversation_example
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 当前节点会话是否可以继续输入
	//
	// example:
	//
	// true
	ConversationWritable *bool `json:"conversationWritable,omitempty" xml:"conversationWritable,omitempty"`
	// 节点负责人展示名称，未知时为空
	//
	// example:
	//
	// 数字员工
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 节点 key（nodeKey）
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 节点展示名称，与节点列表同源
	//
	// example:
	//
	// 查询签单金额
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// 节点类型，返回实际枚举
	//
	// example:
	//
	// AGENT
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 节点结果摘要，不含inline正文
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// 是否有待派发续轮
	//
	// example:
	//
	// false
	ResumePending *bool `json:"resumePending,omitempty" xml:"resumePending,omitempty"`
	// 节点AI工作记录时间线，与GetAgentWorkLog同源，duration单位毫秒；无记录时为空列表
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
	// 节点运行状态，返回实际枚举
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 当前节点关联的待办ID列表，元素为待办ID字符串
	//
	// example:
	//
	// []
	TodoIds []*string `json:"todoIds,omitempty" xml:"todoIds,omitempty" type:"Repeated"`
	// 当前尝试的对话轮序号
	//
	// example:
	//
	// 2
	TurnSeq *int64 `json:"turnSeq,omitempty" xml:"turnSeq,omitempty"`
	// 节点等待原因，未等待时为空
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
	// 节点产物摘要与稳定引用，与顶层artifacts一致；无产物时为空列表
	//
	// example:
	//
	// []
	Artifacts []*GetNodeDetailResponseBodyAgentWorkProcessArtifacts `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Repeated"`
	// 已保存的回复记录，action为REPLY；无记录时为空列表
	//
	// example:
	//
	// []
	Replies []*GetNodeDetailResponseBodyAgentWorkProcessReplies `json:"replies,omitempty" xml:"replies,omitempty" type:"Repeated"`
	// 已保存的技能调用记录，action为SKILL_CALL；无记录时为空列表
	//
	// example:
	//
	// []
	SkillCalls []*GetNodeDetailResponseBodyAgentWorkProcessSkillCalls `json:"skillCalls,omitempty" xml:"skillCalls,omitempty" type:"Repeated"`
	// 已保存的思考记录，action为THINK；无记录时为空列表
	//
	// example:
	//
	// []
	Thoughts []*GetNodeDetailResponseBodyAgentWorkProcessThoughts `json:"thoughts,omitempty" xml:"thoughts,omitempty" type:"Repeated"`
	// 已保存的工具调用记录，action为TOOL_CALL；无记录时为空列表
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
	// 登记产物ID；仅来自结果摘要的产物可能为空
	//
	// example:
	//
	// artifact_example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// 产物在节点结果中的唯一key
	//
	// example:
	//
	// amount_report
	ArtifactKey *string `json:"artifactKey,omitempty" xml:"artifactKey,omitempty"`
	// 产物稳定引用，包含kind、id、url；登记产物URL为重新鉴权入口，可为空
	//
	// example:
	//
	// {"kind":"INLINE","id":"amount_report","url":"inline://amount_report"}
	Ref map[string]*string `json:"ref,omitempty" xml:"ref,omitempty"`
	// 产物摘要，不含inline正文
	//
	// example:
	//
	// 已完成签单金额统计
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 产物展示标题
	//
	// example:
	//
	// 签单金额统计报告
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 产物类型
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
	// 记录类型：THINK、REPLY、TOOL_CALL、SKILL_CALL或ARTIFACT
	//
	// example:
	//
	// REPLY
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 已保存的原始记录内容；工具调用与产物元数据为JSON字符串
	//
	// example:
	//
	// 已完成本轮数据查询
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 该条工作记录耗时，单位毫秒；产物记录为0
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 所属节点key（nodeKey），不是会话ID
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 记录事件时间，ISO-8601格式；源记录无时间时为空字符串
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
	// 记录类型：THINK、REPLY、TOOL_CALL、SKILL_CALL或ARTIFACT
	//
	// example:
	//
	// SKILL_CALL
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 已保存的原始记录内容；工具调用与产物元数据为JSON字符串
	//
	// example:
	//
	// 已完成本轮数据查询
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 该条工作记录耗时，单位毫秒；产物记录为0
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 所属节点key（nodeKey），不是会话ID
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 记录事件时间，ISO-8601格式；源记录无时间时为空字符串
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
	// 记录类型：THINK、REPLY、TOOL_CALL、SKILL_CALL或ARTIFACT
	//
	// example:
	//
	// THINK
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 已保存的原始记录内容；工具调用与产物元数据为JSON字符串
	//
	// example:
	//
	// 已完成本轮数据查询
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 该条工作记录耗时，单位毫秒；产物记录为0
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 所属节点key（nodeKey），不是会话ID
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 记录事件时间，ISO-8601格式；源记录无时间时为空字符串
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
	// 记录类型：THINK、REPLY、TOOL_CALL、SKILL_CALL或ARTIFACT
	//
	// example:
	//
	// TOOL_CALL
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 已保存的原始记录内容；工具调用与产物元数据为JSON字符串
	//
	// example:
	//
	// 已完成本轮数据查询
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 该条工作记录耗时，单位毫秒；产物记录为0
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 所属节点key（nodeKey），不是会话ID
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 记录事件时间，ISO-8601格式；源记录无时间时为空字符串
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
	// 登记产物ID；仅来自结果摘要的产物可能为空
	//
	// example:
	//
	// artifact_example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// 产物在节点结果中的唯一key
	//
	// example:
	//
	// amount_report
	ArtifactKey *string `json:"artifactKey,omitempty" xml:"artifactKey,omitempty"`
	// 产物稳定引用，包含kind、id、url；登记产物URL为重新鉴权入口，可为空
	//
	// example:
	//
	// {"kind":"INLINE","id":"amount_report","url":"inline://amount_report"}
	Ref map[string]*string `json:"ref,omitempty" xml:"ref,omitempty"`
	// 产物摘要，不含inline正文
	//
	// example:
	//
	// 已完成签单金额统计
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 产物展示标题
	//
	// example:
	//
	// 签单金额统计报告
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 产物类型
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
	// 记录类型：THINK、REPLY、TOOL_CALL、SKILL_CALL或ARTIFACT
	//
	// example:
	//
	// REPLY
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 已保存的原始记录内容；工具调用与产物元数据为JSON字符串
	//
	// example:
	//
	// 已完成本轮数据查询
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 该条工作记录耗时，单位毫秒；产物记录为0
	//
	// example:
	//
	// 1200
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 所属节点key（nodeKey），不是会话ID
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 记录事件时间，ISO-8601格式；源记录无时间时为空字符串
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
