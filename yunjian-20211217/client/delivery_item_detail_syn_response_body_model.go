// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliveryItemDetailSynResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeliveryItemDetailSynResponseBody
	GetCode() *int64
	SetData(v bool) *DeliveryItemDetailSynResponseBody
	GetData() *bool
	SetMessage(v string) *DeliveryItemDetailSynResponseBody
	GetMessage() *string
	SetSuccess(v bool) *DeliveryItemDetailSynResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeliveryItemDetailSynResponseBody
	GetTraceId() *string
}

type DeliveryItemDetailSynResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 212cf01016405759151137225e83cd
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeliveryItemDetailSynResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeliveryItemDetailSynResponseBody) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeliveryItemDetailSynResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeliveryItemDetailSynResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeliveryItemDetailSynResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeliveryItemDetailSynResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeliveryItemDetailSynResponseBody) SetCode(v int64) *DeliveryItemDetailSynResponseBody {
	s.Code = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) SetData(v bool) *DeliveryItemDetailSynResponseBody {
	s.Data = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) SetMessage(v string) *DeliveryItemDetailSynResponseBody {
	s.Message = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) SetSuccess(v bool) *DeliveryItemDetailSynResponseBody {
	s.Success = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) SetTraceId(v string) *DeliveryItemDetailSynResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) Validate() error {
	return dara.Validate(s)
}
