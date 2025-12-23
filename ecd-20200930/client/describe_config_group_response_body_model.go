// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeConfigGroupResponseBodyData) *DescribeConfigGroupResponseBody
	GetData() []*DescribeConfigGroupResponseBodyData
	SetPageNumber(v int32) *DescribeConfigGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeConfigGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeConfigGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeConfigGroupResponseBody
	GetTotalCount() *int32
}

type DescribeConfigGroupResponseBody struct {
	// The configuration groups.
	Data []*DescribeConfigGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeConfigGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigGroupResponseBody) GetData() []*DescribeConfigGroupResponseBodyData {
	return s.Data
}

func (s *DescribeConfigGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConfigGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConfigGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfigGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeConfigGroupResponseBody) SetData(v []*DescribeConfigGroupResponseBodyData) *DescribeConfigGroupResponseBody {
	s.Data = v
	return s
}

func (s *DescribeConfigGroupResponseBody) SetPageNumber(v int32) *DescribeConfigGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConfigGroupResponseBody) SetPageSize(v int32) *DescribeConfigGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeConfigGroupResponseBody) SetRequestId(v string) *DescribeConfigGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigGroupResponseBody) SetTotalCount(v int32) *DescribeConfigGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeConfigGroupResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConfigGroupResponseBodyData struct {
	// The number of resources that are bound to the configuration group.
	//
	// example:
	//
	// 4
	BindCount *int32 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// The number of bound cloud computers.
	BindCountMap map[string]*int32 `json:"BindCountMap,omitempty" xml:"BindCountMap,omitempty"`
	// The description of the configuration group.
	//
	// example:
	//
	// ScheduledTask
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the configuration group.
	//
	// example:
	//
	// ccg-0cid8v30an12****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// INNER_TIMER_10_MINUTES_HIBERNATE_NO_UPDATE_DESC
	InnerTimerDesc *string `json:"InnerTimerDesc,omitempty" xml:"InnerTimerDesc,omitempty"`
	// example:
	//
	// INNER_TIMER_10_MINUTES_HIBERNATE_NO_UPDATE
	InnerTimerName *string `json:"InnerTimerName,omitempty" xml:"InnerTimerName,omitempty"`
	IsBind         *bool   `json:"IsBind,omitempty" xml:"IsBind,omitempty"`
	IsUpdate       *bool   `json:"IsUpdate,omitempty" xml:"IsUpdate,omitempty"`
	// The name of the configuration group.
	//
	// example:
	//
	// ScheduledTask
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service type of the configuration group.
	//
	// Valid values:
	//
	// 	- CLOUD_DESKTOP: the cloud computer service.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The state of the configuration group.
	//
	// Valid values:
	//
	// 	- AVAILABLE: The configuration group is available.
	//
	// 	- UNAVAILABLE: The configuration group is deleted.
	//
	// 	- DELETING: The configuration group is being deleted.
	//
	// 	- UPDATING: The configuration group is being modified.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the configuration group.
	//
	// Valid values:
	//
	// 	- Timer: the scheduled task type.
	//
	// example:
	//
	// Timer
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeConfigGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeConfigGroupResponseBodyData) GetBindCount() *int32 {
	return s.BindCount
}

func (s *DescribeConfigGroupResponseBodyData) GetBindCountMap() map[string]*int32 {
	return s.BindCountMap
}

func (s *DescribeConfigGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeConfigGroupResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeConfigGroupResponseBodyData) GetInnerTimerDesc() *string {
	return s.InnerTimerDesc
}

func (s *DescribeConfigGroupResponseBodyData) GetInnerTimerName() *string {
	return s.InnerTimerName
}

func (s *DescribeConfigGroupResponseBodyData) GetIsBind() *bool {
	return s.IsBind
}

func (s *DescribeConfigGroupResponseBodyData) GetIsUpdate() *bool {
	return s.IsUpdate
}

func (s *DescribeConfigGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeConfigGroupResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeConfigGroupResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeConfigGroupResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeConfigGroupResponseBodyData) SetBindCount(v int32) *DescribeConfigGroupResponseBodyData {
	s.BindCount = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetBindCountMap(v map[string]*int32) *DescribeConfigGroupResponseBodyData {
	s.BindCountMap = v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetDescription(v string) *DescribeConfigGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetGroupId(v string) *DescribeConfigGroupResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetInnerTimerDesc(v string) *DescribeConfigGroupResponseBodyData {
	s.InnerTimerDesc = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetInnerTimerName(v string) *DescribeConfigGroupResponseBodyData {
	s.InnerTimerName = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetIsBind(v bool) *DescribeConfigGroupResponseBodyData {
	s.IsBind = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetIsUpdate(v bool) *DescribeConfigGroupResponseBodyData {
	s.IsUpdate = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetName(v string) *DescribeConfigGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetProductType(v string) *DescribeConfigGroupResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetStatus(v string) *DescribeConfigGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) SetType(v string) *DescribeConfigGroupResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeConfigGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
