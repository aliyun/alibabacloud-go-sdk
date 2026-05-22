// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpRequestHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetHttpRequestHeaderModificationRuleRequest
	GetSiteId() *int64
}

type GetHttpRequestHeaderModificationRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetHttpRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *GetHttpRequestHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetHttpRequestHeaderModificationRuleRequest) SetConfigId(v int64) *GetHttpRequestHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleRequest) SetSiteId(v int64) *GetHttpRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
