// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id3MetaVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *Id3MetaVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id3MetaVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *Id3MetaVerifyResponseBodyResultObject) *Id3MetaVerifyResponseBody
	GetResultObject() *Id3MetaVerifyResponseBodyResultObject
}

type Id3MetaVerifyResponseBody struct {
	// Return code: 200 indicates success, others indicate failure.
	//
	// **Important**
	//
	// - This parameter indicates whether the interface was called correctly. For detailed return code explanations, please refer to the error codes.
	//
	// - Check the business verification result through the fields in `ResultObject`.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Interface call return message.
	//
	// **Important**
	//
	// This parameter only indicates whether there was an exception with the interface.
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
	ResultObject *Id3MetaVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id3MetaVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id3MetaVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id3MetaVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id3MetaVerifyResponseBody) GetResultObject() *Id3MetaVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *Id3MetaVerifyResponseBody) SetCode(v string) *Id3MetaVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *Id3MetaVerifyResponseBody) SetMessage(v string) *Id3MetaVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *Id3MetaVerifyResponseBody) SetRequestId(v string) *Id3MetaVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id3MetaVerifyResponseBody) SetResultObject(v *Id3MetaVerifyResponseBodyResultObject) *Id3MetaVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *Id3MetaVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Id3MetaVerifyResponseBodyResultObject struct {
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
	// Image comparison score.
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

func (s Id3MetaVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Id3MetaVerifyResponseBodyResultObject) GetFaceDetail() *string {
	return s.FaceDetail
}

func (s *Id3MetaVerifyResponseBodyResultObject) SetBizCode(v string) *Id3MetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Id3MetaVerifyResponseBodyResultObject) SetFaceDetail(v string) *Id3MetaVerifyResponseBodyResultObject {
	s.FaceDetail = &v
	return s
}

func (s *Id3MetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
