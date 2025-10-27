// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResubmitConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeResubmitConfigResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DescribeResubmitConfigResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeResubmitConfigResponseBodyRules) *DescribeResubmitConfigResponseBody
	GetRules() []*DescribeResubmitConfigResponseBodyRules
}

type DescribeResubmitConfigResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// example:
	//
	// am-8vbyw9awuj141haf9
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8D217417-BBA7-566C-9B9D-FFCBC86112B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job resubmission rules.
	Rules []*DescribeResubmitConfigResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeResubmitConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResubmitConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResubmitConfigResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeResubmitConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResubmitConfigResponseBody) GetRules() []*DescribeResubmitConfigResponseBodyRules {
	return s.Rules
}

func (s *DescribeResubmitConfigResponseBody) SetDBClusterId(v string) *DescribeResubmitConfigResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeResubmitConfigResponseBody) SetRequestId(v string) *DescribeResubmitConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResubmitConfigResponseBody) SetRules(v []*DescribeResubmitConfigResponseBodyRules) *DescribeResubmitConfigResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeResubmitConfigResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResubmitConfigResponseBodyRules struct {
	// Indicates whether out-of-memory (OOM) check is configured.
	//
	// example:
	//
	// false
	ExceedMemoryException *bool `json:"ExceedMemoryException,omitempty" xml:"ExceedMemoryException,omitempty"`
	// The name of the source resource group.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The peak memory usage.
	//
	// example:
	//
	// 32
	PeakMemory *string `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The duration of the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 300
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The name of the destination resource group.
	//
	// example:
	//
	// test_target_group
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s DescribeResubmitConfigResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeResubmitConfigResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeResubmitConfigResponseBodyRules) GetExceedMemoryException() *bool {
	return s.ExceedMemoryException
}

func (s *DescribeResubmitConfigResponseBodyRules) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeResubmitConfigResponseBodyRules) GetPeakMemory() *string {
	return s.PeakMemory
}

func (s *DescribeResubmitConfigResponseBodyRules) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeResubmitConfigResponseBodyRules) GetTargetGroupName() *string {
	return s.TargetGroupName
}

func (s *DescribeResubmitConfigResponseBodyRules) SetExceedMemoryException(v bool) *DescribeResubmitConfigResponseBodyRules {
	s.ExceedMemoryException = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) SetGroupName(v string) *DescribeResubmitConfigResponseBodyRules {
	s.GroupName = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) SetPeakMemory(v string) *DescribeResubmitConfigResponseBodyRules {
	s.PeakMemory = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) SetQueryTime(v string) *DescribeResubmitConfigResponseBodyRules {
	s.QueryTime = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) SetTargetGroupName(v string) *DescribeResubmitConfigResponseBodyRules {
	s.TargetGroupName = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
