// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownStreamBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDownStreamBindingsResponseBody
	GetCode() *int32
	SetData(v *ListDownStreamBindingsResponseBodyData) *ListDownStreamBindingsResponseBody
	GetData() *ListDownStreamBindingsResponseBodyData
	SetMessage(v string) *ListDownStreamBindingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDownStreamBindingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDownStreamBindingsResponseBody
	GetSuccess() *bool
}

type ListDownStreamBindingsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListDownStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C1E0502-0790-4FDB-8C96-6D5C8D9B7***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDownStreamBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDownStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDownStreamBindingsResponseBody) GetData() *ListDownStreamBindingsResponseBodyData {
	return s.Data
}

func (s *ListDownStreamBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDownStreamBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDownStreamBindingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDownStreamBindingsResponseBody) SetCode(v int32) *ListDownStreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetData(v *ListDownStreamBindingsResponseBodyData) *ListDownStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetMessage(v string) *ListDownStreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetRequestId(v string) *ListDownStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetSuccess(v bool) *ListDownStreamBindingsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDownStreamBindingsResponseBodyData struct {
	// The bindings.
	Bindings []*ListDownStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDownStreamBindingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDownStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBodyData) GetBindings() []*ListDownStreamBindingsResponseBodyDataBindings {
	return s.Bindings
}

func (s *ListDownStreamBindingsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDownStreamBindingsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDownStreamBindingsResponseBodyData) SetBindings(v []*ListDownStreamBindingsResponseBodyDataBindings) *ListDownStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) SetMaxResults(v int32) *ListDownStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) SetNextToken(v string) *ListDownStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDownStreamBindingsResponseBodyDataBindings struct {
	// The x-match attribute. Valid values:
	//
	// 	- **all:*	- A headers exchange routes a message to a queue only if all binding attributes of the queue except for x-match match the headers attributes of the message. This value is the default value.
	//
	// 	- **any:*	- A headers exchange routes a message to a queue if one or more binding attributes of the queue except for x-match match the headers attributes of the message.
	//
	// This parameter is available only for headers exchanges.
	//
	// example:
	//
	// test
	Argument *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	// The binding key.
	//
	// 	- If the source exchange is not a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// 	- If the source exchange is a topic exchange, the binding key must meet the following conventions:
	//
	//     	- The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//     	- The binding key cannot start or end with a period (.). If a binding key starts with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be followed by a period (.). If the binding key ends with a number sign (#) or an asterisk (\\*), the number sign (#) or asterisk (\\*) must be preceded by a period (.). If a number sign (#) or an asterisk (\\*) is used in the middle of a binding key, the number sign (#) or asterisk (\\*) must be preceded and followed by a period (.).
	//
	//     	- The binding key must be 1 to 255 characters in length.
	//
	// example:
	//
	// amq.test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the object to which the source exchange is bound. Valid values:
	//
	// 	- **QUEUE**
	//
	// 	- **EXCHANGE**
	//
	// example:
	//
	// QUEUE
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the object to which the source exchange is bound.
	//
	// example:
	//
	// QueueTest
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	// The name of the source exchange.
	//
	// example:
	//
	// test
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListDownStreamBindingsResponseBodyDataBindings) String() string {
	return dara.Prettify(s)
}

func (s ListDownStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) GetArgument() *string {
	return s.Argument
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) GetBindingKey() *string {
	return s.BindingKey
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) GetBindingType() *string {
	return s.BindingType
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) GetDestinationName() *string {
	return s.DestinationName
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) GetSourceExchange() *string {
	return s.SourceExchange
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) Validate() error {
	return dara.Validate(s)
}
