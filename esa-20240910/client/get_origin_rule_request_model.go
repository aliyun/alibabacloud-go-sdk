// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetOriginRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetOriginRuleRequest
	GetSiteId() *int64
}

type GetOriginRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginRuleRequest) GoString() string {
	return s.String()
}

func (s *GetOriginRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetOriginRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginRuleRequest) SetConfigId(v int64) *GetOriginRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetOriginRuleRequest) SetSiteId(v int64) *GetOriginRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetOriginRuleRequest) Validate() error {
	return dara.Validate(s)
}
