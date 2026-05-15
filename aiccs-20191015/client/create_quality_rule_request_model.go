// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateQualityRuleRequest
	GetInstanceId() *string
	SetKeyWords(v []*string) *CreateQualityRuleRequest
	GetKeyWords() []*string
	SetMatchType(v int32) *CreateQualityRuleRequest
	GetMatchType() *int32
	SetName(v string) *CreateQualityRuleRequest
	GetName() *string
	SetRuleTag(v int32) *CreateQualityRuleRequest
	GetRuleTag() *int32
}

type CreateQualityRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	KeyWords []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// This parameter is required.
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RuleTag *int32 `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
}

func (s CreateQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateQualityRuleRequest) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *CreateQualityRuleRequest) GetMatchType() *int32 {
	return s.MatchType
}

func (s *CreateQualityRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateQualityRuleRequest) GetRuleTag() *int32 {
	return s.RuleTag
}

func (s *CreateQualityRuleRequest) SetInstanceId(v string) *CreateQualityRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQualityRuleRequest) SetKeyWords(v []*string) *CreateQualityRuleRequest {
	s.KeyWords = v
	return s
}

func (s *CreateQualityRuleRequest) SetMatchType(v int32) *CreateQualityRuleRequest {
	s.MatchType = &v
	return s
}

func (s *CreateQualityRuleRequest) SetName(v string) *CreateQualityRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateQualityRuleRequest) SetRuleTag(v int32) *CreateQualityRuleRequest {
	s.RuleTag = &v
	return s
}

func (s *CreateQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
