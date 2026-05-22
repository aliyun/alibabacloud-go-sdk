// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWafRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigShrink(v string) *UpdateWafRuleShrinkRequest
	GetConfigShrink() *string
	SetId(v int64) *UpdateWafRuleShrinkRequest
	GetId() *int64
	SetPosition(v int64) *UpdateWafRuleShrinkRequest
	GetPosition() *int64
	SetSiteId(v int64) *UpdateWafRuleShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *UpdateWafRuleShrinkRequest
	GetSiteVersion() *int32
	SetStatus(v string) *UpdateWafRuleShrinkRequest
	GetStatus() *string
}

type UpdateWafRuleShrinkRequest struct {
	// Rule configuration.
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// WAF rule ID, which can be obtained by calling the [ListWafRules](https://help.aliyun.com/document_detail/2878257.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The position of the rule in the rule set.
	//
	// example:
	//
	// 1
	Position *int64 `json:"Position,omitempty" xml:"Position,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site version.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Rule status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateWafRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWafRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWafRuleShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *UpdateWafRuleShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateWafRuleShrinkRequest) GetPosition() *int64 {
	return s.Position
}

func (s *UpdateWafRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateWafRuleShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *UpdateWafRuleShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateWafRuleShrinkRequest) SetConfigShrink(v string) *UpdateWafRuleShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetId(v int64) *UpdateWafRuleShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetPosition(v int64) *UpdateWafRuleShrinkRequest {
	s.Position = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetSiteId(v int64) *UpdateWafRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetSiteVersion(v int32) *UpdateWafRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) SetStatus(v string) *UpdateWafRuleShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateWafRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
