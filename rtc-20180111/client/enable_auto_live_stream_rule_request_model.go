// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoLiveStreamRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppId(v string) *EnableAutoLiveStreamRuleRequest
  GetAppId() *string 
  SetOwnerId(v int64) *EnableAutoLiveStreamRuleRequest
  GetOwnerId() *int64 
  SetRuleId(v int64) *EnableAutoLiveStreamRuleRequest
  GetRuleId() *int64 
}

type EnableAutoLiveStreamRuleRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // eo85****
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 12
  RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableAutoLiveStreamRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoLiveStreamRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableAutoLiveStreamRuleRequest) GetAppId() *string  {
  return s.AppId
}

func (s *EnableAutoLiveStreamRuleRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableAutoLiveStreamRuleRequest) GetRuleId() *int64  {
  return s.RuleId
}

func (s *EnableAutoLiveStreamRuleRequest) SetAppId(v string) *EnableAutoLiveStreamRuleRequest {
  s.AppId = &v
  return s
}

func (s *EnableAutoLiveStreamRuleRequest) SetOwnerId(v int64) *EnableAutoLiveStreamRuleRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableAutoLiveStreamRuleRequest) SetRuleId(v int64) *EnableAutoLiveStreamRuleRequest {
  s.RuleId = &v
  return s
}

func (s *EnableAutoLiveStreamRuleRequest) Validate() error {
  return dara.Validate(s)
}

