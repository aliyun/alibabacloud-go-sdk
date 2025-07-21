// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListBindingsResponseBodyData) *ListBindingsResponseBody
	GetData() *ListBindingsResponseBodyData
	SetRequestId(v string) *ListBindingsResponseBody
	GetRequestId() *string
}

type ListBindingsResponseBody struct {
	// The returned data.
	Data *ListBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E0A71208-3E87-4732-81CC-B18E0B4B1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBody) GetData() *ListBindingsResponseBodyData {
	return s.Data
}

func (s *ListBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBindingsResponseBody) SetData(v *ListBindingsResponseBodyData) *ListBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListBindingsResponseBody) SetRequestId(v string) *ListBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindingsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBindingsResponseBodyData struct {
	// The bindings.
	Bindings []*ListBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
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

func (s ListBindingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyData) GetBindings() []*ListBindingsResponseBodyDataBindings {
	return s.Bindings
}

func (s *ListBindingsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBindingsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBindingsResponseBodyData) SetBindings(v []*ListBindingsResponseBodyDataBindings) *ListBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListBindingsResponseBodyData) SetMaxResults(v int32) *ListBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsResponseBodyData) SetNextToken(v string) *ListBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListBindingsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListBindingsResponseBodyDataBindings struct {
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
	// all
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
	//     	- The binding key can contain letters, digits, hyphens (-), underscores (_), asterisks (\\*), periods (.), number signs (#), forward slashes (/), and at signs (@).
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

func (s ListBindingsResponseBodyDataBindings) String() string {
	return dara.Prettify(s)
}

func (s ListBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyDataBindings) GetArgument() *string {
	return s.Argument
}

func (s *ListBindingsResponseBodyDataBindings) GetBindingKey() *string {
	return s.BindingKey
}

func (s *ListBindingsResponseBodyDataBindings) GetBindingType() *string {
	return s.BindingType
}

func (s *ListBindingsResponseBodyDataBindings) GetDestinationName() *string {
	return s.DestinationName
}

func (s *ListBindingsResponseBodyDataBindings) GetSourceExchange() *string {
	return s.SourceExchange
}

func (s *ListBindingsResponseBodyDataBindings) SetArgument(v string) *ListBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetBindingType(v string) *ListBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) Validate() error {
	return dara.Validate(s)
}
