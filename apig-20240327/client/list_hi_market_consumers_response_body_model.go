// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHiMarketConsumersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHiMarketConsumersResponseBody
	GetCode() *string
	SetData(v *ListHiMarketConsumersResponseBodyData) *ListHiMarketConsumersResponseBody
	GetData() *ListHiMarketConsumersResponseBodyData
	SetMaxResults(v int32) *ListHiMarketConsumersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListHiMarketConsumersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListHiMarketConsumersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHiMarketConsumersResponseBody
	GetRequestId() *string
}

type ListHiMarketConsumersResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListHiMarketConsumersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The maximum number of entries on the current page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. An empty value indicates that the last page has been reached.
	//
	// example:
	//
	// token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListHiMarketConsumersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHiMarketConsumersResponseBody) GetData() *ListHiMarketConsumersResponseBodyData {
	return s.Data
}

func (s *ListHiMarketConsumersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHiMarketConsumersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHiMarketConsumersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHiMarketConsumersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHiMarketConsumersResponseBody) SetCode(v string) *ListHiMarketConsumersResponseBody {
	s.Code = &v
	return s
}

func (s *ListHiMarketConsumersResponseBody) SetData(v *ListHiMarketConsumersResponseBodyData) *ListHiMarketConsumersResponseBody {
	s.Data = v
	return s
}

func (s *ListHiMarketConsumersResponseBody) SetMaxResults(v int32) *ListHiMarketConsumersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHiMarketConsumersResponseBody) SetMessage(v string) *ListHiMarketConsumersResponseBody {
	s.Message = &v
	return s
}

func (s *ListHiMarketConsumersResponseBody) SetNextToken(v string) *ListHiMarketConsumersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHiMarketConsumersResponseBody) SetRequestId(v string) *ListHiMarketConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHiMarketConsumersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHiMarketConsumersResponseBodyData struct {
	// The list of consumers. Each element represents a HiMarket consumer.
	Content []*ListHiMarketConsumersResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListHiMarketConsumersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumersResponseBodyData) GetContent() []*ListHiMarketConsumersResponseBodyDataContent {
	return s.Content
}

func (s *ListHiMarketConsumersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHiMarketConsumersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHiMarketConsumersResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListHiMarketConsumersResponseBodyData) SetContent(v []*ListHiMarketConsumersResponseBodyDataContent) *ListHiMarketConsumersResponseBodyData {
	s.Content = v
	return s
}

func (s *ListHiMarketConsumersResponseBodyData) SetPageNumber(v int32) *ListHiMarketConsumersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHiMarketConsumersResponseBodyData) SetPageSize(v int32) *ListHiMarketConsumersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHiMarketConsumersResponseBodyData) SetTotalCount(v int64) *ListHiMarketConsumersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHiMarketConsumersResponseBodyData) Validate() error {
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

type ListHiMarketConsumersResponseBodyDataContent struct {
	// The consumer ID. You can use this ID to locate the consumer in operations such as GetHiMarketConsumer, DeleteHiMarketConsumer, GetHiMarketConsumerCredential, and ListHiMarketConsumerSubscriptions.
	//
	// example:
	//
	// cs-xxxxxxxxxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The description of the consumer.
	//
	// example:
	//
	// Test consumer
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the consumer is a primary consumer.
	IsPrimary *bool `json:"isPrimary,omitempty" xml:"isPrimary,omitempty"`
	// The name of the consumer.
	//
	// example:
	//
	// consumer-1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListHiMarketConsumersResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumersResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumersResponseBodyDataContent) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListHiMarketConsumersResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListHiMarketConsumersResponseBodyDataContent) GetIsPrimary() *bool {
	return s.IsPrimary
}

func (s *ListHiMarketConsumersResponseBodyDataContent) GetName() *string {
	return s.Name
}

func (s *ListHiMarketConsumersResponseBodyDataContent) SetConsumerId(v string) *ListHiMarketConsumersResponseBodyDataContent {
	s.ConsumerId = &v
	return s
}

func (s *ListHiMarketConsumersResponseBodyDataContent) SetDescription(v string) *ListHiMarketConsumersResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListHiMarketConsumersResponseBodyDataContent) SetIsPrimary(v bool) *ListHiMarketConsumersResponseBodyDataContent {
	s.IsPrimary = &v
	return s
}

func (s *ListHiMarketConsumersResponseBodyDataContent) SetName(v string) *ListHiMarketConsumersResponseBodyDataContent {
	s.Name = &v
	return s
}

func (s *ListHiMarketConsumersResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
