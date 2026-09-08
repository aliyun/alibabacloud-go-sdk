// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *PutEventsRequest
	GetEventBusName() *string
	SetEventList(v []*PutEventsRequestEventList) *PutEventsRequest
	GetEventList() []*PutEventsRequestEventList
}

type PutEventsRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The list of events.
	EventList []*PutEventsRequestEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
}

func (s PutEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutEventsRequest) GoString() string {
	return s.String()
}

func (s *PutEventsRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *PutEventsRequest) GetEventList() []*PutEventsRequestEventList {
	return s.EventList
}

func (s *PutEventsRequest) SetEventBusName(v string) *PutEventsRequest {
	s.EventBusName = &v
	return s
}

func (s *PutEventsRequest) SetEventList(v []*PutEventsRequestEventList) *PutEventsRequest {
	s.EventList = v
	return s
}

func (s *PutEventsRequest) Validate() error {
	if s.EventList != nil {
		for _, item := range s.EventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutEventsRequestEventList struct {
	// The event payload.
	//
	// example:
	//
	// {"orderId": "1001", "amount": 99.9}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The data format.
	//
	// example:
	//
	// application/json
	DataContentType *string `json:"DataContentType,omitempty" xml:"DataContentType,omitempty"`
	// The data schema address.
	//
	// example:
	//
	// https://example.com/schema/v1
	DataSchema *string `json:"DataSchema,omitempty" xml:"DataSchema,omitempty"`
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a2c8f4e-0001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The event source.
	//
	// This parameter is required.
	//
	// example:
	//
	// my.custom.source
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The protocol version.
	//
	// example:
	//
	// 1.0
	SpecVersion *string `json:"SpecVersion,omitempty" xml:"SpecVersion,omitempty"`
	// The event subject.
	//
	// example:
	//
	// order-1001
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The event time.
	//
	// example:
	//
	// 2026-08-04T10:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The event type.
	//
	// This parameter is required.
	//
	// example:
	//
	// order:created
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PutEventsRequestEventList) String() string {
	return dara.Prettify(s)
}

func (s PutEventsRequestEventList) GoString() string {
	return s.String()
}

func (s *PutEventsRequestEventList) GetData() interface{} {
	return s.Data
}

func (s *PutEventsRequestEventList) GetDataContentType() *string {
	return s.DataContentType
}

func (s *PutEventsRequestEventList) GetDataSchema() *string {
	return s.DataSchema
}

func (s *PutEventsRequestEventList) GetId() *string {
	return s.Id
}

func (s *PutEventsRequestEventList) GetSource() *string {
	return s.Source
}

func (s *PutEventsRequestEventList) GetSpecVersion() *string {
	return s.SpecVersion
}

func (s *PutEventsRequestEventList) GetSubject() *string {
	return s.Subject
}

func (s *PutEventsRequestEventList) GetTime() *string {
	return s.Time
}

func (s *PutEventsRequestEventList) GetType() *string {
	return s.Type
}

func (s *PutEventsRequestEventList) SetData(v interface{}) *PutEventsRequestEventList {
	s.Data = v
	return s
}

func (s *PutEventsRequestEventList) SetDataContentType(v string) *PutEventsRequestEventList {
	s.DataContentType = &v
	return s
}

func (s *PutEventsRequestEventList) SetDataSchema(v string) *PutEventsRequestEventList {
	s.DataSchema = &v
	return s
}

func (s *PutEventsRequestEventList) SetId(v string) *PutEventsRequestEventList {
	s.Id = &v
	return s
}

func (s *PutEventsRequestEventList) SetSource(v string) *PutEventsRequestEventList {
	s.Source = &v
	return s
}

func (s *PutEventsRequestEventList) SetSpecVersion(v string) *PutEventsRequestEventList {
	s.SpecVersion = &v
	return s
}

func (s *PutEventsRequestEventList) SetSubject(v string) *PutEventsRequestEventList {
	s.Subject = &v
	return s
}

func (s *PutEventsRequestEventList) SetTime(v string) *PutEventsRequestEventList {
	s.Time = &v
	return s
}

func (s *PutEventsRequestEventList) SetType(v string) *PutEventsRequestEventList {
	s.Type = &v
	return s
}

func (s *PutEventsRequestEventList) Validate() error {
	return dara.Validate(s)
}
