// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetOriginRuleRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetOriginRuleRequest
	GetSiteId() *int64
}

type GetOriginRuleRequest struct {
	// The ID of the configuration. You can call the [ListSiteRoutes](https://help.aliyun.com/document_detail/2866989.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33793140540****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3400350********
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginRuleRequest) GoString() string {
	return s.String()
}

func (s *GetOriginRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetOriginRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginRuleRequest) SetConfigId(v int64) *GetOriginRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetOriginRuleRequest) SetSiteId(v int64) *GetOriginRuleRequest {
	s.SiteId = &v
	return s
}

func (s *GetOriginRuleRequest) Validate() error {
	return dara.Validate(s)
}
