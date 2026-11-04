// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRecordPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventTypes(v string) *CreateEventRecordPlanRequest
	GetEventTypes() *string
	SetName(v string) *CreateEventRecordPlanRequest
	GetName() *string
	SetPreRecordDuration(v int32) *CreateEventRecordPlanRequest
	GetPreRecordDuration() *int32
	SetRecordDuration(v int32) *CreateEventRecordPlanRequest
	GetRecordDuration() *int32
	SetTemplateId(v string) *CreateEventRecordPlanRequest
	GetTemplateId() *string
}

type CreateEventRecordPlanRequest struct {
	// The types of intelligent alert events. Valid values:
	//
	// - 0: all intelligent alert events.
	//
	// - Multiple intelligent alert events: enter the sequence numbers of the intelligent alert events and separate them with commas (,). Example: `10001,10002`.
	//
	// > You can specify up to 10 intelligent alert events at a time.
	//
	// To view all types of intelligent alert events, log on to the IoT Platform console, go to the **Product Details*	- page, and click the **Define Feature*	- tab. For more information, see [View intelligent alert event types](https://help.aliyun.com/document_detail/207701.html).
	//
	// example:
	//
	// 1
	EventTypes *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	// The name of the recording plan.
	//
	// The name must be less than 48 characters in length and can contain letters, digits, and the following special characters: hyphens (-), underscores (_), at signs (@), periods (.), and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// plan1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pre-recording duration in seconds. Default value: 0. Maximum value: 10.
	//
	// example:
	//
	// 5
	PreRecordDuration *int32 `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	// The recording duration in seconds, including the pre-recording duration. Valid values: 10 to 60.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	RecordDuration *int32 `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	// The ID of the time template. You can call the [QueryTimeTemplate](https://help.aliyun.com/document_detail/107890.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fc82774fa749485bad7d719f9670****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateEventRecordPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRecordPlanRequest) GetEventTypes() *string {
	return s.EventTypes
}

func (s *CreateEventRecordPlanRequest) GetName() *string {
	return s.Name
}

func (s *CreateEventRecordPlanRequest) GetPreRecordDuration() *int32 {
	return s.PreRecordDuration
}

func (s *CreateEventRecordPlanRequest) GetRecordDuration() *int32 {
	return s.RecordDuration
}

func (s *CreateEventRecordPlanRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateEventRecordPlanRequest) SetEventTypes(v string) *CreateEventRecordPlanRequest {
	s.EventTypes = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetName(v string) *CreateEventRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetPreRecordDuration(v int32) *CreateEventRecordPlanRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetRecordDuration(v int32) *CreateEventRecordPlanRequest {
	s.RecordDuration = &v
	return s
}

func (s *CreateEventRecordPlanRequest) SetTemplateId(v string) *CreateEventRecordPlanRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateEventRecordPlanRequest) Validate() error {
	return dara.Validate(s)
}
