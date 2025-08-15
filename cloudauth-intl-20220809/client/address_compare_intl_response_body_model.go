// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressCompareIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddressCompareIntlResponseBody
	GetCode() *string
	SetMessage(v string) *AddressCompareIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddressCompareIntlResponseBody
	GetRequestId() *string
	SetResult(v *AddressCompareIntlResponseBodyResult) *AddressCompareIntlResponseBody
	GetResult() *AddressCompareIntlResponseBodyResult
}

type AddressCompareIntlResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C40EC3-5940-5F47-995C-BFE90B70E540
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddressCompareIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s AddressCompareIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddressCompareIntlResponseBody) GoString() string {
	return s.String()
}

func (s *AddressCompareIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddressCompareIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddressCompareIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddressCompareIntlResponseBody) GetResult() *AddressCompareIntlResponseBodyResult {
	return s.Result
}

func (s *AddressCompareIntlResponseBody) SetCode(v string) *AddressCompareIntlResponseBody {
	s.Code = &v
	return s
}

func (s *AddressCompareIntlResponseBody) SetMessage(v string) *AddressCompareIntlResponseBody {
	s.Message = &v
	return s
}

func (s *AddressCompareIntlResponseBody) SetRequestId(v string) *AddressCompareIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddressCompareIntlResponseBody) SetResult(v *AddressCompareIntlResponseBodyResult) *AddressCompareIntlResponseBody {
	s.Result = v
	return s
}

func (s *AddressCompareIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddressCompareIntlResponseBodyResult struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s AddressCompareIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddressCompareIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddressCompareIntlResponseBodyResult) GetData() *string {
	return s.Data
}

func (s *AddressCompareIntlResponseBodyResult) SetData(v string) *AddressCompareIntlResponseBodyResult {
	s.Data = &v
	return s
}

func (s *AddressCompareIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
