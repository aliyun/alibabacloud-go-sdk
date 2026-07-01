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
	// The code of the card SMS template. Valid values:
	//
	// - After you call the [CreateCardSmsTemplate](~~CreateCardSmsTemplate~~) operation, the value of the **TemplateCode*	- response parameter is the code of the newly created card SMS template.
	//
	// - You can also log on to the [Domestic Card SMS](https://dysms.console.aliyun.com/domestic/card) page in the console and view the card SMS template code on the **Template Management*	- tab.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_2****
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
