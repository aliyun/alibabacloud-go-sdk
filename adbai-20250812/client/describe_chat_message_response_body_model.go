// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChatMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeChatMessageResponseBody
	GetCode() *string
	SetCompletedAt(v string) *DescribeChatMessageResponseBody
	GetCompletedAt() *string
	SetContent(v []*DescribeChatMessageResponseBodyContent) *DescribeChatMessageResponseBody
	GetContent() []*DescribeChatMessageResponseBodyContent
	SetCreatedAt(v string) *DescribeChatMessageResponseBody
	GetCreatedAt() *string
	SetData(v *DescribeChatMessageResponseBodyData) *DescribeChatMessageResponseBody
	GetData() *DescribeChatMessageResponseBodyData
	SetDelta(v bool) *DescribeChatMessageResponseBody
	GetDelta() *bool
	SetError(v string) *DescribeChatMessageResponseBody
	GetError() *string
	SetId(v string) *DescribeChatMessageResponseBody
	GetId() *string
	SetIndex(v int64) *DescribeChatMessageResponseBody
	GetIndex() *int64
	SetMessage(v string) *DescribeChatMessageResponseBody
	GetMessage() *string
	SetMsgId(v string) *DescribeChatMessageResponseBody
	GetMsgId() *string
	SetObject(v string) *DescribeChatMessageResponseBody
	GetObject() *string
	SetOutput(v string) *DescribeChatMessageResponseBody
	GetOutput() *string
	SetRequestId(v string) *DescribeChatMessageResponseBody
	GetRequestId() *string
	SetRole(v string) *DescribeChatMessageResponseBody
	GetRole() *string
	SetSequenceNumber(v int64) *DescribeChatMessageResponseBody
	GetSequenceNumber() *int64
	SetSessionId(v string) *DescribeChatMessageResponseBody
	GetSessionId() *string
	SetStatus(v string) *DescribeChatMessageResponseBody
	GetStatus() *string
	SetText(v string) *DescribeChatMessageResponseBody
	GetText() *string
	SetType(v string) *DescribeChatMessageResponseBody
	GetType() *string
}

type DescribeChatMessageResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2026-03-18T02:01:56Z
	CompletedAt *string                                   `json:"CompletedAt,omitempty" xml:"CompletedAt,omitempty"`
	Content     []*DescribeChatMessageResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-03-18T02:01:56Z
	CreatedAt *string                              `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Data      *DescribeChatMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// false
	Delta *bool `json:"Delta,omitempty" xml:"Delta,omitempty"`
	// example:
	//
	// Failed to get sse streaming, please try again later.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 97616
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// msg_294c8b98-dc64-4c82-9417-e03522a631f6
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// example:
	//
	// content
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// example:
	//
	// [{"type": "text", "text": "Skill not found: instance_health_inspection"}]
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 1
	SequenceNumber *int64 `json:"SequenceNumber,omitempty" xml:"SequenceNumber,omitempty"`
	// example:
	//
	// 593b51eef93b443fb2ba2a6dc68b737b
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ：-636 KB
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChatMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChatMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeChatMessageResponseBody) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *DescribeChatMessageResponseBody) GetContent() []*DescribeChatMessageResponseBodyContent {
	return s.Content
}

func (s *DescribeChatMessageResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeChatMessageResponseBody) GetData() *DescribeChatMessageResponseBodyData {
	return s.Data
}

func (s *DescribeChatMessageResponseBody) GetDelta() *bool {
	return s.Delta
}

func (s *DescribeChatMessageResponseBody) GetError() *string {
	return s.Error
}

func (s *DescribeChatMessageResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeChatMessageResponseBody) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeChatMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeChatMessageResponseBody) GetMsgId() *string {
	return s.MsgId
}

func (s *DescribeChatMessageResponseBody) GetObject() *string {
	return s.Object
}

func (s *DescribeChatMessageResponseBody) GetOutput() *string {
	return s.Output
}

func (s *DescribeChatMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChatMessageResponseBody) GetRole() *string {
	return s.Role
}

func (s *DescribeChatMessageResponseBody) GetSequenceNumber() *int64 {
	return s.SequenceNumber
}

func (s *DescribeChatMessageResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeChatMessageResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeChatMessageResponseBody) GetText() *string {
	return s.Text
}

func (s *DescribeChatMessageResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeChatMessageResponseBody) SetCode(v string) *DescribeChatMessageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetCompletedAt(v string) *DescribeChatMessageResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetContent(v []*DescribeChatMessageResponseBodyContent) *DescribeChatMessageResponseBody {
	s.Content = v
	return s
}

func (s *DescribeChatMessageResponseBody) SetCreatedAt(v string) *DescribeChatMessageResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetData(v *DescribeChatMessageResponseBodyData) *DescribeChatMessageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeChatMessageResponseBody) SetDelta(v bool) *DescribeChatMessageResponseBody {
	s.Delta = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetError(v string) *DescribeChatMessageResponseBody {
	s.Error = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetId(v string) *DescribeChatMessageResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetIndex(v int64) *DescribeChatMessageResponseBody {
	s.Index = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetMessage(v string) *DescribeChatMessageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetMsgId(v string) *DescribeChatMessageResponseBody {
	s.MsgId = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetObject(v string) *DescribeChatMessageResponseBody {
	s.Object = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetOutput(v string) *DescribeChatMessageResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetRequestId(v string) *DescribeChatMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetRole(v string) *DescribeChatMessageResponseBody {
	s.Role = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetSequenceNumber(v int64) *DescribeChatMessageResponseBody {
	s.SequenceNumber = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetSessionId(v string) *DescribeChatMessageResponseBody {
	s.SessionId = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetStatus(v string) *DescribeChatMessageResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetText(v string) *DescribeChatMessageResponseBody {
	s.Text = &v
	return s
}

func (s *DescribeChatMessageResponseBody) SetType(v string) *DescribeChatMessageResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeChatMessageResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeChatMessageResponseBodyContent struct {
	Data *DescribeChatMessageResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// false
	Delta *bool `json:"Delta,omitempty" xml:"Delta,omitempty"`
	// example:
	//
	// Failed to get sse streaming, please try again later.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 0
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// msg_294c8b98-dc64-4c82-9417-e03522a631f6
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// example:
	//
	// content
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// example:
	//
	// 1
	SequenceNumber *int64 `json:"SequenceNumber,omitempty" xml:"SequenceNumber,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3 MB\\n-
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChatMessageResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeChatMessageResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeChatMessageResponseBodyContent) GetData() *DescribeChatMessageResponseBodyContentData {
	return s.Data
}

func (s *DescribeChatMessageResponseBodyContent) GetDelta() *bool {
	return s.Delta
}

func (s *DescribeChatMessageResponseBodyContent) GetError() *string {
	return s.Error
}

func (s *DescribeChatMessageResponseBodyContent) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeChatMessageResponseBodyContent) GetMsgId() *string {
	return s.MsgId
}

func (s *DescribeChatMessageResponseBodyContent) GetObject() *string {
	return s.Object
}

func (s *DescribeChatMessageResponseBodyContent) GetSequenceNumber() *int64 {
	return s.SequenceNumber
}

func (s *DescribeChatMessageResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *DescribeChatMessageResponseBodyContent) GetText() *string {
	return s.Text
}

func (s *DescribeChatMessageResponseBodyContent) GetType() *string {
	return s.Type
}

func (s *DescribeChatMessageResponseBodyContent) SetData(v *DescribeChatMessageResponseBodyContentData) *DescribeChatMessageResponseBodyContent {
	s.Data = v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetDelta(v bool) *DescribeChatMessageResponseBodyContent {
	s.Delta = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetError(v string) *DescribeChatMessageResponseBodyContent {
	s.Error = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetIndex(v int64) *DescribeChatMessageResponseBodyContent {
	s.Index = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetMsgId(v string) *DescribeChatMessageResponseBodyContent {
	s.MsgId = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetObject(v string) *DescribeChatMessageResponseBodyContent {
	s.Object = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetSequenceNumber(v int64) *DescribeChatMessageResponseBodyContent {
	s.SequenceNumber = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetStatus(v string) *DescribeChatMessageResponseBodyContent {
	s.Status = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetText(v string) *DescribeChatMessageResponseBodyContent {
	s.Text = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) SetType(v string) *DescribeChatMessageResponseBodyContent {
	s.Type = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeChatMessageResponseBodyContentData struct {
	// example:
	//
	// call_e0e9ee423c7e4ba88d058fd6
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// describeScalingRecommendation
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// [{"type": "text", "text": "Skill not found: instance_health_inspection"}]
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s DescribeChatMessageResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s DescribeChatMessageResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *DescribeChatMessageResponseBodyContentData) GetCallId() *string {
	return s.CallId
}

func (s *DescribeChatMessageResponseBodyContentData) GetName() *string {
	return s.Name
}

func (s *DescribeChatMessageResponseBodyContentData) GetOutput() *string {
	return s.Output
}

func (s *DescribeChatMessageResponseBodyContentData) SetCallId(v string) *DescribeChatMessageResponseBodyContentData {
	s.CallId = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContentData) SetName(v string) *DescribeChatMessageResponseBodyContentData {
	s.Name = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContentData) SetOutput(v string) *DescribeChatMessageResponseBodyContentData {
	s.Output = &v
	return s
}

func (s *DescribeChatMessageResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}

type DescribeChatMessageResponseBodyData struct {
	// example:
	//
	// {"instanceName": "amv-bp1b9y9xhvpzm9p0", "pageNumber": 1, "pageSize": 10}
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// example:
	//
	// call_1891f1689bc44ab292aadff3
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// describeScalingRecommendation
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// [{"type": "text", "text": "Skill not found: instance_health_inspection"}]
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s DescribeChatMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeChatMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeChatMessageResponseBodyData) GetArguments() *string {
	return s.Arguments
}

func (s *DescribeChatMessageResponseBodyData) GetCallId() *string {
	return s.CallId
}

func (s *DescribeChatMessageResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeChatMessageResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *DescribeChatMessageResponseBodyData) SetArguments(v string) *DescribeChatMessageResponseBodyData {
	s.Arguments = &v
	return s
}

func (s *DescribeChatMessageResponseBodyData) SetCallId(v string) *DescribeChatMessageResponseBodyData {
	s.CallId = &v
	return s
}

func (s *DescribeChatMessageResponseBodyData) SetName(v string) *DescribeChatMessageResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeChatMessageResponseBodyData) SetOutput(v string) *DescribeChatMessageResponseBodyData {
	s.Output = &v
	return s
}

func (s *DescribeChatMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
