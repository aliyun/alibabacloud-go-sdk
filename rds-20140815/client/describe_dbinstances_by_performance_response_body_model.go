// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesByPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBInstancesByPerformanceResponseBodyItems) *DescribeDBInstancesByPerformanceResponseBody
	GetItems() *DescribeDBInstancesByPerformanceResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstancesByPerformanceResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBInstancesByPerformanceResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBInstancesByPerformanceResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBInstancesByPerformanceResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstancesByPerformanceResponseBody struct {
	Items *DescribeDBInstancesByPerformanceResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 28
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 23907437-79B9-411A-9EE6-75A8F0F1C619
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 28
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancesByPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByPerformanceResponseBody) GetItems() *DescribeDBInstancesByPerformanceResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstancesByPerformanceResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesByPerformanceResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBInstancesByPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesByPerformanceResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstancesByPerformanceResponseBody) SetItems(v *DescribeDBInstancesByPerformanceResponseBodyItems) *DescribeDBInstancesByPerformanceResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBody) SetPageNumber(v int32) *DescribeDBInstancesByPerformanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancesByPerformanceResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBody) SetRequestId(v string) *DescribeDBInstancesByPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancesByPerformanceResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesByPerformanceResponseBodyItems struct {
	DBInstancePerformance []*DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance `json:"DBInstancePerformance,omitempty" xml:"DBInstancePerformance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesByPerformanceResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByPerformanceResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItems) GetDBInstancePerformance() []*DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance {
	return s.DBInstancePerformance
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItems) SetDBInstancePerformance(v []*DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) *DescribeDBInstancesByPerformanceResponseBodyItems {
	s.DBInstancePerformance = v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItems) Validate() error {
	if s.DBInstancePerformance != nil {
		for _, item := range s.DBInstancePerformance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance struct {
	CPUUsage              *string `json:"CPUUsage,omitempty" xml:"CPUUsage,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DiskUsage             *string `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	IOPSUsage             *string `json:"IOPSUsage,omitempty" xml:"IOPSUsage,omitempty"`
	SessionUsage          *string `json:"SessionUsage,omitempty" xml:"SessionUsage,omitempty"`
}

func (s DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) GetCPUUsage() *string {
	return s.CPUUsage
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) GetDiskUsage() *string {
	return s.DiskUsage
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) GetIOPSUsage() *string {
	return s.IOPSUsage
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) GetSessionUsage() *string {
	return s.SessionUsage
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) SetCPUUsage(v string) *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance {
	s.CPUUsage = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) SetDBInstanceDescription(v string) *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) SetDBInstanceId(v string) *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) SetDiskUsage(v string) *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance {
	s.DiskUsage = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) SetIOPSUsage(v string) *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance {
	s.IOPSUsage = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) SetSessionUsage(v string) *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance {
	s.SessionUsage = &v
	return s
}

func (s *DescribeDBInstancesByPerformanceResponseBodyItemsDBInstancePerformance) Validate() error {
	return dara.Validate(s)
}
