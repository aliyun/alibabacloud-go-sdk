// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListCalendarsResponseBody
	GetRequestId() *string
	SetResponse(v *ListCalendarsResponseBodyResponse) *ListCalendarsResponseBody
	GetResponse() *ListCalendarsResponseBodyResponse
}

type ListCalendarsResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Response  *ListCalendarsResponseBodyResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
}

func (s ListCalendarsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCalendarsResponseBody) GetResponse() *ListCalendarsResponseBodyResponse {
	return s.Response
}

func (s *ListCalendarsResponseBody) SetRequestId(v string) *ListCalendarsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalendarsResponseBody) SetResponse(v *ListCalendarsResponseBodyResponse) *ListCalendarsResponseBody {
	s.Response = v
	return s
}

func (s *ListCalendarsResponseBody) Validate() error {
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCalendarsResponseBodyResponse struct {
	Calendars []*ListCalendarsResponseBodyResponseCalendars `json:"Calendars,omitempty" xml:"Calendars,omitempty" type:"Repeated"`
}

func (s ListCalendarsResponseBodyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBodyResponse) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyResponse) GetCalendars() []*ListCalendarsResponseBodyResponseCalendars {
	return s.Calendars
}

func (s *ListCalendarsResponseBodyResponse) SetCalendars(v []*ListCalendarsResponseBodyResponseCalendars) *ListCalendarsResponseBodyResponse {
	s.Calendars = v
	return s
}

func (s *ListCalendarsResponseBodyResponse) Validate() error {
	if s.Calendars != nil {
		for _, item := range s.Calendars {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCalendarsResponseBodyResponseCalendars struct {
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 0
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// cnNTbWxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// VIEW_DETAIL
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
	// example:
	//
	// 标题
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// example:
	//
	// primary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCalendarsResponseBodyResponseCalendars) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBodyResponseCalendars) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyResponseCalendars) GetDescription() *string {
	return s.Description
}

func (s *ListCalendarsResponseBodyResponseCalendars) GetETag() *string {
	return s.ETag
}

func (s *ListCalendarsResponseBodyResponseCalendars) GetId() *string {
	return s.Id
}

func (s *ListCalendarsResponseBodyResponseCalendars) GetPrivilege() *string {
	return s.Privilege
}

func (s *ListCalendarsResponseBodyResponseCalendars) GetSummary() *string {
	return s.Summary
}

func (s *ListCalendarsResponseBodyResponseCalendars) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListCalendarsResponseBodyResponseCalendars) GetType() *string {
	return s.Type
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetDescription(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Description = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetETag(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.ETag = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetId(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Id = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetPrivilege(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Privilege = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetSummary(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Summary = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetTimeZone(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.TimeZone = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) SetType(v string) *ListCalendarsResponseBodyResponseCalendars {
	s.Type = &v
	return s
}

func (s *ListCalendarsResponseBodyResponseCalendars) Validate() error {
	return dara.Validate(s)
}
