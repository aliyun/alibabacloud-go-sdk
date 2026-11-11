// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorPagesRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetErrorPagesRedirectRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetErrorPagesRedirectRuleRequest
	GetSiteId() *int64
}

type GetErrorPagesRedirectRuleRequest struct {
	// The ID of the configuration. You can call the [ListCompressionRules](https://help.aliyun.com/document_detail/2867498.html) operation to obtain the configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the site. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11223
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetErrorPagesRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErrorPagesRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *GetErrorPagesRedirectRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetErrorPagesRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetErrorPagesRedirectRuleRequest) SetConfigId(v int64) *GetErrorPagesRedirectRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetErrorPagesRedirectRuleRequest) SetSiteId(v int64) *GetErrorPagesRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetErrorPagesRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
