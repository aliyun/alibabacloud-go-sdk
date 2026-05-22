// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteHttpResponseHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteHttpResponseHeaderModificationRuleRequest
	GetSiteId() *int64
}

type DeleteHttpResponseHeaderModificationRuleRequest struct {
	// The configuration ID, which can be obtained by calling the [ListHttpResponseHeaderModificationRules](~~ListHttpResponseHeaderModificationRules~~) operation.
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
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteHttpResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpResponseHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteHttpResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteHttpResponseHeaderModificationRuleRequest) SetConfigId(v int64) *DeleteHttpResponseHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteHttpResponseHeaderModificationRuleRequest) SetSiteId(v int64) *DeleteHttpResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteHttpResponseHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
