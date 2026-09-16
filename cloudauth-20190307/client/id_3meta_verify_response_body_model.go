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
	// The response code. 200 indicates success. Other values indicate failure.
	//
	// **Important**
	//
	// - This parameter indicates whether the API call is successful. For more information about return codes, see error codes.
	//
	// - Check the business verification result in the fields of ResultObject.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// **Important**
	//
	// This parameter only indicates whether the API call is abnormal.
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
	// - 1: verification is consistent.
	//
	// - 2: verification is inconsistent.
	//
	// - 3: no record found.
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
	// The authoritative source verification details. Valid values:
	//
	// - 101: authentication passed.
	//
	// - 201: authentication failed. The name does not match the ID card number.
	//
	// - 202: authentication failed. The person is suspected to be the ID holder.
	//
	// - 203: authentication failed. No photo exists in the database.
	//
	// - 204: authentication failed. The person is not the same individual.
	//
	// - 205: authentication failed. Modeling of the image to be compared failed.
	//
	// - 206: authentication failed. The image format is incorrect.
	//
	// - 207: authentication failed. The uploaded image is too small. Upload a new image.
	//
	// - 208: authentication failed. The quality of the uploaded portrait photo is poor. Upload a new photo.
	//
	// - 301: no record found. The ID number does not exist in the database.
	//
	// - 302: no record found. Verification cannot be performed.
	//
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
