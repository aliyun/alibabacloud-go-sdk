// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v *SearchCloudGtmAddressesResponseBodyAddresses) *SearchCloudGtmAddressesResponseBody
	GetAddresses() *SearchCloudGtmAddressesResponseBodyAddresses
	SetPageNumber(v int32) *SearchCloudGtmAddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmAddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchCloudGtmAddressesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchCloudGtmAddressesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchCloudGtmAddressesResponseBody
	GetTotalPages() *int32
}

type SearchCloudGtmAddressesResponseBody struct {
	Addresses *SearchCloudGtmAddressesResponseBodyAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	// Current page number, starting from **1**, default is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of **100*	- and a default of **20**.
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
	// Total number of address entries that meet the query conditions.
	//
	// example:
	//
	// 15
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s SearchCloudGtmAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBody) GetAddresses() *SearchCloudGtmAddressesResponseBodyAddresses {
	return s.Addresses
}

func (s *SearchCloudGtmAddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmAddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchCloudGtmAddressesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchCloudGtmAddressesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchCloudGtmAddressesResponseBody) SetAddresses(v *SearchCloudGtmAddressesResponseBodyAddresses) *SearchCloudGtmAddressesResponseBody {
	s.Addresses = v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetPageNumber(v int32) *SearchCloudGtmAddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetPageSize(v int32) *SearchCloudGtmAddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetRequestId(v string) *SearchCloudGtmAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetTotalItems(v int32) *SearchCloudGtmAddressesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetTotalPages(v int32) *SearchCloudGtmAddressesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) Validate() error {
	if s.Addresses != nil {
		if err := s.Addresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCloudGtmAddressesResponseBodyAddresses struct {
	Address []*SearchCloudGtmAddressesResponseBodyAddressesAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmAddressesResponseBodyAddresses) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBodyAddresses) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBodyAddresses) GetAddress() []*SearchCloudGtmAddressesResponseBodyAddressesAddress {
	return s.Address
}

func (s *SearchCloudGtmAddressesResponseBodyAddresses) SetAddress(v []*SearchCloudGtmAddressesResponseBodyAddressesAddress) *SearchCloudGtmAddressesResponseBodyAddresses {
	s.Address = v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddresses) Validate() error {
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

type SearchCloudGtmAddressesResponseBodyAddressesAddress struct {
	Address               *string                                                         `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressId             *string                                                         `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	AttributeInfo         *string                                                         `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	AvailableMode         *string                                                         `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	AvailableStatus       *string                                                         `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime            *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp       *int64                                                          `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus          *string                                                         `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement       *string                                                         `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus          *string                                                         `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	HealthTasks           *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	ManualAvailableStatus *string                                                         `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	Name                  *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark                *string                                                         `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Type                  *string                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime            *string                                                         `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp       *int64                                                          `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddress) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddress) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAddress() *string {
	return s.Address
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAddressId() *string {
	return s.AddressId
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetHealthTasks() *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks {
	return s.HealthTasks
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetName() *string {
	return s.Name
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetRemark() *string {
	return s.Remark
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetType() *string {
	return s.Type
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAddress(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.Address = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAddressId(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.AddressId = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAttributeInfo(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.AttributeInfo = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAvailableMode(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.AvailableMode = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAvailableStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetCreateTime(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.CreateTime = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetCreateTimestamp(v int64) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetEnableStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetHealthJudgement(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthJudgement = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetHealthStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetHealthTasks(v *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthTasks = v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetManualAvailableStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.ManualAvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetName(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.Name = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetRemark(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.Remark = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetType(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.Type = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetUpdateTime(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.UpdateTime = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetUpdateTimestamp(v int64) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) Validate() error {
	if s.HealthTasks != nil {
		if err := s.HealthTasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks struct {
	HealthTask []*SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask `json:"HealthTask,omitempty" xml:"HealthTask,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) GetHealthTask() []*SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	return s.HealthTask
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) SetHealthTask(v []*SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks {
	s.HealthTask = v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) Validate() error {
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

type SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask struct {
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName  *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetPort() *int32 {
	return s.Port
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetMonitorStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.MonitorStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetPort(v int32) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.Port = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetTemplateId(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.TemplateId = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetTemplateName(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.TemplateName = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) Validate() error {
	return dara.Validate(s)
}
