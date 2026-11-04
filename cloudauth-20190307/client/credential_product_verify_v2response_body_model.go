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
	// The return code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result object.
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
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CredentialProductVerifyV2ResponseBodyResultObject struct {
	// The additional information in JSON format.
	//
	// example:
	//
	// {
	//
	//  "sameBackgroundDetail": {
	//
	//  // The request ID of the original image request for the similar background
	//
	//  "originalRequestId": "130A2C10-B9EE-4D84-88E3-5384FF03****",
	//
	//  // The merchant ID of the original image request for the similar background
	//
	//  "originalMerchantId": "xxxxxxxx"
	//
	//  }
	//
	// }
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// The risk level. Valid values:
	//
	// - 0: Low risk.
	//
	// - 1: High risk.
	//
	// - 2: Suspicious.
	//
	// example:
	//
	// 1
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The risk score map.
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// The risk tags. Multiple tags are separated by commas. Valid values:
	//
	// - PS: Image manipulation.
	//
	// - LOW_QUALITY_PRODUCT: Low quality (low clarity).
	//
	// - SAME_BACKGROUND: Similar image.
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
