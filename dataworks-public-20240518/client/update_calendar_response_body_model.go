// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCalendarResponseBody
	GetRequestId() *string
	SetVersion(v int32) *UpdateCalendarResponseBody
	GetVersion() *int32
}

type UpdateCalendarResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The updated version number.
	//
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCalendarResponseBody) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateCalendarResponseBody) SetRequestId(v string) *UpdateCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCalendarResponseBody) SetVersion(v int32) *UpdateCalendarResponseBody {
	s.Version = &v
	return s
}

func (s *UpdateCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
