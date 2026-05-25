// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNumberMccMncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeNumberMccMncResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeNumberMccMncResponseBody
	GetCode() *string
	SetData(v *DescribeNumberMccMncResponseBodyData) *DescribeNumberMccMncResponseBody
	GetData() *DescribeNumberMccMncResponseBodyData
	SetMessage(v string) *DescribeNumberMccMncResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeNumberMccMncResponseBody
	GetRequestId() *string
}

type DescribeNumberMccMncResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeNumberMccMncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNumberMccMncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberMccMncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNumberMccMncResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeNumberMccMncResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeNumberMccMncResponseBody) GetData() *DescribeNumberMccMncResponseBodyData {
	return s.Data
}

func (s *DescribeNumberMccMncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeNumberMccMncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNumberMccMncResponseBody) SetAccessDeniedDetail(v string) *DescribeNumberMccMncResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeNumberMccMncResponseBody) SetCode(v string) *DescribeNumberMccMncResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNumberMccMncResponseBody) SetData(v *DescribeNumberMccMncResponseBodyData) *DescribeNumberMccMncResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNumberMccMncResponseBody) SetMessage(v string) *DescribeNumberMccMncResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNumberMccMncResponseBody) SetRequestId(v string) *DescribeNumberMccMncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNumberMccMncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNumberMccMncResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	CountryIso3 *string `json:"CountryIso3,omitempty" xml:"CountryIso3,omitempty"`
	// example:
	//
	// 示例值示例值
	Mcc *string `json:"Mcc,omitempty" xml:"Mcc,omitempty"`
	// example:
	//
	// 示例值示例值
	Mnc *string `json:"Mnc,omitempty" xml:"Mnc,omitempty"`
	// example:
	//
	// true
	Ported *bool `json:"Ported,omitempty" xml:"Ported,omitempty"`
}

func (s DescribeNumberMccMncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberMccMncResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNumberMccMncResponseBodyData) GetCountryIso3() *string {
	return s.CountryIso3
}

func (s *DescribeNumberMccMncResponseBodyData) GetMcc() *string {
	return s.Mcc
}

func (s *DescribeNumberMccMncResponseBodyData) GetMnc() *string {
	return s.Mnc
}

func (s *DescribeNumberMccMncResponseBodyData) GetPorted() *bool {
	return s.Ported
}

func (s *DescribeNumberMccMncResponseBodyData) SetCountryIso3(v string) *DescribeNumberMccMncResponseBodyData {
	s.CountryIso3 = &v
	return s
}

func (s *DescribeNumberMccMncResponseBodyData) SetMcc(v string) *DescribeNumberMccMncResponseBodyData {
	s.Mcc = &v
	return s
}

func (s *DescribeNumberMccMncResponseBodyData) SetMnc(v string) *DescribeNumberMccMncResponseBodyData {
	s.Mnc = &v
	return s
}

func (s *DescribeNumberMccMncResponseBodyData) SetPorted(v bool) *DescribeNumberMccMncResponseBodyData {
	s.Ported = &v
	return s
}

func (s *DescribeNumberMccMncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
