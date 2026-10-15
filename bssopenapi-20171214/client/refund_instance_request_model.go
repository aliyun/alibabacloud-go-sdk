// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RefundInstanceRequest
	GetClientToken() *string
	SetImmediatelyRelease(v string) *RefundInstanceRequest
	GetImmediatelyRelease() *string
	SetInstanceId(v string) *RefundInstanceRequest
	GetInstanceId() *string
	SetProductCode(v string) *RefundInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *RefundInstanceRequest
	GetProductType() *string
}

type RefundInstanceRequest struct {
	// Required for idempotent scenarios. A UUID-generated string. This field is used for idempotence. If the same value is passed in multiple calls, the result of the first call is returned. If a result is returned after the call, you can parse the refund order ID, which indicates that the call is successful.
	//
	// example:
	//
	// 793F021C-B589-1225-82A9-99232AEBE494
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Required for unsubscription scenarios. Valid values: 1: immediately releases the instance. 0: stops the instance first based on the shutdown policy. Only some products support this value. Default value: 1.
	//
	// example:
	//
	// 1
	ImmediatelyRelease *string `json:"ImmediatelyRelease,omitempty" xml:"ImmediatelyRelease,omitempty"`
	// Required for unsubscription scenarios. The instance ID, not the custom name.
	//
	// example:
	//
	// i-bp1etb69sqxgl4*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Required for unsubscription scenarios. The product code. If the error code returned after the call indicates that the product is not supported, submit a ticket to request that the cloud service enable this feature.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Required for unsubscription scenarios. The product type. Unless otherwise specified, pass an empty string. In the debugging tool, you do not need to manually enter a value. Hover over the input box, and a button labeled "Enter empty string" appears. Click this button. When initiating an unsubscription through code, pass an empty string "".
	//
	// example:
	//
	// Empty string
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s RefundInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundInstanceRequest) GoString() string {
	return s.String()
}

func (s *RefundInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RefundInstanceRequest) GetImmediatelyRelease() *string {
	return s.ImmediatelyRelease
}

func (s *RefundInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RefundInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *RefundInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RefundInstanceRequest) SetClientToken(v string) *RefundInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RefundInstanceRequest) SetImmediatelyRelease(v string) *RefundInstanceRequest {
	s.ImmediatelyRelease = &v
	return s
}

func (s *RefundInstanceRequest) SetInstanceId(v string) *RefundInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundInstanceRequest) SetProductCode(v string) *RefundInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *RefundInstanceRequest) SetProductType(v string) *RefundInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *RefundInstanceRequest) Validate() error {
	return dara.Validate(s)
}
