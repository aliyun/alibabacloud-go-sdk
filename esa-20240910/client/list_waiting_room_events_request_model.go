// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *ListWaitingRoomEventsRequest
	GetSiteId() *int64
	SetWaitingRoomEventId(v int64) *ListWaitingRoomEventsRequest
	GetWaitingRoomEventId() *int64
	SetWaitingRoomId(v string) *ListWaitingRoomEventsRequest
	GetWaitingRoomId() *string
}

type ListWaitingRoomEventsRequest struct {
	// This parameter is required.
	SiteId             *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	WaitingRoomEventId *int64 `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
	// This parameter is required.
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s ListWaitingRoomEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomEventsRequest) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomEventsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWaitingRoomEventsRequest) GetWaitingRoomEventId() *int64 {
	return s.WaitingRoomEventId
}

func (s *ListWaitingRoomEventsRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *ListWaitingRoomEventsRequest) SetSiteId(v int64) *ListWaitingRoomEventsRequest {
	s.SiteId = &v
	return s
}

func (s *ListWaitingRoomEventsRequest) SetWaitingRoomEventId(v int64) *ListWaitingRoomEventsRequest {
	s.WaitingRoomEventId = &v
	return s
}

func (s *ListWaitingRoomEventsRequest) SetWaitingRoomId(v string) *ListWaitingRoomEventsRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *ListWaitingRoomEventsRequest) Validate() error {
	return dara.Validate(s)
}
