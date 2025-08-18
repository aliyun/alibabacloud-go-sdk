// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteRedirectRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteRedirectRuleRequest
	GetSiteId() *int64
}

type DeleteRedirectRuleRequest struct {
	// The configuration ID, which can be obtained by calling the [ListRedirectRules](~~ListRedirectRules~~) operation.
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

func (s DeleteRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRedirectRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteRedirectRuleRequest) SetConfigId(v int64) *DeleteRedirectRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteRedirectRuleRequest) SetSiteId(v int64) *DeleteRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
