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
	// Returned result information.
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
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CredentialVerifyIntlResponseBodyResultObject struct {
	// Other information in JSON format.
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
	// Risk result:
	//
	// - **0**: Low risk
	//
	// - **1**: High risk
	//
	// - **2**: Suspicious
	//
	// example:
	//
	// 1
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Risk score map
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// Risk tags, separated by commas (,). Includes:
	//
	// - PS: Image manipulation (Photoshop)
	//
	// - SCREEN_PHOTO: Screen recapture
	//
	// - SCREENSHOT: Screenshot
	//
	// - ORIGINAL_PHOTO: Not original image
	//
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
