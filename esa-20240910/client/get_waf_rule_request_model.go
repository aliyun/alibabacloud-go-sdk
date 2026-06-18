// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetWafRuleRequest
	GetId() *int64
	SetSiteId(v int64) *GetWafRuleRequest
	GetSiteId() *int64
}

type GetWafRuleRequest struct {
	// The ID of the WAF rule. You can get this ID by calling the [ListWafRules](https://help.aliyun.com/document_detail/2878257.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the site. You can get this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetWafRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWafRuleRequest) GoString() string {
	return s.String()
}

func (s *GetWafRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *GetWafRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetWafRuleRequest) SetId(v int64) *GetWafRuleRequest {
	s.Id = &v
	return s
}

func (s *GetWafRuleRequest) SetSiteId(v int64) *GetWafRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetWafRuleRequest) Validate() error {
	return dara.Validate(s)
}
