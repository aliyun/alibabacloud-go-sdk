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
	// Response code.
	//
	// 200: Success.
	//
	// Other: Error code. For error code details, see Error Codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response message.
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
	// Response result.
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
	// Related result information
	ExtFaceInfo *FaceCompareResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// The comparison score between the submitted face photo and the reference face image during the authentication process. Value range: **0*	- to **100**.
	//
	// example:
	//
	// 98
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
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
	// The unique identifier of the authentication request.
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
	// Overall quality score
	//
	// example:
	//
	// 39.04
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// Illumination score
	//
	// example:
	//
	// 0.02
	IlluminationScore *float64 `json:"IlluminationScore,omitempty" xml:"IlluminationScore,omitempty"`
	// Key area occlusion score
	//
	// example:
	//
	// 20
	KaOcclusionScore *float64 `json:"KaOcclusionScore,omitempty" xml:"KaOcclusionScore,omitempty"`
	// Occlusion score
	//
	// example:
	//
	// 50.26
	OcclusionScore *float64 `json:"OcclusionScore,omitempty" xml:"OcclusionScore,omitempty"`
	// Sharpness score
	//
	// example:
	//
	// 86.47
	SharpnessScore          *float64 `json:"SharpnessScore,omitempty" xml:"SharpnessScore,omitempty"`
	TargetFaceQualityScore  *float64 `json:"TargetFaceQualityScore,omitempty" xml:"TargetFaceQualityScore,omitempty"`
	TargetIlluminationScore *float64 `json:"TargetIlluminationScore,omitempty" xml:"TargetIlluminationScore,omitempty"`
	TargetKaOcclusionScore  *float64 `json:"TargetKaOcclusionScore,omitempty" xml:"TargetKaOcclusionScore,omitempty"`
	TargetOcclusionScore    *float64 `json:"TargetOcclusionScore,omitempty" xml:"TargetOcclusionScore,omitempty"`
	TargetSharpnessScore    *float64 `json:"TargetSharpnessScore,omitempty" xml:"TargetSharpnessScore,omitempty"`
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

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetTargetFaceQualityScore() *float64 {
	return s.TargetFaceQualityScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetTargetIlluminationScore() *float64 {
	return s.TargetIlluminationScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetTargetKaOcclusionScore() *float64 {
	return s.TargetKaOcclusionScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetTargetOcclusionScore() *float64 {
	return s.TargetOcclusionScore
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) GetTargetSharpnessScore() *float64 {
	return s.TargetSharpnessScore
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

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetTargetFaceQualityScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.TargetFaceQualityScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetTargetIlluminationScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.TargetIlluminationScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetTargetKaOcclusionScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.TargetKaOcclusionScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetTargetOcclusionScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.TargetOcclusionScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) SetTargetSharpnessScore(v float64) *FaceCompareResponseBodyResultExtFaceInfo {
	s.TargetSharpnessScore = &v
	return s
}

func (s *FaceCompareResponseBodyResultExtFaceInfo) Validate() error {
	return dara.Validate(s)
}
