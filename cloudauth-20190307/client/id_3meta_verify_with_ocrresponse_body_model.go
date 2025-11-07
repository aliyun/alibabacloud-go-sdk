// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyWithOCRResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id3MetaVerifyWithOCRResponseBody
	GetCode() *string
	SetMessage(v string) *Id3MetaVerifyWithOCRResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id3MetaVerifyWithOCRResponseBody
	GetRequestId() *string
	SetResultObject(v *Id3MetaVerifyWithOCRResponseBodyResultObject) *Id3MetaVerifyWithOCRResponseBody
	GetResultObject() *Id3MetaVerifyWithOCRResponseBodyResultObject
}

type Id3MetaVerifyWithOCRResponseBody struct {
	// Return code: 200 indicates success, any other value indicates failure. **Important**
	//
	// - This parameter indicates whether the interface was called correctly. For a detailed explanation of return codes, please refer to the error codes.
	//
	// - The business verification result can be viewed through the fields in ResultObject.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Interface call return message. **Important*	- This parameter only indicates whether there was an exception with the interface.
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
	// Result object.
	ResultObject *Id3MetaVerifyWithOCRResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id3MetaVerifyWithOCRResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyWithOCRResponseBody) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyWithOCRResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id3MetaVerifyWithOCRResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id3MetaVerifyWithOCRResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id3MetaVerifyWithOCRResponseBody) GetResultObject() *Id3MetaVerifyWithOCRResponseBodyResultObject {
	return s.ResultObject
}

func (s *Id3MetaVerifyWithOCRResponseBody) SetCode(v string) *Id3MetaVerifyWithOCRResponseBody {
	s.Code = &v
	return s
}

func (s *Id3MetaVerifyWithOCRResponseBody) SetMessage(v string) *Id3MetaVerifyWithOCRResponseBody {
	s.Message = &v
	return s
}

func (s *Id3MetaVerifyWithOCRResponseBody) SetRequestId(v string) *Id3MetaVerifyWithOCRResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id3MetaVerifyWithOCRResponseBody) SetResultObject(v *Id3MetaVerifyWithOCRResponseBodyResultObject) *Id3MetaVerifyWithOCRResponseBody {
	s.ResultObject = v
	return s
}

func (s *Id3MetaVerifyWithOCRResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Id3MetaVerifyWithOCRResponseBodyResultObject struct {
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
	// Card ocr result.
	//
	// example:
	//
	// {"address":"浙江省杭州市余*****","birthDate":"19901226","certName":"张三","certNo":"1234561990122*****","nationality":"汉","authority":"xxx公安局","startDate":"20201130","endDate":"20301130"}
	CardInfo *string `json:"CardInfo,omitempty" xml:"CardInfo,omitempty"`
	// Face comparison score.
	//
	// example:
	//
	// {
	//
	//  "verifyScore": 50.28594166529785
	//
	// }
	FaceDetail *string `json:"FaceDetail,omitempty" xml:"FaceDetail,omitempty"`
}

func (s Id3MetaVerifyWithOCRResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyWithOCRResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyWithOCRResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Id3MetaVerifyWithOCRResponseBodyResultObject) GetCardInfo() *string {
	return s.CardInfo
}

func (s *Id3MetaVerifyWithOCRResponseBodyResultObject) GetFaceDetail() *string {
	return s.FaceDetail
}

func (s *Id3MetaVerifyWithOCRResponseBodyResultObject) SetBizCode(v string) *Id3MetaVerifyWithOCRResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Id3MetaVerifyWithOCRResponseBodyResultObject) SetCardInfo(v string) *Id3MetaVerifyWithOCRResponseBodyResultObject {
	s.CardInfo = &v
	return s
}

func (s *Id3MetaVerifyWithOCRResponseBodyResultObject) SetFaceDetail(v string) *Id3MetaVerifyWithOCRResponseBodyResultObject {
	s.FaceDetail = &v
	return s
}

func (s *Id3MetaVerifyWithOCRResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
