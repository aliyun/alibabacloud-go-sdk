// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMobilesCardSupportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptType(v string) *QueryMobilesCardSupportRequest
	GetEncryptType() *string
	SetMobiles(v []map[string]interface{}) *QueryMobilesCardSupportRequest
	GetMobiles() []map[string]interface{}
	SetTemplateCode(v string) *QueryMobilesCardSupportRequest
	GetTemplateCode() *string
}

type QueryMobilesCardSupportRequest struct {
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
	Mobiles []map[string]interface{} `json:"Mobiles,omitempty" xml:"Mobiles,omitempty" type:"Repeated"`
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

func (s QueryMobilesCardSupportRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportRequest) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *QueryMobilesCardSupportRequest) GetMobiles() []map[string]interface{} {
	return s.Mobiles
}

func (s *QueryMobilesCardSupportRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QueryMobilesCardSupportRequest) SetEncryptType(v string) *QueryMobilesCardSupportRequest {
	s.EncryptType = &v
	return s
}

func (s *QueryMobilesCardSupportRequest) SetMobiles(v []map[string]interface{}) *QueryMobilesCardSupportRequest {
	s.Mobiles = v
	return s
}

func (s *QueryMobilesCardSupportRequest) SetTemplateCode(v string) *QueryMobilesCardSupportRequest {
	s.TemplateCode = &v
	return s
}

func (s *QueryMobilesCardSupportRequest) Validate() error {
	return dara.Validate(s)
}
