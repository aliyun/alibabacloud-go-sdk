// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceGuardRiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FaceGuardRiskResponseBody
	GetCode() *string
	SetMessage(v string) *FaceGuardRiskResponseBody
	GetMessage() *string
	SetRequestId(v string) *FaceGuardRiskResponseBody
	GetRequestId() *string
	SetResult(v *FaceGuardRiskResponseBodyResult) *FaceGuardRiskResponseBody
	GetResult() *FaceGuardRiskResponseBodyResult
}

type FaceGuardRiskResponseBody struct {
	// The return code. A value of Success indicates that the API operation responded successfully. For more information about how to determine the authentication result, expand the **Return codes*	- section below.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A detailed description of the return code.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 595E387B-3F0E-5C52-BD02-8EFE63D41FD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object
	Result *FaceGuardRiskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceGuardRiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceGuardRiskResponseBody) GoString() string {
	return s.String()
}

func (s *FaceGuardRiskResponseBody) GetCode() *string {
	return s.Code
}

func (s *FaceGuardRiskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FaceGuardRiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceGuardRiskResponseBody) GetResult() *FaceGuardRiskResponseBodyResult {
	return s.Result
}

func (s *FaceGuardRiskResponseBody) SetCode(v string) *FaceGuardRiskResponseBody {
	s.Code = &v
	return s
}

func (s *FaceGuardRiskResponseBody) SetMessage(v string) *FaceGuardRiskResponseBody {
	s.Message = &v
	return s
}

func (s *FaceGuardRiskResponseBody) SetRequestId(v string) *FaceGuardRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceGuardRiskResponseBody) SetResult(v *FaceGuardRiskResponseBodyResult) *FaceGuardRiskResponseBody {
	s.Result = v
	return s
}

func (s *FaceGuardRiskResponseBody) Validate() error {
	return dara.Validate(s)
}

type FaceGuardRiskResponseBodyResult struct {
	// The device risk probability predicted by the Device Guard algorithm. A higher score indicates a higher device risk.
	//
	// Valid values: 0 to 100.
	//
	// example:
	//
	// 0
	GuardRiskScore *float64 `json:"GuardRiskScore,omitempty" xml:"GuardRiskScore,omitempty"`
	// Extended information. This is empty by default.
	//
	// example:
	//
	// {
	//
	//   "code": 200
	//
	//   "message":"success"
	//
	//   "umid":"07d3295d3d597b425d102a7f********",
	//
	//   "sip":"198.51.100.1",
	//
	//   "durationMs" : 4968931
	//
	//   "queryCount":1,
	//
	//   "querySessionCount":1,
	//
	//   "queryUmidCount":1
	//
	//   "platform":"Android
	//
	// }
	RiskExtends *string `json:"RiskExtends,omitempty" xml:"RiskExtends,omitempty"`
	// The device risk tags. Multiple risk tags are separated by commas (**,**). For more information about the risk tags and their meanings, expand the **Risk tags (RiskTags)*	- section below.
	//
	// example:
	//
	// ROOT,VPN,HOOK
	RiskTags *string `json:"RiskTags,omitempty" xml:"RiskTags,omitempty"`
	// The transaction ID.
	//
	// example:
	//
	// hk573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceGuardRiskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FaceGuardRiskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceGuardRiskResponseBodyResult) GetGuardRiskScore() *float64 {
	return s.GuardRiskScore
}

func (s *FaceGuardRiskResponseBodyResult) GetRiskExtends() *string {
	return s.RiskExtends
}

func (s *FaceGuardRiskResponseBodyResult) GetRiskTags() *string {
	return s.RiskTags
}

func (s *FaceGuardRiskResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *FaceGuardRiskResponseBodyResult) SetGuardRiskScore(v float64) *FaceGuardRiskResponseBodyResult {
	s.GuardRiskScore = &v
	return s
}

func (s *FaceGuardRiskResponseBodyResult) SetRiskExtends(v string) *FaceGuardRiskResponseBodyResult {
	s.RiskExtends = &v
	return s
}

func (s *FaceGuardRiskResponseBodyResult) SetRiskTags(v string) *FaceGuardRiskResponseBodyResult {
	s.RiskTags = &v
	return s
}

func (s *FaceGuardRiskResponseBodyResult) SetTransactionId(v string) *FaceGuardRiskResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *FaceGuardRiskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
