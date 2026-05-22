// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRewriteUrlRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetRewriteUrlRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetRewriteUrlRuleRequest
	GetSiteId() *int64
}

type GetRewriteUrlRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetRewriteUrlRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRewriteUrlRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRewriteUrlRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetRewriteUrlRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetRewriteUrlRuleRequest) SetConfigId(v int64) *GetRewriteUrlRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetRewriteUrlRuleRequest) SetSiteId(v int64) *GetRewriteUrlRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetRewriteUrlRuleRequest) Validate() error {
	return dara.Validate(s)
}
