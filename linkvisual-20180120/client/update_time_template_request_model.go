// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTimeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllDay(v int32) *UpdateTimeTemplateRequest
	GetAllDay() *int32
	SetName(v string) *UpdateTimeTemplateRequest
	GetName() *string
	SetTemplateId(v string) *UpdateTimeTemplateRequest
	GetTemplateId() *string
	SetTimeSections(v []*UpdateTimeTemplateRequestTimeSections) *UpdateTimeTemplateRequest
	GetTimeSections() []*UpdateTimeTemplateRequestTimeSections
}

type UpdateTimeTemplateRequest struct {
	// Specifies whether 24/7 recording is enabled.
	//
	// Valid values:
	//
	// - 1: yes.
	//
	// - 0: no. For detailed configurations, see the TimeSections.N.DayOfWeek, TimeSections.N.Begin, and TimeSections.N.End parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AllDay *int32 `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	// The name of the time template.
	//
	// The name must be less than 48 characters in length and can contain letters, digits, and the following special characters: hyphens (-), underscores (_), at signs (@), periods (.), and colons (:).
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the time template. You can call the [QueryTimeTemplate](https://help.aliyun.com/document_detail/107890.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fc82774fa749485bad7d719f9670****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The list of time sections.
	TimeSections []*UpdateTimeTemplateRequestTimeSections `json:"TimeSections,omitempty" xml:"TimeSections,omitempty" type:"Repeated"`
}

func (s UpdateTimeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateRequest) GetAllDay() *int32 {
	return s.AllDay
}

func (s *UpdateTimeTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTimeTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateTimeTemplateRequest) GetTimeSections() []*UpdateTimeTemplateRequestTimeSections {
	return s.TimeSections
}

func (s *UpdateTimeTemplateRequest) SetAllDay(v int32) *UpdateTimeTemplateRequest {
	s.AllDay = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetName(v string) *UpdateTimeTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetTemplateId(v string) *UpdateTimeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTimeTemplateRequest) SetTimeSections(v []*UpdateTimeTemplateRequestTimeSections) *UpdateTimeTemplateRequest {
	s.TimeSections = v
	return s
}

func (s *UpdateTimeTemplateRequest) Validate() error {
	if s.TimeSections != nil {
		for _, item := range s.TimeSections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTimeTemplateRequestTimeSections struct {
	// The start time of the day. Valid values: 0 to 86399. Unit: seconds.
	//
	// example:
	//
	// 20000
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// The day of the week. Valid values: 0 to 6, which indicate Sunday to Saturday.
	//
	// example:
	//
	// 0
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	// The end time of the day. Valid values: 1 to 86399. Unit: seconds.
	//
	// example:
	//
	// 60000
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s UpdateTimeTemplateRequestTimeSections) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimeTemplateRequestTimeSections) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateRequestTimeSections) GetBegin() *int32 {
	return s.Begin
}

func (s *UpdateTimeTemplateRequestTimeSections) GetDayOfWeek() *int32 {
	return s.DayOfWeek
}

func (s *UpdateTimeTemplateRequestTimeSections) GetEnd() *int32 {
	return s.End
}

func (s *UpdateTimeTemplateRequestTimeSections) SetBegin(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.Begin = &v
	return s
}

func (s *UpdateTimeTemplateRequestTimeSections) SetDayOfWeek(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.DayOfWeek = &v
	return s
}

func (s *UpdateTimeTemplateRequestTimeSections) SetEnd(v int32) *UpdateTimeTemplateRequestTimeSections {
	s.End = &v
	return s
}

func (s *UpdateTimeTemplateRequestTimeSections) Validate() error {
	return dara.Validate(s)
}
