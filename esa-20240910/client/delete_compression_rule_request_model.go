// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompressionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteCompressionRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteCompressionRuleRequest
	GetSiteId() *int64
}

type DeleteCompressionRuleRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteCompressionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompressionRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCompressionRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteCompressionRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteCompressionRuleRequest) SetConfigId(v int64) *DeleteCompressionRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteCompressionRuleRequest) SetSiteId(v int64) *DeleteCompressionRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteCompressionRuleRequest) Validate() error {
	return dara.Validate(s)
}
