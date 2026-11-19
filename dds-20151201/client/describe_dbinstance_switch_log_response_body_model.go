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
	// The list of primary/secondary switchover logs.
	LogItems []*DescribeDBInstanceSwitchLogResponseBodyLogItems `json:"LogItems,omitempty" xml:"LogItems,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ECBCA991-XXXX-XXXX-834C-B3E8007F33AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of primary/secondary switchover records returned in the query.
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
	if s.LogItems != nil {
		for _, item := range s.LogItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceSwitchLogResponseBodyLogItems struct {
	// Instance ID of the instance if the instance is a replica set instance, or instance ID of the switched node if the instance is a sharded cluster instance.
	//
	// example:
	//
	// If the instance is a replica set instance: dds-uf68f1b5a57exxxx.
	//
	// If the instance is a sharded cluster instance: d-uf652b73f554xxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The cause code of the switchover. Valid values:
	//
	// - USER_CONSOLE_OPERATION: manual operation
	//
	// - OPERATION_AND_MAINTENANCE: potential risk
	//
	// - MACHINE_DOWNTIME: host offline
	//
	// - PRIMARY_UNHEALTHY: instance failures
	//
	// - SECONDARY_UNHEALTHY: instance failures
	//
	// - MULTIPLE_NODE_FAILURES: instance failures
	//
	// example:
	//
	// USER_CONSOLE_OPERATION
	SwitchCode *string `json:"SwitchCode,omitempty" xml:"SwitchCode,omitempty"`
	// The switchover status. Valid values:
	//
	// - **1**: The switchover is successful.
	//
	// - **0**: The switchover failed.
	//
	// example:
	//
	// 1
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	// The switchover time. The time is in the <i>yyyy-mm-dd</i>T<i>hh:mm:ss</i>Z format (UTC).
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
