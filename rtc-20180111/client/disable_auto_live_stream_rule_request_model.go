// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoLiveStreamRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DisableAutoLiveStreamRuleRequest
	GetAppId() *string
	SetOwnerId(v int64) *DisableAutoLiveStreamRuleRequest
	GetOwnerId() *int64
	SetRuleId(v int64) *DisableAutoLiveStreamRuleRequest
	GetRuleId() *int64
}

type DisableAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableAutoLiveStreamRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableAutoLiveStreamRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *DisableAutoLiveStreamRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableAutoLiveStreamRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DisableAutoLiveStreamRuleRequest) SetAppId(v string) *DisableAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *DisableAutoLiveStreamRuleRequest) SetOwnerId(v int64) *DisableAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableAutoLiveStreamRuleRequest) SetRuleId(v int64) *DisableAutoLiveStreamRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DisableAutoLiveStreamRuleRequest) Validate() error {
	return dara.Validate(s)
}
