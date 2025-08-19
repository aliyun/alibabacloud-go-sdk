// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyWithOCRResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id2MetaVerifyWithOCRResponseBody
	GetCode() *string
	SetMessage(v string) *Id2MetaVerifyWithOCRResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id2MetaVerifyWithOCRResponseBody
	GetRequestId() *string
	SetResultObject(v *Id2MetaVerifyWithOCRResponseBodyResultObject) *Id2MetaVerifyWithOCRResponseBody
	GetResultObject() *Id2MetaVerifyWithOCRResponseBodyResultObject
}

type Id2MetaVerifyWithOCRResponseBody struct {
	// Return code: 200 indicates success, any other value indicates failure.
	//
	// **Important**
	//
	// - This parameter indicates whether the API was called correctly. For detailed return code explanations, please refer to the error codes.
	//
	// - Check the business verification results through the fields in ResultObject.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// API call return message.
	//
	// **Important**
	//
	// This parameter only indicates if there was an exception with the API call.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF03****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object
	ResultObject *Id2MetaVerifyWithOCRResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id2MetaVerifyWithOCRResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyWithOCRResponseBody) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyWithOCRResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id2MetaVerifyWithOCRResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id2MetaVerifyWithOCRResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id2MetaVerifyWithOCRResponseBody) GetResultObject() *Id2MetaVerifyWithOCRResponseBodyResultObject {
	return s.ResultObject
}

func (s *Id2MetaVerifyWithOCRResponseBody) SetCode(v string) *Id2MetaVerifyWithOCRResponseBody {
	s.Code = &v
	return s
}

func (s *Id2MetaVerifyWithOCRResponseBody) SetMessage(v string) *Id2MetaVerifyWithOCRResponseBody {
	s.Message = &v
	return s
}

func (s *Id2MetaVerifyWithOCRResponseBody) SetRequestId(v string) *Id2MetaVerifyWithOCRResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id2MetaVerifyWithOCRResponseBody) SetResultObject(v *Id2MetaVerifyWithOCRResponseBodyResultObject) *Id2MetaVerifyWithOCRResponseBody {
	s.ResultObject = v
	return s
}

func (s *Id2MetaVerifyWithOCRResponseBody) Validate() error {
	return dara.Validate(s)
}

type Id2MetaVerifyWithOCRResponseBodyResultObject struct {
	// Identity verification result:
	//
	// - 1: Consistent
	//
	// - 2: Inconsistent
	//
	// - 3: No record found
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// {"address":"Zhejiang Province, Hangzhou City, Yu*****","birthDate":"19901226","certName":"Zhang San","certNo":"1234561990122*****","nationality":"Han","authority":"xxx Public Security Bureau","startDate":"20201130","endDate":"20301130"}
	//
	// example:
	//
	// OCR读取的身份证信息。
	CardInfo *string `json:"CardInfo,omitempty" xml:"CardInfo,omitempty"`
}

func (s Id2MetaVerifyWithOCRResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyWithOCRResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyWithOCRResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Id2MetaVerifyWithOCRResponseBodyResultObject) GetCardInfo() *string {
	return s.CardInfo
}

func (s *Id2MetaVerifyWithOCRResponseBodyResultObject) SetBizCode(v string) *Id2MetaVerifyWithOCRResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Id2MetaVerifyWithOCRResponseBodyResultObject) SetCardInfo(v string) *Id2MetaVerifyWithOCRResponseBodyResultObject {
	s.CardInfo = &v
	return s
}

func (s *Id2MetaVerifyWithOCRResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
