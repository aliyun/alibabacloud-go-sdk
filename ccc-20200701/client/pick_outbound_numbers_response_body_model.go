// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPickOutboundNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PickOutboundNumbersResponseBody
	GetCode() *string
	SetData(v []*PickOutboundNumbersResponseBodyData) *PickOutboundNumbersResponseBody
	GetData() []*PickOutboundNumbersResponseBodyData
	SetHttpStatusCode(v int32) *PickOutboundNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PickOutboundNumbersResponseBody
	GetMessage() *string
	SetRequestId(v string) *PickOutboundNumbersResponseBody
	GetRequestId() *string
}

type PickOutboundNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*PickOutboundNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PickOutboundNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PickOutboundNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *PickOutboundNumbersResponseBody) GetData() []*PickOutboundNumbersResponseBodyData {
	return s.Data
}

func (s *PickOutboundNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PickOutboundNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PickOutboundNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PickOutboundNumbersResponseBody) SetCode(v string) *PickOutboundNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *PickOutboundNumbersResponseBody) SetData(v []*PickOutboundNumbersResponseBodyData) *PickOutboundNumbersResponseBody {
	s.Data = v
	return s
}

func (s *PickOutboundNumbersResponseBody) SetHttpStatusCode(v int32) *PickOutboundNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PickOutboundNumbersResponseBody) SetMessage(v string) *PickOutboundNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *PickOutboundNumbersResponseBody) SetRequestId(v string) *PickOutboundNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *PickOutboundNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}

type PickOutboundNumbersResponseBodyData struct {
	Callee *PickOutboundNumbersResponseBodyDataCallee `json:"Callee,omitempty" xml:"Callee,omitempty" type:"Struct"`
	Caller *PickOutboundNumbersResponseBodyDataCaller `json:"Caller,omitempty" xml:"Caller,omitempty" type:"Struct"`
}

func (s PickOutboundNumbersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PickOutboundNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponseBodyData) GetCallee() *PickOutboundNumbersResponseBodyDataCallee {
	return s.Callee
}

func (s *PickOutboundNumbersResponseBodyData) GetCaller() *PickOutboundNumbersResponseBodyDataCaller {
	return s.Caller
}

func (s *PickOutboundNumbersResponseBodyData) SetCallee(v *PickOutboundNumbersResponseBodyDataCallee) *PickOutboundNumbersResponseBodyData {
	s.Callee = v
	return s
}

func (s *PickOutboundNumbersResponseBodyData) SetCaller(v *PickOutboundNumbersResponseBodyDataCaller) *PickOutboundNumbersResponseBodyData {
	s.Caller = v
	return s
}

func (s *PickOutboundNumbersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type PickOutboundNumbersResponseBodyDataCallee struct {
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 1388888****
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s PickOutboundNumbersResponseBodyDataCallee) String() string {
	return dara.Prettify(s)
}

func (s PickOutboundNumbersResponseBodyDataCallee) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponseBodyDataCallee) GetCity() *string {
	return s.City
}

func (s *PickOutboundNumbersResponseBodyDataCallee) GetNumber() *string {
	return s.Number
}

func (s *PickOutboundNumbersResponseBodyDataCallee) GetProvince() *string {
	return s.Province
}

func (s *PickOutboundNumbersResponseBodyDataCallee) SetCity(v string) *PickOutboundNumbersResponseBodyDataCallee {
	s.City = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCallee) SetNumber(v string) *PickOutboundNumbersResponseBodyDataCallee {
	s.Number = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCallee) SetProvince(v string) *PickOutboundNumbersResponseBodyDataCallee {
	s.Province = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCallee) Validate() error {
	return dara.Validate(s)
}

type PickOutboundNumbersResponseBodyDataCaller struct {
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 0108989****
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s PickOutboundNumbersResponseBodyDataCaller) String() string {
	return dara.Prettify(s)
}

func (s PickOutboundNumbersResponseBodyDataCaller) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponseBodyDataCaller) GetCity() *string {
	return s.City
}

func (s *PickOutboundNumbersResponseBodyDataCaller) GetNumber() *string {
	return s.Number
}

func (s *PickOutboundNumbersResponseBodyDataCaller) GetProvince() *string {
	return s.Province
}

func (s *PickOutboundNumbersResponseBodyDataCaller) SetCity(v string) *PickOutboundNumbersResponseBodyDataCaller {
	s.City = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCaller) SetNumber(v string) *PickOutboundNumbersResponseBodyDataCaller {
	s.Number = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCaller) SetProvince(v string) *PickOutboundNumbersResponseBodyDataCaller {
	s.Province = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCaller) Validate() error {
	return dara.Validate(s)
}
