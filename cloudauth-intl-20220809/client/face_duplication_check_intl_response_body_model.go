// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceDuplicationCheckIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FaceDuplicationCheckIntlResponseBody
	GetCode() *string
	SetMessage(v string) *FaceDuplicationCheckIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *FaceDuplicationCheckIntlResponseBody
	GetRequestId() *string
	SetResult(v *FaceDuplicationCheckIntlResponseBodyResult) *FaceDuplicationCheckIntlResponseBody
	GetResult() *FaceDuplicationCheckIntlResponseBodyResult
}

type FaceDuplicationCheckIntlResponseBody struct {
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
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	Result *FaceDuplicationCheckIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceDuplicationCheckIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceDuplicationCheckIntlResponseBody) GoString() string {
	return s.String()
}

func (s *FaceDuplicationCheckIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *FaceDuplicationCheckIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FaceDuplicationCheckIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceDuplicationCheckIntlResponseBody) GetResult() *FaceDuplicationCheckIntlResponseBodyResult {
	return s.Result
}

func (s *FaceDuplicationCheckIntlResponseBody) SetCode(v string) *FaceDuplicationCheckIntlResponseBody {
	s.Code = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBody) SetMessage(v string) *FaceDuplicationCheckIntlResponseBody {
	s.Message = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBody) SetRequestId(v string) *FaceDuplicationCheckIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBody) SetResult(v *FaceDuplicationCheckIntlResponseBodyResult) *FaceDuplicationCheckIntlResponseBody {
	s.Result = v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceDuplicationCheckIntlResponseBodyResult struct {
	// Returns the face library face ID and UserID when a duplicate face is detected.
	//
	// example:
	//
	// [
	//
	// {\\"faceGroupCode\\":\\"sg7****uzt\\",\\"faceId\\":\\"f5a921*******9e792ec84c8f0ca592a\\"}
	//
	// ]
	DuplicateFace *string `json:"DuplicateFace,omitempty" xml:"DuplicateFace,omitempty"`
	// The estimated age of the face, which may not be returned if the prediction fails.
	//
	// example:
	//
	// 30
	FaceAge *string `json:"FaceAge,omitempty" xml:"FaceAge,omitempty"`
	// Indicates whether the captured face involves a liveness attack, Y for an attack, N for no attack.
	//
	// Returned when silent liveness detection is enabled.
	//
	// example:
	//
	// N
	FaceAttack *string `json:"FaceAttack,omitempty" xml:"FaceAttack,omitempty"`
	// The probability of a liveness attack detected by silent liveness detection. The value range is 0 to 100.
	//
	// Returned when silent liveness detection is enabled.
	//
	// example:
	//
	// 99
	FaceAttackScore *string `json:"FaceAttackScore,omitempty" xml:"FaceAttackScore,omitempty"`
	// When the verification mode is 1 or 2, returns the 1:1 verification comparison score
	//
	// Comparison score range 0～100.
	//
	// example:
	//
	// 98
	FaceComparisonScore *string `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	// The predicted gender of the face in the image, which may not be returned if the prediction fails.
	//
	// - M: Male
	//
	// - F: Female
	//
	// example:
	//
	// M
	FaceGender *string `json:"FaceGender,omitempty" xml:"FaceGender,omitempty"`
	// Final authentication result, values:
	//
	// - Y: Passed
	//
	// - N: Not passed
	//
	// example:
	//
	// Y
	FacePassed *string `json:"FacePassed,omitempty" xml:"FacePassed,omitempty"`
	// Returns the corresponding FACEID only when the customer sets auto-registration and the face registration is successful.
	//
	// example:
	//
	// 9e792ec84c8f0ca592a
	FaceRegistrationId *string `json:"FaceRegistrationId,omitempty" xml:"FaceRegistrationId,omitempty"`
	// Face registration result
	//
	// - 0- Failed
	//
	// - 1- Succeeded
	//
	// example:
	//
	// 0
	FaceRegistrationResult *int32 `json:"FaceRegistrationResult,omitempty" xml:"FaceRegistrationResult,omitempty"`
	// Description of the authentication result. For more information, see ResultObject.SubCode error code description.
	//
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// Unique identifier of the authentication request.
	//
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceDuplicationCheckIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceDuplicationCheckIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetDuplicateFace() *string {
	return s.DuplicateFace
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetFaceAge() *string {
	return s.FaceAge
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetFaceAttack() *string {
	return s.FaceAttack
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetFaceAttackScore() *string {
	return s.FaceAttackScore
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetFaceComparisonScore() *string {
	return s.FaceComparisonScore
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetFaceGender() *string {
	return s.FaceGender
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetFacePassed() *string {
	return s.FacePassed
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetFaceRegistrationId() *string {
	return s.FaceRegistrationId
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetFaceRegistrationResult() *int32 {
	return s.FaceRegistrationResult
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetDuplicateFace(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.DuplicateFace = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetFaceAge(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.FaceAge = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetFaceAttack(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.FaceAttack = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetFaceAttackScore(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.FaceAttackScore = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetFaceComparisonScore(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.FaceComparisonScore = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetFaceGender(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.FaceGender = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetFacePassed(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.FacePassed = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetFaceRegistrationId(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.FaceRegistrationId = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetFaceRegistrationResult(v int32) *FaceDuplicationCheckIntlResponseBodyResult {
	s.FaceRegistrationResult = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetSubCode(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) SetTransactionId(v string) *FaceDuplicationCheckIntlResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *FaceDuplicationCheckIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
