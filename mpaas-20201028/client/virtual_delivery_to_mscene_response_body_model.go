// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualDeliveryToMsceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMpaasUserGamecenterPaymentVirtualdeliveryResponse(v *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse) *VirtualDeliveryToMsceneResponseBody
	GetMpaasUserGamecenterPaymentVirtualdeliveryResponse() *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse
	SetRequestId(v string) *VirtualDeliveryToMsceneResponseBody
	GetRequestId() *string
	SetResultCode(v string) *VirtualDeliveryToMsceneResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *VirtualDeliveryToMsceneResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *VirtualDeliveryToMsceneResponseBody
	GetSuccess() *bool
}

type VirtualDeliveryToMsceneResponseBody struct {
	MpaasUserGamecenterPaymentVirtualdeliveryResponse *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse `json:"MpaasUserGamecenterPaymentVirtualdeliveryResponse,omitempty" xml:"MpaasUserGamecenterPaymentVirtualdeliveryResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// ac1f0083177615939018778261913
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// SUCCESS
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VirtualDeliveryToMsceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VirtualDeliveryToMsceneResponseBody) GoString() string {
	return s.String()
}

func (s *VirtualDeliveryToMsceneResponseBody) GetMpaasUserGamecenterPaymentVirtualdeliveryResponse() *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse {
	return s.MpaasUserGamecenterPaymentVirtualdeliveryResponse
}

func (s *VirtualDeliveryToMsceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VirtualDeliveryToMsceneResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *VirtualDeliveryToMsceneResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *VirtualDeliveryToMsceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VirtualDeliveryToMsceneResponseBody) SetMpaasUserGamecenterPaymentVirtualdeliveryResponse(v *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse) *VirtualDeliveryToMsceneResponseBody {
	s.MpaasUserGamecenterPaymentVirtualdeliveryResponse = v
	return s
}

func (s *VirtualDeliveryToMsceneResponseBody) SetRequestId(v string) *VirtualDeliveryToMsceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *VirtualDeliveryToMsceneResponseBody) SetResultCode(v string) *VirtualDeliveryToMsceneResponseBody {
	s.ResultCode = &v
	return s
}

func (s *VirtualDeliveryToMsceneResponseBody) SetResultMsg(v string) *VirtualDeliveryToMsceneResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *VirtualDeliveryToMsceneResponseBody) SetSuccess(v bool) *VirtualDeliveryToMsceneResponseBody {
	s.Success = &v
	return s
}

func (s *VirtualDeliveryToMsceneResponseBody) Validate() error {
	if s.MpaasUserGamecenterPaymentVirtualdeliveryResponse != nil {
		if err := s.MpaasUserGamecenterPaymentVirtualdeliveryResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse) GoString() string {
	return s.String()
}

func (s *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse) GetSuccess() *bool {
	return s.Success
}

func (s *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse) SetSuccess(v bool) *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse {
	s.Success = &v
	return s
}

func (s *VirtualDeliveryToMsceneResponseBodyMpaasUserGamecenterPaymentVirtualdeliveryResponse) Validate() error {
	return dara.Validate(s)
}
