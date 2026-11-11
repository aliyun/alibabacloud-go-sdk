// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpIncomingRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpIncomingRequestHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetHttpIncomingRequestHeaderModificationRuleRequest
	GetSiteId() *int64
}

type GetHttpIncomingRequestHeaderModificationRuleRequest struct {
	// The configuration ID. You can call the ListHttpIncomingRequestHeaderModificationRules operation to obtain the configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 433045006266368
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 608665779308176
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetHttpIncomingRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpIncomingRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *GetHttpIncomingRequestHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpIncomingRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetHttpIncomingRequestHeaderModificationRuleRequest) SetConfigId(v int64) *GetHttpIncomingRequestHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleRequest) SetSiteId(v int64) *GetHttpIncomingRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
