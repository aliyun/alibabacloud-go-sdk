// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrMaxResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DocOcrMaxResponseBody
	GetCode() *string
	SetMessage(v string) *DocOcrMaxResponseBody
	GetMessage() *string
	SetRequestId(v string) *DocOcrMaxResponseBody
	GetRequestId() *string
	SetResult(v *DocOcrMaxResponseBodyResult) *DocOcrMaxResponseBody
	GetResult() *DocOcrMaxResponseBodyResult
}

type DocOcrMaxResponseBody struct {
	// Return code
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result
	Result *DocOcrMaxResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DocOcrMaxResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxResponseBody) GoString() string {
	return s.String()
}

func (s *DocOcrMaxResponseBody) GetCode() *string {
	return s.Code
}

func (s *DocOcrMaxResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DocOcrMaxResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocOcrMaxResponseBody) GetResult() *DocOcrMaxResponseBodyResult {
	return s.Result
}

func (s *DocOcrMaxResponseBody) SetCode(v string) *DocOcrMaxResponseBody {
	s.Code = &v
	return s
}

func (s *DocOcrMaxResponseBody) SetMessage(v string) *DocOcrMaxResponseBody {
	s.Message = &v
	return s
}

func (s *DocOcrMaxResponseBody) SetRequestId(v string) *DocOcrMaxResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocOcrMaxResponseBody) SetResult(v *DocOcrMaxResponseBodyResult) *DocOcrMaxResponseBody {
	s.Result = v
	return s
}

func (s *DocOcrMaxResponseBody) Validate() error {
	return dara.Validate(s)
}

type DocOcrMaxResponseBodyResult struct {
	// Card and document recognition result	Only returned when the interface response is successful
	//
	// example:
	//
	// {
	//
	//   "docType": "PPTW01",
	//
	//   "ocrIdInfo": {
	//
	//     "passportNo": "36*******",
	//
	//     "expiryDate": "2032/02/10",
	//
	//     "placeOfBirth": "TAIWAN",
	//
	//     "surname": "CHEN",
	//
	//     "givenname": "LIN-CHUN",
	//
	//     "countryCode": "TWN",
	//
	//     "sex": "F",
	//
	//     "personalNo": "S22********",
	//
	//     "issueDate": "2022/02/10",
	//
	//     "birthDate": "1988/10/04"
	//
	//   }
	//
	// }
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// Whether the authentication passed.
	//
	// - Y: Passed.
	//
	// - N: Not passed.
	//
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// Sub-result code.
	//
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// Authentication ID
	//
	// example:
	//
	// hk573be80f944d95ac812e0*******a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DocOcrMaxResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DocOcrMaxResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *DocOcrMaxResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *DocOcrMaxResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *DocOcrMaxResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *DocOcrMaxResponseBodyResult) SetExtIdInfo(v string) *DocOcrMaxResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *DocOcrMaxResponseBodyResult) SetPassed(v string) *DocOcrMaxResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *DocOcrMaxResponseBodyResult) SetSubCode(v string) *DocOcrMaxResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *DocOcrMaxResponseBodyResult) SetTransactionId(v string) *DocOcrMaxResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *DocOcrMaxResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
