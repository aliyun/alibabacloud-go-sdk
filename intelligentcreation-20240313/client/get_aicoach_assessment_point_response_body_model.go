// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachAssessmentPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnswerList(v []*GetAICoachAssessmentPointResponseBodyAnswerList) *GetAICoachAssessmentPointResponseBody
	GetAnswerList() []*GetAICoachAssessmentPointResponseBodyAnswerList
	SetCitations(v int32) *GetAICoachAssessmentPointResponseBody
	GetCitations() *int32
	SetDocumentId(v string) *GetAICoachAssessmentPointResponseBody
	GetDocumentId() *string
	SetDocumentName(v string) *GetAICoachAssessmentPointResponseBody
	GetDocumentName() *string
	SetGmtCreate(v string) *GetAICoachAssessmentPointResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetAICoachAssessmentPointResponseBody
	GetGmtModified() *string
	SetKbId(v string) *GetAICoachAssessmentPointResponseBody
	GetKbId() *string
	SetKbType(v string) *GetAICoachAssessmentPointResponseBody
	GetKbType() *string
	SetKnowledgeList(v []*string) *GetAICoachAssessmentPointResponseBody
	GetKnowledgeList() []*string
	SetName(v string) *GetAICoachAssessmentPointResponseBody
	GetName() *string
	SetPointId(v string) *GetAICoachAssessmentPointResponseBody
	GetPointId() *string
	SetQuestionDescription(v string) *GetAICoachAssessmentPointResponseBody
	GetQuestionDescription() *string
	SetQuestionSample(v string) *GetAICoachAssessmentPointResponseBody
	GetQuestionSample() *string
	SetRequestId(v string) *GetAICoachAssessmentPointResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetAICoachAssessmentPointResponseBody
	GetStatus() *string
}

type GetAICoachAssessmentPointResponseBody struct {
	AnswerList []*GetAICoachAssessmentPointResponseBodyAnswerList `json:"answerList,omitempty" xml:"answerList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Citations *int32 `json:"citations,omitempty" xml:"citations,omitempty"`
	// example:
	//
	// 1
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
	// example:
	//
	// demo
	DocumentName *string `json:"documentName,omitempty" xml:"documentName,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2025-02-24 12:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	KbId *string `json:"kbId,omitempty" xml:"kbId,omitempty"`
	// example:
	//
	// Cloudcode
	KbType        *string   `json:"kbType,omitempty" xml:"kbType,omitempty"`
	KnowledgeList []*string `json:"knowledgeList,omitempty" xml:"knowledgeList,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	PointId *string `json:"pointId,omitempty" xml:"pointId,omitempty"`
	// example:
	//
	// demo
	QuestionDescription *string `json:"questionDescription,omitempty" xml:"questionDescription,omitempty"`
	// example:
	//
	// demo
	QuestionSample *string `json:"questionSample,omitempty" xml:"questionSample,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4830493A-728F-5F19-BBCC-1443292E9C49
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBody) GetAnswerList() []*GetAICoachAssessmentPointResponseBodyAnswerList {
	return s.AnswerList
}

func (s *GetAICoachAssessmentPointResponseBody) GetCitations() *int32 {
	return s.Citations
}

func (s *GetAICoachAssessmentPointResponseBody) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetAICoachAssessmentPointResponseBody) GetDocumentName() *string {
	return s.DocumentName
}

func (s *GetAICoachAssessmentPointResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAICoachAssessmentPointResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAICoachAssessmentPointResponseBody) GetKbId() *string {
	return s.KbId
}

func (s *GetAICoachAssessmentPointResponseBody) GetKbType() *string {
	return s.KbType
}

func (s *GetAICoachAssessmentPointResponseBody) GetKnowledgeList() []*string {
	return s.KnowledgeList
}

func (s *GetAICoachAssessmentPointResponseBody) GetName() *string {
	return s.Name
}

func (s *GetAICoachAssessmentPointResponseBody) GetPointId() *string {
	return s.PointId
}

func (s *GetAICoachAssessmentPointResponseBody) GetQuestionDescription() *string {
	return s.QuestionDescription
}

func (s *GetAICoachAssessmentPointResponseBody) GetQuestionSample() *string {
	return s.QuestionSample
}

func (s *GetAICoachAssessmentPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICoachAssessmentPointResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAICoachAssessmentPointResponseBody) SetAnswerList(v []*GetAICoachAssessmentPointResponseBodyAnswerList) *GetAICoachAssessmentPointResponseBody {
	s.AnswerList = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetCitations(v int32) *GetAICoachAssessmentPointResponseBody {
	s.Citations = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetDocumentId(v string) *GetAICoachAssessmentPointResponseBody {
	s.DocumentId = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetDocumentName(v string) *GetAICoachAssessmentPointResponseBody {
	s.DocumentName = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetGmtCreate(v string) *GetAICoachAssessmentPointResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetGmtModified(v string) *GetAICoachAssessmentPointResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetKbId(v string) *GetAICoachAssessmentPointResponseBody {
	s.KbId = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetKbType(v string) *GetAICoachAssessmentPointResponseBody {
	s.KbType = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetKnowledgeList(v []*string) *GetAICoachAssessmentPointResponseBody {
	s.KnowledgeList = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetName(v string) *GetAICoachAssessmentPointResponseBody {
	s.Name = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetPointId(v string) *GetAICoachAssessmentPointResponseBody {
	s.PointId = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetQuestionDescription(v string) *GetAICoachAssessmentPointResponseBody {
	s.QuestionDescription = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetQuestionSample(v string) *GetAICoachAssessmentPointResponseBody {
	s.QuestionSample = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetRequestId(v string) *GetAICoachAssessmentPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) SetStatus(v string) *GetAICoachAssessmentPointResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBody) Validate() error {
	if s.AnswerList != nil {
		for _, item := range s.AnswerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachAssessmentPointResponseBodyAnswerList struct {
	AnswerValues []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues `json:"answerValues,omitempty" xml:"answerValues,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnabledKeyword *bool     `json:"enabledKeyword,omitempty" xml:"enabledKeyword,omitempty"`
	NameList       []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	// example:
	//
	// and
	Operators  *string                                                      `json:"operators,omitempty" xml:"operators,omitempty"`
	Parameters []*GetAICoachAssessmentPointResponseBodyAnswerListParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// example:
	//
	// custom
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerList) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) GetAnswerValues() []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	return s.AnswerValues
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) GetEnabledKeyword() *bool {
	return s.EnabledKeyword
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) GetNameList() []*string {
	return s.NameList
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) GetOperators() *string {
	return s.Operators
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) GetParameters() []*GetAICoachAssessmentPointResponseBodyAnswerListParameters {
	return s.Parameters
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) GetType() *string {
	return s.Type
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetAnswerValues(v []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.AnswerValues = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetEnabledKeyword(v bool) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.EnabledKeyword = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetNameList(v []*string) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.NameList = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetOperators(v string) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.Operators = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetParameters(v []*GetAICoachAssessmentPointResponseBodyAnswerListParameters) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.Parameters = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetType(v string) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.Type = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) SetWeight(v int32) *GetAICoachAssessmentPointResponseBodyAnswerList {
	s.Weight = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerList) Validate() error {
	if s.AnswerValues != nil {
		for _, item := range s.AnswerValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues struct {
	// example:
	//
	// demo
	AnswerName *string `json:"answerName,omitempty" xml:"answerName,omitempty"`
	// example:
	//
	// 50
	AnswerWeight  *int32                                                                      `json:"answerWeight,omitempty" xml:"answerWeight,omitempty"`
	KeywordValues []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues `json:"keywordValues,omitempty" xml:"keywordValues,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	KeywordWeight *int32                                                                     `json:"keywordWeight,omitempty" xml:"keywordWeight,omitempty"`
	ScoringRules  []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules `json:"scoringRules,omitempty" xml:"scoringRules,omitempty" type:"Repeated"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) GetAnswerName() *string {
	return s.AnswerName
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) GetAnswerWeight() *int32 {
	return s.AnswerWeight
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) GetKeywordValues() []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues {
	return s.KeywordValues
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) GetKeywordWeight() *int32 {
	return s.KeywordWeight
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) GetScoringRules() []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules {
	return s.ScoringRules
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetAnswerName(v string) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.AnswerName = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetAnswerWeight(v int32) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.AnswerWeight = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetKeywordValues(v []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.KeywordValues = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetKeywordWeight(v int32) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.KeywordWeight = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) SetScoringRules(v []*GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues {
	s.ScoringRules = v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValues) Validate() error {
	if s.KeywordValues != nil {
		for _, item := range s.KeywordValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScoringRules != nil {
		for _, item := range s.ScoringRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 50
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) GetName() *string {
	return s.Name
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) GetWeight() *int32 {
	return s.Weight
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) SetName(v string) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues {
	s.Name = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) SetWeight(v int32) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues {
	s.Weight = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesKeywordValues) Validate() error {
	return dara.Validate(s)
}

type GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) GetName() *string {
	return s.Name
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) SetName(v string) *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules {
	s.Name = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListAnswerValuesScoringRules) Validate() error {
	return dara.Validate(s)
}

type GetAICoachAssessmentPointResponseBodyAnswerListParameters struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 441323200602114284
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListParameters) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachAssessmentPointResponseBodyAnswerListParameters) GoString() string {
	return s.String()
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListParameters) GetName() *string {
	return s.Name
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListParameters) GetValue() *string {
	return s.Value
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListParameters) SetName(v string) *GetAICoachAssessmentPointResponseBodyAnswerListParameters {
	s.Name = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListParameters) SetValue(v string) *GetAICoachAssessmentPointResponseBodyAnswerListParameters {
	s.Value = &v
	return s
}

func (s *GetAICoachAssessmentPointResponseBodyAnswerListParameters) Validate() error {
	return dara.Validate(s)
}
