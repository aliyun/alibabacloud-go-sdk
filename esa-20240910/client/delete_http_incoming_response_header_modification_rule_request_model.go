// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpIncomingResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteHttpIncomingResponseHeaderModificationRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteHttpIncomingResponseHeaderModificationRuleRequest
	GetSiteId() *int64
}

type DeleteHttpIncomingResponseHeaderModificationRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 672344269424192
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteHttpIncomingResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpIncomingResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleRequest) SetConfigId(v int64) *DeleteHttpIncomingResponseHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleRequest) SetSiteId(v int64) *DeleteHttpIncomingResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteHttpIncomingResponseHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}
