// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorHotelEventPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CooperatorHotelEventPushResponseBody
	GetCode() *string
	SetMessage(v string) *CooperatorHotelEventPushResponseBody
	GetMessage() *string
	SetModule(v bool) *CooperatorHotelEventPushResponseBody
	GetModule() *bool
	SetRequestId(v string) *CooperatorHotelEventPushResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CooperatorHotelEventPushResponseBody
	GetSuccess() *bool
}

type CooperatorHotelEventPushResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Module *bool `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CooperatorHotelEventPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelEventPushResponseBody) GoString() string {
	return s.String()
}

func (s *CooperatorHotelEventPushResponseBody) GetCode() *string {
	return s.Code
}

func (s *CooperatorHotelEventPushResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CooperatorHotelEventPushResponseBody) GetModule() *bool {
	return s.Module
}

func (s *CooperatorHotelEventPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CooperatorHotelEventPushResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CooperatorHotelEventPushResponseBody) SetCode(v string) *CooperatorHotelEventPushResponseBody {
	s.Code = &v
	return s
}

func (s *CooperatorHotelEventPushResponseBody) SetMessage(v string) *CooperatorHotelEventPushResponseBody {
	s.Message = &v
	return s
}

func (s *CooperatorHotelEventPushResponseBody) SetModule(v bool) *CooperatorHotelEventPushResponseBody {
	s.Module = &v
	return s
}

func (s *CooperatorHotelEventPushResponseBody) SetRequestId(v string) *CooperatorHotelEventPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *CooperatorHotelEventPushResponseBody) SetSuccess(v bool) *CooperatorHotelEventPushResponseBody {
	s.Success = &v
	return s
}

func (s *CooperatorHotelEventPushResponseBody) Validate() error {
	return dara.Validate(s)
}
