// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorReportComponentSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorReportComponentSummaryResponseBodyData) *GetDoctorReportComponentSummaryResponseBody
	GetData() *GetDoctorReportComponentSummaryResponseBodyData
	SetRequestId(v string) *GetDoctorReportComponentSummaryResponseBody
	GetRequestId() *string
}

type GetDoctorReportComponentSummaryResponseBody struct {
	// The content of the report.
	Data *GetDoctorReportComponentSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorReportComponentSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorReportComponentSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorReportComponentSummaryResponseBody) GetData() *GetDoctorReportComponentSummaryResponseBodyData {
	return s.Data
}

func (s *GetDoctorReportComponentSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorReportComponentSummaryResponseBody) SetData(v *GetDoctorReportComponentSummaryResponseBodyData) *GetDoctorReportComponentSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorReportComponentSummaryResponseBody) SetRequestId(v string) *GetDoctorReportComponentSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorReportComponentSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDoctorReportComponentSummaryResponseBodyData struct {
	// Score.
	//
	// example:
	//
	// 88
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// Optimization suggestions.
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The summary of the report.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetDoctorReportComponentSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorReportComponentSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorReportComponentSummaryResponseBodyData) GetScore() *int32 {
	return s.Score
}

func (s *GetDoctorReportComponentSummaryResponseBodyData) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetDoctorReportComponentSummaryResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *GetDoctorReportComponentSummaryResponseBodyData) SetScore(v int32) *GetDoctorReportComponentSummaryResponseBodyData {
	s.Score = &v
	return s
}

func (s *GetDoctorReportComponentSummaryResponseBodyData) SetSuggestion(v string) *GetDoctorReportComponentSummaryResponseBodyData {
	s.Suggestion = &v
	return s
}

func (s *GetDoctorReportComponentSummaryResponseBodyData) SetSummary(v string) *GetDoctorReportComponentSummaryResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetDoctorReportComponentSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
