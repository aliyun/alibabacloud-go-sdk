// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddressVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *AddressVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddressVerifyIntlResponseBody
	GetRequestId() *string
	SetResultObject(v *AddressVerifyIntlResponseBodyResultObject) *AddressVerifyIntlResponseBody
	GetResultObject() *AddressVerifyIntlResponseBodyResultObject
}

type AddressVerifyIntlResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 86C40EC3-5940-5F47-995C-BFE90B70E540
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
	ResultObject *AddressVerifyIntlResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s AddressVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddressVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *AddressVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddressVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddressVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddressVerifyIntlResponseBody) GetResultObject() *AddressVerifyIntlResponseBodyResultObject {
	return s.ResultObject
}

func (s *AddressVerifyIntlResponseBody) SetCode(v string) *AddressVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *AddressVerifyIntlResponseBody) SetMessage(v string) *AddressVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *AddressVerifyIntlResponseBody) SetRequestId(v string) *AddressVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddressVerifyIntlResponseBody) SetResultObject(v *AddressVerifyIntlResponseBodyResultObject) *AddressVerifyIntlResponseBody {
	s.ResultObject = v
	return s
}

func (s *AddressVerifyIntlResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddressVerifyIntlResponseBodyResultObject struct {
	// The address verification details.
	//
	// example:
	//
	// 0-3000
	AddressInfo *string `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty"`
	// The telecommunications service provider name. Valid values:
	//
	// - CMCC: China Mobile
	//
	// - CTCC: China Telecom
	//
	// - CUCC: China Unicom.
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The verification result. Valid values:
	//
	// - Y: The verified address is within 10 km of the residential address.
	//
	// - N: The verified address is more than 10 km from the residential address.
	//
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The description of the verification result.
	//
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// The unique identifier of the verification request.
	//
	// example:
	//
	// hksb7ba1b28130d24e015d69********
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s AddressVerifyIntlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s AddressVerifyIntlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *AddressVerifyIntlResponseBodyResultObject) GetAddressInfo() *string {
	return s.AddressInfo
}

func (s *AddressVerifyIntlResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *AddressVerifyIntlResponseBodyResultObject) GetPassed() *string {
	return s.Passed
}

func (s *AddressVerifyIntlResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *AddressVerifyIntlResponseBodyResultObject) GetTransactionId() *string {
	return s.TransactionId
}

func (s *AddressVerifyIntlResponseBodyResultObject) SetAddressInfo(v string) *AddressVerifyIntlResponseBodyResultObject {
	s.AddressInfo = &v
	return s
}

func (s *AddressVerifyIntlResponseBodyResultObject) SetIspName(v string) *AddressVerifyIntlResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *AddressVerifyIntlResponseBodyResultObject) SetPassed(v string) *AddressVerifyIntlResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *AddressVerifyIntlResponseBodyResultObject) SetSubCode(v string) *AddressVerifyIntlResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *AddressVerifyIntlResponseBodyResultObject) SetTransactionId(v string) *AddressVerifyIntlResponseBodyResultObject {
	s.TransactionId = &v
	return s
}

func (s *AddressVerifyIntlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
