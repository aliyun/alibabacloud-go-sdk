// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateWafRulesetRequest
	GetName() *string
	SetPhase(v string) *CreateWafRulesetRequest
	GetPhase() *string
	SetSiteId(v int64) *CreateWafRulesetRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateWafRulesetRequest
	GetSiteVersion() *int32
}

type CreateWafRulesetRequest struct {
	// Ruleset name.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// WAF operation phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
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
}

func (s CreateWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *CreateWafRulesetRequest) GetName() *string {
	return s.Name
}

func (s *CreateWafRulesetRequest) GetPhase() *string {
	return s.Phase
}

func (s *CreateWafRulesetRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateWafRulesetRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateWafRulesetRequest) SetName(v string) *CreateWafRulesetRequest {
	s.Name = &v
	return s
}

func (s *CreateWafRulesetRequest) SetPhase(v string) *CreateWafRulesetRequest {
	s.Phase = &v
	return s
}

func (s *CreateWafRulesetRequest) SetSiteId(v int64) *CreateWafRulesetRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWafRulesetRequest) SetSiteVersion(v int32) *CreateWafRulesetRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateWafRulesetRequest) Validate() error {
	return dara.Validate(s)
}
