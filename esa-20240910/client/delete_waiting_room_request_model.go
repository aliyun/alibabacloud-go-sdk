// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DeleteWaitingRoomRequest
	GetSiteId() *int64
	SetWaitingRoomId(v string) *DeleteWaitingRoomRequest
	GetWaitingRoomId() *string
}

type DeleteWaitingRoomRequest struct {
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s DeleteWaitingRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteWaitingRoomRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *DeleteWaitingRoomRequest) SetSiteId(v int64) *DeleteWaitingRoomRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWaitingRoomRequest) SetWaitingRoomId(v string) *DeleteWaitingRoomRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *DeleteWaitingRoomRequest) Validate() error {
	return dara.Validate(s)
}
