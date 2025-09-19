// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevels(v string) *CreateUserSettingRequest
	GetAlertLevels() *string
	SetInvalidWarningKeepDays(v int32) *CreateUserSettingRequest
	GetInvalidWarningKeepDays() *int32
	SetSourceIp(v string) *CreateUserSettingRequest
	GetSourceIp() *string
}

type CreateUserSettingRequest struct {
	// The severities of alerts.
	//
	// example:
	//
	// high,low
	AlertLevels *string `json:"AlertLevels,omitempty" xml:"AlertLevels,omitempty"`
	// The number of days during which you want to retain invalid alerts.
	//
	// example:
	//
	// 7
	InvalidWarningKeepDays *int32 `json:"InvalidWarningKeepDays,omitempty" xml:"InvalidWarningKeepDays,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 112.48.16.***
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSettingRequest) GoString() string {
	return s.String()
}

func (s *CreateUserSettingRequest) GetAlertLevels() *string {
	return s.AlertLevels
}

func (s *CreateUserSettingRequest) GetInvalidWarningKeepDays() *int32 {
	return s.InvalidWarningKeepDays
}

func (s *CreateUserSettingRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateUserSettingRequest) SetAlertLevels(v string) *CreateUserSettingRequest {
	s.AlertLevels = &v
	return s
}

func (s *CreateUserSettingRequest) SetInvalidWarningKeepDays(v int32) *CreateUserSettingRequest {
	s.InvalidWarningKeepDays = &v
	return s
}

func (s *CreateUserSettingRequest) SetSourceIp(v string) *CreateUserSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
