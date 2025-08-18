// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DeleteWaitingRoomEventRequest
	GetSiteId() *int64
	SetWaitingRoomEventId(v int64) *DeleteWaitingRoomEventRequest
	GetWaitingRoomEventId() *int64
}

type DeleteWaitingRoomEventRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The ID of the waiting room event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 302909890***
	WaitingRoomEventId *int64 `json:"WaitingRoomEventId,omitempty" xml:"WaitingRoomEventId,omitempty"`
}

func (s DeleteWaitingRoomEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomEventRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomEventRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteWaitingRoomEventRequest) GetWaitingRoomEventId() *int64 {
	return s.WaitingRoomEventId
}

func (s *DeleteWaitingRoomEventRequest) SetSiteId(v int64) *DeleteWaitingRoomEventRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWaitingRoomEventRequest) SetWaitingRoomEventId(v int64) *DeleteWaitingRoomEventRequest {
	s.WaitingRoomEventId = &v
	return s
}

func (s *DeleteWaitingRoomEventRequest) Validate() error {
	return dara.Validate(s)
}
