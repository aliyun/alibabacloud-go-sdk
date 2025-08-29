// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportABMetricGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentReport(v map[string]*ExperimentReportValue) *ReportABMetricGroupResponseBody
	GetExperimentReport() map[string]*ExperimentReportValue
	SetGroupDimension(v []*string) *ReportABMetricGroupResponseBody
	GetGroupDimension() []*string
	SetRequestId(v string) *ReportABMetricGroupResponseBody
	GetRequestId() *string
}

type ReportABMetricGroupResponseBody struct {
	ExperimentReport map[string]*ExperimentReportValue `json:"ExperimentReport,omitempty" xml:"ExperimentReport,omitempty"`
	GroupDimension   []*string                         `json:"GroupDimension,omitempty" xml:"GroupDimension,omitempty" type:"Repeated"`
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportABMetricGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ReportABMetricGroupResponseBody) GetExperimentReport() map[string]*ExperimentReportValue {
	return s.ExperimentReport
}

func (s *ReportABMetricGroupResponseBody) GetGroupDimension() []*string {
	return s.GroupDimension
}

func (s *ReportABMetricGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportABMetricGroupResponseBody) SetExperimentReport(v map[string]*ExperimentReportValue) *ReportABMetricGroupResponseBody {
	s.ExperimentReport = v
	return s
}

func (s *ReportABMetricGroupResponseBody) SetGroupDimension(v []*string) *ReportABMetricGroupResponseBody {
	s.GroupDimension = v
	return s
}

func (s *ReportABMetricGroupResponseBody) SetRequestId(v string) *ReportABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportABMetricGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
