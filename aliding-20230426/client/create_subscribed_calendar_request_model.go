// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscribedCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSubscribedCalendarRequest
	GetDescription() *string
	SetManagers(v []*string) *CreateSubscribedCalendarRequest
	GetManagers() []*string
	SetName(v string) *CreateSubscribedCalendarRequest
	GetName() *string
	SetSubscribeScope(v *CreateSubscribedCalendarRequestSubscribeScope) *CreateSubscribedCalendarRequest
	GetSubscribeScope() *CreateSubscribedCalendarRequestSubscribeScope
}

type CreateSubscribedCalendarRequest struct {
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Managers    []*string `json:"Managers,omitempty" xml:"Managers,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SubscribeScope *CreateSubscribedCalendarRequestSubscribeScope `json:"SubscribeScope,omitempty" xml:"SubscribeScope,omitempty" type:"Struct"`
}

func (s CreateSubscribedCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscribedCalendarRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSubscribedCalendarRequest) GetManagers() []*string {
	return s.Managers
}

func (s *CreateSubscribedCalendarRequest) GetName() *string {
	return s.Name
}

func (s *CreateSubscribedCalendarRequest) GetSubscribeScope() *CreateSubscribedCalendarRequestSubscribeScope {
	return s.SubscribeScope
}

func (s *CreateSubscribedCalendarRequest) SetDescription(v string) *CreateSubscribedCalendarRequest {
	s.Description = &v
	return s
}

func (s *CreateSubscribedCalendarRequest) SetManagers(v []*string) *CreateSubscribedCalendarRequest {
	s.Managers = v
	return s
}

func (s *CreateSubscribedCalendarRequest) SetName(v string) *CreateSubscribedCalendarRequest {
	s.Name = &v
	return s
}

func (s *CreateSubscribedCalendarRequest) SetSubscribeScope(v *CreateSubscribedCalendarRequestSubscribeScope) *CreateSubscribedCalendarRequest {
	s.SubscribeScope = v
	return s
}

func (s *CreateSubscribedCalendarRequest) Validate() error {
	return dara.Validate(s)
}

type CreateSubscribedCalendarRequestSubscribeScope struct {
	CorpIds             []*string `json:"CorpIds,omitempty" xml:"CorpIds,omitempty" type:"Repeated"`
	OpenConversationIds []*string `json:"OpenConversationIds,omitempty" xml:"OpenConversationIds,omitempty" type:"Repeated"`
	UserIds             []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CreateSubscribedCalendarRequestSubscribeScope) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscribedCalendarRequestSubscribeScope) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) GetCorpIds() []*string {
	return s.CorpIds
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) GetOpenConversationIds() []*string {
	return s.OpenConversationIds
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) SetCorpIds(v []*string) *CreateSubscribedCalendarRequestSubscribeScope {
	s.CorpIds = v
	return s
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) SetOpenConversationIds(v []*string) *CreateSubscribedCalendarRequestSubscribeScope {
	s.OpenConversationIds = v
	return s
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) SetUserIds(v []*string) *CreateSubscribedCalendarRequestSubscribeScope {
	s.UserIds = v
	return s
}

func (s *CreateSubscribedCalendarRequestSubscribeScope) Validate() error {
	return dara.Validate(s)
}
