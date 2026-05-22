// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWafRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteWafRuleRequest
	GetId() *int64
	SetSiteId(v int64) *DeleteWafRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *DeleteWafRuleRequest
	GetSiteVersion() *int32
}

type DeleteWafRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s DeleteWafRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWafRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWafRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteWafRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteWafRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *DeleteWafRuleRequest) SetId(v int64) *DeleteWafRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteWafRuleRequest) SetSiteId(v int64) *DeleteWafRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWafRuleRequest) SetSiteVersion(v int32) *DeleteWafRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *DeleteWafRuleRequest) Validate() error {
	return dara.Validate(s)
}
