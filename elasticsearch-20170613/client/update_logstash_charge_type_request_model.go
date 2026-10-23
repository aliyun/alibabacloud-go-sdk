// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateLogstashChargeTypeRequest
	GetBody() *string
	SetClientToken(v string) *UpdateLogstashChargeTypeRequest
	GetClientToken() *string
}

type UpdateLogstashChargeTypeRequest struct {
	// The billing information of the Logstash instance after conversion, including the subscription duration and billing cycle.
	//
	// example:
	//
	// {"paymentInfo":{"duration":1,"pricingCycle":"Month"},"paymentType":"prepaid"}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The client generates the parameter value. The value must be unique across different requests and can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashChargeTypeRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateLogstashChargeTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLogstashChargeTypeRequest) SetBody(v string) *UpdateLogstashChargeTypeRequest {
	s.Body = &v
	return s
}

func (s *UpdateLogstashChargeTypeRequest) SetClientToken(v string) *UpdateLogstashChargeTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLogstashChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
