// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaSimpleStandardVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Mobile3MetaSimpleStandardVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Mobile3MetaSimpleStandardVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Mobile3MetaSimpleStandardVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) *Mobile3MetaSimpleStandardVerifyResponseBody
	GetResultObject() *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject
}

type Mobile3MetaSimpleStandardVerifyResponseBody struct {
	// Return code: 200 for success, others for failure
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
	// ID of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Mobile3MetaSimpleStandardVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaSimpleStandardVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) GetResultObject() *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) SetCode(v string) *Mobile3MetaSimpleStandardVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) SetMessage(v string) *Mobile3MetaSimpleStandardVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) SetRequestId(v string) *Mobile3MetaSimpleStandardVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) SetResultObject(v *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) *Mobile3MetaSimpleStandardVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Mobile3MetaSimpleStandardVerifyResponseBodyResultObject struct {
	// Verification result:
	//
	// - 1: Consistent (billable)
	//
	// - 2: Inconsistent (billable)
	//
	// - 3: No record found (non-billable)
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Operator name:
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
}

func (s Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) GetIspName() *string {
	return s.IspName
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) SetBizCode(v string) *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) SetIspName(v string) *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject {
	s.IspName = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
