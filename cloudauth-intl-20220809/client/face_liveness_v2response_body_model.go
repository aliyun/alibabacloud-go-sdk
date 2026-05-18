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
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceLivenessV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	ExtFaceInfo *FaceLivenessV2ResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
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
	// example:
	//
	// 18
	FaceAge *int64 `json:"FaceAge,omitempty" xml:"FaceAge,omitempty"`
	// example:
	//
	// Y
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	// example:
	//
	// M
	FaceGender *string `json:"FaceGender,omitempty" xml:"FaceGender,omitempty"`
	// example:
	//
	// 79.04
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// example:
	//
	// 97.43
	IlluminationScore *float64 `json:"IlluminationScore,omitempty" xml:"IlluminationScore,omitempty"`
	// example:
	//
	// 100
	KaOcclusionScore *float64 `json:"KaOcclusionScore,omitempty" xml:"KaOcclusionScore,omitempty"`
	// example:
	//
	// Y
	OcclusionResult *string `json:"OcclusionResult,omitempty" xml:"OcclusionResult,omitempty"`
	// example:
	//
	// 50.26
	OcclusionScore *float64 `json:"OcclusionScore,omitempty" xml:"OcclusionScore,omitempty"`
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
