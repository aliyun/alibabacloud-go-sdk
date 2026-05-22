// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompressionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetCompressionRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetCompressionRuleRequest
	GetSiteId() *int64
}

type GetCompressionRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetCompressionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCompressionRuleRequest) GoString() string {
	return s.String()
}

func (s *GetCompressionRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCompressionRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCompressionRuleRequest) SetConfigId(v int64) *GetCompressionRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetCompressionRuleRequest) SetSiteId(v int64) *GetCompressionRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetCompressionRuleRequest) Validate() error {
	return dara.Validate(s)
}
