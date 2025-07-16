// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBenchmarkTaskReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DescribeBenchmarkTaskReportResponseBody
	GetData() interface{}
	SetReportUrl(v string) *DescribeBenchmarkTaskReportResponseBody
	GetReportUrl() *string
	SetRequestId(v string) *DescribeBenchmarkTaskReportResponseBody
	GetRequestId() *string
}

type DescribeBenchmarkTaskReportResponseBody struct {
	// If the value of ReportType is set to RAW, the details about the stress testing report are returned.
	//
	// example:
	//
	// {
	//
	//     "TimestampList": ["int64"],
	//
	//     "QPSList": ["float32"],
	//
	//     "RTList": [
	//
	//       {
	//
	//         "AVG": "float32",
	//
	//         "TP100": "float32",
	//
	//         "TP99": "float32",
	//
	//         "TP90": "float32",
	//
	//         "TP50": "float32",
	//
	//         "TP10": "float32"
	//
	//       }
	//
	//     ],
	//
	//     "TrafficList": [
	//
	//       {
	//
	//         "Send": "float64",
	//
	//         "Receive": "float64"
	//
	//       }
	//
	//     ],
	//
	//     "StatusCode": {
	//
	//       "200": "uint64",
	//
	//       "450": "uint64",
	//
	//       "500": "uint64"
	//
	//     },
	//
	//     "Count": "uint64",
	//
	//     "Total": "float64",
	//
	//     "MinRT": "float32",
	//
	//     "AvgRT": "float32",
	//
	//     "MaxRT": "float32",
	//
	//     "QPS": "float32",
	//
	//     "TotalSend": "float64",
	//
	//     "TotalReceive": "float64",
	//
	//     "RTDistribution": [
	//
	//       {
	//
	//         "Latency": "float32",
	//
	//         "Percentage": "int"
	//
	//       }
	//
	//     ],
	//
	//     "RTHistogram": [
	//
	//       {
	//
	//         "Count": "int",
	//
	//         "Mark": "float32",
	//
	//         "Frequency": "float32"
	//
	//       }
	//
	//     ]
	//
	//   }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// If the value of ReportType is set to Report, the URL of the stress testing report is returned.
	//
	// example:
	//
	// http://eas-benchmark.oss-cn-chengdu.aliyuncs.com/summary/benchmark-larec-test-015d-10007.html
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBenchmarkTaskReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBenchmarkTaskReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskReportResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DescribeBenchmarkTaskReportResponseBody) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *DescribeBenchmarkTaskReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBenchmarkTaskReportResponseBody) SetData(v interface{}) *DescribeBenchmarkTaskReportResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBenchmarkTaskReportResponseBody) SetReportUrl(v string) *DescribeBenchmarkTaskReportResponseBody {
	s.ReportUrl = &v
	return s
}

func (s *DescribeBenchmarkTaskReportResponseBody) SetRequestId(v string) *DescribeBenchmarkTaskReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBenchmarkTaskReportResponseBody) Validate() error {
	return dara.Validate(s)
}
