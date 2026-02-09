// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *DescribeDBInstancePlansResponseBody
	GetErrorMessage() *string
	SetItems(v *DescribeDBInstancePlansResponseBodyItems) *DescribeDBInstancePlansResponseBody
	GetItems() *DescribeDBInstancePlansResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstancePlansResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBInstancePlansResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBInstancePlansResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDBInstancePlansResponseBody
	GetStatus() *string
	SetTotalRecordCount(v int32) *DescribeDBInstancePlansResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstancePlansResponseBody struct {
	// The error message.
	//
	// This parameter is returned only if the request fails.
	//
	// example:
	//
	// ****
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Items        *DescribeDBInstancePlansResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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
	// 34b32a0a-08ef-4a87-b6be-cdd988888888
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// If the request was successful, **success*	- is returned. If the request failed, this parameter is not returned.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancePlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDBInstancePlansResponseBody) GetItems() *DescribeDBInstancePlansResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstancePlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancePlansResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBInstancePlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancePlansResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstancePlansResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstancePlansResponseBody) SetErrorMessage(v string) *DescribeDBInstancePlansResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetItems(v *DescribeDBInstancePlansResponseBodyItems) *DescribeDBInstancePlansResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetPageNumber(v int32) *DescribeDBInstancePlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancePlansResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetRequestId(v string) *DescribeDBInstancePlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetStatus(v string) *DescribeDBInstancePlansResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancePlansResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancePlansResponseBodyItems struct {
	PlanList []*DescribeDBInstancePlansResponseBodyItemsPlanList `json:"PlanList,omitempty" xml:"PlanList,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePlansResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBodyItems) GetPlanList() []*DescribeDBInstancePlansResponseBodyItemsPlanList {
	return s.PlanList
}

func (s *DescribeDBInstancePlansResponseBodyItems) SetPlanList(v []*DescribeDBInstancePlansResponseBodyItemsPlanList) *DescribeDBInstancePlansResponseBodyItems {
	s.PlanList = v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItems) Validate() error {
	if s.PlanList != nil {
		for _, item := range s.PlanList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancePlansResponseBodyItemsPlanList struct {
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	PlanConfig       *string `json:"PlanConfig,omitempty" xml:"PlanConfig,omitempty"`
	PlanDesc         *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	PlanEndDate      *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	PlanId           *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName         *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	PlanStartDate    *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
	PlanStatus       *string `json:"PlanStatus,omitempty" xml:"PlanStatus,omitempty"`
	PlanType         *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s DescribeDBInstancePlansResponseBodyItemsPlanList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePlansResponseBodyItemsPlanList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanConfig() *string {
	return s.PlanConfig
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanDesc() *string {
	return s.PlanDesc
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanEndDate() *string {
	return s.PlanEndDate
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanScheduleType() *string {
	return s.PlanScheduleType
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanStartDate() *string {
	return s.PlanStartDate
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanStatus() *string {
	return s.PlanStatus
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) GetPlanType() *string {
	return s.PlanType
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetDBInstanceId(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanConfig(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanConfig = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanDesc(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanDesc = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanEndDate(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanEndDate = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanId(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanId = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanName(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanScheduleType(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanScheduleType = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanStartDate(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanStartDate = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanStatus(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanStatus = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) SetPlanType(v string) *DescribeDBInstancePlansResponseBodyItemsPlanList {
	s.PlanType = &v
	return s
}

func (s *DescribeDBInstancePlansResponseBodyItemsPlanList) Validate() error {
	return dara.Validate(s)
}
