// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChatResponseBody
	GetCode() *string
	SetContent(v []*ChatResponseBodyContent) *ChatResponseBody
	GetContent() []*ChatResponseBodyContent
	SetCreated(v string) *ChatResponseBody
	GetCreated() *string
	SetCreatedAt(v string) *ChatResponseBody
	GetCreatedAt() *string
	SetHttpStatusCode(v int32) *ChatResponseBody
	GetHttpStatusCode() *int32
	SetId(v string) *ChatResponseBody
	GetId() *string
	SetMessage(v string) *ChatResponseBody
	GetMessage() *string
	SetObject(v string) *ChatResponseBody
	GetObject() *string
	SetRequestId(v string) *ChatResponseBody
	GetRequestId() *string
	SetRole(v string) *ChatResponseBody
	GetRole() *string
	SetSequenceNumber(v string) *ChatResponseBody
	GetSequenceNumber() *string
	SetSessionId(v string) *ChatResponseBody
	GetSessionId() *string
	SetStatus(v string) *ChatResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ChatResponseBody
	GetSuccess() *bool
	SetText(v string) *ChatResponseBody
	GetText() *string
	SetType(v string) *ChatResponseBody
	GetType() *string
}

type ChatResponseBody struct {
	// example:
	//
	// 200
	Code    *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Content []*ChatResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 1773380609
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// example:
	//
	// 1773380609
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// msg_xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// response
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// example:
	//
	// EA12****-****-****-****-****E5C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 1
	SequenceNumber *string `json:"SequenceNumber,omitempty" xml:"SequenceNumber,omitempty"`
	// example:
	//
	// 176405663****961
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// in_progress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// "hello"
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// message / text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatResponseBody) GetContent() []*ChatResponseBodyContent {
	return s.Content
}

func (s *ChatResponseBody) GetCreated() *string {
	return s.Created
}

func (s *ChatResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ChatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ChatResponseBody) GetId() *string {
	return s.Id
}

func (s *ChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatResponseBody) GetObject() *string {
	return s.Object
}

func (s *ChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBody) GetRole() *string {
	return s.Role
}

func (s *ChatResponseBody) GetSequenceNumber() *string {
	return s.SequenceNumber
}

func (s *ChatResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ChatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatResponseBody) GetText() *string {
	return s.Text
}

func (s *ChatResponseBody) GetType() *string {
	return s.Type
}

func (s *ChatResponseBody) SetCode(v string) *ChatResponseBody {
	s.Code = &v
	return s
}

func (s *ChatResponseBody) SetContent(v []*ChatResponseBodyContent) *ChatResponseBody {
	s.Content = v
	return s
}

func (s *ChatResponseBody) SetCreated(v string) *ChatResponseBody {
	s.Created = &v
	return s
}

func (s *ChatResponseBody) SetCreatedAt(v string) *ChatResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *ChatResponseBody) SetHttpStatusCode(v int32) *ChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChatResponseBody) SetId(v string) *ChatResponseBody {
	s.Id = &v
	return s
}

func (s *ChatResponseBody) SetMessage(v string) *ChatResponseBody {
	s.Message = &v
	return s
}

func (s *ChatResponseBody) SetObject(v string) *ChatResponseBody {
	s.Object = &v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBody) SetRole(v string) *ChatResponseBody {
	s.Role = &v
	return s
}

func (s *ChatResponseBody) SetSequenceNumber(v string) *ChatResponseBody {
	s.SequenceNumber = &v
	return s
}

func (s *ChatResponseBody) SetSessionId(v string) *ChatResponseBody {
	s.SessionId = &v
	return s
}

func (s *ChatResponseBody) SetStatus(v string) *ChatResponseBody {
	s.Status = &v
	return s
}

func (s *ChatResponseBody) SetSuccess(v bool) *ChatResponseBody {
	s.Success = &v
	return s
}

func (s *ChatResponseBody) SetText(v string) *ChatResponseBody {
	s.Text = &v
	return s
}

func (s *ChatResponseBody) SetType(v string) *ChatResponseBody {
	s.Type = &v
	return s
}

func (s *ChatResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatResponseBodyContent struct {
	// example:
	//
	// {"call_id":"call_xxx","name":"get_weather","arguments":"{"city":"Beijing"}"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// ""
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// example:
	//
	// 您好
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ChatResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyContent) GetData() map[string]interface{} {
	return s.Data
}

func (s *ChatResponseBodyContent) GetObject() *string {
	return s.Object
}

func (s *ChatResponseBodyContent) GetText() *string {
	return s.Text
}

func (s *ChatResponseBodyContent) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyContent) SetData(v map[string]interface{}) *ChatResponseBodyContent {
	s.Data = v
	return s
}

func (s *ChatResponseBodyContent) SetObject(v string) *ChatResponseBodyContent {
	s.Object = &v
	return s
}

func (s *ChatResponseBodyContent) SetText(v string) *ChatResponseBodyContent {
	s.Text = &v
	return s
}

func (s *ChatResponseBodyContent) SetType(v string) *ChatResponseBodyContent {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
