// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomResponseCodeRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetCustomResponseCodeRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetCustomResponseCodeRuleRequest
	GetSiteId() *int64
}

type GetCustomResponseCodeRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 424022244554752
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 775724064754208
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetCustomResponseCodeRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomResponseCodeRuleRequest) GoString() string {
	return s.String()
}

func (s *GetCustomResponseCodeRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCustomResponseCodeRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCustomResponseCodeRuleRequest) SetConfigId(v int64) *GetCustomResponseCodeRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetCustomResponseCodeRuleRequest) SetSiteId(v int64) *GetCustomResponseCodeRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetCustomResponseCodeRuleRequest) Validate() error {
	return dara.Validate(s)
}
