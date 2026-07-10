// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaDetailVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Mobile3MetaDetailVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Mobile3MetaDetailVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Mobile3MetaDetailVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Mobile3MetaDetailVerifyResponseBodyResultObject) *Mobile3MetaDetailVerifyResponseBody
	GetResultObject() *Mobile3MetaDetailVerifyResponseBodyResultObject
}

type Mobile3MetaDetailVerifyResponseBody struct {
	// The response code. **200*	- indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
	ResultObject *Mobile3MetaDetailVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile3MetaDetailVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Mobile3MetaDetailVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Mobile3MetaDetailVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Mobile3MetaDetailVerifyResponseBody) GetResultObject() *Mobile3MetaDetailVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetCode(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetMessage(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetRequestId(v string) *Mobile3MetaDetailVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) SetResultObject(v *Mobile3MetaDetailVerifyResponseBodyResultObject) *Mobile3MetaDetailVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Mobile3MetaDetailVerifyResponseBodyResultObject struct {
	// The verification result code. Valid values:
	//
	// - **1**: Consistent.
	//
	// - **2**: Inconsistent.
	//
	// - **3**: No record found.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The carrier name. Valid values:
	//
	// - **CMCC**: China Mobile.
	//
	// - **CUCC**: China Unicom.
	//
	// - **CTCC**: China Telecom.
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The detailed verification result. Valid values:
	//
	// - **101**: Verification passed.
	//
	// - **201**: The phone number is inconsistent with the name, and the phone number is inconsistent with the ID card number.
	//
	// - **202**: The phone number is consistent with the name, but the phone number is inconsistent with the ID card number.
	//
	// - **203**: The phone number is consistent with the ID card number, but the phone number is inconsistent with the name.
	//
	// - **204**: Other inconsistency.
	//
	// - **301**: No record found.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Mobile3MetaDetailVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetIspName(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) SetSubCode(v string) *Mobile3MetaDetailVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
