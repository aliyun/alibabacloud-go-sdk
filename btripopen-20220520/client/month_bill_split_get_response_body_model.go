// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillSplitGetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MonthBillSplitGetResponseBody
	GetCode() *string
	SetMessage(v string) *MonthBillSplitGetResponseBody
	GetMessage() *string
	SetModule(v []*MonthBillSplitGetResponseBodyModule) *MonthBillSplitGetResponseBody
	GetModule() []*MonthBillSplitGetResponseBodyModule
	SetRequestId(v string) *MonthBillSplitGetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MonthBillSplitGetResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MonthBillSplitGetResponseBody
	GetTraceId() *string
}

type MonthBillSplitGetResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  []*MonthBillSplitGetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// example:
	//
	// 407543AF-****-****-****-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce********056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MonthBillSplitGetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MonthBillSplitGetResponseBody) GoString() string {
	return s.String()
}

func (s *MonthBillSplitGetResponseBody) GetCode() *string {
	return s.Code
}

func (s *MonthBillSplitGetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MonthBillSplitGetResponseBody) GetModule() []*MonthBillSplitGetResponseBodyModule {
	return s.Module
}

func (s *MonthBillSplitGetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MonthBillSplitGetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MonthBillSplitGetResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MonthBillSplitGetResponseBody) SetCode(v string) *MonthBillSplitGetResponseBody {
	s.Code = &v
	return s
}

func (s *MonthBillSplitGetResponseBody) SetMessage(v string) *MonthBillSplitGetResponseBody {
	s.Message = &v
	return s
}

func (s *MonthBillSplitGetResponseBody) SetModule(v []*MonthBillSplitGetResponseBodyModule) *MonthBillSplitGetResponseBody {
	s.Module = v
	return s
}

func (s *MonthBillSplitGetResponseBody) SetRequestId(v string) *MonthBillSplitGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *MonthBillSplitGetResponseBody) SetSuccess(v bool) *MonthBillSplitGetResponseBody {
	s.Success = &v
	return s
}

func (s *MonthBillSplitGetResponseBody) SetTraceId(v string) *MonthBillSplitGetResponseBody {
	s.TraceId = &v
	return s
}

func (s *MonthBillSplitGetResponseBody) Validate() error {
	return dara.Validate(s)
}

type MonthBillSplitGetResponseBodyModule struct {
	// example:
	//
	// invoice_third_part_id123
	BillSplitKey *string `json:"bill_split_key,omitempty" xml:"bill_split_key,omitempty"`
	// example:
	//
	// 2023-12-31
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// example:
	//
	// 2023-12-01
	StartDate *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// example:
	//
	// https://alibtrip-oss.oss-cn-hangzhou.aliyuncs.com/********
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MonthBillSplitGetResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MonthBillSplitGetResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MonthBillSplitGetResponseBodyModule) GetBillSplitKey() *string {
	return s.BillSplitKey
}

func (s *MonthBillSplitGetResponseBodyModule) GetEndDate() *string {
	return s.EndDate
}

func (s *MonthBillSplitGetResponseBodyModule) GetStartDate() *string {
	return s.StartDate
}

func (s *MonthBillSplitGetResponseBodyModule) GetUrl() *string {
	return s.Url
}

func (s *MonthBillSplitGetResponseBodyModule) SetBillSplitKey(v string) *MonthBillSplitGetResponseBodyModule {
	s.BillSplitKey = &v
	return s
}

func (s *MonthBillSplitGetResponseBodyModule) SetEndDate(v string) *MonthBillSplitGetResponseBodyModule {
	s.EndDate = &v
	return s
}

func (s *MonthBillSplitGetResponseBodyModule) SetStartDate(v string) *MonthBillSplitGetResponseBodyModule {
	s.StartDate = &v
	return s
}

func (s *MonthBillSplitGetResponseBodyModule) SetUrl(v string) *MonthBillSplitGetResponseBodyModule {
	s.Url = &v
	return s
}

func (s *MonthBillSplitGetResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
