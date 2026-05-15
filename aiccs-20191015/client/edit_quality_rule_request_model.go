// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EditQualityRuleRequest
  GetInstanceId() *string 
  SetKeyWords(v []*string) *EditQualityRuleRequest
  GetKeyWords() []*string 
  SetMatchType(v int32) *EditQualityRuleRequest
  GetMatchType() *int32 
  SetName(v string) *EditQualityRuleRequest
  GetName() *string 
  SetQualityRuleId(v int64) *EditQualityRuleRequest
  GetQualityRuleId() *int64 
  SetRuleTag(v int32) *EditQualityRuleRequest
  GetRuleTag() *int32 
}

type EditQualityRuleRequest struct {
  // This parameter is required.
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // This parameter is required.
  KeyWords []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
  // This parameter is required.
  MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
  // This parameter is required.
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // This parameter is required.
  QualityRuleId *int64 `json:"QualityRuleId,omitempty" xml:"QualityRuleId,omitempty"`
  // This parameter is required.
  RuleTag *int32 `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
}

func (s EditQualityRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EditQualityRuleRequest) GoString() string {
  return s.String()
}

func (s *EditQualityRuleRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EditQualityRuleRequest) GetKeyWords() []*string  {
  return s.KeyWords
}

func (s *EditQualityRuleRequest) GetMatchType() *int32  {
  return s.MatchType
}

func (s *EditQualityRuleRequest) GetName() *string  {
  return s.Name
}

func (s *EditQualityRuleRequest) GetQualityRuleId() *int64  {
  return s.QualityRuleId
}

func (s *EditQualityRuleRequest) GetRuleTag() *int32  {
  return s.RuleTag
}

func (s *EditQualityRuleRequest) SetInstanceId(v string) *EditQualityRuleRequest {
  s.InstanceId = &v
  return s
}

func (s *EditQualityRuleRequest) SetKeyWords(v []*string) *EditQualityRuleRequest {
  s.KeyWords = v
  return s
}

func (s *EditQualityRuleRequest) SetMatchType(v int32) *EditQualityRuleRequest {
  s.MatchType = &v
  return s
}

func (s *EditQualityRuleRequest) SetName(v string) *EditQualityRuleRequest {
  s.Name = &v
  return s
}

func (s *EditQualityRuleRequest) SetQualityRuleId(v int64) *EditQualityRuleRequest {
  s.QualityRuleId = &v
  return s
}

func (s *EditQualityRuleRequest) SetRuleTag(v int32) *EditQualityRuleRequest {
  s.RuleTag = &v
  return s
}

func (s *EditQualityRuleRequest) Validate() error {
  return dara.Validate(s)
}

