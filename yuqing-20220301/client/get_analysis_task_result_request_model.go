// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnalysisTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisId(v int64) *GetAnalysisTaskResultRequest
	GetAnalysisId() *int64
	SetRequestId(v string) *GetAnalysisTaskResultRequest
	GetRequestId() *string
	SetTeamHashId(v string) *GetAnalysisTaskResultRequest
	GetTeamHashId() *string
}

type GetAnalysisTaskResultRequest struct {
	AnalysisId *int64  `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TeamHashId *string `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s GetAnalysisTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAnalysisTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetAnalysisTaskResultRequest) GetAnalysisId() *int64 {
	return s.AnalysisId
}

func (s *GetAnalysisTaskResultRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAnalysisTaskResultRequest) GetTeamHashId() *string {
	return s.TeamHashId
}

func (s *GetAnalysisTaskResultRequest) SetAnalysisId(v int64) *GetAnalysisTaskResultRequest {
	s.AnalysisId = &v
	return s
}

func (s *GetAnalysisTaskResultRequest) SetRequestId(v string) *GetAnalysisTaskResultRequest {
	s.RequestId = &v
	return s
}

func (s *GetAnalysisTaskResultRequest) SetTeamHashId(v string) *GetAnalysisTaskResultRequest {
	s.TeamHashId = &v
	return s
}

func (s *GetAnalysisTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
