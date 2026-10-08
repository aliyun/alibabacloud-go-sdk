// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceHistoryEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceSystemEventSet(v []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) *DescribeRCInstanceHistoryEventsResponseBody
	GetInstanceSystemEventSet() []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet
	SetNextToken(v string) *DescribeRCInstanceHistoryEventsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeRCInstanceHistoryEventsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCInstanceHistoryEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRCInstanceHistoryEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRCInstanceHistoryEventsResponseBody
	GetTotalCount() *int32
}

type DescribeRCInstanceHistoryEventsResponseBody struct {
	InstanceSystemEventSet []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet `json:"InstanceSystemEventSet,omitempty" xml:"InstanceSystemEventSet,omitempty" type:"Repeated"`
	NextToken              *string                                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNumber             *int32                                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *int32                                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) GetInstanceSystemEventSet() []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	return s.InstanceSystemEventSet
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) SetInstanceSystemEventSet(v []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) *DescribeRCInstanceHistoryEventsResponseBody {
	s.InstanceSystemEventSet = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) SetNextToken(v string) *DescribeRCInstanceHistoryEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) SetPageNumber(v int32) *DescribeRCInstanceHistoryEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) SetPageSize(v int32) *DescribeRCInstanceHistoryEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) SetRequestId(v string) *DescribeRCInstanceHistoryEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) SetTotalCount(v int32) *DescribeRCInstanceHistoryEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBody) Validate() error {
	if s.InstanceSystemEventSet != nil {
		for _, item := range s.InstanceSystemEventSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet struct {
	EventCycleStatus  *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus  `json:"EventCycleStatus,omitempty" xml:"EventCycleStatus,omitempty" type:"Struct"`
	EventFinishTime   *string                                                                             `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	EventId           *string                                                                             `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventPublishTime  *string                                                                             `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty"`
	EventType         *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType         `json:"EventType,omitempty" xml:"EventType,omitempty" type:"Struct"`
	ExtendedAttribute *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute `json:"ExtendedAttribute,omitempty" xml:"ExtendedAttribute,omitempty" type:"Struct"`
	ImpactLevel       *string                                                                             `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
	InstanceId        *string                                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NotBefore         *string                                                                             `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	Reason            *string                                                                             `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ReasonCode        *string                                                                             `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ResourceType      *string                                                                             `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetEventCycleStatus() *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus {
	return s.EventCycleStatus
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetEventFinishTime() *string {
	return s.EventFinishTime
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetEventId() *string {
	return s.EventId
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetEventPublishTime() *string {
	return s.EventPublishTime
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetEventType() *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType {
	return s.EventType
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetExtendedAttribute() *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	return s.ExtendedAttribute
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetImpactLevel() *string {
	return s.ImpactLevel
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetNotBefore() *string {
	return s.NotBefore
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetReason() *string {
	return s.Reason
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetEventCycleStatus(v *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.EventCycleStatus = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetEventFinishTime(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.EventFinishTime = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetEventId(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.EventId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetEventPublishTime(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.EventPublishTime = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetEventType(v *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.EventType = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetExtendedAttribute(v *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.ExtendedAttribute = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetImpactLevel(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.ImpactLevel = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetInstanceId(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetNotBefore(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.NotBefore = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetReason(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.Reason = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetReasonCode(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.ReasonCode = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) SetResourceType(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet {
	s.ResourceType = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet) Validate() error {
	if s.EventCycleStatus != nil {
		if err := s.EventCycleStatus.Validate(); err != nil {
			return err
		}
	}
	if s.EventType != nil {
		if err := s.EventType.Validate(); err != nil {
			return err
		}
	}
	if s.ExtendedAttribute != nil {
		if err := s.ExtendedAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus) GetCode() *string {
	return s.Code
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus) GetName() *string {
	return s.Name
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus) SetCode(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus {
	s.Code = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus) SetName(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus {
	s.Name = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType) GetCode() *string {
	return s.Code
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType) GetName() *string {
	return s.Name
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType) SetCode(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType {
	s.Code = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType) SetName(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType {
	s.Name = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute struct {
	CanAccept          *string                                                                                            `json:"CanAccept,omitempty" xml:"CanAccept,omitempty"`
	Code               *string                                                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Device             *string                                                                                            `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskId             *string                                                                                            `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	HostId             *string                                                                                            `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostType           *string                                                                                            `json:"HostType,omitempty" xml:"HostType,omitempty"`
	InactiveDisks      []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks `json:"InactiveDisks,omitempty" xml:"InactiveDisks,omitempty" type:"Repeated"`
	MigrationOptions   []*string                                                                                          `json:"MigrationOptions,omitempty" xml:"MigrationOptions,omitempty" type:"Repeated"`
	OnlineRepairPolicy *string                                                                                            `json:"OnlineRepairPolicy,omitempty" xml:"OnlineRepairPolicy,omitempty"`
	PunishDomain       *string                                                                                            `json:"PunishDomain,omitempty" xml:"PunishDomain,omitempty"`
	PunishType         *string                                                                                            `json:"PunishType,omitempty" xml:"PunishType,omitempty"`
	PunishUrl          *string                                                                                            `json:"PunishUrl,omitempty" xml:"PunishUrl,omitempty"`
	Rack               *string                                                                                            `json:"Rack,omitempty" xml:"Rack,omitempty"`
	ResponseResult     *string                                                                                            `json:"ResponseResult,omitempty" xml:"ResponseResult,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetCanAccept() *string {
	return s.CanAccept
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetCode() *string {
	return s.Code
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetDevice() *string {
	return s.Device
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetHostId() *string {
	return s.HostId
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetHostType() *string {
	return s.HostType
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetInactiveDisks() []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks {
	return s.InactiveDisks
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetMigrationOptions() []*string {
	return s.MigrationOptions
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetOnlineRepairPolicy() *string {
	return s.OnlineRepairPolicy
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetPunishDomain() *string {
	return s.PunishDomain
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetPunishType() *string {
	return s.PunishType
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetPunishUrl() *string {
	return s.PunishUrl
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetRack() *string {
	return s.Rack
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) GetResponseResult() *string {
	return s.ResponseResult
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetCanAccept(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.CanAccept = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetCode(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.Code = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetDevice(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.Device = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetDiskId(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.DiskId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetHostId(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.HostId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetHostType(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.HostType = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetInactiveDisks(v []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.InactiveDisks = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetMigrationOptions(v []*string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.MigrationOptions = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetOnlineRepairPolicy(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.OnlineRepairPolicy = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetPunishDomain(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.PunishDomain = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetPunishType(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.PunishType = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetPunishUrl(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.PunishUrl = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetRack(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.Rack = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) SetResponseResult(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute {
	s.ResponseResult = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute) Validate() error {
	if s.InactiveDisks != nil {
		for _, item := range s.InactiveDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeviceCategory *string `json:"DeviceCategory,omitempty" xml:"DeviceCategory,omitempty"`
	DeviceSize     *string `json:"DeviceSize,omitempty" xml:"DeviceSize,omitempty"`
	DeviceType     *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ReleaseTime    *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) GetDeviceCategory() *string {
	return s.DeviceCategory
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) GetDeviceSize() *string {
	return s.DeviceSize
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) SetCreationTime(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks {
	s.CreationTime = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) SetDeviceCategory(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks {
	s.DeviceCategory = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) SetDeviceSize(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks {
	s.DeviceSize = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) SetDeviceType(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks {
	s.DeviceType = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) SetReleaseTime(v string) *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks) Validate() error {
	return dara.Validate(s)
}
