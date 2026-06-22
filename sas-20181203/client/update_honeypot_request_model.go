// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotId(v string) *UpdateHoneypotRequest
	GetHoneypotId() *string
	SetHoneypotName(v string) *UpdateHoneypotRequest
	GetHoneypotName() *string
	SetLang(v string) *UpdateHoneypotRequest
	GetLang() *string
	SetMeta(v string) *UpdateHoneypotRequest
	GetMeta() *string
}

type UpdateHoneypotRequest struct {
	// The honeypot ID.
	//
	// > You can call the [ListHoneypot](~~ListHoneypot~~) operation to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 945607c2ae2a1a737c04599d6608065688bfc6048d9b9d306ce8dc8191c*****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The custom name of the honeypot.
	//
	// example:
	//
	// ExtMail
	HoneypotName *string `json:"HoneypotName,omitempty" xml:"HoneypotName,omitempty"`
	// The language of the content in the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The custom configuration of the honeypot.
	//
	// >The value of this parameter is obtained from the **Template*	- field returned by the [ListAvailableHoneypot](~~ListAvailableHoneypot~~) operation.
	//
	// example:
	//
	// {\\"burp\\":\\"open\\",\\"webshell\\":\\"open\\",\\"trojan_git\\":\\"close\\",\\"portrait_option\\":\\"true\\"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
}

func (s UpdateHoneypotRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotRequest) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotRequest) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *UpdateHoneypotRequest) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *UpdateHoneypotRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateHoneypotRequest) GetMeta() *string {
	return s.Meta
}

func (s *UpdateHoneypotRequest) SetHoneypotId(v string) *UpdateHoneypotRequest {
	s.HoneypotId = &v
	return s
}

func (s *UpdateHoneypotRequest) SetHoneypotName(v string) *UpdateHoneypotRequest {
	s.HoneypotName = &v
	return s
}

func (s *UpdateHoneypotRequest) SetLang(v string) *UpdateHoneypotRequest {
	s.Lang = &v
	return s
}

func (s *UpdateHoneypotRequest) SetMeta(v string) *UpdateHoneypotRequest {
	s.Meta = &v
	return s
}

func (s *UpdateHoneypotRequest) Validate() error {
	return dara.Validate(s)
}
