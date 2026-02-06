// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomResponseCodeRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteCustomResponseCodeRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteCustomResponseCodeRuleRequest
	GetSiteId() *int64
}

type DeleteCustomResponseCodeRuleRequest struct {
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 434497172875264
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 478016908379824
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteCustomResponseCodeRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomResponseCodeRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomResponseCodeRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteCustomResponseCodeRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteCustomResponseCodeRuleRequest) SetConfigId(v int64) *DeleteCustomResponseCodeRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteCustomResponseCodeRuleRequest) SetSiteId(v int64) *DeleteCustomResponseCodeRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteCustomResponseCodeRuleRequest) Validate() error {
	return dara.Validate(s)
}
