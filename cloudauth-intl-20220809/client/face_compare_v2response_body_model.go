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
	// The return code.
	//
	// 200: succeeded. Other values: error codes. For more information, see error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
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
	Result *FaceCompareV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	// The additional result information.
	ExtFaceInfo *FaceCompareV2ResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// The comparison score between the submitted face image and the reference face image during verification. Value range: 0 to 100.
	//
	// example:
	//
	// 98
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// Indicates whether the verification passed. Valid values:
	//
	// - Y: passed.
	//
	// - N: not passed.
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

func (s FaceCompareV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceCompareV2ResponseBodyResult) GetExtFaceInfo() *FaceCompareV2ResponseBodyResultExtFaceInfo {
	return s.ExtFaceInfo
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

func (s *FaceCompareV2ResponseBodyResult) SetExtFaceInfo(v *FaceCompareV2ResponseBodyResultExtFaceInfo) *FaceCompareV2ResponseBodyResult {
	s.ExtFaceInfo = v
	return s
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
	if s.ExtFaceInfo != nil {
		if err := s.ExtFaceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceCompareV2ResponseBodyResultExtFaceInfo struct {
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

func (s FaceCompareV2ResponseBodyResultExtFaceInfo) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareV2ResponseBodyResultExtFaceInfo) GoString() string {
	return s.String()
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) GetFaceQualityScore() *float64 {
	return s.FaceQualityScore
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) GetIlluminationScore() *float64 {
	return s.IlluminationScore
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) GetKaOcclusionScore() *float64 {
	return s.KaOcclusionScore
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) GetOcclusionScore() *float64 {
	return s.OcclusionScore
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) GetSharpnessScore() *float64 {
	return s.SharpnessScore
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) SetFaceQualityScore(v float64) *FaceCompareV2ResponseBodyResultExtFaceInfo {
	s.FaceQualityScore = &v
	return s
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) SetIlluminationScore(v float64) *FaceCompareV2ResponseBodyResultExtFaceInfo {
	s.IlluminationScore = &v
	return s
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) SetKaOcclusionScore(v float64) *FaceCompareV2ResponseBodyResultExtFaceInfo {
	s.KaOcclusionScore = &v
	return s
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) SetOcclusionScore(v float64) *FaceCompareV2ResponseBodyResultExtFaceInfo {
	s.OcclusionScore = &v
	return s
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) SetSharpnessScore(v float64) *FaceCompareV2ResponseBodyResultExtFaceInfo {
	s.SharpnessScore = &v
	return s
}

func (s *FaceCompareV2ResponseBodyResultExtFaceInfo) Validate() error {
	return dara.Validate(s)
}
