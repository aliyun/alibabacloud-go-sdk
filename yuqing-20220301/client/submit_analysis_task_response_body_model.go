// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisId(v int64) *SubmitAnalysisTaskResponseBody
	GetAnalysisId() *int64
	SetRequestId(v string) *SubmitAnalysisTaskResponseBody
	GetRequestId() *string
	SetResultJson(v string) *SubmitAnalysisTaskResponseBody
	GetResultJson() *string
}

type SubmitAnalysisTaskResponseBody struct {
	AnalysisId *int64  `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultJson *string `json:"resultJson,omitempty" xml:"resultJson,omitempty"`
}

func (s SubmitAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisTaskResponseBody) GetAnalysisId() *int64 {
	return s.AnalysisId
}

func (s *SubmitAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAnalysisTaskResponseBody) GetResultJson() *string {
	return s.ResultJson
}

func (s *SubmitAnalysisTaskResponseBody) SetAnalysisId(v int64) *SubmitAnalysisTaskResponseBody {
	s.AnalysisId = &v
	return s
}

func (s *SubmitAnalysisTaskResponseBody) SetRequestId(v string) *SubmitAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAnalysisTaskResponseBody) SetResultJson(v string) *SubmitAnalysisTaskResponseBody {
	s.ResultJson = &v
	return s
}

func (s *SubmitAnalysisTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
