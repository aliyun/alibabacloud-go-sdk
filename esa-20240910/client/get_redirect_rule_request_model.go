// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetRedirectRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetRedirectRuleRequest
	GetSiteId() *int64
}

type GetRedirectRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRedirectRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetRedirectRuleRequest) SetConfigId(v int64) *GetRedirectRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetRedirectRuleRequest) SetSiteId(v int64) *GetRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
