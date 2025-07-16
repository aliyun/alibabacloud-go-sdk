// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscribedCalendarShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSubscribedCalendarShrinkRequest
	GetDescription() *string
	SetManagersShrink(v string) *CreateSubscribedCalendarShrinkRequest
	GetManagersShrink() *string
	SetName(v string) *CreateSubscribedCalendarShrinkRequest
	GetName() *string
	SetSubscribeScopeShrink(v string) *CreateSubscribedCalendarShrinkRequest
	GetSubscribeScopeShrink() *string
}

type CreateSubscribedCalendarShrinkRequest struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagersShrink *string `json:"Managers,omitempty" xml:"Managers,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SubscribeScopeShrink *string `json:"SubscribeScope,omitempty" xml:"SubscribeScope,omitempty"`
}

func (s CreateSubscribedCalendarShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscribedCalendarShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSubscribedCalendarShrinkRequest) GetManagersShrink() *string {
	return s.ManagersShrink
}

func (s *CreateSubscribedCalendarShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateSubscribedCalendarShrinkRequest) GetSubscribeScopeShrink() *string {
	return s.SubscribeScopeShrink
}

func (s *CreateSubscribedCalendarShrinkRequest) SetDescription(v string) *CreateSubscribedCalendarShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSubscribedCalendarShrinkRequest) SetManagersShrink(v string) *CreateSubscribedCalendarShrinkRequest {
	s.ManagersShrink = &v
	return s
}

func (s *CreateSubscribedCalendarShrinkRequest) SetName(v string) *CreateSubscribedCalendarShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSubscribedCalendarShrinkRequest) SetSubscribeScopeShrink(v string) *CreateSubscribedCalendarShrinkRequest {
	s.SubscribeScopeShrink = &v
	return s
}

func (s *CreateSubscribedCalendarShrinkRequest) Validate() error {
	return dara.Validate(s)
}
