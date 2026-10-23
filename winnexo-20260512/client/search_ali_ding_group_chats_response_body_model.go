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
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether more pages are available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The user group tree node information.
	Items []*SearchAliDingGroupChatsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The starting cursor for the next page.
	//
	// example:
	//
	// 1
	NextCursor *string `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// The request ID.
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
	// The conversation ID. This ID is typically used for JSSDK.
	//
	// example:
	//
	// cid-example
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The conversation type. Valid values:
	//
	// - **0**: one-on-one chat.
	//
	// - **1**: group chat.
	//
	// > For a one-on-one chat, you do not need to specify **openConversationId**. Specify the employee ID in **receiverUserIdList**. For a one-on-one chat between two users, specify the employee ID of either user. For a one-on-one chat between a user and a bot, specify the employee ID of the user.
	//
	// example:
	//
	// INTERNAL_GROUP
	ConversationType *string `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	// Indicates whether the current user has enabled the Do Not Disturb mode.
	//
	// example:
	//
	// false
	Muted *bool `json:"muted,omitempty" xml:"muted,omitempty"`
	// The title of the scheduled meeting.
	//
	// example:
	//
	// Customer project group
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
