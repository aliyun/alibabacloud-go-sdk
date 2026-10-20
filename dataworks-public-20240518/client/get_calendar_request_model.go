// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetCalendarRequest
	GetId() *int64
}

type GetCalendarRequest struct {
	// The ID of the scheduling calendar.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCalendarRequest) GoString() string {
	return s.String()
}

func (s *GetCalendarRequest) GetId() *int64 {
	return s.Id
}

func (s *GetCalendarRequest) SetId(v int64) *GetCalendarRequest {
	s.Id = &v
	return s
}

func (s *GetCalendarRequest) Validate() error {
	return dara.Validate(s)
}
