// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetChatSessionResponseBody
	GetCode() *string
	SetMessage(v string) *GetChatSessionResponseBody
	GetMessage() *string
	SetMessages(v []*GetChatSessionResponseBodyMessages) *GetChatSessionResponseBody
	GetMessages() []*GetChatSessionResponseBodyMessages
	SetRequestId(v string) *GetChatSessionResponseBody
	GetRequestId() *string
	SetSession(v *GetChatSessionResponseBodySession) *GetChatSessionResponseBody
	GetSession() *GetChatSessionResponseBodySession
}

type GetChatSessionResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message  *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Messages []*GetChatSessionResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Session   *GetChatSessionResponseBodySession `json:"session,omitempty" xml:"session,omitempty" type:"Struct"`
}

func (s GetChatSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatSessionResponseBody) GetMessages() []*GetChatSessionResponseBodyMessages {
	return s.Messages
}

func (s *GetChatSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatSessionResponseBody) GetSession() *GetChatSessionResponseBodySession {
	return s.Session
}

func (s *GetChatSessionResponseBody) SetCode(v string) *GetChatSessionResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatSessionResponseBody) SetMessage(v string) *GetChatSessionResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatSessionResponseBody) SetMessages(v []*GetChatSessionResponseBodyMessages) *GetChatSessionResponseBody {
	s.Messages = v
	return s
}

func (s *GetChatSessionResponseBody) SetRequestId(v string) *GetChatSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatSessionResponseBody) SetSession(v *GetChatSessionResponseBodySession) *GetChatSessionResponseBody {
	s.Session = v
	return s
}

func (s *GetChatSessionResponseBody) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatSessionResponseBodyMessages struct {
	// 消息内容
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 该消息之后 LLM 上下文是否已清空
	//
	// example:
	//
	// true
	ContextCleared *bool `json:"contextCleared,omitempty" xml:"contextCleared,omitempty"`
	// 是否来自分享续聊复制的消息
	//
	// example:
	//
	// true
	FromShare *bool `json:"fromShare,omitempty" xml:"fromShare,omitempty"`
	// 消息ID
	//
	// example:
	//
	// exampleId
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 1
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// 类型
	//
	// example:
	//
	// string_value
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// 角色
	//
	// example:
	//
	// string_value
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 分享来源用户名称
	//
	// example:
	//
	// string_value
	ShareUserName *string `json:"shareUserName,omitempty" xml:"shareUserName,omitempty"`
	// 消息状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 追踪ID
	//
	// example:
	//
	// exampleTraceId
	TraceId  *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
	UpdateAt *int64  `json:"updateAt,omitempty" xml:"updateAt,omitempty"`
	// 用户反馈类型
	//
	// example:
	//
	// string_value
	UserFeedback *string `json:"userFeedback,omitempty" xml:"userFeedback,omitempty"`
}

func (s GetChatSessionResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *GetChatSessionResponseBodyMessages) GetContent() *string {
	return s.Content
}

func (s *GetChatSessionResponseBodyMessages) GetContextCleared() *bool {
	return s.ContextCleared
}

func (s *GetChatSessionResponseBodyMessages) GetFromShare() *bool {
	return s.FromShare
}

func (s *GetChatSessionResponseBodyMessages) GetId() *string {
	return s.Id
}

func (s *GetChatSessionResponseBodyMessages) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetChatSessionResponseBodyMessages) GetObject() *string {
	return s.Object
}

func (s *GetChatSessionResponseBodyMessages) GetRole() *string {
	return s.Role
}

func (s *GetChatSessionResponseBodyMessages) GetShareUserName() *string {
	return s.ShareUserName
}

func (s *GetChatSessionResponseBodyMessages) GetStatus() *string {
	return s.Status
}

func (s *GetChatSessionResponseBodyMessages) GetTraceId() *string {
	return s.TraceId
}

func (s *GetChatSessionResponseBodyMessages) GetUpdateAt() *int64 {
	return s.UpdateAt
}

func (s *GetChatSessionResponseBodyMessages) GetUserFeedback() *string {
	return s.UserFeedback
}

func (s *GetChatSessionResponseBodyMessages) SetContent(v string) *GetChatSessionResponseBodyMessages {
	s.Content = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetContextCleared(v bool) *GetChatSessionResponseBodyMessages {
	s.ContextCleared = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetFromShare(v bool) *GetChatSessionResponseBodyMessages {
	s.FromShare = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetId(v string) *GetChatSessionResponseBodyMessages {
	s.Id = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetMetadata(v map[string]interface{}) *GetChatSessionResponseBodyMessages {
	s.Metadata = v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetObject(v string) *GetChatSessionResponseBodyMessages {
	s.Object = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetRole(v string) *GetChatSessionResponseBodyMessages {
	s.Role = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetShareUserName(v string) *GetChatSessionResponseBodyMessages {
	s.ShareUserName = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetStatus(v string) *GetChatSessionResponseBodyMessages {
	s.Status = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetTraceId(v string) *GetChatSessionResponseBodyMessages {
	s.TraceId = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetUpdateAt(v int64) *GetChatSessionResponseBodyMessages {
	s.UpdateAt = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) SetUserFeedback(v string) *GetChatSessionResponseBodyMessages {
	s.UserFeedback = &v
	return s
}

func (s *GetChatSessionResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}

type GetChatSessionResponseBodySession struct {
	// 创建时间
	//
	// example:
	//
	// 1
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 消息ID
	//
	// example:
	//
	// exampleId
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 创建时间是否超过30天
	//
	// example:
	//
	// true
	IsExpired *bool `json:"isExpired,omitempty" xml:"isExpired,omitempty"`
	// 关联对象ID
	//
	// example:
	//
	// exampleObjectId
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// 会话使用的抽象模型名（quick/standard/flagship）
	//
	// example:
	//
	// string_value
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 类型
	//
	// example:
	//
	// string_value
	Object   *string `json:"object,omitempty" xml:"object,omitempty"`
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// operatingObjectName
	//
	// example:
	//
	// string_value
	OperatingObjectName []*string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty" type:"Repeated"`
	// 标题
	//
	// example:
	//
	// 示例标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 1
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetChatSessionResponseBodySession) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionResponseBodySession) GoString() string {
	return s.String()
}

func (s *GetChatSessionResponseBodySession) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetChatSessionResponseBodySession) GetId() *string {
	return s.Id
}

func (s *GetChatSessionResponseBodySession) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *GetChatSessionResponseBodySession) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetChatSessionResponseBodySession) GetModel() *string {
	return s.Model
}

func (s *GetChatSessionResponseBodySession) GetObject() *string {
	return s.Object
}

func (s *GetChatSessionResponseBodySession) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetChatSessionResponseBodySession) GetOperatingObjectName() []*string {
	return s.OperatingObjectName
}

func (s *GetChatSessionResponseBodySession) GetTitle() *string {
	return s.Title
}

func (s *GetChatSessionResponseBodySession) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetChatSessionResponseBodySession) SetCreatedAt(v int64) *GetChatSessionResponseBodySession {
	s.CreatedAt = &v
	return s
}

func (s *GetChatSessionResponseBodySession) SetId(v string) *GetChatSessionResponseBodySession {
	s.Id = &v
	return s
}

func (s *GetChatSessionResponseBodySession) SetIsExpired(v bool) *GetChatSessionResponseBodySession {
	s.IsExpired = &v
	return s
}

func (s *GetChatSessionResponseBodySession) SetMetadata(v map[string]interface{}) *GetChatSessionResponseBodySession {
	s.Metadata = v
	return s
}

func (s *GetChatSessionResponseBodySession) SetModel(v string) *GetChatSessionResponseBodySession {
	s.Model = &v
	return s
}

func (s *GetChatSessionResponseBodySession) SetObject(v string) *GetChatSessionResponseBodySession {
	s.Object = &v
	return s
}

func (s *GetChatSessionResponseBodySession) SetObjectId(v string) *GetChatSessionResponseBodySession {
	s.ObjectId = &v
	return s
}

func (s *GetChatSessionResponseBodySession) SetOperatingObjectName(v []*string) *GetChatSessionResponseBodySession {
	s.OperatingObjectName = v
	return s
}

func (s *GetChatSessionResponseBodySession) SetTitle(v string) *GetChatSessionResponseBodySession {
	s.Title = &v
	return s
}

func (s *GetChatSessionResponseBodySession) SetUpdatedAt(v int64) *GetChatSessionResponseBodySession {
	s.UpdatedAt = &v
	return s
}

func (s *GetChatSessionResponseBodySession) Validate() error {
	return dara.Validate(s)
}
