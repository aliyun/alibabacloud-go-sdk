// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAutomaticFrequencyControlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *SetAutomaticFrequencyControlConfigRequest
	GetActionType() *string
	SetEnable(v string) *SetAutomaticFrequencyControlConfigRequest
	GetEnable() *string
	SetLevel(v string) *SetAutomaticFrequencyControlConfigRequest
	GetLevel() *string
	SetSiteId(v int64) *SetAutomaticFrequencyControlConfigRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *SetAutomaticFrequencyControlConfigRequest
	GetSiteVersion() *int32
}

type SetAutomaticFrequencyControlConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// js
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// normal
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s SetAutomaticFrequencyControlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAutomaticFrequencyControlConfigRequest) GoString() string {
	return s.String()
}

func (s *SetAutomaticFrequencyControlConfigRequest) GetActionType() *string {
	return s.ActionType
}

func (s *SetAutomaticFrequencyControlConfigRequest) GetEnable() *string {
	return s.Enable
}

func (s *SetAutomaticFrequencyControlConfigRequest) GetLevel() *string {
	return s.Level
}

func (s *SetAutomaticFrequencyControlConfigRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetAutomaticFrequencyControlConfigRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *SetAutomaticFrequencyControlConfigRequest) SetActionType(v string) *SetAutomaticFrequencyControlConfigRequest {
	s.ActionType = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigRequest) SetEnable(v string) *SetAutomaticFrequencyControlConfigRequest {
	s.Enable = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigRequest) SetLevel(v string) *SetAutomaticFrequencyControlConfigRequest {
	s.Level = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigRequest) SetSiteId(v int64) *SetAutomaticFrequencyControlConfigRequest {
	s.SiteId = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigRequest) SetSiteVersion(v int32) *SetAutomaticFrequencyControlConfigRequest {
	s.SiteVersion = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigRequest) Validate() error {
	return dara.Validate(s)
}
