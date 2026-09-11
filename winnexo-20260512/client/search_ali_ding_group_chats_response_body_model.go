// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAliDingGroupChatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchAliDingGroupChatsResponseBody
	GetCode() *string
	SetHasMore(v bool) *SearchAliDingGroupChatsResponseBody
	GetHasMore() *bool
	SetItems(v []*SearchAliDingGroupChatsResponseBodyItems) *SearchAliDingGroupChatsResponseBody
	GetItems() []*SearchAliDingGroupChatsResponseBodyItems
	SetMessage(v string) *SearchAliDingGroupChatsResponseBody
	GetMessage() *string
	SetNextCursor(v string) *SearchAliDingGroupChatsResponseBody
	GetNextCursor() *string
	SetRequestId(v string) *SearchAliDingGroupChatsResponseBody
	GetRequestId() *string
}

type SearchAliDingGroupChatsResponseBody struct {
	// 业务状态码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否还有下一页
	//
	// example:
	//
	// false
	HasMore *bool                                       `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items   []*SearchAliDingGroupChatsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 下一页分页游标，末页为空
	//
	// example:
	//
	// 1
	NextCursor *string `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// request-id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SearchAliDingGroupChatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAliDingGroupChatsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAliDingGroupChatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchAliDingGroupChatsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *SearchAliDingGroupChatsResponseBody) GetItems() []*SearchAliDingGroupChatsResponseBodyItems {
	return s.Items
}

func (s *SearchAliDingGroupChatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchAliDingGroupChatsResponseBody) GetNextCursor() *string {
	return s.NextCursor
}

func (s *SearchAliDingGroupChatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAliDingGroupChatsResponseBody) SetCode(v string) *SearchAliDingGroupChatsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBody) SetHasMore(v bool) *SearchAliDingGroupChatsResponseBody {
	s.HasMore = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBody) SetItems(v []*SearchAliDingGroupChatsResponseBodyItems) *SearchAliDingGroupChatsResponseBody {
	s.Items = v
	return s
}

func (s *SearchAliDingGroupChatsResponseBody) SetMessage(v string) *SearchAliDingGroupChatsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBody) SetNextCursor(v string) *SearchAliDingGroupChatsResponseBody {
	s.NextCursor = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBody) SetRequestId(v string) *SearchAliDingGroupChatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchAliDingGroupChatsResponseBodyItems struct {
	// 阿里钉群聊 ID
	//
	// example:
	//
	// cid-example
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// 会话类型
	//
	// example:
	//
	// INTERNAL_GROUP
	ConversationType *string `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	// 当前用户是否开启免打扰
	//
	// example:
	//
	// false
	Muted *bool `json:"muted,omitempty" xml:"muted,omitempty"`
	// 群聊标题
	//
	// example:
	//
	// 客户项目群
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SearchAliDingGroupChatsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s SearchAliDingGroupChatsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SearchAliDingGroupChatsResponseBodyItems) GetChatId() *string {
	return s.ChatId
}

func (s *SearchAliDingGroupChatsResponseBodyItems) GetConversationType() *string {
	return s.ConversationType
}

func (s *SearchAliDingGroupChatsResponseBodyItems) GetMuted() *bool {
	return s.Muted
}

func (s *SearchAliDingGroupChatsResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *SearchAliDingGroupChatsResponseBodyItems) SetChatId(v string) *SearchAliDingGroupChatsResponseBodyItems {
	s.ChatId = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBodyItems) SetConversationType(v string) *SearchAliDingGroupChatsResponseBodyItems {
	s.ConversationType = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBodyItems) SetMuted(v bool) *SearchAliDingGroupChatsResponseBodyItems {
	s.Muted = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBodyItems) SetTitle(v string) *SearchAliDingGroupChatsResponseBodyItems {
	s.Title = &v
	return s
}

func (s *SearchAliDingGroupChatsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
