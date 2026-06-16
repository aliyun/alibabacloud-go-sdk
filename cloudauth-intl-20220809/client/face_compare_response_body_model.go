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
	// The response code.
	//
	// 200: The request was successful.
	//
	// Other values: An error occurred. For more information, see error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
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
	// The returned result.
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceCompareResponseBodyResult struct {
	// The additional result information.
	ExtFaceInfo *FaceCompareResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// The comparison score between the submitted face image and the reference face image during verification. Value range: **0*	- to **100**.
	//
	// example:
	//
	// 98
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// Indicates whether the verification passed.
	//
	// - Y: Passed.
	//
	// - N: Not passed.
	//
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The unique ID of the verification request.
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

func (s *FaceCompareResponseBodyResult) GetExtFaceInfo() *FaceCompareResponseBodyResultExtFaceInfo {
	return s.ExtFaceInfo
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

func (s *FaceCompareResponseBodyResult) SetExtFaceInfo(v *FaceCompareResponseBodyResultExtFaceInfo) *FaceCompareResponseBodyResult {
	s.ExtFaceInfo = v
	return s
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
	if s.ExtFaceInfo != nil {
		if err := s.ExtFaceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceCompareResponseBodyResultExtFaceInfo struct {
	// The overall quality score.
	//
	// example:
	//
	// 39.04
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// The illumination score.
	//
	// example:
	//
	// 0.02
	IlluminationScore *float64 `json:"IlluminationScore,omitempty" xml:"IlluminationScore,omitempty"`
	// The key area occlusion score.
	//
	// example:
	//
	// 20
	KaOcclusionScore *float64 `json:"KaOcclusionScore,omitempty" xml:"KaOcclusionScore,omitempty"`
	// The occlusion score.
	//
	// example:
	//
	// 50.26
	OcclusionScore *float64 `json:"OcclusionScore,omitempty" xml:"OcclusionScore,omitempty"`
	// The sharpness score.
	//
	// example:
	//
	// 86.47
	SharpnessScore *float64 `json:"SharpnessScore,omitempty" xml:"SharpnessScore,omitempty"`
}

func (s FaceCompareResponseBodyResultExtFaceInfo) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareResponseBodyResultExtFaceInfo) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetFaceQualityScore() *float64 {
	return s.FaceQualityScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetIlluminationScore() *float64 {
	return s.IlluminationScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetKaOcclusionScore() *float64 {
	return s.KaOcclusionScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetOcclusionScore() *float64 {
	return s.OcclusionScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetSharpnessScore() *float64 {
	return s.SharpnessScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetFaceQualityScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.FaceQualityScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetIlluminationScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.IlluminationScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetKaOcclusionScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.KaOcclusionScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetOcclusionScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.OcclusionScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetSharpnessScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.SharpnessScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) Validate() error {
	return dara.Validate(s)
}
