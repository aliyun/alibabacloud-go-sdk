// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubscribeProductResponseBody
	GetCode() *string
	SetData(v *SubscribeProductResponseBodyData) *SubscribeProductResponseBody
	GetData() *SubscribeProductResponseBodyData
	SetMessage(v string) *SubscribeProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubscribeProductResponseBody
	GetRequestId() *string
}

type SubscribeProductResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *SubscribeProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s SubscribeProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubscribeProductResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubscribeProductResponseBody) GetData() *SubscribeProductResponseBodyData {
	return s.Data
}

func (s *SubscribeProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubscribeProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubscribeProductResponseBody) SetCode(v string) *SubscribeProductResponseBody {
	s.Code = &v
	return s
}

func (s *SubscribeProductResponseBody) SetData(v *SubscribeProductResponseBodyData) *SubscribeProductResponseBody {
	s.Data = v
	return s
}

func (s *SubscribeProductResponseBody) SetMessage(v string) *SubscribeProductResponseBody {
	s.Message = &v
	return s
}

func (s *SubscribeProductResponseBody) SetRequestId(v string) *SubscribeProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubscribeProductResponseBodyData struct {
	// The ID of the consumer to which the subscription belongs. This is the same identifier as the consumerId path parameter in the request and refers to the same consumer as the consumerName field returned at the same level.
	//
	// example:
	//
	// cs-xxxxxxxxxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The name of the consumer to which the subscription belongs. This refers to the same consumer as the consumerId field returned at the same level.
	//
	// example:
	//
	// consumer-1
	ConsumerName *string `json:"consumerName,omitempty" xml:"consumerName,omitempty"`
	// The ID of the subscribed API product. This is the primary ID of the API product resource and is the same identifier as the productId passed in the request body.
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
	// The type of the subscribed API product. The valid values are not declared in this operation definition. The response parameter with the same name in the related ListProductSubscriptions operation includes the following valid values: REST_API, HTTP_API, MCP_SERVER, AGENT_API, and MODEL_API.
	//
	// example:
	//
	// MCP_SERVER
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// The subscription status. Valid values: PENDING and APPROVED. These valid values are declared in the status request parameter of the related ListProductSubscriptions operation.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The subscription ID. Identifies the API product subscription record corresponding to this subscription. You can use this ID to cancel the subscription by calling UnsubscribeHiMarketProduct or approve the subscription by calling ApproveProductSubscription. Both operations locate the record by using the /himarket/v1/consumers/{consumerId}/subscriptions/{subscriptionId} path.
	//
	// example:
	//
	// sub-xxxxxxxx
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
}

func (s SubscribeProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubscribeProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubscribeProductResponseBodyData) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *SubscribeProductResponseBodyData) GetConsumerName() *string {
	return s.ConsumerName
}

func (s *SubscribeProductResponseBodyData) GetProductId() *string {
	return s.ProductId
}

func (s *SubscribeProductResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *SubscribeProductResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *SubscribeProductResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubscribeProductResponseBodyData) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *SubscribeProductResponseBodyData) SetConsumerId(v string) *SubscribeProductResponseBodyData {
	s.ConsumerId = &v
	return s
}

func (s *SubscribeProductResponseBodyData) SetConsumerName(v string) *SubscribeProductResponseBodyData {
	s.ConsumerName = &v
	return s
}

func (s *SubscribeProductResponseBodyData) SetProductId(v string) *SubscribeProductResponseBodyData {
	s.ProductId = &v
	return s
}

func (s *SubscribeProductResponseBodyData) SetProductName(v string) *SubscribeProductResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *SubscribeProductResponseBodyData) SetProductType(v string) *SubscribeProductResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *SubscribeProductResponseBodyData) SetStatus(v string) *SubscribeProductResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubscribeProductResponseBodyData) SetSubscriptionId(v string) *SubscribeProductResponseBodyData {
	s.SubscriptionId = &v
	return s
}

func (s *SubscribeProductResponseBodyData) Validate() error {
	return dara.Validate(s)
}
