// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatisticReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListStatisticReportResponseBody
	GetRequestId() *string
	SetResult(v []map[string]interface{}) *ListStatisticReportResponseBody
	GetResult() []map[string]interface{}
	SetTotalCount(v int64) *ListStatisticReportResponseBody
	GetTotalCount() *int64
}

type ListStatisticReportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F65C8BB2-C14F-5983-888B-41C4E082D3BC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The statistical reports. Valid values:
	//
	// 	- For more information about the metrics in data quality reports, see the Upload behavioral data section of [Data collection 2.0](https://help.aliyun.com/document_detail/131547.html).
	//
	// 	- For more information about the metrics in application and A/B test reports, see the Core metrics section of [Metrics of statistical reports](https://help.aliyun.com/document_detail/187665.html).
	//
	// 	- For more information about the metrics in query analysis reports, see the Query analysis metrics section of [Metrics of statistical reports](https://help.aliyun.com/document_detail/187665.html).
	//
	// example:
	//
	// []
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 43
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStatisticReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStatisticReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatisticReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStatisticReportResponseBody) GetResult() []map[string]interface{} {
	return s.Result
}

func (s *ListStatisticReportResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListStatisticReportResponseBody) SetRequestId(v string) *ListStatisticReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStatisticReportResponseBody) SetResult(v []map[string]interface{}) *ListStatisticReportResponseBody {
	s.Result = v
	return s
}

func (s *ListStatisticReportResponseBody) SetTotalCount(v int64) *ListStatisticReportResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStatisticReportResponseBody) Validate() error {
	return dara.Validate(s)
}
