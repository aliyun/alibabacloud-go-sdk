// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteHttpRequestHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteHttpRequestHeaderModificationRuleRequest
	GetSiteId() *int64
}

type DeleteHttpRequestHeaderModificationRuleRequest struct {
	// The configuration ID, which can be obtained by calling the [ListHttpRequestHeaderModificationRules](~~ListHttpRequestHeaderModificationRules~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteHttpRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpRequestHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteHttpRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteHttpRequestHeaderModificationRuleRequest) SetConfigId(v int64) *DeleteHttpRequestHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteHttpRequestHeaderModificationRuleRequest) SetSiteId(v int64) *DeleteHttpRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteHttpRequestHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
