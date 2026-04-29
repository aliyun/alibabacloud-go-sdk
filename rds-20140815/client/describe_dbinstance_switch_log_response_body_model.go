// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSwitchLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceSwitchLogResponseBody
	GetDBInstanceName() *string
	SetItems(v *DescribeDBInstanceSwitchLogResponseBodyItems) *DescribeDBInstanceSwitchLogResponseBody
	GetItems() *DescribeDBInstanceSwitchLogResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstanceSwitchLogResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBInstanceSwitchLogResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBInstanceSwitchLogResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBInstanceSwitchLogResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstanceSwitchLogResponseBody struct {
	// example:
	//
	// rdsaiiabnaiiabn
	DBInstanceName *string                                       `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	Items          *DescribeDBInstanceSwitchLogResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 60
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// D1CA494F-CC13-4EB6-8C4D-5352EE4045BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstanceSwitchLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSwitchLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetItems() *DescribeDBInstanceSwitchLogResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSwitchLogResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetDBInstanceName(v string) *DescribeDBInstanceSwitchLogResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetItems(v *DescribeDBInstanceSwitchLogResponseBodyItems) *DescribeDBInstanceSwitchLogResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetPageNumber(v int32) *DescribeDBInstanceSwitchLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetPageRecordCount(v int32) *DescribeDBInstanceSwitchLogResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetRequestId(v string) *DescribeDBInstanceSwitchLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstanceSwitchLogResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceSwitchLogResponseBodyItems struct {
	Item []*DescribeDBInstanceSwitchLogResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceSwitchLogResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSwitchLogResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItems) GetItem() []*DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	return s.Item
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItems) SetItem(v []*DescribeDBInstanceSwitchLogResponseBodyItemsItem) *DescribeDBInstanceSwitchLogResponseBodyItems {
	s.Item = v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceSwitchLogResponseBodyItemsItem struct {
	AffectedSessions  *string `json:"AffectedSessions,omitempty" xml:"AffectedSessions,omitempty"`
	HAStatus          *string `json:"HAStatus,omitempty" xml:"HAStatus,omitempty"`
	SwitchCauseCode   *string `json:"SwitchCauseCode,omitempty" xml:"SwitchCauseCode,omitempty"`
	SwitchCauseDetail *string `json:"SwitchCauseDetail,omitempty" xml:"SwitchCauseDetail,omitempty"`
	SwitchFinishTime  *string `json:"SwitchFinishTime,omitempty" xml:"SwitchFinishTime,omitempty"`
	SwitchId          *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	SwitchStartTime   *string `json:"SwitchStartTime,omitempty" xml:"SwitchStartTime,omitempty"`
	TotalSessions     *string `json:"TotalSessions,omitempty" xml:"TotalSessions,omitempty"`
}

func (s DescribeDBInstanceSwitchLogResponseBodyItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSwitchLogResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) GetAffectedSessions() *string {
	return s.AffectedSessions
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) GetHAStatus() *string {
	return s.HAStatus
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) GetSwitchCauseCode() *string {
	return s.SwitchCauseCode
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) GetSwitchCauseDetail() *string {
	return s.SwitchCauseDetail
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) GetSwitchFinishTime() *string {
	return s.SwitchFinishTime
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) GetSwitchId() *string {
	return s.SwitchId
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) GetSwitchStartTime() *string {
	return s.SwitchStartTime
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) GetTotalSessions() *string {
	return s.TotalSessions
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) SetAffectedSessions(v string) *DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	s.AffectedSessions = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) SetHAStatus(v string) *DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	s.HAStatus = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) SetSwitchCauseCode(v string) *DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	s.SwitchCauseCode = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) SetSwitchCauseDetail(v string) *DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	s.SwitchCauseDetail = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) SetSwitchFinishTime(v string) *DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	s.SwitchFinishTime = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) SetSwitchId(v string) *DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	s.SwitchId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) SetSwitchStartTime(v string) *DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	s.SwitchStartTime = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) SetTotalSessions(v string) *DescribeDBInstanceSwitchLogResponseBodyItemsItem {
	s.TotalSessions = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogResponseBodyItemsItem) Validate() error {
	return dara.Validate(s)
}
