// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyseType(v string) *SubmitAnalysisTaskRequest
	GetAnalyseType() *string
	SetRequestId(v string) *SubmitAnalysisTaskRequest
	GetRequestId() *string
	SetSearchCondition(v *SearchCondition) *SubmitAnalysisTaskRequest
	GetSearchCondition() *SearchCondition
	SetTeamHashId(v string) *SubmitAnalysisTaskRequest
	GetTeamHashId() *string
}

type SubmitAnalysisTaskRequest struct {
	AnalyseType     *string          `json:"analyseType,omitempty" xml:"analyseType,omitempty"`
	RequestId       *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SearchCondition *SearchCondition `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	TeamHashId      *string          `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s SubmitAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisTaskRequest) GetAnalyseType() *string {
	return s.AnalyseType
}

func (s *SubmitAnalysisTaskRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAnalysisTaskRequest) GetSearchCondition() *SearchCondition {
	return s.SearchCondition
}

func (s *SubmitAnalysisTaskRequest) GetTeamHashId() *string {
	return s.TeamHashId
}

func (s *SubmitAnalysisTaskRequest) SetAnalyseType(v string) *SubmitAnalysisTaskRequest {
	s.AnalyseType = &v
	return s
}

func (s *SubmitAnalysisTaskRequest) SetRequestId(v string) *SubmitAnalysisTaskRequest {
	s.RequestId = &v
	return s
}

func (s *SubmitAnalysisTaskRequest) SetSearchCondition(v *SearchCondition) *SubmitAnalysisTaskRequest {
	s.SearchCondition = v
	return s
}

func (s *SubmitAnalysisTaskRequest) SetTeamHashId(v string) *SubmitAnalysisTaskRequest {
	s.TeamHashId = &v
	return s
}

func (s *SubmitAnalysisTaskRequest) Validate() error {
	if s.SearchCondition != nil {
		if err := s.SearchCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
