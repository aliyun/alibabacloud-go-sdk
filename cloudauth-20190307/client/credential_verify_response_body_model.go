// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *CredentialVerifyResponseBodyResultObject) *CredentialVerifyResponseBody
	GetResultObject() *CredentialVerifyResponseBodyResultObject
}

type CredentialVerifyResponseBody struct {
	// The response code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D6163397-15C5-419C-9ACC-B7C83E0B4C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
	ResultObject *CredentialVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CredentialVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialVerifyResponseBody) GetResultObject() *CredentialVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *CredentialVerifyResponseBody) SetCode(v string) *CredentialVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialVerifyResponseBody) SetMessage(v string) *CredentialVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialVerifyResponseBody) SetRequestId(v string) *CredentialVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialVerifyResponseBody) SetResultObject(v *CredentialVerifyResponseBodyResultObject) *CredentialVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *CredentialVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CredentialVerifyResponseBodyResultObject struct {
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
	// - **0**: Low risk.
	//
	// - **1**: High risk.
	//
	// - **2**: Suspicious.
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
	// - COLOR_PRINT: color print copy.
	//
	// - WEB_IMAGE: web image.
	//
	// - SAME_FACE: similar face.
	//
	// - SAME_BACKGROUND: similar background.
	//
	// example:
	//
	// PS,SCREEN_PHOTO
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
	// The authoritative verification details.
	//
	// example:
	//
	// **
	VerifyDetail *string `json:"VerifyDetail,omitempty" xml:"VerifyDetail,omitempty"`
	// The authoritative verification result.
	//
	// example:
	//
	// *
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	// The Qwen interpretation.
	VlResult *CredentialVerifyResponseBodyResultObjectVlResult `json:"VlResult,omitempty" xml:"VlResult,omitempty" type:"Struct"`
}

func (s CredentialVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CredentialVerifyResponseBodyResultObject) GetMaterialInfo() *string {
	return s.MaterialInfo
}

func (s *CredentialVerifyResponseBodyResultObject) GetOcrInfo() *string {
	return s.OcrInfo
}

func (s *CredentialVerifyResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *CredentialVerifyResponseBodyResultObject) GetRiskScore() map[string]*string {
	return s.RiskScore
}

func (s *CredentialVerifyResponseBodyResultObject) GetRiskTag() *string {
	return s.RiskTag
}

func (s *CredentialVerifyResponseBodyResultObject) GetVerifyDetail() *string {
	return s.VerifyDetail
}

func (s *CredentialVerifyResponseBodyResultObject) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *CredentialVerifyResponseBodyResultObject) GetVlResult() *CredentialVerifyResponseBodyResultObjectVlResult {
	return s.VlResult
}

func (s *CredentialVerifyResponseBodyResultObject) SetMaterialInfo(v string) *CredentialVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetOcrInfo(v string) *CredentialVerifyResponseBodyResultObject {
	s.OcrInfo = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetResult(v string) *CredentialVerifyResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetRiskScore(v map[string]*string) *CredentialVerifyResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetRiskTag(v string) *CredentialVerifyResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetVerifyDetail(v string) *CredentialVerifyResponseBodyResultObject {
	s.VerifyDetail = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetVerifyResult(v string) *CredentialVerifyResponseBodyResultObject {
	s.VerifyResult = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) SetVlResult(v *CredentialVerifyResponseBodyResultObjectVlResult) *CredentialVerifyResponseBodyResultObject {
	s.VlResult = v
	return s
}

func (s *CredentialVerifyResponseBodyResultObject) Validate() error {
	if s.VlResult != nil {
		if err := s.VlResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CredentialVerifyResponseBodyResultObjectVlResult struct {
	// Indicates whether the Qwen interpretation is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The Qwen interpretation content.
	//
	// example:
	//
	// {\\"这张图有没有明显的PS特征\\":0,\\"图片是否为正常经营照片\\":1,\\"图片中有没有58、美团、大众点评字样\\":0,\\"这张图有没有网站信息\\":0,\\"图片经营的场景是否为酒店\\":0}
	VlContent *string `json:"VlContent,omitempty" xml:"VlContent,omitempty"`
}

func (s CredentialVerifyResponseBodyResultObjectVlResult) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyResponseBodyResultObjectVlResult) GoString() string {
	return s.String()
}

func (s *CredentialVerifyResponseBodyResultObjectVlResult) GetSuccess() *bool {
	return s.Success
}

func (s *CredentialVerifyResponseBodyResultObjectVlResult) GetVlContent() *string {
	return s.VlContent
}

func (s *CredentialVerifyResponseBodyResultObjectVlResult) SetSuccess(v bool) *CredentialVerifyResponseBodyResultObjectVlResult {
	s.Success = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObjectVlResult) SetVlContent(v string) *CredentialVerifyResponseBodyResultObjectVlResult {
	s.VlContent = &v
	return s
}

func (s *CredentialVerifyResponseBodyResultObjectVlResult) Validate() error {
	return dara.Validate(s)
}
