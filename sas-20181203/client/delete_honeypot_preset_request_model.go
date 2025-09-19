// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotPresetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotPresetId(v string) *DeleteHoneypotPresetRequest
	GetHoneypotPresetId() *string
	SetLang(v string) *DeleteHoneypotPresetRequest
	GetLang() *string
}

type DeleteHoneypotPresetRequest struct {
	// The ID of the honeypot template.
	//
	// > You can call the [ListHoneypotPreset](~~ListHoneypotPreset~~) operation to query the IDs of honeypot templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 84104b7b-a2bc-41df-a190-12298f99xxxx
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
}

func (s DeleteHoneypotPresetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotPresetRequest) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotPresetRequest) GetHoneypotPresetId() *string {
	return s.HoneypotPresetId
}

func (s *DeleteHoneypotPresetRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteHoneypotPresetRequest) SetHoneypotPresetId(v string) *DeleteHoneypotPresetRequest {
	s.HoneypotPresetId = &v
	return s
}

func (s *DeleteHoneypotPresetRequest) SetLang(v string) *DeleteHoneypotPresetRequest {
	s.Lang = &v
	return s
}

func (s *DeleteHoneypotPresetRequest) Validate() error {
	return dara.Validate(s)
}
