// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotPresetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotPresetId(v string) *GetHoneypotPresetRequest
	GetHoneypotPresetId() *string
	SetLang(v string) *GetHoneypotPresetRequest
	GetLang() *string
}

type GetHoneypotPresetRequest struct {
	// The ID of the honeypot template.
	//
	// > You can call the [ListHoneypotPreset](~~ListHoneypotPreset~~) operation to query the IDs of honeypot templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 207ca117-44b9-495d-84e7-50289b4cxxxx
	HoneypotPresetId *string `json:"HoneypotPresetId,omitempty" xml:"HoneypotPresetId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetHoneypotPresetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotPresetRequest) GoString() string {
	return s.String()
}

func (s *GetHoneypotPresetRequest) GetHoneypotPresetId() *string {
	return s.HoneypotPresetId
}

func (s *GetHoneypotPresetRequest) GetLang() *string {
	return s.Lang
}

func (s *GetHoneypotPresetRequest) SetHoneypotPresetId(v string) *GetHoneypotPresetRequest {
	s.HoneypotPresetId = &v
	return s
}

func (s *GetHoneypotPresetRequest) SetLang(v string) *GetHoneypotPresetRequest {
	s.Lang = &v
	return s
}

func (s *GetHoneypotPresetRequest) Validate() error {
	return dara.Validate(s)
}
