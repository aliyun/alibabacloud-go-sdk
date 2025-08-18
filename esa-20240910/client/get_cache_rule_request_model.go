// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetCacheRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetCacheRuleRequest
	GetSiteId() *int64
}

type GetCacheRuleRequest struct {
	// ConfigId of the configuration, which can be obtained by calling the [ListCacheRules](https://help.aliyun.com/document_detail/2866985.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
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

func (s GetCacheRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCacheRuleRequest) GoString() string {
	return s.String()
}

func (s *GetCacheRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCacheRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCacheRuleRequest) SetConfigId(v int64) *GetCacheRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetCacheRuleRequest) SetSiteId(v int64) *GetCacheRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetCacheRuleRequest) Validate() error {
	return dara.Validate(s)
}
