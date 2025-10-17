// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCrossCompareIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FaceCrossCompareIntlResponseBody
	GetCode() *string
	SetMessage(v string) *FaceCrossCompareIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *FaceCrossCompareIntlResponseBody
	GetRequestId() *string
	SetResult(v *FaceCrossCompareIntlResponseBodyResult) *FaceCrossCompareIntlResponseBody
	GetResult() *FaceCrossCompareIntlResponseBodyResult
}

type FaceCrossCompareIntlResponseBody struct {
	// Return code.
	//
	// example:
	//
	// Success
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
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	Result *FaceCrossCompareIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceCrossCompareIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceCrossCompareIntlResponseBody) GoString() string {
	return s.String()
}

func (s *FaceCrossCompareIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *FaceCrossCompareIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FaceCrossCompareIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceCrossCompareIntlResponseBody) GetResult() *FaceCrossCompareIntlResponseBodyResult {
	return s.Result
}

func (s *FaceCrossCompareIntlResponseBody) SetCode(v string) *FaceCrossCompareIntlResponseBody {
	s.Code = &v
	return s
}

func (s *FaceCrossCompareIntlResponseBody) SetMessage(v string) *FaceCrossCompareIntlResponseBody {
	s.Message = &v
	return s
}

func (s *FaceCrossCompareIntlResponseBody) SetRequestId(v string) *FaceCrossCompareIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceCrossCompareIntlResponseBody) SetResult(v *FaceCrossCompareIntlResponseBodyResult) *FaceCrossCompareIntlResponseBody {
	s.Result = v
	return s
}

func (s *FaceCrossCompareIntlResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceCrossCompareIntlResponseBodyResult struct {
	// A to B comparison score, range 0～100.
	//
	// example:
	//
	// 99
	FaceComparisonScoreA2B *float64 `json:"FaceComparisonScoreA2B,omitempty" xml:"FaceComparisonScoreA2B,omitempty"`
	// B to C comparison score, range 0～100.
	//
	// example:
	//
	// 99
	FaceComparisonScoreB2C *float64 `json:"FaceComparisonScoreB2C,omitempty" xml:"FaceComparisonScoreB2C,omitempty"`
	// C to A comparison score, range 0～100.
	//
	// example:
	//
	// 99
	FaceComparisonScoreC2A *float64 `json:"FaceComparisonScoreC2A,omitempty" xml:"FaceComparisonScoreC2A,omitempty"`
	// Final verification result, values:
	//
	// - Y: Pass
	//
	// - N: Fail
	//
	// example:
	//
	// Y
	FacePassed *string `json:"FacePassed,omitempty" xml:"FacePassed,omitempty"`
	// Unique identifier for the authentication request.
	//
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceCrossCompareIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceCrossCompareIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceCrossCompareIntlResponseBodyResult) GetFaceComparisonScoreA2B() *float64 {
	return s.FaceComparisonScoreA2B
}

func (s *FaceCrossCompareIntlResponseBodyResult) GetFaceComparisonScoreB2C() *float64 {
	return s.FaceComparisonScoreB2C
}

func (s *FaceCrossCompareIntlResponseBodyResult) GetFaceComparisonScoreC2A() *float64 {
	return s.FaceComparisonScoreC2A
}

func (s *FaceCrossCompareIntlResponseBodyResult) GetFacePassed() *string {
	return s.FacePassed
}

func (s *FaceCrossCompareIntlResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *FaceCrossCompareIntlResponseBodyResult) SetFaceComparisonScoreA2B(v float64) *FaceCrossCompareIntlResponseBodyResult {
	s.FaceComparisonScoreA2B = &v
	return s
}

func (s *FaceCrossCompareIntlResponseBodyResult) SetFaceComparisonScoreB2C(v float64) *FaceCrossCompareIntlResponseBodyResult {
	s.FaceComparisonScoreB2C = &v
	return s
}

func (s *FaceCrossCompareIntlResponseBodyResult) SetFaceComparisonScoreC2A(v float64) *FaceCrossCompareIntlResponseBodyResult {
	s.FaceComparisonScoreC2A = &v
	return s
}

func (s *FaceCrossCompareIntlResponseBodyResult) SetFacePassed(v string) *FaceCrossCompareIntlResponseBodyResult {
	s.FacePassed = &v
	return s
}

func (s *FaceCrossCompareIntlResponseBodyResult) SetTransactionId(v string) *FaceCrossCompareIntlResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *FaceCrossCompareIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
