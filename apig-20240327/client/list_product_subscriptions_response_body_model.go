// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductSubscriptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProductSubscriptionsResponseBody
	GetCode() *string
	SetData(v *ListProductSubscriptionsResponseBodyData) *ListProductSubscriptionsResponseBody
	GetData() *ListProductSubscriptionsResponseBodyData
	SetMessage(v string) *ListProductSubscriptionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProductSubscriptionsResponseBody
	GetRequestId() *string
}

type ListProductSubscriptionsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListProductSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProductSubscriptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductSubscriptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProductSubscriptionsResponseBody) GetData() *ListProductSubscriptionsResponseBodyData {
	return s.Data
}

func (s *ListProductSubscriptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProductSubscriptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductSubscriptionsResponseBody) SetCode(v string) *ListProductSubscriptionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductSubscriptionsResponseBody) SetData(v *ListProductSubscriptionsResponseBodyData) *ListProductSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductSubscriptionsResponseBody) SetMessage(v string) *ListProductSubscriptionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductSubscriptionsResponseBody) SetRequestId(v string) *ListProductSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductSubscriptionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductSubscriptionsResponseBodyData struct {
	// The subscription list of the API product. The list element is named subscription, and each element represents an API product subscription record.
	Content []*ListProductSubscriptionsResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
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

func (s ListProductSubscriptionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProductSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductSubscriptionsResponseBodyData) GetContent() []*ListProductSubscriptionsResponseBodyDataContent {
	return s.Content
}

func (s *ListProductSubscriptionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductSubscriptionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductSubscriptionsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProductSubscriptionsResponseBodyData) SetContent(v []*ListProductSubscriptionsResponseBodyDataContent) *ListProductSubscriptionsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListProductSubscriptionsResponseBodyData) SetPageNumber(v int32) *ListProductSubscriptionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyData) SetPageSize(v int32) *ListProductSubscriptionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyData) SetTotalCount(v int64) *ListProductSubscriptionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyData) Validate() error {
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

type ListProductSubscriptionsResponseBodyDataContent struct {
	// The ID of the consumer to which the subscription belongs. This ID and the consumerName returned in the same element refer to the same consumer.
	//
	// example:
	//
	// cs-xxxxxxxxxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The name of the consumer to which the subscription belongs. This corresponds to the keyword search for the consumerName request parameter.
	//
	// example:
	//
	// consumer-1
	ConsumerName *string `json:"consumerName,omitempty" xml:"consumerName,omitempty"`
	// The ID of the subscribed API product. This is the primary ID of the API product resource and is the same as the productId path parameter in this operation.
	//
	// example:
	//
	// product-xxxxxxxx
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The name of the subscribed API product. This is the Name property of the API product resource.
	//
	// example:
	//
	// model-service
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// The type of the subscribed API product. Valid values: REST_API, HTTP_API, MCP_SERVER, AGENT_API, and MODEL_API.
	//
	// example:
	//
	// MCP_SERVER
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// The subscription status. Valid values: PENDING and APPROVED. (These values are defined in the enumeration for the status request parameter. This response parameter is not annotated with @enums.)
	//
	// example:
	//
	// APPROVED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The subscription ID, which identifies an API product subscription record. You can use this ID to cancel the subscription by calling UnsubscribeHiMarketProduct or approve the subscription by calling ApproveProductSubscription. Both operations locate the record by using the /himarket/v1/consumers/{consumerId}/subscriptions/{subscriptionId} path.
	//
	// example:
	//
	// sub-xxxxxxxx
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s ListProductSubscriptionsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListProductSubscriptionsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListProductSubscriptionsResponseBodyDataContent) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListProductSubscriptionsResponseBodyDataContent) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *ListProductSubscriptionsResponseBodyDataContent) GetProductId() *string {
	return s.ProductId
}

func (s *ListProductSubscriptionsResponseBodyDataContent) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductSubscriptionsResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *ListProductSubscriptionsResponseBodyDataContent) GetStatus() *string {
	return s.Status
}

func (s *ListProductSubscriptionsResponseBodyDataContent) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListProductSubscriptionsResponseBodyDataContent) SetConsumerId(v string) *ListProductSubscriptionsResponseBodyDataContent {
	s.ConsumerId = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyDataContent) SetConsumerName(v string) *ListProductSubscriptionsResponseBodyDataContent {
	s.ConsumerName = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyDataContent) SetProductId(v string) *ListProductSubscriptionsResponseBodyDataContent {
	s.ProductId = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyDataContent) SetProductName(v string) *ListProductSubscriptionsResponseBodyDataContent {
	s.ProductName = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyDataContent) SetProductType(v string) *ListProductSubscriptionsResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyDataContent) SetStatus(v string) *ListProductSubscriptionsResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyDataContent) SetSubscriptionId(v string) *ListProductSubscriptionsResponseBodyDataContent {
	s.SubscriptionId = &v
	return s
}

func (s *ListProductSubscriptionsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
