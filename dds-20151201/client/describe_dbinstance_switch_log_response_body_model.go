// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSwitchLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceSwitchLogResponseBody
	GetDBInstanceId() *string
	SetLogItems(v []*DescribeDBInstanceSwitchLogResponseBodyLogItems) *DescribeDBInstanceSwitchLogResponseBody
	GetLogItems() []*DescribeDBInstanceSwitchLogResponseBodyLogItems
	SetPageNumber(v int64) *DescribeDBInstanceSwitchLogResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDBInstanceSwitchLogResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDBInstanceSwitchLogResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDBInstanceSwitchLogResponseBody
	GetTotalCount() *int64
}

type DescribeDBInstanceSwitchLogResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// dds-uf68f1b5a57exxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The primary/secondary switchover logs.
	LogItems []*DescribeDBInstanceSwitchLogResponseBodyLogItems `json:"LogItems,omitempty" xml:"LogItems,omitempty" type:"Repeated"`
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ECBCA991-XXXX-XXXX-834C-B3E8007F33AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of primary/secondary switching entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceSwitchLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSwitchLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetLogItems() []*DescribeDBInstanceSwitchLogResponseBodyLogItems {
	return s.LogItems
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceSwitchLogResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetLogItems(v []*DescribeDBInstanceSwitchLogResponseBodyLogItems) *DescribeDBInstanceSwitchLogResponseBody {
	s.LogItems = v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetPageNumber(v int64) *DescribeDBInstanceSwitchLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetPageSize(v int64) *DescribeDBInstanceSwitchLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetRequestId(v string) *DescribeDBInstanceSwitchLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetTotalCount(v int64) *DescribeDBInstanceSwitchLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceSwitchLogResponseBodyLogItems struct {
	// The ID of the replica set instance or the ID of the node on which a primary/secondary switchover is performed.
	//
	// example:
	//
	// dds-uf68f1b5a57exxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The code that indicates the reason of a primary/secondary switchover. Valid values:
	//
	// 	- USER_CONSOLE_OPERATION: The switchover is manually performed.
	//
	// 	- OPERATION_AND_MAINTENANCE: Potential risks exist.
	//
	// 	- MACHINE_DOWNTIME: The host is offline.
	//
	// 	- PRIMARY_UNHEALTHY: An exception occurs on the primary node of the instance.
	//
	// 	- SECONDARY_UNHEALTHY: An exception occurs on the secondary node of the instance.
	//
	// 	- MULTIPLE_NODE_FAILURES: An exception occurs on multiple nodes of the instance.
	//
	// example:
	//
	// USER_CONSOLE_OPERATION
	SwitchCode *string `json:"SwitchCode,omitempty" xml:"SwitchCode,omitempty"`
	// The switchover status. Valid values: **1*	- and **0**. The value 1 indicates a successful primary/secondary switchover and the value 0 indicates a failed primary/secondary switchover.
	//
	// example:
	//
	// 1
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	// The point in time when a primary/secondary switchover was performed. The time follows the ISO 8601 standard in the *yyyy-mm-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-07T18:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s DescribeDBInstanceSwitchLogResponseBodyLogItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSwitchLogResponseBodyLogItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) GetSwitchCode() *string {
	return s.SwitchCode
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) SetNodeId(v string) *DescribeDBInstanceSwitchLogResponseBodyLogItems {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) SetSwitchCode(v string) *DescribeDBInstanceSwitchLogResponseBodyLogItems {
	s.SwitchCode = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) SetSwitchStatus(v string) *DescribeDBInstanceSwitchLogResponseBodyLogItems {
	s.SwitchStatus = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) SetSwitchTime(v string) *DescribeDBInstanceSwitchLogResponseBodyLogItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyLogItems) Validate() error {
	return dara.Validate(s)
}
