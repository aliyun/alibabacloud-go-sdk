// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressVerifyV2IntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddressVerifyV2IntlResponseBody
	GetCode() *string
	SetMessage(v string) *AddressVerifyV2IntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddressVerifyV2IntlResponseBody
	GetRequestId() *string
	SetResult(v *AddressVerifyV2IntlResponseBodyResult) *AddressVerifyV2IntlResponseBody
	GetResult() *AddressVerifyV2IntlResponseBodyResult
}

type AddressVerifyV2IntlResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F971622-38C0-5F56-B2EC-315367979B4F
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddressVerifyV2IntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s AddressVerifyV2IntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddressVerifyV2IntlResponseBody) GoString() string {
	return s.String()
}

func (s *AddressVerifyV2IntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddressVerifyV2IntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddressVerifyV2IntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddressVerifyV2IntlResponseBody) GetResult() *AddressVerifyV2IntlResponseBodyResult {
	return s.Result
}

func (s *AddressVerifyV2IntlResponseBody) SetCode(v string) *AddressVerifyV2IntlResponseBody {
	s.Code = &v
	return s
}

func (s *AddressVerifyV2IntlResponseBody) SetMessage(v string) *AddressVerifyV2IntlResponseBody {
	s.Message = &v
	return s
}

func (s *AddressVerifyV2IntlResponseBody) SetRequestId(v string) *AddressVerifyV2IntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddressVerifyV2IntlResponseBody) SetResult(v *AddressVerifyV2IntlResponseBodyResult) *AddressVerifyV2IntlResponseBody {
	s.Result = v
	return s
}

func (s *AddressVerifyV2IntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddressVerifyV2IntlResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {
	//
	//   "distanceRange": "0-3000",
	//
	//   "ispName": "CTCC",
	//
	//   "phoneStatus": "1"
	//
	// }
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// hksb7ba1b28130d24e015d69********
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s AddressVerifyV2IntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddressVerifyV2IntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddressVerifyV2IntlResponseBodyResult) GetBizCode() *string {
	return s.BizCode
}

func (s *AddressVerifyV2IntlResponseBodyResult) GetDetail() *string {
	return s.Detail
}

func (s *AddressVerifyV2IntlResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *AddressVerifyV2IntlResponseBodyResult) SetBizCode(v string) *AddressVerifyV2IntlResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *AddressVerifyV2IntlResponseBodyResult) SetDetail(v string) *AddressVerifyV2IntlResponseBodyResult {
	s.Detail = &v
	return s
}

func (s *AddressVerifyV2IntlResponseBodyResult) SetTransactionId(v string) *AddressVerifyV2IntlResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *AddressVerifyV2IntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
