// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaDetailStandardVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Mobile3MetaDetailStandardVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Mobile3MetaDetailStandardVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Mobile3MetaDetailStandardVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) *Mobile3MetaDetailStandardVerifyResponseBody
	GetResultObject() *Mobile3MetaDetailStandardVerifyResponseBodyResultObject
}

type Mobile3MetaDetailStandardVerifyResponseBody struct {
	// The return code. **200*	- indicates a successful response.
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
	// Id of the request
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
	ResultObject *Mobile3MetaDetailStandardVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile3MetaDetailStandardVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaDetailStandardVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) GetResultObject() *Mobile3MetaDetailStandardVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) SetCode(v string) *Mobile3MetaDetailStandardVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) SetMessage(v string) *Mobile3MetaDetailStandardVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) SetRequestId(v string) *Mobile3MetaDetailStandardVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) SetResultObject(v *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) *Mobile3MetaDetailStandardVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Mobile3MetaDetailStandardVerifyResponseBodyResultObject struct {
	// The verification result code. Valid values:
	//
	// - **1**: Verification is consistent.
	//
	// - **2**: Verification is inconsistent.
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
	// - **CBCC**: China Broadnet.
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The detailed verification result. Valid values:
	//
	// - 101: Verification passed. All three elements are consistent.
	//
	// - 201: The phone number is inconsistent with both the name and the ID card number.
	//
	// - 202: The phone number is consistent with the name but inconsistent with the ID card number.
	//
	// - 203: The phone number is inconsistent with the name but consistent with the ID card number.
	//
	// - 204: Other inconsistency.
	//
	// - 301: No record found.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Mobile3MetaDetailStandardVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaDetailStandardVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile3MetaDetailStandardVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) SetIspName(v string) *Mobile3MetaDetailStandardVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) SetSubCode(v string) *Mobile3MetaDetailStandardVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
