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
	// 4EB35****87EBA1
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckVerifyLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CheckVerifyLogResponseBodyResult struct {
	// example:
	//
	// {}
	ExtInfo       *string   `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	InterruptPage *string   `json:"InterruptPage,omitempty" xml:"InterruptPage,omitempty"`
	LogInfo       []*string `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
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
	// 1001
	VerifyErrorCode *string `json:"VerifyErrorCode,omitempty" xml:"VerifyErrorCode,omitempty"`
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

func (s *CheckVerifyLogResponseBodyResult) GetLogInfo() []*string {
	return s.LogInfo
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

func (s *CheckVerifyLogResponseBodyResult) SetLogInfo(v []*string) *CheckVerifyLogResponseBodyResult {
	s.LogInfo = v
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
