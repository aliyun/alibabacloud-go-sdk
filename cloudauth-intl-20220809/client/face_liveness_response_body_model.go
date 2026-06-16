// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceLivenessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FaceLivenessResponseBody
	GetCode() *string
	SetMessage(v string) *FaceLivenessResponseBody
	GetMessage() *string
	SetRequestId(v string) *FaceLivenessResponseBody
	GetRequestId() *string
	SetResult(v *FaceLivenessResponseBodyResult) *FaceLivenessResponseBody
	GetResult() *FaceLivenessResponseBodyResult
}

type FaceLivenessResponseBody struct {
	// The return code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned with the result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 42EA58CA-5DF4-55D5-82C4-5E7A40DA62BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *FaceLivenessResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceLivenessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessResponseBody) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBody) GetCode() *string {
	return s.Code
}

func (s *FaceLivenessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FaceLivenessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceLivenessResponseBody) GetResult() *FaceLivenessResponseBodyResult {
	return s.Result
}

func (s *FaceLivenessResponseBody) SetCode(v string) *FaceLivenessResponseBody {
	s.Code = &v
	return s
}

func (s *FaceLivenessResponseBody) SetMessage(v string) *FaceLivenessResponseBody {
	s.Message = &v
	return s
}

func (s *FaceLivenessResponseBody) SetRequestId(v string) *FaceLivenessResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceLivenessResponseBody) SetResult(v *FaceLivenessResponseBodyResult) *FaceLivenessResponseBody {
	s.Result = v
	return s
}

func (s *FaceLivenessResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceLivenessResponseBodyResult struct {
	// The face result information.
	ExtFaceInfo *FaceLivenessResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// Indicates whether the authentication passed. Valid values:
	//
	// - Y: passed.
	//
	// - N: not passed.
	//
	// example:
	//
	// N
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The sub-result code.
	//
	// example:
	//
	// 205
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// The unique ID of the authentication request.
	//
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceLivenessResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBodyResult) GetExtFaceInfo() *FaceLivenessResponseBodyResultExtFaceInfo {
	return s.ExtFaceInfo
}

func (s *FaceLivenessResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *FaceLivenessResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *FaceLivenessResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *FaceLivenessResponseBodyResult) SetExtFaceInfo(v *FaceLivenessResponseBodyResultExtFaceInfo) *FaceLivenessResponseBodyResult {
	s.ExtFaceInfo = v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetPassed(v string) *FaceLivenessResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetSubCode(v string) *FaceLivenessResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *FaceLivenessResponseBodyResult) SetTransactionId(v string) *FaceLivenessResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *FaceLivenessResponseBodyResult) Validate() error {
	if s.ExtFaceInfo != nil {
		if err := s.ExtFaceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceLivenessResponseBodyResultExtFaceInfo struct {
	// The predicted reference age based on the face. The prediction may fail and return no value.
	//
	// example:
	//
	// 18
	FaceAge *int32 `json:"FaceAge,omitempty" xml:"FaceAge,omitempty"`
	// The liveness detection result. Valid values: Y (attack detected) and N (normal).
	//
	// example:
	//
	// Y
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	// The predicted gender based on the face photo. The prediction may fail and return no value. Valid values:
	//
	// - M: male.
	//
	// - F: female.
	//
	// example:
	//
	// M
	FaceGender *string `json:"FaceGender,omitempty" xml:"FaceGender,omitempty"`
	// The face quality score (0 to 100). This value is returned only when the face quality score switch is enabled in the request parameters.
	//
	// example:
	//
	// 87.19
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
	// The occlusion detection result. Valid values: Y (occluded) and N (not occluded). This value is returned only when the occlusion detection switch is enabled.
	//
	// example:
	//
	// Y
	OcclusionResult *string `json:"OcclusionResult,omitempty" xml:"OcclusionResult,omitempty"`
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

func (s FaceLivenessResponseBodyResultExtFaceInfo) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessResponseBodyResultExtFaceInfo) GoString() string {
	return s.String()
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetFaceAge() *int32 {
	return s.FaceAge
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetFaceAttack() *string {
	return s.FaceAttack
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetFaceGender() *string {
	return s.FaceGender
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetFaceQualityScore() *float64 {
	return s.FaceQualityScore
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetIlluminationScore() *float64 {
	return s.IlluminationScore
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetKaOcclusionScore() *float64 {
	return s.KaOcclusionScore
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetOcclusionResult() *string {
	return s.OcclusionResult
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetOcclusionScore() *float64 {
	return s.OcclusionScore
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetSharpnessScore() *float64 {
	return s.SharpnessScore
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceAge(v int32) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceAge = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceAttack(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceAttack = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceGender(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceGender = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetFaceQualityScore(v float64) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.FaceQualityScore = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetIlluminationScore(v float64) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.IlluminationScore = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetKaOcclusionScore(v float64) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.KaOcclusionScore = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetOcclusionResult(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.OcclusionResult = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetOcclusionScore(v float64) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.OcclusionScore = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetSharpnessScore(v float64) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.SharpnessScore = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) Validate() error {
	return dara.Validate(s)
}
