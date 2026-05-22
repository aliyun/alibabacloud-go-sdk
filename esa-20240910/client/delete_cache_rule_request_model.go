// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCacheRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteCacheRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteCacheRuleRequest
	GetSiteId() *int64
}

type DeleteCacheRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteCacheRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCacheRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCacheRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteCacheRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteCacheRuleRequest) SetConfigId(v int64) *DeleteCacheRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteCacheRuleRequest) SetSiteId(v int64) *DeleteCacheRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteCacheRuleRequest) Validate() error {
	return dara.Validate(s)
}
