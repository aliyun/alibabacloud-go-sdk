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
	// example:
	//
	// NMjvQanQgplBSaEI0sL86WnQplB
	CallbackToken *string `json:"CallbackToken,omitempty" xml:"CallbackToken,omitempty"`
	// example:
	//
	// https://www.aliyun.com?callbackToken=100000****&certifyId=shaxxxx&subCode=200
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CardPageNumber *string `json:"CardPageNumber,omitempty" xml:"CardPageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// IDENTITY_CARD
	CardType *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// example:
	//
	// shoot
	DocScanMode *string `json:"DocScanMode,omitempty" xml:"DocScanMode,omitempty"`
	// example:
	//
	// Y
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ID_OCR
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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
