// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWafRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *WafRuleConfig) *UpdateWafRuleRequest
	GetConfig() *WafRuleConfig
	SetId(v int64) *UpdateWafRuleRequest
	GetId() *int64
	SetPosition(v int64) *UpdateWafRuleRequest
	GetPosition() *int64
	SetSiteId(v int64) *UpdateWafRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *UpdateWafRuleRequest
	GetSiteVersion() *int32
	SetStatus(v string) *UpdateWafRuleRequest
	GetStatus() *string
}

type UpdateWafRuleRequest struct {
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateWafRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWafRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWafRuleRequest) GetConfig() *WafRuleConfig {
	return s.Config
}

func (s *UpdateWafRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateWafRuleRequest) GetPosition() *int64 {
	return s.Position
}

func (s *UpdateWafRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateWafRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *UpdateWafRuleRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateWafRuleRequest) SetConfig(v *WafRuleConfig) *UpdateWafRuleRequest {
	s.Config = v
	return s
}

func (s *UpdateWafRuleRequest) SetId(v int64) *UpdateWafRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateWafRuleRequest) SetPosition(v int64) *UpdateWafRuleRequest {
	s.Position = &v
	return s
}

func (s *UpdateWafRuleRequest) SetSiteId(v int64) *UpdateWafRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWafRuleRequest) SetSiteVersion(v int32) *UpdateWafRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateWafRuleRequest) SetStatus(v string) *UpdateWafRuleRequest {
	s.Status = &v
	return s
}

func (s *UpdateWafRuleRequest) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}
