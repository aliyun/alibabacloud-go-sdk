// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventInfo(v []*PutCustomEventRequestEventInfo) *PutCustomEventRequest
	GetEventInfo() []*PutCustomEventRequestEventInfo
	SetRegionId(v string) *PutCustomEventRequest
	GetRegionId() *string
}

type PutCustomEventRequest struct {
	// The event details.
	//
	// This parameter is required.
	EventInfo []*PutCustomEventRequestEventInfo `json:"EventInfo,omitempty" xml:"EventInfo,omitempty" type:"Repeated"`
	RegionId  *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutCustomEventRequest) String() string {
	return dara.Prettify(s)
}

func (s PutCustomEventRequest) GoString() string {
	return s.String()
}

func (s *PutCustomEventRequest) GetEventInfo() []*PutCustomEventRequestEventInfo {
	return s.EventInfo
}

func (s *PutCustomEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutCustomEventRequest) SetEventInfo(v []*PutCustomEventRequestEventInfo) *PutCustomEventRequest {
	s.EventInfo = v
	return s
}

func (s *PutCustomEventRequest) SetRegionId(v string) *PutCustomEventRequest {
	s.RegionId = &v
	return s
}

func (s *PutCustomEventRequest) Validate() error {
	if s.EventInfo != nil {
		for _, item := range s.EventInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutCustomEventRequestEventInfo struct {
	// The event content. Valid values of N: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// IOException
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The event name. Valid values of N: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// myEvent
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The ID of the application group. Valid values of N: 1 to 50.
	//
	// Default value: 0. This value indicates that the event to be reported does not belong to any application group.
	//
	// example:
	//
	// 123456
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the event occurred.
	//
	// Format: `yyyyMMddTHHmmss.SSSZ`.
	//
	// Valid values of N: 1 to 50.
	//
	// >  You can also specify a UNIX timestamp. Example: 1552199984000. Unit: milliseconds.
	//
	// example:
	//
	// 20171013T170923.456+0800
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s PutCustomEventRequestEventInfo) String() string {
	return dara.Prettify(s)
}

func (s PutCustomEventRequestEventInfo) GoString() string {
	return s.String()
}

func (s *PutCustomEventRequestEventInfo) GetContent() *string {
	return s.Content
}

func (s *PutCustomEventRequestEventInfo) GetEventName() *string {
	return s.EventName
}

func (s *PutCustomEventRequestEventInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *PutCustomEventRequestEventInfo) GetTime() *string {
	return s.Time
}

func (s *PutCustomEventRequestEventInfo) SetContent(v string) *PutCustomEventRequestEventInfo {
	s.Content = &v
	return s
}

func (s *PutCustomEventRequestEventInfo) SetEventName(v string) *PutCustomEventRequestEventInfo {
	s.EventName = &v
	return s
}

func (s *PutCustomEventRequestEventInfo) SetGroupId(v string) *PutCustomEventRequestEventInfo {
	s.GroupId = &v
	return s
}

func (s *PutCustomEventRequestEventInfo) SetTime(v string) *PutCustomEventRequestEventInfo {
	s.Time = &v
	return s
}

func (s *PutCustomEventRequestEventInfo) Validate() error {
	return dara.Validate(s)
}
