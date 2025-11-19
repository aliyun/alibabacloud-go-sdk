// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateOrderShrinkRequest
	GetApiKey() *string
	SetAutoPay(v bool) *CreateOrderShrinkRequest
	GetAutoPay() *bool
	SetDynamicProductParamsShrink(v string) *CreateOrderShrinkRequest
	GetDynamicProductParamsShrink() *string
	SetOperationType(v string) *CreateOrderShrinkRequest
	GetOperationType() *string
	SetOrderFrom(v string) *CreateOrderShrinkRequest
	GetOrderFrom() *string
	SetSessionId(v string) *CreateOrderShrinkRequest
	GetSessionId() *string
}

type CreateOrderShrinkRequest struct {
	// This parameter is required.
	ApiKey  *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	AutoPay *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// This parameter is required.
	DynamicProductParamsShrink *string `json:"DynamicProductParams,omitempty" xml:"DynamicProductParams,omitempty"`
	// This parameter is required.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OrderFrom     *string `json:"OrderFrom,omitempty" xml:"OrderFrom,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s CreateOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateOrderShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateOrderShrinkRequest) GetDynamicProductParamsShrink() *string {
	return s.DynamicProductParamsShrink
}

func (s *CreateOrderShrinkRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateOrderShrinkRequest) GetOrderFrom() *string {
	return s.OrderFrom
}

func (s *CreateOrderShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateOrderShrinkRequest) SetApiKey(v string) *CreateOrderShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetAutoPay(v bool) *CreateOrderShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetDynamicProductParamsShrink(v string) *CreateOrderShrinkRequest {
	s.DynamicProductParamsShrink = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetOperationType(v string) *CreateOrderShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetOrderFrom(v string) *CreateOrderShrinkRequest {
	s.OrderFrom = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetSessionId(v string) *CreateOrderShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *CreateOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
