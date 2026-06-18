// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpRequestHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetHttpRequestHeaderModificationRuleRequest
	GetSiteId() *int64
}

type GetHttpRequestHeaderModificationRuleRequest struct {
	// The configuration ID. You can obtain this ID by calling the [ListHttpRequestHeaderModificationRules](https://help.aliyun.com/document_detail/2867483.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the site. You can obtain this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetHttpRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *GetHttpRequestHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetHttpRequestHeaderModificationRuleRequest) SetConfigId(v int64) *GetHttpRequestHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleRequest) SetSiteId(v int64) *GetHttpRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
