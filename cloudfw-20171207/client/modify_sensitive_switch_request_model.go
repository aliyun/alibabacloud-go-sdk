// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySensitiveSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifySensitiveSwitchRequest
	GetLang() *string
	SetSensitiveCategory(v string) *ModifySensitiveSwitchRequest
	GetSensitiveCategory() *string
	SetSwitchStatus(v string) *ModifySensitiveSwitchRequest
	GetSwitchStatus() *string
}

type ModifySensitiveSwitchRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// id_card
	SensitiveCategory *string `json:"SensitiveCategory,omitempty" xml:"SensitiveCategory,omitempty"`
	// example:
	//
	// 1
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s ModifySensitiveSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySensitiveSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifySensitiveSwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifySensitiveSwitchRequest) GetSensitiveCategory() *string {
	return s.SensitiveCategory
}

func (s *ModifySensitiveSwitchRequest) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *ModifySensitiveSwitchRequest) SetLang(v string) *ModifySensitiveSwitchRequest {
	s.Lang = &v
	return s
}

func (s *ModifySensitiveSwitchRequest) SetSensitiveCategory(v string) *ModifySensitiveSwitchRequest {
	s.SensitiveCategory = &v
	return s
}

func (s *ModifySensitiveSwitchRequest) SetSwitchStatus(v string) *ModifySensitiveSwitchRequest {
	s.SwitchStatus = &v
	return s
}

func (s *ModifySensitiveSwitchRequest) Validate() error {
	return dara.Validate(s)
}
