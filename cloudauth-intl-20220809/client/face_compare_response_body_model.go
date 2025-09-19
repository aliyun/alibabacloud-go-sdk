// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCompareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FaceCompareResponseBody
	GetCode() *string
	SetMessage(v string) *FaceCompareResponseBody
	GetMessage() *string
	SetRequestId(v string) *FaceCompareResponseBody
	GetRequestId() *string
	SetResult(v *FaceCompareResponseBodyResult) *FaceCompareResponseBody
	GetResult() *FaceCompareResponseBodyResult
}

type FaceCompareResponseBody struct {
	// The [response code](https://www.alibabacloud.com/help/en/ekyc/latest/facecompare?spm=a3c0i.23458820.2359477120.28.21167d3fzUmXQC#c43fd16d07mae).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed description of the response code.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object
	Result *FaceCompareResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceCompareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareResponseBody) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBody) GetCode() *string {
	return s.Code
}

func (s *FaceCompareResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FaceCompareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceCompareResponseBody) GetResult() *FaceCompareResponseBodyResult {
	return s.Result
}

func (s *FaceCompareResponseBody) SetCode(v string) *FaceCompareResponseBody {
	s.Code = &v
	return s
}

func (s *FaceCompareResponseBody) SetMessage(v string) *FaceCompareResponseBody {
	s.Message = &v
	return s
}

func (s *FaceCompareResponseBody) SetRequestId(v string) *FaceCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceCompareResponseBody) SetResult(v *FaceCompareResponseBodyResult) *FaceCompareResponseBody {
	s.Result = v
	return s
}

func (s *FaceCompareResponseBody) Validate() error {
	return dara.Validate(s)
}

type FaceCompareResponseBodyResult struct {
	// The face comparison score. The value ranges from 0 to 100.
	//
	// example:
	//
	// 98
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// The final authentication result. Valid values:
	//
	// - **Y**: The authentication is passed.
	//
	// - **N**: The authentication failed.
	//
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The transaction ID.
	//
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceCompareResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBodyResult) GetFaceComparisonScore() *float64 {
	return s.FaceComparisonScore
}

func (s *FaceCompareResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *FaceCompareResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *FaceCompareResponseBodyResult) SetFaceComparisonScore(v float64) *FaceCompareResponseBodyResult {
	s.FaceComparisonScore = &v
	return s
}

func (s *FaceCompareResponseBodyResult) SetPassed(v string) *FaceCompareResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceCompareResponseBodyResult) SetTransactionId(v string) *FaceCompareResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *FaceCompareResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
