// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryExplainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExplainResult(v string) *DescribeQueryExplainResponseBody
	GetExplainResult() *string
	SetRequestId(v string) *DescribeQueryExplainResponseBody
	GetRequestId() *string
	SetSql(v string) *DescribeQueryExplainResponseBody
	GetSql() *string
}

type DescribeQueryExplainResponseBody struct {
	// The Explain result.
	//
	// example:
	//
	// +-----------------------------------------------------------------------+
	//
	// | Explain String|
	//
	// +-----------------------------------------------------------------------+
	//
	// | PLAN FRAGMENT 0                                                       |
	//
	// |OUTPUT EXPRS:                                                        |
	//
	// |    name[#1]                                                           |
	//
	// |    age[#2]                                                            |
	//
	// |  PARTITION: UNPARTITIONED                                |
	//
	// |                                                                        |
	//
	// |  VRESULT SINK                                                         |
	//
	// |                                                                        |
	//
	// |  1:VEXCHANGE                                                          |
	//
	// |     offset: 0                                                        |
	//
	// |                                                                        |
	//
	// | PLAN FRAGMENT 1                                                       |
	//
	// |                                                                        |
	//
	// |  PARTITION: HASH_PARTITIONED: id[#0]                                  |
	//
	// |                                                                        |
	//
	// |  STREAM DATA SINK                                                     |
	//
	// |    EXCHANGE ID: 01|
	//
	// |    UNPARTITIONED                                |
	//
	// |                                                                        |
	//
	// |  0:VOlapScanNode                                                      |
	//
	// |     TABLE: example_db.example_tbl(example_tbl)                        |
	//
	// |     PREAGGREGATION: ON                                                |
	//
	// |     PREDICATES: (age[#2] > 18)                                        |
	//
	// |     cardinality=1, avgRowSize=20.0, numNodes=1                        |
	//
	// |     tablet list: 10023, 10025, 10027                |
	//
	// +-----------------------------------------------------------------------+
	ExplainResult *string `json:"ExplainResult,omitempty" xml:"ExplainResult,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SQL statement for which the execution plan is retrieved. Excessively long SQL statements in audit logs may be truncated.
	//
	// example:
	//
	// SELECT 	- FROM example_db.example_tbl
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s DescribeQueryExplainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryExplainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQueryExplainResponseBody) GetExplainResult() *string {
	return s.ExplainResult
}

func (s *DescribeQueryExplainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQueryExplainResponseBody) GetSql() *string {
	return s.Sql
}

func (s *DescribeQueryExplainResponseBody) SetExplainResult(v string) *DescribeQueryExplainResponseBody {
	s.ExplainResult = &v
	return s
}

func (s *DescribeQueryExplainResponseBody) SetRequestId(v string) *DescribeQueryExplainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQueryExplainResponseBody) SetSql(v string) *DescribeQueryExplainResponseBody {
	s.Sql = &v
	return s
}

func (s *DescribeQueryExplainResponseBody) Validate() error {
	return dara.Validate(s)
}
