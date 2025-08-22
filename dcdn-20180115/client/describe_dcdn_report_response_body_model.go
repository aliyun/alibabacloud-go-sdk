// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DescribeDcdnReportResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DescribeDcdnReportResponseBody
	GetRequestId() *string
}

type DescribeDcdnReportResponseBody struct {
	// The content of the operations report.
	//
	// example:
	//
	// "data":[{"deliver":{ "report":{"title":"TopUrlByAcc","format":"table","sape":"","header":["url","traf","traf_rate","acc","acc_rate"]}}, "data":[{"acc":440,"acc_rate":"0.200%","traf":22,"url":"http://example.com","traf_rate":"0.100%"},{"acc":440,"acc_rate":"0.200%","traf":22,"url":"http://example.org","traf_rate":"0.100%"}]}]
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DescribeDcdnReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnReportResponseBody) SetContent(v map[string]interface{}) *DescribeDcdnReportResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnReportResponseBody) SetRequestId(v string) *DescribeDcdnReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnReportResponseBody) Validate() error {
	return dara.Validate(s)
}
