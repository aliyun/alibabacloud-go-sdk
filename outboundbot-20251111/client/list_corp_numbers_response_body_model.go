// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCorpNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCorpNumbersResponseBody
	GetCode() *string
	SetData(v *ListCorpNumbersResponseBodyData) *ListCorpNumbersResponseBody
	GetData() *ListCorpNumbersResponseBodyData
	SetHttpStatusCode(v int32) *ListCorpNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCorpNumbersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCorpNumbersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCorpNumbersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCorpNumbersResponseBody
	GetSuccess() *bool
}

type ListCorpNumbersResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListCorpNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
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
	// 0956D5DA-0978-5DC9-94B0-C68527DA7475
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCorpNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCorpNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCorpNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCorpNumbersResponseBody) GetData() *ListCorpNumbersResponseBodyData {
	return s.Data
}

func (s *ListCorpNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCorpNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCorpNumbersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCorpNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCorpNumbersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCorpNumbersResponseBody) SetCode(v string) *ListCorpNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListCorpNumbersResponseBody) SetData(v *ListCorpNumbersResponseBodyData) *ListCorpNumbersResponseBody {
	s.Data = v
	return s
}

func (s *ListCorpNumbersResponseBody) SetHttpStatusCode(v int32) *ListCorpNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCorpNumbersResponseBody) SetMessage(v string) *ListCorpNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListCorpNumbersResponseBody) SetParams(v []*string) *ListCorpNumbersResponseBody {
	s.Params = v
	return s
}

func (s *ListCorpNumbersResponseBody) SetRequestId(v string) *ListCorpNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCorpNumbersResponseBody) SetSuccess(v bool) *ListCorpNumbersResponseBody {
	s.Success = &v
	return s
}

func (s *ListCorpNumbersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCorpNumbersResponseBodyData struct {
	// The list of tenant numbers.
	CorpNumbers []*ListCorpNumbersResponseBodyDataCorpNumbers `json:"CorpNumbers,omitempty" xml:"CorpNumbers,omitempty" type:"Repeated"`
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
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCorpNumbersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCorpNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCorpNumbersResponseBodyData) GetCorpNumbers() []*ListCorpNumbersResponseBodyDataCorpNumbers {
	return s.CorpNumbers
}

func (s *ListCorpNumbersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCorpNumbersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCorpNumbersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCorpNumbersResponseBodyData) SetCorpNumbers(v []*ListCorpNumbersResponseBodyDataCorpNumbers) *ListCorpNumbersResponseBodyData {
	s.CorpNumbers = v
	return s
}

func (s *ListCorpNumbersResponseBodyData) SetPageNumber(v int32) *ListCorpNumbersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCorpNumbersResponseBodyData) SetPageSize(v int32) *ListCorpNumbersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCorpNumbersResponseBodyData) SetTotalCount(v int32) *ListCorpNumbersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCorpNumbersResponseBodyData) Validate() error {
	if s.CorpNumbers != nil {
		for _, item := range s.CorpNumbers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCorpNumbersResponseBodyDataCorpNumbers struct {
	// The list of instance names associated with the number.
	InstanceNames []*string `json:"InstanceNames,omitempty" xml:"InstanceNames,omitempty" type:"Repeated"`
	// The phone number.
	//
	// example:
	//
	// 057100001234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The SIP trunk ID.
	//
	// example:
	//
	// 159d4b8361804ebda21242216a8981c2
	SipTrunkId *string `json:"SipTrunkId,omitempty" xml:"SipTrunkId,omitempty"`
	// The SIP trunk name.
	//
	// example:
	//
	// trunk_1
	SipTrunkName *string `json:"SipTrunkName,omitempty" xml:"SipTrunkName,omitempty"`
}

func (s ListCorpNumbersResponseBodyDataCorpNumbers) String() string {
	return dara.Prettify(s)
}

func (s ListCorpNumbersResponseBodyDataCorpNumbers) GoString() string {
	return s.String()
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) GetInstanceNames() []*string {
	return s.InstanceNames
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) GetNumber() *string {
	return s.Number
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) GetSipTrunkId() *string {
	return s.SipTrunkId
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) GetSipTrunkName() *string {
	return s.SipTrunkName
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) SetInstanceNames(v []*string) *ListCorpNumbersResponseBodyDataCorpNumbers {
	s.InstanceNames = v
	return s
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) SetNumber(v string) *ListCorpNumbersResponseBodyDataCorpNumbers {
	s.Number = &v
	return s
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) SetSipTrunkId(v string) *ListCorpNumbersResponseBodyDataCorpNumbers {
	s.SipTrunkId = &v
	return s
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) SetSipTrunkName(v string) *ListCorpNumbersResponseBodyDataCorpNumbers {
	s.SipTrunkName = &v
	return s
}

func (s *ListCorpNumbersResponseBodyDataCorpNumbers) Validate() error {
	return dara.Validate(s)
}
