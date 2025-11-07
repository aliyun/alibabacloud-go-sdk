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
	// Return code: 200 for success, others for failure.
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
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
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
	// Unique identifier for real-person authentication.
	//
	// example:
	//
	// sha58aeae7ea2f5ed069530f58df4e6d
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Deletion result. Possible values are as follows:
	//
	// - Y: Deletion successful.
	//
	// - N: Deletion failed.
	//
	// example:
	//
	// N
	DeleteResult *string `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	// Reason for deletion failure
	//
	// - NOT_DELETE_REPEATEDLY: Cannot be deleted repeatedly
	//
	// - NEED_QUERY_VERIFY_RESULT: Need to query the verification result first, then delete
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
