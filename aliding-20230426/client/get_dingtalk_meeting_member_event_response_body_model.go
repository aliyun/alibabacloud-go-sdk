// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetDingtalkMeetingMemberEventResponseBodyData) *GetDingtalkMeetingMemberEventResponseBody
	GetData() []*GetDingtalkMeetingMemberEventResponseBodyData
	SetRequestId(v string) *GetDingtalkMeetingMemberEventResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetDingtalkMeetingMemberEventResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDingtalkMeetingMemberEventResponseBody
	GetVendorType() *string
}

type GetDingtalkMeetingMemberEventResponseBody struct {
	// example:
	//
	// []
	Data []*GetDingtalkMeetingMemberEventResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetDingtalkMeetingMemberEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberEventResponseBody) GetData() []*GetDingtalkMeetingMemberEventResponseBodyData {
	return s.Data
}

func (s *GetDingtalkMeetingMemberEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDingtalkMeetingMemberEventResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDingtalkMeetingMemberEventResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDingtalkMeetingMemberEventResponseBody) SetData(v []*GetDingtalkMeetingMemberEventResponseBodyData) *GetDingtalkMeetingMemberEventResponseBody {
	s.Data = v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBody) SetRequestId(v string) *GetDingtalkMeetingMemberEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBody) SetVendorRequestId(v string) *GetDingtalkMeetingMemberEventResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBody) SetVendorType(v string) *GetDingtalkMeetingMemberEventResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBody) Validate() error {
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

type GetDingtalkMeetingMemberEventResponseBodyData struct {
	ConfModule *string `json:"confModule,omitempty" xml:"confModule,omitempty"`
	// example:
	//
	// event123
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// join
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// example:
	//
	// member
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// example:
	//
	// 1638360000000
	Ts *int64 `json:"ts,omitempty" xml:"ts,omitempty"`
}

func (s GetDingtalkMeetingMemberEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) GetConfModule() *string {
	return s.ConfModule
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) GetEventId() *string {
	return s.EventId
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) GetEventType() *string {
	return s.EventType
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) GetTs() *int64 {
	return s.Ts
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) SetConfModule(v string) *GetDingtalkMeetingMemberEventResponseBodyData {
	s.ConfModule = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) SetEventId(v string) *GetDingtalkMeetingMemberEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) SetEventName(v string) *GetDingtalkMeetingMemberEventResponseBodyData {
	s.EventName = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) SetEventType(v string) *GetDingtalkMeetingMemberEventResponseBodyData {
	s.EventType = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) SetTs(v int64) *GetDingtalkMeetingMemberEventResponseBodyData {
	s.Ts = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}
