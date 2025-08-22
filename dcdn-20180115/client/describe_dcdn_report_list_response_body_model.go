// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeDcdnReportListResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeDcdnReportListResponseBody
	GetRequestId() *string
}

type DescribeDcdnReportListResponseBody struct {
	// The information about the operations report.
	//
	// example:
	//
	// "data": [{"reportId":2,"deliver":{"report":{"title":"DomainPvUv","format":"chart","shape":"line","xAxis":"ds","yAxis":"cnt","legend":"cnt_type","header":["ds","cnt_type","cnt"]}}}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportListResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnReportListResponseBody) SetContent(v string) *DescribeDcdnReportListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnReportListResponseBody) SetRequestId(v string) *DescribeDcdnReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnReportListResponseBody) Validate() error {
	return dara.Validate(s)
}
