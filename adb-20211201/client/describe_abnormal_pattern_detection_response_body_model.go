// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbnormalPatternDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAbnormalPatternDetectionResponseBody
	GetDBClusterId() *string
	SetDetectionItems(v []*DescribeAbnormalPatternDetectionResponseBodyDetectionItems) *DescribeAbnormalPatternDetectionResponseBody
	GetDetectionItems() []*DescribeAbnormalPatternDetectionResponseBodyDetectionItems
	SetRequestId(v string) *DescribeAbnormalPatternDetectionResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeAbnormalPatternDetectionResponseBody
	GetTotalCount() *string
}

type DescribeAbnormalPatternDetectionResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// amv-xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeAbnormalPatternDetectionResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 19B824E0-690D-5A78-9992-5398C2F43694
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAbnormalPatternDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalPatternDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalPatternDetectionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAbnormalPatternDetectionResponseBody) GetDetectionItems() []*DescribeAbnormalPatternDetectionResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeAbnormalPatternDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAbnormalPatternDetectionResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeAbnormalPatternDetectionResponseBody) SetDBClusterId(v string) *DescribeAbnormalPatternDetectionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBody) SetDetectionItems(v []*DescribeAbnormalPatternDetectionResponseBodyDetectionItems) *DescribeAbnormalPatternDetectionResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBody) SetRequestId(v string) *DescribeAbnormalPatternDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBody) SetTotalCount(v string) *DescribeAbnormalPatternDetectionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBody) Validate() error {
	if s.DetectionItems != nil {
		for _, item := range s.DetectionItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAbnormalPatternDetectionResponseBodyDetectionItems struct {
	// The name of the detection item.
	//
	// example:
	//
	// Cost
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The detection result items.
	Results []*DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DescribeAbnormalPatternDetectionResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalPatternDetectionResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItems) GetResults() []*DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	return s.Results
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItems) SetName(v string) *DescribeAbnormalPatternDetectionResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItems) SetResults(v []*DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) *DescribeAbnormalPatternDetectionResponseBodyDetectionItems {
	s.Results = v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItems) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults struct {
	// The IP address of the SQL client that submits the SQL pattern.
	//
	// example:
	//
	// 172.16.133.168
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The description of the detection result.
	//
	// example:
	//
	// Two SQL patterns that have abnormal totalTime metric values are detected. This may result in increased CPU utilization, query slowdown, and degraded system stability. Go to the monitoring page to diagnose the issue and then perform optimization.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The number of failed SQL patterns within the time range.
	//
	// example:
	//
	// 7
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The SQL pattern ID.
	//
	// example:
	//
	// 2803084667741875187
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The number of queries.
	//
	// example:
	//
	// 72
	QueryCount *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The metrics related to the SQL pattern.
	//
	// example:
	//
	// Maximum query duration: 1.43s
	//
	// Maximum peak memory: 20.73 MB
	//
	// Maximum read table data: 10.12 MB
	RelatedMetrics *string `json:"RelatedMetrics,omitempty" xml:"RelatedMetrics,omitempty"`
	// The SQL statement that represents the SQL pattern.
	//
	// example:
	//
	// SELECT `tsid`nFROM `prod_ods_marketing_engine_material`nWHERE `tsid` = ?nLIMIT ?
	SQLPattern *string `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	// The names of tables.
	//
	// example:
	//
	// dw_student_exam.ods_school_queanal
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The name of the database account that is used to submit the query.
	//
	// example:
	//
	// test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetAccessIp() *string {
	return s.AccessIp
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetDetail() *string {
	return s.Detail
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetPatternId() *string {
	return s.PatternId
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetQueryCount() *int64 {
	return s.QueryCount
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetRelatedMetrics() *string {
	return s.RelatedMetrics
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetSQLPattern() *string {
	return s.SQLPattern
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetTables() *string {
	return s.Tables
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) GetUser() *string {
	return s.User
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetAccessIp(v string) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.AccessIp = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetDetail(v string) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.Detail = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetFailedCount(v int64) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.FailedCount = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetPatternId(v string) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.PatternId = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetQueryCount(v int64) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.QueryCount = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetRelatedMetrics(v string) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.RelatedMetrics = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetSQLPattern(v string) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.SQLPattern = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetTables(v string) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.Tables = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) SetUser(v string) *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults {
	s.User = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponseBodyDetectionItemsResults) Validate() error {
	return dara.Validate(s)
}
