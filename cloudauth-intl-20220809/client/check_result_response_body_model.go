// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckResultResponseBody
	GetCode() *string
	SetMessage(v string) *CheckResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckResultResponseBody
	GetRequestId() *string
	SetResult(v *CheckResultResponseBodyResult) *CheckResultResponseBody
	GetResult() *CheckResultResponseBodyResult
}

type CheckResultResponseBody struct {
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
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	Result *CheckResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckResultResponseBody) GetResult() *CheckResultResponseBodyResult {
	return s.Result
}

func (s *CheckResultResponseBody) SetCode(v string) *CheckResultResponseBody {
	s.Code = &v
	return s
}

func (s *CheckResultResponseBody) SetMessage(v string) *CheckResultResponseBody {
	s.Message = &v
	return s
}

func (s *CheckResultResponseBody) SetRequestId(v string) *CheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResultResponseBody) SetResult(v *CheckResultResponseBodyResult) *CheckResultResponseBody {
	s.Result = v
	return s
}

func (s *CheckResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckResultResponseBodyResult struct {
	// Authentication result.
	//
	// example:
	//
	// **
	EkycResult *string `json:"EkycResult,omitempty" xml:"EkycResult,omitempty"`
	// Extended basic information.
	//
	// example:
	//
	// **
	ExtBasicInfo *string `json:"ExtBasicInfo,omitempty" xml:"ExtBasicInfo,omitempty"`
	// Face information.
	//
	// example:
	//
	// **
	ExtFaceInfo *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
	// ID information.
	//
	// example:
	//
	// **
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// Extended information
	//
	// example:
	//
	// {}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// Risk information.
	//
	// example:
	//
	// **
	ExtRiskInfo *string `json:"ExtRiskInfo,omitempty" xml:"ExtRiskInfo,omitempty"`
	// Whether the authentication is passed.
	//
	// - Y: Passed
	//
	// - N: Not passed
	//
	// example:
	//
	// Y
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// Sub-result code.
	//
	// example:
	//
	// ***
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s CheckResultResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CheckResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckResultResponseBodyResult) GetEkycResult() *string {
	return s.EkycResult
}

func (s *CheckResultResponseBodyResult) GetExtBasicInfo() *string {
	return s.ExtBasicInfo
}

func (s *CheckResultResponseBodyResult) GetExtFaceInfo() *string {
	return s.ExtFaceInfo
}

func (s *CheckResultResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *CheckResultResponseBodyResult) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *CheckResultResponseBodyResult) GetExtRiskInfo() *string {
	return s.ExtRiskInfo
}

func (s *CheckResultResponseBodyResult) GetPassed() *string {
	return s.Passed
}

func (s *CheckResultResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *CheckResultResponseBodyResult) SetEkycResult(v string) *CheckResultResponseBodyResult {
	s.EkycResult = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtBasicInfo(v string) *CheckResultResponseBodyResult {
	s.ExtBasicInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtFaceInfo(v string) *CheckResultResponseBodyResult {
	s.ExtFaceInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtIdInfo(v string) *CheckResultResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtInfo(v string) *CheckResultResponseBodyResult {
	s.ExtInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtRiskInfo(v string) *CheckResultResponseBodyResult {
	s.ExtRiskInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetPassed(v string) *CheckResultResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetSubCode(v string) *CheckResultResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CheckResultResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
