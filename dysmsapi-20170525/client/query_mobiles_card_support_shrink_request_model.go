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
	// if can be null:
	// true
	//
	// example:
	//
	// NORMAL
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The list of mobile phone numbers.
	//
	// This parameter is required.
	MobilesShrink *string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_0000
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
