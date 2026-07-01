// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMobilesCardSupportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptType(v string) *QueryMobilesCardSupportShrinkRequest
	GetEncryptType() *string
	SetMobilesShrink(v string) *QueryMobilesCardSupportShrinkRequest
	GetMobilesShrink() *string
	SetTemplateCode(v string) *QueryMobilesCardSupportShrinkRequest
	GetTemplateCode() *string
}

type QueryMobilesCardSupportShrinkRequest struct {
	// The encryption method for the phone number. Valid values:
	//
	// - SHA1: SHA1 encryption.
	//
	// - NORMAL: no encryption. The phone number is transmitted in plaintext.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// NORMAL
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The list of phone numbers.
	//
	// This parameter is required.
	MobilesShrink *string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty"`
	// The code of the card SMS template. To view the code, log on to the console and choose **Domestic Messages*	- > [Template Management](https://dysms.console.aliyun.com/domestic/text/template).
	//
	// >The template must be added and approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_2****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QueryMobilesCardSupportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportShrinkRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *QueryMobilesCardSupportShrinkRequest) GetMobilesShrink() *string {
	return s.MobilesShrink
}

func (s *QueryMobilesCardSupportShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QueryMobilesCardSupportShrinkRequest) SetEncryptType(v string) *QueryMobilesCardSupportShrinkRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryMobilesCardSupportShrinkRequest) SetMobilesShrink(v string) *QueryMobilesCardSupportShrinkRequest {
	s.MobilesShrink = &v
	return s
}

func (s *QueryMobilesCardSupportShrinkRequest) SetTemplateCode(v string) *QueryMobilesCardSupportShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *QueryMobilesCardSupportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
