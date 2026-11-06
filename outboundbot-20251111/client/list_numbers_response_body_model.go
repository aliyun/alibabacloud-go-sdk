// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNumbersResponseBody
	GetCode() *string
	SetData(v *ListNumbersResponseBodyData) *ListNumbersResponseBody
	GetData() *ListNumbersResponseBodyData
	SetHttpStatusCode(v int32) *ListNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListNumbersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListNumbersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListNumbersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNumbersResponseBody
	GetSuccess() *bool
}

type ListNumbersResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *ListNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNumbersResponseBody) GetData() *ListNumbersResponseBodyData {
	return s.Data
}

func (s *ListNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNumbersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNumbersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNumbersResponseBody) SetCode(v string) *ListNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListNumbersResponseBody) SetData(v *ListNumbersResponseBodyData) *ListNumbersResponseBody {
	s.Data = v
	return s
}

func (s *ListNumbersResponseBody) SetHttpStatusCode(v int32) *ListNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNumbersResponseBody) SetMessage(v string) *ListNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListNumbersResponseBody) SetParams(v []*string) *ListNumbersResponseBody {
	s.Params = v
	return s
}

func (s *ListNumbersResponseBody) SetRequestId(v string) *ListNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNumbersResponseBody) SetSuccess(v bool) *ListNumbersResponseBody {
	s.Success = &v
	return s
}

func (s *ListNumbersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNumbersResponseBodyData struct {
	// The list of calling numbers.
	Numbers []*ListNumbersResponseBodyDataNumbers `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNumbersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNumbersResponseBodyData) GetNumbers() []*ListNumbersResponseBodyDataNumbers {
	return s.Numbers
}

func (s *ListNumbersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNumbersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNumbersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNumbersResponseBodyData) SetNumbers(v []*ListNumbersResponseBodyDataNumbers) *ListNumbersResponseBodyData {
	s.Numbers = v
	return s
}

func (s *ListNumbersResponseBodyData) SetPageNumber(v int32) *ListNumbersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListNumbersResponseBodyData) SetPageSize(v int32) *ListNumbersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListNumbersResponseBodyData) SetTotalCount(v int32) *ListNumbersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListNumbersResponseBodyData) Validate() error {
	if s.Numbers != nil {
		for _, item := range s.Numbers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNumbersResponseBodyDataNumbers struct {
	// The phone number.
	//
	// example:
	//
	// 057100001234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The SIP trunk name.
	//
	// example:
	//
	// trunk_example
	SipTrunkName *string `json:"SipTrunkName,omitempty" xml:"SipTrunkName,omitempty"`
}

func (s ListNumbersResponseBodyDataNumbers) String() string {
	return dara.Prettify(s)
}

func (s ListNumbersResponseBodyDataNumbers) GoString() string {
	return s.String()
}

func (s *ListNumbersResponseBodyDataNumbers) GetNumber() *string {
	return s.Number
}

func (s *ListNumbersResponseBodyDataNumbers) GetSipTrunkName() *string {
	return s.SipTrunkName
}

func (s *ListNumbersResponseBodyDataNumbers) SetNumber(v string) *ListNumbersResponseBodyDataNumbers {
	s.Number = &v
	return s
}

func (s *ListNumbersResponseBodyDataNumbers) SetSipTrunkName(v string) *ListNumbersResponseBodyDataNumbers {
	s.SipTrunkName = &v
	return s
}

func (s *ListNumbersResponseBodyDataNumbers) Validate() error {
	return dara.Validate(s)
}
