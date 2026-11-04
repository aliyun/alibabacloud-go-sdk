// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitAuthVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackToken(v string) *InitAuthVerifyRequest
	GetCallbackToken() *string
	SetCallbackUrl(v string) *InitAuthVerifyRequest
	GetCallbackUrl() *string
	SetCardPageNumber(v string) *InitAuthVerifyRequest
	GetCardPageNumber() *string
	SetCardType(v string) *InitAuthVerifyRequest
	GetCardType() *string
	SetDocScanMode(v string) *InitAuthVerifyRequest
	GetDocScanMode() *string
	SetIdSpoof(v string) *InitAuthVerifyRequest
	GetIdSpoof() *string
	SetMetaInfo(v string) *InitAuthVerifyRequest
	GetMetaInfo() *string
	SetOuterOrderNo(v string) *InitAuthVerifyRequest
	GetOuterOrderNo() *string
	SetProductCode(v string) *InitAuthVerifyRequest
	GetProductCode() *string
	SetSceneId(v int64) *InitAuthVerifyRequest
	GetSceneId() *int64
}

type InitAuthVerifyRequest struct {
	// The security token that you generate for anti-replay and anti-tampering verification.
	//
	// If this parameter is specified, the CallbackToken field is included in the callback request sent to CallbackUrl.
	//
	// example:
	//
	// NMjvQanQgplBSaEI0sL86WnQplB
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// The callback URL for OCR results. The default callback request method is GET. The callback URL must start with https. After OCR recognition is completed, a callback request is sent to this URL, and the certifyId and subcode fields are automatically appended.
	//
	// Warning
	//
	// - The accessibility of this URL is verified before the API is invoked. If the specified URL does not support public network access, a 400 error is returned.
	//
	// - The callback is executed immediately after the OCR recognition is invoked, but may be delayed due to network or other issues. Accept the request completion notification on the client side first, and then invoke the query API to obtain the result details.
	//
	// example:
	//
	// https://www.aliyun.com?callbackToken=100000****&certifyId=shaxxxx&subCode=200
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The number of pages of the identity document to be collected by the SDK. Valid values:
	//
	// - 1: the front side.
	//
	// - 2: both the front and back sides.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CardPageNumber *string `json:"CardPageNumber,omitempty" xml:"CardPageNumber,omitempty"`
	// The type of the identity document. Valid values: IDENTITY_CARD.
	//
	// This parameter is required.
	//
	// example:
	//
	// IDENTITY_CARD
	CardType *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// The OCR identity document scanning pattern. Valid values:
	//
	// - shoot: takes a photo. This is the default value.
	//
	// - scan: scans the document.
	//
	// - auto: automatic switchover between photo and scan patterns.
	//
	// example:
	//
	// shoot
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// Specifies whether to enable the anti-counterfeiting detection feature for the identity document. Valid values:
	//
	// - Y: Enabled.
	//
	// - N: Disabled. This is the default value.
	//
	// example:
	//
	// Y
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// The MetaInfo environment parameter, which must be obtained through the client SDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "zimVer": "3.0.0",
	//
	//   "appVersion": "1",
	//
	//   "bioMetaInfo": "4.1.0:1150****,0",
	//
	//   "appName": "com.aliyun.antcloudauth",
	//
	//   "deviceType": "ios",
	//
	//   "osVersion": "iOS 10.3.2",
	//
	//   "apdidToken": "",
	//
	//   "deviceModel": "iPhone9,1"
	//
	// }
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The custom unique business identifier used for subsequent troubleshooting.
	//
	// The value can contain uppercase and lowercase letters and digits, and can be up to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// The product solution to be integrated. Valid values: ID_OCR.
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_OCR
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the authentication scene.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000002996
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s InitAuthVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s InitAuthVerifyRequest) GoString() string {
	return s.String()
}

func (s *InitAuthVerifyRequest) GetCallbackToken() *string {
	return s.CallbackToken
}

func (s *InitAuthVerifyRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *InitAuthVerifyRequest) GetCardPageNumber() *string {
	return s.CardPageNumber
}

func (s *InitAuthVerifyRequest) GetCardType() *string {
	return s.CardType
}

func (s *InitAuthVerifyRequest) GetDocScanMode() *string {
	return s.DocScanMode
}

func (s *InitAuthVerifyRequest) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *InitAuthVerifyRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *InitAuthVerifyRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *InitAuthVerifyRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *InitAuthVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *InitAuthVerifyRequest) SetCallbackToken(v string) *InitAuthVerifyRequest {
	s.CallbackToken = &v
	return s
}

func (s *InitAuthVerifyRequest) SetCallbackUrl(v string) *InitAuthVerifyRequest {
	s.CallbackUrl = &v
	return s
}

func (s *InitAuthVerifyRequest) SetCardPageNumber(v string) *InitAuthVerifyRequest {
	s.CardPageNumber = &v
	return s
}

func (s *InitAuthVerifyRequest) SetCardType(v string) *InitAuthVerifyRequest {
	s.CardType = &v
	return s
}

func (s *InitAuthVerifyRequest) SetDocScanMode(v string) *InitAuthVerifyRequest {
	s.DocScanMode = &v
	return s
}

func (s *InitAuthVerifyRequest) SetIdSpoof(v string) *InitAuthVerifyRequest {
	s.IdSpoof = &v
	return s
}

func (s *InitAuthVerifyRequest) SetMetaInfo(v string) *InitAuthVerifyRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitAuthVerifyRequest) SetOuterOrderNo(v string) *InitAuthVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *InitAuthVerifyRequest) SetProductCode(v string) *InitAuthVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *InitAuthVerifyRequest) SetSceneId(v int64) *InitAuthVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *InitAuthVerifyRequest) Validate() error {
	return dara.Validate(s)
}
