// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityRuleDetailResponseBody
	GetCode() *string
	SetData(v *GetQualityRuleDetailResponseBodyData) *GetQualityRuleDetailResponseBody
	GetData() *GetQualityRuleDetailResponseBodyData
	SetMessage(v string) *GetQualityRuleDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityRuleDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityRuleDetailResponseBody
	GetSuccess() *bool
}

type GetQualityRuleDetailResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQualityRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityRuleDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityRuleDetailResponseBody) GetData() *GetQualityRuleDetailResponseBodyData {
	return s.Data
}

func (s *GetQualityRuleDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityRuleDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityRuleDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityRuleDetailResponseBody) SetCode(v string) *GetQualityRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleDetailResponseBody) SetData(v *GetQualityRuleDetailResponseBodyData) *GetQualityRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityRuleDetailResponseBody) SetMessage(v string) *GetQualityRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleDetailResponseBody) SetRequestId(v string) *GetQualityRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleDetailResponseBody) SetSuccess(v bool) *GetQualityRuleDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleDetailResponseBodyData struct {
	KeyWords       []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	MatchType      *int32    `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleCreateTime *string   `json:"RuleCreateTime,omitempty" xml:"RuleCreateTime,omitempty"`
	RuleId         *int64    `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleTag        *int32    `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
}

func (s GetQualityRuleDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityRuleDetailResponseBodyData) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *GetQualityRuleDetailResponseBodyData) GetMatchType() *int32 {
	return s.MatchType
}

func (s *GetQualityRuleDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetQualityRuleDetailResponseBodyData) GetRuleCreateTime() *string {
	return s.RuleCreateTime
}

func (s *GetQualityRuleDetailResponseBodyData) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetQualityRuleDetailResponseBodyData) GetRuleTag() *int32 {
	return s.RuleTag
}

func (s *GetQualityRuleDetailResponseBodyData) SetKeyWords(v []*string) *GetQualityRuleDetailResponseBodyData {
	s.KeyWords = v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetMatchType(v int32) *GetQualityRuleDetailResponseBodyData {
	s.MatchType = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetName(v string) *GetQualityRuleDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetRuleCreateTime(v string) *GetQualityRuleDetailResponseBodyData {
	s.RuleCreateTime = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetRuleId(v int64) *GetQualityRuleDetailResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) SetRuleTag(v int32) *GetQualityRuleDetailResponseBodyData {
	s.RuleTag = &v
	return s
}

func (s *GetQualityRuleDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
