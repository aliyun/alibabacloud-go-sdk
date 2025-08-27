// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBtripBillInfoAdjustResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BtripBillInfoAdjustResponseBody
	GetCode() *string
	SetMessage(v string) *BtripBillInfoAdjustResponseBody
	GetMessage() *string
	SetRequestId(v string) *BtripBillInfoAdjustResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BtripBillInfoAdjustResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BtripBillInfoAdjustResponseBody
	GetTraceId() *string
}

type BtripBillInfoAdjustResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s BtripBillInfoAdjustResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BtripBillInfoAdjustResponseBody) GoString() string {
	return s.String()
}

func (s *BtripBillInfoAdjustResponseBody) GetCode() *string {
	return s.Code
}

func (s *BtripBillInfoAdjustResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BtripBillInfoAdjustResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BtripBillInfoAdjustResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BtripBillInfoAdjustResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BtripBillInfoAdjustResponseBody) SetCode(v string) *BtripBillInfoAdjustResponseBody {
	s.Code = &v
	return s
}

func (s *BtripBillInfoAdjustResponseBody) SetMessage(v string) *BtripBillInfoAdjustResponseBody {
	s.Message = &v
	return s
}

func (s *BtripBillInfoAdjustResponseBody) SetRequestId(v string) *BtripBillInfoAdjustResponseBody {
	s.RequestId = &v
	return s
}

func (s *BtripBillInfoAdjustResponseBody) SetSuccess(v bool) *BtripBillInfoAdjustResponseBody {
	s.Success = &v
	return s
}

func (s *BtripBillInfoAdjustResponseBody) SetTraceId(v string) *BtripBillInfoAdjustResponseBody {
	s.TraceId = &v
	return s
}

func (s *BtripBillInfoAdjustResponseBody) Validate() error {
	return dara.Validate(s)
}
