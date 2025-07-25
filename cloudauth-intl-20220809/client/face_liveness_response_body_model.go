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
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 42EA58CA-5DF4-55D5-82C4-5E7A40DA62BA
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceLivenessResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type FaceLivenessResponseBodyResult struct {
	ExtFaceInfo *FaceLivenessResponseBodyResultExtFaceInfo `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty" type:"Struct"`
	// example:
	//
	// N
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// example:
	//
	// 205
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
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
	return dara.Validate(s)
}

type FaceLivenessResponseBodyResultExtFaceInfo struct {
	FaceAge *int32 `json:"FaceAge,omitempty" xml:"FaceAge,omitempty"`
	// example:
	//
	// Y
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	FaceGender *string `json:"FaceGender,omitempty" xml:"FaceGender,omitempty"`
	// example:
	//
	// 87.19
	FaceQualityScore *float64 `json:"FaceQualityScore,omitempty" xml:"FaceQualityScore,omitempty"`
	// example:
	//
	// Y
	OcclusionResult *string `json:"OcclusionResult,omitempty" xml:"OcclusionResult,omitempty"`
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

func (s *FaceLivenessResponseBodyResultExtFaceInfo) GetOcclusionResult() *string {
	return s.OcclusionResult
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

func (s *FaceLivenessResponseBodyResultExtFaceInfo) SetOcclusionResult(v string) *FaceLivenessResponseBodyResultExtFaceInfo {
	s.OcclusionResult = &v
	return s
}

func (s *FaceLivenessResponseBodyResultExtFaceInfo) Validate() error {
	return dara.Validate(s)
}
