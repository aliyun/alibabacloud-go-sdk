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
	// The ID of the configuration. You can call the [ListCompressionRules](https://help.aliyun.com/document_detail/2867498.html) operation to query the configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the site. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to query the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
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
