// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyPROResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *Id3MetaVerifyPROResponseBody
	GetCode() *string
	SetMessage(v string) *Id3MetaVerifyPROResponseBody
	GetMessage() *string
	SetRequestId(v string) *Id3MetaVerifyPROResponseBody
	GetRequestId() *string
	SetResultObject(v *Id3MetaVerifyPROResponseBodyResultObject) *Id3MetaVerifyPROResponseBody
	GetResultObject() *Id3MetaVerifyPROResponseBodyResultObject
}

type Id3MetaVerifyPROResponseBody struct {
	// The return code. A value of **200*	- indicates that the API call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned for the API call.
	//
	// 	Notice: This parameter only indicates whether an exception occurs during the API call.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FA2C773-47DB-4156-B1EE-5B047321A939
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	ResultObject *Id3MetaVerifyPROResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s Id3MetaVerifyPROResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyPROResponseBody) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyPROResponseBody) GetCode() *string {
	return s.Code
}

func (s *Id3MetaVerifyPROResponseBody) GetMessage() *string {
	return s.Message
}

func (s *Id3MetaVerifyPROResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Id3MetaVerifyPROResponseBody) GetResultObject() *Id3MetaVerifyPROResponseBodyResultObject {
	return s.ResultObject
}

func (s *Id3MetaVerifyPROResponseBody) SetCode(v string) *Id3MetaVerifyPROResponseBody {
	s.Code = &v
	return s
}

func (s *Id3MetaVerifyPROResponseBody) SetMessage(v string) *Id3MetaVerifyPROResponseBody {
	s.Message = &v
	return s
}

func (s *Id3MetaVerifyPROResponseBody) SetRequestId(v string) *Id3MetaVerifyPROResponseBody {
	s.RequestId = &v
	return s
}

func (s *Id3MetaVerifyPROResponseBody) SetResultObject(v *Id3MetaVerifyPROResponseBodyResultObject) *Id3MetaVerifyPROResponseBody {
	s.ResultObject = v
	return s
}

func (s *Id3MetaVerifyPROResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Id3MetaVerifyPROResponseBodyResultObject struct {
	// The verification result from the authoritative source. Valid values:
	//
	// - **1**: Consistent (billed).
	//
	// - **2**: Inconsistent (billed).
	//
	// - **3**: No record found (not billed).
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// - **verifyScore**: The face comparison score. Valid values: 0 to 1000. A higher score indicates a higher probability that the faces belong to the same person. A score of 700.0 or higher confirms that the faces belong to the same person.
	//
	// - **faceAttack**: Returned when liveness detection is enabled. This value does not affect the verification result.
	//
	// - **invokeChannel**: The identifier of the actual invocation channel. Valid values:
	//
	// 1: Authoritative source. 0: Comprehensive source.
	//
	// example:
	//
	// {
	//
	//  "verifyScore": 810.28,
	//
	// "faceAttack":"N",
	//
	// "invokeChannel": 1
	//
	// }
	FaceDetail *string `json:"FaceDetail,omitempty" xml:"FaceDetail,omitempty"`
	// Indicates whether the whitelist is hit. Valid value: **Y**.
	//
	// example:
	//
	// Y
	HitWhitelist *string `json:"HitWhitelist,omitempty" xml:"HitWhitelist,omitempty"`
	// The verification details from the authoritative source. Valid values:
	//
	// - **101**: Authentication passed.
	//
	// - **201**: Authentication failed. The name does not match the ID card number.
	//
	// - **202**: Authentication failed. The person is suspected to be the ID owner.
	//
	// - **203**: Authentication failed. No photo is found in the database.
	//
	// - **204**: Authentication failed. The faces do not belong to the same person.
	//
	// - **205**: Authentication failed. Image modeling for the image to be compared failed.
	//
	// - **206**: Authentication failed. The image format is incorrect.
	//
	// - **207**: Authentication failed. The uploaded image is too small. Upload another image.
	//
	// - **208**: Authentication failed. The quality of the uploaded face photo is poor. Upload another photo.
	//
	// - **301**: No record found. The ID card number does not exist in the database.
	//
	// - **302**: No record found. Verification cannot be performed.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s Id3MetaVerifyPROResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyPROResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) GetBizCode() *string {
	return s.BizCode
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) GetFaceDetail() *string {
	return s.FaceDetail
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) GetHitWhitelist() *string {
	return s.HitWhitelist
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) SetBizCode(v string) *Id3MetaVerifyPROResponseBodyResultObject {
	s.BizCode = &v
	return s
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) SetFaceDetail(v string) *Id3MetaVerifyPROResponseBodyResultObject {
	s.FaceDetail = &v
	return s
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) SetHitWhitelist(v string) *Id3MetaVerifyPROResponseBodyResultObject {
	s.HitWhitelist = &v
	return s
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) SetSubCode(v string) *Id3MetaVerifyPROResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *Id3MetaVerifyPROResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
