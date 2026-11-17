// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceLivenessV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FaceLivenessV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FaceLivenessV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *FaceLivenessV2ResponseBody
	GetRequestId() *string
	SetResult(v *FaceLivenessV2ResponseBodyResult) *FaceLivenessV2ResponseBody
	GetResult() *FaceLivenessV2ResponseBodyResult
}

type FaceLivenessV2ResponseBody struct {
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
	Result *FaceLivenessV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceLivenessV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FaceLivenessV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FaceLivenessV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FaceLivenessV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceLivenessV2ResponseBody) GetResult() *FaceLivenessV2ResponseBodyResult {
	return s.Result
}

func (s *FaceLivenessV2ResponseBody) SetCode(v string) *FaceLivenessV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FaceLivenessV2ResponseBody) SetMessage(v string) *FaceLivenessV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FaceLivenessV2ResponseBody) SetRequestId(v string) *FaceLivenessV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceLivenessV2ResponseBody) SetResult(v *FaceLivenessV2ResponseBodyResult) *FaceLivenessV2ResponseBody {
	s.Result = v
	return s
}

func (s *FaceLivenessV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceLivenessV2ResponseBodyResult struct {
	// The face result information.
	ExtFaceInfo *FaceLivenessV2ResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// Indicates whether the authentication is passed. Valid values:
	//
	// - Y: passed
	//
	// - N: failed
	//
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
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
	// hksb7ba1b28130d24e015d694361****
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceLivenessV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceLivenessV2ResponseBodyResult) GetExtFaceInfo() *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	return s.ExtFaceInfo
}

func (s *FaceLivenessV2ResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *FaceLivenessV2ResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *FaceLivenessV2ResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *FaceLivenessV2ResponseBodyResult) SetExtFaceInfo(v *FaceLivenessV2ResponseBodyResultExtFaceInfo) *FaceLivenessV2ResponseBodyResult {
	s.ExtFaceInfo = v
	return s
}

func (s *FaceLivenessV2ResponseBodyResult) SetPassed(v string) *FaceLivenessV2ResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResult) SetSubCode(v string) *FaceLivenessV2ResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResult) SetTransactionId(v string) *FaceLivenessV2ResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResult) Validate() error {
	if s.ExtFaceInfo != nil {
		if err := s.ExtFaceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceLivenessV2ResponseBodyResultExtFaceInfo struct {
	// The predicted reference age of the face. If the prediction fails, no value is returned.
	//
	// example:
	//
	// 18
	FaceAge *int64 `json:"FaceAge,omitempty" xml:"FaceAge,omitempty"`
	// The liveness detection result. Valid values: Y (attack) and N (normal).
	//
	// example:
	//
	// Y
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	// Detects the facial attributes of the primary face in the input image. The primary face is the one closest to the camera, with the largest facial area, and in focus. The detection result is returned as a string for downstream processing.
	//
	// The returned fields are defined as follows:
	//
	// **gender**: the gender of the primary person. A value of 1 indicates male, and 0 indicates female.
	//
	// **glasses**: whether the primary person is wearing glasses. A value of 1 indicates that glasses are detected, and 0 indicates that glasses are not detected.
	//
	// **hat**: whether the primary person is wearing a hat (including headscarves and veils). A value of 1 indicates that a hat is detected, and 0 indicates that a hat is not detected.
	//
	// **mask**: whether the primary person is wearing a mask. A value of 1 indicates that a mask is detected, and 0 indicates that a mask is not detected.
	//
	// **tattoo**: whether the primary person has visible tattoos. A value of 1 indicates that a tattoo is detected, and 0 indicates that a tattoo is not detected.
	//
	// **smoking**: whether the primary person is smoking. A value of 1 indicates that smoking is detected, and 0 indicates that smoking is not detected.
	//
	// **others**: whether other people are present besides the primary person. A value of 1 indicates that other people are detected, and 0 indicates that other people are not detected.
	//
	// **env**: the environment where the primary person is located, such as an office, factory, hospital, or vehicle. Valid values: office, hospital, factory, vehical, and other.
	//
	// **expression**: whether the primary person has an exaggerated expression (facial expressions that significantly deviate from a natural or calm state). A value of 1 indicates that an exaggerated expression is detected, and 0 indicates that an exaggerated expression is not detected.
	//
	// **age**: the actual age of the primary person in years.
	//
	// example:
	//
	// {
	//
	//         gender: 1
	//
	//         glasses: 0
	//
	//         hat: 1
	//
	//         mask: 1
	//
	//         tattoo: 0
	//
	//         smoking: 0
	//
	//         others: 1
	//
	//         env：office
	//
	//         expression：1
	//
	//         age：16
	//
	// }
	FaceAttributeInfo *string `json:"FaceAttributeInfo,omitempty" xml:"FaceAttributeInfo,omitempty"`
	// The predicted gender of the face in the image. If the prediction fails, no value is returned. Valid values:
	//
	// - M: male
	//
	// - F: female
	//
	// example:
	//
	// M
	FaceGender *string `json:"FaceGender,omitempty" xml:"FaceGender,omitempty"`
	// The liveness face quality score. Valid values: 0 to 100. A larger value indicates better quality.
	//
	// example:
	//
	// 79.04
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// The algorithm score for illumination in the quality judgment sub-dimension. Valid values: 0 to 100. A larger value indicates better quality.
	//
	// example:
	//
	// 97.43
	IlluminationScore *float64 `json:"IlluminationScore,omitempty" xml:"IlluminationScore,omitempty"`
	// The algorithm score for key area occlusion in the quality judgment sub-dimension. Valid values: 0 to 100. A larger value indicates better quality.
	//
	// example:
	//
	// 100
	KaOcclusionScore *float64 `json:"KaOcclusionScore,omitempty" xml:"KaOcclusionScore,omitempty"`
	// Indicates whether the face is occluded. A value of Y indicates that the face is occluded. A value of N indicates that the face is not occluded.
	//
	// example:
	//
	// Y
	OcclusionResult *string `json:"OcclusionResult,omitempty" xml:"OcclusionResult,omitempty"`
	// The algorithm score for occlusion in the quality judgment sub-dimension. Valid values: 0 to 100. A larger value indicates better quality.
	//
	// example:
	//
	// 50.26
	OcclusionScore *float64 `json:"OcclusionScore,omitempty" xml:"OcclusionScore,omitempty"`
	// The image sharpness score in the quality judgment sub-dimension. Valid values: 0 to 100. A larger value indicates better quality.
	//
	// example:
	//
	// 60.78
	SharpnessScore *float64 `json:"SharpnessScore,omitempty" xml:"SharpnessScore,omitempty"`
}

func (s FaceLivenessV2ResponseBodyResultExtFaceInfo) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessV2ResponseBodyResultExtFaceInfo) GoString() string {
	return s.String()
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetFaceAge() *int64 {
	return s.FaceAge
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetFaceAttack() *string {
	return s.FaceAttack
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetFaceAttributeInfo() *string {
	return s.FaceAttributeInfo
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetFaceGender() *string {
	return s.FaceGender
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetFaceQualityScore() *float64 {
	return s.FaceQualityScore
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetIlluminationScore() *float64 {
	return s.IlluminationScore
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetKaOcclusionScore() *float64 {
	return s.KaOcclusionScore
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetOcclusionResult() *string {
	return s.OcclusionResult
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetOcclusionScore() *float64 {
	return s.OcclusionScore
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) GetSharpnessScore() *float64 {
	return s.SharpnessScore
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetFaceAge(v int64) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.FaceAge = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetFaceAttack(v string) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.FaceAttack = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetFaceAttributeInfo(v string) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.FaceAttributeInfo = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetFaceGender(v string) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.FaceGender = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetFaceQualityScore(v float64) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.FaceQualityScore = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetIlluminationScore(v float64) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.IlluminationScore = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetKaOcclusionScore(v float64) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.KaOcclusionScore = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetOcclusionResult(v string) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.OcclusionResult = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetOcclusionScore(v float64) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.OcclusionScore = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) SetSharpnessScore(v float64) *FaceLivenessV2ResponseBodyResultExtFaceInfo {
	s.SharpnessScore = &v
	return s
}

func (s *FaceLivenessV2ResponseBodyResultExtFaceInfo) Validate() error {
	return dara.Validate(s)
}
