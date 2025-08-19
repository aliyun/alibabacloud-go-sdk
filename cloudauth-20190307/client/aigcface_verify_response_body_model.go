// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIGCFaceVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AIGCFaceVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *AIGCFaceVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *AIGCFaceVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *AIGCFaceVerifyResponseBodyResultObject) *AIGCFaceVerifyResponseBody
	GetResultObject() *AIGCFaceVerifyResponseBodyResultObject
}

type AIGCFaceVerifyResponseBody struct {
	// Return code: 200 for success, others for failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result
	ResultObject *AIGCFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s AIGCFaceVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AIGCFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *AIGCFaceVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *AIGCFaceVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AIGCFaceVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AIGCFaceVerifyResponseBody) GetResultObject() *AIGCFaceVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *AIGCFaceVerifyResponseBody) SetCode(v string) *AIGCFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *AIGCFaceVerifyResponseBody) SetMessage(v string) *AIGCFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *AIGCFaceVerifyResponseBody) SetRequestId(v string) *AIGCFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AIGCFaceVerifyResponseBody) SetResultObject(v *AIGCFaceVerifyResponseBodyResultObject) *AIGCFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *AIGCFaceVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type AIGCFaceVerifyResponseBodyResultObject struct {
	// Unique real-person authentication identifier.
	//
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0****
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Authentication result. Values:
	//
	// ● Y: AIGC-generated face.
	//
	// ● N: Not detected
	//
	// example:
	//
	// Y
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Detection score
	//
	// example:
	//
	// 1.0000
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s AIGCFaceVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s AIGCFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *AIGCFaceVerifyResponseBodyResultObject) GetCertifyId() *string {
	return s.CertifyId
}

func (s *AIGCFaceVerifyResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *AIGCFaceVerifyResponseBodyResultObject) GetScore() *string {
	return s.Score
}

func (s *AIGCFaceVerifyResponseBodyResultObject) SetCertifyId(v string) *AIGCFaceVerifyResponseBodyResultObject {
	s.CertifyId = &v
	return s
}

func (s *AIGCFaceVerifyResponseBodyResultObject) SetResult(v string) *AIGCFaceVerifyResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *AIGCFaceVerifyResponseBodyResultObject) SetScore(v string) *AIGCFaceVerifyResponseBodyResultObject {
	s.Score = &v
	return s
}

func (s *AIGCFaceVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
