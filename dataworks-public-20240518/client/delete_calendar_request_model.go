// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteCalendarRequest
	GetId() *int64
}

type DeleteCalendarRequest struct {
	// The ID of the scheduling calendar.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCalendarRequest) GoString() string {
	return s.String()
}

func (s *DeleteCalendarRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteCalendarRequest) SetId(v int64) *DeleteCalendarRequest {
	s.Id = &v
	return s
}

func (s *DeleteCalendarRequest) Validate() error {
	return dara.Validate(s)
}
