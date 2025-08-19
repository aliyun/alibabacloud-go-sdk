// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialProductVerifyV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialProductVerifyV2ResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialProductVerifyV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialProductVerifyV2ResponseBody
	GetRequestId() *string
	SetResultObject(v *CredentialProductVerifyV2ResponseBodyResultObject) *CredentialProductVerifyV2ResponseBody
	GetResultObject() *CredentialProductVerifyV2ResponseBodyResultObject
}

type CredentialProductVerifyV2ResponseBody struct {
	// Return code: 200 for success, others for failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object.
	ResultObject *CredentialProductVerifyV2ResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CredentialProductVerifyV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialProductVerifyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialProductVerifyV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialProductVerifyV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialProductVerifyV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialProductVerifyV2ResponseBody) GetResultObject() *CredentialProductVerifyV2ResponseBodyResultObject {
	return s.ResultObject
}

func (s *CredentialProductVerifyV2ResponseBody) SetCode(v string) *CredentialProductVerifyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialProductVerifyV2ResponseBody) SetMessage(v string) *CredentialProductVerifyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialProductVerifyV2ResponseBody) SetRequestId(v string) *CredentialProductVerifyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialProductVerifyV2ResponseBody) SetResultObject(v *CredentialProductVerifyV2ResponseBodyResultObject) *CredentialProductVerifyV2ResponseBody {
	s.ResultObject = v
	return s
}

func (s *CredentialProductVerifyV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type CredentialProductVerifyV2ResponseBodyResultObject struct {
	// Additional information in JSON format.
	//
	// example:
	//
	// {
	//
	//  "sameBackgroundDetail": {
	//
	//  // 相似背景对于的原始图请求RequestId
	//
	//  "originalRequestId": "130A2C10-B9EE-4D84-88E3-5384FF03****";
	//
	//  // 相似背景对于的原始图请求商户ID
	//
	//  "originalMerchantId": "xxxxxxxx"
	//
	//  }
	//
	// }
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// - 0: Low risk
	//
	// - 1: High risk
	//
	// - 2: Suspicious
	//
	// example:
	//
	// 1
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Map of risk scores.
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// Risk tags, separated by commas, including:
	//
	// - PS: Image has been photoshopped
	//
	// - LOW_QUALITY_PRODUCT: Low quality (low clarity)
	//
	// - SAME_BACKGROUND: Similar background
	//
	// example:
	//
	// PS,LOW_QUALITY_PRODUCT
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
}

func (s CredentialProductVerifyV2ResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CredentialProductVerifyV2ResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) GetMaterialInfo() *string {
	return s.MaterialInfo
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) GetRiskScore() map[string]*string {
	return s.RiskScore
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) GetRiskTag() *string {
	return s.RiskTag
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) SetMaterialInfo(v string) *CredentialProductVerifyV2ResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) SetResult(v string) *CredentialProductVerifyV2ResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) SetRiskScore(v map[string]*string) *CredentialProductVerifyV2ResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) SetRiskTag(v string) *CredentialProductVerifyV2ResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

func (s *CredentialProductVerifyV2ResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
