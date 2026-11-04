// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventRecordPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventTypes(v string) *UpdateEventRecordPlanRequest
	GetEventTypes() *string
	SetName(v string) *UpdateEventRecordPlanRequest
	GetName() *string
	SetPlanId(v string) *UpdateEventRecordPlanRequest
	GetPlanId() *string
	SetPreRecordDuration(v int32) *UpdateEventRecordPlanRequest
	GetPreRecordDuration() *int32
	SetRecordDuration(v int32) *UpdateEventRecordPlanRequest
	GetRecordDuration() *int32
	SetTemplateId(v string) *UpdateEventRecordPlanRequest
	GetTemplateId() *string
}

type UpdateEventRecordPlanRequest struct {
	// The types of intelligent alert events. Valid values:
	//
	// - 0: All intelligent alert events.
	//
	// - Multiple intelligent alert events: Enter the sequence numbers of the intelligent alert events and separate them with commas (,), for example, `10001,10002`.
	//
	// > You can specify up to 10 intelligent alert events at a time.
	//
	// To view all types of intelligent alert events, log on to the IoT Platform console, go to the **Product Details*	- page, and click the **Define Feature*	- tab. For more information, see [View intelligent alert event types](https://help.aliyun.com/document_detail/207701.html).
	//
	// example:
	//
	// 1
	EventTypes *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	// The name of the video recording plan for intelligent alert events.
	//
	// The name must be less than 48 characters in length and can contain letters, digits, and the following special characters: hyphens (-), underscores (_), at signs (@), periods (.), and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// record1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the video recording plan for intelligent alert events. You can call the [QueryEventRecordPlans](https://help.aliyun.com/document_detail/107907.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b60bf0dd55944c19d7817442bc9****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The pre-recording duration. Valid values: 0 to 10. Unit: seconds.
	//
	// > A pre-recording duration is generated only when the device supports the pre-recording feature. Otherwise, the pre-recording duration is 0.
	//
	// example:
	//
	// 10
	PreRecordDuration *int32 `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	// The video recording duration, including the pre-recording duration. Valid values: 10 to 60. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	RecordDuration *int32 `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
	// The ID of the time template. You can call the [QueryTimeTemplate](https://help.aliyun.com/document_detail/107890.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36f06322442e4e3f85a5ae8862c9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateEventRecordPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventRecordPlanRequest) GetEventTypes() *string {
	return s.EventTypes
}

func (s *UpdateEventRecordPlanRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEventRecordPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateEventRecordPlanRequest) GetPreRecordDuration() *int32 {
	return s.PreRecordDuration
}

func (s *UpdateEventRecordPlanRequest) GetRecordDuration() *int32 {
	return s.RecordDuration
}

func (s *UpdateEventRecordPlanRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateEventRecordPlanRequest) SetEventTypes(v string) *UpdateEventRecordPlanRequest {
	s.EventTypes = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetName(v string) *UpdateEventRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetPlanId(v string) *UpdateEventRecordPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetPreRecordDuration(v int32) *UpdateEventRecordPlanRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetRecordDuration(v int32) *UpdateEventRecordPlanRequest {
	s.RecordDuration = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) SetTemplateId(v string) *UpdateEventRecordPlanRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateEventRecordPlanRequest) Validate() error {
	return dara.Validate(s)
}
