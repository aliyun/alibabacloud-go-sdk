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
	// Return code, **200*	- indicates a successful API response.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
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
	return dara.Validate(s)
}

type Mobile3MetaDetailStandardVerifyResponseBodyResultObject struct {
	// Verification result code:
	//
	// - **1**: Verification matches.
	//
	// - **2**: Verification does not match.
	//
	// - **3**: No record found.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// ISP name:
	//
	// - **CMCC**: China Mobile.
	//
	// - **CUCC**: China Unicom.
	//
	// - **CTCC**: China Telecom.
	//
	// - **CBCC**: China Broadcasting Network.
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// Detailed verification results:
	//
	// - 101: Passed, three elements are consistent.
	//
	// - 201: The phone number does not match the name and ID number.
	//
	// - 202: The phone number matches the name but does not match the ID number.
	//
	// - 203: The phone number does not match the name but matches the ID number.
	//
	// - 204: Other inconsistencies.
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
