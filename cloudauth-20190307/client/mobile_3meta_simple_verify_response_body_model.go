// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaSimpleVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Mobile3MetaSimpleVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Mobile3MetaSimpleVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Mobile3MetaSimpleVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Mobile3MetaSimpleVerifyResponseBodyResultObject) *Mobile3MetaSimpleVerifyResponseBody
	GetResultObject() *Mobile3MetaSimpleVerifyResponseBodyResultObject
}

type Mobile3MetaSimpleVerifyResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result.
	ResultObject *Mobile3MetaSimpleVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile3MetaSimpleVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Mobile3MetaSimpleVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Mobile3MetaSimpleVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Mobile3MetaSimpleVerifyResponseBody) GetResultObject() *Mobile3MetaSimpleVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetCode(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetMessage(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetRequestId(v string) *Mobile3MetaSimpleVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) SetResultObject(v *Mobile3MetaSimpleVerifyResponseBodyResultObject) *Mobile3MetaSimpleVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Mobile3MetaSimpleVerifyResponseBodyResultObject struct {
	// Verification result code:
	//
	// - **1**: Verification consistent.
	//
	// - **2**: Verification inconsistent.
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
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s Mobile3MetaSimpleVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile3MetaSimpleVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) SetIspName(v string) *Mobile3MetaSimpleVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
