// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscribedCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthor(v string) *GetSubscribedCalendarResponseBody
	GetAuthor() *string
	SetCalendarId(v string) *GetSubscribedCalendarResponseBody
	GetCalendarId() *string
	SetDescription(v string) *GetSubscribedCalendarResponseBody
	GetDescription() *string
	SetManagers(v []*string) *GetSubscribedCalendarResponseBody
	GetManagers() []*string
	SetName(v string) *GetSubscribedCalendarResponseBody
	GetName() *string
	SetRequestId(v string) *GetSubscribedCalendarResponseBody
	GetRequestId() *string
	SetSubscribeScope(v *GetSubscribedCalendarResponseBodySubscribeScope) *GetSubscribedCalendarResponseBody
	GetSubscribeScope() *GetSubscribedCalendarResponseBodySubscribeScope
}

type GetSubscribedCalendarResponseBody struct {
	// example:
	//
	// 012345
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// example:
	//
	// M5MjkxNDUxQHVzZXJzLmRpbmd0YWxrLmNv
	CalendarId *string `json:"calendarId,omitempty" xml:"calendarId,omitempty"`
	// example:
	//
	// 中国传统日历
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	Managers    []*string `json:"managers,omitempty" xml:"managers,omitempty" type:"Repeated"`
	// example:
	//
	// 中国传统日历
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId      *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SubscribeScope *GetSubscribedCalendarResponseBodySubscribeScope `json:"subscribeScope,omitempty" xml:"subscribeScope,omitempty" type:"Struct"`
}

func (s GetSubscribedCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscribedCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarResponseBody) GetAuthor() *string {
	return s.Author
}

func (s *GetSubscribedCalendarResponseBody) GetCalendarId() *string {
	return s.CalendarId
}

func (s *GetSubscribedCalendarResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetSubscribedCalendarResponseBody) GetManagers() []*string {
	return s.Managers
}

func (s *GetSubscribedCalendarResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSubscribedCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubscribedCalendarResponseBody) GetSubscribeScope() *GetSubscribedCalendarResponseBodySubscribeScope {
	return s.SubscribeScope
}

func (s *GetSubscribedCalendarResponseBody) SetAuthor(v string) *GetSubscribedCalendarResponseBody {
	s.Author = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetCalendarId(v string) *GetSubscribedCalendarResponseBody {
	s.CalendarId = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetDescription(v string) *GetSubscribedCalendarResponseBody {
	s.Description = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetManagers(v []*string) *GetSubscribedCalendarResponseBody {
	s.Managers = v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetName(v string) *GetSubscribedCalendarResponseBody {
	s.Name = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetRequestId(v string) *GetSubscribedCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscribedCalendarResponseBody) SetSubscribeScope(v *GetSubscribedCalendarResponseBodySubscribeScope) *GetSubscribedCalendarResponseBody {
	s.SubscribeScope = v
	return s
}

func (s *GetSubscribedCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSubscribedCalendarResponseBodySubscribeScope struct {
	CorpIds             []*string `json:"CorpIds,omitempty" xml:"CorpIds,omitempty" type:"Repeated"`
	OpenConversationIds []*string `json:"OpenConversationIds,omitempty" xml:"OpenConversationIds,omitempty" type:"Repeated"`
	UserIds             []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetSubscribedCalendarResponseBodySubscribeScope) String() string {
	return dara.Prettify(s)
}

func (s GetSubscribedCalendarResponseBodySubscribeScope) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) GetCorpIds() []*string {
	return s.CorpIds
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) GetOpenConversationIds() []*string {
	return s.OpenConversationIds
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) SetCorpIds(v []*string) *GetSubscribedCalendarResponseBodySubscribeScope {
	s.CorpIds = v
	return s
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) SetOpenConversationIds(v []*string) *GetSubscribedCalendarResponseBodySubscribeScope {
	s.OpenConversationIds = v
	return s
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) SetUserIds(v []*string) *GetSubscribedCalendarResponseBodySubscribeScope {
	s.UserIds = v
	return s
}

func (s *GetSubscribedCalendarResponseBodySubscribeScope) Validate() error {
	return dara.Validate(s)
}
