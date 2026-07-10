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
	// The response code. A value of 200 indicates success. Other values indicate failure.
	//
	// **Important**
	//
	// - This parameter indicates only whether the API call is successful. For more information about return codes, see error codes.
	//
	// - Check the fields in ResultObject for the business verification result.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message of the API call.
	//
	// **Important**
	//
	// This parameter indicates only whether the API call is abnormal.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF03****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result object.
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
	// The identity verification result. Valid values:
	//
	// - 1: Consistent.
	//
	// - 2: Inconsistent.
	//
	// - 3: No record found.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The face comparison score.
	//
	// example:
	//
	// {
	//
	//  "verifyScore": 50.28594166529785
	//
	// }
	FaceDetail *string `json:"FaceDetail,omitempty" xml:"FaceDetail,omitempty"`
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
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

func (s *Id3MetaVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *Id3MetaVerifyResponseBodyResultObject) SetBizCode(v string) *Id3MetaVerifyResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Id3MetaVerifyResponseBodyResultObject) SetFaceDetail(v string) *Id3MetaVerifyResponseBodyResultObject {
	s.FaceDetail = &v
	return s
}

func (s *Id3MetaVerifyResponseBodyResultObject) SetSubCode(v string) *Id3MetaVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *Id3MetaVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
