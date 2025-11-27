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
	// Details about the instance system event.
	InstanceSystemEventSet []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSet `json:"InstanceSystemEventSet,omitempty" xml:"InstanceSystemEventSet,omitempty" type:"Repeated"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 866F5EB8-4650-4061-87F0-379F6F968BCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instance events.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The lifecycle state of the system event.
	EventCycleStatus *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventCycleStatus `json:"EventCycleStatus,omitempty" xml:"EventCycleStatus,omitempty" type:"Struct"`
	// The time when the system event ended. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-04-01T06:32:31Z
	EventFinishTime *string `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	// The ID of the system event.
	//
	// example:
	//
	// e-uf64yvznlao4jl2c****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time when the system event was published. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-03-30T06:32:31Z
	EventPublishTime *string `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty"`
	// The type of the system event.
	EventType *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetEventType `json:"EventType,omitempty" xml:"EventType,omitempty" type:"Struct"`
	// The extended attribute of the system event.
	ExtendedAttribute *DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttribute `json:"ExtendedAttribute,omitempty" xml:"ExtendedAttribute,omitempty" type:"Struct"`
	// The impact level of the event.
	//
	// example:
	//
	// 100
	ImpactLevel *string `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-yuf59nplc45t2tzn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The start time of the scheduled execution of the system event. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-04-01T06:32:31Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The reason why the system event occurred.
	//
	// example:
	//
	// System maintenance is scheduled due to ***.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The reason code category for the system event.
	//
	// example:
	//
	// VPCMigrationEcs
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The resource type. The value is fixed to INSTANCE.
	//
	// example:
	//
	// custom
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	// The state code of the system event.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The state name of the system event.
	//
	// example:
	//
	// Executed
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
	// The code of the system event type.
	//
	// example:
	//
	// 34
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the system event type.
	//
	// example:
	//
	// InstanceExpiration.Stop
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
	// Indicates whether the event can be handled.
	//
	// example:
	//
	// true
	CanAccept *string `json:"CanAccept,omitempty" xml:"CanAccept,omitempty"`
	// The code of the security violation.
	//
	// example:
	//
	// PR111
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The device name of the local disk.
	//
	// example:
	//
	// /dev/vda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The ID of the local disk.
	//
	// example:
	//
	// rcd-****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the host.
	//
	// example:
	//
	// dh-bp1ewce1gk3iwv2****
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The type of the host. Valid values:
	//
	// 	- **ddh**: dedicated host
	//
	// 	- **managehost**: physical machine in a smart hosting pool
	//
	// example:
	//
	// ddh
	HostType *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	// The inactive disks that have been released and whose data must be cleared.
	InactiveDisks []*DescribeRCInstanceHistoryEventsResponseBodyInstanceSystemEventSetExtendedAttributeInactiveDisks `json:"InactiveDisks,omitempty" xml:"InactiveDisks,omitempty" type:"Repeated"`
	// The migration solutions of the instance.
	MigrationOptions []*string `json:"MigrationOptions,omitempty" xml:"MigrationOptions,omitempty" type:"Repeated"`
	// The online repair policy for the damaged disk. Valid value: IsolateOnly, which indicates that damaged disks are isolated but not repaired.
	//
	// example:
	//
	// IsolateOnly
	OnlineRepairPolicy *string `json:"OnlineRepairPolicy,omitempty" xml:"OnlineRepairPolicy,omitempty"`
	// The illegal domain name.
	//
	// example:
	//
	// 1228.test.com
	PunishDomain *string `json:"PunishDomain,omitempty" xml:"PunishDomain,omitempty"`
	// The type of the penalty.
	//
	// example:
	//
	// ecs_message_alert
	PunishType *string `json:"PunishType,omitempty" xml:"PunishType,omitempty"`
	// The illegal URL.
	//
	// example:
	//
	// http://1228.test.com/1
	PunishUrl *string `json:"PunishUrl,omitempty" xml:"PunishUrl,omitempty"`
	// The rack number of the cloud box.
	//
	// example:
	//
	// A01
	Rack *string `json:"Rack,omitempty" xml:"Rack,omitempty"`
	// The response result of the event. Valid values:
	//
	// 	- **true**: the event was handled.
	//
	// 	- **false**: the event failed to be handled.
	//
	// example:
	//
	// true
	ResponseResult *string `json:"ResponseResult,omitempty" xml:"ResponseResult,omitempty"`
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
	// The time when the disk was created. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-03-26T03:33:56Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The category of the cloud disk or local disk. Valid values:
	//
	// 	- **cloud_efficiency**: ultra disk
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: ESSD
	//
	// 	- **cloud_auto**: Premium ESSD
	//
	// example:
	//
	// cloud_auto
	DeviceCategory *string `json:"DeviceCategory,omitempty" xml:"DeviceCategory,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 40
	DeviceSize *string `json:"DeviceSize,omitempty" xml:"DeviceSize,omitempty"`
	// The disk type. Valid values:
	//
	// 	- **system**: system disk.
	//
	// 	- **data**: data disk.
	//
	// example:
	//
	// data
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The time when the disk was released. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-03-26T03:33:56Z
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
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
