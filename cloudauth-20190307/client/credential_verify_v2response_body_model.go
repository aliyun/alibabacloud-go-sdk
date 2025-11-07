// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialVerifyV2ResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialVerifyV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialVerifyV2ResponseBody
	GetRequestId() *string
	SetResultObject(v *CredentialVerifyV2ResponseBodyResultObject) *CredentialVerifyV2ResponseBody
	GetResultObject() *CredentialVerifyV2ResponseBodyResultObject
}

type CredentialVerifyV2ResponseBody struct {
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
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *CredentialVerifyV2ResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CredentialVerifyV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialVerifyV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialVerifyV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialVerifyV2ResponseBody) GetResultObject() *CredentialVerifyV2ResponseBodyResultObject {
	return s.ResultObject
}

func (s *CredentialVerifyV2ResponseBody) SetCode(v string) *CredentialVerifyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialVerifyV2ResponseBody) SetMessage(v string) *CredentialVerifyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialVerifyV2ResponseBody) SetRequestId(v string) *CredentialVerifyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialVerifyV2ResponseBody) SetResultObject(v *CredentialVerifyV2ResponseBodyResultObject) *CredentialVerifyV2ResponseBody {
	s.ResultObject = v
	return s
}

func (s *CredentialVerifyV2ResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CredentialVerifyV2ResponseBodyResultObject struct {
	// Additional information in JSON format.
	//
	// example:
	//
	// {
	//
	// "sameBackgroundDetail": {
	//
	// // 相似背景对于的原始图请求RequestId
	//
	// "originalRequestId": "130A2C10-B9EE-4D84-88E3-5384FF03****";
	//
	// // 相似背景对于的原始图请求商户ID
	//
	// "originalMerchantId": "xxxxxxxx"
	//
	// }
	//
	// }
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// OCR recognition result.
	//
	// example:
	//
	// {
	//
	//    "certNo":"20216*********",
	//
	//   "certType":"小学教师资格",
	//
	//    "gender":"男",
	//
	//    "subject":"美术",
	//
	//     "name":"李**",
	//
	//     "ext_info":"{}",
	//
	//     "birthDate":"1998-07-28",
	//
	//     "idNo":"620************",
	//
	//     "certOrg":""
	//
	// }
	OcrInfo *string `json:"OcrInfo,omitempty" xml:"OcrInfo,omitempty"`
	// Risk result
	//
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
	// Risk score map.
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// Risk tags, separated by commas (,), including:
	//
	// - PS: Image manipulation.
	//
	// - SCREEN_PHOTO: Screen recapture.
	//
	// - SCREENSHOT: Screenshot.
	//
	// - WATERMARK: Watermark.
	//
	// - SAME_BACKGROUND: Similar background.
	//
	// - ORIGINAL_PHOTO: Not the original image
	//
	// example:
	//
	// PS,SCREEN_PHOTO
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
	// Authority verification details.
	//
	// example:
	//
	// **
	VerifyDetail *string `json:"VerifyDetail,omitempty" xml:"VerifyDetail,omitempty"`
	// Authority verification result
	//
	// example:
	//
	// *
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	// Qwen interpretation.
	VlResult *CredentialVerifyV2ResponseBodyResultObjectVlResult `json:"VlResult,omitempty" xml:"VlResult,omitempty" type:"Struct"`
}

func (s CredentialVerifyV2ResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2ResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2ResponseBodyResultObject) GetMaterialInfo() *string {
	return s.MaterialInfo
}

func (s *CredentialVerifyV2ResponseBodyResultObject) GetOcrInfo() *string {
	return s.OcrInfo
}

func (s *CredentialVerifyV2ResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *CredentialVerifyV2ResponseBodyResultObject) GetRiskScore() map[string]*string {
	return s.RiskScore
}

func (s *CredentialVerifyV2ResponseBodyResultObject) GetRiskTag() *string {
	return s.RiskTag
}

func (s *CredentialVerifyV2ResponseBodyResultObject) GetVerifyDetail() *string {
	return s.VerifyDetail
}

func (s *CredentialVerifyV2ResponseBodyResultObject) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *CredentialVerifyV2ResponseBodyResultObject) GetVlResult() *CredentialVerifyV2ResponseBodyResultObjectVlResult {
	return s.VlResult
}

func (s *CredentialVerifyV2ResponseBodyResultObject) SetMaterialInfo(v string) *CredentialVerifyV2ResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObject) SetOcrInfo(v string) *CredentialVerifyV2ResponseBodyResultObject {
	s.OcrInfo = &v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObject) SetResult(v string) *CredentialVerifyV2ResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObject) SetRiskScore(v map[string]*string) *CredentialVerifyV2ResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObject) SetRiskTag(v string) *CredentialVerifyV2ResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObject) SetVerifyDetail(v string) *CredentialVerifyV2ResponseBodyResultObject {
	s.VerifyDetail = &v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObject) SetVerifyResult(v string) *CredentialVerifyV2ResponseBodyResultObject {
	s.VerifyResult = &v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObject) SetVlResult(v *CredentialVerifyV2ResponseBodyResultObjectVlResult) *CredentialVerifyV2ResponseBodyResultObject {
	s.VlResult = v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObject) Validate() error {
	if s.VlResult != nil {
		if err := s.VlResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CredentialVerifyV2ResponseBodyResultObjectVlResult struct {
	// Qwen interpretation success indicator
	//
	// true: Success
	//
	// false: Failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Image understanding result:
	//
	// - When PromptModel is DEFAULT, the output format refers to the example on the right.
	//
	// - When PromptModel is CUSTOM, the output format follows the agreed format of the Prompt.
	//
	// example:
	//
	// {
	//
	//  "CHK_DOOR_PHOTO":1, -- 是否门头照 1：是 0：否
	//
	//  "CHK_INTERIOR_PHOTO":0, -- 是否内景照 1：是 0：否
	//
	//  "CHK_COUNTER_PHOTO":0 -- 是否柜台照 1：是 0：否
	//
	// }
	VlContent *string `json:"VlContent,omitempty" xml:"VlContent,omitempty"`
}

func (s CredentialVerifyV2ResponseBodyResultObjectVlResult) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2ResponseBodyResultObjectVlResult) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2ResponseBodyResultObjectVlResult) GetSuccess() *bool {
	return s.Success
}

func (s *CredentialVerifyV2ResponseBodyResultObjectVlResult) GetVlContent() *string {
	return s.VlContent
}

func (s *CredentialVerifyV2ResponseBodyResultObjectVlResult) SetSuccess(v bool) *CredentialVerifyV2ResponseBodyResultObjectVlResult {
	s.Success = &v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObjectVlResult) SetVlContent(v string) *CredentialVerifyV2ResponseBodyResultObjectVlResult {
	s.VlContent = &v
	return s
}

func (s *CredentialVerifyV2ResponseBodyResultObjectVlResult) Validate() error {
	return dara.Validate(s)
}
