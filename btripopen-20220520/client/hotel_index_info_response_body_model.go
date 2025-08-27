// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelIndexInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelIndexInfoResponseBody
	GetCode() *string
	SetMessage(v string) *HotelIndexInfoResponseBody
	GetMessage() *string
	SetModule(v *HotelIndexInfoResponseBodyModule) *HotelIndexInfoResponseBody
	GetModule() *HotelIndexInfoResponseBodyModule
	SetRequestId(v string) *HotelIndexInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelIndexInfoResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelIndexInfoResponseBody
	GetTraceId() *string
}

type HotelIndexInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// operation success.
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelIndexInfoResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s HotelIndexInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelIndexInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HotelIndexInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelIndexInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelIndexInfoResponseBody) GetModule() *HotelIndexInfoResponseBodyModule {
	return s.Module
}

func (s *HotelIndexInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelIndexInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelIndexInfoResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelIndexInfoResponseBody) SetCode(v string) *HotelIndexInfoResponseBody {
	s.Code = &v
	return s
}

func (s *HotelIndexInfoResponseBody) SetMessage(v string) *HotelIndexInfoResponseBody {
	s.Message = &v
	return s
}

func (s *HotelIndexInfoResponseBody) SetModule(v *HotelIndexInfoResponseBodyModule) *HotelIndexInfoResponseBody {
	s.Module = v
	return s
}

func (s *HotelIndexInfoResponseBody) SetRequestId(v string) *HotelIndexInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelIndexInfoResponseBody) SetSuccess(v bool) *HotelIndexInfoResponseBody {
	s.Success = &v
	return s
}

func (s *HotelIndexInfoResponseBody) SetTraceId(v string) *HotelIndexInfoResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelIndexInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelIndexInfoResponseBodyModule struct {
	Items []*HotelIndexInfoResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 182873
	PageToken *string `json:"page_token,omitempty" xml:"page_token,omitempty"`
}

func (s HotelIndexInfoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelIndexInfoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelIndexInfoResponseBodyModule) GetItems() []*HotelIndexInfoResponseBodyModuleItems {
	return s.Items
}

func (s *HotelIndexInfoResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotelIndexInfoResponseBodyModule) GetPageToken() *string {
	return s.PageToken
}

func (s *HotelIndexInfoResponseBodyModule) SetItems(v []*HotelIndexInfoResponseBodyModuleItems) *HotelIndexInfoResponseBodyModule {
	s.Items = v
	return s
}

func (s *HotelIndexInfoResponseBodyModule) SetPageSize(v int32) *HotelIndexInfoResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *HotelIndexInfoResponseBodyModule) SetPageToken(v string) *HotelIndexInfoResponseBodyModule {
	s.PageToken = &v
	return s
}

func (s *HotelIndexInfoResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelIndexInfoResponseBodyModuleItems struct {
	// example:
	//
	// 182873
	HotelId *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// example:
	//
	// 0
	HotelStatus *string `json:"hotel_status,omitempty" xml:"hotel_status,omitempty"`
}

func (s HotelIndexInfoResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s HotelIndexInfoResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *HotelIndexInfoResponseBodyModuleItems) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelIndexInfoResponseBodyModuleItems) GetHotelStatus() *string {
	return s.HotelStatus
}

func (s *HotelIndexInfoResponseBodyModuleItems) SetHotelId(v string) *HotelIndexInfoResponseBodyModuleItems {
	s.HotelId = &v
	return s
}

func (s *HotelIndexInfoResponseBodyModuleItems) SetHotelStatus(v string) *HotelIndexInfoResponseBodyModuleItems {
	s.HotelStatus = &v
	return s
}

func (s *HotelIndexInfoResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
