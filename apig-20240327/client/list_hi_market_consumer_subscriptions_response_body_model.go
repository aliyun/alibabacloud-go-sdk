// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHiMarketConsumerSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHiMarketConsumerSubscriptionsResponseBody
	GetCode() *string
	SetData(v *ListHiMarketConsumerSubscriptionsResponseBodyData) *ListHiMarketConsumerSubscriptionsResponseBody
	GetData() *ListHiMarketConsumerSubscriptionsResponseBodyData
	SetMaxResults(v int32) *ListHiMarketConsumerSubscriptionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListHiMarketConsumerSubscriptionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListHiMarketConsumerSubscriptionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHiMarketConsumerSubscriptionsResponseBody
	GetRequestId() *string
}

type ListHiMarketConsumerSubscriptionsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListHiMarketConsumerSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The cursor for the next page. An empty value indicates that the last page is reached.
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

func (s ListHiMarketConsumerSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumerSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) GetData() *ListHiMarketConsumerSubscriptionsResponseBodyData {
	return s.Data
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) SetCode(v string) *ListHiMarketConsumerSubscriptionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) SetData(v *ListHiMarketConsumerSubscriptionsResponseBodyData) *ListHiMarketConsumerSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) SetMaxResults(v int32) *ListHiMarketConsumerSubscriptionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) SetMessage(v string) *ListHiMarketConsumerSubscriptionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) SetNextToken(v string) *ListHiMarketConsumerSubscriptionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) SetRequestId(v string) *ListHiMarketConsumerSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHiMarketConsumerSubscriptionsResponseBodyData struct {
	// The list of API product subscriptions for the consumer. The list element is named subscription, and each element represents an API product subscription record.
	Content []*ListHiMarketConsumerSubscriptionsResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
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

func (s ListHiMarketConsumerSubscriptionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumerSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) GetContent() []*ListHiMarketConsumerSubscriptionsResponseBodyDataContent {
	return s.Content
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) SetContent(v []*ListHiMarketConsumerSubscriptionsResponseBodyDataContent) *ListHiMarketConsumerSubscriptionsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) SetPageNumber(v int32) *ListHiMarketConsumerSubscriptionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) SetPageSize(v int32) *ListHiMarketConsumerSubscriptionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) SetTotalCount(v int64) *ListHiMarketConsumerSubscriptionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyData) Validate() error {
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

type ListHiMarketConsumerSubscriptionsResponseBodyDataContent struct {
	// The ID of the consumer to which the subscription belongs. This is the same identifier as the consumerId path parameter of this API operation.
	//
	// example:
	//
	// cs-xxxxxxxxxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The name of the consumer to which the subscription belongs.
	//
	// example:
	//
	// consumer-1
	ConsumerName *string `json:"consumerName,omitempty" xml:"consumerName,omitempty"`
	// The ID of the subscribed API product, which is the primary ID of the API product resource.
	//
	// example:
	//
	// product-xxxxxxxx
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The name of the subscribed API product.
	//
	// example:
	//
	// model-service
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// The type of the subscribed API product. This response parameter does not declare enums. The related ListProductSubscriptions API operation declares the following valid values for the parameter with the same name: REST_API, HTTP_API, MCP_SERVER, AGENT_API, and MODEL_API.
	//
	// example:
	//
	// MCP_SERVER
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// The subscription status. Valid values: PENDING and APPROVED. This set of valid values is declared in the enums of the status request parameter in this API operation, and this response parameter does not declare enums.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The subscription ID. You can use this ID to unsubscribe from the product by calling UnsubscribeHiMarketProduct or approve the subscription by calling ApproveProductSubscription. Both operations locate the subscription by using the /himarket/v1/consumers/{consumerId}/subscriptions/{subscriptionId} path.
	//
	// example:
	//
	// sub-xxxxxxxx
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s ListHiMarketConsumerSubscriptionsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumerSubscriptionsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) GetProductId() *string {
	return s.ProductId
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) GetProductName() *string {
	return s.ProductName
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) GetStatus() *string {
	return s.Status
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) SetConsumerId(v string) *ListHiMarketConsumerSubscriptionsResponseBodyDataContent {
	s.ConsumerId = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) SetConsumerName(v string) *ListHiMarketConsumerSubscriptionsResponseBodyDataContent {
	s.ConsumerName = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) SetProductId(v string) *ListHiMarketConsumerSubscriptionsResponseBodyDataContent {
	s.ProductId = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) SetProductName(v string) *ListHiMarketConsumerSubscriptionsResponseBodyDataContent {
	s.ProductName = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) SetProductType(v string) *ListHiMarketConsumerSubscriptionsResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) SetStatus(v string) *ListHiMarketConsumerSubscriptionsResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) SetSubscriptionId(v string) *ListHiMarketConsumerSubscriptionsResponseBodyDataContent {
	s.SubscriptionId = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
