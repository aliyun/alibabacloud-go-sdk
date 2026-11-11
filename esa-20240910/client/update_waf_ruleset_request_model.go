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
	// The ID of the WAF rule set. You can call the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation to obtain the rule set ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the site. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The target status of the rule set.
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
