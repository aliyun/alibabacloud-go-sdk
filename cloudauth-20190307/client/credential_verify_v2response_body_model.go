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
	// The return code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
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
	// The additional information in JSON format.
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
	// The OCR recognition result.
	//
	// 	Danger: Deprecated.
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
	// The risk result. Valid values:
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
	// The risk tags, separated by commas (,). Valid values:
	//
	// - PS: image manipulation.
	//
	// - SCREEN_PHOTO: screen recapture.
	//
	// - SCREENSHOT: screenshot.
	//
	// - WATERMARK: watermark.
	//
	// - SAME_BACKGROUND: similar background.
	//
	// - ORIGINAL_PHOTO: non-original image.
	//
	// example:
	//
	// PS,SCREEN_PHOTO
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
	// The authoritative verification details.
	//
	// 	Danger: Deprecated.
	//
	// example:
	//
	// **
	VerifyDetail *string `json:"VerifyDetail,omitempty" xml:"VerifyDetail,omitempty"`
	// The authoritative verification result.
	//
	// 	Danger: Deprecated.
	//
	// example:
	//
	// *
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	// This feature is offline. This parameter no longer takes effect.
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
	// This feature is offline. This parameter no longer takes effect.
	//
	// example:
	//
	// -
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// This feature is offline. This parameter no longer takes effect.
	//
	// example:
	//
	// -
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
