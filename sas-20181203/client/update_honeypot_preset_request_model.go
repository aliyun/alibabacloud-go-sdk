// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotPresetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotImageName(v string) *UpdateHoneypotPresetRequest
	GetHoneypotImageName() *string
	SetHoneypotPresetId(v string) *UpdateHoneypotPresetRequest
	GetHoneypotPresetId() *string
	SetLang(v string) *UpdateHoneypotPresetRequest
	GetLang() *string
	SetMeta(v string) *UpdateHoneypotPresetRequest
	GetMeta() *string
	SetPresetName(v string) *UpdateHoneypotPresetRequest
	GetPresetName() *string
}

type UpdateHoneypotPresetRequest struct {
	// The name of the image that is used for the honeypot.
	//
	// example:
	//
	// metabase
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The ID of the honeypot template.
	//
	// > You can call the [ListHoneypotPreset](~~ListHoneypotPreset~~) operation to query the IDs of honeypot templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// f75eddce-e9d3-4a88-af95-b10b6f65xxxx
	HoneypotPresetId *string `json:"HoneypotPresetId,omitempty" xml:"HoneypotPresetId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The custom configurations of the honeypot template. The value is a JSON string that contains the following fields:
	//
	// 	- **portrait_option**: Social Source Tracing
	//
	// 	- **burp**: Burp-specific Defense
	//
	// 	- **trojan_git**: Git-specific Defense
	//
	// example:
	//
	// {"portrait_option":true,"burp":"open"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The custom name of the honeypot template.
	//
	// example:
	//
	// apc_web_python
	PresetName *string `json:"PresetName,omitempty" xml:"PresetName,omitempty"`
}

func (s UpdateHoneypotPresetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotPresetRequest) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotPresetRequest) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *UpdateHoneypotPresetRequest) GetHoneypotPresetId() *string {
	return s.HoneypotPresetId
}

func (s *UpdateHoneypotPresetRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateHoneypotPresetRequest) GetMeta() *string {
	return s.Meta
}

func (s *UpdateHoneypotPresetRequest) GetPresetName() *string {
	return s.PresetName
}

func (s *UpdateHoneypotPresetRequest) SetHoneypotImageName(v string) *UpdateHoneypotPresetRequest {
	s.HoneypotImageName = &v
	return s
}

func (s *UpdateHoneypotPresetRequest) SetHoneypotPresetId(v string) *UpdateHoneypotPresetRequest {
	s.HoneypotPresetId = &v
	return s
}

func (s *UpdateHoneypotPresetRequest) SetLang(v string) *UpdateHoneypotPresetRequest {
	s.Lang = &v
	return s
}

func (s *UpdateHoneypotPresetRequest) SetMeta(v string) *UpdateHoneypotPresetRequest {
	s.Meta = &v
	return s
}

func (s *UpdateHoneypotPresetRequest) SetPresetName(v string) *UpdateHoneypotPresetRequest {
	s.PresetName = &v
	return s
}

func (s *UpdateHoneypotPresetRequest) Validate() error {
	return dara.Validate(s)
}
