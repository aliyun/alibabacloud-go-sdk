// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleName(v string) *ListWaitingRoomRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListWaitingRoomRulesRequest
	GetSiteId() *int64
	SetWaitingRoomId(v string) *ListWaitingRoomRulesRequest
	GetWaitingRoomId() *string
	SetWaitingRoomRuleId(v int64) *ListWaitingRoomRulesRequest
	GetWaitingRoomRuleId() *int64
}

type ListWaitingRoomRulesRequest struct {
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	WaitingRoomId     *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
	WaitingRoomRuleId *int64  `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s ListWaitingRoomRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListWaitingRoomRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWaitingRoomRulesRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *ListWaitingRoomRulesRequest) GetWaitingRoomRuleId() *int64 {
	return s.WaitingRoomRuleId
}

func (s *ListWaitingRoomRulesRequest) SetRuleName(v string) *ListWaitingRoomRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListWaitingRoomRulesRequest) SetSiteId(v int64) *ListWaitingRoomRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWaitingRoomRulesRequest) SetWaitingRoomId(v string) *ListWaitingRoomRulesRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *ListWaitingRoomRulesRequest) SetWaitingRoomRuleId(v int64) *ListWaitingRoomRulesRequest {
	s.WaitingRoomRuleId = &v
	return s
}

func (s *ListWaitingRoomRulesRequest) Validate() error {
	return dara.Validate(s)
}
