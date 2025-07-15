// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BankMetaVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *BankMetaVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *BankMetaVerifyIntlResponseBody
	GetRequestId() *string
	SetResultObject(v *BankMetaVerifyIntlResponseBodyResultObject) *BankMetaVerifyIntlResponseBody
	GetResultObject() *BankMetaVerifyIntlResponseBodyResultObject
}

type BankMetaVerifyIntlResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4EB35****87EBA1
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *BankMetaVerifyIntlResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s BankMetaVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *BankMetaVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BankMetaVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BankMetaVerifyIntlResponseBody) GetResultObject() *BankMetaVerifyIntlResponseBodyResultObject {
	return s.ResultObject
}

func (s *BankMetaVerifyIntlResponseBody) SetCode(v string) *BankMetaVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetMessage(v string) *BankMetaVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetRequestId(v string) *BankMetaVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) SetResultObject(v *BankMetaVerifyIntlResponseBodyResultObject) *BankMetaVerifyIntlResponseBody {
	s.ResultObject = v
	return s
}

func (s *BankMetaVerifyIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type BankMetaVerifyIntlResponseBodyResultObject struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s BankMetaVerifyIntlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyIntlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) SetBizCode(v string) *BankMetaVerifyIntlResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) SetSubCode(v string) *BankMetaVerifyIntlResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *BankMetaVerifyIntlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
