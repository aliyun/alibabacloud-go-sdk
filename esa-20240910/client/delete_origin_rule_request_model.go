// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteOriginRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteOriginRuleRequest
	GetSiteId() *int64
}

type DeleteOriginRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteOriginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteOriginRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteOriginRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteOriginRuleRequest) SetConfigId(v int64) *DeleteOriginRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteOriginRuleRequest) SetSiteId(v int64) *DeleteOriginRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteOriginRuleRequest) Validate() error {
	return dara.Validate(s)
}
