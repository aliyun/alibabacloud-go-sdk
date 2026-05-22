// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetCacheRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetCacheRuleRequest
	GetSiteId() *int64
}

type GetCacheRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetCacheRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCacheRuleRequest) GoString() string {
	return s.String()
}

func (s *GetCacheRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCacheRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCacheRuleRequest) SetConfigId(v int64) *GetCacheRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetCacheRuleRequest) SetSiteId(v int64) *GetCacheRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetCacheRuleRequest) Validate() error {
	return dara.Validate(s)
}
