// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile2MetaVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Mobile2MetaVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Mobile2MetaVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Mobile2MetaVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Mobile2MetaVerifyResponseBodyResultObject) *Mobile2MetaVerifyResponseBody
	GetResultObject() *Mobile2MetaVerifyResponseBodyResultObject
}

type Mobile2MetaVerifyResponseBody struct {
	// Return code: 200 for success, others for failure.
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
	// Request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
	ResultObject *Mobile2MetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile2MetaVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Mobile2MetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile2MetaVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Mobile2MetaVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Mobile2MetaVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Mobile2MetaVerifyResponseBody) GetResultObject() *Mobile2MetaVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Mobile2MetaVerifyResponseBody) SetCode(v string) *Mobile2MetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile2MetaVerifyResponseBody) SetMessage(v string) *Mobile2MetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile2MetaVerifyResponseBody) SetRequestId(v string) *Mobile2MetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile2MetaVerifyResponseBody) SetResultObject(v *Mobile2MetaVerifyResponseBodyResultObject) *Mobile2MetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Mobile2MetaVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type Mobile2MetaVerifyResponseBodyResultObject struct {
	// Verification result:
	//
	// - 1: Consistent verification
	//
	// - 2: Inconsistent verification
	//
	// - 3: No record found
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Operator name:
	//
	// - CMCC: China Mobile
	//
	// - CUCC: China Unicom
	//
	// - CTCC: China Telecom
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s Mobile2MetaVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Mobile2MetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile2MetaVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Mobile2MetaVerifyResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *Mobile2MetaVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile2MetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile2MetaVerifyResponseBodyResultObject) SetIspName(v string) *Mobile2MetaVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *Mobile2MetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
