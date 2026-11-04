// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceVerifyResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFaceVerifyResultResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteFaceVerifyResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFaceVerifyResultResponseBody
	GetRequestId() *string
	SetResultObject(v *DeleteFaceVerifyResultResponseBodyResultObject) *DeleteFaceVerifyResultResponseBody
	GetResultObject() *DeleteFaceVerifyResultResponseBodyResultObject
}

type DeleteFaceVerifyResultResponseBody struct {
	// The return code. A value of 200 indicates success, and other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result information.
	ResultObject *DeleteFaceVerifyResultResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DeleteFaceVerifyResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceVerifyResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFaceVerifyResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFaceVerifyResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceVerifyResultResponseBody) GetResultObject() *DeleteFaceVerifyResultResponseBodyResultObject {
	return s.ResultObject
}

func (s *DeleteFaceVerifyResultResponseBody) SetCode(v string) *DeleteFaceVerifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBody) SetMessage(v string) *DeleteFaceVerifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBody) SetRequestId(v string) *DeleteFaceVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBody) SetResultObject(v *DeleteFaceVerifyResultResponseBodyResultObject) *DeleteFaceVerifyResultResponseBody {
	s.ResultObject = v
	return s
}

func (s *DeleteFaceVerifyResultResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteFaceVerifyResultResponseBodyResultObject struct {
	// The unique identifier for ID Verification.
	//
	// example:
	//
	// sha58aeae7ea2f5ed069530f58df4e6d
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The deletion result. Valid values:
	//
	// - Y: The deletion succeeded.
	//
	// - N: The deletion failed.
	//
	// example:
	//
	// N
	DeleteResult *string `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	// The reason for the deletion failure. Valid values:
	//
	// - NOT_DELETE_REPEATEDLY: The data cannot be repeatedly deleted.
	//
	// - NEED_QUERY_VERIFY_RESULT: You must query the verification result before you delete the data.
	//
	// example:
	//
	// NOT_DELETE_REPEATEDLY
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s DeleteFaceVerifyResultResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceVerifyResultResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) GetDeleteResult() *string {
	return s.DeleteResult
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) GetFailReason() *string {
	return s.FailReason
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) SetCertifyId(v string) *DeleteFaceVerifyResultResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) SetDeleteResult(v string) *DeleteFaceVerifyResultResponseBodyResultObject {
	s.DeleteResult = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) SetFailReason(v string) *DeleteFaceVerifyResultResponseBodyResultObject {
	s.FailReason = &v
	return s
}

func (s *DeleteFaceVerifyResultResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
