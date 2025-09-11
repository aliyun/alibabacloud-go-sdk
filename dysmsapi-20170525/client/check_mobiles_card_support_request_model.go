// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMobilesCardSupportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobiles(v []map[string]interface{}) *CheckMobilesCardSupportRequest
	GetMobiles() []map[string]interface{}
	SetTemplateCode(v string) *CheckMobilesCardSupportRequest
	GetTemplateCode() *string
}

type CheckMobilesCardSupportRequest struct {
	// The list of mobile phone numbers that receive messages.
	//
	// This parameter is required.
	Mobiles []map[string]interface{} `json:"Mobiles,omitempty" xml:"Mobiles,omitempty" type:"Repeated"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CheckMobilesCardSupportRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckMobilesCardSupportRequest) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportRequest) GetMobiles() []map[string]interface{} {
	return s.Mobiles
}

func (s *CheckMobilesCardSupportRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CheckMobilesCardSupportRequest) SetMobiles(v []map[string]interface{}) *CheckMobilesCardSupportRequest {
	s.Mobiles = v
	return s
}

func (s *CheckMobilesCardSupportRequest) SetTemplateCode(v string) *CheckMobilesCardSupportRequest {
	s.TemplateCode = &v
	return s
}

func (s *CheckMobilesCardSupportRequest) Validate() error {
	return dara.Validate(s)
}
