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
	// Security Token, used for anti-replay and anti-tampering checks. If this parameter is passed, the CallbackToken field will be displayed in the callback address.
	//
	// example:
	//
	// NMjvQanQgplBSaEI0sL86WnQplB
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// - The callback notification address for the authentication result, which must start with https.
	//
	// - The platform will call back this address after completing the authentication and automatically add the certifyId and passed fields, example: https://www.aliyun.com?certifyId=xxxx&passed=T
	//
	// - Warning
	//
	// The callback is triggered only when the authentication is completed. If the authentication is abandoned, interrupted abnormally, or not performed, no notification will be sent. It is recommended that when you receive the callback notification, if necessary, you can obtain detailed authentication information through the query interface.
	//
	// example:
	//
	// https://www.aliyun.com
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// Number of card pages collected by the SDK
	//
	// - You can input 1 or 2; input 1 to collect the front side, input 2 to collect both the front and back sides.
	//
	// - If the verification type is ID period (VerifyMeta value is ID_PERIOD), you must input 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CardPageNumber *string `json:"CardPageNumber,omitempty" xml:"CardPageNumber,omitempty"`
	// Type of identification
	//
	// - Resident Second Generation ID Card: IDENTITY_CARD
	//
	// This parameter is required.
	//
	// example:
	//
	// IDENTITY_CARD
	CardType *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// Enumeration of photo-taking methods (manual/auto)
	//
	// - Take a photo: shoot
	//
	// - Scan: scan
	//
	// - Auto switch: auto
	//
	// example:
	//
	// shoot
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// A unique business identifier you define, used for subsequent troubleshooting.
	//
	// Supports a combination of 32 alphanumeric characters, please ensure uniqueness.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// MetaInfo environment parameter, which needs to be obtained through the client SDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"zimVer":"3.0.0","appVersion": "1","bioMetaInfo": "4.1.0:1150****,0","appName": "com.aliyun.antcloudauth","deviceType": "ios","osVersion": "iOS 10.3.2","apdidToken": "","deviceModel": "iPhone9,1"}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// Verification method, value:
	//
	// - OCR_VERIFY: OCR recognition and verification mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// OCR_VERIFY
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Whether to temporarily store the images collected by the app.
	//
	// - Y: Yes
	//
	// - N: No
	//
	// - If \\"Yes\\" is selected here, the query interface will support returning the card image information.
	//
	// This parameter is required.
	//
	// example:
	//
	// Y
	PictureSave *string `json:"PictureSave,omitempty" xml:"PictureSave,omitempty"`
	// Verification type, value:
	//
	// - Identity two elements (name + ID number): ID_2_META
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
