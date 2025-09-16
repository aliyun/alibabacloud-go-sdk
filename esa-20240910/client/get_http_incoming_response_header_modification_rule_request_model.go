// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpIncomingResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpIncomingResponseHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetHttpIncomingResponseHeaderModificationRuleRequest
	GetSiteId() *int64
}

type GetHttpIncomingResponseHeaderModificationRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 430509230649344
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 347168101647504
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetHttpIncomingResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpIncomingResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *GetHttpIncomingResponseHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpIncomingResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetHttpIncomingResponseHeaderModificationRuleRequest) SetConfigId(v int64) *GetHttpIncomingResponseHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleRequest) SetSiteId(v int64) *GetHttpIncomingResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
