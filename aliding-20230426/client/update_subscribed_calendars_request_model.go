// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscribedCalendarsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *UpdateSubscribedCalendarsRequest
	GetCalendarId() *string
	SetDescription(v string) *UpdateSubscribedCalendarsRequest
	GetDescription() *string
	SetManagers(v []*string) *UpdateSubscribedCalendarsRequest
	GetManagers() []*string
	SetName(v string) *UpdateSubscribedCalendarsRequest
	GetName() *string
	SetSubscribeScope(v *UpdateSubscribedCalendarsRequestSubscribeScope) *UpdateSubscribedCalendarsRequest
	GetSubscribeScope() *UpdateSubscribedCalendarsRequestSubscribeScope
}

type UpdateSubscribedCalendarsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// M5MjkxNDUxQHVzZXJzLmRpbmd0YWxrLmNv
	CalendarId     *string                                         `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
	Description    *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Managers       []*string                                       `json:"Managers,omitempty" xml:"Managers,omitempty" type:"Repeated"`
	Name           *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	SubscribeScope *UpdateSubscribedCalendarsRequestSubscribeScope `json:"SubscribeScope,omitempty" xml:"SubscribeScope,omitempty" type:"Struct"`
}

func (s UpdateSubscribedCalendarsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscribedCalendarsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *UpdateSubscribedCalendarsRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSubscribedCalendarsRequest) GetManagers() []*string {
	return s.Managers
}

func (s *UpdateSubscribedCalendarsRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSubscribedCalendarsRequest) GetSubscribeScope() *UpdateSubscribedCalendarsRequestSubscribeScope {
	return s.SubscribeScope
}

func (s *UpdateSubscribedCalendarsRequest) SetCalendarId(v string) *UpdateSubscribedCalendarsRequest {
	s.CalendarId = &v
	return s
}

func (s *UpdateSubscribedCalendarsRequest) SetDescription(v string) *UpdateSubscribedCalendarsRequest {
	s.Description = &v
	return s
}

func (s *UpdateSubscribedCalendarsRequest) SetManagers(v []*string) *UpdateSubscribedCalendarsRequest {
	s.Managers = v
	return s
}

func (s *UpdateSubscribedCalendarsRequest) SetName(v string) *UpdateSubscribedCalendarsRequest {
	s.Name = &v
	return s
}

func (s *UpdateSubscribedCalendarsRequest) SetSubscribeScope(v *UpdateSubscribedCalendarsRequestSubscribeScope) *UpdateSubscribedCalendarsRequest {
	s.SubscribeScope = v
	return s
}

func (s *UpdateSubscribedCalendarsRequest) Validate() error {
	if s.SubscribeScope != nil {
		if err := s.SubscribeScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSubscribedCalendarsRequestSubscribeScope struct {
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s UpdateSubscribedCalendarsRequestSubscribeScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscribedCalendarsRequestSubscribeScope) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsRequestSubscribeScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *UpdateSubscribedCalendarsRequestSubscribeScope) SetUserIds(v []*string) *UpdateSubscribedCalendarsRequestSubscribeScope {
	s.UserIds = v
	return s
}

func (s *UpdateSubscribedCalendarsRequestSubscribeScope) Validate() error {
	return dara.Validate(s)
}
