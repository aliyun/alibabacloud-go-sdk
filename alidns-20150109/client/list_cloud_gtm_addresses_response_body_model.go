// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v *ListCloudGtmAddressesResponseBodyAddresses) *ListCloudGtmAddressesResponseBody
	GetAddresses() *ListCloudGtmAddressesResponseBodyAddresses
	SetPageNumber(v int32) *ListCloudGtmAddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmAddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCloudGtmAddressesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListCloudGtmAddressesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListCloudGtmAddressesResponseBody
	GetTotalPages() *int32
}

type ListCloudGtmAddressesResponseBody struct {
	Addresses *ListCloudGtmAddressesResponseBodyAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	// Current page number, starting from **1**, default is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of 100 and a default of 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Get the total number of addresses in the address list.
	//
	// example:
	//
	// 30
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCloudGtmAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressesResponseBody) GetAddresses() *ListCloudGtmAddressesResponseBodyAddresses {
	return s.Addresses
}

func (s *ListCloudGtmAddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmAddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmAddressesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListCloudGtmAddressesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListCloudGtmAddressesResponseBody) SetAddresses(v *ListCloudGtmAddressesResponseBodyAddresses) *ListCloudGtmAddressesResponseBody {
	s.Addresses = v
	return s
}

func (s *ListCloudGtmAddressesResponseBody) SetPageNumber(v int32) *ListCloudGtmAddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBody) SetPageSize(v int32) *ListCloudGtmAddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBody) SetRequestId(v string) *ListCloudGtmAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBody) SetTotalItems(v int32) *ListCloudGtmAddressesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBody) SetTotalPages(v int32) *ListCloudGtmAddressesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBody) Validate() error {
	if s.Addresses != nil {
		if err := s.Addresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmAddressesResponseBodyAddresses struct {
	Address []*ListCloudGtmAddressesResponseBodyAddressesAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s ListCloudGtmAddressesResponseBodyAddresses) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressesResponseBodyAddresses) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressesResponseBodyAddresses) GetAddress() []*ListCloudGtmAddressesResponseBodyAddressesAddress {
	return s.Address
}

func (s *ListCloudGtmAddressesResponseBodyAddresses) SetAddress(v []*ListCloudGtmAddressesResponseBodyAddressesAddress) *ListCloudGtmAddressesResponseBodyAddresses {
	s.Address = v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddresses) Validate() error {
	if s.Address != nil {
		for _, item := range s.Address {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudGtmAddressesResponseBodyAddressesAddress struct {
	Address               *string                                                       `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressId             *string                                                       `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	AttributeInfo         *string                                                       `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	AvailableMode         *string                                                       `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	AvailableStatus       *string                                                       `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime            *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp       *int64                                                        `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus          *string                                                       `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement       *string                                                       `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus          *string                                                       `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	HealthTasks           *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	ManualAvailableStatus *string                                                       `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	Name                  *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark                *string                                                       `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Type                  *string                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime            *string                                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp       *int64                                                        `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s ListCloudGtmAddressesResponseBodyAddressesAddress) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressesResponseBodyAddressesAddress) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetAddress() *string {
	return s.Address
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetAddressId() *string {
	return s.AddressId
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetHealthTasks() *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks {
	return s.HealthTasks
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetName() *string {
	return s.Name
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetRemark() *string {
	return s.Remark
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetType() *string {
	return s.Type
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetAddress(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.Address = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetAddressId(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.AddressId = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetAttributeInfo(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.AttributeInfo = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetAvailableMode(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.AvailableMode = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetAvailableStatus(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.AvailableStatus = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetCreateTime(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.CreateTime = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetCreateTimestamp(v int64) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetEnableStatus(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.EnableStatus = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetHealthJudgement(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthJudgement = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetHealthStatus(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthStatus = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetHealthTasks(v *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthTasks = v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetManualAvailableStatus(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.ManualAvailableStatus = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetName(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.Name = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetRemark(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.Remark = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetType(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.Type = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetUpdateTime(v string) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) SetUpdateTimestamp(v int64) *ListCloudGtmAddressesResponseBodyAddressesAddress {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddress) Validate() error {
	if s.HealthTasks != nil {
		if err := s.HealthTasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks struct {
	HealthTask []*ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask `json:"HealthTask,omitempty" xml:"HealthTask,omitempty" type:"Repeated"`
}

func (s ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) GetHealthTask() []*ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	return s.HealthTask
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) SetHealthTask(v []*ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks {
	s.HealthTask = v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) Validate() error {
	if s.HealthTask != nil {
		for _, item := range s.HealthTask {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask struct {
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName  *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetPort() *int32 {
	return s.Port
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetMonitorStatus(v string) *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.MonitorStatus = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetPort(v int32) *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.Port = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetTemplateId(v string) *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.TemplateId = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetTemplateName(v string) *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.TemplateName = &v
	return s
}

func (s *ListCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) Validate() error {
	return dara.Validate(s)
}
