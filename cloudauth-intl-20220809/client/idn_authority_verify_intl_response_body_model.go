// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdnAuthorityVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IdnAuthorityVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *IdnAuthorityVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *IdnAuthorityVerifyIntlResponseBody
	GetRequestId() *string
	SetResult(v *IdnAuthorityVerifyIntlResponseBodyResult) *IdnAuthorityVerifyIntlResponseBody
	GetResult() *IdnAuthorityVerifyIntlResponseBodyResult
}

type IdnAuthorityVerifyIntlResponseBody struct {
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
	// The returned result.
	Result *IdnAuthorityVerifyIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s IdnAuthorityVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IdnAuthorityVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *IdnAuthorityVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *IdnAuthorityVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IdnAuthorityVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IdnAuthorityVerifyIntlResponseBody) GetResult() *IdnAuthorityVerifyIntlResponseBodyResult {
	return s.Result
}

func (s *IdnAuthorityVerifyIntlResponseBody) SetCode(v string) *IdnAuthorityVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *IdnAuthorityVerifyIntlResponseBody) SetMessage(v string) *IdnAuthorityVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *IdnAuthorityVerifyIntlResponseBody) SetRequestId(v string) *IdnAuthorityVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *IdnAuthorityVerifyIntlResponseBody) SetResult(v *IdnAuthorityVerifyIntlResponseBodyResult) *IdnAuthorityVerifyIntlResponseBody {
	s.Result = v
	return s
}

func (s *IdnAuthorityVerifyIntlResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IdnAuthorityVerifyIntlResponseBodyResult struct {
	// The detailed verification results from the data source are described as follows (using the Indonesian data source as an example):
	//
	// - **govId, fullName, dob**: A comparison score equal to 1.0 indicates a complete match with the official data source. A score lower than 1.0 indicates a mismatch.
	//
	// - **selfiePhoto**: A comparison score greater than 0.8 indicates a match with the official data source. A score equal to or lower than 0.8 indicates a mismatch.
	//
	// - **liveness**: A score higher than 0.95 indicates a liveness detection risk.
	//
	// - **imgManipulationScore**: A score higher than 0.95 indicates an image tampering risk.
	//
	// example:
	//
	// {
	//
	//   "govId": 1.0,
	//
	//   "fullName": 1.0,
	//
	//   "dob": 0.9,
	//
	//   "selfiePhoto": 0.8777,
	//
	//   "liveness": 0.1152,
	//
	//   "imgManipulationScore": 0.2253
	//
	// }
	ExtSourceInfo *string `json:"ExtSourceInfo,omitempty" xml:"ExtSourceInfo,omitempty"`
	// Indicates whether the verification is passed. Valid values:
	//
	// - Y: passed.
	//
	// - N: not passed.
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
	// hk573be80f944d95ac812e0*******a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s IdnAuthorityVerifyIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s IdnAuthorityVerifyIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) GetExtSourceInfo() *string {
	return s.ExtSourceInfo
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) SetExtSourceInfo(v string) *IdnAuthorityVerifyIntlResponseBodyResult {
	s.ExtSourceInfo = &v
	return s
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) SetPassed(v string) *IdnAuthorityVerifyIntlResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) SetSubCode(v string) *IdnAuthorityVerifyIntlResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) SetTransactionId(v string) *IdnAuthorityVerifyIntlResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *IdnAuthorityVerifyIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
