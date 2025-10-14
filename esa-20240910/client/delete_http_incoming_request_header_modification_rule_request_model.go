// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpIncomingRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteHttpIncomingRequestHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteHttpIncomingRequestHeaderModificationRuleRequest
	GetSiteId() *int64
}

type DeleteHttpIncomingRequestHeaderModificationRuleRequest struct {
	// The configuration ID. You can call the ListHttpIncomingRequestHeaderModificationRules operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 427428371703808
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54362329990032
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteHttpIncomingRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpIncomingRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleRequest) SetConfigId(v int64) *DeleteHttpIncomingRequestHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleRequest) SetSiteId(v int64) *DeleteHttpIncomingRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteHttpIncomingRequestHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
