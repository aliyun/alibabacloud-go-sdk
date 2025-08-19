// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id2MetaVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Id2MetaVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id2MetaVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Id2MetaVerifyResponseBodyResultObject) *Id2MetaVerifyResponseBody
	GetResultObject() *Id2MetaVerifyResponseBodyResultObject
}

type Id2MetaVerifyResponseBody struct {
	// Return code, **200*	- indicates successful API response.
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
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *Id2MetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id2MetaVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id2MetaVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id2MetaVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id2MetaVerifyResponseBody) GetResultObject() *Id2MetaVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Id2MetaVerifyResponseBody) SetCode(v string) *Id2MetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetMessage(v string) *Id2MetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetRequestId(v string) *Id2MetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaVerifyResponseBody) SetResultObject(v *Id2MetaVerifyResponseBodyResultObject) *Id2MetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Id2MetaVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type Id2MetaVerifyResponseBodyResultObject struct {
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
}

func (s Id2MetaVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Id2MetaVerifyResponseBodyResultObject) SetBizCode(v string) *Id2MetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Id2MetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
