// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRewriteUrlRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteRewriteUrlRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteRewriteUrlRuleRequest
	GetSiteId() *int64
}

type DeleteRewriteUrlRuleRequest struct {
	// The configuration ID, which can be obtained by calling the [ListRewriteUrlRules](~~ListRewriteUrlRules~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteRewriteUrlRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRewriteUrlRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRewriteUrlRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteRewriteUrlRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteRewriteUrlRuleRequest) SetConfigId(v int64) *DeleteRewriteUrlRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteRewriteUrlRuleRequest) SetSiteId(v int64) *DeleteRewriteUrlRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteRewriteUrlRuleRequest) Validate() error {
	return dara.Validate(s)
}
