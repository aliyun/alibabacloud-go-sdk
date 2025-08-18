// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateWafRulesetRequest
	GetId() *int64
	SetSiteId(v int64) *UpdateWafRulesetRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *UpdateWafRulesetRequest
	GetSiteVersion() *int32
	SetStatus(v string) *UpdateWafRulesetRequest
	GetStatus() *string
}

type UpdateWafRulesetRequest struct {
	// ID of the WAF ruleset, which can be obtained by calling the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
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
	// The target status to change for the ruleset.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *UpdateWafRulesetRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateWafRulesetRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateWafRulesetRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *UpdateWafRulesetRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateWafRulesetRequest) SetId(v int64) *UpdateWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *UpdateWafRulesetRequest) SetSiteId(v int64) *UpdateWafRulesetRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWafRulesetRequest) SetSiteVersion(v int32) *UpdateWafRulesetRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateWafRulesetRequest) SetStatus(v string) *UpdateWafRulesetRequest {
	s.Status = &v
	return s
}

func (s *UpdateWafRulesetRequest) Validate() error {
	return dara.Validate(s)
}
