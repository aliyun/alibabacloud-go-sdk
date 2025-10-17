// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHALogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeHALogsResponseBody
	GetDBInstanceName() *string
	SetDBInstanceType(v string) *DescribeHALogsResponseBody
	GetDBInstanceType() *string
	SetHaLogItems(v []*DescribeHALogsResponseBodyHaLogItems) *DescribeHALogsResponseBody
	GetHaLogItems() []*DescribeHALogsResponseBodyHaLogItems
	SetHaStatus(v int32) *DescribeHALogsResponseBody
	GetHaStatus() *int32
	SetItemsNumbers(v int32) *DescribeHALogsResponseBody
	GetItemsNumbers() *int32
	SetPageNumber(v int32) *DescribeHALogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHALogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHALogsResponseBody
	GetRequestId() *string
	SetTotalRecords(v int32) *DescribeHALogsResponseBody
	GetTotalRecords() *int32
}

type DescribeHALogsResponseBody struct {
	// example:
	//
	// pc-a*************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// polardb_mysql_rw
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The failover logs.
	HaLogItems []*DescribeHALogsResponseBodyHaLogItems `json:"HaLogItems,omitempty" xml:"HaLogItems,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	HaStatus *int32 `json:"HaStatus,omitempty" xml:"HaStatus,omitempty"`
	// example:
	//
	// 10
	ItemsNumbers *int32 `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 160
	TotalRecords *int32 `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeHALogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHALogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHALogsResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeHALogsResponseBody) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeHALogsResponseBody) GetHaLogItems() []*DescribeHALogsResponseBodyHaLogItems {
	return s.HaLogItems
}

func (s *DescribeHALogsResponseBody) GetHaStatus() *int32 {
	return s.HaStatus
}

func (s *DescribeHALogsResponseBody) GetItemsNumbers() *int32 {
	return s.ItemsNumbers
}

func (s *DescribeHALogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHALogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHALogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHALogsResponseBody) GetTotalRecords() *int32 {
	return s.TotalRecords
}

func (s *DescribeHALogsResponseBody) SetDBInstanceName(v string) *DescribeHALogsResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeHALogsResponseBody) SetDBInstanceType(v string) *DescribeHALogsResponseBody {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeHALogsResponseBody) SetHaLogItems(v []*DescribeHALogsResponseBodyHaLogItems) *DescribeHALogsResponseBody {
	s.HaLogItems = v
	return s
}

func (s *DescribeHALogsResponseBody) SetHaStatus(v int32) *DescribeHALogsResponseBody {
	s.HaStatus = &v
	return s
}

func (s *DescribeHALogsResponseBody) SetItemsNumbers(v int32) *DescribeHALogsResponseBody {
	s.ItemsNumbers = &v
	return s
}

func (s *DescribeHALogsResponseBody) SetPageNumber(v int32) *DescribeHALogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHALogsResponseBody) SetPageSize(v int32) *DescribeHALogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHALogsResponseBody) SetRequestId(v string) *DescribeHALogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHALogsResponseBody) SetTotalRecords(v int32) *DescribeHALogsResponseBody {
	s.TotalRecords = &v
	return s
}

func (s *DescribeHALogsResponseBody) Validate() error {
	if s.HaLogItems != nil {
		for _, item := range s.HaLogItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHALogsResponseBodyHaLogItems struct {
	// The reason code of the failover.
	//
	// example:
	//
	// Platform.Ha.AuroraService.ManualOperations
	SwitchCauseCode *string `json:"SwitchCauseCode,omitempty" xml:"SwitchCauseCode,omitempty"`
	// The reason of the failover.
	//
	// example:
	//
	// Platform.Ha.ManuallyTriggered
	SwitchCauseDetail *string `json:"SwitchCauseDetail,omitempty" xml:"SwitchCauseDetail,omitempty"`
	// The time when the failover ended.
	//
	// example:
	//
	// 2025-05-20T03:09:56Z
	SwitchFinishTime *string `json:"SwitchFinishTime,omitempty" xml:"SwitchFinishTime,omitempty"`
	// example:
	//
	// e571f897-9b3c-4012-9470-88333832dec4
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The time when the failover started.
	//
	// example:
	//
	// 2025-05-20T03:09:45Z
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" xml:"SwitchStartTime,omitempty"`
}

func (s DescribeHALogsResponseBodyHaLogItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeHALogsResponseBodyHaLogItems) GoString() string {
	return s.String()
}

func (s *DescribeHALogsResponseBodyHaLogItems) GetSwitchCauseCode() *string {
	return s.SwitchCauseCode
}

func (s *DescribeHALogsResponseBodyHaLogItems) GetSwitchCauseDetail() *string {
	return s.SwitchCauseDetail
}

func (s *DescribeHALogsResponseBodyHaLogItems) GetSwitchFinishTime() *string {
	return s.SwitchFinishTime
}

func (s *DescribeHALogsResponseBodyHaLogItems) GetSwitchId() *string {
	return s.SwitchId
}

func (s *DescribeHALogsResponseBodyHaLogItems) GetSwitchStartTime() *string {
	return s.SwitchStartTime
}

func (s *DescribeHALogsResponseBodyHaLogItems) SetSwitchCauseCode(v string) *DescribeHALogsResponseBodyHaLogItems {
	s.SwitchCauseCode = &v
	return s
}

func (s *DescribeHALogsResponseBodyHaLogItems) SetSwitchCauseDetail(v string) *DescribeHALogsResponseBodyHaLogItems {
	s.SwitchCauseDetail = &v
	return s
}

func (s *DescribeHALogsResponseBodyHaLogItems) SetSwitchFinishTime(v string) *DescribeHALogsResponseBodyHaLogItems {
	s.SwitchFinishTime = &v
	return s
}

func (s *DescribeHALogsResponseBodyHaLogItems) SetSwitchId(v string) *DescribeHALogsResponseBodyHaLogItems {
	s.SwitchId = &v
	return s
}

func (s *DescribeHALogsResponseBodyHaLogItems) SetSwitchStartTime(v string) *DescribeHALogsResponseBodyHaLogItems {
	s.SwitchStartTime = &v
	return s
}

func (s *DescribeHALogsResponseBodyHaLogItems) Validate() error {
	return dara.Validate(s)
}
