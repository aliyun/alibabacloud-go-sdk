// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataShareInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDataShareInstancesResponseBodyItems) *DescribeDataShareInstancesResponseBody
	GetItems() *DescribeDataShareInstancesResponseBodyItems
	SetPageNumber(v int32) *DescribeDataShareInstancesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDataShareInstancesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDataShareInstancesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDataShareInstancesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDataShareInstancesResponseBody struct {
	Items *DescribeDataShareInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D5**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDataShareInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBody) GetItems() *DescribeDataShareInstancesResponseBodyItems {
	return s.Items
}

func (s *DescribeDataShareInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataShareInstancesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDataShareInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataShareInstancesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDataShareInstancesResponseBody) SetItems(v *DescribeDataShareInstancesResponseBodyItems) *DescribeDataShareInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetPageNumber(v int32) *DescribeDataShareInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetPageRecordCount(v int32) *DescribeDataShareInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetRequestId(v string) *DescribeDataShareInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeDataShareInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataShareInstancesResponseBodyItems struct {
	DBInstance []*DescribeDataShareInstancesResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDataShareInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBodyItems) GetDBInstance() []*DescribeDataShareInstancesResponseBodyItemsDBInstance {
	return s.DBInstance
}

func (s *DescribeDataShareInstancesResponseBodyItems) SetDBInstance(v []*DescribeDataShareInstancesResponseBodyItemsDBInstance) *DescribeDataShareInstancesResponseBodyItems {
	s.DBInstance = v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItems) Validate() error {
	if s.DBInstance != nil {
		for _, item := range s.DBInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataShareInstancesResponseBodyItemsDBInstance struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceMode  *string `json:"DBInstanceMode,omitempty" xml:"DBInstanceMode,omitempty"`
	DataShareStatus *string `json:"DataShareStatus,omitempty" xml:"DataShareStatus,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDataShareInstancesResponseBodyItemsDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataShareInstancesResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) GetDBInstanceMode() *string {
	return s.DBInstanceMode
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) GetDataShareStatus() *string {
	return s.DataShareStatus
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) GetDescription() *string {
	return s.Description
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDBInstanceMode(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DBInstanceMode = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDataShareStatus(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.DataShareStatus = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetDescription(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.Description = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDataShareInstancesResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDataShareInstancesResponseBodyItemsDBInstance) Validate() error {
	return dara.Validate(s)
}
