// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeUpStreamBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListExchangeUpStreamBindingsResponseBody
	GetCode() *int32
	SetData(v *ListExchangeUpStreamBindingsResponseBodyData) *ListExchangeUpStreamBindingsResponseBody
	GetData() *ListExchangeUpStreamBindingsResponseBodyData
	SetMessage(v string) *ListExchangeUpStreamBindingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListExchangeUpStreamBindingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExchangeUpStreamBindingsResponseBody
	GetSuccess() *bool
}

type ListExchangeUpStreamBindingsResponseBody struct {
	// The return code. A value of 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListExchangeUpStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 2DCCCE88-BC82-4A4F-AF5E-9A759672B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListExchangeUpStreamBindingsResponseBody) GetData() *ListExchangeUpStreamBindingsResponseBodyData {
	return s.Data
}

func (s *ListExchangeUpStreamBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExchangeUpStreamBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExchangeUpStreamBindingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetCode(v int32) *ListExchangeUpStreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetData(v *ListExchangeUpStreamBindingsResponseBodyData) *ListExchangeUpStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetMessage(v string) *ListExchangeUpStreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetRequestId(v string) *ListExchangeUpStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetSuccess(v bool) *ListExchangeUpStreamBindingsResponseBody {
	s.Success = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExchangeUpStreamBindingsResponseBodyData struct {
	// The bindings.
	Bindings []*ListExchangeUpStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// The maximum number of results returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that indicates the position where the current query ends. An empty value indicates that all data has been read.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) GetBindings() []*ListExchangeUpStreamBindingsResponseBodyDataBindings {
	return s.Bindings
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetBindings(v []*ListExchangeUpStreamBindingsResponseBodyDataBindings) *ListExchangeUpStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetMaxResults(v int32) *ListExchangeUpStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetNextToken(v string) *ListExchangeUpStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) Validate() error {
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

type ListExchangeUpStreamBindingsResponseBodyDataBindings struct {
	// The x-match property. Valid values:
	//
	// - **all**: The default value. All key-value pairs in the message header must match.
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
	//   - The key can contain only letters, digits, hyphens (-), underscores (_), periods (.), forward slashes (/), and at signs (@).
	//
	//   - The key must be 1 to 255 characters in length.
	//
	// - If the source exchange is a topic exchange:
	//
	//   - The key can contain letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	//   - The key cannot start or end with a period (.). A number sign (#) or an asterisk (\\*) must be preceded by a period if it is not at the start of the key. It must be followed by a period if it is not at the end of the key.
	//
	//   - The key must be 1 to 255 characters in length.
	//
	// example:
	//
	// amq.dle.test
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	// The type of the destination object. Valid values:
	//
	// - **QUEUE**
	//
	// - **EXCHANGE**
	//
	// example:
	//
	// EXCHANGE
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The destination name.
	//
	// example:
	//
	// test
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	// The name of the source exchange.
	//
	// example:
	//
	// dle
	SourceExchange *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBodyDataBindings) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) GetArgument() *string {
	return s.Argument
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) GetBindingKey() *string {
	return s.BindingKey
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) GetBindingType() *string {
	return s.BindingType
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) GetDestinationName() *string {
	return s.DestinationName
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) GetSourceExchange() *string {
	return s.SourceExchange
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) Validate() error {
	return dara.Validate(s)
}
