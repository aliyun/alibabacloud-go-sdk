// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompressionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetCompressionRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetCompressionRuleRequest
	GetSiteId() *int64
}

type GetCompressionRuleRequest struct {
	// Configuration ID, which can be obtained by calling the [ListCompressionRules](https://help.aliyun.com/document_detail/2867498.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34003500310****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetCompressionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCompressionRuleRequest) GoString() string {
	return s.String()
}

func (s *GetCompressionRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCompressionRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCompressionRuleRequest) SetConfigId(v int64) *GetCompressionRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetCompressionRuleRequest) SetSiteId(v int64) *GetCompressionRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetCompressionRuleRequest) Validate() error {
	return dara.Validate(s)
}
