// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPhase(v string) *GetWafFilterRequest
	GetPhase() *string
	SetSiteId(v int64) *GetWafFilterRequest
	GetSiteId() *int64
	SetTarget(v string) *GetWafFilterRequest
	GetTarget() *string
	SetType(v string) *GetWafFilterRequest
	GetType() *string
}

type GetWafFilterRequest struct {
	// The WAF phase. Specifies the WAF phase for which to query the match engine information.
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The target. Defines the application target of the match engine.
	//
	// example:
	//
	// characteristics
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The rule type.
	//
	// example:
	//
	// http_custom_cc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWafFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterRequest) GoString() string {
	return s.String()
}

func (s *GetWafFilterRequest) GetPhase() *string {
	return s.Phase
}

func (s *GetWafFilterRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetWafFilterRequest) GetTarget() *string {
	return s.Target
}

func (s *GetWafFilterRequest) GetType() *string {
	return s.Type
}

func (s *GetWafFilterRequest) SetPhase(v string) *GetWafFilterRequest {
	s.Phase = &v
	return s
}

func (s *GetWafFilterRequest) SetSiteId(v int64) *GetWafFilterRequest {
	s.SiteId = &v
	return s
}

func (s *GetWafFilterRequest) SetTarget(v string) *GetWafFilterRequest {
	s.Target = &v
	return s
}

func (s *GetWafFilterRequest) SetType(v string) *GetWafFilterRequest {
	s.Type = &v
	return s
}

func (s *GetWafFilterRequest) Validate() error {
	return dara.Validate(s)
}
