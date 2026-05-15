// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityRuleListResponseBody
	GetCode() *string
	SetData(v *GetQualityRuleListResponseBodyData) *GetQualityRuleListResponseBody
	GetData() *GetQualityRuleListResponseBodyData
	SetMessage(v string) *GetQualityRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityRuleListResponseBody
	GetSuccess() *bool
}

type GetQualityRuleListResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQualityRuleListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityRuleListResponseBody) GetData() *GetQualityRuleListResponseBodyData {
	return s.Data
}

func (s *GetQualityRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityRuleListResponseBody) SetCode(v string) *GetQualityRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleListResponseBody) SetData(v *GetQualityRuleListResponseBodyData) *GetQualityRuleListResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityRuleListResponseBody) SetMessage(v string) *GetQualityRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleListResponseBody) SetRequestId(v string) *GetQualityRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleListResponseBody) SetSuccess(v bool) *GetQualityRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleListResponseBodyData struct {
	PageNo          *int32                                               `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QualityRuleList []*GetQualityRuleListResponseBodyDataQualityRuleList `json:"QualityRuleList,omitempty" xml:"QualityRuleList,omitempty" type:"Repeated"`
	Total           *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQualityRuleListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQualityRuleListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQualityRuleListResponseBodyData) GetQualityRuleList() []*GetQualityRuleListResponseBodyDataQualityRuleList {
	return s.QualityRuleList
}

func (s *GetQualityRuleListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQualityRuleListResponseBodyData) SetPageNo(v int32) *GetQualityRuleListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQualityRuleListResponseBodyData) SetPageSize(v int32) *GetQualityRuleListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQualityRuleListResponseBodyData) SetQualityRuleList(v []*GetQualityRuleListResponseBodyDataQualityRuleList) *GetQualityRuleListResponseBodyData {
	s.QualityRuleList = v
	return s
}

func (s *GetQualityRuleListResponseBodyData) SetTotal(v int64) *GetQualityRuleListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQualityRuleListResponseBodyData) Validate() error {
	if s.QualityRuleList != nil {
		for _, item := range s.QualityRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQualityRuleListResponseBodyDataQualityRuleList struct {
	KeyWords       []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	MatchType      *int32    `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleCreateTime *string   `json:"RuleCreateTime,omitempty" xml:"RuleCreateTime,omitempty"`
	RuleId         *int64    `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleTag        *int32    `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
}

func (s GetQualityRuleListResponseBodyDataQualityRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleListResponseBodyDataQualityRuleList) GoString() string {
	return s.String()
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) GetMatchType() *int32 {
	return s.MatchType
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) GetName() *string {
	return s.Name
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) GetRuleCreateTime() *string {
	return s.RuleCreateTime
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) GetRuleTag() *int32 {
	return s.RuleTag
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetKeyWords(v []*string) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.KeyWords = v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetMatchType(v int32) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.MatchType = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetName(v string) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.Name = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetRuleCreateTime(v string) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.RuleCreateTime = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetRuleId(v int64) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.RuleId = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) SetRuleTag(v int32) *GetQualityRuleListResponseBodyDataQualityRuleList {
	s.RuleTag = &v
	return s
}

func (s *GetQualityRuleListResponseBodyDataQualityRuleList) Validate() error {
	return dara.Validate(s)
}
