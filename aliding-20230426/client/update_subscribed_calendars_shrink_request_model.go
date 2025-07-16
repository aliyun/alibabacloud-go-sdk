// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscribedCalendarsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *UpdateSubscribedCalendarsShrinkRequest
	GetCalendarId() *string
	SetDescription(v string) *UpdateSubscribedCalendarsShrinkRequest
	GetDescription() *string
	SetManagersShrink(v string) *UpdateSubscribedCalendarsShrinkRequest
	GetManagersShrink() *string
	SetName(v string) *UpdateSubscribedCalendarsShrinkRequest
	GetName() *string
	SetSubscribeScopeShrink(v string) *UpdateSubscribedCalendarsShrinkRequest
	GetSubscribeScopeShrink() *string
}

type UpdateSubscribedCalendarsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// M5MjkxNDUxQHVzZXJzLmRpbmd0YWxrLmNv
	CalendarId           *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagersShrink       *string `json:"Managers,omitempty" xml:"Managers,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SubscribeScopeShrink *string `json:"SubscribeScope,omitempty" xml:"SubscribeScope,omitempty"`
}

func (s UpdateSubscribedCalendarsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscribedCalendarsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsShrinkRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *UpdateSubscribedCalendarsShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSubscribedCalendarsShrinkRequest) GetManagersShrink() *string {
	return s.ManagersShrink
}

func (s *UpdateSubscribedCalendarsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSubscribedCalendarsShrinkRequest) GetSubscribeScopeShrink() *string {
	return s.SubscribeScopeShrink
}

func (s *UpdateSubscribedCalendarsShrinkRequest) SetCalendarId(v string) *UpdateSubscribedCalendarsShrinkRequest {
	s.CalendarId = &v
	return s
}

func (s *UpdateSubscribedCalendarsShrinkRequest) SetDescription(v string) *UpdateSubscribedCalendarsShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSubscribedCalendarsShrinkRequest) SetManagersShrink(v string) *UpdateSubscribedCalendarsShrinkRequest {
	s.ManagersShrink = &v
	return s
}

func (s *UpdateSubscribedCalendarsShrinkRequest) SetName(v string) *UpdateSubscribedCalendarsShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSubscribedCalendarsShrinkRequest) SetSubscribeScopeShrink(v string) *UpdateSubscribedCalendarsShrinkRequest {
	s.SubscribeScopeShrink = &v
	return s
}

func (s *UpdateSubscribedCalendarsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
