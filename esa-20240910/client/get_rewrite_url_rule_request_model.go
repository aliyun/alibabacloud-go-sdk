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
	// Configuration ID, which can be obtained by calling the [ListRewriteUrlRules](https://help.aliyun.com/document_detail/2867480.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
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
