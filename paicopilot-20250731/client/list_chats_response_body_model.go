// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChats(v []*Chat) *ListChatsResponseBody
	GetChats() []*Chat
	SetRequestId(v string) *ListChatsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListChatsResponseBody
	GetTotalCount() *int64
}

type ListChatsResponseBody struct {
	Chats []*Chat `json:"Chats,omitempty" xml:"Chats,omitempty" type:"Repeated"`
	// example:
	//
	// 44553E9A-******-37ADC33FE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatsResponseBody) GetChats() []*Chat {
	return s.Chats
}

func (s *ListChatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListChatsResponseBody) SetChats(v []*Chat) *ListChatsResponseBody {
	s.Chats = v
	return s
}

func (s *ListChatsResponseBody) SetRequestId(v string) *ListChatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatsResponseBody) SetTotalCount(v int64) *ListChatsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChatsResponseBody) Validate() error {
	if s.Chats != nil {
		for _, item := range s.Chats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
