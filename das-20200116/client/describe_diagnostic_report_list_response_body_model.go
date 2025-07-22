// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDiagnosticReportListResponseBody
	GetCode() *string
	SetData(v string) *DescribeDiagnosticReportListResponseBody
	GetData() *string
	SetMessage(v string) *DescribeDiagnosticReportListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDiagnosticReportListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDiagnosticReportListResponseBody
	GetSuccess() *string
	SetSynchro(v string) *DescribeDiagnosticReportListResponseBody
	GetSynchro() *string
}

type DescribeDiagnosticReportListResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information of the diagnostics reports. Valid values:
	//
	// 	- **total**: the number of diagnostics reports.
	//
	// 	- **score**: the health score.
	//
	// 	- **diagnosticTime**: the time when the diagnostics report was generated. The time is displayed in UTC.
	//
	// 	- **startTime**: the start time of the query. The time is displayed in UTC.
	//
	// 	- **endTime**: the end time of the query. The time is displayed in UTC.
	//
	// example:
	//
	// {     "total": 1,     "list": [       {         "score": 100,         "diagnosticTime": "2022-11-14T08:17:00Z",         "startTime": "2022-11-14T07:16:59Z",         "endTime": "2022-11-14T08:16:59Z"       }     ]   }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D00DB161-FEF6-5428-B37A-8D29A4C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Synchro *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DescribeDiagnosticReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDiagnosticReportListResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeDiagnosticReportListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDiagnosticReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosticReportListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDiagnosticReportListResponseBody) GetSynchro() *string {
	return s.Synchro
}

func (s *DescribeDiagnosticReportListResponseBody) SetCode(v string) *DescribeDiagnosticReportListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetData(v string) *DescribeDiagnosticReportListResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetMessage(v string) *DescribeDiagnosticReportListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetRequestId(v string) *DescribeDiagnosticReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetSuccess(v string) *DescribeDiagnosticReportListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) SetSynchro(v string) *DescribeDiagnosticReportListResponseBody {
	s.Synchro = &v
	return s
}

func (s *DescribeDiagnosticReportListResponseBody) Validate() error {
	return dara.Validate(s)
}
