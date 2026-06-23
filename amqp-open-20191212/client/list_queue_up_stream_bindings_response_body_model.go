// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueUpStreamBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListQueueUpStreamBindingsResponseBodyData) *ListQueueUpStreamBindingsResponseBody
	GetData() *ListQueueUpStreamBindingsResponseBodyData
	SetRequestId(v string) *ListQueueUpStreamBindingsResponseBody
	GetRequestId() *string
}

type ListQueueUpStreamBindingsResponseBody struct {
	// The returned data.
	Data *ListQueueUpStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8BFB1C9D-08A2-4859-A47C-403C9EFA2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBody) GetData() *ListQueueUpStreamBindingsResponseBodyData {
	return s.Data
}

func (s *ListQueueUpStreamBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueueUpStreamBindingsResponseBody) SetData(v *ListQueueUpStreamBindingsResponseBodyData) *ListQueueUpStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBody) SetRequestId(v string) *ListQueueUpStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueueUpStreamBindingsResponseBodyData struct {
	// The bindings.
	Bindings []*ListQueueUpStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// The maximum number of results returned.
	//
	// example:
	//
	// 1
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current results. An empty value indicates that all results have been returned.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBodyData) GetBindings() []*ListQueueUpStreamBindingsResponseBodyDataBindings {
	return s.Bindings
}

func (s *ListQueueUpStreamBindingsResponseBodyData) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListQueueUpStreamBindingsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetBindings(v []*ListQueueUpStreamBindingsResponseBodyDataBindings) *ListQueueUpStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetMaxResults(v string) *ListQueueUpStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetNextToken(v string) *ListQueueUpStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) Validate() error {
	if s.Bindings != nil {
		for _, item := range s.Bindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueueUpStreamBindingsResponseBodyDataBindings struct {
	// The x-match property. Valid values:
	//
	// - **all**: This is the default value. All key-value pairs in the message header must match.
	//
	// - **any**: At least one key-value pair in the message header must match.
	//
	// This parameter is valid only for headers exchanges. It is invalid for other types of exchanges.
	//
	// example:
	//
	// all
	Argument *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	// The binding key.
	//
	// - If the source exchange is not a topic exchange:
	//
	//   - The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//   - The binding key must be 1 to 255 characters in length.
	//
	// - If the source exchange is a topic exchange:
	//
	//   - The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//   - The binding key cannot start or end with a period (.). If a number sign (#) or an asterisk (\\*) is at the beginning of the key, it must be followed by a period (.). If it is at the end of the key, it must be preceded by a period (.). If it is in the middle of the key, it must be enclosed by periods (.).
	//
	//   - The binding key must be 1 to 255 characters in length.
	//
	// example:
	//
	// amq.test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the destination object. Valid values:
	//
	// - **QUEUE**
	//
	// - **EXCHANGE**
	//
	// example:
	//
	// QUEUE
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The name of the destination.
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

func (s ListQueueUpStreamBindingsResponseBodyDataBindings) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) GetArgument() *string {
	return s.Argument
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) GetBindingKey() *string {
	return s.BindingKey
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) GetBindingType() *string {
	return s.BindingType
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) GetDestinationName() *string {
	return s.DestinationName
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) GetSourceExchange() *string {
	return s.SourceExchange
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) Validate() error {
	return dara.Validate(s)
}
