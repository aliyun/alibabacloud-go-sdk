// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCompareV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FaceCompareV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FaceCompareV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *FaceCompareV2ResponseBody
	GetRequestId() *string
	SetResult(v *FaceCompareV2ResponseBodyResult) *FaceCompareV2ResponseBody
	GetResult() *FaceCompareV2ResponseBodyResult
}

type FaceCompareV2ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceCompareV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceCompareV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FaceCompareV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FaceCompareV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FaceCompareV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceCompareV2ResponseBody) GetResult() *FaceCompareV2ResponseBodyResult {
	return s.Result
}

func (s *FaceCompareV2ResponseBody) SetCode(v string) *FaceCompareV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FaceCompareV2ResponseBody) SetMessage(v string) *FaceCompareV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FaceCompareV2ResponseBody) SetRequestId(v string) *FaceCompareV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceCompareV2ResponseBody) SetResult(v *FaceCompareV2ResponseBodyResult) *FaceCompareV2ResponseBody {
	s.Result = v
	return s
}

func (s *FaceCompareV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceCompareV2ResponseBodyResult struct {
	// example:
	//
	// 98
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceCompareV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceCompareV2ResponseBodyResult) GetFaceComparisonScore() *float64 {
	return s.FaceComparisonScore
}

func (s *FaceCompareV2ResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *FaceCompareV2ResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *FaceCompareV2ResponseBodyResult) SetFaceComparisonScore(v float64) *FaceCompareV2ResponseBodyResult {
	s.FaceComparisonScore = &v
	return s
}

func (s *FaceCompareV2ResponseBodyResult) SetPassed(v string) *FaceCompareV2ResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceCompareV2ResponseBodyResult) SetTransactionId(v string) *FaceCompareV2ResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *FaceCompareV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
