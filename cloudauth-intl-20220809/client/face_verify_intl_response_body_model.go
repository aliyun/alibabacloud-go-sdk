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
	// The response code.
	//
	// example:
	//
	// Success
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
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response result.
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
	// The face ID, user ID, and comparison score of the duplicate face in the face library when a duplicate face is retrieved.
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
	// The predicted reference age of the face. This parameter may not be returned if the prediction fails.
	//
	// example:
	//
	// 30
	FaceAge *int64 `json:"FaceAge,omitempty" xml:"FaceAge,omitempty"`
	// Indicates whether the collected face involves a liveness attack. Valid values:
	//
	// - Y: attack detected.
	//
	// - N: no attack detected.
	//
	// This parameter is returned only when passive liveness detection is enabled.
	//
	// example:
	//
	// N
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	// The probability of a liveness attack detected by passive liveness detection. Valid values: 0 to 100.
	//
	// This parameter is returned only when passive liveness detection is enabled.
	//
	// example:
	//
	// 99
	FaceAttackScore *float64 `json:"FaceAttackScore,omitempty" xml:"FaceAttackScore,omitempty"`
	// The comparison score between the submitted face photo and the reference face image during authentication. Valid values: 0 to 100.
	//
	// example:
	//
	// 95.0
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// The predicted gender of the face image. This parameter may not be returned if the prediction fails. Valid values:
	//
	// - M: male.
	//
	// - F: female.
	//
	// example:
	//
	// M
	FaceGender *string `json:"FaceGender,omitempty" xml:"FaceGender,omitempty"`
	// The final authentication result. Valid values:
	//
	// - Y: passed.
	//
	// - N: failed.
	//
	// example:
	//
	// Y
	FacePassed *string `json:"FacePassed,omitempty" xml:"FacePassed,omitempty"`
	// The FACEID returned when face registration succeeds. This parameter is returned only when automatic registration is enabled by the customer.
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
	// The unique identifier of the authentication request.
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
	// The attribute detection result of the main face, returned as a string. The result includes information such as gender, glasses, hat, mask, tattoo, smoking, other people, environment, expression, and age.
	//
	// example:
	//
	// {"gender":1,"glasses":0,"hat":1,"mask":1,"tattoo":0,"smoking":0,"others":1,"env":"office","expression":1,"age":16}
	FaceAttributeInfo *string `json:"FaceAttributeInfo,omitempty" xml:"FaceAttributeInfo,omitempty"`
	// The quality score of the live face. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 39.04
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// The algorithm score for illumination, which is a sub-dimension of quality judgment. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 97.43
	IlluminationScore *float64 `json:"IlluminationScore,omitempty" xml:"IlluminationScore,omitempty"`
	// The algorithm score for key area occlusion, which is a sub-dimension of quality judgment. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 100
	KaOcclusionScore *float64 `json:"KaOcclusionScore,omitempty" xml:"KaOcclusionScore,omitempty"`
	// The algorithm score for occlusion, which is a sub-dimension of quality judgment. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 50.26
	OcclusionScore *float64 `json:"OcclusionScore,omitempty" xml:"OcclusionScore,omitempty"`
	// The image sharpness score, which is a sub-dimension of quality judgment. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 86.47
	SharpnessScore *float64 `json:"SharpnessScore,omitempty" xml:"SharpnessScore,omitempty"`
	// The comprehensive quality score of the reference face image. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 39.04
	TargetFaceQualityScore *float64 `json:"TargetFaceQualityScore,omitempty" xml:"TargetFaceQualityScore,omitempty"`
	// The illumination score of the reference face image. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 0.02
	TargetIlluminationScore *float64 `json:"TargetIlluminationScore,omitempty" xml:"TargetIlluminationScore,omitempty"`
	// The key area occlusion score of the reference face image. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 20
	TargetKaOcclusionScore *float64 `json:"TargetKaOcclusionScore,omitempty" xml:"TargetKaOcclusionScore,omitempty"`
	// The occlusion score of the reference face image. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 50.26
	TargetOcclusionScore *float64 `json:"TargetOcclusionScore,omitempty" xml:"TargetOcclusionScore,omitempty"`
	// The sharpness score of the reference face image. Valid values: 0 to 100. A higher value indicates better quality.
	//
	// example:
	//
	// 86.47
	TargetSharpnessScore *float64 `json:"TargetSharpnessScore,omitempty" xml:"TargetSharpnessScore,omitempty"`
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
