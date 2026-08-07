// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAuthCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GenerateAuthCodeResponseBody
	GetAuthCode() *string
	SetRequestId(v string) *GenerateAuthCodeResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GenerateAuthCodeResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GenerateAuthCodeResponseBody
	GetVendorType() *string
}

type GenerateAuthCodeResponseBody struct {
	// example:
	//
	// temporary-auth-code
	AuthCode *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GenerateAuthCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAuthCodeResponseBody) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GenerateAuthCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAuthCodeResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GenerateAuthCodeResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GenerateAuthCodeResponseBody) SetAuthCode(v string) *GenerateAuthCodeResponseBody {
	s.AuthCode = &v
	return s
}

func (s *GenerateAuthCodeResponseBody) SetRequestId(v string) *GenerateAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAuthCodeResponseBody) SetVendorRequestId(v string) *GenerateAuthCodeResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GenerateAuthCodeResponseBody) SetVendorType(v string) *GenerateAuthCodeResponseBody {
	s.VendorType = &v
	return s
}

func (s *GenerateAuthCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
