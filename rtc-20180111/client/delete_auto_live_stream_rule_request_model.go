// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoLiveStreamRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAutoLiveStreamRuleRequest
	GetAppId() *string
	SetOwnerId(v int64) *DeleteAutoLiveStreamRuleRequest
	GetOwnerId() *int64
	SetRuleId(v int64) *DeleteAutoLiveStreamRuleRequest
	GetRuleId() *int64
}

type DeleteAutoLiveStreamRuleRequest struct {
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

func (s DeleteAutoLiveStreamRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoLiveStreamRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAutoLiveStreamRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAutoLiveStreamRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteAutoLiveStreamRuleRequest) SetAppId(v string) *DeleteAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleRequest) SetOwnerId(v int64) *DeleteAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleRequest) SetRuleId(v int64) *DeleteAutoLiveStreamRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleRequest) Validate() error {
	return dara.Validate(s)
}
