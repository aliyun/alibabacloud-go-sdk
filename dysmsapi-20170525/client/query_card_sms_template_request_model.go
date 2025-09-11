// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCardSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateCode(v string) *QueryCardSmsTemplateRequest
	GetTemplateCode() *string
}

type QueryCardSmsTemplateRequest struct {
	// The code of the message template.
	//
	// You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_4139
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QueryCardSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCardSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QueryCardSmsTemplateRequest) SetTemplateCode(v string) *QueryCardSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *QueryCardSmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
