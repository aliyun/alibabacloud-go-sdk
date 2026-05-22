// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteWafRulesetRequest
	GetId() *int64
	SetSiteId(v int64) *DeleteWafRulesetRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *DeleteWafRulesetRequest
	GetSiteVersion() *int32
}

type DeleteWafRulesetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s DeleteWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *DeleteWafRulesetRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteWafRulesetRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteWafRulesetRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *DeleteWafRulesetRequest) SetId(v int64) *DeleteWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *DeleteWafRulesetRequest) SetSiteId(v int64) *DeleteWafRulesetRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWafRulesetRequest) SetSiteVersion(v int32) *DeleteWafRulesetRequest {
	s.SiteVersion = &v
	return s
}

func (s *DeleteWafRulesetRequest) Validate() error {
	return dara.Validate(s)
}
