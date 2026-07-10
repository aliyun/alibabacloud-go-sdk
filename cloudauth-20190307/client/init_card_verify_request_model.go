// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitCardVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackToken(v string) *InitCardVerifyRequest
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitCardVerifyRequest
	GetCallbackUrl() *string
	SetCardPageNumber(v string) *InitCardVerifyRequest
	GetCardPageNumber() *string
	SetCardType(v string) *InitCardVerifyRequest
	GetCardType() *string
	SetDocScanMode(v string) *InitCardVerifyRequest
	GetDocScanMode() *string
	SetMerchantBizId(v string) *InitCardVerifyRequest
	GetMerchantBizId() *string
	SetMetaInfo(v string) *InitCardVerifyRequest
	GetMetaInfo() *string
	SetModel(v string) *InitCardVerifyRequest
	GetModel() *string
	SetPictureSave(v string) *InitCardVerifyRequest
	GetPictureSave() *string
	SetVerifyMeta(v string) *InitCardVerifyRequest
	GetVerifyMeta() *string
}

type InitCardVerifyRequest struct {
	// The security token used for anti-replay and anti-tampering verification. If you specify this parameter, the CallbackToken field is included in the callback URL.
	//
	// example:
	//
	// NMjvQanQgplBSaEI0sL86WnQplB
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// The callback URL for authentication results. The URL must start with https. After the authentication is complete, the system sends a callback to this URL with the certifyId and passed fields automatically appended. Example: https://www.aliyun.com?certifyId=xxxx&passed=T
	//
	// > **Warning*	- The callback is triggered only when the authentication is complete. No notification is sent if the authentication is abandoned, interrupted, or not performed. After you receive the callback notification, call the query operation to obtain the authentication details if needed.
	//
	// example:
	//
	// https://www.aliyun.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The number of card pages to be collected by the SDK. Valid values:
	//
	// - 1: collects the front side only.
	//
	// - 2: collects both the front and back sides.
	//
	// - If the verification type is ID card validity period (VerifyMeta is set to ID_PERIOD), set this parameter to 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CardPageNumber *string `json:"CardPageNumber,omitempty" xml:"CardPageNumber,omitempty"`
	// The document type. Valid values:
	//
	// - IDENTITY_CARD: resident identity card.
	//
	// This parameter is required.
	//
	// example:
	//
	// IDENTITY_CARD
	CardType *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// The photo capture mode (manual or automatic). Valid values:
	//
	// - shoot: manual capture
	//
	// - scan: scan mode
	//
	// - auto: automatic switchover.
	//
	// example:
	//
	// shoot
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// A custom business unique identifier that you define for subsequent troubleshooting. The value is a combination of letters and digits up to 32 characters in length. Make sure the value is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The MetaInfo environment parameter. Obtain this value by using the client SDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"zimVer":"3.0.0","appVersion": "1","bioMetaInfo": "4.1.0:1150****,0","appName": "com.aliyun.antcloudauth","deviceType": "ios","osVersion": "iOS 10.3.2","apdidToken": "","deviceModel": "iPhone9,1"}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The verification mode. Valid values:
	//
	// - OCR_VERIFY: OCR recognition and authentication mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// OCR_VERIFY
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies whether to temporarily store images collected by the app. Valid values:
	//
	// - Y: Yes.
	//
	// - N: No.
	//
	// - If you set this parameter to Y, the query operation returns card image information.
	//
	// This parameter is required.
	//
	// example:
	//
	// Y
	PictureSave *string `json:"PictureSave,omitempty" xml:"PictureSave,omitempty"`
	// The verification type. Valid values:
	//
	// - ID_2_META: two-factor identity verification (name + ID card number).
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_2_META
	VerifyMeta *string `json:"VerifyMeta,omitempty" xml:"VerifyMeta,omitempty"`
}

func (s InitCardVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s InitCardVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitCardVerifyRequest) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitCardVerifyRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitCardVerifyRequest) GetCardPageNumber() *string {
	return s.CardPageNumber
}

func (s *InitCardVerifyRequest) GetCardType() *string {
	return s.CardType
}

func (s *InitCardVerifyRequest) GetDocScanMode() *string {
	return s.DocScanMode
}

func (s *InitCardVerifyRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *InitCardVerifyRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitCardVerifyRequest) GetModel() *string {
	return s.Model
}

func (s *InitCardVerifyRequest) GetPictureSave() *string {
	return s.PictureSave
}

func (s *InitCardVerifyRequest) GetVerifyMeta() *string {
	return s.VerifyMeta
}

func (s *InitCardVerifyRequest) SetCallbackToken(v string) *InitCardVerifyRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitCardVerifyRequest) SetCallbackUrl(v string) *InitCardVerifyRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitCardVerifyRequest) SetCardPageNumber(v string) *InitCardVerifyRequest {
	s.CardPageNumber = &v
	return s
}

func (s *InitCardVerifyRequest) SetCardType(v string) *InitCardVerifyRequest {
	s.CardType = &v
	return s
}

func (s *InitCardVerifyRequest) SetDocScanMode(v string) *InitCardVerifyRequest {
	s.DocScanMode = &v
	return s
}

func (s *InitCardVerifyRequest) SetMerchantBizId(v string) *InitCardVerifyRequest {
	s.MerchantBizId = &v
	return s
}

func (s *InitCardVerifyRequest) SetMetaInfo(v string) *InitCardVerifyRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitCardVerifyRequest) SetModel(v string) *InitCardVerifyRequest {
	s.Model = &v
	return s
}

func (s *InitCardVerifyRequest) SetPictureSave(v string) *InitCardVerifyRequest {
	s.PictureSave = &v
	return s
}

func (s *InitCardVerifyRequest) SetVerifyMeta(v string) *InitCardVerifyRequest {
	s.VerifyMeta = &v
	return s
}

func (s *InitCardVerifyRequest) Validate() error {
	return dara.Validate(s)
}
