// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSystemEventAttributeResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeSystemEventAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSystemEventAttributeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSystemEventAttributeResponseBody
	GetSuccess() *string
	SetSystemEvents(v *DescribeSystemEventAttributeResponseBodySystemEvents) *DescribeSystemEventAttributeResponseBody
	GetSystemEvents() *DescribeSystemEventAttributeResponseBodySystemEvents
}

type DescribeSystemEventAttributeResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, `success` is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 60912C8D-B340-4253-ADE7-61ACDFD25CFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values: True: The call is successful. false: The call fails.
	//
	// example:
	//
	// true
	Success      *string                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	SystemEvents *DescribeSystemEventAttributeResponseBodySystemEvents `json:"SystemEvents,omitempty" xml:"SystemEvents,omitempty" type:"Struct"`
}

func (s DescribeSystemEventAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSystemEventAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSystemEventAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemEventAttributeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSystemEventAttributeResponseBody) GetSystemEvents() *DescribeSystemEventAttributeResponseBodySystemEvents {
	return s.SystemEvents
}

func (s *DescribeSystemEventAttributeResponseBody) SetCode(v string) *DescribeSystemEventAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) SetMessage(v string) *DescribeSystemEventAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) SetRequestId(v string) *DescribeSystemEventAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) SetSuccess(v string) *DescribeSystemEventAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) SetSystemEvents(v *DescribeSystemEventAttributeResponseBodySystemEvents) *DescribeSystemEventAttributeResponseBody {
	s.SystemEvents = v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) Validate() error {
	if s.SystemEvents != nil {
		if err := s.SystemEvents.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSystemEventAttributeResponseBodySystemEvents struct {
	SystemEvent []*DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent `json:"SystemEvent,omitempty" xml:"SystemEvent,omitempty" type:"Repeated"`
}

func (s DescribeSystemEventAttributeResponseBodySystemEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventAttributeResponseBodySystemEvents) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeResponseBodySystemEvents) GetSystemEvent() []*DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	return s.SystemEvent
}

func (s *DescribeSystemEventAttributeResponseBodySystemEvents) SetSystemEvent(v []*DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) *DescribeSystemEventAttributeResponseBodySystemEvents {
	s.SystemEvent = v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEvents) Validate() error {
	if s.SystemEvent != nil {
		for _, item := range s.SystemEvent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent struct {
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Time         *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetContent() *string {
	return s.Content
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetId() *string {
	return s.Id
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetLevel() *string {
	return s.Level
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetName() *string {
	return s.Name
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetProduct() *string {
	return s.Product
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetStatus() *string {
	return s.Status
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GetTime() *int64 {
	return s.Time
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetContent(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Content = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetGroupId(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetId(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Id = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetInstanceName(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.InstanceName = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetLevel(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetName(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetProduct(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetRegionId(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.RegionId = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetResourceId(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.ResourceId = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetStatus(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetTime(v int64) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Time = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) Validate() error {
	return dara.Validate(s)
}
