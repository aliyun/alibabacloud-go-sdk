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
	// ConfigId of the configuration, which can be obtained by calling the [ListOriginRules](https://help.aliyun.com/document_detail/2866989.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33793140540****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3400350********
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
