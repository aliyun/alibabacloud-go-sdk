// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *ListWaitingRoomsRequest
	GetSiteId() *int64
	SetWaitingRoomId(v string) *ListWaitingRoomsRequest
	GetWaitingRoomId() *string
}

type ListWaitingRoomsRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120876698010528
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The ID of the waiting room. Specify this parameter to query the information about a specific waiting room.
	//
	// example:
	//
	// 6a51d5bc6460887abd1291dc7d4d****
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s ListWaitingRoomsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWaitingRoomsRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *ListWaitingRoomsRequest) SetSiteId(v int64) *ListWaitingRoomsRequest {
	s.SiteId = &v
	return s
}

func (s *ListWaitingRoomsRequest) SetWaitingRoomId(v string) *ListWaitingRoomsRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *ListWaitingRoomsRequest) Validate() error {
	return dara.Validate(s)
}
