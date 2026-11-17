// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertDirectConversationToWorkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConvertDirectConversationToWorkResponseBody
	GetCode() *string
	SetConversionId(v string) *ConvertDirectConversationToWorkResponseBody
	GetConversionId() *string
	SetConvertedRounds(v int32) *ConvertDirectConversationToWorkResponseBody
	GetConvertedRounds() *int32
	SetHistoryScope(v string) *ConvertDirectConversationToWorkResponseBody
	GetHistoryScope() *string
	SetMessage(v string) *ConvertDirectConversationToWorkResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConvertDirectConversationToWorkResponseBody
	GetRequestId() *string
	SetRetryable(v bool) *ConvertDirectConversationToWorkResponseBody
	GetRetryable() *bool
	SetSourceEpoch(v int64) *ConvertDirectConversationToWorkResponseBody
	GetSourceEpoch() *int64
	SetSourceSessionId(v string) *ConvertDirectConversationToWorkResponseBody
	GetSourceSessionId() *string
	SetStatus(v string) *ConvertDirectConversationToWorkResponseBody
	GetStatus() *string
	SetTargetSessionId(v string) *ConvertDirectConversationToWorkResponseBody
	GetTargetSessionId() *string
	SetThroughSequence(v int64) *ConvertDirectConversationToWorkResponseBody
	GetThroughSequence() *int64
	SetWorkMode(v string) *ConvertDirectConversationToWorkResponseBody
	GetWorkMode() *string
}

type ConvertDirectConversationToWorkResponseBody struct {
	// 业务状态码：成功为 200，失败沿现有 OpenAPI/POP 错误合同返回；转换是否整体完成须同时检查 status
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 转换 ID，查询及恢复过程中保持不变
	//
	// example:
	//
	// exampleConversionId
	ConversionId *string `json:"conversionId,omitempty" xml:"conversionId,omitempty"`
	// 固定快照中转换的完整 direct 问答轮数，不受发送接口 directMaxHistoryRounds 限制
	//
	// example:
	//
	// 2
	ConvertedRounds *int32 `json:"convertedRounds,omitempty" xml:"convertedRounds,omitempty"`
	// 转换历史范围，固定为 direct；不包含普通 work 历史
	//
	// example:
	//
	// direct
	HistoryScope *string `json:"historyScope,omitempty" xml:"historyScope,omitempty"`
	// 错误描述，成功时为空
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 是否允许继续原转换；false 表示不另起执行，仍允许用同一源 sessionId 查询。重试或继续始终复用同一转换、目标及快照
	//
	// example:
	//
	// false
	Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
	// 冻结快照对应的源 direct 历史 epoch
	//
	// example:
	//
	// 0
	SourceEpoch *int64 `json:"sourceEpoch,omitempty" xml:"sourceEpoch,omitempty"`
	// 源会话 ID，与请求 sessionId 一致
	//
	// example:
	//
	// exampleSourceSessionId
	SourceSessionId *string `json:"sourceSessionId,omitempty" xml:"sourceSessionId,omitempty"`
	// 转换状态：processing 表示处理中；succeeded 表示模型导入、目标登记、展示投影及两侧就绪均完成。失败按 OpenAPI/POP 错误合同处理，不能仅凭 HTTP 200 判断转换成功
	//
	// example:
	//
	// succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 新的 work 目标会话 ID，与源 ID 不同；processing 时可能返回预分配 ID，仅 succeeded 后可使用；重复查询或恢复复用原目标
	//
	// example:
	//
	// exampleTargetSessionId
	TargetSessionId *string `json:"targetSessionId,omitempty" xml:"targetSessionId,omitempty"`
	// 冻结快照包含的最后一个 direct 轮次序号；源会话后续轮次不自动加入目标
	//
	// example:
	//
	// 2
	ThroughSequence *int64 `json:"throughSequence,omitempty" xml:"throughSequence,omitempty"`
	// 目标会话的工作模式，固定为 work
	//
	// example:
	//
	// work
	WorkMode *string `json:"workMode,omitempty" xml:"workMode,omitempty"`
}

func (s ConvertDirectConversationToWorkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertDirectConversationToWorkResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertDirectConversationToWorkResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConvertDirectConversationToWorkResponseBody) GetConversionId() *string {
	return s.ConversionId
}

func (s *ConvertDirectConversationToWorkResponseBody) GetConvertedRounds() *int32 {
	return s.ConvertedRounds
}

func (s *ConvertDirectConversationToWorkResponseBody) GetHistoryScope() *string {
	return s.HistoryScope
}

func (s *ConvertDirectConversationToWorkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConvertDirectConversationToWorkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertDirectConversationToWorkResponseBody) GetRetryable() *bool {
	return s.Retryable
}

func (s *ConvertDirectConversationToWorkResponseBody) GetSourceEpoch() *int64 {
	return s.SourceEpoch
}

func (s *ConvertDirectConversationToWorkResponseBody) GetSourceSessionId() *string {
	return s.SourceSessionId
}

func (s *ConvertDirectConversationToWorkResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ConvertDirectConversationToWorkResponseBody) GetTargetSessionId() *string {
	return s.TargetSessionId
}

func (s *ConvertDirectConversationToWorkResponseBody) GetThroughSequence() *int64 {
	return s.ThroughSequence
}

func (s *ConvertDirectConversationToWorkResponseBody) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ConvertDirectConversationToWorkResponseBody) SetCode(v string) *ConvertDirectConversationToWorkResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetConversionId(v string) *ConvertDirectConversationToWorkResponseBody {
	s.ConversionId = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetConvertedRounds(v int32) *ConvertDirectConversationToWorkResponseBody {
	s.ConvertedRounds = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetHistoryScope(v string) *ConvertDirectConversationToWorkResponseBody {
	s.HistoryScope = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetMessage(v string) *ConvertDirectConversationToWorkResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetRequestId(v string) *ConvertDirectConversationToWorkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetRetryable(v bool) *ConvertDirectConversationToWorkResponseBody {
	s.Retryable = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetSourceEpoch(v int64) *ConvertDirectConversationToWorkResponseBody {
	s.SourceEpoch = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetSourceSessionId(v string) *ConvertDirectConversationToWorkResponseBody {
	s.SourceSessionId = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetStatus(v string) *ConvertDirectConversationToWorkResponseBody {
	s.Status = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetTargetSessionId(v string) *ConvertDirectConversationToWorkResponseBody {
	s.TargetSessionId = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetThroughSequence(v int64) *ConvertDirectConversationToWorkResponseBody {
	s.ThroughSequence = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) SetWorkMode(v string) *ConvertDirectConversationToWorkResponseBody {
	s.WorkMode = &v
	return s
}

func (s *ConvertDirectConversationToWorkResponseBody) Validate() error {
	return dara.Validate(s)
}
