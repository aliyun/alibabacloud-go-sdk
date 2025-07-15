// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialVerifyIntlResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialVerifyIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialVerifyIntlResponseBody
	GetRequestId() *string
	SetResultObject(v *CredentialVerifyIntlResponseBodyResultObject) *CredentialVerifyIntlResponseBody
	GetResultObject() *CredentialVerifyIntlResponseBodyResultObject
}

type CredentialVerifyIntlResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *CredentialVerifyIntlResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CredentialVerifyIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyIntlResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialVerifyIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialVerifyIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialVerifyIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialVerifyIntlResponseBody) GetResultObject() *CredentialVerifyIntlResponseBodyResultObject {
	return s.ResultObject
}

func (s *CredentialVerifyIntlResponseBody) SetCode(v string) *CredentialVerifyIntlResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialVerifyIntlResponseBody) SetMessage(v string) *CredentialVerifyIntlResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialVerifyIntlResponseBody) SetRequestId(v string) *CredentialVerifyIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialVerifyIntlResponseBody) SetResultObject(v *CredentialVerifyIntlResponseBodyResultObject) *CredentialVerifyIntlResponseBody {
	s.ResultObject = v
	return s
}

func (s *CredentialVerifyIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type CredentialVerifyIntlResponseBodyResultObject struct {
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// example:
	//
	// 1
	Result    *string            `json:"Result,omitempty" xml:"Result,omitempty"`
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// example:
	//
	// PS
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
}

func (s CredentialVerifyIntlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyIntlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CredentialVerifyIntlResponseBodyResultObject) GetMaterialInfo() *string {
	return s.MaterialInfo
}

func (s *CredentialVerifyIntlResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *CredentialVerifyIntlResponseBodyResultObject) GetRiskScore() map[string]*string {
	return s.RiskScore
}

func (s *CredentialVerifyIntlResponseBodyResultObject) GetRiskTag() *string {
	return s.RiskTag
}

func (s *CredentialVerifyIntlResponseBodyResultObject) SetMaterialInfo(v string) *CredentialVerifyIntlResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *CredentialVerifyIntlResponseBodyResultObject) SetResult(v string) *CredentialVerifyIntlResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *CredentialVerifyIntlResponseBodyResultObject) SetRiskScore(v map[string]*string) *CredentialVerifyIntlResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *CredentialVerifyIntlResponseBodyResultObject) SetRiskTag(v string) *CredentialVerifyIntlResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

func (s *CredentialVerifyIntlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
