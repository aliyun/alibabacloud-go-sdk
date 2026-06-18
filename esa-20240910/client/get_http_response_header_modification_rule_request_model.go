// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpResponseHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetHttpResponseHeaderModificationRuleRequest
	GetSiteId() *int64
}

type GetHttpResponseHeaderModificationRuleRequest struct {
	// The configuration ID. You can obtain the configuration ID by calling the [ListHttpResponseHeaderModificationRules](https://help.aliyun.com/document_detail/2867483.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The site ID. You can obtain the site ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetHttpResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *GetHttpResponseHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetHttpResponseHeaderModificationRuleRequest) SetConfigId(v int64) *GetHttpResponseHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleRequest) SetSiteId(v int64) *GetHttpResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
