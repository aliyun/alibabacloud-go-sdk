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
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The waiting room ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 25133f536f1b1f6b6091f6a92c614dd4
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
