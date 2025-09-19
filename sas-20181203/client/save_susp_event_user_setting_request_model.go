// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSuspEventUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *SaveSuspEventUserSettingRequest
	GetFrom() *string
	SetLevelsOn(v string) *SaveSuspEventUserSettingRequest
	GetLevelsOn() *string
}

type SaveSuspEventUserSettingRequest struct {
	// The data source of the exception. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The severities of alert notifications. Valid values:
	//
	// 	- **remind**
	//
	// 	- **suspicious**
	//
	// 	- **serious**
	//
	// example:
	//
	// suspicious,serious,remind
	LevelsOn *string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty"`
}

func (s SaveSuspEventUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSuspEventUserSettingRequest) GoString() string {
	return s.String()
}

func (s *SaveSuspEventUserSettingRequest) GetFrom() *string {
	return s.From
}

func (s *SaveSuspEventUserSettingRequest) GetLevelsOn() *string {
	return s.LevelsOn
}

func (s *SaveSuspEventUserSettingRequest) SetFrom(v string) *SaveSuspEventUserSettingRequest {
	s.From = &v
	return s
}

func (s *SaveSuspEventUserSettingRequest) SetLevelsOn(v string) *SaveSuspEventUserSettingRequest {
	s.LevelsOn = &v
	return s
}

func (s *SaveSuspEventUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
