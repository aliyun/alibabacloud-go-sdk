// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePendingMaintenanceActionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePendingMaintenanceActionsResponseBody
	GetRequestId() *string
	SetTypeList(v []*DescribePendingMaintenanceActionsResponseBodyTypeList) *DescribePendingMaintenanceActionsResponseBody
	GetTypeList() []*DescribePendingMaintenanceActionsResponseBodyTypeList
}

type DescribePendingMaintenanceActionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 314127C2-B6C1-4F58-B1F6-E6B645******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of pending events.
	TypeList []*DescribePendingMaintenanceActionsResponseBodyTypeList `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s DescribePendingMaintenanceActionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePendingMaintenanceActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePendingMaintenanceActionsResponseBody) GetTypeList() []*DescribePendingMaintenanceActionsResponseBodyTypeList {
	return s.TypeList
}

func (s *DescribePendingMaintenanceActionsResponseBody) SetRequestId(v string) *DescribePendingMaintenanceActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsResponseBody) SetTypeList(v []*DescribePendingMaintenanceActionsResponseBodyTypeList) *DescribePendingMaintenanceActionsResponseBody {
	s.TypeList = v
	return s
}

func (s *DescribePendingMaintenanceActionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePendingMaintenanceActionsResponseBodyTypeList struct {
	// The number of pending events.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The task type of pending events. Valid values:
	//
	// 	- **DatabaseSoftwareUpgrading**: database software upgrades
	//
	// 	- **DatabaseHardwareMaintenance**: hardware maintenance and upgrades
	//
	// 	- **DatabaseStorageUpgrading**: database storage upgrades
	//
	// 	- **DatabaseProxyUpgrading**: minor version upgrades of the proxy
	//
	// example:
	//
	// DatabaseSoftwareUpgrading
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribePendingMaintenanceActionsResponseBodyTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribePendingMaintenanceActionsResponseBodyTypeList) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionsResponseBodyTypeList) GetCount() *int32 {
	return s.Count
}

func (s *DescribePendingMaintenanceActionsResponseBodyTypeList) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribePendingMaintenanceActionsResponseBodyTypeList) SetCount(v int32) *DescribePendingMaintenanceActionsResponseBodyTypeList {
	s.Count = &v
	return s
}

func (s *DescribePendingMaintenanceActionsResponseBodyTypeList) SetTaskType(v string) *DescribePendingMaintenanceActionsResponseBodyTypeList {
	s.TaskType = &v
	return s
}

func (s *DescribePendingMaintenanceActionsResponseBodyTypeList) Validate() error {
	return dara.Validate(s)
}
