// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConversationResponseBody
	GetCode() *string
	SetConversationId(v string) *CreateConversationResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *CreateConversationResponseBody
	GetCreatedAt() *string
	SetMessage(v string) *CreateConversationResponseBody
	GetMessage() *string
	SetMetadata(v map[string]interface{}) *CreateConversationResponseBody
	GetMetadata() map[string]interface{}
	SetRequestId(v string) *CreateConversationResponseBody
	GetRequestId() *string
	SetTitle(v string) *CreateConversationResponseBody
	GetTitle() *string
}

type CreateConversationResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 新建会话ID
	//
	// example:
	//
	// exampleConversationId
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 会话创建时间戳（秒）
	//
	// example:
	//
	// 1
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 错误描述，成功时为空
	Message  *string                `json:"message,omitempty" xml:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 会话标题（已按调用方语言国际化）
	//
	// example:
	//
	// 示例标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConversationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConversationResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateConversationResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateConversationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConversationResponseBody) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *CreateConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConversationResponseBody) GetTitle() *string {
	return s.Title
}

func (s *CreateConversationResponseBody) SetCode(v string) *CreateConversationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConversationResponseBody) SetConversationId(v string) *CreateConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateConversationResponseBody) SetCreatedAt(v string) *CreateConversationResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateConversationResponseBody) SetMessage(v string) *CreateConversationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConversationResponseBody) SetMetadata(v map[string]interface{}) *CreateConversationResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateConversationResponseBody) SetRequestId(v string) *CreateConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConversationResponseBody) SetTitle(v string) *CreateConversationResponseBody {
	s.Title = &v
	return s
}

func (s *CreateConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
