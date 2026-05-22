// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetRedirectRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetRedirectRuleRequest
	GetSiteId() *int64
}

type GetRedirectRuleRequest struct {
	// Configuration ID. It can be obtained by calling the [ListRedirectRules](~~ListRedirectRules~~) interface.
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
	// 34003500310****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRedirectRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetRedirectRuleRequest) SetConfigId(v int64) *GetRedirectRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetRedirectRuleRequest) SetSiteId(v int64) *GetRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
