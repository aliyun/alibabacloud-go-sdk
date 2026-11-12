// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErrorPagesRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteErrorPagesRedirectRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteErrorPagesRedirectRuleRequest
	GetSiteId() *int64
}

type DeleteErrorPagesRedirectRuleRequest struct {
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to query the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteErrorPagesRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteErrorPagesRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteErrorPagesRedirectRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteErrorPagesRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteErrorPagesRedirectRuleRequest) SetConfigId(v int64) *DeleteErrorPagesRedirectRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteErrorPagesRedirectRuleRequest) SetSiteId(v int64) *DeleteErrorPagesRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteErrorPagesRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
