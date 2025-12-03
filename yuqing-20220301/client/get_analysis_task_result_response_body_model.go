// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnalysisTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisId(v int64) *GetAnalysisTaskResultResponseBody
	GetAnalysisId() *int64
	SetRequestId(v string) *GetAnalysisTaskResultResponseBody
	GetRequestId() *string
	SetResultJson(v string) *GetAnalysisTaskResultResponseBody
	GetResultJson() *string
}

type GetAnalysisTaskResultResponseBody struct {
	AnalysisId *int64  `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultJson *string `json:"resultJson,omitempty" xml:"resultJson,omitempty"`
}

func (s GetAnalysisTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAnalysisTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAnalysisTaskResultResponseBody) GetAnalysisId() *int64 {
	return s.AnalysisId
}

func (s *GetAnalysisTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAnalysisTaskResultResponseBody) GetResultJson() *string {
	return s.ResultJson
}

func (s *GetAnalysisTaskResultResponseBody) SetAnalysisId(v int64) *GetAnalysisTaskResultResponseBody {
	s.AnalysisId = &v
	return s
}

func (s *GetAnalysisTaskResultResponseBody) SetRequestId(v string) *GetAnalysisTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAnalysisTaskResultResponseBody) SetResultJson(v string) *GetAnalysisTaskResultResponseBody {
	s.ResultJson = &v
	return s
}

func (s *GetAnalysisTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}
