// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FaceVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *FaceVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *FaceVerifyIntlResponseBody
	GetRequestId() *string
	SetResult(v *FaceVerifyIntlResponseBodyResult) *FaceVerifyIntlResponseBody
	GetResult() *FaceVerifyIntlResponseBodyResult
}

type FaceVerifyIntlResponseBody struct {
	// The return code.
	//
	// example:
	//
	// Success
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
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *FaceVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *FaceVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *FaceVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FaceVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceVerifyIntlResponseBody) GetResult() *FaceVerifyIntlResponseBodyResult {
	return s.Result
}

func (s *FaceVerifyIntlResponseBody) SetCode(v string) *FaceVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *FaceVerifyIntlResponseBody) SetMessage(v string) *FaceVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *FaceVerifyIntlResponseBody) SetRequestId(v string) *FaceVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceVerifyIntlResponseBody) SetResult(v *FaceVerifyIntlResponseBodyResult) *FaceVerifyIntlResponseBody {
	s.Result = v
	return s
}

func (s *FaceVerifyIntlResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceVerifyIntlResponseBodyResult struct {
	// The face ID, user ID, and comparison score of the corresponding face in the face library when a duplicate face is found during retrieval.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "faceGroupCode": "sg7****uzt",
	//
	//         "faceId": "f5a921*******9e792ec84c8f0ca592a",
	//
	//         "merchantUserId": "fa****01"
	//
	//     }
	//
	// ]
	DuplicateFace *string `json:"DuplicateFace,omitempty" xml:"DuplicateFace,omitempty"`
	// The additional face result information.
	ExtFaceInfo *FaceVerifyIntlResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// The predicted reference age of the face. Prediction may fail and the value may not be returned in some cases.
	//
	// example:
	//
	// 30
	FaceAge *int64 `json:"FaceAge,omitempty" xml:"FaceAge,omitempty"`
	// Indicates whether the captured face involves a liveness attack. Valid values:
	//
	// - Y: attack detected.
	//
	// - N: no attack detected.
	//
	// Returned when passive liveness detection is enabled.
	//
	// example:
	//
	// N
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	// The probability of a passive liveness detection attack on the face. Value range: 0 to 100. Returned when passive liveness detection is enabled.
	//
	// example:
	//
	// 99
	FaceAttackScore *float64 `json:"FaceAttackScore,omitempty" xml:"FaceAttackScore,omitempty"`
	// The comparison score between the submitted face image and the reference face image during verification. Value range: 0 to 100.
	//
	// example:
	//
	// 95.0
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// The predicted gender of the face image. Prediction may fail and the value may not be returned in some cases. Valid values:
	//
	// - M: male.
	//
	// - F: female.
	//
	// example:
	//
	// M
	FaceGender *string `json:"FaceGender,omitempty" xml:"FaceGender,omitempty"`
	// The final verification result. Valid values:
	//
	// - Y: passed.
	//
	// - N: not passed.
	//
	// example:
	//
	// Y
	FacePassed *string `json:"FacePassed,omitempty" xml:"FacePassed,omitempty"`
	// The corresponding FACEID returned only when the customer has enabled automatic registration and the face is registered successfully.
	//
	// example:
	//
	// 9e792ec84c8f0ca592a
	FaceRegistrationId *string `json:"FaceRegistrationId,omitempty" xml:"FaceRegistrationId,omitempty"`
	// The face registration result. Valid values:
	//
	// - 0: failed.
	//
	// - 1: succeeded.
	//
	// example:
	//
	// 0
	FaceRegistrationResult *int64 `json:"FaceRegistrationResult,omitempty" xml:"FaceRegistrationResult,omitempty"`
	// The sub-result code.
	//
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// The unique identifier of the verification request.
	//
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceVerifyIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceVerifyIntlResponseBodyResult) GetDuplicateFace() *string {
	return s.DuplicateFace
}

func (s *FaceVerifyIntlResponseBodyResult) GetExtFaceInfo() *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	return s.ExtFaceInfo
}

func (s *FaceVerifyIntlResponseBodyResult) GetFaceAge() *int64 {
	return s.FaceAge
}

func (s *FaceVerifyIntlResponseBodyResult) GetFaceAttack() *string {
	return s.FaceAttack
}

func (s *FaceVerifyIntlResponseBodyResult) GetFaceAttackScore() *float64 {
	return s.FaceAttackScore
}

func (s *FaceVerifyIntlResponseBodyResult) GetFaceComparisonScore() *float64 {
	return s.FaceComparisonScore
}

func (s *FaceVerifyIntlResponseBodyResult) GetFaceGender() *string {
	return s.FaceGender
}

func (s *FaceVerifyIntlResponseBodyResult) GetFacePassed() *string {
	return s.FacePassed
}

func (s *FaceVerifyIntlResponseBodyResult) GetFaceRegistrationId() *string {
	return s.FaceRegistrationId
}

func (s *FaceVerifyIntlResponseBodyResult) GetFaceRegistrationResult() *int64 {
	return s.FaceRegistrationResult
}

func (s *FaceVerifyIntlResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *FaceVerifyIntlResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *FaceVerifyIntlResponseBodyResult) SetDuplicateFace(v string) *FaceVerifyIntlResponseBodyResult {
	s.DuplicateFace = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetExtFaceInfo(v *FaceVerifyIntlResponseBodyResultExtFaceInfo) *FaceVerifyIntlResponseBodyResult {
	s.ExtFaceInfo = v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetFaceAge(v int64) *FaceVerifyIntlResponseBodyResult {
	s.FaceAge = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetFaceAttack(v string) *FaceVerifyIntlResponseBodyResult {
	s.FaceAttack = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetFaceAttackScore(v float64) *FaceVerifyIntlResponseBodyResult {
	s.FaceAttackScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetFaceComparisonScore(v float64) *FaceVerifyIntlResponseBodyResult {
	s.FaceComparisonScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetFaceGender(v string) *FaceVerifyIntlResponseBodyResult {
	s.FaceGender = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetFacePassed(v string) *FaceVerifyIntlResponseBodyResult {
	s.FacePassed = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetFaceRegistrationId(v string) *FaceVerifyIntlResponseBodyResult {
	s.FaceRegistrationId = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetFaceRegistrationResult(v int64) *FaceVerifyIntlResponseBodyResult {
	s.FaceRegistrationResult = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetSubCode(v string) *FaceVerifyIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) SetTransactionId(v string) *FaceVerifyIntlResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResult) Validate() error {
	if s.ExtFaceInfo != nil {
		if err := s.ExtFaceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceVerifyIntlResponseBodyResultExtFaceInfo struct {
	FaceAttributeInfo *string `json:"FaceAttributeInfo,omitempty" xml:"FaceAttributeInfo,omitempty"`
	// The liveness face quality score. Value range: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 39.04
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// The algorithm score for illumination as a sub-dimension of quality assessment. Value range: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 97.43
	IlluminationScore *float64 `json:"IlluminationScore,omitempty" xml:"IlluminationScore,omitempty"`
	// The algorithm score for key area occlusion as a sub-dimension of quality assessment. Value range: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 100
	KaOcclusionScore *float64 `json:"KaOcclusionScore,omitempty" xml:"KaOcclusionScore,omitempty"`
	// The algorithm score for occlusion as a sub-dimension of quality assessment. Value range: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 50.26
	OcclusionScore *float64 `json:"OcclusionScore,omitempty" xml:"OcclusionScore,omitempty"`
	// The image sharpness score as a sub-dimension of quality assessment. Value range: 0 to 100. A higher value indicates better quality.
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

func (s FaceVerifyIntlResponseBodyResultExtFaceInfo) String() string {
	return dara.Prettify(s)
}

func (s FaceVerifyIntlResponseBodyResultExtFaceInfo) GoString() string {
	return s.String()
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetFaceAttributeInfo() *string {
	return s.FaceAttributeInfo
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetFaceQualityScore() *float64 {
	return s.FaceQualityScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetIlluminationScore() *float64 {
	return s.IlluminationScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetKaOcclusionScore() *float64 {
	return s.KaOcclusionScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetOcclusionScore() *float64 {
	return s.OcclusionScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetSharpnessScore() *float64 {
	return s.SharpnessScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetTargetFaceQualityScore() *float64 {
	return s.TargetFaceQualityScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetTargetIlluminationScore() *float64 {
	return s.TargetIlluminationScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetTargetKaOcclusionScore() *float64 {
	return s.TargetKaOcclusionScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetTargetOcclusionScore() *float64 {
	return s.TargetOcclusionScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) GetTargetSharpnessScore() *float64 {
	return s.TargetSharpnessScore
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetFaceAttributeInfo(v string) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.FaceAttributeInfo = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetFaceQualityScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.FaceQualityScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetIlluminationScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.IlluminationScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetKaOcclusionScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.KaOcclusionScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetOcclusionScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.OcclusionScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetSharpnessScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.SharpnessScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetTargetFaceQualityScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.TargetFaceQualityScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetTargetIlluminationScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.TargetIlluminationScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetTargetKaOcclusionScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.TargetKaOcclusionScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetTargetOcclusionScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.TargetOcclusionScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) SetTargetSharpnessScore(v float64) *FaceVerifyIntlResponseBodyResultExtFaceInfo {
	s.TargetSharpnessScore = &v
	return s
}

func (s *FaceVerifyIntlResponseBodyResultExtFaceInfo) Validate() error {
	return dara.Validate(s)
}
