// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMessageResponseBody
	GetRequestId() *string
	SetResult(v *ListMessageResponseBodyResult) *ListMessageResponseBody
	GetResult() *ListMessageResponseBodyResult
}

type ListMessageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *ListMessageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageResponseBody) GetResult() *ListMessageResponseBodyResult {
	return s.Result
}

func (s *ListMessageResponseBody) SetRequestId(v string) *ListMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageResponseBody) SetResult(v *ListMessageResponseBodyResult) *ListMessageResponseBody {
	s.Result = v
	return s
}

func (s *ListMessageResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessageResponseBodyResult struct {
	// Indicates whether the current page is followed by another page. Valid values:
	//
	// 	- true: The current page is followed by another page.
	//
	// 	- false: The current page is not followed by another page.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// Details about the messages.
	MessageList []*ListMessageResponseBodyResultMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
}

func (s ListMessageResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMessageResponseBodyResult) GetMessageList() []*ListMessageResponseBodyResultMessageList {
	return s.MessageList
}

func (s *ListMessageResponseBodyResult) SetHasMore(v bool) *ListMessageResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListMessageResponseBodyResult) SetMessageList(v []*ListMessageResponseBodyResultMessageList) *ListMessageResponseBodyResult {
	s.MessageList = v
	return s
}

func (s *ListMessageResponseBodyResult) Validate() error {
	if s.MessageList != nil {
		for _, item := range s.MessageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessageResponseBodyResultMessageList struct {
	// The message body. The value is a JSON string.
	//
	// example:
	//
	// test
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the message group.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// qt***
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The ID of the user who sent the message.
	//
	// example:
	//
	// yi***
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The type of the message.
	//
	// example:
	//
	// 10002
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMessageResponseBodyResultMessageList) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponseBodyResultMessageList) GoString() string {
	return s.String()
}

func (s *ListMessageResponseBodyResultMessageList) GetData() *string {
	return s.Data
}

func (s *ListMessageResponseBodyResultMessageList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListMessageResponseBodyResultMessageList) GetMessageId() *string {
	return s.MessageId
}

func (s *ListMessageResponseBodyResultMessageList) GetSenderId() *string {
	return s.SenderId
}

func (s *ListMessageResponseBodyResultMessageList) GetType() *int32 {
	return s.Type
}

func (s *ListMessageResponseBodyResultMessageList) SetData(v string) *ListMessageResponseBodyResultMessageList {
	s.Data = &v
	return s
}

func (s *ListMessageResponseBodyResultMessageList) SetGroupId(v string) *ListMessageResponseBodyResultMessageList {
	s.GroupId = &v
	return s
}

func (s *ListMessageResponseBodyResultMessageList) SetMessageId(v string) *ListMessageResponseBodyResultMessageList {
	s.MessageId = &v
	return s
}

func (s *ListMessageResponseBodyResultMessageList) SetSenderId(v string) *ListMessageResponseBodyResultMessageList {
	s.SenderId = &v
	return s
}

func (s *ListMessageResponseBodyResultMessageList) SetType(v int32) *ListMessageResponseBodyResultMessageList {
	s.Type = &v
	return s
}

func (s *ListMessageResponseBodyResultMessageList) Validate() error {
	return dara.Validate(s)
}
