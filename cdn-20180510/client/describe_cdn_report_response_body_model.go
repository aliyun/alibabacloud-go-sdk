// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *DescribeCdnReportResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *DescribeCdnReportResponseBody
	GetRequestId() *string
}

type DescribeCdnReportResponseBody struct {
	// The content of the operations report.
	//
	// example:
	//
	// "data":[{"deliver":{"report":{"title":"TopUrlByAcc","format":"table","shape":"","header":["url","traf","traf_rate","acc","acc_rate"]}},"data":[{"acc":440,"acc_rate":"0.200%","traf":22,"url":"http://demo.com","traf_rate":"0.100%"},{"acc":440,"acc_rate":"0.200%","traf":22,"url":"http://demo.com","traf_rate":"0.100%"}]}]}}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *DescribeCdnReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnReportResponseBody) SetContent(v map[string]interface{}) *DescribeCdnReportResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCdnReportResponseBody) SetRequestId(v string) *DescribeCdnReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnReportResponseBody) Validate() error {
	return dara.Validate(s)
}
