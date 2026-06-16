// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVerifyLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckVerifyLogResponseBody
	GetCode() *string
	SetMessage(v string) *CheckVerifyLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckVerifyLogResponseBody
	GetRequestId() *string
	SetResult(v *CheckVerifyLogResponseBodyResult) *CheckVerifyLogResponseBody
	GetResult() *CheckVerifyLogResponseBodyResult
}

type CheckVerifyLogResponseBody struct {
	// The backend error code.
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
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *CheckVerifyLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckVerifyLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckVerifyLogResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVerifyLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckVerifyLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckVerifyLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckVerifyLogResponseBody) GetResult() *CheckVerifyLogResponseBodyResult {
	return s.Result
}

func (s *CheckVerifyLogResponseBody) SetCode(v string) *CheckVerifyLogResponseBody {
	s.Code = &v
	return s
}

func (s *CheckVerifyLogResponseBody) SetMessage(v string) *CheckVerifyLogResponseBody {
	s.Message = &v
	return s
}

func (s *CheckVerifyLogResponseBody) SetRequestId(v string) *CheckVerifyLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckVerifyLogResponseBody) SetResult(v *CheckVerifyLogResponseBodyResult) *CheckVerifyLogResponseBody {
	s.Result = v
	return s
}

func (s *CheckVerifyLogResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckVerifyLogResponseBodyResult struct {
	// The extended information.
	//
	// example:
	//
	// {}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// The last page where the authentication was interrupted. Valid values:
	//
	// - Page not started
	//
	// - OCR guide page
	//
	// - OCR camera authorization
	//
	// - OCR document capture page
	//
	// - OCR recognition loading
	//
	// - OCR recognition result editing page
	//
	// - OCR recognition result editing loading
	//
	// - Liveness detection guide page
	//
	// - Liveness detection camera authorization page
	//
	// - Liveness detection page
	//
	// - Liveness detection downgrade page
	//
	// - Liveness detection retry
	//
	// - Liveness detection loading.
	//
	// example:
	//
	// OCR拍摄证件页面
	InterruptPage *string `json:"InterruptPage,omitempty" xml:"InterruptPage,omitempty"`
	// The last page where the authentication was interrupted, in English. Valid values:
	//
	// - LOADING
	//
	// - GUIDE
	//
	// - FACE
	//
	// - OCR_SCAN
	//
	// - OCR_RESULT
	//
	// - NFC_INPUT
	//
	// - NFC_READ.
	//
	// example:
	//
	// LOADING
	InterruptPageEn *string `json:"InterruptPageEn,omitempty" xml:"InterruptPageEn,omitempty"`
	// The SDK runtime log details.
	LogInfo []*string `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
	// The SDK runtime trace log details in English. The format of this field is the same as **LogInfo**.
	LogInfoEn []*string `json:"LogInfoEn,omitempty" xml:"LogInfoEn,omitempty" type:"Repeated"`
	// The SDK runtime log statistics details.
	//
	// example:
	//
	// {
	//
	//           "faceOverTimes": 0,
	//
	//           "hasFaceOverTimes": false,
	//
	//           "hasFacePermissionRefuse": false,
	//
	//           "hasOcrEdit": true,
	//
	//           "hasOcrEditOverTimes": false,
	//
	//           "hasOcrOverTimes": true,
	//
	//           "hasOcrPermissionRefuse": false,
	//
	//           "ocrEditOverTimes": 0,
	//
	//           "ocrEditTimes": 1,
	//
	//           "ocrOverTimes": 1,
	//
	//           "pageStayTimeInfo": {
	//
	//             "LOADING": "1615",
	//
	//             "OCR_SCAN": "37446",
	//
	//             "OCR_RESULT": "1338",
	//
	//             "FACE": "8707"
	//
	//           }
	//
	//         }
	LogStatisticsInfo *string `json:"LogStatisticsInfo,omitempty" xml:"LogStatisticsInfo,omitempty"`
	// Indicates whether the authentication is passed. Valid values:
	//
	// - Y: Passed.
	//
	// - N: Not passed.
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
	// The error code for authentication interruption. Valid values:
	//
	// - 1000: The user completed the face verification process, and the authentication result is passed.
	//
	// - 1001: The user completed the face verification process, and the authentication result is not passed.
	//
	// - 1002: System error.
	//
	// - 1003: SDK initialization failed. Check whether the client time is correct.
	//
	// - 1004: Camera permission error.
	//
	// - 1005: Network error.
	//
	// - 1006: The user exited.
	//
	// - 1007: Invalid TransactionId.
	//
	// - 1009: Client timestamp error.
	//
	// - 1011: Incorrect document type submitted.
	//
	// - 1012: Key information of the recognized document is missing or format validation failed.
	//
	// - 1013: Poor image quality.
	//
	// - 1014: The number of errors exceeded the upper limit.
	//
	// - 1015: The Android system version is too low.
	//
	// - 1016: Camera permission not obtained.
	//
	// - 9999: The authentication process is suspected to be interrupted.
	//
	// example:
	//
	// 1001
	VerifyErrorCode *string `json:"VerifyErrorCode,omitempty" xml:"VerifyErrorCode,omitempty"`
	// The authentication status. Valid values:
	//
	// - 0: finished. The authentication is complete.
	//
	// - 1: unfinished. The authentication is interrupted.
	//
	// - 2: notstart. The authentication has not started.
	//
	// example:
	//
	// 1
	VerifyStatus *string `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
}

func (s CheckVerifyLogResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CheckVerifyLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckVerifyLogResponseBodyResult) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *CheckVerifyLogResponseBodyResult) GetInterruptPage() *string {
	return s.InterruptPage
}

func (s *CheckVerifyLogResponseBodyResult) GetInterruptPageEn() *string {
	return s.InterruptPageEn
}

func (s *CheckVerifyLogResponseBodyResult) GetLogInfo() []*string {
	return s.LogInfo
}

func (s *CheckVerifyLogResponseBodyResult) GetLogInfoEn() []*string {
	return s.LogInfoEn
}

func (s *CheckVerifyLogResponseBodyResult) GetLogStatisticsInfo() *string {
	return s.LogStatisticsInfo
}

func (s *CheckVerifyLogResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *CheckVerifyLogResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *CheckVerifyLogResponseBodyResult) GetVerifyErrorCode() *string {
	return s.VerifyErrorCode
}

func (s *CheckVerifyLogResponseBodyResult) GetVerifyStatus() *string {
	return s.VerifyStatus
}

func (s *CheckVerifyLogResponseBodyResult) SetExtInfo(v string) *CheckVerifyLogResponseBodyResult {
	s.ExtInfo = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetInterruptPage(v string) *CheckVerifyLogResponseBodyResult {
	s.InterruptPage = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetInterruptPageEn(v string) *CheckVerifyLogResponseBodyResult {
	s.InterruptPageEn = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetLogInfo(v []*string) *CheckVerifyLogResponseBodyResult {
	s.LogInfo = v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetLogInfoEn(v []*string) *CheckVerifyLogResponseBodyResult {
	s.LogInfoEn = v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetLogStatisticsInfo(v string) *CheckVerifyLogResponseBodyResult {
	s.LogStatisticsInfo = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetPassed(v string) *CheckVerifyLogResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetSubCode(v string) *CheckVerifyLogResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetVerifyErrorCode(v string) *CheckVerifyLogResponseBodyResult {
	s.VerifyErrorCode = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) SetVerifyStatus(v string) *CheckVerifyLogResponseBodyResult {
	s.VerifyStatus = &v
	return s
}

func (s *CheckVerifyLogResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
