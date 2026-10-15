// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInquiryPriceRefundInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *InquiryPriceRefundInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *InquiryPriceRefundInstanceRequest
	GetInstanceId() *string
	SetProductCode(v string) *InquiryPriceRefundInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *InquiryPriceRefundInstanceRequest
	GetProductType() *string
}

type InquiryPriceRefundInstanceRequest struct {
	// Required for idempotent scenarios. A UUID-generated string. This field is used for idempotence. If the same value is passed in multiple calls, the result of the first call is returned.
	//
	// example:
	//
	// 793F021C-B589-1225-82A9-99232AEBE494
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Required for unsubscription scenarios. The instance ID.
	//
	// example:
	//
	// i-bp1etb69sqxgl4*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Required for unsubscription scenarios. The product code.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Required for unsubscription scenarios. The product type. Pass an empty string ("") unless otherwise specified.
	//
	// example:
	//
	// ”“
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s InquiryPriceRefundInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s InquiryPriceRefundInstanceRequest) GoString() string {
	return s.String()
}

func (s *InquiryPriceRefundInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InquiryPriceRefundInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InquiryPriceRefundInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InquiryPriceRefundInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *InquiryPriceRefundInstanceRequest) SetClientToken(v string) *InquiryPriceRefundInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *InquiryPriceRefundInstanceRequest) SetInstanceId(v string) *InquiryPriceRefundInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *InquiryPriceRefundInstanceRequest) SetProductCode(v string) *InquiryPriceRefundInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *InquiryPriceRefundInstanceRequest) SetProductType(v string) *InquiryPriceRefundInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *InquiryPriceRefundInstanceRequest) Validate() error {
	return dara.Validate(s)
}
