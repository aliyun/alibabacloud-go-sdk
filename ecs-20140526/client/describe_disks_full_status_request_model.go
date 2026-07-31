// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisksFullStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventTime(v *DescribeDisksFullStatusRequestEventTime) *DescribeDisksFullStatusRequest
	GetEventTime() *DescribeDisksFullStatusRequestEventTime
	SetDiskId(v []*string) *DescribeDisksFullStatusRequest
	GetDiskId() []*string
	SetEventId(v []*string) *DescribeDisksFullStatusRequest
	GetEventId() []*string
	SetEventType(v string) *DescribeDisksFullStatusRequest
	GetEventType() *string
	SetHealthStatus(v string) *DescribeDisksFullStatusRequest
	GetHealthStatus() *string
	SetOwnerAccount(v string) *DescribeDisksFullStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDisksFullStatusRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDisksFullStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDisksFullStatusRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDisksFullStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDisksFullStatusRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDisksFullStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDisksFullStatusRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeDisksFullStatusRequest
	GetStatus() *string
	SetTag(v []*DescribeDisksFullStatusRequestTag) *DescribeDisksFullStatusRequest
	GetTag() []*DescribeDisksFullStatusRequestTag
}

type DescribeDisksFullStatusRequest struct {
	EventTime *DescribeDisksFullStatusRequestEventTime `json:"EventTime,omitempty" xml:"EventTime,omitempty" type:"Struct"`
	// The block storage ID. Valid values of N: 1 to 100.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId []*string `json:"DiskId,omitempty" xml:"DiskId,omitempty" type:"Repeated"`
	// The event ID. Valid values of N: 1 to 100.
	//
	// example:
	//
	// e-bp67acfmxazb4p****
	EventId []*string `json:"EventId,omitempty" xml:"EventId,omitempty" type:"Repeated"`
	// The event type of the block storage device. Valid values:
	//
	// - Degraded: The block storage performance is degraded.
	//
	// - SeverelyDegraded: The block storage performance is severely degraded.
	//
	// - Stalled: The block storage performance is severely impacted.
	//
	// - ErrorDetected: A local disk is damaged.
	//
	// example:
	//
	// Stalled
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The health status of the block storage device. Valid values:
	//
	// - Impaired: temporarily unreadable and unwritable.
	//
	// - Warning: degraded service.
	//
	// - Initializing: being initialized.
	//
	// - InsufficientData: insufficient data.
	//
	// - NotApplicable: not applicable.
	//
	// example:
	//
	// Warning
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the query result. Valid values: positive integers.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the block storage device. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the block storage resource belongs. When you use this parameter to filter resources, the resource count cannot exceed 1,000.
	//
	// example:
	//
	// rg-aek2kkmhmhs****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The lifecycle status of the block storage device. For more information, see [Disk status table](https://help.aliyun.com/document_detail/25689.html). Valid values:
	//
	// - In_use: in use.
	//
	// - Available: to be attached.
	//
	// - Attaching: being attached.
	//
	// - Detaching: being detached.
	//
	// - Creating: being created.
	//
	// - ReIniting: being initialized.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*DescribeDisksFullStatusRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDisksFullStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusRequest) GetEventTime() *DescribeDisksFullStatusRequestEventTime {
	return s.EventTime
}

func (s *DescribeDisksFullStatusRequest) GetDiskId() []*string {
	return s.DiskId
}

func (s *DescribeDisksFullStatusRequest) GetEventId() []*string {
	return s.EventId
}

func (s *DescribeDisksFullStatusRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDisksFullStatusRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeDisksFullStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDisksFullStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDisksFullStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDisksFullStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDisksFullStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDisksFullStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDisksFullStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDisksFullStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDisksFullStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDisksFullStatusRequest) GetTag() []*DescribeDisksFullStatusRequestTag {
	return s.Tag
}

func (s *DescribeDisksFullStatusRequest) SetEventTime(v *DescribeDisksFullStatusRequestEventTime) *DescribeDisksFullStatusRequest {
	s.EventTime = v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetDiskId(v []*string) *DescribeDisksFullStatusRequest {
	s.DiskId = v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetEventId(v []*string) *DescribeDisksFullStatusRequest {
	s.EventId = v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetEventType(v string) *DescribeDisksFullStatusRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetHealthStatus(v string) *DescribeDisksFullStatusRequest {
	s.HealthStatus = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetOwnerAccount(v string) *DescribeDisksFullStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetOwnerId(v int64) *DescribeDisksFullStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetPageNumber(v int32) *DescribeDisksFullStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetPageSize(v int32) *DescribeDisksFullStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetRegionId(v string) *DescribeDisksFullStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetResourceGroupId(v string) *DescribeDisksFullStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetResourceOwnerAccount(v string) *DescribeDisksFullStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetResourceOwnerId(v int64) *DescribeDisksFullStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetStatus(v string) *DescribeDisksFullStatusRequest {
	s.Status = &v
	return s
}

func (s *DescribeDisksFullStatusRequest) SetTag(v []*DescribeDisksFullStatusRequestTag) *DescribeDisksFullStatusRequest {
	s.Tag = v
	return s
}

func (s *DescribeDisksFullStatusRequest) Validate() error {
	if s.EventTime != nil {
		if err := s.EventTime.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDisksFullStatusRequestEventTime struct {
	// The end of the time range during which to query events.
	//
	// Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2018-05-08T02:48:52Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start of the time range during which to query events.
	//
	// Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2018-05-06T02:43:10Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeDisksFullStatusRequestEventTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusRequestEventTime) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusRequestEventTime) GetEnd() *string {
	return s.End
}

func (s *DescribeDisksFullStatusRequestEventTime) GetStart() *string {
	return s.Start
}

func (s *DescribeDisksFullStatusRequestEventTime) SetEnd(v string) *DescribeDisksFullStatusRequestEventTime {
	s.End = &v
	return s
}

func (s *DescribeDisksFullStatusRequestEventTime) SetStart(v string) *DescribeDisksFullStatusRequestEventTime {
	s.Start = &v
	return s
}

func (s *DescribeDisksFullStatusRequestEventTime) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksFullStatusRequestTag struct {
	// The tag key attached to the block storage resource. N specifies that you can set one or more tag keys. The value of N in this parameter corresponds to the value of N in the `Tag.N.Value` parameter to form a key-value pair. Valid values of N: 1 to 20.
	//
	// If you use a single tag to filter resources, the number of resources with the tag cannot exceed 1,000. If you use multiple tags to filter resources, the number of resources that are attached to all specified tags cannot exceed 1,000.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value attached to the block storage resource. N specifies that you can set one or more tag values. The value of N in this parameter corresponds to the value of N in the `Tag.N.Key` parameter to form a key-value pair. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDisksFullStatusRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDisksFullStatusRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDisksFullStatusRequestTag) SetKey(v string) *DescribeDisksFullStatusRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDisksFullStatusRequestTag) SetValue(v string) *DescribeDisksFullStatusRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDisksFullStatusRequestTag) Validate() error {
	return dara.Validate(s)
}
