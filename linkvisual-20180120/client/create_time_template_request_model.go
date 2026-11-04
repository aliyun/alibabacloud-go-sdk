// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTimeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllDay(v int32) *CreateTimeTemplateRequest
	GetAllDay() *int32
	SetName(v string) *CreateTimeTemplateRequest
	GetName() *string
	SetTimeSections(v []*CreateTimeTemplateRequestTimeSections) *CreateTimeTemplateRequest
	GetTimeSections() []*CreateTimeTemplateRequestTimeSections
}

type CreateTimeTemplateRequest struct {
	// Specifies whether to enable 24/7 recording. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No. You must configure the TimeSections.N.DayOfWeek, TimeSections.N.Begin, and TimeSections.N.End parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AllDay *int32 `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	// The name of the time template.
	//
	// The name must be fewer than 48 characters in length and can contain letters, digits, and the following special characters: hyphens (-), underscores (_), at signs (@), periods (.), and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of time ranges.
	TimeSections []*CreateTimeTemplateRequestTimeSections `json:"TimeSections,omitempty" xml:"TimeSections,omitempty" type:"Repeated"`
}

func (s CreateTimeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTimeTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateRequest) GetAllDay() *int32 {
	return s.AllDay
}

func (s *CreateTimeTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateTimeTemplateRequest) GetTimeSections() []*CreateTimeTemplateRequestTimeSections {
	return s.TimeSections
}

func (s *CreateTimeTemplateRequest) SetAllDay(v int32) *CreateTimeTemplateRequest {
	s.AllDay = &v
	return s
}

func (s *CreateTimeTemplateRequest) SetName(v string) *CreateTimeTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateTimeTemplateRequest) SetTimeSections(v []*CreateTimeTemplateRequestTimeSections) *CreateTimeTemplateRequest {
	s.TimeSections = v
	return s
}

func (s *CreateTimeTemplateRequest) Validate() error {
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

type CreateTimeTemplateRequestTimeSections struct {
	// The start time of each day. Valid values: 0 to 86399. Unit: seconds.
	//
	// example:
	//
	// 20000
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// The day of the week. Valid values: 0 to 6, which represent Sunday through Saturday.
	//
	// example:
	//
	// 0
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	// The end time of each day. Valid values: 1 to 86399. Unit: seconds.
	//
	// example:
	//
	// 60000
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s CreateTimeTemplateRequestTimeSections) String() string {
	return dara.Prettify(s)
}

func (s CreateTimeTemplateRequestTimeSections) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateRequestTimeSections) GetBegin() *int32 {
	return s.Begin
}

func (s *CreateTimeTemplateRequestTimeSections) GetDayOfWeek() *int32 {
	return s.DayOfWeek
}

func (s *CreateTimeTemplateRequestTimeSections) GetEnd() *int32 {
	return s.End
}

func (s *CreateTimeTemplateRequestTimeSections) SetBegin(v int32) *CreateTimeTemplateRequestTimeSections {
	s.Begin = &v
	return s
}

func (s *CreateTimeTemplateRequestTimeSections) SetDayOfWeek(v int32) *CreateTimeTemplateRequestTimeSections {
	s.DayOfWeek = &v
	return s
}

func (s *CreateTimeTemplateRequestTimeSections) SetEnd(v int32) *CreateTimeTemplateRequestTimeSections {
	s.End = &v
	return s
}

func (s *CreateTimeTemplateRequestTimeSections) Validate() error {
	return dara.Validate(s)
}
