// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v []interface{}) *CreateDeliveryPlanResponseBody
	GetArguments() []interface{}
	SetRequestId(v string) *CreateDeliveryPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDeliveryPlanResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *CreateDeliveryPlanResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateDeliveryPlanResponseBody
	GetVendorType() *string
}

type CreateDeliveryPlanResponseBody struct {
	// example:
	//
	// []
	Arguments []interface{} `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CreateDeliveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanResponseBody) GetArguments() []interface{} {
	return s.Arguments
}

func (s *CreateDeliveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeliveryPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDeliveryPlanResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateDeliveryPlanResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateDeliveryPlanResponseBody) SetArguments(v []interface{}) *CreateDeliveryPlanResponseBody {
	s.Arguments = v
	return s
}

func (s *CreateDeliveryPlanResponseBody) SetRequestId(v string) *CreateDeliveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeliveryPlanResponseBody) SetSuccess(v bool) *CreateDeliveryPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDeliveryPlanResponseBody) SetVendorRequestId(v string) *CreateDeliveryPlanResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateDeliveryPlanResponseBody) SetVendorType(v string) *CreateDeliveryPlanResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateDeliveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
